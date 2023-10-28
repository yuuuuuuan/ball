package ballgame

type Time struct {
	tick int
	time int
}

func NewTime() *Time {
	return &Time{
		tick: 0,
		time: 0,
	}
}

func (t *Time) Update() {
	t.tick++
	if t.tick == 59 {
		t.tick = 0
		t.time++
	}
}
