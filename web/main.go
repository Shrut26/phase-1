package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://lco.dev:3000/learn?coursename=reactjs&paymentId=lsdgbaskn329"

func main() {

	// response, err := http.Get(url)
	// if err != nil {
	// 	panic(err)
	// } else {
	// 	fmt.Printf("%T \n", response)
	// 	fmt.Println(response)
	// 	defer response.Body.Close()
	// 	data, err := ioutil.ReadAll(response.Body)
	// 	if err != nil {
	// 		panic(err)
	// 	} else {
	// 		fmt.Println(string(data))
	// 	}
	// }

	//parsing the url
	result, err := url.Parse(myurl)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(result.RawQuery)
		qparams := result.Query()
		fmt.Println(qparams)
	}

	//constructing a URL
	newurl := &url.URL{
		Scheme:   "http",
		Host:     "github.io",
		Path:     "repo",
		RawQuery: "user=shrut&payment=0374013",
	}
	anotherurl := newurl.String()
	fmt.Println(anotherurl)
}
