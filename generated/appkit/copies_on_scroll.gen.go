
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [copiesOnScroll] class.
var copiesOnScrollClass _copiesOnScrollClass

func init() {
	copiesOnScrollClass = _copiesOnScrollClass{objc.GetClass("copiesOnScroll")}
}

type _copiesOnScrollClass struct {
	objc.Class
}

// An interface definition for the [copiesOnScroll] class.
type IcopiesOnScroll interface {
	ID() objc.ID
}

type copiesOnScroll struct {
	id objc.ID
}

func copiesOnScrollFrom(ptr unsafe.Pointer) copiesOnScroll {
	return copiesOnScroll{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ copiesOnScroll) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _copiesOnScrollClass) Alloc() copiesOnScroll {
	rv := objc.Send[copiesOnScroll](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _copiesOnScrollClass) New() copiesOnScroll {
	rv := objc.Send[copiesOnScroll](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewcopiesOnScroll creates and returns a new initialized instance.
func NewcopiesOnScroll() copiesOnScroll {
	return copiesOnScrollClass.New()
}

// Init initializes the instance.
func (c_ copiesOnScroll) Init() copiesOnScroll {
	rv := objc.Send[copiesOnScroll](c_.ID(), selInit)
	return rv
}
