// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRGeneralCommissioningClusterSetRegulatoryConfigResponseParams] class.
var (
	MTRGeneralCommissioningClusterSetRegulatoryConfigResponseParamsClass     _MTRGeneralCommissioningClusterSetRegulatoryConfigResponseParamsClass
	MTRGeneralCommissioningClusterSetRegulatoryConfigResponseParamsClassOnce sync.Once
)

func getMTRGeneralCommissioningClusterSetRegulatoryConfigResponseParamsClass() _MTRGeneralCommissioningClusterSetRegulatoryConfigResponseParamsClass {
	MTRGeneralCommissioningClusterSetRegulatoryConfigResponseParamsClassOnce.Do(func() {
		MTRGeneralCommissioningClusterSetRegulatoryConfigResponseParamsClass = _MTRGeneralCommissioningClusterSetRegulatoryConfigResponseParamsClass{objc.GetClass("MTRGeneralCommissioningClusterSetRegulatoryConfigResponseParams")}
	})
	return MTRGeneralCommissioningClusterSetRegulatoryConfigResponseParamsClass
}

type _MTRGeneralCommissioningClusterSetRegulatoryConfigResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRGeneralCommissioningClusterSetRegulatoryConfigResponseParams] class.
type IMTRGeneralCommissioningClusterSetRegulatoryConfigResponseParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGeneralCommissioningClusterSetRegulatoryConfigResponseParams
type MTRGeneralCommissioningClusterSetRegulatoryConfigResponseParams struct {
	objectivec.Object
}

// MTRGeneralCommissioningClusterSetRegulatoryConfigResponseParamsFrom constructs a [MTRGeneralCommissioningClusterSetRegulatoryConfigResponseParams] from an unsafe.Pointer.
func MTRGeneralCommissioningClusterSetRegulatoryConfigResponseParamsFrom(ptr unsafe.Pointer) MTRGeneralCommissioningClusterSetRegulatoryConfigResponseParams {
	return MTRGeneralCommissioningClusterSetRegulatoryConfigResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRGeneralCommissioningClusterSetRegulatoryConfigResponseParamsClass) Alloc() MTRGeneralCommissioningClusterSetRegulatoryConfigResponseParams {
	rv := objc.Send[MTRGeneralCommissioningClusterSetRegulatoryConfigResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRGeneralCommissioningClusterSetRegulatoryConfigResponseParamsClass) New() MTRGeneralCommissioningClusterSetRegulatoryConfigResponseParams {
	rv := objc.Send[MTRGeneralCommissioningClusterSetRegulatoryConfigResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRGeneralCommissioningClusterSetRegulatoryConfigResponseParams) Init() MTRGeneralCommissioningClusterSetRegulatoryConfigResponseParams {
	rv := objc.Send[MTRGeneralCommissioningClusterSetRegulatoryConfigResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRGeneralCommissioningClusterSetRegulatoryConfigResponseParams) Autorelease() MTRGeneralCommissioningClusterSetRegulatoryConfigResponseParams {
	rv := objc.Send[MTRGeneralCommissioningClusterSetRegulatoryConfigResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRGeneralCommissioningClusterSetRegulatoryConfigResponseParams creates a new MTRGeneralCommissioningClusterSetRegulatoryConfigResponseParams instance.
func NewMTRGeneralCommissioningClusterSetRegulatoryConfigResponseParams() MTRGeneralCommissioningClusterSetRegulatoryConfigResponseParams {
	return getMTRGeneralCommissioningClusterSetRegulatoryConfigResponseParamsClass().New()
}




