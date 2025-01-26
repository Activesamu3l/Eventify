# Eventify

A simple REST API for creating, managing, and registering for events. Built with Go, Gin framework, and SQLite.

## Features

- User authentication with JWT.
- CRUD operations for events.
- Event registration and cancellation.
- Secure password hashing using bcrypt.

## Tech Stack

- **Backend**: Go (Gin Framework)
- **Database**: SQLite
- **Authentication**: JWT
- **Hashing**: Bcrypt

## Endpoints

### Authentication
- **POST** `/signup`: Create a new user.
- **POST** `/login`: Login and obtain a JWT.

### Event Management
- **GET** `/events`: Retrieve all events.
- **GET** `/events/:id`: Retrieve a single event by ID.
- **POST** `/events`: Create a new event (requires authentication).
- **PUT** `/events/:id`: Update an event (requires authentication).
- **DELETE** `/events/:id`: Delete an event (requires authentication).

### Event Registration
- **POST** `/events/:id/register`: Register for an event (requires authentication).
- **DELETE** `/events/:id/register`: Cancel registration for an event (requires authentication).

## Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/eventify.git
   cd eventify

2. Install dependencies:
    ```bash
    go mod tidy

3. Run the application:
    ```bash
    go run main.go

4. The API will be available at http://localhost:8080.
