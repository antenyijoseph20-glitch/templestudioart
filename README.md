# ASCII Art Web

## Description

ASCII Art Web is a web-based implementation of the ASCII Art project written in Go. It allows users to enter text through a web interface, choose one of the available banner styles, and generate the corresponding ASCII art directly in the browser.

The application reuses the ASCII rendering engine developed in the previous CLI version while providing a simple and user-friendly web interface.

---

## Features

* Generate ASCII art from user input.
* Supports three banner styles:

  * **standard**
  * **shadow**
  * **thinkertoy**
* Clean web interface built with HTML templates.
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

```
ascii-art-web/
│
├── bannerfiles/
│   ├── standard.txt
│   ├── shadow.txt
│   └── thinkertoy.txt
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

---

## Running the Application

Start the server:

```bash
go run .
```

Open your browser and visit:

```
http://localhost:8080
```

---

## Usage

1. Enter the text you want to convert.
2. Select one of the available banners.
3. Click **Generate ASCII Art**.
4. The generated ASCII art will be displayed on the same page.

---

## HTTP Endpoints

### GET /

Displays the main page containing the input form.

Response:

* **200 OK**

---

### POST /ascii-art

Processes the submitted text and selected banner.

Response:

* **200 OK** – ASCII art generated successfully.
* **400 Bad Request** – Invalid input or unsupported banner.
* **404 Not Found** – Required template or banner file is missing.
* **405 Method Not Allowed** – Unsupported HTTP method.
* **500 Internal Server Error** – Unexpected server error.

---

## Implementation Details

The application is divided into small, reusable components.

### main.go

* Starts the HTTP server.
* Registers all routes.

### handlers.go

* Handles incoming HTTP requests.
* Validates request methods.
* Reads form data.
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

* Go
* HTML
* Go HTML Templates
* net/http

---

## Future Improvements

* Add automated unit tests.
* Improve the visual design using CSS.
* Add download functionality for generated ASCII art.
* Improve logging and diagnostics.

---

## License

This project was developed for educational purposes as part of the 01/edu Go curriculum.

