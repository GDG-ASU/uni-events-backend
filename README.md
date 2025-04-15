
**University Event Management System**

**Description:**

This system is designed to facilitate event participation and management within a university environment. The platform supports three user roles: **Students**, **Club Owners**, and **Admins**.

- **Students** can browse and join events organized by various university clubs.
- **Club Owners** can create and manage events for their respective clubs.
- **Admins** have full control over the system, including managing users, clubs, and events.
The goal is to streamline event participation, foster student engagement, and provide an organized platform for managing university-related events.


![image](https://github.com/user-attachments/assets/c97010b3-0790-4251-b03e-82e89d403215)

---

### **Tech Stack**
#### **Backend:**
- **Language:** Go (Golang)
- **Framework:** Echo
- **ORM:** GORM
- **Database:** Neon (PostgreSQL)
- **Caching:** Redis
- **Authentication:** Clerk (for user management and authentication)
- **Authorization**: Casbin
#### **Frontend:**
- **Development Environment:** Android Studio
- Language
#### **Deployment:**
<Will be discussed later>

---

**Overview of Casbin ABAC **

#### **Key Concepts:**
- **User Attributes:** Information about the user (e.g., role, department, club membership).
- **Resource Attributes:** Information about the resource (e.g., event details, club details).
- **Action Attributes:** Operations that can be performed (e.g., create, read, update, delete).
- **Policies:** Define who can perform what actions on which resources, based on attributes.
---

**Folder structure for GO (Backend)**

```
university-event-app/
│
├── cmd/
│   └── server/
│       └── main.go                   # Main entrypoint (starts Echo server)
│
├── config/
│   ├── config.go                     # Loads env vars, DB, Redis, Clerk configs
│   ├── clerk.go                      # Clerk initialization/config
│   ├── db.go                         # GORM DB connection setup
│   ├── redis.go                      # Redis client setup
│   └── .env.example                  # Example environment variables
│
├── internal/
│   ├── api/
│   │   ├── auth/                     # Auth endpoints (login, logout, token verify)
│   │   │   └── handler.go
│   │   ├── events/                   # CRUD endpoints for events
│   │   │   └── handler.go
│   │   ├── clubs/                    # Club creation, editing, listing
│   │   │   └── handler.go
│   │   ├── users/                    # Student, admin, club-owner details
│   │   │   └── handler.go
│   │   └── admin/                    # Admin-only actions
│   │       └── handler.go
│   │
│   ├── models/                       # GORM models
│   │   ├── user.go                   # User model with role (student/owner/admin)
│   │   ├── club.go
│   │   ├── event.go
│   │   ├── student_event_join.go    # Many-to-many relation table
│   │   └── base.go                  # Common fields (ID, timestamps)
│   │
│   ├── services/                     # Business logic
│   │   ├── auth_service.go
│   │   ├── event_service.go
│   │   ├── club_service.go
│   │   └── user_service.go
│   │
│   ├── middlewares/                 # Echo middleware
│   │   ├── auth_middleware.go       # Clerk token verify
│   │   ├── casbin_middleware.go     # Casbin enforcement
│   │   └── logger.go
│   │
│   ├── repositories/                # DB access layer
│   │   ├── user_repo.go
│   │   ├── event_repo.go
│   │   └── club_repo.go
│   │
│   └── policies/                    # Casbin ABAC configs
│       ├── model.conf               # ABAC model definition
│       ├── policy.csv               # CSV-based rules (for testing)
│       └── enforcer.go             # Casbin enforcer setup
│
├── pkg/
│   ├── utils/                       # Utility helpers
│   │   ├── token.go                 # JWT/Clerk token parsing
│   │   └── response.go             # Standard JSON response wrapper
│   └── logger/                     # Logger setup (Zap, Logrus, etc.)
│       └── logger.go
│

│
├── casbin/
│   ├── model.conf                  # ABAC model definition (duplicate for clarity)
│   └── policy.csv                  # Static policy rules (optional)
│

│
├── Dockerfile                      # Go backend container
├── docker-compose.yml              # Compose setup for backend + Redis + DB
├── .dockerignore
├── .env                            # Your actual env vars
├── go.mod
├── go.sum
└── README.md
```
---

### **Base URL**
All endpoints are prefixed with `/api/v1` 

 **Authentication (Clerk)**

- User registration and login are handled on the frontend using Clerk SDK.
- The backend uses Clerk JWT for verifying requests.
- Example of backend endpoint for authenticated user:
    - `GET /users/getme`  – returns current user profile (used to determine role/ID).
---

### **Students**
- `GET /students/me/events`  – get a list of events the student has joined.
- `POST /students/events/:eventID/join`  – student joins a specific event.
- `DELETE /students/events/:eventID/leave`  – student leaves a specific event.
---

### **Clubs**
- `GET /clubs`  – list all available clubs (public).
- `GET /clubs/:id`  – get specific club details.
- `GET /clubs/:id/events`  – list events for a specific club.
---

### **Club Owners**
- `GET /my-club`  – get the club managed by the currently logged-in club owner.
- `POST /my-club/events`  – create a new event for the club.
- `PUT /my-club/events/:eventID`  – update a specific event.
- `DELETE /my-club/events/:eventID`  – delete a specific event.
- `GET /my-club/events/:eventID/students`  – list students who joined the event.
---

### **Events**
- `GET /events`  – list all public events (for students to browse).
- `GET /events/:id`  – get detailed info for a specific event.
---

### **Admins**
- `GET /admin/users`  – list all users in the system.
- `PUT /admin/users/:userID/role`  – change a user’s role (e.g., student → club owner).
- `GET /admin/clubs`  – list all clubs.
- `POST /admin/clubs`  – create a new club.
- `DELETE /admin/clubs/:clubID`  – delete a club.
- `GET /admin/events`  – list all events in the system.
- `DELETE /admin/events/:eventID`  – delete any event.

