package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"sort"
	"strconv"
	"time"
)

func guessNumber() {
	randomNumber := strconv.Itoa(10)
	iterations := 0
	for true {
		iterations += 1
		var guess string
		fmt.Print("Podaj liczbę: ")
		fmt.Scanf("%v\n", &guess)
		if guess == "koniec" {
			fmt.Println("Żegnaj")
			break
		}
		if guess == randomNumber {
			fmt.Printf("Gratulacje zgadłeś za %v razem! \n", iterations)
			saveScore(iterations)
			playAgain()
			break
		}
		if guess > randomNumber {
			fmt.Println("Za duża")
		}
		if guess < randomNumber {
			fmt.Println("Za mała")
		}
	}
}

func saveScore(score int) {
	save := yesOrNo("Czy chcesz zapisać wynik?")
	if save == true {
		var name string
		fmt.Print("Podaj imie: ")
		fmt.Scanf("%s\n", &name)
		addScore(name, score)
		if checkForGlobalRecord(score, getRanking()) == true {
			fmt.Println("Osiągnąłeś nowy najlepszy wynik!")
		} else if checkForLocalRecord(name, score, getRanking()) {
			fmt.Println("Pobiłeś swój poprzedni wynik!")
		}
	}
}

func getRanking() []record {
	var hallOfFame []record
	file, err1 := ioutil.ReadFile("guess_game/hallOfFame.json")
	if err1 != nil {
		log.Fatal(err1)
	}
	err2 := json.Unmarshal(file, &hallOfFame)
	if err2 != nil {
		log.Fatal(err2)
	}
	sort.Slice(hallOfFame, func(i, j int) bool {
		return hallOfFame[i].Score < hallOfFame[j].Score
	})
	return hallOfFame
}

func checkForLocalRecord(name string, score int, hallOfFame []record) bool {
	for _, value := range hallOfFame {
		if value.Name == name && score < value.Score {
			return true
		}
	}
	return false
}

func checkForGlobalRecord(score int, hallOfFame []record) bool {
	bestScore := getTheBestScore(hallOfFame)
	fmt.Println(bestScore)
	if score < bestScore {
		return true
	}
	return false
}

func getTheBestScore(hallOfFame []record) int {
	return hallOfFame[0].Score
}

type record struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
	Date  string `json:"date"`
}

func addScore(name string, score int) {
	var hallOfFame []record
	file, err1 := ioutil.ReadFile("guess_game/hallOfFame.json")
	if err1 != nil {
		log.Fatal(err1)
	}
	err2 := json.Unmarshal(file, &hallOfFame)
	if err2 != nil {
		log.Fatal(err2)
	}
	records := append(hallOfFame, record{name, score, time.Now().Format("01-02-2006")})
	result, _ := json.Marshal(records)
	err3 := ioutil.WriteFile("guess_game/hallOfFame.json", result, 0644)
	if err3 != nil {
		log.Fatal(err3)
	}
}

func yesOrNo(message string) bool {
	var response string
	for true {
		fmt.Printf("%s [T/N]\n", message)
		fmt.Scanf("%s", &response)
		if response == "T" || response == "tak" || response == "Tak" || response == "t" {
			return true
		}
		if response == "N" || response == "nie" || response == "Nie" || response == "n" {
			return false
		}
	}
	return false
}

func playAgain() {
	again := yesOrNo("Czy chcesz zagrać ponownie?")
	if again == true {
		guessNumber()
	} else {
		fmt.Println("Dzięki za gre!")
		printHallOfFame(getRanking())
	}
}

func printHallOfFame(hallOfFame []record) {
	fmt.Printf("%13s %3s\n", "name", "score")
	for index, value := range hallOfFame {
		fmt.Printf("%2d. %10s %3d\n", index+1, value.Name, value.Score)
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	fmt.Println("Teraz będziesz zgadywać liczbę, którą wylosowałem (wpisz \"koniec\" żeby zakończyć gre)")
	guessNumber()
}
