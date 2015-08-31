package gomui

import (
	"golang.org/x/mobile/event/touch"
	"golang.org/x/mobile/exp/f32"
	"golang.org/x/mobile/exp/sprite"
	"golang.org/x/mobile/exp/sprite/clock"
)

type Button struct {
	node              *sprite.Node
	state             bState
	enable            bool
	ntex, ttext, dtex sprite.SubTex
}

type bState int

const (
	stateNormal   bState = iota
	stateTouched  bState
	stateDisabled bState
)

func NewButton(w, h int) *Button {
	b := &Button{
		node:  &sprite.Node{},
		state: stateNormal,
		ch:    make(chan touch.Event),
	}

	lastState := b.state
	b.node.Arranger = arrangerFunc(func(e sprite.Engine, n *sprite.Node, t clock.Time) {
		if lastState == b.state {
			return
		}
		lastState = b.state

		var tex sprite.SubTex
		switch b.state {
		case stateNormal:
			tex = b.ntex
		case stateTouched:
			tex = b.ttex
		case stateDisabled:
			tex = b.dtex
		}
		e.SetSubTex(b.node, tex)
		e.SetTransform(b.node, f32.Affine{
			{float64(tex.R.Max.X), 0, 0},
			{0, float64(tex.R.Max.Y), 0},
		})
	})

	return b
}

func (b *Button) SetSubTex(normal, touched, disabled sprite.SubTex) {
	b.ntex = normal
	b.ttex = touched
	b.dtex = disabled
}

func (b *Button) SetEnable(enable bool) {
	if enable {
		b.state = stateNormal
	} else {
		b.state = stateDisabled
	}
}

func (b *Button) contains(x, y float32) bool {
}

func (b *Button) Touch(e touch.Event) {
	switch {
	case b.state == stateDisabled || b.contains():
		return
	case b.state == stateNormal && e.Type == touch.TypeBegin:
		b.state = stateTouched
	case b.state == stateTouched && e.Type == touch.TypeEnd:
		b.state = stateNormal
	}
}

func (b *Button) Event() <-chan touch.Event {
	return b.ch
}

func (b *Button) Node() *sprite.Node {
	return b.node
}
