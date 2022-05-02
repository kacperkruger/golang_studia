package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"strings"
)

var days = [31]string{
	"Error",
	"przewodowy",
	"niebieski",
	"czerowny",
	"głodona",
	"spragniony",
	"brudny",
	"wesoły",
	"szczęśliwy",
	"smutny",
	"duży",
	"mały",
	"nietuzinkowy",
	"czarny",
	"kolorowy",
	"brzydki",
	"słodki",
	"niecodzienny",
	"katastroficzny",
	"zamykany",
	"podświetlany",
	"zmieszany",
	"piękny",
	"miękki",
	"tolerancyjny",
	"zimny",
	"ciepły",
	"letni",
	"biały",
	"zamykany",
	"okrągły",
}

var names = map[rune]string{
	'A': "stokrotka",
	'B': "łosoś",
	'C': "gibon",
	'D': "gepard",
	'E': "tygrys",
	'F': "niedźwiedź",
	'G': "samolot",
	'H': "samochód",
	'I': "nietoperz",
	'J': "mrówka",
	'K': "żyrafa",
	'L': "słoń",
	'M': "okoń",
	'N': "woda",
	'O': "mleko",
	'P': "sok pomarańczowy",
	'R': "sok jabłkowy",
	'S': "jabłko",
	'T': "dżem",
	'U': "masło orzechowe",
	'V': "masło",
	'W': "dynia",
	'X': "pomidor",
	'Y': "sałatka grecka",
	'Z': "astronauta",
}

var surnames = map[rune]string{
	'A': "z plecakiem",
	'B': "pod ziemią",
	'C': "pod wodą",
	'D': "z ciechocinka",
	'E': "z bydgoszczy",
	'F': "z dużym mieczem",
	'K': "z bydgoszczy",
}

func firstLetterValidation(firstLetter rune) (rune, error) {
	if firstLetter < 65 || firstLetter > 90 {
		return firstLetter, errors.New("invalid name or surname")
	}
	return firstLetter, nil
}

func dayValidation(day int) (int, error) {
	if day < 0 || day > 31 {
		return 0, errors.New("wrong day")
	}
	return day, nil
}

func getDay() (int, error) {
	var day int
	flag.IntVar(&day, "day", 0, "day of birth")
	flag.Parse()
	if day == 0 {
		fmt.Print("Enter day of birth: ")
		fmt.Scanf("%d\n", &day)
	}
	return dayValidation(day)
}

func getName() (rune, error) {
	var nameChar rune
	var name string
	flag.StringVar(&name, "name", "", "first letter of name")
	flag.Parse()
	if name == "" {
		fmt.Print("Enter your name: ")
		fmt.Scanf("%s\n", &name)
	}
	name = strings.ToUpper(name)
	nameChar = rune(name[0])
	return firstLetterValidation(nameChar)
}

func getSurname() (rune, error) {
	var surnameChar rune
	var surname string
	flag.StringVar(&surname, "surname", "", "first letter of surname")
	flag.Parse()
	if surname == "" {
		fmt.Print("Enter your surname: ")
		fmt.Scanf("%s\n", &surname)
	}
	surname = strings.ToUpper(surname)
	surnameChar = rune(surname[0])
	return firstLetterValidation(surnameChar)
}

func logError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	log.SetPrefix("nick generator: ")
	log.SetFlags(0)
	day, errDay := getDay()
	logError(errDay)
	name, errName := getName()
	logError(errName)
	surname, errSurname := getSurname()
	logError(errSurname)
	fmt.Println(days[day], names[name], surnames[surname])
}
