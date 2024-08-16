# Lightweight Bank API

Lightweight Bank API is a simple banking application built with Go. It provides basic functionalities such as user registration, login, and money transfer between accounts.

## Table of Contents

- [Installation](#installation)
- [Configuration](#configuration)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [License](#license)

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/barisaskaleli/lightweight-bank.git
    cd lightweight-bank
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

3. Set up the environment variables:
   Copy the `.env.example` file to `.env` and adjust the values as needed.
    ```sh
    cp .env.example .env
    ```

## Configuration

The application uses environment variables for configuration. Below is a list of the variables you need to set in the `.env` file:

- `SERVER_PORT`: The port on which the server will run.
- `JWT_SECRET`: Secret key for JWT token generation.
- `DB_HOST`: Database host.
- `DB_PORT`: Database port.
- `DB_USER`: Database user.
- `DB_PASSWORD`: Database password.
- `DB_DATABASE`: Database name.
- `TRANSFER_FEE`: Fee for money transfers.
- `TRANSFER_FEE_ENABLED`: Enable or disable transfer fee.
- `SMS_INFO_ENABLED`: Enable or disable SMS notifications.

## Usage

1. Run the application via docker:
    ```sh
    docker compose up -d --build
    ```
2. Or run by the golang:
    ```sh
    go run cmd/server.go
    ```

2. The server will start on the port specified in the `.env` file.

## API Endpoints

Swagger documentation is available at
````
http://localhost:3000/swagger/index.html
````

### User Registration

- **URL**: `/register`
- **Method**: `POST`
- **Request Body**:
    ```json
    {
        "name": "John",
        "surname": "Doe",
        "email": "john.doe@example.com",
        "password": "password123"
    }
    ```
- **Response**:
    ```json
    {
        "id": 1,
        "account_number": "1234567890",
        "name": "John",
        "surname": "Doe",
        "email": "john.doe@example.com",
        "balance": 0.0
    }
    ```

### User Login

- **URL**: `/login`
- **Method**: `POST`
- **Request Body**:
    ```json
    {
        "email": "john.doe@example.com",
        "password": "password123"
    }
    ```
- **Response**:
    ```json
    {
        "token": "jwt_token_here"
    }
    ```

### Money Transfer

- **URL**: `/transfer`
- **Method**: `POST`
- **Request Body**:
    ```json
    {
        "sender_account_number": "1234567890",
        "receiver_account_number": "0987654321",
        "amount": 100.0
    }
    ```
- **Response**:
    ```json
    {
        "status": true,
        "message": "Transfer successful",
        "sender_balance": 900.0,
        "receiver_balance": 1100.0,
        "fee": 10.0
    }
    ```

## License

This project is licensed under the MIT License.