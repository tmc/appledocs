
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [gestureRecognizers] class.
var gestureRecognizersClass _gestureRecognizersClass

func init() {
	gestureRecognizersClass = _gestureRecognizersClass{objc.GetClass("gestureRecognizers")}
}

type _gestureRecognizersClass struct {
	objc.Class
}

// An interface definition for the [gestureRecognizers] class.
type IgestureRecognizers interface {
	ID() objc.ID
}

type gestureRecognizers struct {
	id objc.ID
}

func gestureRecognizersFrom(ptr unsafe.Pointer) gestureRecognizers {
	return gestureRecognizers{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (g_ gestureRecognizers) ID() objc.ID {
	return g_.id
}

// Alloc allocates a new instance without initialization.
func (gc _gestureRecognizersClass) Alloc() gestureRecognizers {
	rv := objc.Send[gestureRecognizers](objc.ID(gc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (gc _gestureRecognizersClass) New() gestureRecognizers {
	rv := objc.Send[gestureRecognizers](objc.ID(gc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewgestureRecognizers creates and returns a new initialized instance.
func NewgestureRecognizers() gestureRecognizers {
	return gestureRecognizersClass.New()
}

// Init initializes the instance.
func (g_ gestureRecognizers) Init() gestureRecognizers {
	rv := objc.Send[gestureRecognizers](g_.ID(), selInit)
	return rv
}
