// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterIdentify] class.
var (
	MTRBaseClusterIdentifyClass     _MTRBaseClusterIdentifyClass
	MTRBaseClusterIdentifyClassOnce sync.Once
)

func getMTRBaseClusterIdentifyClass() _MTRBaseClusterIdentifyClass {
	MTRBaseClusterIdentifyClassOnce.Do(func() {
		MTRBaseClusterIdentifyClass = _MTRBaseClusterIdentifyClass{objc.GetClass("MTRBaseClusterIdentify")}
	})
	return MTRBaseClusterIdentifyClass
}

type _MTRBaseClusterIdentifyClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterIdentify] class.
type IMTRBaseClusterIdentify interface {
	IMTRGenericBaseCluster
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterIdentify
type MTRBaseClusterIdentify struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterIdentifyFrom constructs a [MTRBaseClusterIdentify] from an unsafe.Pointer.
func MTRBaseClusterIdentifyFrom(ptr unsafe.Pointer) MTRBaseClusterIdentify {
	return MTRBaseClusterIdentify{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterIdentifyClass) Alloc() MTRBaseClusterIdentify {
	rv := objc.Send[MTRBaseClusterIdentify](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterIdentifyClass) New() MTRBaseClusterIdentify {
	rv := objc.Send[MTRBaseClusterIdentify](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterIdentify) Init() MTRBaseClusterIdentify {
	rv := objc.Send[MTRBaseClusterIdentify](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterIdentify) Autorelease() MTRBaseClusterIdentify {
	rv := objc.Send[MTRBaseClusterIdentify](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterIdentify creates a new MTRBaseClusterIdentify instance.
func NewMTRBaseClusterIdentify() MTRBaseClusterIdentify {
	return getMTRBaseClusterIdentifyClass().New()
}




