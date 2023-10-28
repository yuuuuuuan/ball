package ballgame

type Footstep struct {
	size  int
	pos   *Position
	state int
}

func NewFootstep(x float64, footstepsize int, footstepstate int) *Footstep {
	p := &Position{
		x: x,
		y: ScreenHeight,
	}
	f := &Footstep{
		size:  footstepsize,
		pos:   p,
		state: footstepstate,
	}
	return f
}

func (f *Footstep) Update(g *Game) {

}
