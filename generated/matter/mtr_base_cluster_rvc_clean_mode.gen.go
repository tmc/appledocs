// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterRVCCleanMode] class.
var (
	MTRBaseClusterRVCCleanModeClass     _MTRBaseClusterRVCCleanModeClass
	MTRBaseClusterRVCCleanModeClassOnce sync.Once
)

func getMTRBaseClusterRVCCleanModeClass() _MTRBaseClusterRVCCleanModeClass {
	MTRBaseClusterRVCCleanModeClassOnce.Do(func() {
		MTRBaseClusterRVCCleanModeClass = _MTRBaseClusterRVCCleanModeClass{objc.GetClass("MTRBaseClusterRVCCleanMode")}
	})
	return MTRBaseClusterRVCCleanModeClass
}

type _MTRBaseClusterRVCCleanModeClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterRVCCleanMode] class.
type IMTRBaseClusterRVCCleanMode interface {
	IMTRGenericBaseCluster
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterRVCCleanMode
type MTRBaseClusterRVCCleanMode struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterRVCCleanModeFrom constructs a [MTRBaseClusterRVCCleanMode] from an unsafe.Pointer.
func MTRBaseClusterRVCCleanModeFrom(ptr unsafe.Pointer) MTRBaseClusterRVCCleanMode {
	return MTRBaseClusterRVCCleanMode{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterRVCCleanModeClass) Alloc() MTRBaseClusterRVCCleanMode {
	rv := objc.Send[MTRBaseClusterRVCCleanMode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterRVCCleanModeClass) New() MTRBaseClusterRVCCleanMode {
	rv := objc.Send[MTRBaseClusterRVCCleanMode](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterRVCCleanMode) Init() MTRBaseClusterRVCCleanMode {
	rv := objc.Send[MTRBaseClusterRVCCleanMode](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterRVCCleanMode) Autorelease() MTRBaseClusterRVCCleanMode {
	rv := objc.Send[MTRBaseClusterRVCCleanMode](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterRVCCleanMode creates a new MTRBaseClusterRVCCleanMode instance.
func NewMTRBaseClusterRVCCleanMode() MTRBaseClusterRVCCleanMode {
	return getMTRBaseClusterRVCCleanModeClass().New()
}




