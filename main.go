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

func (g *grid) checkValidMove(x, y int) bool {
	if g.checkFilled(x, y) {
		return false
	} else if x < 0 || y < 0 {
		return false
	} else if x >= g.size || y >= g.size {
		return false
	} else if g.current != nil {
		if x == g.current[0]+1 && y == g.current[1] {
			return true
		} else if x == g.current[0]-1 && y == g.current[1] {
			return true
		} else if x == g.current[0] && y == g.current[1]+1 {
			return true
		} else if x == g.current[0] && y == g.current[1]-1 {
			return true
		}
	}
	return true
}

func (g *grid) move(x, y int) {
	g.current = []int{x, y}
}

func drawGrid(g grid) {
	fmt.Println("\n    0   1   2   3   4")
	for i := 0; i < g.size; i++ {

		fmt.Print("  ---------------------\n", i, " ")

		for j := 0; j < g.size; j++ {

			if g.current == nil {
				fmt.Print("|   ")
			} else {
				if g.checkFilled(i, j) {
					fmt.Print("| X ")
				} else if j == g.current[0] && i == g.current[1] {
					fmt.Print("| O ")
				} else if j == g.current[0]+1 && i == g.current[1] {
					fmt.Print("| # ")
				} else if j == g.current[0]-1 && i == g.current[1] {
					fmt.Print("| # ")
				} else if j == g.current[0] && i == g.current[1]+1 {
					fmt.Print("| # ")
				} else if j == g.current[0] && i == g.current[1]-1 {
					fmt.Print("| # ")
				} else {
					fmt.Print("|   ")
				}
			}
		}
		fmt.Println("|")
	}
	fmt.Println("  ---------------------")

}

func main() {
	//Init
	gameGrid := grid{5, nil, make([][]int, 0)}
	reader := bufio.NewReader(os.Stdin)

	//First move
	for {
		drawGrid(gameGrid)
		//Get input
		fmt.Print("Set start location: ")
		input, _ := reader.ReadString('\n')
		inputSplit := strings.Split(input, ",")
		x, err := strconv.Atoi(inputSplit[0])
		y, err := strconv.Atoi(strings.TrimSpace(inputSplit[1]))

		//Validation
		if err != nil {
			fmt.Println("Please input: number,number")
		} else if gameGrid.checkValidMove(x, y) {
			fmt.Println("Start point:", x, y)
			gameGrid.move(x, y)
			break
		} else {
			fmt.Println("Invalid move")
		}
	}

	//Other moves
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

/*
    0   1   2   3   4
  ---------------------
0 |   |   |   |   |   |
  ---------------------
1 |   |   |   |   |   |
  ---------------------
2 |   |   |   |   |   |
  ---------------------
3 |   |   |   |   |   |
  ---------------------
4 |   |   |   |   |   |
  ---------------------
*/
