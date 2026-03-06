package main

type Preson struct {
	Name    string
	Age     int
	Married bool
	Hobbies []string
	Address Address
}

type Address struct {
	Country string
	City    string
}
