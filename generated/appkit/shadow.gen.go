
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Shadow] class.
var ShadowClass _ShadowClass

func init() {
	ShadowClass = _ShadowClass{objc.GetClass("NSShadow")}
}

type _ShadowClass struct {
	objc.Class
}

// An interface definition for the [Shadow] class.
type IShadow interface {
	ID() objc.ID
}

type Shadow struct {
	id objc.ID
}

func ShadowFrom(ptr unsafe.Pointer) Shadow {
	return Shadow{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ Shadow) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _ShadowClass) Alloc() Shadow {
	rv := objc.Send[Shadow](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _ShadowClass) New() Shadow {
	rv := objc.Send[Shadow](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewShadow creates and returns a new initialized instance.
func NewShadow() Shadow {
	return ShadowClass.New()
}

// Init initializes the instance.
func (s_ Shadow) Init() Shadow {
	rv := objc.Send[Shadow](s_.ID(), selInit)
	return rv
}
