package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://instagram.com.com",
		"http://strackoverflow.com",
		"http://twitter.com",
	}

	c := make(chan string) //channel to communicate with go routine and manage diffrent go routines

	for _, given := range links {
		go checkLink(given, c) //using go routine if yu just add go keyword it wont work
	}
	
	for  {
		go checkLink(<-c,c)
	}
}



func checkLink(link string, c chan string) {     //c is a channel type string 
	_, err :=http.Get(link)
	 if err!= nil{
		fmt.Println(link,"might be down!")
		c <- link
		return
	}

	fmt.Println(link,"is up!")
	c <- link
}