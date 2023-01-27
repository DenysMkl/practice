package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

func main() {
	urls := []string{}
	for i:= 1; i <= 10; i++{
		urls = append(urls,  "https://jsonplaceholder.typicode.com/posts/"+strconv.Itoa(i))
	}

	err := os.MkdirAll("storage/posts", 0750)
	if err != nil{
		fmt.Println("error")
	}


	channel := make(chan string)
	
	for _, url := range urls{
		go func (url string, channel chan string)  {
			channel <- Get_data(url)
		}(url, channel)
	
	}

	var counter int = 1
	for i := range channel{
		file, _ := os.Create("storage/posts/"+strconv.Itoa(counter)+".txt")
		counter+=1
		file.WriteString(i)
		file.Close()
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