package main

import "time"
import "os"
import "strconv"
import "github.com/nsf/termbox-go"

type Game struct {
  State int
  Level int

  Counter int

  Board *Board

  Player *Player 
  Swarm *Swarm
}

const (
  StateNewGame = 0
  SelectedWeaponColor = termbox.ColorGreen
)

type Level struct {
  Monsters []Monster
}

func NewGame() *Game {
  g := &Game{Board: NewBoard(), State: StateNewGame, Player: NewPlayer()}
  g.NextLevel()
  return g
}

func (g *Game) Draw() {
  g.Board.DrawBorder()
  g.DrawStatus()

  g.Player.Draw(g.Board)

  for _, monster := range g.Swarm.Monsters {
    monster.Draw(g.Board)
  }
}

func (g *Game) NextLevel() {
  g.Level++
  g.Swarm = LoadLevel(g.Level)

  g.Player.Reset()
}

func (g *Game) Update() {
  if g.Swarm.AllDead() {
    g.NextLevel()
  }

  g.Player.Update()
  g.Swarm.Update()

}

func (g *Game) DrawStatus() {
  g.Board.DrawLine(BoardWidth + 1, 0, g.Player.Name)
  g.Board.DrawLine(BoardWidth + 1, 1, "X: " + strconv.Itoa(g.Player.X) + " Y: " + strconv.Itoa(g.Player.Y))
  g.Board.DrawLine(BoardWidth + 1, 2, "Lives: " +  itoa(g.Player.Lives) + " Score: " + itoa(g.Player.Points))


  weaponsStartRow := 8
  g.Board.DrawLine(BoardWidth + 1, weaponsStartRow, "Weapons: ")
  for i, weapon := range g.Player.AvailableWeapons {
    if i == g.Player.WeaponIndex {
      g.Board.DrawColoredLine(BoardWidth + 1, weaponsStartRow + i, itoa(i+1) + ") " + weapon.Name(), SelectedWeaponColor)
    } else {
      g.Board.DrawLine(BoardWidth + 1, weaponsStartRow + i, itoa(i+1) + ") " + weapon.Name())
    }
  }
}
  
func (g *Game) Play() {
  g.Board.SetInputHandler(g.Player)

  for {
    time.Sleep(time.Millisecond * 75)

    g.Update()
    g.Draw()

    g.CollisionDetection()

    g.Board.Sync()
    g.Counter++
  }
}

func (g *Game) PlayerDeath() {
  g.Player.Lives -= 1

  if g.Player.Lives <= 0 {
    os.Exit(1)
  }
}

func (g *Game) CollisionDetection() {
  for _, bullet := range g.Player.Bullets {
    if bullet.Alive {
      for _, monster := range g.Swarm.Monsters {
        if monster.Alive() && monster.InShape(bullet.X, bullet.Y) {
          monster.Lives--
          bullet.Alive = false
          g.Player.Points++
        }
      }
    }
  }

  for _, monster := range g.Swarm.Monsters {
    if monster.Alive() && IsCollision(monster.X, monster.Y, monster.Width, monster.Height,
      g.Player.X, g.Player.Y, g.Player.Width, g.Player.Height) {
      g.PlayerDeath()
    }

    for _, bullet := range monster.Bullets {
      if bullet.Alive && IsCollision(g.Player.X, g.Player.Y, g.Player.Width, g.Player.Height, bullet.X, bullet.Y, 1, 1) {
        g.PlayerDeath()
        bullet.Alive = false
      }
    }
  }

  for _, monster := range g.Swarm.Monsters {
    if monster.Alive() && monster.Y + monster.Height > BoardHeight {
      os.Exit(1)
    }
  }
}

func itoa(in int) string {
  return strconv.Itoa(in)
}

func IsCollision(x1, y1, width1, height1, x2, y2, width2, height2 int) bool {
  return (x1 < x2 + width2) && (x1 + width1 > x2) && (y2 < y1 + height1) && (y1 < y2 + height2)
}