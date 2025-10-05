# Final-Project-Golang-Bootcamp
##Event Ticketing API — Final Project (Golang Bootcamp)

This project is a **RESTful API** built with **Golang** and **Gin framework**, designed to handle an event ticketing flow from event creation to ticket generation.  
It demonstrates a full **end-to-end flow** involving event management, queueing, ordering, and ticket retrieval.

## 🧩 Tech Stack

- **Language:** Go (Golang)
- **Framework:** Gin
- **Database:** PostgreSQL
- **Authentication:** JWT (JSON Web Token)
- **Environment Management:** godotenv


---

## 🔑 Environment Variables

Create a `.env` file in the root directory:

```env
PORT=8080
DB_USER=yourusername
DB_PASS=yourpassword
DB_NAME=yourdbname
DB_HOST=localhost
DB_PORT=5432
JWT_SECRET_KEY=your_secret_key

# 1. Install dependencies
go mod tidy

# 2. Run the application
go run main.go

# 3. Server will start on
http://localhost:8080

🧭 Full API Endpoint List


🔄 End-to-End Flow Overview
Step	Action	Endpoint	Description
1️⃣	Create Event	POST /events	Admin creates a new event
2️⃣	Create Queue	POST /queues	User joins an event queue
3️⃣	Create Order	POST /orders	User makes an order after being served
4️⃣	Update Queue	PUT /queues/:QueueID	Admin updates queue status to served
5️⃣	Get Ticket	GET /tickets?order_id=<order_id>	User retrieves generated tickets
6️⃣	Get Order	GET /orders?user_id=<user_id>	User retrieves all their orders

🧩 Example Flow (E2E)
1️⃣ Create Event

Endpoint: POST /events

{
  "event_code": "EVT123",
  "name": "Music Festival 2025",
  "location": "Jakarta Convention Center",
  "date": "2025-12-01",
  "quota": 100,
  "description": "A fun music experience!"
}


✅ Response:

{
  "message": "Event created successfully",
  "event": { ... }
}

2️⃣ Create Queue

Endpoint: POST /queues

{
  "user_id": "USR001",
  "event_id": "EVT123",
  "status": "waiting"
}


✅ Response:

{
  "message": "Queue created successfully",
  "Queue Data": { ... }
}

3️⃣ Create Order

Endpoint: POST /orders

{
  "user_id": "USR001",
  "event_id": "EVT123",
  "ticket_count": 2,
  "payment_method": "credit_card",
  "total_price": 500000
}


✅ Response:

{
  "message": "Order created successfully",
  "order_id": "ORD001"
}

4️⃣ Update Queue to Served

Endpoint: PUT /queues/:QueueID

{
  "status": "served"
}


✅ Response:

{
  "message": "Queue updated successfully",
  "queue": { ... }
}

5️⃣ Get Ticket

Endpoint: GET /tickets?order_id=ORD001
✅ Response:

{
  "tickets": [
    {
      "ticket_id": "TCK001",
      "order_id": "ORD001",
      "ticket_number": "T12345",
      "price": 250000
    }
  ]
}

6️⃣ Get Orders by User ID

Endpoint: GET /orders?user_id=USR001
✅ Response:

[
  {
    "order_id": "ORD001",
    "event_id": "EVT123",
    "ticket_count": 2,
    "total_price": 500000,
    "payment_method": "credit_card"
  }
]

🧠 Key Features

✅ CRUD operations for events, queues, orders, and tickets
✅ Custom validation for dates (e.g., YYYY-MM-DD format)
✅ Quota and queue management
✅ JWT authentication with role-based authorization
✅ Detailed error handling and logging
✅ Easy deployment setup for Railway / Render
✅ Clean controller separation and modular design

🧪 Example Use Case

Admin creates a new event (POST /events)
User joins the event queue (POST /queues)
Admin updates queue to served (PUT /queues/:id)
User creates an order (POST /orders)
System automatically generates tickets
User retrieves their tickets (GET /tickets) and orders (GET /orders)