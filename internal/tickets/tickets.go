package tickets

import (
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

	str := string(file)

	arrStrAux := strings.Split(str, ";")

	var arrStr [][]string

	for _, v := range arrStrAux {

		line := strings.Split(v, ",")

		arrStr = append(arrStr, line)

	}

	return arrStr, nil
}

// ejemplo 1
// func GetTotalTickets(destination string) (string, error) {

// Una función que calcule cuántas personas viajan a un país determinado.

// }

// // ejemplo 2
// func GetMornings(time string) (int, error) {}

// // ejemplo 3
// func AverageDestination(destination string, total int) (int, error) {}
