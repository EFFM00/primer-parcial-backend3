package main

import (
	"fmt"

	"github.com/EFFM00/primer-parcial-backend3/internal/tickets"
)

func main() {

	datos, _ := tickets.GetArray()
	total := tickets.GetTotalTickets("Brazil", datos)

	percPassengersByDestin := tickets.AverageDestination("Brazil", datos)

	fmt.Println(total)
	fmt.Printf("%0.2f perc", percPassengersByDestin)

}
