package main

import (
	"bytes"
	_ "embed"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	// カボチャの画像を埋め込む
	//go:embed pumpkin.png
	pumpkinPng   []byte
	pumpkinImage *ebiten.Image
	theta        float64
)

type Game struct{}

// Updateはゲームを1ティック更新します
func (g *Game) Update() error {
	theta += 0.05
	return nil
}

// Draw はゲーム画面を1フレーム分描画します
func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Rotate(theta)
	screen.DrawImage(pumpkinImage, op)
}

// Layout はゲームの論理的な画面サイズを返します
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
func main() {
	// png画像を *ebiten.Image に変換します
	img, _, _ := image.Decode(bytes.NewReader(pumpkinPng))
	pumpkinImage = ebiten.NewImageFromImage(img)
	ebiten.RunGame(&Game{})
}
