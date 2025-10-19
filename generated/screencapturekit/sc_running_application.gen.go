// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SCRunningApplication] class.
var (
	sCRunningApplicationClass     _SCRunningApplicationClass
	sCRunningApplicationClassOnce sync.Once
)

func getSCRunningApplicationClass() _SCRunningApplicationClass {
	sCRunningApplicationClassOnce.Do(func() {
		sCRunningApplicationClass = _SCRunningApplicationClass{objc.GetClass("SCRunningApplication")}
	})
	return sCRunningApplicationClass
}

type _SCRunningApplicationClass struct {
	class objc.Class
}

// An interface definition for the [SCRunningApplication] class.
type ISCRunningApplication interface {
	objectivec.IObject
}

// An instance that represents an app running on a device.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCRunningApplication
type SCRunningApplication struct {
	objectivec.Object
}

// SCRunningApplicationFrom constructs a [SCRunningApplication] from an unsafe.Pointer.
//
// An instance that represents an app running on a device.
func SCRunningApplicationFrom(ptr unsafe.Pointer) SCRunningApplication {
	return SCRunningApplication{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SCRunningApplicationClass) Alloc() SCRunningApplication {
	rv := objc.Send[SCRunningApplication](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SCRunningApplicationClass) New() SCRunningApplication {
	rv := objc.Send[SCRunningApplication](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SCRunningApplication) Init() SCRunningApplication {
	rv := objc.Send[SCRunningApplication](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SCRunningApplication) Autorelease() SCRunningApplication {
	rv := objc.Send[SCRunningApplication](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSCRunningApplication creates a new SCRunningApplication instance.
func NewSCRunningApplication() SCRunningApplication {
	return getSCRunningApplicationClass().New()
}




