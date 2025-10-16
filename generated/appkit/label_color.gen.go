
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [labelColor] class.
var labelColorClass _labelColorClass

func init() {
	labelColorClass = _labelColorClass{objc.GetClass("labelColor")}
}

type _labelColorClass struct {
	objc.Class
}

// An interface definition for the [labelColor] class.
type IlabelColor interface {
	ID() objc.ID
}

type labelColor struct {
	id objc.ID
}

func labelColorFrom(ptr unsafe.Pointer) labelColor {
	return labelColor{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (l_ labelColor) ID() objc.ID {
	return l_.id
}

// Alloc allocates a new instance without initialization.
func (lc _labelColorClass) Alloc() labelColor {
	rv := objc.Send[labelColor](objc.ID(lc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (lc _labelColorClass) New() labelColor {
	rv := objc.Send[labelColor](objc.ID(lc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewlabelColor creates and returns a new initialized instance.
func NewlabelColor() labelColor {
	return labelColorClass.New()
}

// Init initializes the instance.
func (l_ labelColor) Init() labelColor {
	rv := objc.Send[labelColor](l_.ID(), selInit)
	return rv
}
