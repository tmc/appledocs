// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTROperationalCredentialsClusterUpdateNOCParams] class.
var (
	MTROperationalCredentialsClusterUpdateNOCParamsClass     _MTROperationalCredentialsClusterUpdateNOCParamsClass
	MTROperationalCredentialsClusterUpdateNOCParamsClassOnce sync.Once
)

func getMTROperationalCredentialsClusterUpdateNOCParamsClass() _MTROperationalCredentialsClusterUpdateNOCParamsClass {
	MTROperationalCredentialsClusterUpdateNOCParamsClassOnce.Do(func() {
		MTROperationalCredentialsClusterUpdateNOCParamsClass = _MTROperationalCredentialsClusterUpdateNOCParamsClass{objc.GetClass("MTROperationalCredentialsClusterUpdateNOCParams")}
	})
	return MTROperationalCredentialsClusterUpdateNOCParamsClass
}

type _MTROperationalCredentialsClusterUpdateNOCParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTROperationalCredentialsClusterUpdateNOCParams] class.
type IMTROperationalCredentialsClusterUpdateNOCParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterUpdateNOCParams
type MTROperationalCredentialsClusterUpdateNOCParams struct {
	objectivec.Object
}

// MTROperationalCredentialsClusterUpdateNOCParamsFrom constructs a [MTROperationalCredentialsClusterUpdateNOCParams] from an unsafe.Pointer.
func MTROperationalCredentialsClusterUpdateNOCParamsFrom(ptr unsafe.Pointer) MTROperationalCredentialsClusterUpdateNOCParams {
	return MTROperationalCredentialsClusterUpdateNOCParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROperationalCredentialsClusterUpdateNOCParamsClass) Alloc() MTROperationalCredentialsClusterUpdateNOCParams {
	rv := objc.Send[MTROperationalCredentialsClusterUpdateNOCParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROperationalCredentialsClusterUpdateNOCParamsClass) New() MTROperationalCredentialsClusterUpdateNOCParams {
	rv := objc.Send[MTROperationalCredentialsClusterUpdateNOCParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROperationalCredentialsClusterUpdateNOCParams) Init() MTROperationalCredentialsClusterUpdateNOCParams {
	rv := objc.Send[MTROperationalCredentialsClusterUpdateNOCParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROperationalCredentialsClusterUpdateNOCParams) Autorelease() MTROperationalCredentialsClusterUpdateNOCParams {
	rv := objc.Send[MTROperationalCredentialsClusterUpdateNOCParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROperationalCredentialsClusterUpdateNOCParams creates a new MTROperationalCredentialsClusterUpdateNOCParams instance.
func NewMTROperationalCredentialsClusterUpdateNOCParams() MTROperationalCredentialsClusterUpdateNOCParams {
	return getMTROperationalCredentialsClusterUpdateNOCParamsClass().New()
}




