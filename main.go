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

//TODO double check that this does what its supposed to
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
	if g.checkFilled(x, y) { //picked move
		return false
	} else if x < 0 || y < 0 { //out of bounds
		return false
	} else if x >= g.size || y >= g.size { //out of bounds
		return false
	} else if g.current == nil { //starting move
		return true
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
	return false
}

func (g *grid) setPosition(x, y int) {
	g.current = []int{x, y}
}

func (g *grid) move(x, y int) {
	//get direction
	//x,y guarenteed to be valid
	if x == g.current[0]+1 && y == g.current[1] { //right
		for i := g.current[0]; i < g.size; i++ {
			g.fill(i, y)
			//if next is filled, setpos and break
			if i == g.size-1 || g.checkFilled(i+1, y) {
				g.setPosition(i, y)
				break
			}
		}
	} else if x == g.current[0]-1 && y == g.current[1] { //left
		for i := g.current[0]; i > -1; i-- {
			fmt.Println(i, y)
			g.fill(i, y)
			//if next is filled, setpos and break
			if i == 0 || g.checkFilled(i-1, y) {
				g.setPosition(i, y)
				break
			}
		}
	} else if x == g.current[0] && y == g.current[1]+1 { //down
		for i := g.current[1]; i < g.size; i++ {
			g.fill(x, i)
			//if next is filled, setpos and break
			if i == g.size-1 || g.checkFilled(x, i+1) {
				g.setPosition(x, i)
				break
			}
		}
	} else if x == g.current[0] && y == g.current[1]-1 { //up
		for i := g.current[1]; i > -1; i-- {
			g.fill(x, i)
			//if next is filled, setpos and break
			if i == 0 || g.checkFilled(x, i-1) {
				g.setPosition(x, i)
				break
			}
		}
	}
}

func drawGrid(g grid) {
	fmt.Println("\n    0   1   2   3   4")
	for i := 0; i < g.size; i++ {

		fmt.Print("  ---------------------\n", i, " ")

		for j := 0; j < g.size; j++ {

			if g.current == nil {
				fmt.Print("|   ")
			} else {
				if j == g.current[0] && i == g.current[1] {
					fmt.Print("| O ")
				} else if g.checkFilled(j, i) {
					fmt.Print("| X ")
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

func getInput() (int, int, error) {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	inputSplit := strings.Split(input, ",")
	x, err := strconv.Atoi(inputSplit[0])
	if err != nil {
		return 0, 0, err
	}
	y, err := strconv.Atoi(strings.TrimSpace(inputSplit[1]))
	if err != nil {
		return 0, 0, err
	}
	return x, y, nil
}

func main() {
	//Init
	gameGrid := grid{5, nil, make([][]int, 0)}

	//First move
	for {
		drawGrid(gameGrid)
		//Get input
		fmt.Println("Place starting move:")
		x, y, err := getInput()

		//Validate
		if err != nil {
			fmt.Println("ERROR: Please input: number,number\nTry again.")
		} else if gameGrid.checkValidMove(x, y) {
			gameGrid.setPosition(x, y)
			break
		} else {
			fmt.Println("Invalid move")
		}
	}

	fmt.Println("Start position set at: ", gameGrid.current)

	//Other moves
	for {
		//TODO check win con should be moved to after a move is completed
		//Check win con
		if gameGrid.won() {
			fmt.Println("YOU WIN!")
			break
		}

		//Get input
		drawGrid(gameGrid)
		fmt.Println("Place next move:")
		x, y, err := getInput()
		if err != nil {
			fmt.Println("Please input: number,number\nTry again")
			continue
		}

		//Validate move
		if gameGrid.checkValidMove(x, y) {
			//move
			gameGrid.move(x, y)
		} else {
			fmt.Println("Invalid move position\nTry again.")
		}

		//Post-move conditions
		//check for cannot move
		//check for win condition
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
