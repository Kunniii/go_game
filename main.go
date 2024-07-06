package main

import (
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	spriteSize = 16
	screenH    = spriteSize * 30
	screenW    = spriteSize * 17
)

type Game struct {
	PlayerImage *ebiten.Image
	X, Y        float64
}

func (g *Game) Update() error {

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.X -= 2
		if g.X < 0 {
			g.X = 0
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.X += 2
		if g.X > screenH-spriteSize {
			g.X = screenH - spriteSize
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.Y += 2
		if g.Y > screenW-spriteSize {
			g.Y = screenW - spriteSize
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.Y -= 2
		if g.Y < 0 {
			g.Y = 0
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyF11) {
		if ebiten.IsFullscreen() {
			ebiten.SetFullscreen(false)
		} else {
			ebiten.SetFullscreen(true)
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{120, 180, 255, 255})

	opts := ebiten.DrawImageOptions{}
	opts.GeoM.Translate(g.X, g.Y)

	screen.DrawImage(g.PlayerImage.SubImage(image.Rect(0, 0, 16, 16)).(*ebiten.Image), &opts)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return screenH, screenW
}

func main() {
	ebiten.SetWindowSize(screenH*3, screenW*3)
	ebiten.SetWindowTitle("Go Game")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	playerImage, _, err := ebitenutil.NewImageFromFile("assets/images/ninja.png")

	if err != nil {
		log.Fatal("Player image cannot be loaded!", err)
	}

	if err := ebiten.RunGame(&Game{PlayerImage: playerImage, X: 0, Y: 0}); err != nil {
		log.Fatal(err)
	}

}
