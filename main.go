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

func (g *grid) multifill(walls [][]int) {
	for _, wall := range walls {
		g.filled = append(g.filled, []int{wall[0], wall[1]})
	}
}

func (g *grid) fill(x, y int) {
	g.filled = append(g.filled, []int{x, y})
}

func (g *grid) checkWin() bool {
	return len(g.filled) == g.size*g.size-1
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

func (g *grid) checkStuck() bool {
	x := g.current[0]
	y := g.current[1]
	return (g.checkFilled(x+1, y) || x+1 == g.size) && (g.checkFilled(x-1, y) || x-1 == -1) && (g.checkFilled(x, y+1) || y+1 == g.size) && (g.checkFilled(x, y-1) || y-1 == -1)
}

func (g *grid) setPosition(x, y int) {
	g.current = []int{x, y}
}

/*
Performs desired move in direction of input parameters x and y.
Fills out all cells between current position and ending position.
Input parameters validated by parent function.
*/
func (g *grid) move(x, y int) {
	//Same code for each case of direction
	if x == g.current[0]+1 && y == g.current[1] { //right
		for i := g.current[0]; i < g.size; i++ {
			//if next is filled, setpos and break
			if i == g.size-1 || g.checkFilled(i+1, y) {
				g.setPosition(i, y)
				break
			}
			g.fill(i, y)
		}
	} else if x == g.current[0]-1 && y == g.current[1] { //left
		for i := g.current[0]; i > -1; i-- {
			//if next is filled, setpos and break
			if i == 0 || g.checkFilled(i-1, y) {
				g.setPosition(i, y)
				break
			}
			g.fill(i, y)
		}
	} else if x == g.current[0] && y == g.current[1]+1 { //down
		for i := g.current[1]; i < g.size; i++ {
			//if next is filled, setpos and break
			if i == g.size-1 || g.checkFilled(x, i+1) {
				g.setPosition(x, i)
				break
			}
			g.fill(x, i)
		}
	} else if x == g.current[0] && y == g.current[1]-1 { //up
		for i := g.current[1]; i > -1; i-- {
			//if next is filled, setpos and break
			if i == 0 || g.checkFilled(x, i-1) {
				g.setPosition(x, i)
				break
			}
			g.fill(x, i)
		}
	}
}

/*
Draws grid defined by input g
Takes into account all g's attributes including
pre-placed walls, current position, move options, post-move walls
*/
func drawGrid(g grid) {
	fmt.Println("\n    0   1   2   3   4")
	for i := 0; i < g.size; i++ {

		fmt.Print("  ---------------------\n", i, " ")

		for j := 0; j < g.size; j++ {
			//Pre-gamestart rendering
			if g.current == nil {
				if g.checkFilled(j, i) {
					fmt.Print("| X ")
				} else {
					fmt.Print("|   ")
				}
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

/*
Gets terminal input and tries to parse format: (integer,integer)
Returns the integers and error from conversions
*/
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

/*
Instantiates set of levels and iterates through each level as the player completes them.
*/
func main() {
	//Set of levels
	levels := []struct {
		g     grid
		walls [][]int
	}{
		{
			grid{5, nil, make([][]int, 0)},
			[][]int{{2, 0}, {3, 0}, {4, 0}, {0, 4}, {1, 4}},
		},
		{
			grid{5, nil, make([][]int, 0)},
			[][]int{{0, 3}, {2, 4}, {3, 2}, {4, 0}},
		},
	}

	//Plays levels individually
	for index, level := range levels {
		level.g.multifill(level.walls)
		for !play(level.g) {
			fmt.Println("No more moves available. Please try again.\nRestarting Level")
		}
		fmt.Println("You win!")
		if index != len(levels)-1 {
			fmt.Println("Moving onto the next level")
		} else {
			fmt.Println("No more levels to complete")
		}
	}

}

/*
Plays the boardstate defined by input grid gameGrid.
Reads input, performs move on the board state, and renders board
Returns whether or not the player has sucessfully completed the level
*/
func play(gameGrid grid) bool {
	//First move
	for {
		drawGrid(gameGrid)
		//Get input
		fmt.Println("Place starting move (x,y):")
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
		//Get input
		drawGrid(gameGrid)
		fmt.Println("Place next move (x,y):")
		x, y, err := getInput()
		if err != nil {
			fmt.Println("Please input: number,number\nTry again")
			continue
		}

		//Validate move
		if gameGrid.checkValidMove(x, y) {
			gameGrid.move(x, y)
		} else {
			fmt.Println("Invalid move position\nTry again.")
		}

		//Post move checking
		if gameGrid.checkWin() {
			return true
		} else if gameGrid.checkStuck() {
			return false
		}
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
