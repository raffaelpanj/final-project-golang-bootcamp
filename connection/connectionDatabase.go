package connection

import (
	"database/sql"
	"final-project-golang-bootcamp/models"
	"fmt"
	"log"
	"time"
	"os"
	_ "github.com/lib/pq"
)

var(
	Db *sql.DB
	err error
)

// ConnectDB establishes a connection to the PostgreSQL database using environment variables.
func ConnectDB() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// Take from railway
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		// fallback to local if env not set (so it can run locally as well)
		dsn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname)
	}

	Db, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	err = Db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the database!")
}

// SelectAllUsersByRole retrieves all users with a specific role.
func SelectAllUsersByRole(role string)([]models.User, error){
	var results []models.User
	sqlStatement := `SELECT * FROM users WHERE role=$1`
	rows, err := Db.Query(sqlStatement, role)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next(){
		var user = models.User{}
		err = rows.Scan(&user.ID,&user.Name,&user.Email,&user.Password,&user.CreatedAt,&user.Role)
		if err != nil {
			return nil, err
		}
		results = append(results, user)
	}
	return results, nil
}

// InsertUser adds a new user to the users table and returns the generated user_id.
func InsertUser(name string, email string, password string, role string) (string, error){
	var user_id string
	sqlStatement := `
		INSERT INTO users (nama, email, password, role)
		VALUES ($1, $2, $3, $4)
		RETURNING user_id;
	`
	err = Db.QueryRow(sqlStatement, name, email, password, role).Scan(&user_id)
	if err != nil {
		return "", err
	}
	return user_id, nil
}

// SelectUser retrieves a single user by email.
func SelectUser(email string)(models.User, error){
	var user models.User
	sqlStatement := `SELECT * FROM users WHERE email=$1`
	err := Db.QueryRow(sqlStatement, email).Scan(&user.ID,&user.Name,&user.Email,&user.Password,&user.CreatedAt,&user.Role)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

// InsertEvent inserts a new event into the events table.
func InsertEvent(eventCode string, nama string, lokasi string, date string, quota int, description string) (string, error) {
	var event_id string
	layout := "2006-01-02"
	dateTime, err := time.Parse(layout, date)
	if err != nil {
		return "", err
	}
	sqlStatement := `
	INSERT INTO events (event_code, nama, location, date, quota, description)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING event_id`
	err = Db.QueryRow(sqlStatement, eventCode, nama, lokasi, dateTime, quota, description).Scan(&event_id)
	if err != nil {
		return "", err
	}
	return event_id, nil
}

// SelectEventById retrieves a single event by its event_id.
func SelectEventById(id string) (models.Event, error){
	var event = models.Event{}
	var createdAt time.Time
	var eventTime time.Time
	sqlStatement := `SELECT * FROM events WHERE event_id=$1`
	err := Db.QueryRow(sqlStatement, id).Scan(&event.EventID, &event.EventCode, &event.Name, &event.Location, &eventTime, &event.Quota, &event.Description, &createdAt)
		if err != nil {
		fmt.Printf("Error: %v\n", err)
		return models.Event{}, err
	}
	event.Date = eventTime.Format("2006-01-02")
	event.CreatedAt = createdAt.Format("2006-01-02 15:04:05")
	return event, nil
}

// UpdateEventById updates event details by event_id.
func UpdateEventById(id string, nama string, location string, date string, quota int, description string) (int64, error){
	layout := "2006-01-02"
	dateTime, err := time.Parse(layout, date)
	if err != nil {
		return 0, err
	}
	sqlStatmentCheckQuota := `SELECT quota FROM events WHERE event_id=$1`
	var currentQuota int
	err = Db.QueryRow(sqlStatmentCheckQuota, id).Scan(&currentQuota)
	if err != nil || currentQuota > quota {
		return 99, err
	}

	sqlStatement := `
	UPDATE events SET nama = $1, location = $2, date = $3,
	quota = $4, description = $5
	WHERE event_id = $6`

	result, err := Db.Exec(sqlStatement, nama, location, dateTime, quota, description, id)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil{
		return 0, err
	}
	return rowsAffected, nil
}

// InsertOrder creates an order and corresponding tickets in a transaction.
func InsertOrder(userId int, eventId int, ticketCount int, paymentMethod string, totalPrice int, eventCode string) (string, error) {
	tx, err := Db.Begin()
	var orderId string

	if err != nil {
		return "", err
	}
	defer tx.Rollback()

	// Insert order data
	sqlInsertOrder := `
	INSERT INTO orders (user_id, event_id, total_price, payment_method)
	VALUES ($1, $2, $3, $4) RETURNING order_id`
	err = tx.QueryRow(sqlInsertOrder, userId, eventId, totalPrice, paymentMethod).Scan(&orderId)
	if err != nil {
		return "", err
	}

	// Check last ticket number for the event
	var lastTicket string
	sqlCheckLastTicketNumber := `
	SELECT ticket_number FROM tickets 
	WHERE ticket_number LIKE $1 || '%' 
	ORDER BY ticket_id DESC LIMIT 1
	`

	err = tx.QueryRow(sqlCheckLastTicketNumber, eventCode).Scan(&lastTicket)
	if err != nil {
		if err == sql.ErrNoRows {
			lastTicket = ""
		} else {
			return "", err
		}
	}
	lastNumber := 0
	if lastTicket != "" {
		fmt.Sscanf(lastTicket, eventCode+"%d", &lastNumber)
	}
	
	// Insert tickets for the order
	sqlInsertTickets := `
	INSERT INTO tickets (order_id, ticket_number, price)
	VALUES ($1, $2, $3)`

	price := totalPrice / ticketCount
	remainder := totalPrice % ticketCount
	for i := 1; i <= ticketCount; i++ {
		lastNumber++
		ticketNumber := fmt.Sprintf("%s%d", eventCode, lastNumber)
		if(remainder > 0 && i == ticketCount){
			price += remainder
		}
		_, err = tx.Exec(sqlInsertTickets, orderId, ticketNumber, price)
		if err != nil {
			return "", err
		}
	}

	// Update remaining event quota
	sqlUpdateEventQuota := `
	UPDATE events 
	SET quota = quota - $1 
	WHERE event_id = $2 AND quota >= $1`
	res, err := tx.Exec(sqlUpdateEventQuota, ticketCount, eventId)
	if err != nil {
		return "", err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return "", err
	}

	if rowsAffected == 0 {
		return "", fmt.Errorf("not enough quota")
	}

	err = tx.Commit()
	if err != nil {
		return "", err
	}
	return orderId, nil
}

// SelectOrdersByUserId retrieves all orders for a specific user.
func SelectOrdersByUserId(userId string) ([]models.Order, error){
	var results []models.Order
	sqlStatement := `SELECT order_id, event_id, user_id, total_price, payment_method FROM orders WHERE user_id = $1`
	rows, err := Db.Query(sqlStatement, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next(){
		var order = models.Order{}
		err = rows.Scan(&order.OrderID, &order.EventID, &order.UserID, &order.TotalPrice, &order.PaymentMethod)
		if err != nil {
			panic(err)
		}
		results = append(results, order)
	}
	return results, nil
}

// SelectEventCodeById retrieves an event_code by event_id.
func SelectEventCodeById(eventId int) (string, error) {
	var eventCode string
	sqlStatement := `SELECT event_code FROM events WHERE event_id=$1`
	err := Db.QueryRow(sqlStatement, eventId).Scan(&eventCode)
	if err != nil {
		return "", err
	}
	return eventCode, nil
}

// InsertQueue inserts a new queue entry for an event.
func InsertQueue(userId string, eventId string, status string) (models.Queue, error) {
	var queue models.Queue
	var lastNumber = 0
	sqlCheckLastNumber := `SELECT queue_number FROM queues WHERE event_id = $1 ORDER BY queue_number DESC LIMIT 1`
	err := Db.QueryRow(sqlCheckLastNumber, eventId).Scan(&lastNumber)
	if err != nil && err != sql.ErrNoRows {
		return queue, err
	}
	lastNumber++

	sqlInsertOrder := `INSERT INTO queues (user_id, event_id, queue_number, status) 
	VALUES ($1, $2, $3, $4) RETURNING queue_id, user_id, queue_number, event_id, status`
	err = Db.QueryRow(sqlInsertOrder, userId, eventId, lastNumber, status).Scan(&queue.QueueID, &queue.UserID, &queue.QueueNumber, &queue.EventID, &queue.Status)
	if err != nil {
		return queue, err
	}
	return queue, nil
}

// UpdateQueueById updates the status of a queue by its ID.
func UpdateQueueById(id string, status string) (int64, error){
	sqlStatement := `
	UPDATE queues SET status = $1
	WHERE queue_id = $2`
	result, err := Db.Exec(sqlStatement, status, id)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}

// GetQueueById retrieves queue data by queue_id.
func GetQueueById(id string) (models.Queue, error) {
	var queue models.Queue
	sqlStatement := `SELECT queue_id, user_id, queue_number, event_id, status FROM queues WHERE queue_id = $1`
	err := Db.QueryRow(sqlStatement, id).Scan(&queue.QueueID, &queue.UserID, &queue.QueueNumber, &queue.EventID, &queue.Status)
	if err != nil {
		return queue, err
	}
	return queue, nil
}

// GetQueueByEventIdAndStatus retrieves queues for an event filtered by status.
func GetQueueByEventIdAndStatus(eventId string, status string) ([]models.Queue, error){
	var results []models.Queue

	sqlStatement := `SELECT queue_id, user_id, queue_number, event_id, status FROM queues WHERE event_id = $1 and status = $2`
	rows, err := Db.Query(sqlStatement, eventId, status)

		if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next(){
		var queue = models.Queue{}
		err = rows.Scan(&queue.QueueID, &queue.UserID, &queue.QueueNumber, &queue.EventID, &queue.Status)
		if err != nil {
			panic(err)
		}
		results = append(results, queue)
	}
	return results, nil
}

// SelectTicketByOrderId retrieves tickets related to an order.
func SelectTicketByOrderId(orderId string) ([]models.Ticket, error){
	var results []models.Ticket
	sqlStatement := `SELECT ticket_id, order_id, ticket_number, price FROM tickets WHERE order_id = $1`
	rows, err := Db.Query(sqlStatement, orderId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next(){
		var ticket = models.Ticket{}
		err = rows.Scan(&ticket.TicketID, &ticket.OrderID, &ticket.TicketNumber, &ticket.Price)
		if err != nil {
			panic(err)
		}
		results = append(results, ticket)
	}
	return results, nil
}

// SelectTicketById retrieves a ticket by its ticket_id.
func SelectTicketById(ticketId string) (models.Ticket, error){
	var ticket = models.Ticket{}
	sqlStatement := `SELECT ticket_id, order_id, ticket_number, price FROM tickets WHERE ticket_id=$1`
	err := Db.QueryRow(sqlStatement, ticketId).Scan(&ticket.TicketID, &ticket.OrderID, &ticket.TicketNumber, &ticket.Price)
	if err != nil {
		return ticket, err
	}
	return ticket, nil
}