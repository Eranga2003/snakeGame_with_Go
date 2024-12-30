package main

import (
    "fmt"
    "os"
    "time"
    "math/rand"
    "github.com/nsf/termbox-go"
)

var width, height int // These will now be dynamic

type Point struct {
    x, y int
}

var snake []Point
var direction Point
var food Point
var gameOver bool

// Initialize game state
func initGame() {
    snake = []Point{{x: 5, y: 5}, {x: 4, y: 5}, {x: 3, y: 5}}
    direction = Point{x: 1, y: 0}
    food = Point{x: randInt(1, width-2), y: randInt(1, height-2)}
    gameOver = false
}

// Draw game state
func draw() {
    termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

    // Draw snake
    for _, s := range snake {
        termbox.SetCell(s.x, s.y, 'O', termbox.ColorGreen, termbox.ColorBlack)
    }

    // Draw food
    termbox.SetCell(food.x, food.y, '*', termbox.ColorRed, termbox.ColorBlack)

    // Draw borders
    for i := 0; i < width; i++ {
        termbox.SetCell(i, 0, '#', termbox.ColorWhite, termbox.ColorBlack)
        termbox.SetCell(i, height-1, '#', termbox.ColorWhite, termbox.ColorBlack)
    }
    for i := 0; i < height; i++ {
        termbox.SetCell(0, i, '#', termbox.ColorWhite, termbox.ColorBlack)
        termbox.SetCell(width-1, i, '#', termbox.ColorWhite, termbox.ColorBlack)
    }

    termbox.Flush()
}

// Handle input from user
func processInput() {
    event := termbox.PollEvent()
    switch event.Type {
    case termbox.EventKey:
        switch event.Key {
        case termbox.KeyEsc:
            gameOver = true
        case termbox.KeyArrowUp:
            if direction.y != 1 {
                direction = Point{x: 0, y: -1}
            }
        case termbox.KeyArrowDown:
            if direction.y != -1 {
                direction = Point{x: 0, y: 1}
            }
        case termbox.KeyArrowLeft:
            if direction.x != 1 {
                direction = Point{x: -1, y: 0}
            }
        case termbox.KeyArrowRight:
            if direction.x != -1 {
                direction = Point{x: 1, y: 0}
            }
        }
    }
}

// Update game state
func update() {
    head := snake[0]
    newHead := Point{x: head.x + direction.x, y: head.y + direction.y}

    // Check if snake hits walls
    if newHead.x <= 0 || newHead.x >= width-1 || newHead.y <= 0 || newHead.y >= height-1 {
        gameOver = true
        return
    }

    // Check if snake hits itself
    for _, s := range snake {
        if s.x == newHead.x && s.y == newHead.y {
            gameOver = true
            return
        }
    }

    // Move the snake
    snake = append([]Point{newHead}, snake[:len(snake)-1]...)

    // Check if snake eats food
    if newHead.x == food.x && newHead.y == food.y {
        snake = append([]Point{newHead}, snake...) // Add new head
        food = Point{x: randInt(1, width-2), y: randInt(1, height-2)} // New food
    }
}

// Generate random number
func randInt(min, max int) int {
    return min + rand.Intn(max-min+1)
}

func main() {
    // Initialize termbox
    if err := termbox.Init(); err != nil {
        fmt.Println("Failed to initialize termbox:", err)
        os.Exit(1)
    }
    defer termbox.Close()

    // Get terminal dimensions
    width, height = termbox.Size()
    
    initGame()

    // Game loop
    for !gameOver {
        draw()
        processInput()
        update()
        time.Sleep(100 * time.Millisecond)
    }

    // Game Over screen
    termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
    termbox.SetCell(width/2-4, height/2, 'G', termbox.ColorRed, termbox.ColorBlack)
    termbox.SetCell(width/2-3, height/2, 'A', termbox.ColorRed, termbox.ColorBlack)
    termbox.SetCell(width/2-2, height/2, 'M', termbox.ColorRed, termbox.ColorBlack)
    termbox.SetCell(width/2-1, height/2, 'E', termbox.ColorRed, termbox.ColorBlack)
    termbox.Flush()
    time.Sleep(2 * time.Second)
}
