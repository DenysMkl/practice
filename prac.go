package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main(){
	
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil{
		fmt.Println("Error")
	}
	body, err := ioutil.ReadAll(resp.Body)
	sb := string(body)
	fmt.Println(sb)
	
}

