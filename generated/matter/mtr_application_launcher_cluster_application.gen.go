// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRApplicationLauncherClusterApplication] class.
var (
	MTRApplicationLauncherClusterApplicationClass     _MTRApplicationLauncherClusterApplicationClass
	MTRApplicationLauncherClusterApplicationClassOnce sync.Once
)

func getMTRApplicationLauncherClusterApplicationClass() _MTRApplicationLauncherClusterApplicationClass {
	MTRApplicationLauncherClusterApplicationClassOnce.Do(func() {
		MTRApplicationLauncherClusterApplicationClass = _MTRApplicationLauncherClusterApplicationClass{objc.GetClass("MTRApplicationLauncherClusterApplication")}
	})
	return MTRApplicationLauncherClusterApplicationClass
}

type _MTRApplicationLauncherClusterApplicationClass struct {
	class objc.Class
}

// An interface definition for the [MTRApplicationLauncherClusterApplication] class.
type IMTRApplicationLauncherClusterApplication interface {
	IMTRApplicationLauncherClusterApplicationStruct
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRApplicationLauncherClusterApplication
type MTRApplicationLauncherClusterApplication struct {
	MTRApplicationLauncherClusterApplicationStruct
}

// MTRApplicationLauncherClusterApplicationFrom constructs a [MTRApplicationLauncherClusterApplication] from an unsafe.Pointer.
func MTRApplicationLauncherClusterApplicationFrom(ptr unsafe.Pointer) MTRApplicationLauncherClusterApplication {
	return MTRApplicationLauncherClusterApplication{
		MTRApplicationLauncherClusterApplicationStruct: MTRApplicationLauncherClusterApplicationStructFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRApplicationLauncherClusterApplicationClass) Alloc() MTRApplicationLauncherClusterApplication {
	rv := objc.Send[MTRApplicationLauncherClusterApplication](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRApplicationLauncherClusterApplicationClass) New() MTRApplicationLauncherClusterApplication {
	rv := objc.Send[MTRApplicationLauncherClusterApplication](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRApplicationLauncherClusterApplication) Init() MTRApplicationLauncherClusterApplication {
	rv := objc.Send[MTRApplicationLauncherClusterApplication](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRApplicationLauncherClusterApplication) Autorelease() MTRApplicationLauncherClusterApplication {
	rv := objc.Send[MTRApplicationLauncherClusterApplication](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRApplicationLauncherClusterApplication creates a new MTRApplicationLauncherClusterApplication instance.
func NewMTRApplicationLauncherClusterApplication() MTRApplicationLauncherClusterApplication {
	return getMTRApplicationLauncherClusterApplicationClass().New()
}




