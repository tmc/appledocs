// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SCWindow] class.
var (
	sCWindowClass     _SCWindowClass
	sCWindowClassOnce sync.Once
)

func getSCWindowClass() _SCWindowClass {
	sCWindowClassOnce.Do(func() {
		sCWindowClass = _SCWindowClass{objc.GetClass("SCWindow")}
	})
	return sCWindowClass
}

type _SCWindowClass struct {
	class objc.Class
}

// An interface definition for the [SCWindow] class.
type ISCWindow interface {
	objectivec.IObject
}

// An instance that represents an onscreen window.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCWindow
type SCWindow struct {
	objectivec.Object
}

// SCWindowFrom constructs a [SCWindow] from an unsafe.Pointer.
//
// An instance that represents an onscreen window.
func SCWindowFrom(ptr unsafe.Pointer) SCWindow {
	return SCWindow{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SCWindowClass) Alloc() SCWindow {
	rv := objc.Send[SCWindow](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SCWindowClass) New() SCWindow {
	rv := objc.Send[SCWindow](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SCWindow) Init() SCWindow {
	rv := objc.Send[SCWindow](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SCWindow) Autorelease() SCWindow {
	rv := objc.Send[SCWindow](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSCWindow creates a new SCWindow instance.
func NewSCWindow() SCWindow {
	return getSCWindowClass().New()
}




