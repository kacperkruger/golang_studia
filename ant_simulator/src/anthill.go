package ant_simulator_src

import (
	"fmt"
	"math/rand"
	"time"
)

func createAnt(x int, y int) Ant {
	a := Ant{x, y, &Leaf{0, 0}}
	return a
}

func createLeaf(x int, y int) Leaf {
	l := Leaf{x, y}
	return l
}

func createBoard(width int, height int) ([][]*Ant, [][]*Leaf, [][]bool) {
	boardOfAnts := make([][]*Ant, height)
	boardOfLeafs := make([][]*Leaf, height)
	boardOfStacks := make([][]bool, height)
	for i := 0; i < height; i++ {
		boardOfAnts[i] = make([]*Ant, width)
		boardOfLeafs[i] = make([]*Leaf, width)
		boardOfStacks[i] = make([]bool, width)
	}
	return boardOfAnts, boardOfLeafs, boardOfStacks
}

func placeAnt(ant *Ant, board [][]*Ant) {
	board[ant.x][ant.y] = ant
}

func placeLeaf(leaf *Leaf, board [][]*Leaf) {
	board[leaf.x][leaf.y] = leaf
}

func generateAnts(width int, height int) []Ant {
	var listOfAnts []Ant
	for i := 0; i < width*height*50/100; i++ {
		ant := createAnt(rand.Intn(width-1)+1, rand.Intn(height-1)+1)
		listOfAnts = append(listOfAnts, ant)
	}
	return listOfAnts
}

func generateLeafs(width int, height int) []Leaf {
	var listOfLeafs []Leaf
	amountOfLeafs := 0
	for i := 0; i < width*height*80/100; i++ {
		leaf := createLeaf(rand.Intn(width-1)+1, rand.Intn(height-1)+1)
		amountOfLeafs += 1
		listOfLeafs = append(listOfLeafs, leaf)
	}
	return listOfLeafs
}

func initAnthill(width int, height int) ([][]*Ant, [][]*Leaf, [][]bool, []Ant, []Leaf, int) {
	boardOfAnts, boardOfLeafs, boardOfStacks := createBoard(width, height)
	amountOfLeafs := 0
	listOfAnts := generateAnts(width, height)
	for _, ant := range listOfAnts {
		placeAnt(&ant, boardOfAnts)
	}
	listOfLeafs := generateLeafs(width, height)
	for _, leaf := range listOfLeafs {
		placeLeaf(&leaf, boardOfLeafs)
		amountOfLeafs += 1
	}
	return boardOfAnts, boardOfLeafs, boardOfStacks, listOfAnts, listOfLeafs, amountOfLeafs
}

func move(boardOfAnts [][]*Ant, boardOfLeafs [][]*Leaf, boardOfStacks [][]bool, amountOfLeafs int, ant Ant, dx int, dy int) int {
	boardOfAnts[ant.GetY()][ant.GetX()] = nil
	ant.Move(dx, dy)
	newX := ant.GetX()
	newY := ant.GetY()
	boardOfAnts[newY][newX] = &ant
	leaf := *boardOfLeafs[newX][newY]
	if (leaf.GetX() != 0 || leaf.GetY() != 0) && !IsCarryingLeaf(ant) && !boardOfStacks[newY][newX] {
		ant.PickUpLeaf(&leaf)
		boardOfLeafs[newY][newX] = nil
		fmt.Printf("I found leaf %v\n", &ant)
	}
	if (leaf.GetX() != 0 || leaf.GetY() != 0) && IsCarryingLeaf(ant) {
		ant.DropLeaf()
		boardOfStacks[newY][newX] = true
		amountOfLeafs -= 2
		fmt.Println("Dropped the leaf")
	}
	return amountOfLeafs
}

func Play(width int, height int) {
	rand.Seed(time.Now().UnixNano())
	boardOfAnts, boardOfLeafs, boardOfStacks, listOfAnts, listOfLeafs, amountOfLeafs := initAnthill(width, height)
	fmt.Println(amountOfLeafs)
	printAnts(boardOfAnts)
	printLeafs(boardOfLeafs)
	printListOfLeafs(listOfLeafs)
leafs:
	for {
		for _, ant := range listOfAnts {
			dx := rand.Intn(2) - 1
			dy := rand.Intn(2) - 1
			if ant.GetX()+dx < width && ant.GetY()+dy < height && ant.GetX()+dx != 0 && ant.GetY()+dy != 0 {
				amountOfLeafs = move(boardOfAnts, boardOfLeafs, boardOfStacks, amountOfLeafs, ant, dx, dy)
				if amountOfLeafs <= 0 {
					break leafs
				}
			}
		}
	}

	printAnts(boardOfAnts)
	printLeafs(boardOfLeafs)
	printStacks(boardOfStacks)
}

func printListOfLeafs(listOfLeafs []Leaf) {
	fmt.Println("List of leafs")
	for _, value := range listOfLeafs {
		fmt.Print(value)
	}
}

func printAnts(boardOfAnts [][]*Ant) {
	fmt.Println("Ants")
	for _, value := range boardOfAnts {
		for _, val := range value {
			if val == nil {
				fmt.Print("nil")
			} else {
				ant := *val
				if ant.GetY() != 0 || ant.GetX() != 0 {
					fmt.Print("🐜")
				}
			}
		}
		fmt.Println()
	}
}

func printLeafs(boardOfLeafs [][]*Leaf) {
	fmt.Println("Leafs")
	for _, value := range boardOfLeafs {
		for _, val := range value {
			if val == nil {
				fmt.Print("nil")
			} else {
				leaf := *val
				if leaf.GetY() != 0 || leaf.GetX() != 0 {
					fmt.Print("🍁")
				}

			}
		}
		fmt.Println()
	}
}

func printStacks(boardOfStacks [][]bool) {
	fmt.Println("Stacks")
	for _, value := range boardOfStacks {
		fmt.Println(value)
	}
}
