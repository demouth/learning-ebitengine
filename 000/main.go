package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct{}

// Updateはゲームを1ティック更新します
func (g *Game) Update() error {
	// ゲームロジックをここに書きますが、今はまだ何もしません
	return nil
}

// Draw はゲーム画面を1フレーム分描画します
func (g *Game) Draw(screen *ebiten.Image) {
	// 今はまだ何も描画しないので空のままにしておきます
}

// Layout はゲームの論理的な画面サイズを返します
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
func main() {
	ebiten.RunGame(&Game{})
}
