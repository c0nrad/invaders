package main

// import "fmt"

type Swarm struct {
  Rows, Cols int

  MoveRight bool

  Monsters []*Monster 
}

func NewSwarm(rows int, m Monster) *Swarm {
  startX := 1
  startY := 1
  shapeWidth := len(m.Shape[0])
  shapeHeight := len(m.Shape)

  monstersInCol := (BoardWidth / (shapeWidth +2)) - 1

  monsters := []*Monster{}

  for x := startX; x < BoardWidth - shapeWidth - 2; x += (shapeWidth + 2) {
    for y := startY ; y < rows * (shapeHeight+1) + startY; y += shapeHeight + 1 {

      m := NewMonster(x, y, m.Lives, m.FirePercent, m.Shape, m.Weapon)
      monsters = append(monsters, m)
    }
  }
  return &Swarm{Monsters: monsters, Rows: rows, Cols: monstersInCol}
}

func (s *Swarm) AllDead() bool {
  for _, m := range s.Monsters {
    if m.Alive() {
      return false
    }
  }
  return true
}

func (s *Swarm) RightMost() int {
  maxX := 0
  for _, m := range s.Monsters {
    x := m.X + len(m.Shape[0])
    if x >= maxX {
      maxX = x
    }
  }
  return maxX
}

func (s *Swarm) LeftMost() int {
  minX := BoardWidth
  for _, m := range s.Monsters {
    x := m.X 
    if x < minX {
      minX = x
    }
  }
  return minX
}

func (s *Swarm) Update() {
  maxX := s.RightMost()
  minX := s.LeftMost()

  deltaX := 1
  deltaY := 0
  if maxX >= BoardWidth -1 {
    s.MoveRight = ! s.MoveRight
    deltaY = 1
  } else if minX <= 1 {
    s.MoveRight = ! s.MoveRight
    deltaY = 1
  }

  if s.MoveRight {
    deltaX = 1
  } else {
    deltaX = -1
  }

  for _, m := range s.Monsters {
    m.X += deltaX 
    m.Y += deltaY

    m.Update()
  }

}