package main

import "github.com/nsf/termbox-go"
import "strings"
import "os"

const (
  BoardWidth = 50
  BoardHeight = 50

  MaxEvents = 1000
)

const (
  CornerPiece = "+"
  UpPiece = "|"
  FlatPiece = "-"
)

const (
  MoveUp = iota
  MoveRight
  MoveDown
  MoveLeft
  Spacebar
)

type InputHandler interface {
  Handle(termbox.Event)
}

type Board struct {
  Width, Height int

  EventQueue chan termbox.Event
  InputHandler InputHandler
}

func Wait() {
  termbox.PollEvent()
}

func NewBoard() *Board {
  err := termbox.Init()
  if err != nil {
    panic(err)
  }
  termbox.Sync()
  b := &Board{Width: BoardWidth, Height: BoardHeight, EventQueue: make(chan termbox.Event, MaxEvents)}
  b.DrawBorder()
  go b.HandleInput()
  return b
}

func (b *Board) SetInputHandler(i InputHandler) {
  b.InputHandler = i
}

func (b *Board) HandleInput() {
  for {
    event := termbox.PollEvent()
    if event.Type == termbox.EventKey && (event.Key == termbox.KeyEsc || event.Ch == 'q') {
      os.Exit(0)
    }

    b.InputHandler.Handle(event)
  }
}

func (b *Board) Input() {
  termbox.PollEvent()
}

func (b *Board) DrawShape(x, y int, shape []string) {
  for dy, line := range shape {
    b.DrawLine(x, y+dy, line)
  }
}

func (b *Board) DrawLine(x, y int, line string) {
  for dx, c := range line {
    termbox.SetCell(x + dx, y, c,termbox.ColorDefault, termbox.ColorDefault)
  }
}

func (b *Board) DrawColoredLine(x, y int, line string, color termbox.Attribute) {
  for dx, c := range line {
    termbox.SetCell(x + dx, y, c, color, termbox.ColorDefault)
  }
}

func (b *Board) Sync() {
  termbox.Sync()
}

func (b *Board) DrawBorder() {
  topLine := CornerPiece + strings.Repeat(FlatPiece, BoardWidth-2) + CornerPiece
  b.DrawLine(0, 0, topLine)

  middleLine := UpPiece + strings.Repeat(" ", b.Width - 2) + UpPiece 
  for i := 1; i < b.Height; i++ {
    b.DrawLine(0, i, middleLine)
  }

  b.DrawLine(0, b.Height, topLine)
}