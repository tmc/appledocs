// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterRVCRunMode] class.
var (
	MTRBaseClusterRVCRunModeClass     _MTRBaseClusterRVCRunModeClass
	MTRBaseClusterRVCRunModeClassOnce sync.Once
)

func getMTRBaseClusterRVCRunModeClass() _MTRBaseClusterRVCRunModeClass {
	MTRBaseClusterRVCRunModeClassOnce.Do(func() {
		MTRBaseClusterRVCRunModeClass = _MTRBaseClusterRVCRunModeClass{objc.GetClass("MTRBaseClusterRVCRunMode")}
	})
	return MTRBaseClusterRVCRunModeClass
}

type _MTRBaseClusterRVCRunModeClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterRVCRunMode] class.
type IMTRBaseClusterRVCRunMode interface {
	IMTRGenericBaseCluster
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterRVCRunMode
type MTRBaseClusterRVCRunMode struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterRVCRunModeFrom constructs a [MTRBaseClusterRVCRunMode] from an unsafe.Pointer.
func MTRBaseClusterRVCRunModeFrom(ptr unsafe.Pointer) MTRBaseClusterRVCRunMode {
	return MTRBaseClusterRVCRunMode{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterRVCRunModeClass) Alloc() MTRBaseClusterRVCRunMode {
	rv := objc.Send[MTRBaseClusterRVCRunMode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterRVCRunModeClass) New() MTRBaseClusterRVCRunMode {
	rv := objc.Send[MTRBaseClusterRVCRunMode](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterRVCRunMode) Init() MTRBaseClusterRVCRunMode {
	rv := objc.Send[MTRBaseClusterRVCRunMode](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterRVCRunMode) Autorelease() MTRBaseClusterRVCRunMode {
	rv := objc.Send[MTRBaseClusterRVCRunMode](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterRVCRunMode creates a new MTRBaseClusterRVCRunMode instance.
func NewMTRBaseClusterRVCRunMode() MTRBaseClusterRVCRunMode {
	return getMTRBaseClusterRVCRunModeClass().New()
}




