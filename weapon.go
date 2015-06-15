package main

import "github.com/nsf/termbox-go"

var Rainbow = []termbox.Attribute{termbox.ColorRed, termbox.ColorGreen, termbox.ColorYellow, termbox.ColorBlue, termbox.ColorMagenta, termbox.ColorCyan}

var RainbowIndex = 0
func NextRainbowColor() termbox.Attribute {
  color := Rainbow[RainbowIndex]
  RainbowIndex = (RainbowIndex + 1) % len(Rainbow)
  return color
}

type Weapon interface {
  Name() string
  Fire(x, y, direction int) []*Bullet
}

var WeaponMap = map[int]Weapon {
  0: PlainWeapon{"Plain", "*", termbox.ColorDefault},
  1: RainbowWeapon{"Rainbow", "*"},
  2: TridentWeapon{"Trident", "*"},
}

type PlainWeapon struct {
  name string
  shape string
  color termbox.Attribute
}

func (w PlainWeapon) Name() string {
  return w.name
}

func (w PlainWeapon) Fire(x, y, direction int) []*Bullet {
  return []*Bullet{NewBullet(x, y, direction, w.shape, w.color)}
}

type RainbowWeapon struct {
  name string
  shape string
}

func (w RainbowWeapon) Name() string {
  return w.name
}

func (w RainbowWeapon) Fire(x, y, direction int) []*Bullet {
  return []*Bullet{NewBullet(x, y, direction, w.shape, NextRainbowColor())}
}

type TridentWeapon struct {
  name string
  shape string
}

func (w TridentWeapon) Name() string {
  return w.name
}

func (w TridentWeapon) Fire(x, y, direction int) []*Bullet {
  return []*Bullet{
    NewBullet(x-3, y, direction, w.shape, NextRainbowColor()),
    NewBullet(x, y, direction, w.shape, NextRainbowColor()),
    NewBullet(x+3, y, direction, w.shape, NextRainbowColor()),
  }
}


