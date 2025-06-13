package main

import "fmt"

type Person struct {
	Name      string
	Surname   string
	Email     string
	Zip       string
	City      string
	Number    string
	Relations map[string]string
}

var peopleBook = make(map[string]Person)

func addPerson() {
	var name, surname, email, zip, city, number string

	fmt.Print("Enter name: ")
	fmt.Scanf("%s\n", &name)

	fmt.Print("Enter surname: ")
	fmt.Scanf("%s\n", &surname)

	fmt.Print("Enter email: ")
	fmt.Scanf("%s\n", &email)

	fmt.Print("Enter ZIP code: ")
	fmt.Scanf("%s\n", &zip)

	fmt.Print("Enter city: ")
	fmt.Scanf("%s\n", &city)

	fmt.Print("Enter phone number: ")
	fmt.Scanf("%s\n", &number)

	person := Person{
		Name:    name,
		Surname: surname,

		Email:     email,
		Zip:       zip,
		City:      city,
		Number:    number,
		Relations: make(map[string]string),
	}

	for {
		var relType string
		fmt.Print("Enter relationship type (or type 'done' to finish): ")
		fmt.Scanf("%s\n", &relType)

		if relType == "done" {
			break
		}

		var relName string
		fmt.Printf("Enter name of the related person for '%s': ", relType)
		fmt.Scanf("%s\n", &relName)

		person.Relations[relType] = relName
	}

	peopleBook[name] = person
	fmt.Println("Person added successfully!")
}

func viewPeople() {
	for _, person := range peopleBook {
		fmt.Printf("\nName: %s %s\n", person.Name, person.Surname)
		fmt.Printf("Email: %s\n", person.Email)
		fmt.Printf("City: %s, ZIP: %s, Phone: %s\n", person.City, person.Zip, person.Number)
		fmt.Println("Relationships:")
		for rel, name := range person.Relations {
			fmt.Printf("  %s: %s\n", rel, name)
		}
	}
}

func addressBook() {
	for {
		fmt.Println("\n--- People Book ---")
		fmt.Println("1. Add Person")
		fmt.Println("2. View All People")
		fmt.Println("3. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanf("%d\n", &choice)

		switch choice {
		case 1:
			addPerson()
		case 2:
			viewPeople()
		case 3:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice, try again.")
		}
	}
}
