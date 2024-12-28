package main

// Version 6.0
// Invaders forked from https://github.com/sausheong/invaders
// Modified by RangerSmyth & David T Keane
// Date: 27-28th December 2024
// Seasons Greetings and Happy Holidays from the developers

/*
------------------------- HELP SECTION (CONFIGURATION) -------------------------

    Game Settings:

    - windowWidth, windowHeight: Adjust the dimensions of the game window.
    - aliensPerRow: Set the number of aliens per row.
    - aliensStartCol: Set the starting column position for aliens.
    - alienSize: Adjust the size of the alien sprites.
    - bombProbability: Modify the probability of aliens dropping bombs (0.0 to 1.0).
    - bombSpeed: Change the speed at which bombs fall.
    - barrierYPosition: Adjust the vertical position of the barriers.
    - playerYPosition: Set the initial vertical position of the player's cannon.

    Audio Settings:

    - You can adjust the volume of each sound effect by modifying the volume
      when creating the audio player in the loadAudio function.
      Example:
          audioStream, err := audioContext.NewPlayer(wavStream)
          audioStream.SetVolume(0.5) // Sets volume to 50%
      The volume parameter is a float64 between 0.0 (silent) and 1.0 (full volume).

    Control Settings:

    - The game controls are currently hardcoded in the Update() function.
    - To change the controls, modify the ebiten.IsKeyPressed() and
      inpututil.IsKeyJustPressed() functions within Update().
      For example, to change the key for moving the cannon to the right:
          if ebiten.IsKeyPressed(ebiten.KeyD) { // Change from KeyArrowRight to KeyD
              laserCannon.Position.X += 10
          }

    Adding a Settings Panel:

    - To add a settings panel that can be opened/closed with a button:
      1. Create a new struct for managing settings.
      2. Implement drawing functions for the panel.
      3. Add a button to toggle the panel in the main game's Draw() function.
      4. Handle input for the settings panel in the Update() function.

    Using a Switch Statement (Illustrative Example):

    - Go does not have a traditional switch statement for types like in C++ or Java.
    - You can use a type switch or a series of if-else statements for similar functionality.
      Example (Type Switch):
          switch v := someVariable.(type) {
          case int:
              fmt.Println("Integer:", v)
          case string:
              fmt.Println("String:", v)
          default:
              fmt.Println("Unknown Type")
          }
      Example (If-Else):
          if someVariable == "option1" {
              // Do something
          } else if someVariable == "option2" {
              // Do something else
          }

    Managing Multiple Changes and Ideas:

    - It's recommended to use a version control system like Git to manage changes.
    - Break down each change or idea into smaller, manageable tasks.
    - Use comments and documentation to keep track of changes and future plans.
    - Consider creating separate Go files for different parts of the game (e.g., settings, UI, game logic)
      to improve organization and readability.

--------------------------------------------------------------------------------
*/

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"os/user"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

var (
	windowWidth      = 800
	windowHeight     = 600
	aliensPerRow     = 8
	aliensStartCol   = 100
	alienSize        = 30
	bombProbability  = 0.005
	bombSpeed        = 10
	barrierYPosition = 300
	playerYPosition  = 400

	gameOverMessageYOffset = 100
	finalScoreYOffset      = 150
	playAgainYOffset       = 100
	closeGameYOffset       = 100
	highScoresTitleYOffset = 160
	highScoresListYOffset  = 30
	highScoresListSpacing  = 20
)

var (
	src           *ebiten.Image
	background    *ebiten.Image
	backgroundEnd *ebiten.Image
	cannonSprite  = image.Rect(20, 47, 38, 59)
	cannonExplode = image.Rect(0, 47, 16, 57)
	alien1Sprite  = image.Rect(0, 0, 20, 14)
	alien1aSprite = image.Rect(20, 0, 40, 14)
	alien2Sprite  = image.Rect(0, 14, 20, 26)
	alien2aSprite = image.Rect(20, 14, 40, 26)
	alien3Sprite  = image.Rect(0, 27, 20, 40)
	alien3aSprite = image.Rect(20, 27, 40, 40)
	alienExplode  = image.Rect(0, 60, 16, 68)
	beamSprite    = image.Rect(20, 60, 22, 65)
	bombSprite    = image.Rect(0, 70, 10, 79)
	barrierSprite = image.Rect(40, 0, 50, 10)
)

type Sprite struct {
	size     image.Rectangle
	Filter   *ebiten.Image
	FilterA  *ebiten.Image
	FilterE  *ebiten.Image
	Position image.Point
	Status   bool
	Points   int
}

var (
	aliens      = []Sprite{}
	bombs       = []Sprite{}
	barriers    = []Sprite{}
	laserCannon Sprite
	beam        Sprite
)

var (
	laserSound         *audio.Player
	explosionSound     *audio.Player
	gameOverSound      *audio.Player
	backgroundSound    *audio.Player
	endGameSound       *audio.Player
	shipExplosionSound *audio.Player
)

var audioContext *audio.Context

var (
	gameFont     font.Face
	gameOverFont font.Face
)

var (
	highScores []HighScore
	playerName string
)

const maxHighScores = 5

type HighScore struct {
	Name  string
	Score int
}

func createAlien(x, y int, sprite, alt image.Rectangle, points int) (s Sprite) {
	s = Sprite{
		size:     sprite,
		Filter:   src.SubImage(sprite).(*ebiten.Image),
		FilterA:  src.SubImage(alt).(*ebiten.Image),
		FilterE:  src.SubImage(alienExplode).(*ebiten.Image),
		Position: image.Pt(x, y),
		Status:   true,
		Points:   points,
	}
	return
}
func loadFont(path string, size float64) font.Face {
	fontBytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	ttfFont, err := truetype.Parse(fontBytes)
	if err != nil {
		log.Fatal(err)
	}
	return truetype.NewFace(ttfFont, &truetype.Options{
		Size: size,
		DPI:  72,
	})
}
func loadAudio(path string) *audio.Player {
	fileBytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	var audioStream *audio.Player
	if path[len(path)-3:] == "wav" {
		wavStream, err := wav.DecodeWithSampleRate(audioContext.SampleRate(), bytes.NewReader(fileBytes))
		if err != nil {
			log.Fatal(err)
		}
		audioStream, err = audioContext.NewPlayer(wavStream)
		if err != nil {
			log.Fatal(err)
		}

	} else if path[len(path)-3:] == "mp3" {
		mp3Stream, err := mp3.DecodeWithSampleRate(audioContext.SampleRate(), bytes.NewReader(fileBytes))
		if err != nil {
			log.Fatal(err)
		}
		audioStream, err = audioContext.NewPlayer(mp3Stream)
		if err != nil {
			log.Fatal(err)
		}

	}

	audioStream.SetVolume(0.5)

	return audioStream
}
func createBarrier(x, y int) (s Sprite) {
	s = Sprite{
		size:     barrierSprite,
		Filter:   src.SubImage(barrierSprite).(*ebiten.Image),
		Position: image.Pt(x, y),
		Status:   true,
	}
	return
}

func initGame() {
	imgFile, _, err := ebitenutil.NewImageFromFile("imgs/sprites.png")
	if err != nil {
		panic(err)
	}
	src = imgFile

	bg, _, err := ebitenutil.NewImageFromFile("imgs/bg.png")
	if err != nil {
		panic(err)
	}
	background = bg

	bgEnd, _, err := ebitenutil.NewImageFromFile("imgs/background-end3.png")
	if err != nil {
		log.Fatal("Error loading background-end3.jpg:", err) // Or panic
	}
	backgroundEnd = bgEnd

	laserCannon = Sprite{
		size:     cannonSprite,
		Filter:   src.SubImage(cannonSprite).(*ebiten.Image),
		FilterE:  src.SubImage(cannonExplode).(*ebiten.Image),
		Position: image.Pt(50, playerYPosition),
		Status:   true,
	}

	beam = Sprite{
		size:     beamSprite,
		Filter:   src.SubImage(beamSprite).(*ebiten.Image),
		Position: image.Pt(laserCannon.Position.X+7, 250),
		Status:   false,
	}

	rows := 5
	cols := 12
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			x := aliensStartCol + col*(alienSize+10)
			y := 30 + row*30
			points := 10
			if row == 0 {
				points = 30
				aliens = append(aliens, createAlien(x, y, alien1Sprite, alien1aSprite, points))
			} else if row == 1 || row == 2 {
				points = 20
				aliens = append(aliens, createAlien(x, y, alien2Sprite, alien2aSprite, points))
			} else if row == 3 || row == 4 {
				points = 10
				aliens = append(aliens, createAlien(x, y, alien3Sprite, alien3aSprite, points))
			}
		}
	}
	barrierWidth := (windowWidth - 100) / 4
	barriers = append(barriers, createBarrier(100, barrierYPosition))
	barriers = append(barriers, createBarrier(100+barrierWidth, barrierYPosition))
	barriers = append(barriers, createBarrier(100+2*barrierWidth, barrierYPosition))

	gameFont = loadFont("font/font.ttf", 24)
	gameOverFont = loadFont("font/font.ttf", 56)

	laserSound = loadAudio("files/laser.wav")
	explosionSound = loadAudio("files/explosion.wav")
	gameOverSound = loadAudio("files/game-over.mp3")
	backgroundSound = loadAudio("files/background.wav")
	endGameSound = loadAudio("files/end-game.mp3")
	shipExplosionSound = loadAudio("files/explosion-sound.mp3")

	loadHighScores()
}

func loadHighScores() {
	highScores = []HighScore{}
	content, err := ioutil.ReadFile("files/highscores.txt")
	if err == nil {
		lines := strings.Split(string(content), "\n")

		for _, line := range lines {
			if strings.TrimSpace(line) == "" {
				continue
			}
			parts := strings.Split(line, ",")
			if len(parts) != 2 {
				continue
			}
			name := parts[0]
			score, err := strconv.Atoi(parts[1])
			if err != nil {
				continue
			}
			highScores = append(highScores, HighScore{Name: name, Score: score})
		}
	}
	sortHighScores()
}

func saveHighScores() {
	var sb strings.Builder
	for _, score := range highScores {
		sb.WriteString(score.Name)
		sb.WriteString(",")
		sb.WriteString(strconv.Itoa(score.Score))
		sb.WriteString("\n")
	}
	ioutil.WriteFile("files/highscores.txt", []byte(sb.String()), 0644)
}
func sortHighScores() {
	sort.Slice(highScores, func(i, j int) bool {
		return highScores[i].Score > highScores[j].Score
	})
	if len(highScores) > maxHighScores {
		highScores = highScores[:maxHighScores]
	}
}

func addHighScore(score int) {
	if len(highScores) == maxHighScores && score <= highScores[maxHighScores-1].Score {
		return
	}

	for _, existingScore := range highScores {
		if existingScore.Name == playerName && existingScore.Score == score {
			return
		}
	}

	highScores = append(highScores, HighScore{Name: playerName, Score: score})

	sortHighScores()
	saveHighScores()
}

//   This is the end of Part 1

//   Part 1 Summary:
//   This part of the code sets up the game environment, including:
// - Importing necessary packages for game development, audio handling, and file operations.
// - Defining game parameters such as window dimensions, alien configurations, and game over screen settings.
// - Initializing sprites, which are the visual elements of the game like aliens, the player's cannon, and

// This is the Start of Part 2

type Game struct { // Main Game struct — add fields here!
	loop             int
	beamShot         bool
	gameOver         bool
	alienDirection   int
	score            int
	startScreen      *ebiten.Image
	gameFont         font.Face
	gameOverFont     font.Face
	lives            int
	isPaused         bool
	gameOverTimer    int
	showGameOverText bool // Fields correctly placed in the main Game struct
}

func (g *Game) Update() error { // Correct Update function – no local Game struct
	if g.gameOver {
		g.gameOverTimer++ // Now refers to g.gameOverTimer of the *main* Game struct
		if g.gameOverTimer%60 == 0 {
			g.showGameOverText = !g.showGameOverText
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			g.isPaused = false
			g.resetGame()
			return nil
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
			os.Exit(0)
		}
		return nil
	}

	if g.isPaused {
		return nil
	}

	if !g.gameOver {
		if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
			laserCannon.Position.X += 10
		}
		if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
			laserCannon.Position.X -= 10
		}
		if ebiten.IsKeyPressed(ebiten.KeyDown) {
			playerYPosition = min(windowHeight-50, playerYPosition+5)
			for i := range barriers {
				barriers[i].Position.Y = barrierYPosition + 5
			}
		}
		if ebiten.IsKeyPressed(ebiten.KeyUp) {
			playerYPosition = max(100, playerYPosition-5)
			for i := range barriers {
				barriers[i].Position.Y = barrierYPosition - 5
			}
		}

		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			if !beam.Status {
				g.beamShot = true
				if laserSound != nil {
					laserSound.Rewind()
					laserSound.Play()
				}
			}
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
			g.gameOver = true
			g.isPaused = true
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
			g.isPaused = !g.isPaused
		}

		if g.beamShot {
			beam.Position.X = laserCannon.Position.X + 7
			beam.Status = true
			g.beamShot = false
		}
	}

	if aliens[0].Position.X < alienSize || aliens[aliensPerRow-1].Position.X > windowWidth-(2*alienSize) {
		g.alienDirection = g.alienDirection * -1
		for i := 0; i < len(aliens); i++ {
			aliens[i].Position.Y = aliens[i].Position.Y + 10
		}
	}
	return nil
}

// Part 2: Game Rendering and Logic

func (g *Game) drawGameOverScreen(screen *ebiten.Image) {
	// Check if backgroundEnd is loaded
	if backgroundEnd != nil {
		// Calculate scale factors for the game over background
		bgWidth, bgHeight := backgroundEnd.Bounds().Dx(), backgroundEnd.Bounds().Dy()
		xScale := float64(windowWidth) / float64(bgWidth)
		yScale := float64(windowHeight) / float64(bgHeight)
		// Draw the game over background image with scaling
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(xScale, yScale)
		screen.DrawImage(backgroundEnd, op) // Draw the image
	} else {
		// If backgroundEnd is not loaded, use a default black background
		// screen.Fill(color.Black) // Background colour, change to modify - Remove the // at the start of this line to have a black background and no image.
	}

	// Play game over sound
	if gameOverSound != nil && !gameOverSound.IsPlaying() {
		gameOverSound.Rewind()
		gameOverSound.Play()
	}

	// Define box parameters (adjust these as needed)
	boxWidth := 400
	boxHeight := 500                       // Increased height to make room for scores
	boxX := (windowWidth - boxWidth) / 2   // Center the box horizontally
	boxY := (windowHeight - boxHeight) / 4 // Center vertically, adjust as needed

	// Draw a filled rectangle for the box
	ebitenutil.DrawRect(screen, float64(boxX), float64(boxY), float64(boxWidth), float64(boxHeight), color.RGBA{0, 0, 0, 0}) // Semi-transparent black

	// Define the message and the "Try Again" button text

	if g.showGameOverText { // Draw text conditionally
		// Define the message and the "Try Again" button text

		message := fmt.Sprintf("GAME OVER!\n\nFinal score: %d", g.score)
		tryAgain := "Press Enter to Play again"
		closeGame := "Press Esc to close the game"

		// Get text bounds to calculate the center position
		messageBounds := text.BoundString(g.gameOverFont, message)
		tryAgainBounds := text.BoundString(g.gameFont, tryAgain)
		closeGameBounds := text.BoundString(g.gameFont, closeGame)
		highScoreTitleBounds := text.BoundString(g.gameFont, "High Scores:") // Get bounds for title

		// Calculate positions relative to the box
		x := boxX + (boxWidth-messageBounds.Dx())/2
		y := boxY + 30 // Adjust vertical position within the box

		// Position for "Play Again" text
		xTryAgain := boxX + (boxWidth-tryAgainBounds.Dx())/2
		yTryAgain := y + messageBounds.Dy() + playAgainYOffset - 50 // Increased spacing

		// Position for "Close Game" text - put it below "Play Again"
		xCloseGame := boxX + (boxWidth-closeGameBounds.Dx())/2
		yCloseGame := yTryAgain + tryAgainBounds.Dy() + closeGameYOffset - 70 // Decreased spacing

		// Position for "High Scores" title
		xHighScoreTitle := boxX + (boxWidth-highScoreTitleBounds.Dx())/2
		yHighScoreTitle := yCloseGame + closeGameBounds.Dy() + highScoresTitleYOffset - 100 // Position below "Close Game"

		// Draw the text
		text.Draw(screen, message, g.gameOverFont, x, y, color.White)
		text.Draw(screen, tryAgain, g.gameFont, xTryAgain, yTryAgain, color.White)
		text.Draw(screen, closeGame, g.gameFont, xCloseGame, yCloseGame, color.White)
		text.Draw(screen, "High Scores:", g.gameFont, xHighScoreTitle, yHighScoreTitle, color.White) // High scores title

		// Draw the high scores list
		yHighScore := yHighScoreTitle + highScoreTitleBounds.Dy() + highScoresListSpacing + 10 // Start below the title
		for i, score := range highScores {
			scoreText := fmt.Sprintf("%d. %s: %d", i+1, score.Name, score.Score)
			scoreTextBounds := text.BoundString(g.gameFont, scoreText)
			xHighScore := boxX + (boxWidth-scoreTextBounds.Dx())/2 // Center each score within the box
			text.Draw(screen, scoreText, g.gameFont, xHighScore, yHighScore, color.White)
			yHighScore += scoreTextBounds.Dy() + 5
		}
	}

}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.gameOver {
		g.drawGameOverScreen(screen)
		return
	}

	g.drawGameScreen(screen)
}

func (g *Game) drawGameScreen(screen *ebiten.Image) {
	bgWidth, bgHeight := background.Bounds().Dx(), background.Bounds().Dy()
	xScale := float64(windowWidth) / float64(bgWidth)
	yScale := float64(windowHeight) / float64(bgHeight)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(xScale, yScale)
	screen.DrawImage(background, op)

	for _, barrier := range barriers {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(barrier.Position.X), float64(barrier.Position.Y))
		screen.DrawImage(barrier.Filter, op)
	}

	for i := 0; i < len(aliens); i++ {
		aliens[i].Position.X = aliens[i].Position.X + 5*g.alienDirection
		if aliens[i].Status {
			if collide(aliens[i], beam) {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(aliens[i].Position.X), float64(aliens[i].Position.Y))
				screen.DrawImage(aliens[i].FilterE, op)
				aliens[i].Status = false
				g.score += aliens[i].Points
				if explosionSound != nil {
					explosionSound.Rewind()
					explosionSound.Play()
				}
				resetBeam()
			} else {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(aliens[i].Position.X), float64(aliens[i].Position.Y))
				if g.loop%2 == 0 {
					screen.DrawImage(aliens[i].Filter, op)
				} else {
					screen.DrawImage(aliens[i].FilterA, op)
				}
			}

			if rand.Float64() < bombProbability {
				dropBomb(aliens[i])
			}
		}
	}

	for i := 0; i < len(bombs); i++ {
		bombs[i].Position.Y = bombs[i].Position.Y + bombSpeed
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(bombs[i].Position.X), float64(bombs[i].Position.Y))
		screen.DrawImage(bombs[i].Filter, op)
		if collide(bombs[i], laserCannon) {
			g.lives--
			if g.lives <= 0 {
				g.gameOver = true
				addHighScore(g.score)
				if endGameSound != nil {
					endGameSound.Rewind()
					endGameSound.Play()
				}
			} else {
				resetBeam()
				laserCannon.Position.Y = playerYPosition
			}
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(laserCannon.Position.X), float64(laserCannon.Position.Y))
			screen.DrawImage(laserCannon.FilterE, op)
			if shipExplosionSound != nil {
				shipExplosionSound.Rewind()
				shipExplosionSound.Play()
			}
		}
	}
	if !g.gameOver {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(laserCannon.Position.X), float64(laserCannon.Position.Y))
		screen.DrawImage(laserCannon.Filter, op)
	}

	if beam.Status {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(beam.Position.X), float64(beam.Position.Y))
		screen.DrawImage(beam.Filter, op)
		beam.Position.Y -= 10
	}
	if beam.Position.Y < 0 {
		resetBeam()
	}

	for i := range aliens {
		if aliens[i].Position.Y > playerYPosition-50 {
			g.gameOver = true
			addHighScore(g.score)
			if endGameSound != nil {
				endGameSound.Rewind()
				endGameSound.Play()
			}
		}
	}
	g.loop++
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Score: %d    Lives: %d", g.score, g.lives))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return windowWidth, windowHeight
}

func dropBomb(alien Sprite) {
	torpedo := Sprite{
		size:     bombSprite,
		Filter:   src.SubImage(bombSprite).(*ebiten.Image),
		Position: image.Pt(alien.Position.X+7, alien.Position.Y),
		Status:   true,
	}

	bombs = append(bombs, torpedo)
}

func resetBeam() {
	beam.Status = false
	beam.Position.Y = 250
}

func collide(s1, s2 Sprite) bool {
	spriteA := image.Rect(s1.Position.X, s1.Position.Y, s1.Position.X+s1.size.Dx(), s1.Position.Y+s1.size.Dy())
	spriteB := image.Rect(s2.Position.X, s2.Position.Y, s2.Position.X+s1.size.Dx(), s2.Position.Y+s1.size.Dy())
	if spriteA.Min.X < spriteB.Max.X && spriteA.Max.X > spriteB.Min.X &&
		spriteA.Min.Y < spriteB.Max.Y && spriteA.Max.Y > spriteB.Min.Y {
		return true
	}
	return false
}

func (g *Game) resetGame() {
	g.loop = 0
	g.beamShot = false
	g.gameOver = false
	g.alienDirection = 1
	g.score = 0
	g.lives = 3

	aliens = []Sprite{}
	bombs = []Sprite{}

	rows := 5
	cols := 12
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			x := aliensStartCol + col*(alienSize+10)
			y := 30 + row*30
			points := 10
			if row == 0 {
				points = 30
				aliens = append(aliens, createAlien(x, y, alien1Sprite, alien1aSprite, points))
			} else if row == 1 || row == 2 {
				points = 20
				aliens = append(aliens, createAlien(x, y, alien2Sprite, alien2aSprite, points))
			} else if row == 3 || row == 4 {
				points = 10
				aliens = append(aliens, createAlien(x, y, alien3Sprite, alien3aSprite, points))
			}
		}
	}

	laserCannon.Position = image.Pt(50, playerYPosition)
	beam.Position = image.Pt(laserCannon.Position.X+7, 250)

	if backgroundSound != nil {
		backgroundSound.Rewind()
		backgroundSound.Play()
	}
}

// 	 End of Part 2

// 	 Part 2 Summary:

//   This section defines the core game logic and rendering functions. Key components include:
// - Game struct: Holds the game state variables like score, lives, and game loop counter.
// - Update() function: Handles game logic updates, including player input, alien movement, and game over conditions.
// - drawGameOverScreen() function: Renders the game over screen with the final score, high scores, and options to restart or quit.
// - Draw() function: The main rendering function that calls either drawGameOverScreen() or drawGameScreen() based on the game state.
// - drawGameScreen() function: Renders the game elements like the background, barriers, aliens, bombs, laser cannon, and beam.
// - Layout() function: Defines the game's screen layout.
// - Helper functions: Include collision detection (collide), bomb dropping (dropBomb), beam resetting (resetBeam), and game reset (resetGame).

// 	 This is the Start of Part 3

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("Space Invaders")

	audioContext = audio.NewContext(48000)

	laserSound = loadAudio("files/laser.wav")
	explosionSound = loadAudio("files/explosion.wav")
	gameOverSound = loadAudio("files/game-over.mp3")
	backgroundSound = loadAudio("files/background.wav")
	endGameSound = loadAudio("files/end-game.mp3")
	shipExplosionSound = loadAudio("files/explosion-sound.mp3")

	user, err := user.Current()
	if err == nil {
		playerName = user.Username
	} else {
		playerName = "Player"
	}

	game := &Game{
		loop:             0,
		beamShot:         false,
		gameOver:         false,
		alienDirection:   1,
		score:            0,
		gameFont:         loadFont("font/font.ttf", 24),
		gameOverFont:     loadFont("font/font.ttf", 28),
		lives:            3,
		isPaused:         false,
		gameOverTimer:    0,
		showGameOverText: true, // Initial state
	}
	initGame()
	if backgroundSound != nil {
		backgroundSound.Rewind()
		backgroundSound.Play()
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

// This is the End of Part 3
