package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	// Getrequest()
	// Postrequest()
	PerformPostForm()
}

func Getrequest() {
	const myurl = "http://localhost:8000/get"
	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Hello")
		content, err := ioutil.ReadAll(response.Body)
		if err != nil {
			panic(err)
		} else {
			fmt.Println(string(content))
		}
	}
	response.Body.Close()
}

func Postrequest() {
	const myurl = "http://localhost:8000/post"

	requestBody := strings.NewReader(`
	    {
			"coursename" : "golang",
			"price" : "123",
			"place" : "youtube"
		}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	} else {
		defer response.Body.Close()
		content, err := ioutil.ReadAll(response.Body)

		if err != nil {
			panic(err)
		} else {
			fmt.Println(string(content))
		}
	}
}

func PerformPostForm() {
	const myurl = "http://localhost:8000/postform"

	//formdata
	data := url.Values{}
	data.Add("firstname", "shrut")
	data.Add("lastname", "aggarwal")
	data.Add("email", "aggarwal.4@iitj.ac.in")
	response, err := http.PostForm(myurl, data)

	if err != nil {
		fmt.Println("1")
		panic(err)
	} else {
		defer response.Body.Close()
		content, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("2")

			panic(err)
		} else {
			fmt.Println(string(content))
		}
	}

}
