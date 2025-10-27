// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
)





// The class instance for the [ScrubberFlowLayout] class.
var (
	ScrubberFlowLayoutClass     _ScrubberFlowLayoutClass
	ScrubberFlowLayoutClassOnce sync.Once
)

func getScrubberFlowLayoutClass() _ScrubberFlowLayoutClass {
	ScrubberFlowLayoutClassOnce.Do(func() {
		ScrubberFlowLayoutClass = _ScrubberFlowLayoutClass{objc.GetClass("NSScrubberFlowLayout")}
	})
	return ScrubberFlowLayoutClass
}

type _ScrubberFlowLayoutClass struct {
	class objc.Class
}





// An interface definition for the [ScrubberFlowLayout] class.
type IScrubberFlowLayout interface {
	IScrubberLayout
	

	// properties:
	ItemSize() corefoundation.CGSize
	SetItemSize(value corefoundation.CGSize)
	ItemSpacing() float64
	SetItemSpacing(value float64)


	

	// methods:
	InvalidateLayoutForItemsAtIndexes(invalidItemIndexes foundation.IndexSet)


}





// Alloc allocates a new instance without initialization.
func (sc _ScrubberFlowLayoutClass) Alloc() ScrubberFlowLayout {
	rv := objc.Send[ScrubberFlowLayout](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _ScrubberFlowLayoutClass) New() ScrubberFlowLayout {
	rv := objc.Send[ScrubberFlowLayout](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ ScrubberFlowLayout) Init() ScrubberFlowLayout {
	rv := objc.Send[ScrubberFlowLayout](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ ScrubberFlowLayout) Autorelease() ScrubberFlowLayout {
	rv := objc.Send[ScrubberFlowLayout](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewScrubberFlowLayout creates a new ScrubberFlowLayout instance.
func NewScrubberFlowLayout() ScrubberFlowLayout {
	return getScrubberFlowLayoutClass().New()
}





// A concrete layout object that arranges items end-to-end in a linear strip.
//
// To set the size of items on a per-item basis, ensure that your scrubber delegate conforms to the protocol, and provides an implementation of the method.


// A concrete layout object that arranges items end-to-end in a linear strip.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberFlowLayout
type ScrubberFlowLayout struct {
	ScrubberLayout
}

// ScrubberFlowLayoutFrom constructs a [ScrubberFlowLayout] from an unsafe.Pointer.
//
// A concrete layout object that arranges items end-to-end in a linear strip.
func ScrubberFlowLayoutFrom(ptr unsafe.Pointer) ScrubberFlowLayout {
	return ScrubberFlowLayout{
		ScrubberLayout: ScrubberLayoutFrom(ptr),
	}
}




















// Informs the scrubber that it should perform a new layout pass for the items at the specified indexes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberFlowLayout/invalidateLayoutForItems(at:)
func (s_ ScrubberFlowLayout) InvalidateLayoutForItemsAtIndexes(invalidItemIndexes foundation.IndexSet) {
	objc.Send[objc.ID](s_.ID, objc.Sel("invalidateLayoutForItemsAtIndexes:"), invalidItemIndexes)
}







// The frame size for each item in the scrubber.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberFlowLayout/itemSize
func (s_ ScrubberFlowLayout) ItemSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](s_.ID, objc.Sel("itemSize"))
	return rv
}


// The frame size for each item in the scrubber.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberFlowLayout/itemSize
func (s_ ScrubberFlowLayout) SetItemSize(value corefoundation.CGSize) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setItemSize:"), value)
}


// The horizontal spacing between items, specified in points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberFlowLayout/itemSpacing
func (s_ ScrubberFlowLayout) ItemSpacing() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("itemSpacing"))
	return rv
}


// The horizontal spacing between items, specified in points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberFlowLayout/itemSpacing
func (s_ ScrubberFlowLayout) SetItemSpacing(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setItemSpacing:"), value)
}








