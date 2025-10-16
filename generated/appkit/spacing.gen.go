
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [spacing] class.
var spacingClass _spacingClass

func init() {
	spacingClass = _spacingClass{objc.GetClass("spacing")}
}

type _spacingClass struct {
	objc.Class
}

// An interface definition for the [spacing] class.
type Ispacing interface {
	ID() objc.ID
}

type spacing struct {
	id objc.ID
}

func spacingFrom(ptr unsafe.Pointer) spacing {
	return spacing{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ spacing) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _spacingClass) Alloc() spacing {
	rv := objc.Send[spacing](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _spacingClass) New() spacing {
	rv := objc.Send[spacing](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// Newspacing creates and returns a new initialized instance.
func Newspacing() spacing {
	return spacingClass.New()
}

// Init initializes the instance.
func (s_ spacing) Init() spacing {
	rv := objc.Send[spacing](s_.ID(), selInit)
	return rv
}
