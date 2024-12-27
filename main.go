package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"io/ioutil"
	"log"
	"math/rand"
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

// parameters
var windowWidth, windowHeight = 800, 600
var aliensPerRow = 8
var aliensStartCol = 100
var alienSize = 30
var bombProbability = 0.005
var bombSpeed = 10

// sprites
var src *ebiten.Image
var background *ebiten.Image
var cannonSprite = image.Rect(20, 47, 38, 59)
var cannonExplode = image.Rect(0, 47, 16, 57)
var alien1Sprite = image.Rect(0, 0, 20, 14)
var alien1aSprite = image.Rect(20, 0, 40, 14)
var alien2Sprite = image.Rect(0, 14, 20, 26)
var alien2aSprite = image.Rect(20, 14, 40, 26)
var alien3Sprite = image.Rect(0, 27, 20, 40)
var alien3aSprite = image.Rect(20, 27, 40, 40)
var alienExplode = image.Rect(0, 60, 16, 68)
var beamSprite = image.Rect(20, 60, 22, 65)
var bombSprite = image.Rect(0, 70, 10, 79)

// Sprite represents a sprite in the game
type Sprite struct {
	size     image.Rectangle // the sprite size
	Filter   *ebiten.Image   // normal filter used to draw the sprite
	FilterA  *ebiten.Image   // alternate filter used to draw the sprite
	FilterE  *ebiten.Image   // exploded filter used to draw the sprite
	Position image.Point     // top left position of the sprite
	Status   bool            // alive or dead
	Points   int             // number of points if destroyed
}

var aliens = []Sprite{}
var bombs = []Sprite{}

// sprite for laser cannon
var laserCannon Sprite

// sprite for the laser beam
var beam Sprite

// sound effects
var laserSound *audio.Player
var explosionSound *audio.Player
var gameOverSound *audio.Player
var backgroundSound *audio.Player

// audio context
var audioContext *audio.Context

// gameFont
var gameFont font.Face

// used for creating alien sprites
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
func loadFont(path string) font.Face {
	fontBytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	ttfFont, err := truetype.Parse(fontBytes)
	if err != nil {
		log.Fatal(err)
	}
	return truetype.NewFace(ttfFont, &truetype.Options{
		Size: 24, // Adjust the font size as needed
		DPI:  72, // Adjust DPI if needed
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

	return audioStream
}
func initGame() {
	// load sprites image
	imgFile, _, err := ebitenutil.NewImageFromFile("imgs/sprites.png")
	if err != nil {
		panic(err)
	}
	src = imgFile

	// load background image
	bg, _, err := ebitenutil.NewImageFromFile("imgs/bg.png")
	if err != nil {
		panic(err)
	}
	background = bg
	laserCannon = Sprite{
		size:     cannonSprite,
		Filter:   src.SubImage(cannonSprite).(*ebiten.Image),
		FilterE:  src.SubImage(cannonExplode).(*ebiten.Image),
		Position: image.Pt(50, 400),
		Status:   true,
	}

	// sprite for the laser beam
	beam = Sprite{
		size:     beamSprite,
		Filter:   src.SubImage(beamSprite).(*ebiten.Image),
		Position: image.Pt(laserCannon.Position.X+7, 250),
		Status:   false,
	}

	// populate the aliens
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

}

// Game struct
type Game struct {
	loop           int           // game loop
	beamShot       bool          // the instance where the beam is shot
	gameOver       bool          // end of game
	alienDirection int           // direction where alien is heading
	score          int           // number of points scored in the game so far
	startScreen    *ebiten.Image // Start screen image
	gameFont       font.Face
}

func (g *Game) Update() error {
	if g.gameOver {
		if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
			g.resetGame()
		}
		return nil
	}

	// if game is not over, handle input
	if !g.gameOver {
		// move laser cannon left and right
		if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
			laserCannon.Position.X += 10
		}
		if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
			laserCannon.Position.X -= 10
		}

		// fire beam
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			if !beam.Status {
				g.beamShot = true
				if laserSound != nil {
					laserSound.Rewind()
					laserSound.Play()
				}

			}
		}
		// exit the game
		if inpututil.IsKeyJustPressed(ebiten.KeyQ) || inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
			g.gameOver = true
		}

		// if the beam is shot, place the beam at start of the cannon
		if g.beamShot {
			beam.Position.X = laserCannon.Position.X + 7
			beam.Status = true
			g.beamShot = false
		}
	}

	// move the aliens back and forth
	if aliens[0].Position.X < alienSize || aliens[aliensPerRow-1].Position.X > windowWidth-(2*alienSize) {
		g.alienDirection = g.alienDirection * -1
		for i := 0; i < len(aliens); i++ {
			aliens[i].Position.Y = aliens[i].Position.Y + 10
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// if the game is over, draw the background with the game over message
	if g.gameOver {
		// Draw a black background
		screen.Fill(color.Black)

		if gameOverSound != nil {
			gameOverSound.Rewind()
			gameOverSound.Play()
		}
		// Define the message and the "Try Again" button text
		message := fmt.Sprintf("GAME OVER!\nFinal score: %d", g.score)
		tryAgain := "Press ESC to Try Again"

		// Get text bounds to calculate the center position
		messageBounds := text.BoundString(g.gameFont, message)
		tryAgainBounds := text.BoundString(g.gameFont, tryAgain)

		// Calculate the center position of the text
		x := (windowWidth - messageBounds.Dx()) / 2
		y := (windowHeight - messageBounds.Dy()) / 2

		// Calculate the center position of the "Try Again" Button
		xTryAgain := (windowWidth - tryAgainBounds.Dx()) / 2
		yTryAgain := y + messageBounds.Dy() + 20 // move the "Try Again" button 20 px bellow

		// Draw text at the center of the screen
		text.Draw(screen, message, g.gameFont, x, y, color.White)
		text.Draw(screen, tryAgain, g.gameFont, xTryAgain, yTryAgain, color.White)
		return
	}

	// Calculate scale factors
	bgWidth, bgHeight := background.Bounds().Dx(), background.Bounds().Dy()
	xScale := float64(windowWidth) / float64(bgWidth)
	yScale := float64(windowHeight) / float64(bgHeight)

	// Apply scale transform
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(xScale, yScale)
	// draw the background
	screen.DrawImage(background, op)

	// process aliens
	for i := 0; i < len(aliens); i++ {
		aliens[i].Position.X = aliens[i].Position.X + 5*g.alienDirection
		if aliens[i].Status {
			// if alien is hit by a laser beam
			if collide(aliens[i], beam) {
				// draw the explosion
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(aliens[i].Position.X), float64(aliens[i].Position.Y))
				screen.DrawImage(aliens[i].FilterE, op)
				// alien dies, player scores points
				aliens[i].Status = false
				g.score += aliens[i].Points
				// Play explosion sound effect
				if explosionSound != nil {
					explosionSound.Rewind()
					explosionSound.Play()
				}
				// reset the laser beam
				resetBeam()
			} else {
				// show alternating alients
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(aliens[i].Position.X), float64(aliens[i].Position.Y))
				if g.loop%2 == 0 {
					screen.DrawImage(aliens[i].Filter, op)
				} else {
					screen.DrawImage(aliens[i].FilterA, op)
				}
			}

			// drop torpedoes
			if rand.Float64() < bombProbability {
				dropBomb(aliens[i])
			}
		}
	}

	// draw bombs, if laser cannon is hit, game over
	for i := 0; i < len(bombs); i++ {
		bombs[i].Position.Y = bombs[i].Position.Y + bombSpeed
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(bombs[i].Position.X), float64(bombs[i].Position.Y))
		screen.DrawImage(bombs[i].Filter, op)
		if collide(bombs[i], laserCannon) {
			g.gameOver = true
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(laserCannon.Position.X), float64(laserCannon.Position.Y))
			screen.DrawImage(laserCannon.FilterE, op)
		}
	}
	// draw the laser cannon unless it's been destroyed
	if !g.gameOver {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(laserCannon.Position.X), float64(laserCannon.Position.Y))
		screen.DrawImage(laserCannon.Filter, op)
	}

	// keep drawing the beam as it moves every loop
	if beam.Status {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(beam.Position.X), float64(beam.Position.Y))
		screen.DrawImage(beam.Filter, op)
		beam.Position.Y -= 10
	}
	// if the beam leaves the window reset it
	if beam.Position.Y < 0 {
		resetBeam()
	}

	// if the aliens reach the position of the cannon, it's game over!
	if aliens[0].Position.Y > 180 {
		g.gameOver = true
	}
	g.loop++
	// show the score on the top left corner of the screen
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Score: %d", g.score))

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
	g.loop = 0           // game loop
	g.beamShot = false   // the instance where the beam is shot
	g.gameOver = false   // end of game
	g.alienDirection = 1 // direction where alien is heading
	g.score = 0          // number of points scored in the game so far

	aliens = []Sprite{}
	bombs = []Sprite{}

	// populate the aliens
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

	laserCannon.Position = image.Pt(50, 400)
	beam.Position = image.Pt(laserCannon.Position.X+7, 250)
	if backgroundSound != nil {
		backgroundSound.Rewind()
		backgroundSound.Play()
	}
}
func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("Space Invaders")

	// Initialize the audio context
	audioContext = audio.NewContext(48000)

	// Load audio files
	laserSound = loadAudio("files/laser.wav")
	explosionSound = loadAudio("files/explosion.wav")
	gameOverSound = loadAudio("files/game-over.mp3")
	backgroundSound = loadAudio("files/background.wav")

	game := &Game{
		loop:           0,
		beamShot:       false,
		gameOver:       false,
		alienDirection: 1,
		score:          0,
		gameFont:       loadFont("font/font.ttf"),
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
