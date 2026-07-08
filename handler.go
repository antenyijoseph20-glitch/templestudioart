package main

import (
	"html/template"
	"log"
	"net/http"
)

// PageData holds the data sent to the HTML template.
type PageData struct {
	Text   string
	Banner string
	Result string
	Error  string
}

// homeHandler serves the home page.
func homeHandler(w http.ResponseWriter, r *http.Request) {

	// Allow only the home route.
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Allow only GET requests.
	if r.Method != http.MethodGet {
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Load the template.
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "400 Template Not Found", http.StatusNotFound)
		return
	}

	// Show an empty page.
	data := PageData{
		Banner: "standard",
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// asciiArtHandler processes the submitted form.
func asciiArtHandler(w http.ResponseWriter, r *http.Request) {

	// Allow only POST requests.
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	log.Printf("Client: %s", r.RemoteAddr)

	// Read form values.
	text := r.FormValue("text")
	log.Printf("Input: %q", text)
	// if text == "" {
	// 	http.Error(w, "400 Bad Request: empty input", http.StatusBadRequest)
	// 	return
	// }
	if text == "" {
		log.Println("Empty input received")
		http.Error(w, "400 Bad Request: empty input", http.StatusBadRequest)
		return
	}

	banner := r.FormValue("banner")
	if banner == "" {
		banner = "standard"
	}
	log.Printf("Banner: %s", banner)

	switch banner {
	case "standard", "shadow", "thinkertoy":
		// Valid banner.
	default:
		http.Error(w, "400 Bad Request: invalid banner", http.StatusBadRequest)
		return
	}

	// Build banner file path.
	bannerFile := "bannerfiles/" + banner + ".txt"

	// Generate ASCII art.
	result, err := renderArt(text, bannerFile)
	if err != nil {
		log.Printf("Render error: %v", err)
	} else {
		log.Printf("ASCII art generated successfully")
	}

	data := PageData{
		Text:   text,
		Banner: banner,
	}

	if err != nil {
		data.Error = err.Error()
	} else {
		data.Result = result
	}

	// Load template.
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	// Display the page.
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// log.Println("Generating ASCII Art")
	// log.Printf("Banner selected: %s", banner)
}

func DownloadHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	//ascii := r.FormValue("ascii")
	ascii := r.FormValue("ascii")

	log.Printf("Download requested by %s", r.RemoteAddr)
	log.Printf("Downloaded file size: %d bytes", len(ascii))

	w.Header().Set("Content-Disposition", `attachment; filename="ascii-art.txt"`)
	w.Header().Set("Content-Type", "text/plain")

	w.Write([]byte(ascii))
}
