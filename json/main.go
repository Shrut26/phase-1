package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"courseprice"`
	Platform string   `json:"courseplatform"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	// EncodeJson()
	decodeJson()
}

func EncodeJson() {
	lcocourses := []course{
		{"ReactJS", 299, "youtube", "1234567890", []string{"web-dev", "js"}},
		{"Mern", 199, "youtube1", "18651865", []string{"mongo", "nodejs"}},
		{"pern", 399, "youtube2", "46841684553", nil},
	}

	finalJson, err := json.MarshalIndent(lcocourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

func decodeJson() {
	jsonDatafromweb := []byte(`
	{
                "coursename": "ReactJS",
                "courseprice": 299,
                "courseplatform": "youtube",
                "tags": [
                        "web-dev","js"
                ]
        }
	`)

	var lcoCourse course

	isValid := json.Valid(jsonDatafromweb)

	if isValid {
		fmt.Println("Valid")
		json.Unmarshal(jsonDatafromweb, &lcoCourse)
		fmt.Printf("%#v \n", lcoCourse)
	} else {
		fmt.Println("Invalid Json")
	}

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDatafromweb, &myOnlineData)
	fmt.Printf("%#v \n", myOnlineData)

	for key := range myOnlineData {
		fmt.Println(key, myOnlineData[key])
	}

}
