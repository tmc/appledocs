
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ScrubberTextItemView] class.
var ScrubberTextItemViewClass _ScrubberTextItemViewClass

func init() {
	ScrubberTextItemViewClass = _ScrubberTextItemViewClass{objc.GetClass("NSScrubberTextItemView")}
}

type _ScrubberTextItemViewClass struct {
	objc.Class
}

// An interface definition for the [ScrubberTextItemView] class.
type IScrubberTextItemView interface {
	ID() objc.ID
}

type ScrubberTextItemView struct {
	id objc.ID
}

func ScrubberTextItemViewFrom(ptr unsafe.Pointer) ScrubberTextItemView {
	return ScrubberTextItemView{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ ScrubberTextItemView) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _ScrubberTextItemViewClass) Alloc() ScrubberTextItemView {
	rv := objc.Send[ScrubberTextItemView](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _ScrubberTextItemViewClass) New() ScrubberTextItemView {
	rv := objc.Send[ScrubberTextItemView](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewScrubberTextItemView creates and returns a new initialized instance.
func NewScrubberTextItemView() ScrubberTextItemView {
	return ScrubberTextItemViewClass.New()
}

// Init initializes the instance.
func (s_ ScrubberTextItemView) Init() ScrubberTextItemView {
	rv := objc.Send[ScrubberTextItemView](s_.ID(), selInit)
	return rv
}
// The text field that the scrubber item uses to display its text. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubberTextItemView/textField
func (s_ ScrubberTextItemView) TextField() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID(), objc.RegisterName("textField"))
	return rv
}
