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
	x            int
	y            int
	scale        float64 = 1
)

type Game struct{}

// Updateはゲームを1ティック更新します
func (g *Game) Update() error {
	theta += 0.05
	x += 2
	y += 3
	x = x % 640 // xの値が640を超えたら0に戻す
	y = y % 480 // yの値が480を超えたら0に戻す
	scale += 0.01
	if scale > 2 {
		scale = 0.5
	}
	return nil
}

// Draw はゲーム画面を1フレーム分描画します
func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	bounds := pumpkinImage.Bounds() // 画像の大きさを取得する
	op.GeoM.Translate(-float64(bounds.Dx())/2, -float64(bounds.Dy())/2)
	op.GeoM.Rotate(theta)
	op.GeoM.Scale(scale, scale)
	op.GeoM.Translate(float64(bounds.Dx())/2, float64(bounds.Dy())/2)
	op.GeoM.Translate(float64(x), float64(y))
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
