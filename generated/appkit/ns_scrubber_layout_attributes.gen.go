// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ScrubberLayoutAttributes] class.
var (
	ScrubberLayoutAttributesClass     _ScrubberLayoutAttributesClass
	ScrubberLayoutAttributesClassOnce sync.Once
)

func getScrubberLayoutAttributesClass() _ScrubberLayoutAttributesClass {
	ScrubberLayoutAttributesClassOnce.Do(func() {
		ScrubberLayoutAttributesClass = _ScrubberLayoutAttributesClass{objc.GetClass("NSScrubberLayoutAttributes")}
	})
	return ScrubberLayoutAttributesClass
}

type _ScrubberLayoutAttributesClass struct {
	class objc.Class
}

// An interface definition for the [ScrubberLayoutAttributes] class.
type IScrubberLayoutAttributes interface {
	objectivec.IObject
}

// The layout of a scrubber item.
//
// A layout attributes object is the model for the layout of a single item in a scrubber control. If you require model attributes in addition to those provided by this class, create a subclass and add appropriate attributes. Subclasses must implement , and the protocol.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberLayoutAttributes
type ScrubberLayoutAttributes struct {
	objectivec.Object
}

// ScrubberLayoutAttributesFrom constructs a [ScrubberLayoutAttributes] from an unsafe.Pointer.
//
// The layout of a scrubber item.
func ScrubberLayoutAttributesFrom(ptr unsafe.Pointer) ScrubberLayoutAttributes {
	return ScrubberLayoutAttributes{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _ScrubberLayoutAttributesClass) Alloc() ScrubberLayoutAttributes {
	rv := objc.Send[ScrubberLayoutAttributes](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _ScrubberLayoutAttributesClass) New() ScrubberLayoutAttributes {
	rv := objc.Send[ScrubberLayoutAttributes](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ ScrubberLayoutAttributes) Init() ScrubberLayoutAttributes {
	rv := objc.Send[ScrubberLayoutAttributes](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ ScrubberLayoutAttributes) Autorelease() ScrubberLayoutAttributes {
	rv := objc.Send[ScrubberLayoutAttributes](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewScrubberLayoutAttributes creates a new ScrubberLayoutAttributes instance.
func NewScrubberLayoutAttributes() ScrubberLayoutAttributes {
	return getScrubberLayoutAttributesClass().New()
}




// Creates a new layout attributes object for the specified scrubber item index.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberLayoutAttributes/init(forItemAt:)
func NewScrubberLayoutAttributesForItemAtIndex(index int) ScrubberLayoutAttributes {
	rv := objc.Send[ScrubberLayoutAttributes](objc.ID(getScrubberLayoutAttributesClass().class), objc.Sel("layoutAttributesForItemAtIndex:"), index)
	return rv
}


// Creates a new layout attributes object for the specified scrubber item index.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberLayoutAttributes/init(forItemAt:)
func (sc _ScrubberLayoutAttributesClass) LayoutAttributesForItemAtIndex(index int) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("layoutAttributesForItemAtIndex:"), index)
	return rv
}

// The item’s alpha value.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberLayoutAttributes/alpha
func (s_ ScrubberLayoutAttributes) Alpha() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("alpha"))
	return rv
}


// SetAlpha sets the value of the alpha property.
// The item’s alpha value.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberLayoutAttributes/alpha
func (s_ ScrubberLayoutAttributes) SetAlpha(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAlpha:"), value)
}

// The frame of the scrubber item.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberLayoutAttributes/frame
func (s_ ScrubberLayoutAttributes) Frame() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](s_.ID, objc.Sel("frame"))
	return rv
}


// SetFrame sets the value of the frame property.
// The frame of the scrubber item.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberLayoutAttributes/frame
func (s_ ScrubberLayoutAttributes) SetFrame(value coregraphics.CGRect) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setFrame:"), value)
}

// The index of the scrubber item that is represented by the item’s layout attributes.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberLayoutAttributes/itemIndex
func (s_ ScrubberLayoutAttributes) ItemIndex() int {
	rv := objc.Send[int](s_.ID, objc.Sel("itemIndex"))
	return rv
}


// SetItemIndex sets the value of the itemIndex property.
// The index of the scrubber item that is represented by the item’s layout attributes.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberLayoutAttributes/itemIndex
func (s_ ScrubberLayoutAttributes) SetItemIndex(value int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setItemIndex:"), value)
}

// Returns an integer that can be used as a table address in a hash table structure.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/hash
func (s_ ScrubberLayoutAttributes) Hash() int {
	rv := objc.Send[int](s_.ID, objc.Sel("hash"))
	return rv
}


// SetHash sets the value of the hash property.
// Returns an integer that can be used as a table address in a hash table structure.

//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/hash
func (s_ ScrubberLayoutAttributes) SetHash(value int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setHash:"), value)
}


