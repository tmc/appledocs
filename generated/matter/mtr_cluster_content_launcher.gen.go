// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterContentLauncher] class.
var (
	MTRClusterContentLauncherClass     _MTRClusterContentLauncherClass
	MTRClusterContentLauncherClassOnce sync.Once
)

func getMTRClusterContentLauncherClass() _MTRClusterContentLauncherClass {
	MTRClusterContentLauncherClassOnce.Do(func() {
		MTRClusterContentLauncherClass = _MTRClusterContentLauncherClass{objc.GetClass("MTRClusterContentLauncher")}
	})
	return MTRClusterContentLauncherClass
}

type _MTRClusterContentLauncherClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterContentLauncher] class.
type IMTRClusterContentLauncher interface {
	IMTRGenericCluster
	// properties:
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterContentLauncher
type MTRClusterContentLauncher struct {
	MTRGenericCluster
}

// MTRClusterContentLauncherFrom constructs a [MTRClusterContentLauncher] from an unsafe.Pointer.
func MTRClusterContentLauncherFrom(ptr unsafe.Pointer) MTRClusterContentLauncher {
	return MTRClusterContentLauncher{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterContentLauncherClass) Alloc() MTRClusterContentLauncher {
	rv := objc.Send[MTRClusterContentLauncher](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterContentLauncherClass) New() MTRClusterContentLauncher {
	rv := objc.Send[MTRClusterContentLauncher](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterContentLauncher) Init() MTRClusterContentLauncher {
	rv := objc.Send[MTRClusterContentLauncher](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterContentLauncher) Autorelease() MTRClusterContentLauncher {
	rv := objc.Send[MTRClusterContentLauncher](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterContentLauncher creates a new MTRClusterContentLauncher instance.
func NewMTRClusterContentLauncher() MTRClusterContentLauncher {
	return getMTRClusterContentLauncherClass().New()
}
