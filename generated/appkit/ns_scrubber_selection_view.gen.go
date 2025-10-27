// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [ScrubberSelectionView] class.
var (
	ScrubberSelectionViewClass     _ScrubberSelectionViewClass
	ScrubberSelectionViewClassOnce sync.Once
)

func getScrubberSelectionViewClass() _ScrubberSelectionViewClass {
	ScrubberSelectionViewClassOnce.Do(func() {
		ScrubberSelectionViewClass = _ScrubberSelectionViewClass{objc.GetClass("NSScrubberSelectionView")}
	})
	return ScrubberSelectionViewClass
}

type _ScrubberSelectionViewClass struct {
	class objc.Class
}





// An interface definition for the [ScrubberSelectionView] class.
type IScrubberSelectionView interface {
	IScrubberArrangedView
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (sc _ScrubberSelectionViewClass) Alloc() ScrubberSelectionView {
	rv := objc.Send[ScrubberSelectionView](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _ScrubberSelectionViewClass) New() ScrubberSelectionView {
	rv := objc.Send[ScrubberSelectionView](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ ScrubberSelectionView) Init() ScrubberSelectionView {
	rv := objc.Send[ScrubberSelectionView](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ ScrubberSelectionView) Autorelease() ScrubberSelectionView {
	rv := objc.Send[ScrubberSelectionView](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewScrubberSelectionView creates a new ScrubberSelectionView instance.
func NewScrubberSelectionView() ScrubberSelectionView {
	return getScrubberSelectionViewClass().New()
}





// An abstract base class for specifying the appearance of a highlighted or selected item in a scrubber.
//
// Create a subclass to customize the selection or highlight appearance of an item in your scrubber control. You need to return an instance of your subclass from the method on .


// An abstract base class for specifying the appearance of a highlighted or selected item in a scrubber.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberSelectionView
type ScrubberSelectionView struct {
	ScrubberArrangedView
}

// ScrubberSelectionViewFrom constructs a [ScrubberSelectionView] from an unsafe.Pointer.
//
// An abstract base class for specifying the appearance of a highlighted or selected item in a scrubber.
func ScrubberSelectionViewFrom(ptr unsafe.Pointer) ScrubberSelectionView {
	return ScrubberSelectionView{
		ScrubberArrangedView: ScrubberArrangedViewFrom(ptr),
	}
}































