// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterAccountLogin] class.
var (
	MTRBaseClusterAccountLoginClass     _MTRBaseClusterAccountLoginClass
	MTRBaseClusterAccountLoginClassOnce sync.Once
)

func getMTRBaseClusterAccountLoginClass() _MTRBaseClusterAccountLoginClass {
	MTRBaseClusterAccountLoginClassOnce.Do(func() {
		MTRBaseClusterAccountLoginClass = _MTRBaseClusterAccountLoginClass{objc.GetClass("MTRBaseClusterAccountLogin")}
	})
	return MTRBaseClusterAccountLoginClass
}

type _MTRBaseClusterAccountLoginClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterAccountLogin] class.
type IMTRBaseClusterAccountLogin interface {
	IMTRGenericBaseCluster
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterAccountLogin
type MTRBaseClusterAccountLogin struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterAccountLoginFrom constructs a [MTRBaseClusterAccountLogin] from an unsafe.Pointer.
func MTRBaseClusterAccountLoginFrom(ptr unsafe.Pointer) MTRBaseClusterAccountLogin {
	return MTRBaseClusterAccountLogin{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterAccountLoginClass) Alloc() MTRBaseClusterAccountLogin {
	rv := objc.Send[MTRBaseClusterAccountLogin](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterAccountLoginClass) New() MTRBaseClusterAccountLogin {
	rv := objc.Send[MTRBaseClusterAccountLogin](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterAccountLogin) Init() MTRBaseClusterAccountLogin {
	rv := objc.Send[MTRBaseClusterAccountLogin](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterAccountLogin) Autorelease() MTRBaseClusterAccountLogin {
	rv := objc.Send[MTRBaseClusterAccountLogin](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterAccountLogin creates a new MTRBaseClusterAccountLogin instance.
func NewMTRBaseClusterAccountLogin() MTRBaseClusterAccountLogin {
	return getMTRBaseClusterAccountLoginClass().New()
}




