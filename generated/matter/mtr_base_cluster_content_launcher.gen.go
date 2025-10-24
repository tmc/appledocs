// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterContentLauncher] class.
var (
	MTRBaseClusterContentLauncherClass     _MTRBaseClusterContentLauncherClass
	MTRBaseClusterContentLauncherClassOnce sync.Once
)

func getMTRBaseClusterContentLauncherClass() _MTRBaseClusterContentLauncherClass {
	MTRBaseClusterContentLauncherClassOnce.Do(func() {
		MTRBaseClusterContentLauncherClass = _MTRBaseClusterContentLauncherClass{objc.GetClass("MTRBaseClusterContentLauncher")}
	})
	return MTRBaseClusterContentLauncherClass
}

type _MTRBaseClusterContentLauncherClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterContentLauncher] class.
type IMTRBaseClusterContentLauncher interface {
	IMTRGenericBaseCluster
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterContentLauncher
type MTRBaseClusterContentLauncher struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterContentLauncherFrom constructs a [MTRBaseClusterContentLauncher] from an unsafe.Pointer.
func MTRBaseClusterContentLauncherFrom(ptr unsafe.Pointer) MTRBaseClusterContentLauncher {
	return MTRBaseClusterContentLauncher{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterContentLauncherClass) Alloc() MTRBaseClusterContentLauncher {
	rv := objc.Send[MTRBaseClusterContentLauncher](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterContentLauncherClass) New() MTRBaseClusterContentLauncher {
	rv := objc.Send[MTRBaseClusterContentLauncher](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterContentLauncher) Init() MTRBaseClusterContentLauncher {
	rv := objc.Send[MTRBaseClusterContentLauncher](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterContentLauncher) Autorelease() MTRBaseClusterContentLauncher {
	rv := objc.Send[MTRBaseClusterContentLauncher](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterContentLauncher creates a new MTRBaseClusterContentLauncher instance.
func NewMTRBaseClusterContentLauncher() MTRBaseClusterContentLauncher {
	return getMTRBaseClusterContentLauncherClass().New()
}




