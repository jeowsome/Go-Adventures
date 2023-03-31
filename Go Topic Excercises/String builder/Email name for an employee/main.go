package main

import (
	"fmt"
	"log"
	"reflect"
)

// Write the code to concatenate `firstName`, `lastName`,
// and `domain` into an email address using the strings.Builder
func createEmail(firstName, lastName, domain string) strings.Builder {
	var builder strings.Builder
	builder.WriteByte(firstName[0])
	builder.WriteString(lastName)
	builder.WriteString(domain)
	return builder
}

// DO NOT delete or modify the code within the main() function!
func main() {
	var firstName, lastName, domain string
	fmt.Scanln(&firstName, &lastName, &domain)

	email := createEmail(firstName, lastName, domain)
	if !isBuilder(email) {
		log.Fatal("Wrong! the createEmail function must return a strings.Builder type")
	}
	fmt.Println(email.String())
}
