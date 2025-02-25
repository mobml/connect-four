package main

import (
	"github.com/hajimehoshi/ebiten/v2"
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
	currentPlayer = "X"
)

type Game struct{}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.Key1) {
		dropPiece(0, currentPlayer)
		switchPlayer()
	} else if inpututil.IsKeyJustPressed(ebiten.Key2) {
		dropPiece(1, currentPlayer)
		switchPlayer()
	} else if inpututil.IsKeyJustPressed(ebiten.Key3) {
		dropPiece(2, currentPlayer)
		switchPlayer()
	} else if inpututil.IsKeyJustPressed(ebiten.Key4) {
		dropPiece(3, currentPlayer)
		switchPlayer()
	} else if inpututil.IsKeyJustPressed(ebiten.Key5) {
		dropPiece(4, currentPlayer)
		switchPlayer()
	} else if inpututil.IsKeyJustPressed(ebiten.Key6) {
		dropPiece(5, currentPlayer)
		switchPlayer()
	} else if inpututil.IsKeyJustPressed(ebiten.Key7) {
		dropPiece(6, currentPlayer)
		switchPlayer()
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

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Connect Four")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
