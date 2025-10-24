// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterApplicationLauncher] class.
var (
	MTRBaseClusterApplicationLauncherClass     _MTRBaseClusterApplicationLauncherClass
	MTRBaseClusterApplicationLauncherClassOnce sync.Once
)

func getMTRBaseClusterApplicationLauncherClass() _MTRBaseClusterApplicationLauncherClass {
	MTRBaseClusterApplicationLauncherClassOnce.Do(func() {
		MTRBaseClusterApplicationLauncherClass = _MTRBaseClusterApplicationLauncherClass{objc.GetClass("MTRBaseClusterApplicationLauncher")}
	})
	return MTRBaseClusterApplicationLauncherClass
}

type _MTRBaseClusterApplicationLauncherClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterApplicationLauncher] class.
type IMTRBaseClusterApplicationLauncher interface {
	IMTRGenericBaseCluster
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterApplicationLauncher
type MTRBaseClusterApplicationLauncher struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterApplicationLauncherFrom constructs a [MTRBaseClusterApplicationLauncher] from an unsafe.Pointer.
func MTRBaseClusterApplicationLauncherFrom(ptr unsafe.Pointer) MTRBaseClusterApplicationLauncher {
	return MTRBaseClusterApplicationLauncher{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterApplicationLauncherClass) Alloc() MTRBaseClusterApplicationLauncher {
	rv := objc.Send[MTRBaseClusterApplicationLauncher](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterApplicationLauncherClass) New() MTRBaseClusterApplicationLauncher {
	rv := objc.Send[MTRBaseClusterApplicationLauncher](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterApplicationLauncher) Init() MTRBaseClusterApplicationLauncher {
	rv := objc.Send[MTRBaseClusterApplicationLauncher](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterApplicationLauncher) Autorelease() MTRBaseClusterApplicationLauncher {
	rv := objc.Send[MTRBaseClusterApplicationLauncher](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterApplicationLauncher creates a new MTRBaseClusterApplicationLauncher instance.
func NewMTRBaseClusterApplicationLauncher() MTRBaseClusterApplicationLauncher {
	return getMTRBaseClusterApplicationLauncherClass().New()
}




