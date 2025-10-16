
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [shouldDrawColor] class.
var shouldDrawColorClass _shouldDrawColorClass

func init() {
	shouldDrawColorClass = _shouldDrawColorClass{objc.GetClass("shouldDrawColor")}
}

type _shouldDrawColorClass struct {
	objc.Class
}

// An interface definition for the [shouldDrawColor] class.
type IshouldDrawColor interface {
	ID() objc.ID
}

type shouldDrawColor struct {
	id objc.ID
}

func shouldDrawColorFrom(ptr unsafe.Pointer) shouldDrawColor {
	return shouldDrawColor{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ shouldDrawColor) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _shouldDrawColorClass) Alloc() shouldDrawColor {
	rv := objc.Send[shouldDrawColor](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _shouldDrawColorClass) New() shouldDrawColor {
	rv := objc.Send[shouldDrawColor](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewshouldDrawColor creates and returns a new initialized instance.
func NewshouldDrawColor() shouldDrawColor {
	return shouldDrawColorClass.New()
}

// Init initializes the instance.
func (s_ shouldDrawColor) Init() shouldDrawColor {
	rv := objc.Send[shouldDrawColor](s_.ID(), selInit)
	return rv
}
