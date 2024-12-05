package models

type Player struct {
	Name         string
	TelegramTag  string
	Position     string
	PassLevel    int
	ShootLevel   int
	DefenceLevel int
}

func (p *Player) OverallSkill() int {
	return p.PassLevel + p.ShootLevel + p.DefenceLevel
}
