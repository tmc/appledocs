// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ScrubberTextItemView] class.
var (
	scrubberTextItemViewClass     _ScrubberTextItemViewClass
	scrubberTextItemViewClassOnce sync.Once
)

func getScrubberTextItemViewClass() _ScrubberTextItemViewClass {
	scrubberTextItemViewClassOnce.Do(func() {
		scrubberTextItemViewClass = _ScrubberTextItemViewClass{objc.GetClass("NSScrubberTextItemView")}
	})
	return scrubberTextItemViewClass
}

type _ScrubberTextItemViewClass struct {
	class objc.Class
}

// An interface definition for the [ScrubberTextItemView] class.
type IScrubberTextItemView interface {
	IScrubberItemView
}

// A concrete view subclass for displaying text for an item in a scrubber. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberTextItemView

type ScrubberTextItemView struct {
	ScrubberItemView
}

// ScrubberTextItemViewFrom constructs a [ScrubberTextItemView] from an unsafe.Pointer.
//
// A concrete view subclass for displaying text for an item in a scrubber.
func ScrubberTextItemViewFrom(ptr unsafe.Pointer) ScrubberTextItemView {
	return ScrubberTextItemView{
		ScrubberItemView: ScrubberItemViewFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (sc _ScrubberTextItemViewClass) Alloc() ScrubberTextItemView {
	rv := objc.Send[ScrubberTextItemView](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (sc _ScrubberTextItemViewClass) New() ScrubberTextItemView {
	rv := objc.Send[ScrubberTextItemView](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ ScrubberTextItemView) Init() ScrubberTextItemView {
	rv := objc.Send[ScrubberTextItemView](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ ScrubberTextItemView) Autorelease() ScrubberTextItemView {
	rv := objc.Send[ScrubberTextItemView](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewScrubberTextItemView creates a new ScrubberTextItemView instance.
func NewScrubberTextItemView() ScrubberTextItemView {
	return getScrubberTextItemViewClass().New()
}




