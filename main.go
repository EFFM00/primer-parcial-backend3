package main

import (
	"fmt"

	"github.com/EFFM00/primer-parcial-backend3/internal/tickets"
)

func main() {

	datos, _ := tickets.GetArray()
	total := tickets.GetTotalTickets("China", datos)

	fmt.Println(total)

}
