package main

import (
	"fmt"
	"log"
	"net/http"
)


func thankuHandler (res http.ResponseWriter, req *http.Request){
	
	if req.URL.Path != "/thank-u"{
		http.Error(res, "404 not found", 404)
		return
	}
	if req.Method != "GET"{
		http.Error(res, "Unavaible  for this method", 405)
		return
	}

	fmt.Fprint(res, "thank you")


}



func formHandler (res http.ResponseWriter, req *http.Request){
	if req.URL.Path != "/form"{
		http.Error(res, "404 not found", 404)
		return
	}
	// if req.Method != "POST"{
	// 	http.Error(res, "Unavaible  for this method", 405)
	// 	return
	// }

	fmt.Fprint(res, "form submitted succc")
	
	name :=  req.FormValue("name")
	email := req.FormValue("email")
	
	fmt.Fprintf(res, "Name: %s \n" , name)
	fmt.Fprintf(res, "Email: %s \n", email)




}

func main (){


	server := http.FileServer(http.Dir("./static"))
	http.Handle( "/", server )
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/thank-u", thankuHandler)


	err := http.ListenAndServe(":8081", nil)
	if  err != nil {
		log.Fatal("Could not start the web server : ",err)
		} else{
			fmt.Println("Web Server is running on port 8081")
	}
}