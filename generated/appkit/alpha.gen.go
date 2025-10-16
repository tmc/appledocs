
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [alpha] class.
var alphaClass _alphaClass

func init() {
	alphaClass = _alphaClass{objc.GetClass("alpha")}
}

type _alphaClass struct {
	objc.Class
}

// An interface definition for the [alpha] class.
type Ialpha interface {
	ID() objc.ID
}

type alpha struct {
	id objc.ID
}

func alphaFrom(ptr unsafe.Pointer) alpha {
	return alpha{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (a_ alpha) ID() objc.ID {
	return a_.id
}

// Alloc allocates a new instance without initialization.
func (ac _alphaClass) Alloc() alpha {
	rv := objc.Send[alpha](objc.ID(ac.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ac _alphaClass) New() alpha {
	rv := objc.Send[alpha](objc.ID(ac.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// Newalpha creates and returns a new initialized instance.
func Newalpha() alpha {
	return alphaClass.New()
}

// Init initializes the instance.
func (a_ alpha) Init() alpha {
	rv := objc.Send[alpha](a_.ID(), selInit)
	return rv
}
