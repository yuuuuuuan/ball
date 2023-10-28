package ballgame

const (
	easy = 20
	mid  = 40
	hard = 60
)

type Level struct {
	level int
}

func NewLevel() *Level {
	return &Level{
		level: 0,
	}
}

func (l *Level) Update(t *Time) {
	switch t.time {
	case easy:
		{
			l.level++
		}
	case mid:
		{
			l.level++
		}
	case hard:
		{
			l.level++
		}
	}
}
