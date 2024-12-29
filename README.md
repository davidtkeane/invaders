# <h1 align="center">👾 Space Invaders 🚀</h1>


<h2 align="center"><img src="https://camo.githubusercontent.com/33fc802b38252c211851cc922f01bad895d203ece43756c966c98c49ebc21594/68747470733a2f2f636f756e742e6765746c6f6c692e636f6d2f6765742f4066756a697761726163686f6b693f7468656d653d61736f756c"></h2>

## Table of Contents

-   [Features](#features)
-   [Personal Note](#personal-note)
-   [About](#about)
-   [Screenshots](#screenshots)
-   [Installation](#installation)
    -   [Method 1: Using `install_go.sh` (Recommended)](#method-1-using-install_gosh-recommended)
    -   [Method 2: Manual Installation](#method-2-manual-installation)
-   [Running the Game](#running-the-game)
-   [Folder Structure](#folder-structure)
-   [Gameplay](#gameplay)
-   [Configuration](#configuration)
-   [Acknowledgments](#acknowledgments)
-   [Contributing](#contributing)
-   [License](#license)
-   [Space-Invaders](#space-invaders)

## Features

-   Classic Space Invaders gameplay.
-   Multiple alien types with different behaviors.
-   Player-controlled laser cannon with shooting.
-   Destructible barriers.
-   Sound effects for laser fire, explosions, and game over.
-   Background music during gameplay.
-   High score tracking (top 5 scores).
-   Game over screen with final score and high scores display.
-   Option to restart the game or exit.

## Personal Note

<img src="https://camo.githubusercontent.com/b199fabd947beb75a0113cae47b1dd8d1c8be22cd4802cd24b1c8863a4533e3f/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f6d6164655f776974682d2545322539442541342d7265643f7374796c653d666f722d7468652d6261646765266c6162656c436f6c6f723d6f72616e6765" alt="Made with ❤️">

I wanted to play a game of Space Invaders and I started to code, then I remembered GitHub and I did a quick search and found the repo [Space Invaders Game](https://github.com/sausheong/invaders) by [sausheong](https://github.com/sausheong). It is written in Go language and I loved it. So, I asked an AI Chatbot to help add sounds, a background image, to add more lives, and to save the game's high scores. I tried to make it feel more like an arcade machine with the end game sounds and music, with the flashing scoreboard. Two install scripts were created, one to download a GitHub repo and install everything for you, the 2nd is an install Go script that will install Go if it is not installed, then the script will install the dependencies needed for the game.

A big thank you to [sausheong](https://github.com/sausheong) for giving me a platform to learn and play with.

Please find [sausheong](https://github.com/sausheong)'s instructions and how [sausheong](https://github.com/sausheong) made the game below the updated `README.md`. <br>
I have added ` <---> ` to separate both `README.md` files.

Thank you and take care.

## About

<h1 align="left">Hello There! and Welcome to My Version of Invaders <img src="https://i.imgur.com/8tT5L21.gif" width="28px" alt="🤙"></h1>

This project recreates the classic arcade game Space Invaders using Go and the Ebiten game library. It's a simplified version, focusing on core gameplay elements like alien movement, laser fire, and collision detection. The project demonstrates 2D game development principles in Go, showcasing sprite animation, game loops, and basic input handling. No external game engine is used; the game is built directly using Ebiten's drawing and input functions for a minimalist and educational approach. It leverages Go's concurrency features (goroutines) for handling user input and game logic simultaneously. This project also serves as a modern homage to the golden age of arcade games, offering a fun and nostalgic coding exercise for those new to Go and game development.

## Screenshots

<h2 align="left">🤙 Space Invaders Main Screen</h2>

<a href="https://i.imgur.com/x3SLXY5.png"><img src="https://i.imgur.com/x3SLXY5.png" alt="Space Invaders Screenshot" width="50%"></a>

<h2 align="left">🤙 Space Invaders End-Game Screen</h2>

<a href="https://i.imgur.com/aDKAPKS.png"><img src="https://i.imgur.com/aDKAPKS.png" alt="End Game Screen" width="50%"  align="center" ></a>

## Installation

![GitHub last commit](https://img.shields.io/github/last-commit/davidtkeane/invaders?style=flat-square)
![GitHub issues](https://img.shields.io/github/issues-raw/davidtkeane/invaders?style=flat-square)
![Go Version](https://img.shields.io/github/go-mod/go-version/davidtkeane/invaders)
![GitHub commit status](https://img.shields.io/github/checks-status/davidtkeane/invaders/c4fabaedaf92e7821514903768518b03ff2e297a)
![GitHub Sponsors](https://img.shields.io/github/sponsors/davidtkeane)

### Method 1: Using `install_go.sh` (Recommended)

1.  **Download:** Clone the repository to download the `install_go.sh` script and the game files to your computer.

    ```bash
    git clone [https://github.com/davidtkeane/invaders.git](https://github.com/davidtkeane/invaders.git)
    ```

    **Enter Folder:** Enter the files folder and then move to number 2.

    ```bash
    cd invaders/files
    ```

2.  **Make Executable:** Open a terminal and navigate to the directory where you downloaded the files. Run the following command to make the script executable:

    ```bash
    chmod +x install_go.sh
    ```

3.  **Run the Script:** Execute the script with the following command:

    ```bash
    ./install_go.sh
    ```

    **The script after installing will then ask if you want to play invaders:**

    If you say no and want to run the script from the command line use this command:

    ```bash
    go run main.go
    ```
    <br>
    The script will perform the following actions:

    -   Check if Go is already installed on your system.
    -   If Go is not installed:
        -   Attempt to automatically download and install Go version 1.21.5 (you can modify this in the script if needed).
        -   Verify the SHA256 checksum of the downloaded Go archive for security.
        -   Extract the Go archive to `/usr/local`.
        -   Temporarily add Go to your `PATH` for the current terminal session.
        -   Print instructions on how to make the `PATH` change permanent (so you can use Go in any terminal session).
        -   Test the Go installation by printing the version.
    -   Install the necessary Go packages for the game using `go mod tidy`, `go mod download`, and `go get`.
    -   Prompt you whether you want to run the game immediately.

### Method 2: Manual Installation

1.  **Install Go:**

    -   If you don't have Go installed, go to the official Go website: [https://go.dev/doc/install](https://go.dev/doc/install)
    -   Download the Go installer for your operating system (macOS, Windows, or Linux).
    -   Run the installer and follow the on-screen instructions.
    -   **Important:** Make sure that the Go `bin` directory (e.g., `/usr/local/go/bin` or `C:\Go\bin`) is added to your system's `PATH` environment variable. This allows you to run Go commands from any terminal location. You might need to restart your terminal or even your computer for the `PATH` changes to take effect.
    -   **Verify Installation:** Open a terminal and run `go version`. You should see the installed Go version printed.

2.  **Download Game Files:**

    -   Download the `main.go` file, the `install_go.sh` script and all the asset files (images, audio files, font file) from the game's repository.
    -   Organize the files into the correct directory structure as shown in the "Folder Structure" section below.

3.  **Install Dependencies:**

    -   Open a terminal and navigate to the directory where you saved the game files (the directory containing `main.go`).
    -   Run the following commands in order:

        ```bash
        go mod init example.com/spaceinvaders [invalid URL removed]
        go mod tidy
        go mod download golang.org/x/image
        go get github.com/golang/freetype/truetype [invalid URL removed]
        go get github.com/hajimehoshi/ebiten/v2 [invalid URL removed]
        go get github.com/hajimehoshi/ebiten/v2/audio [invalid URL removed]
        go get github.com/hajimehoshi/ebiten/v2/audio/wav [invalid URL removed]
        go get github.com/hajimehoshi/ebiten/v2/audio/mp3 [invalid URL removed]
        go get github.com/hajimehoshi/ebiten/v2/ebitenutil [invalid URL removed]
        go get github.com/hajimehoshi/ebiten/v2/inpututil [invalid URL removed]
        go get github.com/hajimehoshi/ebiten/v2/text [invalid URL removed]
        go get golang.org/x/image/font
        ```

    -   **Note:** The `go get` commands will download the specified packages and their dependencies to your Go workspace (`$GOPATH/pkg/mod`). These packages are not directly placed in your project folder. The `go.mod` and `go.sum` files keep track of the exact versions used in your project.

## Running the Game

Once you have installed the game using either method, you can run it with the following command in the terminal:

```bash
go run main.go
````

### Folder Structure

Ensure that all asset files are in their respective folders as follows:

```bash
$ tree
.
├── README.md
├── files
│   ├── background.wav      # 🔊 Background music
│   ├── end-game.mp3        # 🔊 Game over music
│   ├── game-over2.wav		# 🔊 Game over music
│   ├── game-over3.mp3		# 🔊 Game over music
│   ├── explosion.wav       # 🔊 Explosion sound effect
│   ├── explosion-sound.mp3 # 🔊 Explosion sound effect
│   ├── game-over.mp3       # 🔊 Game over sound effect
│   ├── girlfriend.txt      # 📄 Text file (Easter egg message)
│   ├── highscores.txt      # 💾 High scores data
│   ├── install_go.sh       # 💻 Installation script (Bash)
│   ├── install_github.sh   # 💻 Installation script for Github repo's (Bash)
│   └── laser.wav           # 🔊 Laser sound effect
├── font
│   └── font.ttf            # 🔤 Font file for text
├── go.mod                  # 📄 Go module file
├── go.sum                  # 📄 Go module checksum file
├── images                  # 📂 Original Images
│   ├── collision.png       # 💥 Image of collision in the game
│   ├── invaders.gif        # 👾 Animated GIF of the Invaders game
│   ├── space-invaders.png  # 👾 Space Invaders game title image
│   ├── sprites-position.png# 👾 Positioned sprites image
│   └── sprites.png         # 👾 Collection of game sprites
├── imgs                    # 📂 Resized Images
│   ├── background-end3.jpg # 🖼️ Game over background image
│   ├── bg.png              # 🌌 Background image
│   ├── sprites.png         # 👾 Spritesheet image
│   └── start.png           # ▶️ Start screen image
└── main.go                 # 👾 Main Go source file

5 directories, 25 files
```

**File Explanations:**

  - **`main.go`:** The main Go source code file for the game. It contains the game logic, rendering functions, and initialization code.
  - **`install_go.sh`:** A Bash script that automates the installation of Go, the required packages, and optionally runs the game.
  - **`imgs/`:**
      - `background-end3.jpg`: The background image used on the game over screen.
      - `bg.png`: The background image used during gameplay.
      - `sprites.png`: A spritesheet containing images of the aliens, cannon, laser beam, bombs, and barriers.
  - **`files/`:**
      - `.wav`, `.mp3`: Audio files for various sound effects (laser, explosion, game over) and background music.
      - `highscores.txt`: Stores the high score data.
      - `girlfriend.txt`: A text file containing a message printed by `install_go.sh`.
  - **`font/`:**
      - `font.ttf`: The font file used to render text in the game.
  - **`go.mod`:** The Go module file, which lists the project's dependencies (Ebiten, image libraries, etc.).
  - **`go.sum`:** Contains checksums of the downloaded dependencies for security and verification.

## Gameplay 🎮

  - **Move Cannon:** Use the left and right arrow keys ⬅️➡️ to move the laser cannon.
  - **Fire:** Press the Spacebar 🚀 to fire the laser beam.
  - **Pause:** Press the Esc key ⏸️ to pause/unpause the game.
  - **Quit:** Press Q ❌ to quit the game.
  - **Game Over:** The game ends when the aliens reach the bottom of the screen ⬇️ or when the player loses all lives 💔.

**Configuration**
You can adjust various game settings in the `main.go` file. The `HELP SECTION` at the top of the file provides detailed instructions on how to configure these settings.

**Acknowledgments**

  - **Original Game:** This game is based on the classic Space Invaders arcade game.
  - **Ebiten Library:** Developed and maintained by [Hajime Hoshi](https://www.google.com/url?sa=E&source=gmail&q=https://github.com/hajimehoshi).
  - **Original Fork:** This project is a fork of [github.com/sausheong/invaders](https://github.com/sausheong/invaders).
  - **Developers:**
      - [RangerSmyth](https://www.google.com/url?sa=E&source=gmail&q=https://github.com/RangerSmyth)
      - [David T Keane](https://github.com/davidtkeane)

## Contributing

Contributions to this project are welcome\! If you find any bugs or want to add new features, please feel free to submit a pull request or open an issue on the GitHub repository.

## License

This project is licensed under the MIT License - see the [LICENSE](https://www.google.com/url?sa=E&source=gmail&q=LICENSE) file for details. (You'll need to add a LICENSE file to your repository if you want to specify the MIT License or any other license).

## Space-Invaders

## How I Updated the Space Invaders Game

<details>
<summary><b>Click to Expand</b></summary>

This project started as a fork of [sausheong's Space Invaders game](https://github.com/sausheong/invaders), which was itself inspired by an Ebiten example. The original game provided a solid foundation, but I wanted to enhance it and make it more like a complete arcade experience while also using this project as a learning opportunity. Here's a summary of the updates and improvements I made, with the help of an AI chatbot:

**1. Enhanced Game Over Screen:**

-   **Background Image:** Added a dedicated background image (`background-end3.jpg`) for a more visually appealing game over screen. *Now includes 3 images to choose from!*
-   **Information Box:** Implemented a semi-transparent box to frame the game over information, improving readability and organization.
-   **Text Formatting and Positioning:**
    -   Adjusted the font size for the "GAME OVER" message.
    -   Carefully calculated and positioned the "GAME OVER", final score, "Play Again", "Close Game", and high scores list elements within the box.
    -   Ensured proper spacing between text elements.
-   **High Scores:**
    -   Implemented a persistent high score system that loads and saves the top 5 scores to `files/highscores.txt`.
    -   Prevented duplicate entries with the same name and score.
    -   Displayed the high scores list on the game over screen.

**2. Gameplay Improvements:**

-   **Barriers:** Added destructible barriers to the game screen, providing cover for the player.
-   **Audio:** Integrated multiple sound effects, including laser fire, alien and ship explosions, and game over sounds, to enhance the arcade experience. Also included the functionality to loop background music during gameplay. *Multiple audio files have been included to choose from*
-   **Lives:** Added extra lives to the player for extended gameplay.

**3. Code Enhancements:**

-   **Refactoring:**
    -   Split the `Draw()` function into `drawGameOverScreen()` and `drawGameScreen()` to improve code organization and readability.
-   **Error Handling:** Added better error handling, especially when loading external resources like images and audio files.
-   **Comments:** Included detailed comments throughout the code to explain the logic and functionality, making it easier to understand and modify.
-   **Configuration:** Added a `HELP SECTION` at the beginning of the code with clear instructions on how to configure various game settings, audio, controls, and more.

**4. Installation and Documentation:**

-   **`install_go.sh`:** Created a Bash script to automate the installation of Go, necessary packages, and dependencies. The script also offers an option to run the game after installation.
-   **`README.md`:** Wrote a comprehensive `README.md` file that includes:
    -   A detailed description of the game and its features.
    -   Instructions for both automatic (using `install_go.sh`) and manual installation.
    -   A breakdown of the project's folder structure and file purposes.
    -   Gameplay instructions.
    -   Configuration guidelines.
    -   Acknowledgments to the original developers and libraries used.
    -   Contributing guidelines.

**5. Collaboration and Learning:**

-   This project was a collaborative effort between myself (RangerSmyth) and another developer (David T Keane).
-   We also heavily utilized an AI chatbot as a tool to help us with debugging, refactoring, and generating ideas for improvements.
-   The process involved a lot of back-and-forth, testing, and iterative refinement to get to the final version.
</details>

## Customization

### Changing Background Images and Audio

You can easily customize the background images and audio files used in the game:

<details>
<summary><b>Click to Expand</b></summary>

#### Background Images

1.  **Game Background:**
    -   Replace the `imgs/bg.png` file with your desired background image.
    -   Make sure the new image is in PNG format and has the appropriate dimensions (800x600 pixels are recommended to fit the game window).

2.  **Game Over Background:**
    -   The game over screen uses the image in `imgs/background-end3.jpg`
    -   Replace this file with your preferred image, ensuring it's in JPG format. I have provided 3 different backgrounds to choose from, to change the background used, you will need to edit this line in the `initGame()` function:

    ```go
    bgEnd, _, err := ebitenutil.NewImageFromFile("imgs/background-end3.jpg") // Change the file name here
    ```

#### Audio Files

1.  **Sound Effects:**
    -   The sound effects are located in the `files/` directory.
    -   You can replace the existing `.wav` or `.mp3` files (e.g., `laser.wav`, `explosion.wav`, `game-over.mp3`) with your own sound effects.

2.  **Background Music:**
    -   The background music during gameplay is `files/background.wav`. Replace this file with your desired background music.
    -   The game over screen music is `files/end-game.mp3`. Replace this with your desired game over music.

**Important:**

-   Make sure the replacement images and audio files have the correct file extensions (e.g., `.png`, `.jpg`, `.wav`, `.mp3`).
-   The game is designed to work with specific audio formats (WAV and MP3). Using other formats might require code modifications.

### Modifying Game Settings

You can adjust various game parameters in the `main.go` file. Here's how to modify some of the key settings:

<details>
<summary><b>Click to Expand</b></summary>

1.  **Window Size:**

    -   Modify `windowWidth` and `windowHeight` variables at the beginning of the `main.go` file.

    ```go
    var (
        windowWidth      = 800
        windowHeight     = 600
    )
    ```

2.  **Alien Parameters:**

    -   `aliensPerRow`: Number of aliens in each row.
    -   `aliensStartCol`: Starting horizontal position of the aliens.
    -   `alienSize`: Size of each alien sprite.

    ```go
    var (
        aliensPerRow     = 8
        aliensStartCol   = 100
        alienSize        = 30
    )
    ```

3.  **Bomb Parameters:**

    -   `bombProbability`: Probability of an alien dropping a bomb in each frame (0.0 to 1.0).
    -   `bombSpeed`: Speed at which bombs fall.

    ```go
    var (
        bombProbability  = 0.005
        bombSpeed        = 10
    )
    ```

4.  **Barrier and Player Position:**

    -   `barrierYPosition`: Vertical position of the barriers.
    -   `playerYPosition`: Initial vertical position of the player's cannon.

    ```go
    var (
        barrierYPosition = 300
        playerYPosition  = 400
    )
    ```

5.  **Game Over Screen Text Offsets:**

    -   These variables control the vertical spacing of the text elements on the game over screen:

    ```go
    var (
        gameOverMessageYOffset = 100 
        finalScoreYOffset      = 150 
        playAgainYOffset       = 100 
        closeGameYOffset       = 100 
        highScoresTitleYOffset = 160 
        highScoresListYOffset  = 30  
        highScoresListSpacing  = 20  
    )
    ```

6.  **Other Settings:**

    -   You can find many other settings and parameters throughout the code (e.g., player lives, alien movement speed, laser beam speed). Look for comments that explain what each variable does.

**Important:**

-   Be careful when modifying the code. Make sure you understand what each variable does before changing it.
-   It's recommended to create a backup of your `main.go` file before making significant changes.
-   Test your changes thoroughly after modifying any settings.

</details>
</details>

## Future Improvements

-   Adding a settings panel that can be opened/closed with a button
-   Adding animation to the game
-   Adding the ability to have 2 players at the same time
-   Adding a pause menu with options
-   Adding volume controls for the sound effects
-   Implementing more complex alien movement patterns.
-   Adding more levels and increasing difficulty.

<h4 align="center"> This project has been a great learning experience, and I hope you enjoy playing this updated version of Space Invaders! 🚀 </h4>



				DDDD     AA   V     V EEEEE
				D   D   A  A  V     V E
				D    D AAAAAA V     V EEEEE
				D   D  A    A  V   V  E
				DDDD   A    A   VVV   EEEEE

\<---\>

<h1 align="center">👾 Space Invaders by Sausheong 🚀 </h1>

![Space Invaders](images/space-invaders.png)


The Start of Space Invaders by [sausheong](https://github.com/sausheong/invaders)

# Writing Space Invaders with Go

The earliest memory I had of arcade video games was watching my older brother and cousins going at the video game machines at Genting Highlands. While our parents were at the other types of games Genting Highlands was more popularly known for, we were generally let loose to play arcade games to our hearts' content. 

Those were the magical days of Pac-Man, Space Invaders, Galaxian, Donkey Kong, Frogger, Centipede and many, many more. Days of blinking lights, intense music, frantic tugs at the joystick, furious mashing of the buttons then the groans of dismay as the last life was lost. 

As with many aspiring programmers starting out, one of my secret dreams was always to recreate that magic, to write the next big game. And as with many as well programmers, I failed many times miserably. Even though eventually I succeeded in writing some simple games, I came to realise that even seemingly simple games are in fact not easy to write.

Of course that didn't stop me from trying my hand at it once again. This time round I tried the grand old dame of arcade video games -- _Space Invaders_.

## Space Invaders

[Space Invaders](https://en.wikipedia.org/wiki/Space_Invaders) was one of the most successful arcade video games during the [golden age of arcade video games](https://en.wikipedia.org/wiki/Golden_age_of_arcade_video_games). It was first released in 1978 and is generally accepted as start of the golden age, which lasted from late 70s to the early 90s. Even after the decline of arcade video games, it simply transcended the medium and moved to video game consoles.

![Space Invaders](images/space-invaders.png)

The premise of the game, in case you're not already familiar, is quite simple. As the player you control a laser cannon to battle rows of alien invaders. The laser cannon can only move horizontally across the bottom of the screen as the aliens move back and forth, and slowly descending upon you. The aliens try to destroy you by dropping torpedos at you, while you are partially protected by a few stationary defending bunkers. The game ends if the aliens reach the cannon or all your cannons are destroyed.

As I said, simple. 

## Engine-less

Most game development uses some sort of game engine or at least a graphics engine, but the game I wrote uses neither. Instead what I tried to do is to build a game by creating individual frames and displaying them rapidly one after another. Essentially, this is a variation of the [flocking simulation](https://sausheong.github.io/posts/flocking-with-go/) I wrote earlier. 

The idea is quite straightforward -- a while back I stumbled on a [hack](https://www.iterm2.com/documentation-images.html) in iTerm2 that allows me to display images on the screen and that led me, one thing to another, displaying a number of images one after another, resulting in an animation.

Which, if you think about it, is what a simple game like Space Invaders is all about -- an animation that can be controlled by a user.

Let's look at the code.

## Sprites

In computer graphics, sprites are independent objects that are added on to a background. Sprites were, not unexpectedly, first used in arcade video games and were often generated by hardware. In our case, we're using a simple but popular technique of using a single sprite sheet and taking various parts of the sprite sheet to be a separate sprite.

![sprites](/images/sprites.png)

The image above is a magnified version of the `sprites.png` sprite sheet file. 

```go
// sprites
var src = getImage("imgs/sprites.png")
var background = getImage("imgs/bg.png")
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
```
Each sprite is represented by an `image.Rectangle` position of the corresponding sprite image in the `sprites.png` file. For example, the `alien1Sprite` shows a Rectangle with the upper left position of `(0,0)` and the lower right position of `(20,14)`.

![position](/images/sprites-position.png)

We'll see how this is being used in a short while.

We also see that we load up the 2 image files `sprites.png` and `bg.png`. This function simply gets an `image.Image` from an image file.

```go
func getImage(filePath string) image.Image {
	imgFile, err := os.Open(filePath)
	defer imgFile.Close()
	if err != nil {
		fmt.Println("Cannot read file:", err)
	}
	img, _, err := image.Decode(imgFile)
	if err != nil {
		fmt.Println("Cannot decode file:", err)
	}
	return img
}
```

Now that we have the positions of the sprites, let's see the `Sprite` struct itself.

```go
// Sprite represents a sprite in the game
type Sprite struct {
	size     image.Rectangle // the sprite size
	Filter   *gift.GIFT      // normal filter used to draw the sprite
	FilterA  *gift.GIFT      // alternate filter used to draw the sprite
	FilterE  *gift.GIFT      // exploded filter used to draw the sprite
	Position image.Point     // top left position of the sprite
	Status   bool            // alive or dead
	Points   int             // number of points if destroyed
}
```

The sprites are represented by:

1. The size of the sprite, which is the Rectangle we defined earlier
2. 3 images filters which we will later use for drawing the sprites
3. The position of the sprite to draw on the background
4. A representation of the status of the sprite, and 
5. The number of points if the sprite is destroyed (this is only applicable for aliens).

The image filters we are using are from the excellent [Go Image Filtering Toolkit (GIFT)](https://github.com/disintegration/gift). We are underutilising this library since we're only using this to draw the sprite on the background. Each sprite has 3 filters, a normal filter for drawing a normal sprite, an alternate filter, which draws an alternate form of the sprite (only for aliens, so far) and an exploded filter, which draws the sprite when it's exploded (or died). We're using the alternate form of the sprite to animate the sprite.

Let's look at the definition of different sprites in the game.

```go
var aliens = []Sprite{}

// sprite for laser cannon
var laserCannon = Sprite{
	size:     cannonSprite,
	Filter:   gift.New(gift.Crop(cannonSprite)),
	Position: image.Pt(50, 250),
	Status:   true,
}

// sprite for the laser beam
var beam = Sprite{
	size:     beamSprite,
	Filter:   gift.New(gift.Crop(beamSprite)),
	Position: image.Pt(laserCannon.Position.X+7, 250),
	Status:   false,
}
```

This is a simplified version of Space Invaders, so we only have 4 types of sprites -- aliens, the bombs the aliens drop, the laser cannon, and the laser beam that shoots out of the cannon. The variable `aliens` is an array of alien sprites, `bombs` is an array of bombs dropped while `laserCannon` is the laser cannon sprite and `beam` is the laser beam sprite. As you can see from the code, the filter crops the part of the sprite sheet according to the defined Rectangle earlier.

We're not creating the aliens yet here, but we'll need to in a while, which we will be using a function:

```go
// used for creating alien sprites
func createAlien(x, y int, sprite, alt image.Rectangle, points int) (s Sprite) {
	s = Sprite{
		size:     sprite,
		Filter:   gift.New(gift.Crop(sprite)),
		FilterA:  gift.New(gift.Crop(alt)),
		FilterE:  gift.New(gift.Crop(alienExplode)),
		Position: image.Pt(x, y),
		Status:   true,
		Points:   points,
	}
	return
}
```

We wil be passing the Rectangles to the function to create the correct alien sprite and it's alternate form (for animating the sprite) but all aliens explode the same way. Bombs are created in the game loop itself and we'll see it later.

So much for the sprites, let's look at the `main` function of the game.

## The action starts

This is a game that starts from the terminal, and the whole game is animated on the terminal. Therefore, controlling the terminal is very important. I used the popular [termbox-go](https://github.com/nsf/termbox-go) library to give me this control. Termbox, as with the GIFT library, is probably an overkill since it's a much more powerful library than is needed here.

```go
err := termbox.Init()
if err != nil {
	panic(err)
}
```

I start off the main function by initializing termbox-go. The game actually has 2 independent loops:

* The first is the control of the laser cannon by the user and this is through input on the keyboard (left and right arrows)
* The second is the rest of the game, in what is called a game loop, explained later. This runs regardless of whatever the user does. It includes continual movement of the aliens as they descend while dropping bombs to destroy all life on earth as well as the upward movement of the laser beam as it hurtles towards the aliens to blast them to bits.

This means there are 2 concurrently running threads, and in this case, goroutines.

```go
// poll for keyboard events in another goroutine
events := make(chan termbox.Event, 1000)
go func() {
	for {
		events <- termbox.PollEvent()
	}
}()
```

We start a separate goroutine that polls for events and stuffs them into a buffered channel. We're using a liberally large buffer here, it can really be smaller, but the larger the buffer, the relatively smoother the event captures are so I didn't really try to figure out the optimal size.

Most games have a start screen, where you are asked to place credits to press a button to begin. We have one too.

```go
	// show the start screen
	startScreen := getImage("imgs/start.png")
	printImage(startScreen)
start:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Ch == 's' || ev.Ch == 'S' {
				break start
			}
			if ev.Ch == 'q' {
				gameOver = true
				break start
			}
		}
	}
```

We're polling the keyboard again, and we're waiting for some to either press the `s` or `S` to begin the game, or `q` to quit the game.

Next, we populate the `aliens` array. This is pretty straightforward, we simply want 3 rows of different aliens.

```go
// populate the aliens
for i := aliensStartCol; i < aliensStartCol+(alienSize*aliensPerRow); i += alienSize {
	aliens = append(aliens, createAlien(i, 30, alien1Sprite, alien1aSprite, 30))
}
for i := aliensStartCol; i < aliensStartCol+(30*aliensPerRow); i += alienSize {
	aliens = append(aliens, createAlien(i, 55, alien2Sprite, alien2aSprite, 20))
}
for i := aliensStartCol; i < aliensStartCol+(30*aliensPerRow); i += alienSize {
	aliens = append(aliens, createAlien(i, 80, alien3Sprite, alien3aSprite, 10))
}
```

## Game loop

Now that we've laid the groundwork, let's get into the main game loop. Most games run in a what is commonly called a [game loop](http://gameprogrammingpatterns.com/game-loop.html). A game loop is a game software development pattern that is often the heart of a game. It's an infinite loop that updates and redraws to animate game play. In our Space Invaders game loop we use a variable named `gameOver` to indicate that the loop should continue until the game ends (either triggered by the player or when the aliens win).

The game loop is rather long so we'll break it up into a few parts. Let's look at the first part, which is used to capture the keyboard events from the player.


```go
// if any of the keyboard events are captured
select {
case ev := <-events:
	if ev.Type == termbox.EventKey {
		if ev.Key == termbox.KeyCtrlQ {
			gameOver = true
		}
		if ev.Key == termbox.KeySpace {
			if beam.Status == false {
				beamShot = true
			}
		}
		if ev.Key == termbox.KeyArrowRight {
			laserCannon.Position.X += 10
		}
		if ev.Key == termbox.KeyArrowLeft {
			laserCannon.Position.X -= 10
		}
	}

default:

}
```

Whenever the `events` buffered channel has something, it will be received into the `ev` variable where we try to determine what kind of event it is. Pressing `Ctrl-Q` ends the game, while pressing the space bar fires a laser beam. Pressing the left or right arrow buttons move the laser cannon left or right accordingly. 

You might notice that we have to have a default in the select. This is because if we don't have the default, the select will block, and the loop can only proceed whenever the user presses a key!

Next, we'll draw the laser cannon and alien sprites. We start off with an empty image and drawing the background on to it.

```go
// create background
dst := image.NewRGBA(image.Rect(0, 0, windowWidth, windowHeight))
gift.New().Draw(dst, background)
```

## Aliens are coming!

We start off with the aliens first.

```go
// process aliens
for i := 0; i < len(aliens); i++ {
	aliens[i].Position.X = aliens[i].Position.X + 5*alienDirection
	if aliens[i].Status {
		// if alien is hit by a laser beam
		if collide(aliens[i], beam) {
			// draw the explosion
			aliens[i].FilterE.DrawAt(dst, src, image.Pt(aliens[i].Position.X, aliens[i].Position.Y), gift.OverOperator)
			// alien dies, player scores points
			aliens[i].Status = false
			score += aliens[i].Points
			// reset the laser beam
			resetBeam()
		} else {
			// show alternating alients
			if loop%2 == 0 {
				aliens[i].Filter.DrawAt(dst, src, image.Pt(aliens[i].Position.X, aliens[i].Position.Y), gift.OverOperator)
			} else {
				aliens[i].FilterA.DrawAt(dst, src, image.Pt(aliens[i].Position.X, aliens[i].Position.Y), gift.OverOperator)
			}
			// drop torpedoes
			if rand.Float64() < bombProbability {
				dropBomb(aliens[i])
			}			
		}
	}
}
```

To determine where the aliens should be going, we multiply the horizontal (X) position of the alien with the variable `alienDirection`. 

We also use the `Status` of the alien to determine if it's alive or dead. If it's alive, we check if it has collided with a laser beam. If yes, it's dead. We draw the explosion sprite, set the `Status` to false, rack up the player points and reset the laser beam. Resetting the laser beam just means we set the beam's `Status` back to `false` and place it at the vertical (Y) level the same as the cannon.

```go
func resetBeam() {
	beam.Status = false
	beam.Position.Y = 250
}
```

## Collision physics

If the alien didn't collide with the laser beam, we display either the normal sprite or the alternate sprite. This gives us the animation of a moving alien.

Let's take a quick look at the collision physics, which is concentrated in the `collide()` function.

```go
func collide(s1, s2 Sprite) bool {
	spriteA := image.Rect(s1.Position.X, s1.Position.Y, s1.Position.X+s1.size.Dx(), s1.Position.Y+s1.size.Dy())
	spriteB := image.Rect(s2.Position.X, s2.Position.Y, s2.Position.X+s1.size.Dx(), s2.Position.Y+s1.size.Dy())
	if spriteA.Min.X < spriteB.Max.X && spriteA.Max.X > spriteB.Min.X &&
		spriteA.Min.Y < spriteB.Max.Y && spriteA.Max.Y > spriteB.Min.Y {
		return true
	}
	return false
}
```

Give the two sprites are boxed within the two rectangles, the sprites are considered to have collided if all these conditions are met:

*  spriteA.Min.X < spriteB.Max.X
*  spriteA.Max.X > spriteB.Min.X
*  spriteA.Min.Y < spriteB.Max.Y
*  spriteA.Max.Y > spriteB.Min.Y

![collision](images/collision.png)

## Bombs away

The aliens drop bombs as they descend on the laser cannon. Obviously we don't want it to drop bombs continually, so we use a probability to determine if the bomb should be dropped or not.

```go
// drop torpedoes
if rand.Float64() < bombProbability {
	dropBomb(aliens[i])
}	
```

Dropping a bomb here means we create a new `Bomb` sprite and set it to start at where the alien is located, then adding it to the array of bombs.

```go
func dropBomb(alien Sprite) {
	torpedo := Sprite{
		size:     bombSprite,
		Filter:   gift.New(gift.Crop(bombSprite)),
		Position: image.Pt(alien.Position.X+7, alien.Position.Y),
		Status:   true,
	}

	bombs = append(bombs, torpedo)
}
```

Now that we have drawn the aliens (or its death by explosion), we check if it has moved outside of the window. It it has, we reverse the direction and move the aliens down.

```go
// move the aliens back and forth
if aliens[0].Position.X < alienSize || aliens[aliensPerRow-1].Position.X > windowWidth-(2*alienSize) {
	alienDirection = alienDirection * -1
	for i := 0; i < len(aliens); i++ {
		aliens[i].Position.Y = aliens[i].Position.Y + 10
	}
}
```

We also need to move the bombs downwards on its deadly descend, according to the `bombSpeed`.

```go
// draw bombs, if laser cannon is hit, game over
for i := 0; i < len(bombs); i++ {
	bombs[i].Position.Y = bombs[i].Position.Y + bombSpeed
	bombs[i].Filter.DrawAt(dst, src, image.Pt(bombs[i].Position.X, bombs[i].Position.Y), gift.OverOperator)
	if collide(bombs[i], laserCannon) {
		gameOver = true
		laserCannon.FilterE.DrawAt(dst, src, image.Pt(laserCannon.Position.X, laserCannon.Position.Y), gift.OverOperator)
	}
}
```
As the bombs drop, we need to check if it collides with the laser cannon. If it did, it's game over, and we draw the exploding cannon sprite.

That's it for the aliens and their bombs in the game loop. Let's look at the laser cannon and its laser beam next. 

## Laser cannon and beams

The laser cannon is relatively simple. We continue drawing the cannon as long as it's not been destroyed.

```go
// draw the laser cannon unless it's been destroyed
if !gameOver {
	laserCannon.Filter.DrawAt(dst, src, image.Pt(laserCannon.Position.X, laserCannon.Position.Y), gift.OverOperator)
}
```

As for the laser beam, since there is only one variable to represent it you'd probably realise that the cannon can only shoot one laser beam at a time. This is mostly for simplicity's sake.

We use the `beamShot` variable to determine if the player has pressed the space bar, and therefore shot the laser cannon. 

```go
// if the beam is shot, place the beam at start of the cannon
if beamShot {
	beam.Position.X = laserCannon.Position.X + 7
	beam.Status = true
	beamShot = false
}
```

If yes, we place the beam at the tip of the laser cannon, set the beam's `Status` to `true` to indicate that the beam is in action and we should show it, then set `beamShot` back to false.

We need to check the beam's `Status` and if it's in action, we draw it and then move it up. If the beam's position is outside the window, we reset the beam again.

```go
// keep drawing the beam as it moves every loop
if beam.Status {
	beam.Filter.DrawAt(dst, src, image.Pt(beam.Position.X, beam.Position.Y), gift.OverOperator)
	beam.Position.Y -= 10
}

// if the beam leaves the window reset it
if beam.Position.Y < 0 {
	resetBeam()
}
```

# Wrapping up and printing the image

Just before we end the game loop, we print the image to the screen, print the score and increment the loop.

```go
printImage(dst)
fmt.Println("\n\nSCORE:", score)
loop++
```
So how _exactly_ do we print the image on the screen? This is done through a simple hack on ITerm2 (so this only works on iTerm2, sorry!)

```go
// this only works for iTerm2!
func printImage(img image.Image) {
	var buf bytes.Buffer
	png.Encode(&buf, img)
	imgBase64Str := base64.StdEncoding.EncodeToString(buf.Bytes())
	fmt.Printf("\x1b[2;0H\x1b]1337;File=inline=1:%s\a", imgBase64Str)
}
```

And that's it! All in all, including comments, the code is less than 300 lines. If it seems pretty short, you should remember the original entire compiled binary ROM for the game is less than 110kb in size, including images and music! And mine is a simpler clone without the full game features (multiple lives, defender bunkers, ufo etc) and also includes 2 external libraries.

This is how it looks on my computer. How does it look on yours?

![invaders](images/invaders.gif)

# Code

The code for this is all here in this GitHub repository.

https://github.com/sausheong/invaders

