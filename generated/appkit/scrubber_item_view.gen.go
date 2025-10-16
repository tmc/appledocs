
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ScrubberItemView] class.
var ScrubberItemViewClass _ScrubberItemViewClass

func init() {
	ScrubberItemViewClass = _ScrubberItemViewClass{objc.GetClass("NSScrubberItemView")}
}

type _ScrubberItemViewClass struct {
	objc.Class
}

// An interface definition for the [ScrubberItemView] class.
type IScrubberItemView interface {
	ID() objc.ID
}

type ScrubberItemView struct {
	id objc.ID
}

func ScrubberItemViewFrom(ptr unsafe.Pointer) ScrubberItemView {
	return ScrubberItemView{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ ScrubberItemView) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _ScrubberItemViewClass) Alloc() ScrubberItemView {
	rv := objc.Send[ScrubberItemView](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _ScrubberItemViewClass) New() ScrubberItemView {
	rv := objc.Send[ScrubberItemView](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewScrubberItemView creates and returns a new initialized instance.
func NewScrubberItemView() ScrubberItemView {
	return ScrubberItemViewClass.New()
}

// Init initializes the instance.
func (s_ ScrubberItemView) Init() ScrubberItemView {
	rv := objc.Send[ScrubberItemView](s_.ID(), selInit)
	return rv
}
