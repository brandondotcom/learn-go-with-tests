package util

type Address struct {
	streetNumber int
	streetName   string
}

type Person struct {
	name       string
	address    Address
	addressRef *Address
}
