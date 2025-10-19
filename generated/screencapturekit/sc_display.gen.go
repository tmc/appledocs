// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SCDisplay] class.
var (
	sCDisplayClass     _SCDisplayClass
	sCDisplayClassOnce sync.Once
)

func getSCDisplayClass() _SCDisplayClass {
	sCDisplayClassOnce.Do(func() {
		sCDisplayClass = _SCDisplayClass{objc.GetClass("SCDisplay")}
	})
	return sCDisplayClass
}

type _SCDisplayClass struct {
	class objc.Class
}

// An interface definition for the [SCDisplay] class.
type ISCDisplay interface {
	objectivec.IObject
}

// An instance that represents a display device.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCDisplay
type SCDisplay struct {
	objectivec.Object
}

// SCDisplayFrom constructs a [SCDisplay] from an unsafe.Pointer.
//
// An instance that represents a display device.
func SCDisplayFrom(ptr unsafe.Pointer) SCDisplay {
	return SCDisplay{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SCDisplayClass) Alloc() SCDisplay {
	rv := objc.Send[SCDisplay](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SCDisplayClass) New() SCDisplay {
	rv := objc.Send[SCDisplay](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SCDisplay) Init() SCDisplay {
	rv := objc.Send[SCDisplay](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SCDisplay) Autorelease() SCDisplay {
	rv := objc.Send[SCDisplay](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSCDisplay creates a new SCDisplay instance.
func NewSCDisplay() SCDisplay {
	return getSCDisplayClass().New()
}




