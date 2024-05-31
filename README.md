# Book Management System

A comprehensive Book Management System built with Go. This system allows users to perform CRUD operations on book resources.

## Overview

This project leverages the power of Go to create a robust and efficient system for managing books. The backend is developed using Go with Gorilla Mux for routing, and GORM for database interactions.

## Features

- **Create Books:** Add new books to the system with details such as ISBN, title, author, and publication date.
- **Read Books:** View a list of all books or get detailed information about a specific book.
- **Update Books:** Modify existing book details to keep the system up-to-date.
- **Delete Books:** Remove books from the system when they are no longer needed.

## Technologies Used

### Backend
- **Go:** High-performance language for backend development.
- **Gorilla Mux:** Routing library for handling HTTP requests.
- **GORM:** ORM library for database interactions.
- **MySQL:** Database for storing book information.

## Project Structure

### Backend
- **main.go:** Entry point of the application.
- **routes:** Contains route definitions.
- **controllers:** Handles business logic for each endpoint.
- **models:** Defines the data structures.
- **utils:** Utility functions.
- **config:** Configuration settings for the application.

## Getting Started

### Prerequisites
- Go (version 1.22.2 or higher)
- MySQL

### Installation

1. **Clone the Repository:**
   ```bash
   git clone https://github.com/Rizz-33/book-management-system.git
   ```

2. **Navigate to the Project Directory:**
   ```bash
   cd book-management-system
   ```

3. **Setup the Backend:**
   - Ensure MySQL is running and create a database for the application.
   - Update the configuration settings in the \`config\` package.
   - Build and run the backend application:
     ```bash
     go run main.go
     ```

### Running the Application

The backend server will be available at `http://localhost:9010`.

## API Endpoints

### Books

- **Get all books:**
  - ```GET /books```
  
- **Get a single book:**
  - ```GET /books/{id}```
    
- **Create a book:**
  - ```POST /books```
  - Sample request body:
    ```json
    {
      "isbn": "978-3-16-148410-0",
      "title": "Book Title",
      "author": "Author Name",
      "publication_date": "2023-05-01"
    }
    ```

- **Update a book:**
  - ```PUT /books/{id}```
  - Sample request body:
    ```json
    {
      "isbn": "978-3-16-148410-0",
      "title": "Updated Book Title",
      "author": "Updated Author Name",
      "publication_date": "2024-01-01"
    }
    ```

- **Delete a book:**
  - ```DELETE /books/{id}```

## References

- Referred to various YouTube tutorials to understand best practices for building RESTful APIs with Go and integrating GORM for database interactions.
