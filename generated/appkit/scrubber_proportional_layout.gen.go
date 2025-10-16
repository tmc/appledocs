
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ScrubberProportionalLayout] class.
var ScrubberProportionalLayoutClass _ScrubberProportionalLayoutClass

func init() {
	ScrubberProportionalLayoutClass = _ScrubberProportionalLayoutClass{objc.GetClass("NSScrubberProportionalLayout")}
}

type _ScrubberProportionalLayoutClass struct {
	objc.Class
}

// An interface definition for the [ScrubberProportionalLayout] class.
type IScrubberProportionalLayout interface {
	ID() objc.ID
}

type ScrubberProportionalLayout struct {
	id objc.ID
}

func ScrubberProportionalLayoutFrom(ptr unsafe.Pointer) ScrubberProportionalLayout {
	return ScrubberProportionalLayout{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ ScrubberProportionalLayout) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _ScrubberProportionalLayoutClass) Alloc() ScrubberProportionalLayout {
	rv := objc.Send[ScrubberProportionalLayout](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _ScrubberProportionalLayoutClass) New() ScrubberProportionalLayout {
	rv := objc.Send[ScrubberProportionalLayout](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewScrubberProportionalLayout creates and returns a new initialized instance.
func NewScrubberProportionalLayout() ScrubberProportionalLayout {
	return ScrubberProportionalLayoutClass.New()
}

// Init initializes the instance.
func (s_ ScrubberProportionalLayout) Init() ScrubberProportionalLayout {
	rv := objc.Send[ScrubberProportionalLayout](s_.ID(), selInit)
	return rv
}
