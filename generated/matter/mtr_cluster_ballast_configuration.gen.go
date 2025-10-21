// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterBallastConfiguration] class.
var (
	MTRClusterBallastConfigurationClass     _MTRClusterBallastConfigurationClass
	MTRClusterBallastConfigurationClassOnce sync.Once
)

func getMTRClusterBallastConfigurationClass() _MTRClusterBallastConfigurationClass {
	MTRClusterBallastConfigurationClassOnce.Do(func() {
		MTRClusterBallastConfigurationClass = _MTRClusterBallastConfigurationClass{objc.GetClass("MTRClusterBallastConfiguration")}
	})
	return MTRClusterBallastConfigurationClass
}

type _MTRClusterBallastConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterBallastConfiguration] class.
type IMTRClusterBallastConfiguration interface {
	IMTRGenericCluster
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterBallastConfiguration
type MTRClusterBallastConfiguration struct {
	MTRGenericCluster
}

// MTRClusterBallastConfigurationFrom constructs a [MTRClusterBallastConfiguration] from an unsafe.Pointer.
func MTRClusterBallastConfigurationFrom(ptr unsafe.Pointer) MTRClusterBallastConfiguration {
	return MTRClusterBallastConfiguration{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterBallastConfigurationClass) Alloc() MTRClusterBallastConfiguration {
	rv := objc.Send[MTRClusterBallastConfiguration](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterBallastConfigurationClass) New() MTRClusterBallastConfiguration {
	rv := objc.Send[MTRClusterBallastConfiguration](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterBallastConfiguration) Init() MTRClusterBallastConfiguration {
	rv := objc.Send[MTRClusterBallastConfiguration](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterBallastConfiguration) Autorelease() MTRClusterBallastConfiguration {
	rv := objc.Send[MTRClusterBallastConfiguration](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterBallastConfiguration creates a new MTRClusterBallastConfiguration instance.
func NewMTRClusterBallastConfiguration() MTRClusterBallastConfiguration {
	return getMTRClusterBallastConfigurationClass().New()
}




