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
| **POST** | `/register/customer` | Register a new customer  | ❌ Public |
| **POST** | `/register/admin`    | Register a new admin     | ❌ Public |
| **POST** | `/login`             | Log in and get JWT token | ❌ Public |

🎫 EventController

Manage event data (CRUD operations).

| Method   | Endpoint           | Description        | Auth            |
| :------- | :----------------- | :----------------- | :-------------- |
| **POST** | `/events`          | Create a new event | 🔒 Admin        |
| **GET**  | `/events/:EventID` | Get event by ID    | ✅ Authenticated |
| **PUT**  | `/events/:EventID` | Update event by ID | 🔒 Admin        |

⏳ QueueController

Handles user queues for events.

| Method   | Endpoint                                | Description                                    | Auth            |
| :------- | :-------------------------------------- | :--------------------------------------------- | :-------------- |
| **POST** | `/queues`                               | Create a new queue for an event                | ✅ Authenticated |
| **PUT**  | `/queues/:QueueID`                      | Update queue status (e.g., waiting → served)   | 🔒 Admin        |
| **GET**  | `/queues/:QueueID`                      | Get queue by queue ID                          | ✅ Authenticated |
| **GET**  | `/queues?event_id={id}&status={status}` | Get all queues filtered by event ID and status | 🔒 Admin        |

🛒 OrderController

Handles user orders and ticket generation.

| Method   | Endpoint               | Description                           | Auth            |
| :------- | :--------------------- | :------------------------------------ | :-------------- |
| **POST** | `/orders`              | Create a new order (generate tickets) | ✅ Authenticated |
| **GET**  | `/orders?user_id={id}` | Get all orders by user ID             | ✅ Authenticated |

🎟️ TicketController

Retrieve ticket information by order or ticket ID.

| Method  | Endpoint                 | Description             | Auth            |
| :------ | :----------------------- | :---------------------- | :-------------- |
| **GET** | `/tickets?order_id={id}` | Get tickets by order ID | ✅ Authenticated |
| **GET** | `/tickets/:TicketID`     | Get ticket by ticket ID | ✅ Authenticated |

👥 UserController

Retrieve users by role (for admin management).

| Method  | Endpoint                | Description                                           | Auth     |
| :------ | :---------------------- | :---------------------------------------------------- | :------- |
| **GET** | `/users/role/:UserRole` | Get all users with a specific role (admin / customer) | 🔒 Admin |



## 🔄 End-to-End Flow Overview
| Step | Action               | Endpoint                           | Description                            |
| :--: | :------------------- | :--------------------------------- | :------------------------------------- |
|  1️⃣ | **Register & Login** | `/register` → `/login`             | User registers and obtains a JWT token |
|      |                      |                                    |you can try create as Admin And User for| 
|      |                      |                                    |better experience                                        |
|  2️⃣ | **Create Event**     | `POST /events`                     | Admin creates a new event              |
|  3️⃣ | **Create Queue**     | `POST /queues`                     | User joins an event queue              |
|  4️⃣ | **Create Order**     | `POST /orders`                     | User makes an order after being served |
|  5️⃣ | **Update Queue**     | `PUT /queues/:QueueID`             | Admin updates queue status to “served” |
|  6️⃣ | **Get Ticket**       | `GET /tickets?order_id=<order_id>` | User retrieves generated tickets       |
|  7️⃣ | **Get Orders**       | `GET /orders?user_id=<user_id>`    | User retrieves all their orders        |


## 🧩 Example and Request API
### Below are all available API endpoints with sample requests and responses.
🧑‍💼 Auth Endpoints

🔹 Register New Customer

POST /register/customer
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

POST /register/admin
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

POST /login
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

POST /events
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

GET /events/:EventID
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

PUT /events/:EventID
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

POST /queues
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

PUT /queues/:QueueID
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

GET /queues/:9
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

GET /queues?event_id=26&status=served
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

POST /orders
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

GET /orders?user_id=4
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

GET /tickets?order_id=4
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

GET /tickets/:TicketID
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

GET /users/role/:UserRole
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