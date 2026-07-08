# ASCII Art Web

## Description

ASCII Art Web is a web-based implementation of the ASCII Art project written in Go. It allows users to enter text through a web interface, choose one of the available banner styles, generate the corresponding ASCII art, and download the generated output as a text file.

The application reuses the ASCII rendering engine developed in the previous CLI version while providing a clean, responsive, and user-friendly web interface.

---

## Features

* Generate ASCII art from user input.
* Supports three banner styles:

  * **standard**
  * **shadow**
  * **thinkertoy**
* Clean web interface built with HTML templates.
* Responsive styling using external CSS.
* Download generated ASCII art as a `.txt` file.
* Automated unit tests for rendering functions and HTTP handlers.
* Request logging and diagnostics.
* Proper HTTP status code handling.
* Input validation for unsupported characters.
* Cross-platform newline (`\n` and `\r\n`) handling.
* Error handling for invalid requests, missing files, and server errors.

---

## Authors

* Antenyi Joseph Ochohepo

---

## Requirements

* Go 1.20 or later

---

## Project Structure

```text
ascii-art-web/
│
├── bannerfiles/
│   ├── standard.txt
│   ├── shadow.txt
│   └── thinkertoy.txt
│
├── static/
│   └── css/
│       └── style.css
│
├── templates/
│   └── index.html
│
├── main.go
├── handlers.go
├── renderart.go
├── generateart.go
├── loadbanner.go
├── validateart.go
├── split.go
├── render_test.go
├── handler_test.go
├── go.mod
└── README.md
```

---

## Installation

Clone the repository:

```bash
git clone <repository-url>
```

Move into the project directory:

```bash
cd ascii-art-web
```

Install project dependencies:

```bash
go mod tidy
```

---

## Running the Application

Start the web server:

```bash
go run .
```

Open your browser and visit:

```text
http://localhost:8080
```

---

## Running the Tests

Run all automated unit tests:

```bash
go test -v
```

---

## Usage

1. Enter the text you want to convert.
2. Select one of the available banner styles.
3. Click **Generate ASCII Art**.
4. View the generated ASCII art on the page.
5. Click **Download** to save the generated ASCII art as a `.txt` file.

---

## HTTP Endpoints

### GET /

Displays the main page containing the input form.

**Response**

* **200 OK**

---

### POST /ascii-art

Processes the submitted text and selected banner.

**Response**

* **200 OK** – ASCII art generated successfully.
* **400 Bad Request** – Invalid input or unsupported banner.
* **404 Not Found** – Required template or banner file is missing.
* **405 Method Not Allowed** – Unsupported HTTP method.
* **500 Internal Server Error** – Unexpected server error.

---

### POST /download

Downloads the generated ASCII art as a text file.

**Response**

* **200 OK** – File downloaded successfully.
* **405 Method Not Allowed** – Unsupported HTTP method.

---

## Implementation Details

The application is divided into reusable components.

### main.go

* Starts the HTTP server.
* Registers all application routes.
* Serves static files (CSS).

### handlers.go

* Handles incoming HTTP requests.
* Validates request methods.
* Reads submitted form data.
* Processes ASCII art generation.
* Handles file downloads.
* Renders HTML templates.

### renderart.go

* Coordinates the ASCII generation process.
* Normalizes input line endings.
* Validates user input.
* Loads banner files.
* Calls the rendering engine.

### loadbanner.go

* Reads banner files.
* Builds a map of printable ASCII characters.

### validateart.go

* Ensures only printable ASCII characters are accepted.

### split.go

* Splits user input into individual lines.

### generateart.go

* Generates the final ASCII art output using the selected banner.

### render_test.go

* Tests the ASCII rendering functionality.
* Verifies valid rendering.
* Verifies invalid banner handling.
* Verifies multiline rendering.
* Verifies empty input behavior.

### handler_test.go

* Tests HTTP handlers.
* Verifies valid and invalid request methods.
* Tests HTTP status codes.
* Tests the download endpoint.

---

## Logging

The application records useful runtime information to assist with debugging and diagnostics, including:

* Server startup.
* Client IP address.
* User input.
* Selected banner.
* Successful ASCII art generation.
* Download requests.
* Error messages.

Example log output:

```text
2026/07/08 11:24:00 Client: 127.0.0.1:51528
2026/07/08 11:24:00 Input: "my studio art"
2026/07/08 11:24:00 Banner: standard
2026/07/08 11:24:00 ASCII art generated successfully
2026/07/08 11:24:05 Download requested by 127.0.0.1:51528
```

---

## Error Handling

The application validates:

* Unsupported HTTP methods.
* Invalid routes.
* Invalid banner names.
* Unsupported characters.
* Empty user input.
* Missing banner files.
* Missing HTML templates.

Appropriate HTTP status codes are returned for each situation.

---

## Technologies Used

* Go (Golang)
* HTML5
* CSS3
* Go HTML Templates
* `net/http`
* `html/template`
* Go `testing` package
* `net/http/httptest`
* Go `log` package

---

## License

This project was developed for educational purposes as part of the 01/edu Go curriculum.
