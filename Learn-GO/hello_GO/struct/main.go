package main

import "fmt"

type Person struct {
	FristName string
	LastName  string
	Age       int
}

type Contact struct {
	Email   string
	PhoneNo int
}

type Address struct {
	House int
	Area  string
	State string
}

type Employ struct {
	Employ_detail Person
	Employ_Contact Contact
	Employ_Address Address
}

func main() {
	// var Om Person

	// Om.FristName = "Om"
	// Om.LastName = "SHA"
	// Om.Age = 21

	// fmt.Println("Person Om", Om)

	// // Second Method
	// Person1 := Person{
	// 	FristName: "Akash",
	// 	LastName:  "Bikash",
	// 	Age:       23,
	// }

	// fmt.Println("Person1 details : ", Person1)

	// // Third Method (new Keyword)
	// var person2 = new(Person)
	// person2.FristName = "Agarwal"
	// person2.LastName = "Singh"
	// person2.Age = 44

	// fmt.Println("Person2 with new keyword: ", person2)
	//
	Employ1 := Employ {
		Employ_detail: Person{
			FristName: "Akash",
		 	LastName: "Bikash",
		 	Age: 23,
		},
		// Employ_Contact: Contact{
		// 	Email: "akash@example.com",
		// 	PhoneNo: 123123,
		// },
		Employ_Address: Address{
			Area: "Moti Mahal, Rajasthan",
			House: 112,
			State: "Rajasthan",
		},
	}
	Employ1.Employ_Contact.Email = "example@example.com"
	Employ1.Employ_Contact.PhoneNo = 123123
	
	fmt.Println("Employ details:", Employ1)
	fmt.Println("Employ house: ", Employ1.Employ_Address.House)
}
