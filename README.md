# Bank-Challenge

This project is a challenge related to bank transactions in a digital bank called "bank-challenge". The service is a composition of a microservice written in golang, together with an instance of the "Postgresql" database. To test the service, just clone this repository and enter the project directory and execute the following command:

``` make run-all ``` to start the bank-challenge and postgres container.

``` make stop-all ``` to stop all containers.

In addition, in the project's "docs" directory, we have a postman collection that can be used to test the platform. Among the available endpoints are:

- Accounts:
    - GET: /accounts - Gets all accounts registred (Should be Authenticated).
    - GET: /accounts/{account_id}/balance - Gets the balance of specific account (Should be Authenticated).
    - POST: /accounts - Creates a new account.

- Login:
    - POST: /login - Gets the access token.

- Transfers:
    - GET: /tranfers - Gets the list of transfers realized to specific account id (Should be Authenticated).
    - POST: /transfers - Makes a money transfer money to another account (Should be Authenticated).