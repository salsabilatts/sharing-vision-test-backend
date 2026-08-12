# Sharing Vision - Article Management Backend

Backend REST API for the Sharing Vision Article Management System coding assessment.

The backend provides article management functionality including creating, retrieving, updating, deleting, pagination, status filtering, and article validation.

## Tech Stack

- Go
- Gin
- GORM
- MySQL
- golang-migrate
- godotenv

## Features

- Create article
- Get article by ID
- Get articles with pagination
- Filter articles by status
- Update article
- Delete article
- Validate article data
- MySQL database migration
- Postman API Collection

## Article Status

The application supports the following article statuses:

~~~~text
publish
draft
thrash
~~~~

## Validation

Before creating or updating an article, the backend validates:

- Title: required, minimum 20 characters
- Content: required, minimum 200 characters
- Category: required, minimum 3 characters
- Status: required and must be one of:
~~~~text
publish
draft
thrash
~~~~
## Project Structure
~~~~
sharing-vision-test-backend/
│
├── cmd/
│   └── server/
│       └── main.go
│
├── config/
│   └── database.go
│
├── internal/
│   ├── handler/
│   │   └── post_handler.go
│   │
│   ├── model/
│   │   └── post.go
│   │
│   ├── repository/
│   │   └── post_repository.go
│   │
│   └── service/
│       └── post_service.go
│
├── migrations/
│   ├── 000001_create_posts_table.up.sql
│   └── 000001_create_posts_table.down.sql
│
├── postman/
│   └── Article API.postman_collection.json
│
├── .env.example
├── .gitignore
├── go.mod
├── go.sum
└── README.md
~~~~

## Prerequisites

Make sure the following are installed:

- Go
- MySQL
- golang-migrate
- Git
- Postman (optional, for API testing)
  
## Getting Started
### 1. Clone the Repository
~~~~
git clone https://github.com/salsabilatts/sharing-vision-test-backend.git
cd sharing-vision-test-backend
~~~~
### 2. Create the MySQL Database
Open MySQL:
~~~~
mysql -u root -p
~~~~
Create the database:
~~~~
CREATE DATABASE article;
~~~~
Verify the database:
~~~~
SHOW DATABASES;
~~~~
Then exit MySQL:
~~~~
exit;
~~~~

### 3. Configure Environment Variables

Copy the example environment file:
~~~~
cp .env.example .env
~~~~
Update .env with your local MySQL configuration.

Example:
~~~~
APP_PORT=8080

DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_password
DB_NAME=article
~~~~

### 4. Install Go Dependencies
~~~~
go mod download
~~~~
### 5. Run Database Migration

From the backend directory:
~~~~
migrate -path ./migrations -database "mysql://root:YOUR_PASSWORD@tcp(localhost:3306)/article" up
~~~~
Replace YOUR_PASSWORD with your local MySQL root password.

The migration creates the posts table required by the application.

### 6. Run the Backend
~~~~
go run ./cmd/server
~~~~
The API will run on:
~~~~
http://localhost:8080
~~~~
## Health Check

To verify that the backend is running:
~~~~
GET http://localhost:8080/health
~~~~
Expected response:
~~~~
{
  "status": "ok"
}
~~~~
## API Endpoints
| Method | Endpoint                                 | Description                  |
| ------ | ---------------------------------------- | ---------------------------- |
| POST   | `/article/`                              | Create a new article         |
| GET    | `/article/:limit/:offset`                | Get articles with pagination |
| GET    | `/article/:limit/:offset?status=publish` | Get filtered articles        |
| GET    | `/article/:id`                           | Get an article by ID         |
| PUT    | `/article/:id`                           | Update an article            |
| PATCH  | `/article/:id`                           | Update an article            |
| DELETE | `/article/:id`                           | Delete an article            |
Create Article

POST /article/

Request body:

{
  "title": "Example Article Title With More Than Twenty",
  "content": "Example article content that is longer than two hundred characters and satisfies the validation rules implemented by the backend service.",
  "category": "Technology",
  "status": "draft"
}
Get Articles
GET /article/10/0
