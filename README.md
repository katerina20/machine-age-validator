# Machine Age Validator Microservice

## Overview

This microservice is designed to validate the ages of machines in a manufacturing company's inventory system. It identifies machines with unreasonably long ages by comparing them against other machines in the submitted list using IQR algorithm.

## Getting Started

### Installation and running
Make sure Go is installed

1. Clone the repository:
`git clone https://github.com/katerina20/machine-age-validator`

2. Navigate to the project directory:
`cd machine-age-validator`

3. Install dependencies:
`go mod download`

4. Start the service:
`go run main.go`

The service will be available at `http://localhost:8080`.

## Usage

To use the microservice, send a POST request to `/validate` with a JSON body containing a list of machine IDs and their corresponding age strings.

For age can be use: days, weeks, months and years.

Example request:

```json
[{"id": "machine1", "age": "1 year"}, {"id": "machine2", "age": "200 days"}, {"id": "machine3", "age": "90 years"}, {"id": "machine4", "age": "12 month"}, {"id": "machine5", "age": "300 days"}, {"id": "machine6", "age": "3 month"}]
```

Example response:

```json
{"id": "machine3", "age": "90 years"}
```

### Things to do
1. Improve detection algorithm to use with any data size (even with small dataset).
2. Save coefficient in the db for more exact result.
3. Add more test cases to be sure the results are expected.
4. Create API documentation for endpoints, data formats, and usage examples.
5. Implement security to use within the organization.
