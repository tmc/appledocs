// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneralcommissioningclustersetregulatoryconfigresponseparams/debugtext
func (m_ MTRGeneralCommissioningClusterSetRegulatoryConfigResponseParams) DebugText() string {
	rv := objc.Send[string](m_.ID, objc.Sel("debugText"))
	return rv
}


// SetDebugText sets the value of the debugText property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneralcommissioningclustersetregulatoryconfigresponseparams/debugtext
func (m_ MTRGeneralCommissioningClusterSetRegulatoryConfigResponseParams) SetDebugText(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDebugText:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneralcommissioningclustersetregulatoryconfigresponseparams/errorcode
func (m_ MTRGeneralCommissioningClusterSetRegulatoryConfigResponseParams) ErrorCode() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("errorCode"))
	return rv
}


// SetErrorCode sets the value of the errorCode property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneralcommissioningclustersetregulatoryconfigresponseparams/errorcode
func (m_ MTRGeneralCommissioningClusterSetRegulatoryConfigResponseParams) SetErrorCode(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setErrorCode:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneralcommissioningclustersetregulatoryconfigresponseparams/timedinvoketimeoutms
func (m_ MTRGeneralCommissioningClusterSetRegulatoryConfigResponseParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneralcommissioningclustersetregulatoryconfigresponseparams/timedinvoketimeoutms
func (m_ MTRGeneralCommissioningClusterSetRegulatoryConfigResponseParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



