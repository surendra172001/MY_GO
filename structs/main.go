package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	narendra := person{
		firstName: "Narendra",
		lastName:  "Pandey",
		contact: contactInfo{
			email:   "narendra05102010@gmail.com",
			zipCode: 400101,
		},
	}
	fmt.Printf("%+v\n", narendra)
	s := "surendra"
	fmt.Println(*&s)
}
