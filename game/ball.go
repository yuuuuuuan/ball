package ballgame

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"image/color"
)

type Ball struct {
	size  int
	pos   *Position
	alive bool
	img   *ebiten.Image
}

func NewBall(size int) *Ball {
	rect := image.Rect(0, 0, size, size)
	img := ebiten.NewImageWithOptions(rect, nil)
	img.Fill(color.White)
	p := &Position{
		x: 50,
		y: 50,
	}
	ball := &Ball{
		size:  size,
		pos:   p,
		alive: true,
		img:   img,
	}
	return ball
}

func (b *Ball) Update(i *Input) {
	switch i.state {
	case keyleft:
		{
			b.pos.x--
		}
	case keyright:
		{
			b.pos.x++
		}
	case keynone:
		{

		}
	}
}

func (b *Ball) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(b.pos.x, b.pos.y)
	screen.DrawImage(b.img, op)
}
