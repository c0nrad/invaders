package main

import "github.com/nsf/termbox-go"

type Bullet struct {
  X, Y int

  Alive bool

  DirectionDelta int

  Shape string 
  Color termbox.Attribute
}

func NewBullet(x, y, direction int, shape string, color termbox.Attribute) *Bullet {
  return &Bullet{X: x, Y:y, DirectionDelta: direction, Alive: true, Shape: shape, Color: color}
}

func (b *Bullet) Update() {
  if b.Alive {
    b.Y += b.DirectionDelta
  }
  b.BoundsCheck()
}

func (b *Bullet) BoundsCheck() {
  if b.Y <= 0 || b.Y >= BoardHeight {
    b.Alive = false
  }
}

func (b *Bullet) Draw(board *Board) {
  if b.Alive{
    board.DrawColoredLine(b.X, b.Y, b.Shape, b.Color)
  }
}