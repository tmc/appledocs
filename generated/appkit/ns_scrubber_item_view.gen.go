// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [ScrubberItemView] class.
var (
	ScrubberItemViewClass     _ScrubberItemViewClass
	ScrubberItemViewClassOnce sync.Once
)

func getScrubberItemViewClass() _ScrubberItemViewClass {
	ScrubberItemViewClassOnce.Do(func() {
		ScrubberItemViewClass = _ScrubberItemViewClass{objc.GetClass("NSScrubberItemView")}
	})
	return ScrubberItemViewClass
}

type _ScrubberItemViewClass struct {
	class objc.Class
}





// An interface definition for the [ScrubberItemView] class.
type IScrubberItemView interface {
	IScrubberArrangedView
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (sc _ScrubberItemViewClass) Alloc() ScrubberItemView {
	rv := objc.Send[ScrubberItemView](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _ScrubberItemViewClass) New() ScrubberItemView {
	rv := objc.Send[ScrubberItemView](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ ScrubberItemView) Init() ScrubberItemView {
	rv := objc.Send[ScrubberItemView](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ ScrubberItemView) Autorelease() ScrubberItemView {
	rv := objc.Send[ScrubberItemView](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewScrubberItemView creates a new ScrubberItemView instance.
func NewScrubberItemView() ScrubberItemView {
	return getScrubberItemViewClass().New()
}





// An item at a specific index position in the scrubber.


// An item at a specific index position in the scrubber.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberItemView
type ScrubberItemView struct {
	ScrubberArrangedView
}

// ScrubberItemViewFrom constructs a [ScrubberItemView] from an unsafe.Pointer.
//
// An item at a specific index position in the scrubber.
func ScrubberItemViewFrom(ptr unsafe.Pointer) ScrubberItemView {
	return ScrubberItemView{
		ScrubberArrangedView: ScrubberArrangedViewFrom(ptr),
	}
}































