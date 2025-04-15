# Stress Test Go

`go-stress-test` is a simple tool to perform stress tests on URLs by sending multiple concurrent requests and generating a final report with the results.

## Prerequisites

- [Docker](https://www.docker.com/) installed on your machine.

## How to Build and Run the Project with Docker

### 1. Build the Docker Image

In the root directory of the project, run the following command to build the Docker image:

```bash
docker build -t go-stress-test .
```

### 2. Run the Stress Test

After building the image, you can run the stress test by passing the required parameters via the command line:

```bash
docker run --rm go-stress-test -requests=<number> -url=<URL> -concurrency=<number>
```

#### Required Parameters:
- `-requests`: Total number of requests to be made.
- `-url`: The URL to be tested.
- `-concurrency`: Number of concurrent requests.

#### Example:
```bash
docker run --rm go-stress-test -requests=20 -url="https://www.google.com" -concurrency=5
```

### 3. Final Report

After execution, the program will display a final report in the terminal containing:
- **Tested URL**: The target URL of the test.
- **Total execution time**: The total time taken to complete the test.
- **Total number of requests made**: The total number of requests sent.
- **Number of successful requests**: The number of requests that succeeded.
- **Number of failed requests**: The number of requests that failed.
