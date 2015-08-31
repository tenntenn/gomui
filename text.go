package gomui

import (
	"image"
	"image/draw"
	"math"
	"strings"

	sfont "golang.org/x/exp/shiny/font"
	"golang.org/x/image/math/fixed"
	"golang.org/x/mobile/exp/sprite"
)

type TextTexture struct {
	fg, bg  image.Color
	rgba    image.RGBA
	Spacing float64
	Face    sfont.Face
}

func NewTextTexture(bounds image.Rectangle, face sfont.Face) *TextTexture {
	return &TextTexture{
		fg:      image.Black,
		bg:      image.White,
		rgba:    image.NewRGBA(bounds),
		Face:    face,
		Spacing: 1.5,
	}
}

func (t *TextTexture) SetColor(fg, bg image.Color) {
	t.fg, t.bg = fg, bg
}

func (t *TextTexture) SetBounds(bounds image.Rectangle) {
	t.rgba = image.NewRGBA(bounds)
}

func (t *TextTexture) Create(eng sprite.Engine, text string) (sprite.SubTex, error) {
	draw.Draw(t.rgba, t.rgba.Bounds(), t.bg, image.ZP, draw.Src)
	d := &sfont.Drawer{
		Dst:  t.rgba,
		Src:  t.fg,
		Face: t.Face,
	}

	dy := int(math.Ceil(t.Face.Size * t.Spacing))
	for i, s := range strings.Split(text, "\n") {
		d.Dot = fixed.P(0, int(t.Face.Size*0.8)+dy*i)
		d.DrawString(s)
	}

	tex, err := eng.LoadTexture(t.rgba)
	if err != nil {
		return sprite.SubTex{}, err
	}

	return sprite.SubTex{tex, t.rgba.Bounds()}, nil
}
