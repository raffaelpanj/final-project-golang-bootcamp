package connection

import (
	"database/sql"
	"final-project-golang-bootcamp/models"
	"fmt"

	_ "github.com/lib/pq"
	"time"

	// "os"
	"log"
	// "final-project-golang-bootcamp/models"
)
const (
	host	= "localhost"
	port 	= 5432
	user	= "postgres"
	password= "2010512058"
	dbname	= "db-events-sql"
)

var(
	Db *sql.DB
	err error
	logger = log.Default()
)

func ConnectDB() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	Db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	err = Db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the database!")
}

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
			panic(err)
		}
		results = append(results, user)
	}
	return results, nil
}

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

func UpdateEventById(id string, nama string, location string, date string, quota int, description string) (int64, error){
	layout := "2006-01-02"
	dateTime, err := time.Parse(layout, date)
	if err != nil {
		return 0, err
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


func InsertOrder(userId int, eventId int, ticketCount int, paymentMethod string, totalPrice int, eventCode string) (string, error) {
	tx, err := Db.Begin()
	var orderId string

	if err != nil {
		return "", err
	}
	defer tx.Rollback()

	// Insert to Order table
	sqlInsertOrder := `
	INSERT INTO orders (user_id, event_id, total_price, payment_method)
	VALUES ($1, $2, $3, $4) RETURNING order_id`
	err = tx.QueryRow(sqlInsertOrder, userId, eventId, totalPrice, paymentMethod).Scan(&orderId)
	if err != nil {
		return "", err
	}

	//Insert to Tickets table
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

	// Update Events Quota
	sqlUpdateEventQuota := `
	UPDATE events 
	SET quota = quota - $1 
	WHERE event_id = $2 AND quota >= $1`
	_, err = tx.Exec(sqlUpdateEventQuota, ticketCount, eventId)
	if err != nil {
		return "", err
	}

	err = tx.Commit()
	if err != nil {
		return "", err
	}
	return orderId, nil
}

func GetEventCode(eventId int) (string, error) {
	var eventCode string
	sqlStatement := `SELECT event_code FROM events WHERE event_id=$1`
	err := Db.QueryRow(sqlStatement, eventId).Scan(&eventCode)
	if err != nil {
		return "", err
	}
	return eventCode, nil
}
func InsertQueue(userId int, eventId int, status string) (models.Queue, error) {
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