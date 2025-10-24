// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterIdentify] class.
var (
	MTRClusterIdentifyClass     _MTRClusterIdentifyClass
	MTRClusterIdentifyClassOnce sync.Once
)

func getMTRClusterIdentifyClass() _MTRClusterIdentifyClass {
	MTRClusterIdentifyClassOnce.Do(func() {
		MTRClusterIdentifyClass = _MTRClusterIdentifyClass{objc.GetClass("MTRClusterIdentify")}
	})
	return MTRClusterIdentifyClass
}

type _MTRClusterIdentifyClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterIdentify] class.
type IMTRClusterIdentify interface {
	IMTRGenericCluster
	// properties:
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterIdentify
type MTRClusterIdentify struct {
	MTRGenericCluster
}

// MTRClusterIdentifyFrom constructs a [MTRClusterIdentify] from an unsafe.Pointer.
func MTRClusterIdentifyFrom(ptr unsafe.Pointer) MTRClusterIdentify {
	return MTRClusterIdentify{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterIdentifyClass) Alloc() MTRClusterIdentify {
	rv := objc.Send[MTRClusterIdentify](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterIdentifyClass) New() MTRClusterIdentify {
	rv := objc.Send[MTRClusterIdentify](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterIdentify) Init() MTRClusterIdentify {
	rv := objc.Send[MTRClusterIdentify](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterIdentify) Autorelease() MTRClusterIdentify {
	rv := objc.Send[MTRClusterIdentify](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterIdentify creates a new MTRClusterIdentify instance.
func NewMTRClusterIdentify() MTRClusterIdentify {
	return getMTRClusterIdentifyClass().New()
}
