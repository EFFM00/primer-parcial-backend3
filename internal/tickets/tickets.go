package tickets

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// estrutura que recibirá los datos del csv
type Ticket struct {
	Id          int     `json:"ID"`
	Name        string  `json:"Name"`
	Email       string  `json:"Email"`
	Destination string  `json:"Destination"`
	Departure   string  `json:"Departure"`
	Price       float64 `json:"Price"`
}

// estrutura que recibirá los datos y
// guardará una lista de tickets vendidos por
// la compañía en las últimas 24 hours
type Airline struct {
	ID      string   `json:"ID"`
	Name    string   `json:"Nombre"`
	Tickets []Ticket `json:"Tickets"`
}

// manejo de error en lista vacía
var ErrNoTickets = errors.New("Empty list")

// ---------------------------------------------------------------
// Requisito 01:
// Una función que calcule cuántas personas viajan a un país determinado.
func (a Airline) GetTotalTickets(destination string) (int, error) {

	if len(a.Tickets) == 0 {
		return 0, ErrNoTickets
	}

	count := 0
	for _, ticket := range a.Tickets {
		if strings.ToUpper(ticket.Destination) == strings.ToUpper(destination) {
			count++
		}
	}

	return count, nil
}

// ---------------------------------------------------------------
// Requisito 02:
// Calcular cuántas personas viajan en Dawn (0 → 6), Morning (7 → 12), Afternoon (13 → 19) y Night (20 → 23).

type flighShift uint

var flighShifts = [...]string{"Dawn", "Morning", "Afternoon", "Night"}

const (
	Dawn flighShift = iota //asigna valores secuenciales comenzando desde cero.
	Morning
	Afternoon
	Night
)

func (f flighShift) String() string {
	return flighShifts[f]
}

func (f flighShift) Num() string {
	switch f {
	case Dawn:
		return "0"
	case Morning:
		return "1"
	case Afternoon:
		return "2"
	case Night:
		return "3"
	}
	return ""
}

// manejo de error en el turno
var ErrTimeFormat = errors.New(`Enter a valid time format.`)

// Se valida el parámetro de turno y se lo convierte en un formato de tiempo.
// Luego, se  cuenta el número de tickets vendidos para el horario especificado.
func (a Airline) GetCountByPeriod(time string) (int, error) {
	var flighShift string

	// parametro de tiempo
	if time != Dawn.String() && time != Morning.String() && time != Afternoon.String() && time != Night.String() && time != Dawn.Num() && time != Morning.Num() && time != Afternoon.Num() && time != Night.Num() {
		if strings.Contains(time, ":") {

			format := regexp.MustCompile(`\d\d:\d\d`)
			match := format.MatchString(time)

			if !match {
				return 0, ErrTimeFormat
			}
			parts := strings.SplitN(time, ":", 2)
			var hours, minutes int
			var err error
			hours, err = strconv.Atoi(parts[0])
			if err != nil {
				return 0, err
			}
			minutes, err = strconv.Atoi(parts[1])
			if err != nil {
				return 0, err
			}

			if hours < 0 || hours > 23 || minutes > 59 {
				return 0, ErrTimeFormat
			}
			if hours >= 0 && hours <= 6 {
				flighShift = "Dawn"
			} else if hours >= 7 && hours <= 12 {
				flighShift = "Morning"
			} else if hours >= 13 && hours <= 19 {
				flighShift = "Afternoon"
			} else {
				flighShift = "Night"
			}
		} else {
			num, err := strconv.ParseInt(time, 10, 0)
			if err != nil {
				return 0, err
			}
			if num > 3 {
				return 0, ErrTimeFormat
			}
			flighShift = time
		}
	} else {
		flighShift = time
	}

	count := 0
	for _, v := range a.Tickets {
		departureAux := strings.SplitN(v.Departure, ":", 2)
		hoursAux, err := strconv.Atoi(departureAux[0])
		if err != nil {
			return 0, err
		}
		if (hoursAux >= 0 && hoursAux <= 6) && (flighShift == "Dawn" || time == Dawn.Num()) {
			count++
		} else if (hoursAux >= 7 && hoursAux <= 12) && (flighShift == "Morning" || time == Morning.Num()) {
			count++
		} else if (hoursAux >= 13 && hoursAux <= 19) && (flighShift == "Afternoon" || time == Afternoon.Num()) {
			count++
		} else if (hoursAux >= 20 && hoursAux <= 23) && (flighShift == "Night" || time == Night.Num()) {
			count++
		}
	}

	return count, nil

}

// ---------------------------------------------------------------
// Requisito 03:
//

func (a Airline) AverageDestination(destination string) (float64, error) {

	totalPassengers := len(a.Tickets)

	passengersToDestin, err := a.GetTotalTickets(destination)

	if err != nil {
		return 0.0, err
	}

	percentage := (float64(passengersToDestin) / float64(totalPassengers)) * 100

	return float64(percentage), nil
}

// --------------------
// abre un archivo CSV en la ruta especificada y lee su contenido para crear una lista de objetos Ticket.
func OpenCSV(path string) ([]Ticket, error) {

	f, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer f.Close()

	r := csv.NewReader(f)

	var ticketsList []Ticket

	for {
		record, err := r.Read()

		if err == io.EOF {

			fmt.Printf("Se ha llegado al final del archivo: %s \n", err)
			break
		}

		id, err := strconv.Atoi(record[0])

		price, err := strconv.ParseFloat(record[5], 64)

		if err != nil {
			return nil, err

		}

		ticketsList = append(ticketsList, Ticket{
			Id:          id,
			Name:        record[1],
			Email:       record[2],
			Destination: record[3],
			Departure:   record[4],
			Price:       price,
		})
	}

	return ticketsList, nil
}
