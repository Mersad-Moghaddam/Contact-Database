Go Database Project
A simple database project built with Go, providing a RESTful API for managing data.

Table of Contents

Introduction

Getting Started

Endpoints

Database Schema

API Documentation

Contributing

License

Introduction

This project is a simple database project built with Go, providing a RESTful API for managing data. The project uses the gorilla/mux router and the database/sql package to interact with the database.

Getting Started
To get started with this project, follow these steps:

Clone the repository: git clone https://github.com/your-username/go_database.git

Install dependencies: go get -u github.com/gorilla/mux

Create a new database and schema using the schema.sql file

Update the database.go file with your database credentials

Run the API: go run main.go

Endpoints

The following endpoints are available:

GET /data: Get all data

GET /data/{id}: Get a specific data record by ID

POST /data: Create a new data record

PATCH /data/{id}: Update a specific data record

DELETE /data/{id}: Delete a specific data record

Database Schema:

The database schema is defined in the schema.sql file. The schema consists of a single table with the following columns:

id: primary key

name: string

email: string

API Documentation

For more information on the API endpoints and parameters, see the API Documentation.

Contributing
Contributions are welcome! If you'd like to contribute to this project, please fork the repository and submit a pull request.

License
This project is licensed under the MIT License. See the LICENSE file for more information.
