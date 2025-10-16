
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Sound] class.
var SoundClass _SoundClass

func init() {
	SoundClass = _SoundClass{objc.GetClass("NSSound")}
}

type _SoundClass struct {
	objc.Class
}

// An interface definition for the [Sound] class.
type ISound interface {
	ID() objc.ID
}

type Sound struct {
	id objc.ID
}

func SoundFrom(ptr unsafe.Pointer) Sound {
	return Sound{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ Sound) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _SoundClass) Alloc() Sound {
	rv := objc.Send[Sound](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _SoundClass) New() Sound {
	rv := objc.Send[Sound](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewSound creates and returns a new initialized instance.
func NewSound() Sound {
	return SoundClass.New()
}

// Init initializes the instance.
func (s_ Sound) Init() Sound {
	rv := objc.Send[Sound](s_.ID(), selInit)
	return rv
}
