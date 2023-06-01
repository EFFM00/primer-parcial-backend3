package tickets

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Ticket struct {
	id          string
	name        string
	email       string
	destination string
	timeFlight  string
	price       float64
}

func GetArray() ([]Ticket, error) {

	file, err := os.ReadFile("./tickets.csv")

	if err != nil {
		return nil, err
	}

	arrStrAux := strings.Split(string(file), "\n")

	var arrStr []Ticket

	for _, v := range arrStrAux {

		line := strings.Split(v, ",")

		id := line[0]

		name := line[1]

		email := line[2]

		destination := line[3]

		timeFlight := line[4]

		price, err := strconv.ParseFloat(line[5], 64)
		if err != nil {
			fmt.Println("error parsing the price")
			continue
		}

		newTicket := Ticket{
			id:          id,
			name:        name,
			email:       email,
			destination: destination,
			price:       price,
			timeFlight:  timeFlight,
		}

		arrStr = append(arrStr, newTicket)

	}

	return arrStr, nil
}

func GetTotalTickets(destination string, datos []Ticket) int {

	count := 0

	for _, ticket := range datos {

		if ticket.destination == destination {
			count += 1
		}

	}

	return count
}

// // ejemplo 2
// func GetMornings(time string) (int, error) {}

func AverageDestination(destination string, datos []Ticket) float64 {

	total := float64(len(datos))
	passengersToDestin := float64(GetTotalTickets(destination, datos))

	return (passengersToDestin * 100) / total

}
