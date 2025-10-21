// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterRVCCleanMode] class.
var (
	MTRClusterRVCCleanModeClass     _MTRClusterRVCCleanModeClass
	MTRClusterRVCCleanModeClassOnce sync.Once
)

func getMTRClusterRVCCleanModeClass() _MTRClusterRVCCleanModeClass {
	MTRClusterRVCCleanModeClassOnce.Do(func() {
		MTRClusterRVCCleanModeClass = _MTRClusterRVCCleanModeClass{objc.GetClass("MTRClusterRVCCleanMode")}
	})
	return MTRClusterRVCCleanModeClass
}

type _MTRClusterRVCCleanModeClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterRVCCleanMode] class.
type IMTRClusterRVCCleanMode interface {
	IMTRGenericCluster
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterRVCCleanMode
type MTRClusterRVCCleanMode struct {
	MTRGenericCluster
}

// MTRClusterRVCCleanModeFrom constructs a [MTRClusterRVCCleanMode] from an unsafe.Pointer.
func MTRClusterRVCCleanModeFrom(ptr unsafe.Pointer) MTRClusterRVCCleanMode {
	return MTRClusterRVCCleanMode{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterRVCCleanModeClass) Alloc() MTRClusterRVCCleanMode {
	rv := objc.Send[MTRClusterRVCCleanMode](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterRVCCleanModeClass) New() MTRClusterRVCCleanMode {
	rv := objc.Send[MTRClusterRVCCleanMode](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterRVCCleanMode) Init() MTRClusterRVCCleanMode {
	rv := objc.Send[MTRClusterRVCCleanMode](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterRVCCleanMode) Autorelease() MTRClusterRVCCleanMode {
	rv := objc.Send[MTRClusterRVCCleanMode](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterRVCCleanMode creates a new MTRClusterRVCCleanMode instance.
func NewMTRClusterRVCCleanMode() MTRClusterRVCCleanMode {
	return getMTRClusterRVCCleanModeClass().New()
}




