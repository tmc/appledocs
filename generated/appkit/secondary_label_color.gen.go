
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [secondaryLabelColor] class.
var secondaryLabelColorClass _secondaryLabelColorClass

func init() {
	secondaryLabelColorClass = _secondaryLabelColorClass{objc.GetClass("secondaryLabelColor")}
}

type _secondaryLabelColorClass struct {
	objc.Class
}

// An interface definition for the [secondaryLabelColor] class.
type IsecondaryLabelColor interface {
	ID() objc.ID
}

type secondaryLabelColor struct {
	id objc.ID
}

func secondaryLabelColorFrom(ptr unsafe.Pointer) secondaryLabelColor {
	return secondaryLabelColor{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ secondaryLabelColor) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _secondaryLabelColorClass) Alloc() secondaryLabelColor {
	rv := objc.Send[secondaryLabelColor](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _secondaryLabelColorClass) New() secondaryLabelColor {
	rv := objc.Send[secondaryLabelColor](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewsecondaryLabelColor creates and returns a new initialized instance.
func NewsecondaryLabelColor() secondaryLabelColor {
	return secondaryLabelColorClass.New()
}

// Init initializes the instance.
func (s_ secondaryLabelColor) Init() secondaryLabelColor {
	rv := objc.Send[secondaryLabelColor](s_.ID(), selInit)
	return rv
}
