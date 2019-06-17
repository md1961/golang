package main

import (
    "time"
    "fmt"
    "os"
    "bufio"
    "math/rand"
)


const (
    WIDTH = 40 //80
    HEIGHT = 15
    PCT_INITIAL_ALIVE = 25

    DEAD_DISPLAY = '.'
    LIFE_DISPLAY = 'X'
)


type Universe [][]bool


func NewUniverse(width int, height int) Universe {
    newUniverse := make(Universe, height)
    for i := range newUniverse {
        newUniverse[i] = make([]bool, width)
    }
    return newUniverse
}

func (u Universe) String() string {
    var cell byte
    buf := make([]byte, 0, (u.Width() + 1) * u.Height())
    for _, row := range u {
        for _, isAlive := range row {
            cell = DEAD_DISPLAY
            if isAlive {
                cell = LIFE_DISPLAY
            }
            buf = append(buf, cell)
        }
        buf = append(buf, '\n')
    }
    return string(buf)
}

func (u Universe) Seed() {
    for _, row := range u {
        for i, _ := range row {
            if rand.Intn(100) < PCT_INITIAL_ALIVE {
                row[i] = true
            }
        }
    }
}

func (u Universe) Width() int {
    return len(u[0])
}

func (u Universe) Height() int {
    return len(u)
}

func (u Universe) Alive(x, y int) bool {
    x = (x + u.Width())  % u.Width()
    y = (y + u.Height()) % u.Height()
    return u[y][x]
}

func (u Universe) Neighbors(x, y int) int {
    cell_offsets := []int{-1, 0, 1}
    count := 0
    for _, dy := range cell_offsets {
        for _, dx := range cell_offsets {
            if dx == 0 && dy == 0 {
                continue
            }
            if u.Alive(x + dx, y + dy) {
                count++
            }
        }
    }
    return count
}

func (u Universe) Next(x, y int) bool {
    n := u.Neighbors(x, y)
    return n == 3 || (u.Alive(x, y) && n == 2)
}

func (u Universe) Set(x, y int, value bool) {
    u[y][x] = value
}

func (u Universe) NextUniverse() Universe {
    nextUniverse := NewUniverse(u.Width(), u.Height())
    for y := 0; y < u.Height(); y++ {
        for x := 0; x < u.Width(); x++ {
            nextUniverse.Set(x, y, u.Next(x, y))
        }
    }
    return nextUniverse
}


func main() {
    rand.Seed(time.Now().UnixNano())
    stdin := bufio.NewScanner(os.Stdin)

    universe := NewUniverse(WIDTH, HEIGHT)
    universe.Seed()

    for {
        fmt.Print(universe.String())

        stdin.Scan()
        input := stdin.Text()
        if len(input) > 0 && input[0] == 'q' {
            break
        }

        universe = universe.NextUniverse()
    }
}
