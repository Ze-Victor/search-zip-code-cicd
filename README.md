# Zip Code Search

This is a project to retrieve address information based on a zip code.

## Prerequisites

- Go 1.21.1 installed
- Project dependencies installed

## Installation

To install the project, follow these steps:

1. Clone this repository and navigate to the main directory:
   ```bash
    $ git clone https://github.com/Ze-Victor/search-zip-code.git
    $ cd search-zip-code
   ```

2. Install project dependencies:

```bash
    $ go mod tidy
```

## Usage

1. Start the server:
```bash
    $ make
```

2. Make requests to the API using an HTTP client or web browser. Make sure to generate an authentication token before making the request and include it in the request header.

To generate the token, make a request to the `\auth` route passing `email` and `password`:

```bash
$ curl -X POST http://localhost:8001/api/v1/auth -H "Content-Type: application/json" -d '{"email": "your_email@test.com", "password": "your_password"}'
```

Make the request to get street data. Remember to replace `YOUR_AUTHENTICATION_TOKEN` with the generated JWT token:

```bash
$ curl -X GET http://localhost:8001/api/v1/cep/:cep_buscado -H "Content-Type: application/json" -H "Authorization: YOUR_AUTHENTICATION_TOKEN"
```

## Documentation

API Documentation

The API documentation is available at [Swagger UI](http://localhost:8001/api/v1/swagger/index.html) after starting the server.

## Unit Tests

To run the project tests, navigate to the main directory and use the command:

```bash
    $ cd search-zip-code
    $ make test
```

## End-to-end Test

End-to-end testing requires API initiation. To run these tests, start the server, open a terminal, navigate to the _"test"_ directory, and run the following command:

```bash
    $ cd test
    $ go test -v
```

## Health Endpoint
This endpoint can be accessed to check the health status of the system. It returns an HTTP status 200 if the system is healthy and an HTTP status 500 if the system is in an unhealthy state. Access:

- http://localhost:8001/api/v1/health


## Metrics Endpoint
This endpoint provides metrics about the system, such as CPU and memory usage. The metrics are returned in a specific format suitable for monitoring and analysis. Access:

- http://localhost:8001/api/v1/metrics


## Justification for Technology Choice

I chose to use Go as my main technology due to its widespread adoption by the team I am seeking a position with, as well as my familiarity with the language. Additionally, Go offers a variety of significant benefits, including its coherence and efficiency in handling concurrency, fast compilation, static typing, simplicity and clarity in syntax, as well as its runtime efficiency.

Additionally, I used the gin-gonic framework for routing. Besides having worked with it before, it is a simple and straightforward framework, suitable for small projects.

## Project Architecture

Go projects do not have a specific pattern or well-defined architecture, often leaving it to the team to establish a structure that best suits the project's development flow. As this was a small project with a single developer, I sought to create a structure that was easy to develop and understand, following some patterns from the Go project ecosystem.

My basic structure was as follows:

- **Main Directory**: Project configuration files are stored here.

- **cmd**: This is where the main applications of the project are located. In this case, we only have the API, which has the main file `main.go` and `routes.go` which contains the API routes.

- **config**: Global project configuration files.

- **internal**: Internal code for the application, not intended to be shared with external applications.

- **docs**: Contains project documentation, generated with gin-swagger.

- **test**: Unit tests are at the file level they are testing, but I created an end-to-end test to verify the entire functionality of address search through the ZIP code. It ensures that the system component works as expected.

With this structure, I sought simplicity, clarity, flexibility, and the ability to clearly separate responsibilities within the project. This can help facilitate code development, testing, and maintenance.

## Development Strategy

### Configuration and Change Management Planning

- I created the main branch, which will be the final branch of the project, and started a Go project in it.

- Next, I set up the development environment to receive new features.

- From the develop branch, features were created, which, after completion and testing, were integrated into the development environment and finally into the main branch.

- This flow helps maintain development organization and facilitates task division.

### Project Schedule

#### Day 1:
- Validation of request and ZIP code.
- Retrieval of ZIP code from simulated file.
- Replacement of numbers with zeros until finding the ZIP code.
- Code review.

#### Day 2:
- Unit tests and end-to-end tests.
- Project structuring for better adaptation to development.
- Code review.

#### Day 3:
- Code quality with static code analysis (linter).
- Makefile to automate compilation and execution tasks.
- Documentation with Swagger.
- Logger structuring.

#### Day 4:
- API authorization.
- API design.

_Note: The challenge file mentioned receiving and returning JSON. However, the task was to SEARCH for an address through a ZIP code, so it made sense to be a GET request. However, Getters do not take a body in the request. Therefore, I performed the search by sending the ZIP code as a parameter.
Code review.
Documentation review._

### Day 5:

- Endpoint for application health.
- Endpoint for application metrics.
- Rewriting test end to end
- README Rewrite
- Code Review

### Dia 06:

- Database Connection (SQLite)
- Refactoring of Code
- Code Review
- Test Review
- Documentation Review

## Answer to Question 02:

When we type an address (http://www.netshoes.com.br) in the browser, it starts a complex process. First, it needs to know where that address is hosted, i.e., what its domain is. To do this, it asks DNS for help, which works like a phone book, i.e., a domain server. This domain server resolves the address sent by the server and returns an IP that will be used by the browser to communicate client-server.

With the IP in hand, the browser and the server establish a connection. From then on, the browser (client) makes a request to the server through an HTTP request, and the server, upon receiving this request, processes the information and returns a response to the client (HTTP Response). Following a flow, it would be like this:

- DNS resolution.
- Establishment of TCP/IP connection.
- Sending the HTTP Request.
- Processing the request by the server.
- Sending the HTTP Response.
