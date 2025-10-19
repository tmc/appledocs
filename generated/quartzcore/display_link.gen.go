// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [DisplayLink] class.
var (
	displayLinkClass     _DisplayLinkClass
	displayLinkClassOnce sync.Once
)

func getDisplayLinkClass() _DisplayLinkClass {
	displayLinkClassOnce.Do(func() {
		displayLinkClass = _DisplayLinkClass{objc.GetClass("CADisplayLink")}
	})
	return displayLinkClass
}

type _DisplayLinkClass struct {
	class objc.Class
}

// An interface definition for the [DisplayLink] class.
type IDisplayLink interface {
	objectivec.IObject
	AddToRunLoopForMode(runloop unsafe.Pointer, mode unsafe.Pointer)
}

// A timer object that allows your app to synchronize its drawing to the refresh rate of the display. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CADisplayLink
type DisplayLink struct {
	objectivec.Object
}

// DisplayLinkFrom constructs a [DisplayLink] from an unsafe.Pointer.
//
// A timer object that allows your app to synchronize its drawing to the refresh rate of the display.
func DisplayLinkFrom(ptr unsafe.Pointer) DisplayLink {
	return DisplayLink{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (dc _DisplayLinkClass) Alloc() DisplayLink {
	rv := objc.Send[DisplayLink](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DisplayLinkClass) New() DisplayLink {
	rv := objc.Send[DisplayLink](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DisplayLink) Init() DisplayLink {
	rv := objc.Send[DisplayLink](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DisplayLink) Autorelease() DisplayLink {
	rv := objc.Send[DisplayLink](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDisplayLink creates a new DisplayLink instance.
func NewDisplayLink() DisplayLink {
	return getDisplayLinkClass().New()
}


// Registers the display link with a run loop. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CADisplayLink/add(to:forMode:)
func (d_ DisplayLink) AddToRunLoopForMode(runloop unsafe.Pointer, mode unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("addToRunLoop:forMode:"), runloop, mode)
}


