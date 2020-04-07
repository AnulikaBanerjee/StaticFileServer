package main
//any executable code should have package name main

import (
	"net/http"
	"os"
	// In go, unused package imports are not allowed. It will throw an error
)

func main(){
	//get the directory you are at currently:
	dir,_ := os.Getwd() //lot like pwd command. This function returns a second parameter for err
	http.ListenAndServe(":3030",http.FileServer(http.Dir(dir))) 
	// here we are pointin to the path where the code lies.
	/*ListenAndServe starts an HTTP server with a given address and handler.
		The handler is usually nil, which means to use DefaultServeMux.
		  Creating custom server:
		  s := &http.Server{
			Addr:           ":8080",
			Handler:        myHandler,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}
		log.Fatal(s.ListenAndServe())

	FileServer returns a handler that serves HTTP requests with the contents of the file system rooted at root.
	*/

}