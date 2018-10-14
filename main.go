package main

type contactInfo struct {
	phone string
	address string
}

type person struct {
	firstName string
	lastName  string
	contact contactInfo
}

func main() {
	// jone := person{"Jone", "Doe"}

	// jone := person{firstName: "Jone", lastName: "Doe"}

	// var jone person

	// jone.firstName = "Jone"
	// jone.lastName = "Doe"
	// jone.contact.phone = "01096901954"
	// jone.contact.address = "imbaba"
	// fmt.Println(jone)
	// fmt.Printf("%+v", jone)
}
