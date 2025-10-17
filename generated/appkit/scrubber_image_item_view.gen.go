
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ScrubberImageItemView] class.
var ScrubberImageItemViewClass _ScrubberImageItemViewClass

func init() {
	ScrubberImageItemViewClass = _ScrubberImageItemViewClass{objc.GetClass("NSScrubberImageItemView")}
}

type _ScrubberImageItemViewClass struct {
	objc.Class
}

// An interface definition for the [ScrubberImageItemView] class.
type IScrubberImageItemView interface {
	ID() objc.ID
}

type ScrubberImageItemView struct {
	id objc.ID
}

func ScrubberImageItemViewFrom(ptr unsafe.Pointer) ScrubberImageItemView {
	return ScrubberImageItemView{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ ScrubberImageItemView) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _ScrubberImageItemViewClass) Alloc() ScrubberImageItemView {
	rv := objc.Send[ScrubberImageItemView](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _ScrubberImageItemViewClass) New() ScrubberImageItemView {
	rv := objc.Send[ScrubberImageItemView](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewScrubberImageItemView creates and returns a new initialized instance.
func NewScrubberImageItemView() ScrubberImageItemView {
	return ScrubberImageItemViewClass.New()
}

// Init initializes the instance.
func (s_ ScrubberImageItemView) Init() ScrubberImageItemView {
	rv := objc.Send[ScrubberImageItemView](s_.ID(), selInit)
	return rv
}
// The alignment of the image within the scrubber item. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubberImageItemView/imageAlignment
func (s_ ScrubberImageItemView) ImageAlignment() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID(), objc.RegisterName("imageAlignment"))
	return rv
}
// SetImageAlignment sets the value of the imageAlignment property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubberImageItemView/imageAlignment
func (s_ ScrubberImageItemView) SetImageAlignment(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setImageAlignment:"), value)
}
// The image view that the scrubber item uses to display its image. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubberImageItemView/imageView
func (s_ ScrubberImageItemView) ImageView() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID(), objc.RegisterName("imageView"))
	return rv
}
