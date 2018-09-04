package main

import (
	"../myPackage"
	"fmt"
	"strings"
)

type NewGame struct {
	available []bool
	move      bool
}

var XOboard = [][]string{
	{"1", "2", "3"},
	{"4", "5", "6"},
	{"7", "8", "9"},
}

func main() {
	game := NewGame{}
	game.available = setPointers(game.available)
	game.move = true

	for {
		drawBoard()

		if game.move {
			game.available = AImove(game.available)
		} else {
			game.available = playerMove(game.available)
		}

		if endGame() {
			finishGame(game.move)
			break
		}

		game.move = !game.move
	}
}

func AImove(available []bool) []bool {
	move := strategyAI(available)

	changeXOboard(move, "X")
	available[move-1] = false
	fmt.Println("My move will be on square", move)

	return available
}

func changeXOboard(move int, typo string) {
	switch move {
	case 1:
		XOboard[0][0] = typo
	case 2:
		XOboard[0][1] = typo
	case 3:
		XOboard[0][2] = typo
	case 4:
		XOboard[1][0] = typo
	case 5:
		XOboard[1][1] = typo
	case 6:
		XOboard[1][2] = typo
	case 7:
		XOboard[2][0] = typo
	case 8:
		XOboard[2][1] = typo
	case 9:
		XOboard[2][2] = typo
	}
}

func drawBoard() {
	fmt.Println()

	for i := 0; i < len(XOboard); i++ {
		fmt.Printf("%s\n", strings.Join(XOboard[i], " "))
	}

	fmt.Println()
}

func endGame() bool {
	win :=
		XOboard[0][0] == XOboard[0][1] && XOboard[0][0] == XOboard[0][2] ||
			XOboard[1][0] == XOboard[1][1] && XOboard[1][0] == XOboard[1][2] ||
			XOboard[2][0] == XOboard[2][1] && XOboard[2][0] == XOboard[2][2] ||
			XOboard[0][0] == XOboard[1][0] && XOboard[0][0] == XOboard[2][0] ||
			XOboard[0][1] == XOboard[1][1] && XOboard[0][1] == XOboard[2][1] ||
			XOboard[0][2] == XOboard[1][2] && XOboard[0][2] == XOboard[2][2] ||
			XOboard[0][0] == XOboard[1][1] && XOboard[0][0] == XOboard[2][2] ||
			XOboard[0][2] == XOboard[1][1] && XOboard[0][2] == XOboard[2][0]

	return win
}

func finishGame(aiWin bool) {
	fmt.Println("\n\tIt's a Game!")
	drawBoard()

	if aiWin {
		fmt.Println("You lose...")
	} else {
		fmt.Println("You WIN! Congratulations!")
	}
}

func playerMove(available []bool) []bool {
	move := myPackage.GetIntFromInput("What's your move? Type number of field to put O on it!\n", 1)

	if move == 0 {
		move = 1
	}

	if move <= 9 && available[move-1] {
		changeXOboard(move, "O")
		available[move-1] = false
		fmt.Println("Well, ok.")
	} else {
		fmt.Println("The field is unavailable. We will choose available field for you then.")
		for i := 0; i < 9; i++ {
			if available[i] {
				changeXOboard(i, "O")
				available[i] = false
			}
		}
	}

	return available
}

func setPointers(available []bool) []bool {
	for i := 0; i < 9; i++ {
		available = append(available, true)
	}

	return available
}

func strategyAI(available []bool) int {
	x := string("X")
	o := string("O")

	// Start moves
	if strategyStart(available) {
		return strategyStartMove()
	}

	// Winning moves
	result := strategyCheckEnd(x, available)
	if result != 0 {
		return result
	}

	// Defend moves
	result = strategyCheckEnd(o, available)
	if result != 0 {
		return result
	}

	/*
				Attack moves on progress ...

		return strategyMoveAttack(available, x, o)

	*/

	// Random AI move
	return strategyMoveRandom(available)
}

func strategyCheckEnd(typo string, available []bool) int {

	switch {

	case available[4] && (XOboard[0][0] == typo && XOboard[2][2] == typo ||
		XOboard[0][1] == typo && XOboard[2][1] == typo ||
		XOboard[0][2] == typo && XOboard[2][0] == typo ||
		XOboard[1][0] == typo && XOboard[1][2] == typo):
		return 5

	case available[0] && (XOboard[0][1] == typo && XOboard[0][2] == typo ||
		XOboard[1][0] == typo && XOboard[2][0] == typo ||
		XOboard[1][1] == typo && XOboard[2][2] == typo):
		return 1

	case available[2] && (XOboard[0][0] == typo && XOboard[0][1] == typo ||
		XOboard[1][2] == typo && XOboard[2][2] == typo ||
		XOboard[1][1] == typo && XOboard[2][0] == typo):
		return 3

	case available[6] && (XOboard[0][0] == typo && XOboard[1][0] == typo ||
		XOboard[2][1] == typo && XOboard[2][2] == typo ||
		XOboard[1][1] == typo && XOboard[0][2] == typo):
		return 7

	case available[8] && (XOboard[2][0] == typo && XOboard[2][1] == typo ||
		XOboard[0][2] == typo && XOboard[1][2] == typo ||
		XOboard[1][1] == typo && XOboard[0][0] == typo):
		return 9

	case available[1] && (XOboard[0][0] == typo && XOboard[0][2] == typo ||
		XOboard[1][1] == typo && XOboard[2][1] == typo):
		return 2

	case available[3] && (XOboard[0][0] == typo && XOboard[2][0] == typo ||
		XOboard[1][1] == typo && XOboard[1][2] == typo):
		return 4

	case available[5] && (XOboard[0][2] == typo && XOboard[2][2] == typo ||
		XOboard[1][1] == typo && XOboard[1][0] == typo):
		return 6

	case available[7] && (XOboard[2][0] == typo && XOboard[2][2] == typo ||
		XOboard[1][1] == typo && XOboard[0][1] == typo):
		return 8
	}

	return 0
}

func strategyMoveRandom(available []bool) int {
	var mass []int

	for i := 0; i < 9; i++ {
		if available[i] {
			mass = append(mass, i+1)
		}
	}

	return mass[myPackage.GetRandomInt(len(mass))]
}

func strategyStart(available []bool) bool {
	allAvailable := true

	for i := 0; i < 9; i++ {
		if !available[i] {
			allAvailable = false
		}
	}

	return allAvailable
}

func strategyStartMove() int {
	mass := []int{1, 3, 7, 9, 5}
	return mass[myPackage.GetRandomInt(5)]
}

/*
		ON PROGRESS - Attack moves

func strategyMoveAttack(available []bool, x, o string) int {

		1,9  .. 2,8  .. 3,7  .. 4,6  ..
		1	2	3		2, 3, 4, 7, 5, 9
		4	5	6
		7	8	9

	switch {
	case XOboard[0][0] == x:

		switch {
		case available[8] && XOboard[]:
			return 9
		}

	case XOboard[1][1] == x:

		switch {
		case available[0] &&  XOboard[2][2] != o:

			return 1
		}

	}


	switch {

	case available[0] && (
		XOboard[0][1] == o && XOboard[0][2] == o ||
			XOboard[1][0] == o && XOboard[2][0] == o ||
			XOboard[1][1] == o && XOboard[2][2] == o  ) :
		return 1

	case available[2] && (
		XOboard[0][0] == o && XOboard[0][1] == o ||
			XOboard[1][2] == o && XOboard[2][2] == o ||
			XOboard[1][1] == o && XOboard[2][0] == o  ) :
		return 3

	case available[6] && (
		XOboard[0][0] == o && XOboard[1][0] == o ||
			XOboard[2][1] == o && XOboard[2][2] == o ||
			XOboard[1][1] == o && XOboard[0][2] == o  ) :
		return 7

	case available[8] && (
		XOboard[2][0] == o && XOboard[2][1] == x ||
			XOboard[0][2] == x && XOboard[1][2] == x ||
			XOboard[1][1] == x && XOboard[0][0] == x  ) :
		return 9

	case available[1] && (
		XOboard[0][0] == x && XOboard[0][2] == x ||
			XOboard[1][1] == x && XOboard[2][1] == x  ) :
		return 2

	case available[3] && (
		XOboard[0][0] == x && XOboard[2][0] == x ||
			XOboard[1][1] == x && XOboard[1][2] == x  ) :
		return 4

	case available[5] && (
		XOboard[0][2] == x && XOboard[2][2] == x ||
			XOboard[1][1] == x && XOboard[1][0] == x  ) :
		return 6

	case available[7] && (
		XOboard[2][0] == x && XOboard[2][2] == x ||
			XOboard[1][1] == x && XOboard[0][1] == x  ) :
		return 8

	case available[4] && (
		XOboard[0][0] == x && XOboard[2][2] == x ||
			XOboard[0][1] == x && XOboard[2][1] == x ||
			XOboard[0][2] == x && XOboard[2][0] == x ||
			XOboard[1][0] == x && XOboard[1][2] == x  ) :
		return 5
	}


	return strategyMoveRandom(available)
}
*/
