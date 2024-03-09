# Go Dummy Bank API

This is a simple bank API implemented in Go using `net/http` standard library.

## Features

- User management: Create, read, update, delete user accounts
- Account management: Create, read, update, delete bank accounts, deposit, withdraw, transfer funds
- Transaction management: retrieve transaction history

## Prerequisites

- Go programming language installed on your system
- MySQL or another compatible database installed and running
- `.env` file containing environment variables (e.g., database connection details, server port)

## Installation

1. Clone this repository to your local machine:

```bash
git clone <repository-url>
```

2. Navigate to the project directory:

```bash
cd go-dummy-bank-api

```

3. Create a .env file in the project root and add the following environment variables:

```bash
PORT=:8080
USERNAME=<DB_USERNAME>
PASSWORD=<DB_PASSWORD>
```

## Usage

1. Install dependencies:

```bash
go mod tidy
```

2. Run server

```bash
make run
```

3. Access the API endpoints using tools like cURL or Postman.

## API endpoints

- `POST /users`: Create a new user
- `GET /users/{id}`: Retrieve user details by ID
- `PUT /users/{id}`: Update user details by ID
- `DELETE /users/{id}`: Delete user by ID
- `POST /accounts`: Create a new account for a user
- `GET /accounts/{id}`: Retrieve account details by ID
- `PUT /accounts/{id}`: Update account details by ID
- `DELETE /accounts/{id}`: Delete account by ID
- `POST /accounts/{id}/deposit`: Deposit funds into the account
- `POST /accounts/{id}/withdraw`: Withdraw funds from the account
- `POST /accounts/{fromId}/transfer/{toId}`: Transfer funds between accounts
- `GET /transitions/{id}`: Retrieve transaction history for an account by account id
