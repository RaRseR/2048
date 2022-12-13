package twenty48

import (
	"github.com/hajimehoshi/ebiten/v2"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const (
	ScreenWidth  = 360
	ScreenHeight = 360
	boardSize    = 4
)

// Game represents a game state.
type Game struct {
	input      *Input
	board      *Board
	boardImage *ebiten.Image
}

// NewGame generates a new Game object.
func NewGame() (*Game, error) {
	g := &Game{
		input: NewInput(),
	}
	var err error
	g.board, err = NewBoard(boardSize)
	if err != nil {
		return nil, err
	}
	return g, nil
}

// Layout implements ebiten.Game's Layout.
func (game *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

// Update updates the current game state.
func (game *Game) Update() error {
	game.input.Update()
	if err := game.board.Update(game.input); err != nil {
		return err
	}
	return nil
}

// Draw draws the current game to the given screen.
func (game *Game) Draw(screen *ebiten.Image) {
	if game.boardImage == nil {
		w, h := game.board.Size()
		game.boardImage = ebiten.NewImage(w, h)
	}
	screen.Fill(backgroundColor)
	game.board.Draw(game.boardImage)
	op := &ebiten.DrawImageOptions{}
	sw, sh := screen.Size()
	bw, bh := game.boardImage.Size()
	x := (sw - bw) / 2
	y := (sh - bh) / 2
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(game.boardImage, op)
}
