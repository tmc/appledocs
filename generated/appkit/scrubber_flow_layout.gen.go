
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ScrubberFlowLayout] class.
var ScrubberFlowLayoutClass _ScrubberFlowLayoutClass

func init() {
	ScrubberFlowLayoutClass = _ScrubberFlowLayoutClass{objc.GetClass("NSScrubberFlowLayout")}
}

type _ScrubberFlowLayoutClass struct {
	objc.Class
}

// An interface definition for the [ScrubberFlowLayout] class.
type IScrubberFlowLayout interface {
	ID() objc.ID
}

type ScrubberFlowLayout struct {
	id objc.ID
}

func ScrubberFlowLayoutFrom(ptr unsafe.Pointer) ScrubberFlowLayout {
	return ScrubberFlowLayout{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ ScrubberFlowLayout) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _ScrubberFlowLayoutClass) Alloc() ScrubberFlowLayout {
	rv := objc.Send[ScrubberFlowLayout](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _ScrubberFlowLayoutClass) New() ScrubberFlowLayout {
	rv := objc.Send[ScrubberFlowLayout](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewScrubberFlowLayout creates and returns a new initialized instance.
func NewScrubberFlowLayout() ScrubberFlowLayout {
	return ScrubberFlowLayoutClass.New()
}

// Init initializes the instance.
func (s_ ScrubberFlowLayout) Init() ScrubberFlowLayout {
	rv := objc.Send[ScrubberFlowLayout](s_.ID(), selInit)
	return rv
}
