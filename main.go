package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"log"
	"math/rand"
)

const (
	screenWidth  = 900.0
	screenHeight = 600.0
	paddleWidth  = 20.0
	paddleHeight = 80.0
	ballSize     = 20.0
)

type Game struct{}

var (
	player1Y     = screenHeight / 2.0
	player2Y     = screenHeight / 2.0
	ballX, ballY = screenWidth / 2.0, screenHeight / 2.0
	ballSpeedX   = 4.0
	ballSpeedY   = 4.0
	player1Speed = 4.0
	player2Speed = 4.0
	ballColor    = color.RGBA{255, 255, 255, 0}
)

func (g *Game) Update() error {
	log.Printf("ballX: %f, ballY: %f", ballX, ballY)
	if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
		err := ebiten.Termination
		return err
	}

	// Update ball position
	ballX += ballSpeedX
	ballY += ballSpeedY

	// Ball collisions with walls
	if ballY <= 0 || ballY >= screenHeight-ballSize {
		ballSpeedY = -ballSpeedY
		ballColor = color.RGBA{R: uint8(rand.Int() % 255), G: uint8(rand.Int() % 255), B: uint8(rand.Int() % 255), A: 255}
	}

	// Ball collisions with paddles
	if (ballX <= paddleWidth && ballX >= 0 && ballY >= player1Y && ballY <= player1Y+paddleHeight) ||
		(ballX >= screenWidth-paddleWidth-ballSize && ballX <= screenWidth-ballSize && ballY >= player2Y && ballY <= player2Y+paddleHeight) {
		ballSpeedX = -ballSpeedX
		ballColor = color.RGBA{R: uint8(rand.Int() % 255), G: uint8(rand.Int() % 255), B: uint8(rand.Int() % 255), A: 255}
	}

	// Move paddles
	if ebiten.IsKeyPressed(ebiten.KeyW) && player1Y > 0 {
		player1Y -= player1Speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) && player1Y+paddleHeight < screenHeight {
		player1Y += player1Speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) && player2Y > 0 {
		player2Y -= player2Speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) && player2Y+paddleHeight < screenHeight {
		player2Y += player2Speed
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Draw paddles
	vector.DrawFilledRect(screen, 0, float32(player1Y), paddleWidth, paddleHeight, color.White, true)
	vector.DrawFilledRect(screen, screenWidth-paddleWidth, float32(player2Y), paddleWidth, paddleHeight, color.White, false)

	// Draw ball
	vector.DrawFilledCircle(screen, float32(ballX+ballSize/2), float32(ballY+ballSize/2), ballSize/2, ballColor, true)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return int(screenWidth), int(screenHeight)
}

func main() {
	ebiten.SetWindowSize(int(screenWidth), int(screenHeight))
	ebiten.SetWindowTitle("Pong in Golang with Ebiten")

	game := &Game{}
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
