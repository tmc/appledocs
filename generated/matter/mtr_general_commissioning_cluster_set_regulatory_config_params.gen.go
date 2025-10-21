// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRGeneralCommissioningClusterSetRegulatoryConfigParams] class.
var (
	MTRGeneralCommissioningClusterSetRegulatoryConfigParamsClass     _MTRGeneralCommissioningClusterSetRegulatoryConfigParamsClass
	MTRGeneralCommissioningClusterSetRegulatoryConfigParamsClassOnce sync.Once
)

func getMTRGeneralCommissioningClusterSetRegulatoryConfigParamsClass() _MTRGeneralCommissioningClusterSetRegulatoryConfigParamsClass {
	MTRGeneralCommissioningClusterSetRegulatoryConfigParamsClassOnce.Do(func() {
		MTRGeneralCommissioningClusterSetRegulatoryConfigParamsClass = _MTRGeneralCommissioningClusterSetRegulatoryConfigParamsClass{objc.GetClass("MTRGeneralCommissioningClusterSetRegulatoryConfigParams")}
	})
	return MTRGeneralCommissioningClusterSetRegulatoryConfigParamsClass
}

type _MTRGeneralCommissioningClusterSetRegulatoryConfigParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRGeneralCommissioningClusterSetRegulatoryConfigParams] class.
type IMTRGeneralCommissioningClusterSetRegulatoryConfigParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGeneralCommissioningClusterSetRegulatoryConfigParams
type MTRGeneralCommissioningClusterSetRegulatoryConfigParams struct {
	objectivec.Object
}

// MTRGeneralCommissioningClusterSetRegulatoryConfigParamsFrom constructs a [MTRGeneralCommissioningClusterSetRegulatoryConfigParams] from an unsafe.Pointer.
func MTRGeneralCommissioningClusterSetRegulatoryConfigParamsFrom(ptr unsafe.Pointer) MTRGeneralCommissioningClusterSetRegulatoryConfigParams {
	return MTRGeneralCommissioningClusterSetRegulatoryConfigParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRGeneralCommissioningClusterSetRegulatoryConfigParamsClass) Alloc() MTRGeneralCommissioningClusterSetRegulatoryConfigParams {
	rv := objc.Send[MTRGeneralCommissioningClusterSetRegulatoryConfigParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRGeneralCommissioningClusterSetRegulatoryConfigParamsClass) New() MTRGeneralCommissioningClusterSetRegulatoryConfigParams {
	rv := objc.Send[MTRGeneralCommissioningClusterSetRegulatoryConfigParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRGeneralCommissioningClusterSetRegulatoryConfigParams) Init() MTRGeneralCommissioningClusterSetRegulatoryConfigParams {
	rv := objc.Send[MTRGeneralCommissioningClusterSetRegulatoryConfigParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRGeneralCommissioningClusterSetRegulatoryConfigParams) Autorelease() MTRGeneralCommissioningClusterSetRegulatoryConfigParams {
	rv := objc.Send[MTRGeneralCommissioningClusterSetRegulatoryConfigParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRGeneralCommissioningClusterSetRegulatoryConfigParams creates a new MTRGeneralCommissioningClusterSetRegulatoryConfigParams instance.
func NewMTRGeneralCommissioningClusterSetRegulatoryConfigParams() MTRGeneralCommissioningClusterSetRegulatoryConfigParams {
	return getMTRGeneralCommissioningClusterSetRegulatoryConfigParamsClass().New()
}




