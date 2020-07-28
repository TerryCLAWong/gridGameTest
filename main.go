package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type grid struct {
	size    int
	current []int
	filled  [][]int
}

func (g *grid) fill(x, y int) {
	g.filled = append(g.filled, []int{x, y})
}

func (g *grid) won() bool {
	return len(g.filled) == g.size*g.size
}

func (g *grid) checkFilled(x, y int) bool {
	for _, spot := range g.filled {
		if spot[0] == x && spot[1] == y {
			return true
		}
	}
	return false
}

func (g *grid) move(x, y int) {
	g.current = []int{x, y}
}

func drawGrid(g grid) {
	if g.current == nil {
		for i := 0; i < g.size; i++ {
			fmt.Println("---------------------")
			fmt.Println("|   |   |   |   |   |")
		}
		fmt.Println("---------------------")
	} else {

		for i := 0; i < g.size; i++ {
			fmt.Println("---------------------")

			for j := 0; j < g.size; j++ {
				if g.checkFilled(i, j) {
					fmt.Print("| X ")
					//else if check option TODO
				} else if i == g.current[0] && j == g.current[1] {
					fmt.Print("| O ")
				} else {
					fmt.Print("|   ")
				}
			}
			fmt.Println("|")
		}
		fmt.Println("---------------------")
	}
}

func main() {
	//Init
	gameGrid := grid{5, nil, make([][]int, 0)}
	reader := bufio.NewReader(os.Stdin)

	//First move
	drawGrid(gameGrid)
	fmt.Print("Set start location: ")
	input, _ := reader.ReadString('\n')
	inputSplit := strings.Split(input, ",")
	x, err := strconv.Atoi(inputSplit[0])
	y, err := strconv.Atoi(strings.TrimSpace(inputSplit[1]))
	if err != nil {
		fmt.Println("Please input number,number")
	}
	gameGrid.move(x, y)

	//Loop
	for {
		//Check win con
		if gameGrid.won() {
			fmt.Println("YOU WIN!")
			break
		}

		drawGrid(gameGrid)

		//Get input
		fmt.Print("Input: ")
		input, _ := reader.ReadString('\n')
		fmt.Println(input)
		/*
			Print board

			Get input
			Check input validity
				if ok, continue
				if not, print bad message and get input again

		*/

	}
}
