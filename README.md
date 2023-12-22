# Go.Products.API

This repository is made to document CRUD Operation of a simple Products API in Golang. The API uses GORM for database operations and MySQL as the database.

## Prerequisites

Before running the application, make sure you have the following installed on your system:

- Golang
- GORM
- MySQL

## Installation

1. Clone the repository: `git clone https://github.com/nilanjanb3/Go.Products.API`
2. Change into the project directory: `cd Go.Products.API`
3. Install the dependencies: `go get -d ./...`

## Configuration

The application requires a MySQL database to be set up. Update the database connection details in the `constants.go` file located in the `src` directory.

## Usage

To run the application, use the following command:

```bash
go run *.go
```

The application will be accessible at http://localhost:3000.

API Endpoints
-------------

The following API endpoints are available:

- `GET /`: Retrieves the homepage.
- `GET /products`: Retrieves all products.
- `GET /product/{id}`: Retrieves a specific product by ID.
- `POST /products`: Creates a new product.
- `PUT /product/{id}`: Updates a product by ID.
- `DELETE /product/{id}`: Deletes a product by ID.

License
-------

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.