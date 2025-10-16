
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ScrubberArrangedView] class.
var ScrubberArrangedViewClass _ScrubberArrangedViewClass

func init() {
	ScrubberArrangedViewClass = _ScrubberArrangedViewClass{objc.GetClass("NSScrubberArrangedView")}
}

type _ScrubberArrangedViewClass struct {
	objc.Class
}

// An interface definition for the [ScrubberArrangedView] class.
type IScrubberArrangedView interface {
	ID() objc.ID
	ApplyLayoutAttributes(layoutAttributes unsafe.Pointer)
}

type ScrubberArrangedView struct {
	id objc.ID
}

func ScrubberArrangedViewFrom(ptr unsafe.Pointer) ScrubberArrangedView {
	return ScrubberArrangedView{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ ScrubberArrangedView) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _ScrubberArrangedViewClass) Alloc() ScrubberArrangedView {
	rv := objc.Send[ScrubberArrangedView](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _ScrubberArrangedViewClass) New() ScrubberArrangedView {
	rv := objc.Send[ScrubberArrangedView](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewScrubberArrangedView creates and returns a new initialized instance.
func NewScrubberArrangedView() ScrubberArrangedView {
	return ScrubberArrangedViewClass.New()
}

// Init initializes the instance.
func (s_ ScrubberArrangedView) Init() ScrubberArrangedView {
	rv := objc.Send[ScrubberArrangedView](s_.ID(), selInit)
	return rv
}
// Updates the layout of the arranged view to respect the provided layout attributes. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubberArrangedView/apply(_:)
func (s_ ScrubberArrangedView) ApplyLayoutAttributes(layoutAttributes unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("applyLayoutAttributes:"), layoutAttributes)
}
// A Boolean value that specifies whether the view is currently highlighted. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubberArrangedView/isHighlighted
func (s_ ScrubberArrangedView) Highlighted() bool {
	rv := objc.Send[bool](s_.ID(), objc.RegisterName("highlighted"))
	return rv
}
// SetHighlighted sets the value of the highlighted property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubberArrangedView/isHighlighted
func (s_ ScrubberArrangedView) SetHighlighted(value bool) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setHighlighted:"), value)
}
// A Boolean value that specifies whether the current view is selected. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubberArrangedView/isSelected
func (s_ ScrubberArrangedView) Selected() bool {
	rv := objc.Send[bool](s_.ID(), objc.RegisterName("selected"))
	return rv
}
// SetSelected sets the value of the selected property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubberArrangedView/isSelected
func (s_ ScrubberArrangedView) SetSelected(value bool) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setSelected:"), value)
}
