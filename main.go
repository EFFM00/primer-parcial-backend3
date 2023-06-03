package main

import (
	"fmt"

	"github.com/EFFM00/primer-parcial-backend3/internal/tickets"
	"github.com/google/uuid" //Esta biblioteca permite generar identificadores únicos para objetos
)

var Airline tickets.Airline

func main() {
	ticketsList, err := tickets.OpenCSV("./tickets.csv")
	if err != nil {
		fmt.Println(err)
	}

	Airline = tickets.Airline{
		ID:      uuid.New().String(),
		Name:    "Aerolineas Argentinas",
		Tickets: ticketsList,
	}

	//--------------------------------------------------------------
	//Requisito 1

	channelTotalCountry := make(chan int)

	go func() {
		for {
			totalCountry, err := Airline.GetTotalTickets("Brazil")

			if err != nil {
				fmt.Println(err)
			}

			channelTotalCountry <- totalCountry

		}
	}()

	fmt.Println("--------------------------------------------------------------")

	totalCountryResult := <-channelTotalCountry
	fmt.Println("Total de tickets para Brasil: ", totalCountryResult)

	//--------------------------------------------------------------
	// Requisito 2

	channelTicketsDawn := make(chan int)
	channelTicketsMorning := make(chan int)
	channelTicketsAfternoon := make(chan int)
	channelTicketsNight := make(chan int)

	go func() {
		for {
			ticketsDawn, err := Airline.GetCountByPeriod("0")

			if err != nil {
				fmt.Println(err)
			}

			channelTicketsDawn <- ticketsDawn

		}
	}()

	ticketsDawnResult := <-channelTicketsDawn

	go func() {
		for {
			ticketsMorning, err := Airline.GetCountByPeriod("Morning")

			if err != nil {
				fmt.Println(err)
			}

			channelTicketsMorning <- ticketsMorning
		}
	}()

	ticketsMorningResult := <-channelTicketsMorning

	go func() {
		for {
			ticketsAfternoon, err := Airline.GetCountByPeriod("14:00")

			if err != nil {
				fmt.Println(err)
			}

			channelTicketsAfternoon <- ticketsAfternoon
		}
	}()

	ticketsAfternoonResult := <-channelTicketsAfternoon

	go func() {
		for {
			ticketsNight, err := Airline.GetCountByPeriod("3")

			if err != nil {
				fmt.Println(err)
			}

			channelTicketsNight <- ticketsNight

		}
	}()

	ticketsNightResult := <-channelTicketsNight

	fmt.Println("--------------------------------------------------------------")
	fmt.Println("Personas que viajan en la madrugada:", ticketsDawnResult)
	fmt.Println("Personas que viajan en la mañana:", ticketsMorningResult)
	fmt.Println("Personas que viajan en la tarde:", ticketsAfternoonResult)
	fmt.Println("Personas que viajan en la noche:", ticketsNightResult)
	//--------------------------------------------------------------

	totalTickets := ticketsDawnResult + ticketsMorningResult + ticketsAfternoonResult + ticketsNightResult

	fmt.Println("--------------------------------------------------------------")
	fmt.Println("Total de tickets en un día:", totalTickets)

	//--------------------------------------------------------------
	// Requisito 3

	channelPercentagePassengers := make(chan float64)

	go func() {
		for {
			percentagePassengers, err := Airline.AverageDestination("Brazil")

			if err != nil {
				fmt.Println(err)
			}

			channelPercentagePassengers <- percentagePassengers
		}
	}()

	percentagePassengersResult := <-channelPercentagePassengers

	fmt.Println("--------------------------------------------------------------")
	fmt.Printf("Porcentaje de pasajeros hacia Brasil: %0.2f \n", percentagePassengersResult)
	fmt.Println("--------------------------------------------------------------")

	//--------------------------------------------------------------
	// Requisito 4
	// Implementado a lo largo del proyecto con la utilización de goroutines y canales
}
