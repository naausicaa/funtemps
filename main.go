package main

import (
	"flag"
	"fmt"

	"github.com/naausicaa/funtemps/conv"
	"github.com/naausicaa/funtemps/funfacts"
)

// Definerer flag-variablene i hoved-"scope"
var (
	fahr    float64
	kelvin  float64
	celsius float64
	out     string
	funf    string
	t       string
)

// Bruker init (som anbefalt i dokumentasjonen) for å sikre at flagvariablene
// er initialisert.

func init() {

	// Definerer og initialiserer flagg-variablene
	flag.Float64Var(&fahr, "F", 0.0, "temperatur i grader Fahrenheit")
	flag.Float64Var(&kelvin, "K", 0.0, "temperatur i grader Kelvin")
	flag.Float64Var(&celsius, "C", 0.0, "temperatur i grader Celsius")

	flag.StringVar(&out, "out", "C", "beregne temperatur i F - farhenheit")

	flag.StringVar(&funf, "funfacts", "sun", "\"fun-facts\" om sun - Solen, luna - Månen og terra - Jorden")
	flag.StringVar(&t, "t", "C", "bestemmer hvilken temperaturskala som skal brukes når funfacts skal vises")

}

/*
Program som kan tar innput i form av (kun en om gangen) temperaturene Celsius, Fahrenheit og Kalvin
fra kommandolinje i form av flagg

F.eks.

"-F 100 -out C"

	vil returnere 100 grader Fahrenheit til Celsius.

Og returnerer temperaturen spesifisert av input i "-out X"
*/

func main() {
	fmt.Println("Starting the application...") // Test print

	flag.Parse() // Parse has to be run after decleration of flags and before their access to the program/main

	/*
	   Enkelte kombinasjoner skal ikke være gyldige og da må kontrollstrukturer
	   brukes for å utelukke ugyldige kombinasjoner:
	   -F, -C, -K kan ikke brukes samtidig
	   disse tre kan brukes med -out, men ikke med -funfacts
	   -funfacts kan brukes kun med -t
	   ...
	*/

	//Fahrenheit til Celsius
	if out == "C" && isFlagPassed("F") {
		celsius = conv.FahrenheitToCelsius(fahr)
		fmt.Printf("%.2f°F er %.2f°C", fahr, celsius) //Forsøk på å formattere float output
	}

	//Fahrenheit til Kelvin
	if out == "K" && isFlagPassed("F") {
		kelvin = conv.FahrenheitToKelvin(fahr)
		fmt.Printf("%.2f °F er %.2f °K", fahr, kelvin)
	}

	//Celsius til Fahrenheit
	if out == "F" && isFlagPassed("C") {
		fahr = conv.CelsiusToFahrenheit(celsius)
		fmt.Printf("%.2f°C er %.2f°F", celsius, fahr)
	}

	//Celsius til Kelvin
	if out == "K" && isFlagPassed("C") {
		kelvin = conv.CelsiusToKelvin(celsius)
		fmt.Printf("%.2f°C er %.2f°K", celsius, kelvin)
	}

	//Kelvin til Fahrenheit
	if out == "F" && isFlagPassed("K") {
		fahr = conv.KelvinToFahrenheit(kelvin)
		fmt.Printf("%.2f°K er %.2f°F", kelvin, fahr)
	}

	//Kelvin til Celsius
	if out == "C" && isFlagPassed("K") {
		celsius = conv.KelvinToCelsius(kelvin)
		fmt.Printf("%.2f°K er %.2f°C", kelvin, celsius)
	}

	// FUNFACTS
	// Sun facts
	if funf == "sun" && isFlagPassed("funfacts") {
		sunFact := funfacts.GetFunFacts(funf)
		if t == "C" { // if user typed Celsius

			fmt.Printf("%v %v %v", sunFact[0], 15000000, "°C.\n") // Skal bytte til printf
			fmt.Println(sunFact[1], conv.KelvinToCelsius(5778), "°C.")
		} else if t == "K" { // if user typed Kelvin
			fmt.Println(sunFact[0], conv.CelsiusToKelvin(15000000), "°K.\n", sunFact[1], 5778, "°K.")
		} else if t == "F" { // if user typed Fahrenheit
			fmt.Println(sunFact[0], conv.CelsiusToFahrenheit(15000000), "°F.\n", sunFact[1], conv.KelvinToFahrenheit(5778), "°F.")
		}
	}

	// Luna facts
	if funf == "luna" && isFlagPassed("funfacts") {
		lunafact := funfacts.GetFunFacts(funf)
		if t == "C" {
			fmt.Println(lunafact[0], -183, "°C.\n", lunafact[1], 106, "°C.")
		} else if t == "K" {
			fmt.Println(lunafact[0], conv.CelsiusToKelvin(-183), "°K.\n", lunafact[1], conv.CelsiusToKelvin(106), "°K.")
		} else if t == "F" {
			fmt.Println(lunafact[0], conv.CelsiusToFahrenheit(-183), "°F.\n", lunafact[1], conv.CelsiusToFahrenheit(106), "°F.")
		}
	}

	//Terra facts
	if funf == "terra" && isFlagPassed("funfacts") {
		terrafact := funfacts.GetFunFacts(funf)
		if t == "C" {
			fmt.Println(terrafact[0], 56.7, "°C.\n", terrafact[1], -89.4, "°C.")
		} else if t == "K" {
			fmt.Println(terrafact[0], 329.82, "°K.\n", terrafact[1], conv.CelsiusToKelvin(-89.4), "°K.")
		} else if t == "F" {
			fmt.Println(terrafact[0], 134, "°F.\n", terrafact[1], conv.CelsiusToFahrenheit(-89.4), "°F.")
		}
	}
}

// Funksjonen sjekker om flagget er spesifisert på kommandolinje

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
