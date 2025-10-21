// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterRVCRunMode] class.
var (
	MTRClusterRVCRunModeClass     _MTRClusterRVCRunModeClass
	MTRClusterRVCRunModeClassOnce sync.Once
)

func getMTRClusterRVCRunModeClass() _MTRClusterRVCRunModeClass {
	MTRClusterRVCRunModeClassOnce.Do(func() {
		MTRClusterRVCRunModeClass = _MTRClusterRVCRunModeClass{objc.GetClass("MTRClusterRVCRunMode")}
	})
	return MTRClusterRVCRunModeClass
}

type _MTRClusterRVCRunModeClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterRVCRunMode] class.
type IMTRClusterRVCRunMode interface {
	IMTRGenericCluster
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterRVCRunMode
type MTRClusterRVCRunMode struct {
	MTRGenericCluster
}

// MTRClusterRVCRunModeFrom constructs a [MTRClusterRVCRunMode] from an unsafe.Pointer.
func MTRClusterRVCRunModeFrom(ptr unsafe.Pointer) MTRClusterRVCRunMode {
	return MTRClusterRVCRunMode{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterRVCRunModeClass) Alloc() MTRClusterRVCRunMode {
	rv := objc.Send[MTRClusterRVCRunMode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterRVCRunModeClass) New() MTRClusterRVCRunMode {
	rv := objc.Send[MTRClusterRVCRunMode](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterRVCRunMode) Init() MTRClusterRVCRunMode {
	rv := objc.Send[MTRClusterRVCRunMode](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterRVCRunMode) Autorelease() MTRClusterRVCRunMode {
	rv := objc.Send[MTRClusterRVCRunMode](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterRVCRunMode creates a new MTRClusterRVCRunMode instance.
func NewMTRClusterRVCRunMode() MTRClusterRVCRunMode {
	return getMTRClusterRVCRunModeClass().New()
}




