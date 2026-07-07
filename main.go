// package main

// import(
// 	"fmt"
// 	"os"
// )

// func main(){
// 	if len(os.Args) < 2|| len(os.Args)> 3{
// 		fmt.Println("Usage: go run . <text> <bannerfiles>")
// 		return
// 	}
// 	input := os.Args[1]

// 	bannerFile := "bannerfiles/standard.txt"
// 	result, err := renderArt(input, bannerFile)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Print(result)

// }
package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", homeHandler)

	http.HandleFunc("/ascii-art", asciiArtHandler)

	fmt.Println("Server running on http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
