package main

import "math/rand"

type Monster struct {
  X, Y int
  Width, Height int

  Lives int

  Weapon Weapon
  FirePercent int

  Shape []string
  Bullets []*Bullet
}

func NewMonster(x, y, lives, firePercent int, shape []string, weapon Weapon) *Monster {
  width := len(shape[0])
  height := len(shape)
  m := &Monster{X: x, Y: y, Width: width, Height: height, 
    Lives: lives, Shape: shape, Weapon: weapon, FirePercent: firePercent}
  return m
}

func (m *Monster) Draw(board *Board) {
  if m.Alive() {
    board.DrawShape(m.X, m.Y, m.Shape) 
  }
  for _, b := range m.Bullets {
    b.Draw(board)
  }
}

func (m *Monster) Alive() bool {
  return m.Lives > 0
}

func (m *Monster) FireWeapon() {
  x := m.X + m.Width / 2
  y := m.Y
  m.Bullets = append(m.Bullets, m.Weapon.Fire(x, y, 1)...)
}

func (m *Monster) InShape(x, y int) bool {
  if (x >= m.X && x <= m.X + m.Width) && (y >= m.Y && y <= m.Y + m.Height) {
    return true
  }
  return false
}

func (m *Monster) Update() {

  if m.Alive() {
    if rand.Intn(100) < m.FirePercent {
      m.FireWeapon()
    }
  }

  for _, b := range m.Bullets {
    b.Update()
  }
}