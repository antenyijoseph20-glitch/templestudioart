// // package main

// // import(
// // 	"fmt"
// // 	"os"
// // )

// // func main(){
// // 	if len(os.Args) < 2|| len(os.Args)> 3{
// // 		fmt.Println("Usage: go run . <text> <bannerfiles>")
// // 		return
// // 	}
// // 	input := os.Args[1]

// // 	bannerFile := "bannerfiles/standard.txt"
// // 	result, err := renderArt(input, bannerFile)
// // 	if err != nil {
// // 		fmt.Println(err)
// // 		return
// // 	}
// // 	fmt.Print(result)

// // }
package main

import (
	//"fmt"
	"log"
	"net/http"
)

func main() {
	http.Handle(
		"/static/",
		http.StripPrefix(
			"/static/",
			http.FileServer(http.Dir("static")),
		),
	)

	http.HandleFunc("/", homeHandler)

	http.HandleFunc("/ascii-art", asciiArtHandler)

	http.HandleFunc("/download", DownloadHandler)

	log.Println("Server running on http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println(err)
	}
}

// func DownloadHandler(w http.ResponseWriter, r *http.Request) {

//     if r.Method != http.MethodPost {
//         http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
//         return
//     }

//     ascii := r.FormValue("ascii")

//     w.Header().Set("Content-Disposition", `attachment; filename="ascii-art.txt"`)
//     w.Header().Set("Content-Type", "text/plain")

//     w.Write([]byte(ascii))
// }
