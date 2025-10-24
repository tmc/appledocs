// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterApplicationLauncher] class.
var (
	MTRClusterApplicationLauncherClass     _MTRClusterApplicationLauncherClass
	MTRClusterApplicationLauncherClassOnce sync.Once
)

func getMTRClusterApplicationLauncherClass() _MTRClusterApplicationLauncherClass {
	MTRClusterApplicationLauncherClassOnce.Do(func() {
		MTRClusterApplicationLauncherClass = _MTRClusterApplicationLauncherClass{objc.GetClass("MTRClusterApplicationLauncher")}
	})
	return MTRClusterApplicationLauncherClass
}

type _MTRClusterApplicationLauncherClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterApplicationLauncher] class.
type IMTRClusterApplicationLauncher interface {
	IMTRGenericCluster
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterApplicationLauncher
type MTRClusterApplicationLauncher struct {
	MTRGenericCluster
}

// MTRClusterApplicationLauncherFrom constructs a [MTRClusterApplicationLauncher] from an unsafe.Pointer.
func MTRClusterApplicationLauncherFrom(ptr unsafe.Pointer) MTRClusterApplicationLauncher {
	return MTRClusterApplicationLauncher{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterApplicationLauncherClass) Alloc() MTRClusterApplicationLauncher {
	rv := objc.Send[MTRClusterApplicationLauncher](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterApplicationLauncherClass) New() MTRClusterApplicationLauncher {
	rv := objc.Send[MTRClusterApplicationLauncher](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterApplicationLauncher) Init() MTRClusterApplicationLauncher {
	rv := objc.Send[MTRClusterApplicationLauncher](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterApplicationLauncher) Autorelease() MTRClusterApplicationLauncher {
	rv := objc.Send[MTRClusterApplicationLauncher](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterApplicationLauncher creates a new MTRClusterApplicationLauncher instance.
func NewMTRClusterApplicationLauncher() MTRClusterApplicationLauncher {
	return getMTRClusterApplicationLauncherClass().New()
}




