package main

import (
	"fmt"

	"github.com/bootcamp-go/desafio-go-bases/internal/tickets"
)

func main() {

	datos, _ := tickets.GetArray()
	total := tickets.GetTotalTickets("China", datos)

	fmt.Println(total)

}
