package main

import (
	"fmt"
	"strings"
)

type address struct {
	street, city, state string
	zip                 int
}

type Employee struct {
	home address
	work address
}

func (emp *Employee) toString() string {
	var str strings.Builder
	str.WriteString("Home Address : " + emp.home.street)
	str.WriteString("\n Work Address : " + emp.work.street)
	return str.String()
}

func main() {

	fmt.Println("Strunct demo...")
	emp := &Employee{home: address{"street Test1", "city Test1", "tx", 75063}}
	emp.work = address{"work street", "city Test1", "tx", 75021}
	fmt.Println("Emp : ", (*emp).toString())
}
