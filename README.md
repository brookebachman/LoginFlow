# LoginFlow

Installation and Running the Project

Prerequisites
Go 1.24.2 or higher installed on your machine. You can check your Go version by running:

go version

Docker (optional, for running in Docker container).

SQLite is used as the default database, and it will be created automatically when the API starts if it doesn't exist.

git clone https://github.com/yourusername/loginflow.git
cd login_form_cork

2. Install Dependencies
   Run the following to install the Go dependencies:

go mod tidy

3. Running the API Locally
   Once the dependencies are installed, you can run the API locally with the following command:

go run main.go

This will start the API server on http://localhost:8080. You can test the endpoints using a tool like Postman or cURL.

4. Running the API with Docker (Optional)
   If you prefer to run the API inside a Docker container, you can follow these steps:

Build the Docker Image

docker build -t login-api .

docker run -p 8080:8080 --name login-api login-api
