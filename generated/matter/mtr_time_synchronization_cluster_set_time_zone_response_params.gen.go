// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRTimeSynchronizationClusterSetTimeZoneResponseParams] class.
var (
	MTRTimeSynchronizationClusterSetTimeZoneResponseParamsClass     _MTRTimeSynchronizationClusterSetTimeZoneResponseParamsClass
	MTRTimeSynchronizationClusterSetTimeZoneResponseParamsClassOnce sync.Once
)

func getMTRTimeSynchronizationClusterSetTimeZoneResponseParamsClass() _MTRTimeSynchronizationClusterSetTimeZoneResponseParamsClass {
	MTRTimeSynchronizationClusterSetTimeZoneResponseParamsClassOnce.Do(func() {
		MTRTimeSynchronizationClusterSetTimeZoneResponseParamsClass = _MTRTimeSynchronizationClusterSetTimeZoneResponseParamsClass{objc.GetClass("MTRTimeSynchronizationClusterSetTimeZoneResponseParams")}
	})
	return MTRTimeSynchronizationClusterSetTimeZoneResponseParamsClass
}

type _MTRTimeSynchronizationClusterSetTimeZoneResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRTimeSynchronizationClusterSetTimeZoneResponseParams] class.
type IMTRTimeSynchronizationClusterSetTimeZoneResponseParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterSetTimeZoneResponseParams
type MTRTimeSynchronizationClusterSetTimeZoneResponseParams struct {
	objectivec.Object
}

// MTRTimeSynchronizationClusterSetTimeZoneResponseParamsFrom constructs a [MTRTimeSynchronizationClusterSetTimeZoneResponseParams] from an unsafe.Pointer.
func MTRTimeSynchronizationClusterSetTimeZoneResponseParamsFrom(ptr unsafe.Pointer) MTRTimeSynchronizationClusterSetTimeZoneResponseParams {
	return MTRTimeSynchronizationClusterSetTimeZoneResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTimeSynchronizationClusterSetTimeZoneResponseParamsClass) Alloc() MTRTimeSynchronizationClusterSetTimeZoneResponseParams {
	rv := objc.Send[MTRTimeSynchronizationClusterSetTimeZoneResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTimeSynchronizationClusterSetTimeZoneResponseParamsClass) New() MTRTimeSynchronizationClusterSetTimeZoneResponseParams {
	rv := objc.Send[MTRTimeSynchronizationClusterSetTimeZoneResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTimeSynchronizationClusterSetTimeZoneResponseParams) Init() MTRTimeSynchronizationClusterSetTimeZoneResponseParams {
	rv := objc.Send[MTRTimeSynchronizationClusterSetTimeZoneResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTimeSynchronizationClusterSetTimeZoneResponseParams) Autorelease() MTRTimeSynchronizationClusterSetTimeZoneResponseParams {
	rv := objc.Send[MTRTimeSynchronizationClusterSetTimeZoneResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTimeSynchronizationClusterSetTimeZoneResponseParams creates a new MTRTimeSynchronizationClusterSetTimeZoneResponseParams instance.
func NewMTRTimeSynchronizationClusterSetTimeZoneResponseParams() MTRTimeSynchronizationClusterSetTimeZoneResponseParams {
	return getMTRTimeSynchronizationClusterSetTimeZoneResponseParamsClass().New()
}




// Initialize an MTRTimeSynchronizationClusterSetTimeZoneResponseParams with a response-value dictionary of the sort that MTRDeviceResponseHandler would receive.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterSetTimeZoneResponseParams/init(responseValue:)
func NewMTRTimeSynchronizationClusterSetTimeZoneResponseParamsWithResponseValueError(responseValue unsafe.Pointer, error_ unsafe.Pointer) MTRTimeSynchronizationClusterSetTimeZoneResponseParams {
	instance := getMTRTimeSynchronizationClusterSetTimeZoneResponseParamsClass().Alloc()
	rv := objc.Send[MTRTimeSynchronizationClusterSetTimeZoneResponseParams](instance.ID, objc.Sel("initWithResponseValue:error:"), responseValue, error_)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterSetTimeZoneResponseParams/dstOffsetRequired
func (m_ MTRTimeSynchronizationClusterSetTimeZoneResponseParams) DstOffsetRequired() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("dstOffsetRequired"))
	return rv
}


// SetDstOffsetRequired sets the value of the dstOffsetRequired property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterSetTimeZoneResponseParams/dstOffsetRequired
func (m_ MTRTimeSynchronizationClusterSetTimeZoneResponseParams) SetDstOffsetRequired(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDstOffsetRequired:"), value)
}


