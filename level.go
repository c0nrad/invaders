package main

func LoadLevel(level int) *Swarm {
  if level == 4 || level == 8 || level == 12 || level == 13 {
   
    return LoadBoss(level/4)
  }

  rows := (level % 4) 
  monsterType := level / 4  
  monsterShape := MonsterShapeMap[monsterType]
  monsterWeapon := WeaponMap[monsterType]
  monsterFirePercent := level / 2
  monsterLives := monsterType * 2 + 1 

  m := NewMonster(0, 0, monsterLives, monsterFirePercent,
    monsterShape, monsterWeapon)

  swarm := NewSwarm(rows, *m)
  return swarm
}

func LoadBoss(index int) *Swarm {
  m := &Monster{}
  if index == 1 {
    m = NewMonster(2, 2, 100, 45, BossShape[index-1], WeaponMap[1])
  } else if index == 2{
    m = NewMonster(2, 2, 200, 45, BossShape[index-1], WeaponMap[1])
  } else if index == 3{
    m = NewMonster(2, 2, 300, 45, BossShape[index-1], WeaponMap[2])
  }


  return &Swarm{Monsters: []*Monster{m}}
}

var MonsterShapeMap = map[int][]string{
  0: []string{"|-|", " v "},
  1: []string{"||-||", " VvV "},
  2: []string{"|-|=|-|", " V v V "},
}

var BossShape = map[int][]string{
  0: []string{
   "||-----|| ||-----||",
   "    ---|| ||---    ",
   "      -|| ||-      ",
   "       || ||       ",
   "         V         ",
  },

  1: []string{
    "|                 |",
    "|                 |",
    "|------V-V-V------|",
    "|                 |",
    "|                 |",
  },

  2: []string{
    "         /\\         ",
    "        /  \\        ",
    "       /    \\       ",
    "      /      \\      ",
    "     /        \\     ",
    "    /          \\    ",
    "   /            \\   ", 
    "  /              \\  ",
    " /                \\ ",
    "/                  \\",
    "--------------------",
  },
}
