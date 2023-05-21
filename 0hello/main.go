package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	// reader := bufio.NewReader(os.Stdin)
	// input, _ := reader.ReadString('\n')

	// fmt.Println(input)
	// numberrating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	// if err != nil {
	// 	fmt.Println("yes", err)
	// } else {
	// 	numberrating += 1
	// 	fmt.Println((numberrating))
	// }
	// presentTime := time.Now()
	// fmt.Println(presentTime)
	// var ptr *int
	// fmt.Println(ptr)
	// mynuber := 23
	// var ptr *int = &mynuber
	// fmt.Println(*ptr)
	// *ptr *= 2
	// fmt.Println(*ptr)
	// fmt.Println(mynuber)

	// var arr [4]int
	// arr[0] = 1
	// fmt.Println(arr)

	// var numbers = [3]int{3, 4, 5}
	// fmt.Println(numbers)

	// var arr = []string{}
	// fmt.Printf("%T", arr)

	// arr = append(arr, "mangoes", "apples")
	// fmt.Println(arr)

	// arr1 := make([]int, 4)
	// arr1[0] = 1
	// arr1[1] = 2
	// arr1[2] = 6
	// arr1[3] = 4

	// arr1 = append(arr1, 5)
	// fmt.Println(arr1)

	// sort.Ints(arr1)
	// fmt.Println(arr1)

	// lauguages := make(map[string]string)

	// lauguages["JS"] = "Javascript"
	// lauguages["py"] = "python"
	// lauguages["rb"] = "ruby"

	// fmt.Println(lauguages["JS"])

	// delete(lauguages, "JS")
	// fmt.Println(lauguages)

	// for key, _ := range lauguages {
	// 	fmt.Println(key, lauguages[key])
	// }

	// shrutayu := User{"shrutayu", "aggarwal.4@iitj.ac.in", true, 20}
	// fmt.Println(shrutayu.Email)
	// fmt.Printf("%+v \n", shrutayu)
	// fmt.Println(shrutayu.GetStatus())
	// shrutayu.NewMail()
	// fmt.Println(shrutayu.Email)

	// rand.Seed(time.Now().Unix())
	// number := rand.Intn(6) + 1
	// fmt.Println(number)

	// switch number {
	// case 1:
	// 	fmt.Println("One")
	// 	fallthrough
	// case 2:
	// 	fmt.Println("two")
	// 	fallthrough
	// case 3:
	// 	fmt.Println("three")
	// default:
	// 	fmt.Println(number)

	// }
	// defer fmt.Println("World")
	// defer fmt.Println("again")
	// defer fmt.Println("again2")

	// fmt.Println("Hello")
	// letsee()

	content := "My first string"
	file, err := os.Create("./file.txt")
	fmt.Printf("%T", file)
	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}

	fmt.Println(length)
	file.Close()
	read("./file.txt")
}

func read(file string) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(string(content))
	}
}

// func letsee() {
// 	for i := 0; i < 5; i++ {
// 		defer fmt.Println(i)
// 	}
// }

// type User struct {
// 	Name     string
// 	Email    string
// 	Verified bool
// 	Age      int
// }

// func (u User) GetStatus() bool {
// 	return u.Verified
// }

// func (u User) NewMail() {
// 	u.Email = "hello@ok.in"
// }
