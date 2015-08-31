package gomui

import (
	"image"

	sfont "golang.org/x/exp/shiny/font"
	"golang.org/x/mobile/exp/f32"
	"golang.org/x/mobile/exp/sprite"
	"golang.org/x/mobile/exp/sprite/clock"
)

type Label struct {
	*TextTexture
	node *sprite.Node
	Text string
}

func NewLabel(w, h int, face sfont.Face) *Label {
	l := &Label{
		TextTexture: NewTextTexture(image.Rect(0, 0, w, h), face),
		node:        &sprite.Node{},
	}

	var lastText string
	l.node.Arranger = arrangerFunc(func(e sprite.Engine, n *sprite.Node, t clock.Time) {
		if l.Text == lastText {
			return
		}
		lastText = l.Text
		if l.Text == "" {
			e.SetSubTex(l.node, sprite.SubTex{})
		} else {
			if tex, err := l.Create(e, l.Text); err == nil {
				e.SetSubTex(l.node, tex)
				e.SetTransform(l.node, f32.Affine{
					{float64(w), 0, 0},
					{0, float64(h), 0},
				})
			} else {
				panic(err)
			}
		}
	})

	return l
}

func (l *Label) Node() *sprite.Node {
	return l.node
}
