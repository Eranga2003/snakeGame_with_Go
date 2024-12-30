package main

import (
    "fmt"               // To print text messages to the terminal
    "os"                // To handle errors and program exits
    "time"              // For pausing the game loop (sleep)
    "math/rand"         // For generating random numbers (used to spawn food)
    "github.com/nsf/termbox-go" // Terminal handling library for graphical interface
)

const (
    width  = 30  // Width of the game screen (number of columns)
    height = 10  // Height of the game screen (number of rows)
)

type Point struct {
    x, y int  // A struct to represent a point (position) on the screen
}

var snake []Point       // The snake is represented as a slice of Points
var direction Point     // The direction of the snake's movement (left, right, up, down)
var food Point          // Position of the food on the screen
var gameOver bool       // Boolean flag to check if the game is over

// Function to initialize the game state
func initGame() {
    // Initial snake position (a small snake)
    snake = []Point{{x: 5, y: 5}, {x: 4, y: 5}, {x: 3, y: 5}}
    // Initial movement direction (right)
    direction = Point{x: 1, y: 0}
    // Initial food position (fixed position, later to be randomized)
    food = Point{x: 10, y: 5}
    gameOver = false  // The game is not over at the start
}

// Function to draw the game state on the screen
func draw() {
    termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)  // Clear the terminal screen

    // Draw the snake (each segment is represented as 'O' in green)
    for _, s := range snake {
        termbox.SetCell(s.x, s.y, 'O', termbox.ColorGreen, termbox.ColorBlack)
    }

    // Draw the food (represented as '*' in red)
    termbox.SetCell(food.x, food.y, '*', termbox.ColorRed, termbox.ColorBlack)

    // Draw the game borders using '#' (walls)
    for i := 0; i < width; i++ {
        termbox.SetCell(i, 0, '#', termbox.ColorWhite, termbox.ColorBlack)   // Top border
        termbox.SetCell(i, height-1, '#', termbox.ColorWhite, termbox.ColorBlack) // Bottom border
    }
    for i := 0; i < height; i++ {
        termbox.SetCell(0, i, '#', termbox.ColorWhite, termbox.ColorBlack)   // Left border
        termbox.SetCell(width-1, i, '#', termbox.ColorWhite, termbox.ColorBlack) // Right border
    }

    termbox.Flush()  // Render all the changes to the screen
}

// Function to process user input (keyboard events)
func processInput() {
    event := termbox.PollEvent()  // Poll for user input (key press)
    switch event.Type {
    case termbox.EventKey:  // If a key was pressed
        switch event.Key {
        case termbox.KeyEsc:  // If ESC key is pressed, game ends
            gameOver = true
        case termbox.KeyArrowUp:  // If up arrow key is pressed, move snake up
            if direction.y != 1 {
                direction = Point{x: 0, y: -1}
            }
        case termbox.KeyArrowDown:  // If down arrow key is pressed, move snake down
            if direction.y != -1 {
                direction = Point{x: 0, y: 1}
            }
        case termbox.KeyArrowLeft:  // If left arrow key is pressed, move snake left
            if direction.x != 1 {
                direction = Point{x: -1, y: 0}
            }
        case termbox.KeyArrowRight:  // If right arrow key is pressed, move snake right
            if direction.x != -1 {
                direction = Point{x: 1, y: 0}
            }
        }
    }
}

// Function to update the game state (snake movement, collision checks, etc.)
func update() {
    head := snake[0]  // Get the current head of the snake
    newHead := Point{x: head.x + direction.x, y: head.y + direction.y}  // Calculate the new head position

    // Check for collision with walls (game over condition)
    if newHead.x <= 0 || newHead.x >= width-1 || newHead.y <= 0 || newHead.y >= height-1 {
        gameOver = true
        return
    }

    // Check if the snake collides with itself (game over condition)
    for _, s := range snake {
        if s.x == newHead.x && s.y == newHead.y {
            gameOver = true
            return
        }
    }

    // Move the snake: Add new head to the front and remove the tail
    snake = append([]Point{newHead}, snake[:len(snake)-1]...)

    // Check if the snake eats the food
    if newHead.x == food.x && newHead.y == food.y {
        // If food is eaten, grow the snake (add new head)
        snake = append([]Point{newHead}, snake...)
        // Generate a new random position for the food
        food = Point{x: randInt(1, width-2), y: randInt(1, height-2)}
    }
}

// Function to generate a random integer between min and max (inclusive)
func randInt(min, max int) int {
    return min + rand.Intn(max-min+1)
}

func main() {
    // Initialize termbox (to start drawing on the terminal)
    if err := termbox.Init(); err != nil {
        fmt.Println("Failed to initialize termbox:", err)  // Print error if initialization fails
        os.Exit(1)  // Exit the program if initialization fails
    }
    defer termbox.Close()  // Ensure that termbox resources are released when the program ends

    initGame()  // Initialize the game state

    // Main game loop
    for !gameOver {
        draw()        // Draw the current game state on the screen
        processInput() // Get and process user input (key press)
        update()       // Update the game state based on input and logic
        time.Sleep(100 * time.Millisecond) // Slow down the game to make it playable
    }

    // Game Over screen
    termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)  // Clear the screen
    // Display "GAME OVER" message at the center of the screen
    termbox.SetCell(width/2-4, height/2, 'G', termbox.ColorRed, termbox.ColorBlack)
    termbox.SetCell(width/2-3, height/2, 'A', termbox.ColorRed, termbox.ColorBlack)
    termbox.SetCell(width/2-2, height/2, 'M', termbox.ColorRed, termbox.ColorBlack)
    termbox.SetCell(width/2-1, height/2, 'E', termbox.ColorRed, termbox.ColorBlack)
    termbox.Flush()  // Render the "GAME OVER" message
    time.Sleep(2 * time.Second)  // Wait for a while before exiting
}
