// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ScrubberImageItemView] class.
var (
	scrubberImageItemViewClass     _ScrubberImageItemViewClass
	scrubberImageItemViewClassOnce sync.Once
)

func getScrubberImageItemViewClass() _ScrubberImageItemViewClass {
	scrubberImageItemViewClassOnce.Do(func() {
		scrubberImageItemViewClass = _ScrubberImageItemViewClass{objc.GetClass("NSScrubberImageItemView")}
	})
	return scrubberImageItemViewClass
}

type _ScrubberImageItemViewClass struct {
	class objc.Class
}

// An interface definition for the [ScrubberImageItemView] class.
type IScrubberImageItemView interface {
	IScrubberItemView
}

// A concrete view subclass for displaying images in a scrubber items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubberImageItemView
type ScrubberImageItemView struct {
	ScrubberItemView
}

// ScrubberImageItemViewFrom constructs a [ScrubberImageItemView] from an unsafe.Pointer.
//
// A concrete view subclass for displaying images in a scrubber items.
func ScrubberImageItemViewFrom(ptr unsafe.Pointer) ScrubberImageItemView {
	return ScrubberImageItemView{
		ScrubberItemView: ScrubberItemViewFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (sc _ScrubberImageItemViewClass) Alloc() ScrubberImageItemView {
	rv := objc.Send[ScrubberImageItemView](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _ScrubberImageItemViewClass) New() ScrubberImageItemView {
	rv := objc.Send[ScrubberImageItemView](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ ScrubberImageItemView) Init() ScrubberImageItemView {
	rv := objc.Send[ScrubberImageItemView](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ ScrubberImageItemView) Autorelease() ScrubberImageItemView {
	rv := objc.Send[ScrubberImageItemView](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewScrubberImageItemView creates a new ScrubberImageItemView instance.
func NewScrubberImageItemView() ScrubberImageItemView {
	return getScrubberImageItemViewClass().New()
}




