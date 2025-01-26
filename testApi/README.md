# Simple Banking API

This is a Go-based API application for managing user transactions. It provides endpoints for user registration, top-up, payment, transfer, and transaction reporting.

## Table of Contents
- [Prerequisites](#prerequisites)
- [Running the Application](#running-the-application)
- [API Endpoints](#api-endpoints)
  - [User Registration](#user-registration)
  - [Top-Up](#top-up)
  - [Payment](#payment)
  - [Transfer](#transfer)
  - [Transaction Report](#transaction-report)

## Prerequisites

- Go (version 1.16 or higher)
- PostgreSQL (or your preferred database)
- Required Go packages (use `go get` to install)

## Running the Application

   ```bash
   go run main.go
   ```

The application will start on `http://localhost:8080` (or the port specified in your code).

## API Endpoints

### User Registration

- **Endpoint**: `POST /register`
- **Request Body**:
  ```json
  {
    "first_name": "John",
    "last_name": "Doe",
    "phone_number": "1234567890",
    "address": "123 Main St",
    "pin": "1234"
  }
  ```
- **Response**:
  - **201 Created**:
    ```json
    {
      "status": "SUCCESS",
      "result": {
        "user_id": "uuid",
        "first_name": "John",
        "last_name": "Doe",
        "phone_number": "1234567890",
        "address": "123 Main St",
        "created_at": "2023-01-01T00:00:00Z"
      }
    }
    ```
  - **409 Conflict** (if phone number already registered):
    ```json
    {
      "error": "Phone number already registered"
    }
    ```

### Top-Up

- **Endpoint**: `POST /topup`
- **Request Body**:
  ```json
  {
    "amount": 100.00
  }
  ```
- **Response**:
  - **201 Created**:
    ```json
    {
      "status": "SUCCESS",
      "result": {
        "transaction_id": "uuid",
        "amount_top_up": 100.00,
        "balance_before": 50.00,
        "balance_after": 150.00,
        "created_date": "2023-01-01T00:00:00Z"
      }
    }
    ```
  - **400 Bad Request** (if balance is not enough):
    ```json
    {
      "message": "Balance is not enough"
    }
    ```

### Payment

- **Endpoint**: `POST /payment`
- **Request Body**:
  ```json
  {
    "amount": 50.00,
    "remarks": "Payment for services"
  }
  ```
- **Response**:
  - **201 Created**:
    ```json
    {
      "status": "SUCCESS",
      "result": {
        "transaction_id": "uuid",
        "amount": 50.00,
        "remarks": "Payment for services",
        "balance_before": 150.00,
        "balance_after": 100.00,
        "created_date": "2023-01-01T00:00:00Z"
      }
    }
    ```
  - **400 Bad Request** (if balance is not enough):
    ```json
    {
      "message": "Balance is not enough"
    }
    ```

### Transfer

- **Endpoint**: `POST /transfer`
- **Request Body**:
  ```json
  {
    "target_user": "target-uuid",
    "amount": 30.00,
    "remarks": "Transfer to friend"
  }
  ```
- **Response**:
  - **201 Created**:
    ```json
    {
      "status": "SUCCESS",
      "result": {
        "transaction_id": "uuid",
        "amount": 30.00,
        "remarks": "Transfer to friend",
        "balance_before": 100.00,
        "balance_after": 70.00,
        "created_date": "2023-01-01T00:00:00Z"
      }
    }
    ```
  - **400 Bad Request** (if balance is not enough):
    ```json
    {
      "message": "Balance is not enough"
    }
    ```

### Transaction Report

- **Endpoint**: `GET /transaction-report`
- **Response**:
  - **200 OK**:
    ```json
    {
      "status": "SUCCESS",
      "result": [
        {
          "transaction_id": "uuid",
          "amount": 100.00,
          "transaction_type": "CREDIT",
          "created_date": "2023-01-01T00:00:00Z"
        },
        {
          "transaction_id": "uuid",
          "amount": 50.00,
          "transaction_type": "DEBIT",
          "created_date": "2023-01-02T00:00:00Z"
        }
      ]
    }
    ```

