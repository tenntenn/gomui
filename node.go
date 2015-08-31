package gomui

import (
	"image"

	"golang.org/x/mobile/event/touch"
	"golang.org/x/mobile/exp/sprite"
	"golang.org/x/mobile/exp/sprite/clock"
)

type Node interface {
	Node() *sprite.Node
	Bounds() image.Rectangle
}

func AbsTransform(node Node) {
	panic("TODO")
	/*
		parents := []Node{node}

		n := node
		for n.Node().Parent != nil {
			parents = append(parents, n.Parent)
			n = n.Parent
		}

		for i := len(parents); i >= 0; i-- {
		}
	*/
}

type arrangerFunc func(e sprite.Engine, n *sprite.Node, t clock.Time)

func (a arrangerFunc) Arrange(e sprite.Engine, n *sprite.Node, t clock.Time) { a(e, n, t) }

type Touchable interface {
	Touch(e touch.Event)
	Event() <-chan touch.Event
}
