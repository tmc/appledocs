// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTROperationalCredentialsClusterNOCResponseParams] class.
var (
	MTROperationalCredentialsClusterNOCResponseParamsClass     _MTROperationalCredentialsClusterNOCResponseParamsClass
	MTROperationalCredentialsClusterNOCResponseParamsClassOnce sync.Once
)

func getMTROperationalCredentialsClusterNOCResponseParamsClass() _MTROperationalCredentialsClusterNOCResponseParamsClass {
	MTROperationalCredentialsClusterNOCResponseParamsClassOnce.Do(func() {
		MTROperationalCredentialsClusterNOCResponseParamsClass = _MTROperationalCredentialsClusterNOCResponseParamsClass{objc.GetClass("MTROperationalCredentialsClusterNOCResponseParams")}
	})
	return MTROperationalCredentialsClusterNOCResponseParamsClass
}

type _MTROperationalCredentialsClusterNOCResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTROperationalCredentialsClusterNOCResponseParams] class.
type IMTROperationalCredentialsClusterNOCResponseParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterNOCResponseParams
type MTROperationalCredentialsClusterNOCResponseParams struct {
	objectivec.Object
}

// MTROperationalCredentialsClusterNOCResponseParamsFrom constructs a [MTROperationalCredentialsClusterNOCResponseParams] from an unsafe.Pointer.
func MTROperationalCredentialsClusterNOCResponseParamsFrom(ptr unsafe.Pointer) MTROperationalCredentialsClusterNOCResponseParams {
	return MTROperationalCredentialsClusterNOCResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROperationalCredentialsClusterNOCResponseParamsClass) Alloc() MTROperationalCredentialsClusterNOCResponseParams {
	rv := objc.Send[MTROperationalCredentialsClusterNOCResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROperationalCredentialsClusterNOCResponseParamsClass) New() MTROperationalCredentialsClusterNOCResponseParams {
	rv := objc.Send[MTROperationalCredentialsClusterNOCResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROperationalCredentialsClusterNOCResponseParams) Init() MTROperationalCredentialsClusterNOCResponseParams {
	rv := objc.Send[MTROperationalCredentialsClusterNOCResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROperationalCredentialsClusterNOCResponseParams) Autorelease() MTROperationalCredentialsClusterNOCResponseParams {
	rv := objc.Send[MTROperationalCredentialsClusterNOCResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROperationalCredentialsClusterNOCResponseParams creates a new MTROperationalCredentialsClusterNOCResponseParams instance.
func NewMTROperationalCredentialsClusterNOCResponseParams() MTROperationalCredentialsClusterNOCResponseParams {
	return getMTROperationalCredentialsClusterNOCResponseParamsClass().New()
}




