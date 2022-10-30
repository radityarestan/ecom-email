# Ecommerce Email Service

This is a simple email service that sends emails to customers when they make a register to the app.

## How to run the app
1. Clone the repo
2. Verify that you have running docker-compose on the ecom-repo service
3. Create `app.env` file in the root directory of the repo. You can copy the content of `app.env.example` and fill the value with the environment variables.
   For NSQ, please use the same value as you write on the ecom-repo. For sender email and password, please use your [application password](https://myaccount.google.com/security) *you must enabled 2fa on your google account.
4. Run `go mod download` to download the dependencies
5. Run `go run main.go` to run the app
