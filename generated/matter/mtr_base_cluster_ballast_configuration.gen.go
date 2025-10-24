// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterBallastConfiguration] class.
var (
	MTRBaseClusterBallastConfigurationClass     _MTRBaseClusterBallastConfigurationClass
	MTRBaseClusterBallastConfigurationClassOnce sync.Once
)

func getMTRBaseClusterBallastConfigurationClass() _MTRBaseClusterBallastConfigurationClass {
	MTRBaseClusterBallastConfigurationClassOnce.Do(func() {
		MTRBaseClusterBallastConfigurationClass = _MTRBaseClusterBallastConfigurationClass{objc.GetClass("MTRBaseClusterBallastConfiguration")}
	})
	return MTRBaseClusterBallastConfigurationClass
}

type _MTRBaseClusterBallastConfigurationClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterBallastConfiguration] class.
type IMTRBaseClusterBallastConfiguration interface {
	IMTRGenericBaseCluster
	// properties:
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterBallastConfiguration
type MTRBaseClusterBallastConfiguration struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterBallastConfigurationFrom constructs a [MTRBaseClusterBallastConfiguration] from an unsafe.Pointer.
func MTRBaseClusterBallastConfigurationFrom(ptr unsafe.Pointer) MTRBaseClusterBallastConfiguration {
	return MTRBaseClusterBallastConfiguration{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterBallastConfigurationClass) Alloc() MTRBaseClusterBallastConfiguration {
	rv := objc.Send[MTRBaseClusterBallastConfiguration](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterBallastConfigurationClass) New() MTRBaseClusterBallastConfiguration {
	rv := objc.Send[MTRBaseClusterBallastConfiguration](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterBallastConfiguration) Init() MTRBaseClusterBallastConfiguration {
	rv := objc.Send[MTRBaseClusterBallastConfiguration](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterBallastConfiguration) Autorelease() MTRBaseClusterBallastConfiguration {
	rv := objc.Send[MTRBaseClusterBallastConfiguration](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterBallastConfiguration creates a new MTRBaseClusterBallastConfiguration instance.
func NewMTRBaseClusterBallastConfiguration() MTRBaseClusterBallastConfiguration {
	return getMTRBaseClusterBallastConfigurationClass().New()
}
