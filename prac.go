package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
)

func main() {

	var urls []string = []string{}
	for i:= 1; i <=100; i++{
		urls = append(urls, "https://jsonplaceholder.typicode.com/posts/"+strconv.Itoa(i))	
	}
	
	var wg sync.WaitGroup

	for _, url := range urls{
		wg.Add(1)
		go func (url string)  {
			Get_data(url)	
			wg.Done()		
		}(url)
	}
	wg.Wait()

}


func Get_data(link string) {
	
	content, err := http.Get(link)
	if err != nil{
		fmt.Println("Error")
	}
	data, err := ioutil.ReadAll(content.Body)
	if err != nil{
		fmt.Println("Error")
	}
	defer content.Body.Close()

	fmt.Println(string(data))

}