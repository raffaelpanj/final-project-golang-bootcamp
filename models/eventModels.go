package models

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	CreatedAt string
}

type Event struct {
	EventID		string `json:"event_id"`
	Name	 	string `json:"name" binding:"required"`
	Location	string `json:"location" binding:"required"`
	Quota	 	int `json:"quota" binding:"required"`
	Date	 	string `json:"date" binding:"required"`
	EventCode	 string `json:"event_code" binding:"required"`
	Description	 string `json:"description" binding:"required"`
	CreatedAt   string `json:"created_at"`
}

type UpdateEvent struct {
	EventID		string `json:"event_id"`
	Name	 	string `json:"name" binding:"required"`
	Location	string `json:"location" binding:"required"`
	Quota	 	int `json:"quota" binding:"required"`
	Date	 	string `json:"date" binding:"required"`
	EventCode	string `json:"event_code"`
	Description	string `json:"description" binding:"required"`
	CreatedAt   string `json:"created_at"`
}

type Queue struct {
	QueueID		int `json:"queue_id"`
	QueueNumber int `json:"queue_number"`
	EventID 	int	`json:"event_id" binding:"required"`
	UserID  	int	`json:"user_id" binding:"required"`
	Status 		string `json:"status" binding:"required"`
}

type Order struct {
	OrderID		int `json:"order_id"`
	EventID 	int	`json:"event_id"`
	UserID  	int	`json:"user_id"`
	TotalPrice	int `json:"total_price"`
	PaymentMethod string `json:"payment_method"`
}

type Ticket struct {
	OrderID		int `json:"order"`
	TicketNumber int `json:"ticket_number"`
	Price	int `json:"price"`
}

type OrderRequest struct {
    UserID   int     `json:"user_id" binding:"required"`
    EventID  int     `json:"event_id" binding:"required"`
    TicketCount int  `json:"ticket_count" binding:"required"`
	PaymentMethod string  `json:"payment_method" binding:"required"`
	TotalPrice int  `json:"total_price" binding:"required"`
}