package main 

import "github.com/nsf/termbox-go"

type Player struct {
  X, Y int
  Shape []string
  
  Height, Width int

  Lives int
  Name string 
  Points int

  WeaponIndex int
  AvailableWeapons []Weapon

  Bullets []*Bullet
}

const (
  StartingLives = 3
  StartingName = "Player 1"
)

var PlayerShape = []string{"  ^  ", "<| |>"}

func NewPlayer() *Player {
  x :=  BoardWidth/2 - len(PlayerShape[0])/2
  y := BoardHeight - len(PlayerShape) - 1
  height := len(PlayerShape)
  width := len(PlayerShape[0])
  weaponIndex := 0
  availableWeapons := []Weapon{WeaponMap[0], WeaponMap[1], WeaponMap[2]}

  p := &Player{X: x , Y: y, Height: height, Width: width, 
    Shape: PlayerShape, Lives: StartingLives, Name: StartingName, 
    WeaponIndex: weaponIndex, AvailableWeapons: availableWeapons}
  return p
}

func (p *Player) SwapWeapons(index int) {
  if index >= len(p.AvailableWeapons)  {
    return
  }
  p.WeaponIndex = index
}

func (p *Player) Handle(event termbox.Event) {
  if event.Type != termbox.EventKey {
    return
  }

  scale := 3
  // fmt.Println(event.Ch)

  if event.Key == termbox.KeyArrowUp {
    p.Y -= scale
  } else if event.Key == termbox.KeyArrowRight {
    p.X += scale
  } else if event.Key == termbox.KeyArrowDown {
    p.Y += scale
  } else if event.Key == termbox.KeyArrowLeft {
    p.X -= scale
  } else if event.Key == termbox.KeySpace {
    p.FireWeapon()
  } else if event.Ch >= '0' && event.Ch <= '9'{
    p.SwapWeapons(int(byte(event.Ch) - byte('0')) - 1)
  }

  p.BoundsCheck()
}

func (p *Player) Reset() {
  p.X =  BoardWidth/2 - len(PlayerShape[0])/2
  p.Y = BoardHeight - len(PlayerShape) - 1
  p.Bullets = []*Bullet{}
}

func (p *Player) CurrentWeapon() Weapon {
  return p.AvailableWeapons[p.WeaponIndex]
}

func (p *Player) FireWeapon() {
  x := p.X + p.Width / 2
  y := p.Y
  p.Bullets = append(p.Bullets, p.CurrentWeapon().Fire(x, y, -1)...)
}

func (p *Player) BoundsCheck() {
  if p.X <= 0 {
    p.X = 1
  }

  if p.X + p.Width >= BoardWidth - 1 {
    p.X = BoardWidth - p.Width - 1
  }

  if p.Y <= 1 {
    p.Y = 1
  }

  if p.Y + p.Height >= BoardHeight  {
    p.Y = BoardHeight - p.Height 
  }
}

func (p *Player) Update() {
  for _, b := range p.Bullets {
    b.Update()
  }
}

func (p *Player) Draw(board *Board) {
  board.DrawShape(p.X, p.Y, p.Shape)
  for _, b := range p.Bullets {
    b.Draw(board)
  }
}