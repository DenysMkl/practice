package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func main() {

	var urls []string = []string{}
	for i:= 1; i <=100; i++{
		urls = append(urls, "https://jsonplaceholder.typicode.com/posts/"+strconv.Itoa(i))	
	}
	channel := make(chan string)
	

	for _, url := range urls{
		
		go func (url string, channel chan string)  {
			channel <- Get_data(url)

		}(url, channel)
		for i := range channel{
			fmt.Println(i)
		}
	}
	

}


func Get_data(link string) string{
	
	content, err := http.Get(link)
	if err != nil{
		fmt.Println("Error")
	}
	data, err := ioutil.ReadAll(content.Body)
	if err != nil{
		fmt.Println("Error")
	}
	defer content.Body.Close()

	return string(data)

}