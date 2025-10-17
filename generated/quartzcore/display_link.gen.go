// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [DisplayLink] class.
var displayLinkClass = _DisplayLinkClass{objc.GetClass("CADisplayLink")}

type _DisplayLinkClass struct {
	class objc.Class
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

// Registers the display link with a run loop. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CADisplayLink/add(to:forMode:)
func (d_ DisplayLink) AddToRunLoopForMode(runloop unsafe.Pointer, mode unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("addToRunLoop:forMode:"), runloop, mode)
}


