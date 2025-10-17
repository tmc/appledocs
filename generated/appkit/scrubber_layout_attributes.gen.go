
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ScrubberLayoutAttributes] class.
var ScrubberLayoutAttributesClass _ScrubberLayoutAttributesClass

func init() {
	ScrubberLayoutAttributesClass = _ScrubberLayoutAttributesClass{objc.GetClass("NSScrubberLayoutAttributes")}
}

type _ScrubberLayoutAttributesClass struct {
	objc.Class
}

// An interface definition for the [ScrubberLayoutAttributes] class.
type IScrubberLayoutAttributes interface {
	ID() objc.ID
}

type ScrubberLayoutAttributes struct {
	id objc.ID
}

func ScrubberLayoutAttributesFrom(ptr unsafe.Pointer) ScrubberLayoutAttributes {
	return ScrubberLayoutAttributes{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ ScrubberLayoutAttributes) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _ScrubberLayoutAttributesClass) Alloc() ScrubberLayoutAttributes {
	rv := objc.Send[ScrubberLayoutAttributes](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _ScrubberLayoutAttributesClass) New() ScrubberLayoutAttributes {
	rv := objc.Send[ScrubberLayoutAttributes](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewScrubberLayoutAttributes creates and returns a new initialized instance.
func NewScrubberLayoutAttributes() ScrubberLayoutAttributes {
	return ScrubberLayoutAttributesClass.New()
}

// Init initializes the instance.
func (s_ ScrubberLayoutAttributes) Init() ScrubberLayoutAttributes {
	rv := objc.Send[ScrubberLayoutAttributes](s_.ID(), selInit)
	return rv
}
// The item’s alpha value. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubberLayoutAttributes/alpha
func (s_ ScrubberLayoutAttributes) Alpha() float64 {
	rv := objc.Send[float64](s_.ID(), objc.RegisterName("alpha"))
	return rv
}
// SetAlpha sets the value of the alpha property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubberLayoutAttributes/alpha
func (s_ ScrubberLayoutAttributes) SetAlpha(value float64) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setAlpha:"), value)
}
// The frame of the scrubber item. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubberLayoutAttributes/frame
func (s_ ScrubberLayoutAttributes) Frame() foundation.Rect {
	rv := objc.Send[foundation.Rect](s_.ID(), objc.RegisterName("frame"))
	return rv
}
// SetFrame sets the value of the frame property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubberLayoutAttributes/frame
func (s_ ScrubberLayoutAttributes) SetFrame(value foundation.Rect) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setFrame:"), value)
}
// The index of the scrubber item that is represented by the item’s layout attributes. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubberLayoutAttributes/itemIndex
func (s_ ScrubberLayoutAttributes) ItemIndex() int {
	rv := objc.Send[int](s_.ID(), objc.RegisterName("itemIndex"))
	return rv
}
// SetItemIndex sets the value of the itemIndex property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubberLayoutAttributes/itemIndex
func (s_ ScrubberLayoutAttributes) SetItemIndex(value int) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setItemIndex:"), value)
}
