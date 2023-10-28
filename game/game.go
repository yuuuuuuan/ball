package ballgame

import "github.com/hajimehoshi/ebiten/v2"

const (
	ScreenWidth  = 840
	ScreenHeight = 1200
)

const (
	Ballsize = 30
)

type Position struct {
	x float64
	y float64
}

type Game struct {
	input *Input
	ball  *Ball
	time  *Time
	level *Level
}

func NewGame() (*Game, error) {
	g := &Game{
		input: NewInput(),
		ball:  NewBall(Ballsize),
		time:  NewTime(),
		level: NewLevel(),
	}
	return g, nil
}

func (g *Game) Update() error {
	g.time.Update()
	g.input.Update()
	g.level.Update(g.time)
	g.ball.Update(g.input)
	return nil
}

func (g *Game) Layout(int, int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.ball.Draw(screen)
}
