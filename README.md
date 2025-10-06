# Final-Project-Golang-Bootcamp
##Event Ticketing API — Final Project (Golang Bootcamp)

This project is a **RESTful API** built with **Golang** and **Gin framework**, designed to handle an event ticketing flow from event creation to ticket generation. It demonstrates a full **end-to-end flow** involving event management, queueing, ordering, and ticket retrieval.

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

| Method   | Endpoint             | Description              | Auth     |
| :------- | :------------------- | :----------------------- | :------- |
| **POST** | `/users/registerCustomer` | Register a new customer  | ❌ Public |
| **POST** | `/users/registerAdmin`    | Register a new admin     | ❌ Public |
| **POST** | `/users/login`             | Log in and get JWT token | ❌ Public |

🎫 EventController

Manage event data (CRUD operations).

| Method   | Endpoint           | Description        | Auth            |
| :------- | :----------------- | :----------------- | :-------------- |
| **POST** | `/event`          | Create a new event | 🔒 Admin        |
| **GET**  | `/event/:EventID` | Get event by ID    | ✅ Authenticated |
| **PUT**  | `/event/:EventID` | Update event by ID | 🔒 Admin        |

⏳ QueueController

Handles user queues for events.

| Method   | Endpoint                                | Description                                    | Auth            |
| :------- | :-------------------------------------- | :--------------------------------------------- | :-------------- |
| **POST** | `/createQueue`                               | Create a new queue for an event                | ✅ Authenticated |
| **PUT**  | `/queue/:QueueID`                      | Update queue status (e.g., waiting → served)   | 🔒 Admin        |
| **GET**  | `/queue/:QueueID`                      | Get queue by queue ID                          | ✅ Authenticated |
| **GET**  | `/queue?event_id={id}&status={status}` | Get all queues filtered by event ID and status | ✅ Authenticated |

🛒 OrderController

Handles user orders and ticket generation.

| Method   | Endpoint               | Description                           | Auth            |
| :------- | :--------------------- | :------------------------------------ | :-------------- |
| **POST** | `/createOrder`              | Create a new order (generate tickets) | 🔒 Admin       |
| **GET**  | `/order?user_id={id}` | Get all orders by user ID             | ✅ Authenticated |

🎟️ TicketController

Retrieve ticket information by order or ticket ID.

| Method  | Endpoint                 | Description             | Auth            |
| :------ | :----------------------- | :---------------------- | :-------------- |
| **GET** | `/ticket?order_id={id}` | Get tickets by order ID | ✅ Authenticated |
| **GET** | `/ticket/:TicketID`     | Get ticket by ticket ID | ✅ Authenticated |

👥 UserController

Retrieve users by role (for admin management).

| Method  | Endpoint                | Description                                           | Auth     |
| :------ | :---------------------- | :---------------------------------------------------- | :------- |
| **GET** | `/users/:UserRole` | Get all users with a specific role (admin / customer) | 🔒 Admin |



## 🔄 End-to-End Flow Overview
| Step | Action               | Endpoint                           | Description                            |
| :--: | :------------------- | :--------------------------------- | :------------------------------------- |
|  1️⃣ | **Register & Login** | `/users/registerCustomer` or         | User registers and obtains a JWT token |
|      |                      |  `/users/registerAdmin` → `/login` |you can try create as Admin And User for| 
|      |                      |                                    |better experience                                        |
|  2️⃣ | **Create Event**     | `POST /event`                     | Admin creates a new event              |
|  3️⃣ | **Create Queue**     | `POST /createQueue`                     | User joins an event queue              |
|  4️⃣ | **Create Order**     | `POST /createOrder`                     | User makes an order after being served |
|  5️⃣ | **Update Queue**     | `PUT /queue/:QueueID`             | Admin updates queue status to “served” |
|  6️⃣ | **Get Ticket**       | `GET /ticket?order_id=<order_id>` | User retrieves generated tickets       |
|  7️⃣ | **Get Orders**       | `GET /order?user_id=<user_id>`    | User retrieves all their orders        |


## 🧩 Example and Request API
### Below are all available API endpoints with sample requests and responses.
🧑‍💼 Auth Endpoints

🔹 Register New Customer

POST /users/registerCustomer
```json
{
  "name": "CustomerHandsome",
  "email": "custHand@gmail.com",
  "password": "12345678"
}

```

✅ Response
```json
{
    "message": "User registered successfully",
    "user_id": "8"
}
```


🔹 Register New Admin

POST /users/registerAdmin
```json
{
  "name": "AdminHandsome",
  "email": "adminHand@gmail.com",
  "password": "12345678"
}
```

✅ Response
```json
{
    "message": "User registered successfully",
    "user_id": "9"
}
```

🔹 Login

POST /users/login
```json
{
  "email": "raffael@example.com",
  "password": "password123"
}
```

✅ Response
```json
{
  "message": "Login successful",
  "token": "<JWT_TOKEN>"
}
```

🎫 Event Endpoints

🔹 Create Event (Admin Only)

POST /event

Header: Authorization: Bearer <JWT_TOKEN>

```json
{
    "name": "Metal Concert",
    "location": "Kuta, Bali",
    "quota": 100,
    "date": "2025-12-01",
    "event_code": "BP",
    "description": "This Concert will be held at Bali"
}
```


✅ Response

```json
{
    "event": {
        "event_id": "26",
        "name": "Metal Concert",
        "location": "Kuta, Bali",
        "quota": 100,
        "date": "2025-12-01",
        "event_code": "BP",
        "description": "This Concert will be held at Bali",
        "created_at": "2025-10-05 18:47:00"
    },
    "message": "Event created successfully"
}
```

🔹 Get Event by ID

GET /event/:EventID

Header: Authorization: Bearer <JWT_TOKEN>

✅ Response

```json
{
    "Event": {
        "event_id": "26",
        "name": "Metal Concert",
        "location": "Kuta, Bali",
        "quota": 100,
        "date": "2025-12-01",
        "event_code": "BP",
        "description": "This Concert will be held at Bali",
        "created_at": "2025-10-05 18:47:00"
    }
}
```

🔹 Update Event (Admin Only)

PUT /event/:EventID

PUT /event/1

Header: Authorization: Bearer <JWT_TOKEN>

```json
{
    "name": "Metal Concert",
    "location": "Seminyak, Bali",
    "quota": 100,
    "date": "2029-12-01",
    "description": "This Concert will be held at Bali"
}
```


✅ Response

```json
{
    "Event": {
        "event_id": "26",
        "name": "Metal Concert",
        "location": "Seminyak, Bali",
        "quota": 100,
        "date": "2029-12-01",
        "description": "This Concert will be held at Bali"
    },
    "message": "Event updated successfully"
}
```

⏳ Queue Endpoints

🔹 Create Queue

POST /createQueue

Header: Authorization: Bearer <JWT_TOKEN>

```json
{
    "user_id": "4",
    "event_id": "26",
    "status": "waiting"
}
```


✅ Response

```json
{
    "Queue Data": {
        "queue_id": "9",
        "queue_number": 1,
        "event_id": "26",
        "user_id": "4",
        "status": "waiting"
    },
    "message": "Queue created successfully"
}
```

🔹 Update Queue (Admin Only)

PUT /queue/:QueueID

PUT /queue/1

Header: Authorization: Bearer <JWT_TOKEN>

```json
{
  "status": "served"
}
```


✅ Response

```json
{
    "message": "Queue updated successfully",
    "queue": {
        "queue_id": "9",
        "status": "served"
    }
}
```

🔹 Get Queue by ID

GET /queue/:QueueID

GET /queue/1

Header: Authorization: Bearer <JWT_TOKEN>

✅ Response

```json
{
    "queue": {
        "queue_id": "9",
        "queue_number": 1,
        "event_id": "26",
        "user_id": "4",
        "status": "served"
    }
}
```

🔹 Get Queues by Event & Status (Admin Only)

GET /queues?event_id=X&status=X

GET /queues?event_id=1&status=served

Header: Authorization: Bearer <JWT_TOKEN>

✅ Response

```json
{
    "queues": [
        {
            "queue_id": "9",
            "queue_number": 1,
            "event_id": "26",
            "user_id": "4",
            "status": "served"
        }
    ]
}
```

🛒 Order Endpoints
🔹 Create Order

POST /createOrder

Header: Authorization: Bearer <JWT_TOKEN>

```json
{
    "user_id": 4,
    "event_id": 26,
    "ticket_count": 50,
    "payment_method": "QRIS",
    "total_price": 25000000
}
```


✅ Response

```json
{
    "message": "Order created successfully",
    "order_id": "14"
}
```

🔹 Get Orders by User ID

GET /order?user_id=X

GET /order?user_id=1

Header: Authorization: Bearer <JWT_TOKEN>

✅ Response

```json
[
  {
        "order_id": "14",
        "event_id": "26",
        "user_id": "4",
        "total_price": 25000000,
        "payment_method": "QRIS"
  }
]
```

🎟️ Ticket Endpoints

🔹 Get Ticket by Order ID

GET /ticket?order_id=X

GET /ticket?order_id=1

Header: Authorization: Bearer <JWT_TOKEN>

✅ Response

```json
{
   "tickets": [
    {
        "ticket_id": "482",
        "order_id": "14",
        "ticket_number": "BP1",
        "price": 500000
    },
    {
        "ticket_id": "483",
        "order_id": "14",
        "ticket_number": "BP2",
        "price": 500000
    },
    {}
  ]
}
```

🔹 Get Ticket by Ticket ID

GET /ticket/:TicketID

GET /ticket/1

Header: Authorization: Bearer <JWT_TOKEN>

✅ Response

```json
{
    "ticket": {
        "ticket_id": "482",
        "order_id": "14",
        "ticket_number": "BP1",
        "price": 500000
    }
}
```

👥 User Endpoints (Admin Only)
🔹 Get Users by Role

GET /users/:UserRole

GET /users/customer

Header: Authorization: Bearer <JWT_TOKEN>

✅ Response

```json
{
  "users": [
        {
            "id": 1,
            "name": "Ch1riss",
            "email": "chris1012@gma",
            "password": "****",
            "role": "customer",
            "CreatedAt": "2025-10-05T16:31:51.553898Z"
        },
        {
            "id": 2,
            "name": "CustomerHandsome",
            "email": "custHand@gmail.com",
            "password": "****",
            "role": "customer",
            "CreatedAt": "2025-10-05T18:44:04.084968Z"
        }
  ]
}
```



## 🧠 Key Features

✅ CRUD operations for events, queues, orders, and tickets

✅ Custom validation for dates (e.g., YYYY-MM-DD format)

✅ Quota and queue management

✅ JWT authentication with role-based authorization

✅ Detailed error handling and logging

✅ Easy deployment setup for Railway / Render

✅ Clean controller separation and modular design
