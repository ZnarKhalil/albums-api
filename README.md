# Albums API

This project is a RESTful API to learn Golang, using the Gin framework and MySQL for data persistence.

It allows users to manage a music album database, with functionality to add, retrieve, and filter albums, inspired by ideas from Go tutorials pages:

1. [Database Access Tutorial](https://go.dev/doc/tutorial/database-access)
2. [Web Service with Gin Tutorial](https://go.dev/doc/tutorial/web-service-gin)

## Features

- **Retrieve all albums** or filter by artist.
- **Retrieve a single album** by its ID.
- **Add new albums** to the database.

## Table of Contents

- [Requirements](#requirements)
- [Installation](#installation)
- [Running the Project](#running-the-project)
- [API Endpoints](#api-endpoints)
- [Project Structure](#project-structure)

## Requirements

- [Go](https://golang.org/doc/install) 1.16+
- [MySQL](https://dev.mysql.com/downloads/mysql/)
- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)
- [godotenv](https://github.com/joho/godotenv) for managing environment variables

## Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/ZnarKhalil/albums-api.git
   cd albums-api
   ```

2. **Copy the example environment file:**

   ```bash
   cp .env.example .env
   ```

3. **Update the .env file with your MySQL database credentials**

4. **Download dependencies:**

   ```bash
   go mod tidy
   ```

5. **Set up the database:** by creating a MySQL database named recordings and a table album

   ```sql
   CREATE DATABASE recordings;
   USE recordings;

   CREATE TABLE album (
       id INT AUTO_INCREMENT PRIMARY KEY,
       title VARCHAR(255) NOT NULL,
       artist VARCHAR(255) NOT NULL,
       price DECIMAL(10, 2) NOT NULL
   );
   ```

## Running the Project

To start the server, run:

    go run main.go

The server will run on localhost:8080 by default.

## API Endpoints

1. **Retrieve All Albums:**

- URL: /albums
- Method: GET
- Query Parameter (optional): artist
- Description: Retrieves all albums, or filters albums by the specified artist.

  ```bash
  curl http://localhost:8080/albums?artist=John%20Coltrane
  ```

2. **Add a New Album:**

- URL: /albums
- Method: POST
- Description: Adds a new album to the database.
- Request Body:

  ```json
  {
    "title": "The Modern Sound of Betty Carter",
    "artist": "Betty Carter",
    "price": 49.99
  }
  ```

  ```bash
  curl -X POST http://localhost:8080/albums -H "Content-Type: application/json" -d '{
  "title": "The Modern Sound of Betty Carter",
  "artist": "Betty Carter",
  "price": 49.99
  }'
  ```

3. **Retrieve a Single Album by ID:**

- URL: /albums/:id
- Method: GET
- Description: Retrieves the album with the specified ID.

  ```bash
  curl http://localhost:8080/albums/2
  ```

## Project Structure

The project follows a standard Go project structure:

    library/
    ├── main.go                # Entry point of the application
    ├── db/
    │   └── db.go              # Database connection setup
    ├── handlers/
    │   └── album.go           # Handler functions for album-related endpoints
    ├── models/
    │   └── album.go           # Album struct and database functions
    ├── .env.example           # Environment variables file
    ├── go.mod                 # Go module file
    └── README.md              # Project documentation
