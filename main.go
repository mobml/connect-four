package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"log"
)

const (
	screenWidth  = 560
	screenHeight = 480
	rows         = 6
	columns      = 7
	cellSize     = 80
)

var (
	colorBlue     = color.RGBA{160, 206, 217, 255}
	colorGreen    = color.RGBA{173, 247, 182, 255}
	colorYellow   = color.RGBA{255, 238, 147, 255}
	board         [rows][columns]string
	winner        = ""
	currentPlayer = "X"
	gameState     = StatePlaying
)

type GameState int

const (
	StatePlaying = iota
	StateGameOver
)

type Game struct {
	mousePressed bool
}

func (g *Game) Update() error {
	x, _ := ebiten.CursorPosition()

	switch gameState {
	case StatePlaying:
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && !g.mousePressed {
			col := x / cellSize
			if col >= 0 && col < columns {
				dropPiece(col, currentPlayer)
				switchPlayer()
			}
			g.mousePressed = true
		} else if !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			g.mousePressed = false
		}

		if checkVictory("X") {
			winner = "X"
			gameState = StateGameOver
		} else if checkVictory("O") {
			winner = "O"
			gameState = StateGameOver
		}
	case StateGameOver:
		if inpututil.IsKeyJustPressed(ebiten.KeyR) {
			resetBoard()
			currentPlayer = "X"
			winner = ""
			gameState = StatePlaying
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(colorBlue)

	for row := 0; row < rows; row++ {
		for col := 0; col < columns; col++ {
			x := col * cellSize
			y := row * cellSize

			vector.DrawFilledCircle(screen, float32(x+cellSize/2),
				float32(y+cellSize/2), float32(cellSize/2)-5, color.White, true)

			switch board[row][col] {
			case "X":
				vector.DrawFilledCircle(screen, float32(x+cellSize/2),
					float32(y+cellSize/2), float32(cellSize/2)-10, colorGreen, true)
			case "O":
				vector.DrawFilledCircle(screen, float32(x+cellSize/2),
					float32(y+cellSize/2), float32(cellSize/2)-10, colorYellow, true)
			}
		}
	}
	if gameState == StateGameOver {
		msg := fmt.Sprintf("El jugador %s ha ganado\nPresiona R para reiniciar",
			If(winner == "X", "1", "2"))
		ebitenutil.DebugPrintAt(screen, msg, 10, 10)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func dropPiece(pos int, value string) {
	for row := rows - 1; row >= 0; row-- {
		if board[row][pos] == "" {
			board[row][pos] = value
			break
		}

	}
}

func switchPlayer() {
	if currentPlayer == "X" {
		currentPlayer = "O"
	} else {
		currentPlayer = "X"
	}
}

func checkVictory(player string) bool {
	//horizontal
	for row := 0; row < rows; row++ {
		for col := 0; col <= columns-4; col++ {
			if board[row][col] == player &&
				board[row][col+1] == player &&
				board[row][col+2] == player &&
				board[row][col+3] == player {
				return true
			}
		}
	}

	//vertical
	for row := 0; row <= rows-4; row++ {
		for col := 0; col < columns; col++ {
			if board[row][col] == player &&
				board[row+1][col] == player &&
				board[row+2][col] == player &&
				board[row+3][col] == player {
				return true
			}
		}
	}

	//right diagonal
	for row := 0; row <= rows-4; row++ {
		for col := 0; col <= columns-4; col++ {
			if board[row][col] == player &&
				board[row+1][col+1] == player &&
				board[row+2][col+2] == player &&
				board[row+3][col+3] == player {
				return true
			}
		}
	}

	//left diagonal
	for row := 0; row <= rows-4; row++ {
		for col := 3; col < columns; col++ {
			if board[row][col] == player &&
				board[row+1][col-1] == player &&
				board[row+2][col-2] == player &&
				board[row+3][col-3] == player {
				return true
			}
		}
	}

	return false
}

func resetBoard() {
	for row := 0; row < rows; row++ {
		for col := 0; col < columns; col++ {
			board[row][col] = ""
		}
	}
}

func If[T any](cond bool, a, b T) T {
	if cond {
		return a
	}
	return b
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Connect Four")
	resetBoard()
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
