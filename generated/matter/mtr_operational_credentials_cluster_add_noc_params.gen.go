// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTROperationalCredentialsClusterAddNOCParams] class.
var (
	MTROperationalCredentialsClusterAddNOCParamsClass     _MTROperationalCredentialsClusterAddNOCParamsClass
	MTROperationalCredentialsClusterAddNOCParamsClassOnce sync.Once
)

func getMTROperationalCredentialsClusterAddNOCParamsClass() _MTROperationalCredentialsClusterAddNOCParamsClass {
	MTROperationalCredentialsClusterAddNOCParamsClassOnce.Do(func() {
		MTROperationalCredentialsClusterAddNOCParamsClass = _MTROperationalCredentialsClusterAddNOCParamsClass{objc.GetClass("MTROperationalCredentialsClusterAddNOCParams")}
	})
	return MTROperationalCredentialsClusterAddNOCParamsClass
}

type _MTROperationalCredentialsClusterAddNOCParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTROperationalCredentialsClusterAddNOCParams] class.
type IMTROperationalCredentialsClusterAddNOCParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterAddNOCParams
type MTROperationalCredentialsClusterAddNOCParams struct {
	objectivec.Object
}

// MTROperationalCredentialsClusterAddNOCParamsFrom constructs a [MTROperationalCredentialsClusterAddNOCParams] from an unsafe.Pointer.
func MTROperationalCredentialsClusterAddNOCParamsFrom(ptr unsafe.Pointer) MTROperationalCredentialsClusterAddNOCParams {
	return MTROperationalCredentialsClusterAddNOCParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROperationalCredentialsClusterAddNOCParamsClass) Alloc() MTROperationalCredentialsClusterAddNOCParams {
	rv := objc.Send[MTROperationalCredentialsClusterAddNOCParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROperationalCredentialsClusterAddNOCParamsClass) New() MTROperationalCredentialsClusterAddNOCParams {
	rv := objc.Send[MTROperationalCredentialsClusterAddNOCParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROperationalCredentialsClusterAddNOCParams) Init() MTROperationalCredentialsClusterAddNOCParams {
	rv := objc.Send[MTROperationalCredentialsClusterAddNOCParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROperationalCredentialsClusterAddNOCParams) Autorelease() MTROperationalCredentialsClusterAddNOCParams {
	rv := objc.Send[MTROperationalCredentialsClusterAddNOCParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROperationalCredentialsClusterAddNOCParams creates a new MTROperationalCredentialsClusterAddNOCParams instance.
func NewMTROperationalCredentialsClusterAddNOCParams() MTROperationalCredentialsClusterAddNOCParams {
	return getMTROperationalCredentialsClusterAddNOCParamsClass().New()
}




