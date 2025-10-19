// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ScrubberFlowLayout] class.
var (
	scrubberFlowLayoutClass     _ScrubberFlowLayoutClass
	scrubberFlowLayoutClassOnce sync.Once
)

func getScrubberFlowLayoutClass() _ScrubberFlowLayoutClass {
	scrubberFlowLayoutClassOnce.Do(func() {
		scrubberFlowLayoutClass = _ScrubberFlowLayoutClass{objc.GetClass("NSScrubberFlowLayout")}
	})
	return scrubberFlowLayoutClass
}

type _ScrubberFlowLayoutClass struct {
	class objc.Class
}

// An interface definition for the [ScrubberFlowLayout] class.
type IScrubberFlowLayout interface {
	IScrubberLayout
}

// A concrete layout object that arranges items end-to-end in a linear strip.
//
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

// Alloc allocates a new instance without initialization.
func (sc _ScrubberFlowLayoutClass) Alloc() ScrubberFlowLayout {
	rv := objc.Send[ScrubberFlowLayout](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




