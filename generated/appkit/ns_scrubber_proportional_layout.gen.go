// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ScrubberProportionalLayout] class.
var (
	ScrubberProportionalLayoutClass     _ScrubberProportionalLayoutClass
	ScrubberProportionalLayoutClassOnce sync.Once
)

func getScrubberProportionalLayoutClass() _ScrubberProportionalLayoutClass {
	ScrubberProportionalLayoutClassOnce.Do(func() {
		ScrubberProportionalLayoutClass = _ScrubberProportionalLayoutClass{objc.GetClass("NSScrubberProportionalLayout")}
	})
	return ScrubberProportionalLayoutClass
}

type _ScrubberProportionalLayoutClass struct {
	class objc.Class
}

// An interface definition for the [ScrubberProportionalLayout] class.
type IScrubberProportionalLayout interface {
	IScrubberLayout
	// properties:
	NumberOfVisibleItems() int
	SetNumberOfVisibleItems(value int)
	// methods:
}

// A concrete layout object that sizes each item to some fraction of the scrubber’s visible size.


// A concrete layout object that sizes each item to some fraction of the scrubber’s visible size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberProportionalLayout
type ScrubberProportionalLayout struct {
	ScrubberLayout
}

// ScrubberProportionalLayoutFrom constructs a [ScrubberProportionalLayout] from an unsafe.Pointer.
//
// A concrete layout object that sizes each item to some fraction of the scrubber’s visible size.
func ScrubberProportionalLayoutFrom(ptr unsafe.Pointer) ScrubberProportionalLayout {
	return ScrubberProportionalLayout{
		ScrubberLayout: ScrubberLayoutFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (sc _ScrubberProportionalLayoutClass) Alloc() ScrubberProportionalLayout {
	rv := objc.Send[ScrubberProportionalLayout](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _ScrubberProportionalLayoutClass) New() ScrubberProportionalLayout {
	rv := objc.Send[ScrubberProportionalLayout](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ ScrubberProportionalLayout) Init() ScrubberProportionalLayout {
	rv := objc.Send[ScrubberProportionalLayout](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ ScrubberProportionalLayout) Autorelease() ScrubberProportionalLayout {
	rv := objc.Send[ScrubberProportionalLayout](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewScrubberProportionalLayout creates a new ScrubberProportionalLayout instance.
func NewScrubberProportionalLayout() ScrubberProportionalLayout {
	return getScrubberProportionalLayoutClass().New()
}



// The number of items visible in the scrubber at once.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberproportionallayout/numberofvisibleitems
func (s_ ScrubberProportionalLayout) NumberOfVisibleItems() int {
	rv := objc.Send[int](s_.ID, objc.Sel("numberOfVisibleItems"))
	return rv
}


// The number of items visible in the scrubber at once.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrubberproportionallayout/numberofvisibleitems
func (s_ ScrubberProportionalLayout) SetNumberOfVisibleItems(value int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setNumberOfVisibleItems:"), value)
}



