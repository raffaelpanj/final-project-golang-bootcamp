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
```
## Full API Endpoint List
🧑‍💼 AuthController
Handles registration and login for both customers and admins.

Method	Endpoint	Description	Auth
POST	/register/customer	Register a new customer	❌ Public
POST	/register/admin	Register a new admin	❌ Public
POST	/login	Log in and get JWT token	❌ Public

🎫 EventController
Manage event data (CRUD operations).

Method	Endpoint	Description	Auth
POST	/events	Create a new event	🔒 Admin
GET	/events/:EventID	Get event by ID	✅ Authenticated
PUT	/events/:EventID	Update event by ID	🔒 Admin

⏳ QueueController
Handles user queues for events.

Method	Endpoint	Description	Auth
POST	/queues	Create a new queue for an event	✅ Authenticated
PUT	/queues/:QueueID	Update queue status (e.g., waiting → served)	🔒 Admin
GET	/queues/:QueueID	Get queue by queue ID	✅ Authenticated
GET	/queues?event_id={id}&status={status}	Get all queues filtered by event ID and status	🔒 Admin

🛒 OrderController
Handles user orders and ticket generation.

Method	Endpoint	Description	Auth
POST	/orders	Create a new order (generate tickets)	✅ Authenticated
GET	/orders?user_id={id}	Get all orders by user ID	✅ Authenticated

🎟️ TicketController
Retrieve ticket information by order or ticket ID.

Method	Endpoint	Description	Auth
GET	/tickets?order_id={id}	Get tickets by order ID	✅ Authenticated
GET	/tickets/:TicketID	Get ticket by ticket ID	✅ Authenticated

👥 UserController
Retrieve users by role (for admin management).

Method	Endpoint	Description	Auth
GET	/users/role/:UserRole	Get all users with a specific role (admin / customer)	🔒 Admin


## 🔄 End-to-End Flow Overview
Step	Action	Endpoint	Description
1️⃣	Register	POST /register/customer	User registers a new account
2️⃣	Login	POST /login	User logs in and gets a JWT token
3️⃣	Create Event	POST /events	Admin creates a new event
4️⃣	Create Queue	POST /queues	User joins an event queue
5️⃣	Create Order	POST /orders	User makes an order after being served
6️⃣	Update Queue	PUT /queues/:QueueID	Admin updates queue status to served
7️⃣	Get Ticket	GET /tickets?order_id=<order_id>	User retrieves generated tickets
8️⃣	Get Order	GET /orders?user_id=<user_id>	User retrieves all their orders


🧩 Example Flow (E2E)
1️⃣ Register New Customer

Endpoint: POST /register/customer
{
  "name": "Raffael",
  "email": "raffael@example.com",
  "password": "password123"
}


✅ Response:

{
  "message": "Customer registered successfully",
  "user_id": "USR001"
}

2️⃣ Login

Endpoint: POST /login

{
  "email": "raffael@example.com",
  "password": "password123"
}


✅ Response:

{
  "message": "Login successful",
  "token": "<JWT_TOKEN>"
}

3️⃣ Create Event (Admin Only)

Endpoint: POST /events
Header: Authorization: Bearer <JWT_TOKEN>

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

4️⃣ Create Queue

Endpoint: POST /queues
Header: Authorization: Bearer <JWT_TOKEN>

{
  "user_id": "USR001",
  "event_id": "EVT123",
  "status": "waiting"
}


✅ Response:

{
  "message": "Queue created successfully",
  "queue_data": { ... }
}

5️⃣ Create Order

Endpoint: POST /orders
Header: Authorization: Bearer <JWT_TOKEN>

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

6️⃣ Update Queue to Served (Admin Only)

Endpoint: PUT /queues/:QueueID
Header: Authorization: Bearer <JWT_TOKEN>

{
  "status": "served"
}


✅ Response:

{
  "message": "Queue updated successfully",
  "queue": { ... }
}

7️⃣ Get Ticket by Order ID

Endpoint: GET /tickets?order_id=ORD001
Header: Authorization: Bearer <JWT_TOKEN>
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

8️⃣ Get Orders by User ID

Endpoint: GET /orders?user_id=USR001
Header: Authorization: Bearer <JWT_TOKEN>
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

## 🧠 Key Features

✅ CRUD operations for events, queues, orders, and tickets
✅ Custom validation for dates (e.g., YYYY-MM-DD format)
✅ Quota and queue management
✅ JWT authentication with role-based authorization
✅ Detailed error handling and logging
✅ Easy deployment setup for Railway / Render
✅ Clean controller separation and modular design