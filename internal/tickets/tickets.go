package tickets

import (
	"fmt"
	"os"
	"strings"
)

type Ticket struct {
}

func GetArray() ([][]string, error) {

	file, err := os.ReadFile("./tickets.csv")

	if err != nil {
		return nil, err
	}

	arrStrAux := strings.Split(string(file), ";")

	var arrStr [][]string

	for _, v := range arrStrAux {

		line := strings.Split(v, ",")

		fmt.Printf("%s\n\n", line)

		arrStr = append(arrStr, line)

	}

	return arrStr, nil
}

// ejemplo 1
func GetTotalTickets(destination string, datos [][]string) int {

	// Una función que calcule cuántas personas viajan a un país determinado.

	count := 0

	for i, ticket := range datos {

		fmt.Println(i)
		fmt.Println(ticket)

		// if v[3] == destination {
		// 	count += 1
		// }

	}

	return count
}

// // ejemplo 2
// func GetMornings(time string) (int, error) {}

// // ejemplo 3
// func AverageDestination(destination string, total int) (int, error) {}
