// // package main

// // import (
// // 	"encoding/json"
// // 	"fmt"
// // 	"io"
// // 	"log"
// // 	"net/http"
// // )

// type Post struct {
// 	UserId int `json:"userId"`
// 	Id int `json:"id"`
// 	Title string `json:"title"`
// 	Body string `json:"body"`

// }

// // func main() {
// //     url := "https://jsonplaceholder.typicode.com/posts/1"

// // 	res, err := http.Get(url)

// // 	if err != nil {
// // 		log.Fatal(err)
// // 	}
// //     defer res.Body.Close()

// // 	if res.StatusCode == http.StatusOK {
// // 		bodyBytes, err := io.ReadAll(res.Body)
// // 		if err != nil {
// // 		log.Fatal(err)
// // 		}

// // 		newPost := Post{}

//	// json.Unmarshal(bodyBytes, &newPost)

// // 		fmt.Println(newPost)
// // 	}

// // }

// /*

// same as:

// */

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"
// )

// type Post struct {
// 	UserId int `json:"userId"`
// 	Id int `json:"id"`
// 	Title string `json:"title"`
// 	Body string `json:"body"`

// }

// func main() {
//     url := "https://jsonplaceholder.typicode.com/posts/1"

// 	res, err := http.Get(url)

// 	if err != nil {
// 		log.Fatal(err)
// 	}
//     defer res.Body.Close()

// 	if res.StatusCode == http.StatusOK {

// 		newPost := Post{}

// 		json.NewDecoder(res.Body).Decode(&newPost)

// 		fmt.Println(newPost)

// 	}

// }

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Post struct {
	UserId int `json:"userId"`
	Id int `json:"id"`
	Title string `json:"title"`
	Body string `json:"body"`

}

func main () {
	url := "https://jsonplaceholder.typicode.com/posts/1"

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	if response.StatusCode == http.StatusOK {
		bodyBtyes, err := io.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		firstPost := Post{}
		json.Unmarshal(bodyBtyes, &firstPost)

		fmt.Printf("%T", response.Body) // output: *http.http2gzipReader
		fmt.Println(response.Body) // output: &{[] {0xc00024c180} 0xc0001a0008 <nil>}
		fmt.Println("####################")
		fmt.Println(bodyBtyes)
	}

	defer response.Body.Close()

}
















