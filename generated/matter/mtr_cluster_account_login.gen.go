// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterAccountLogin] class.
var (
	MTRClusterAccountLoginClass     _MTRClusterAccountLoginClass
	MTRClusterAccountLoginClassOnce sync.Once
)

func getMTRClusterAccountLoginClass() _MTRClusterAccountLoginClass {
	MTRClusterAccountLoginClassOnce.Do(func() {
		MTRClusterAccountLoginClass = _MTRClusterAccountLoginClass{objc.GetClass("MTRClusterAccountLogin")}
	})
	return MTRClusterAccountLoginClass
}

type _MTRClusterAccountLoginClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterAccountLogin] class.
type IMTRClusterAccountLogin interface {
	IMTRGenericCluster
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterAccountLogin
type MTRClusterAccountLogin struct {
	MTRGenericCluster
}

// MTRClusterAccountLoginFrom constructs a [MTRClusterAccountLogin] from an unsafe.Pointer.
func MTRClusterAccountLoginFrom(ptr unsafe.Pointer) MTRClusterAccountLogin {
	return MTRClusterAccountLogin{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterAccountLoginClass) Alloc() MTRClusterAccountLogin {
	rv := objc.Send[MTRClusterAccountLogin](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterAccountLoginClass) New() MTRClusterAccountLogin {
	rv := objc.Send[MTRClusterAccountLogin](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterAccountLogin) Init() MTRClusterAccountLogin {
	rv := objc.Send[MTRClusterAccountLogin](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterAccountLogin) Autorelease() MTRClusterAccountLogin {
	rv := objc.Send[MTRClusterAccountLogin](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterAccountLogin creates a new MTRClusterAccountLogin instance.
func NewMTRClusterAccountLogin() MTRClusterAccountLogin {
	return getMTRClusterAccountLoginClass().New()
}




