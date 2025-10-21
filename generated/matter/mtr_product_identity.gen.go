// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRProductIdentity] class.
var (
	MTRProductIdentityClass     _MTRProductIdentityClass
	MTRProductIdentityClassOnce sync.Once
)

func getMTRProductIdentityClass() _MTRProductIdentityClass {
	MTRProductIdentityClassOnce.Do(func() {
		MTRProductIdentityClass = _MTRProductIdentityClass{objc.GetClass("MTRProductIdentity")}
	})
	return MTRProductIdentityClass
}

type _MTRProductIdentityClass struct {
	class objc.Class
}

// An interface definition for the [MTRProductIdentity] class.
type IMTRProductIdentity interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRProductIdentity
type MTRProductIdentity struct {
	objectivec.Object
}

// MTRProductIdentityFrom constructs a [MTRProductIdentity] from an unsafe.Pointer.
func MTRProductIdentityFrom(ptr unsafe.Pointer) MTRProductIdentity {
	return MTRProductIdentity{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRProductIdentityClass) Alloc() MTRProductIdentity {
	rv := objc.Send[MTRProductIdentity](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRProductIdentityClass) New() MTRProductIdentity {
	rv := objc.Send[MTRProductIdentity](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRProductIdentity) Init() MTRProductIdentity {
	rv := objc.Send[MTRProductIdentity](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRProductIdentity) Autorelease() MTRProductIdentity {
	rv := objc.Send[MTRProductIdentity](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRProductIdentity creates a new MTRProductIdentity instance.
func NewMTRProductIdentity() MTRProductIdentity {
	return getMTRProductIdentityClass().New()
}




