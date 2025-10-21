// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTROvenCavityOperationalStateClusterOperationalCommandResponseParams] class.
var (
	MTROvenCavityOperationalStateClusterOperationalCommandResponseParamsClass     _MTROvenCavityOperationalStateClusterOperationalCommandResponseParamsClass
	MTROvenCavityOperationalStateClusterOperationalCommandResponseParamsClassOnce sync.Once
)

func getMTROvenCavityOperationalStateClusterOperationalCommandResponseParamsClass() _MTROvenCavityOperationalStateClusterOperationalCommandResponseParamsClass {
	MTROvenCavityOperationalStateClusterOperationalCommandResponseParamsClassOnce.Do(func() {
		MTROvenCavityOperationalStateClusterOperationalCommandResponseParamsClass = _MTROvenCavityOperationalStateClusterOperationalCommandResponseParamsClass{objc.GetClass("MTROvenCavityOperationalStateClusterOperationalCommandResponseParams")}
	})
	return MTROvenCavityOperationalStateClusterOperationalCommandResponseParamsClass
}

type _MTROvenCavityOperationalStateClusterOperationalCommandResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTROvenCavityOperationalStateClusterOperationalCommandResponseParams] class.
type IMTROvenCavityOperationalStateClusterOperationalCommandResponseParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenCavityOperationalStateClusterOperationalCommandResponseParams
type MTROvenCavityOperationalStateClusterOperationalCommandResponseParams struct {
	objectivec.Object
}

// MTROvenCavityOperationalStateClusterOperationalCommandResponseParamsFrom constructs a [MTROvenCavityOperationalStateClusterOperationalCommandResponseParams] from an unsafe.Pointer.
func MTROvenCavityOperationalStateClusterOperationalCommandResponseParamsFrom(ptr unsafe.Pointer) MTROvenCavityOperationalStateClusterOperationalCommandResponseParams {
	return MTROvenCavityOperationalStateClusterOperationalCommandResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROvenCavityOperationalStateClusterOperationalCommandResponseParamsClass) Alloc() MTROvenCavityOperationalStateClusterOperationalCommandResponseParams {
	rv := objc.Send[MTROvenCavityOperationalStateClusterOperationalCommandResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROvenCavityOperationalStateClusterOperationalCommandResponseParamsClass) New() MTROvenCavityOperationalStateClusterOperationalCommandResponseParams {
	rv := objc.Send[MTROvenCavityOperationalStateClusterOperationalCommandResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROvenCavityOperationalStateClusterOperationalCommandResponseParams) Init() MTROvenCavityOperationalStateClusterOperationalCommandResponseParams {
	rv := objc.Send[MTROvenCavityOperationalStateClusterOperationalCommandResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROvenCavityOperationalStateClusterOperationalCommandResponseParams) Autorelease() MTROvenCavityOperationalStateClusterOperationalCommandResponseParams {
	rv := objc.Send[MTROvenCavityOperationalStateClusterOperationalCommandResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROvenCavityOperationalStateClusterOperationalCommandResponseParams creates a new MTROvenCavityOperationalStateClusterOperationalCommandResponseParams instance.
func NewMTROvenCavityOperationalStateClusterOperationalCommandResponseParams() MTROvenCavityOperationalStateClusterOperationalCommandResponseParams {
	return getMTROvenCavityOperationalStateClusterOperationalCommandResponseParamsClass().New()
}


// Initialize an MTROvenCavityOperationalStateClusterOperationalCommandResponseParams with a response-value dictionary of the sort that MTRDeviceResponseHandler would receive.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenCavityOperationalStateClusterOperationalCommandResponseParams/init(responseValue:)
func NewMTROvenCavityOperationalStateClusterOperationalCommandResponseParamsWithResponseValueError(responseValue unsafe.Pointer, error_ unsafe.Pointer) MTROvenCavityOperationalStateClusterOperationalCommandResponseParams {
	instance := getMTROvenCavityOperationalStateClusterOperationalCommandResponseParamsClass().Alloc()
	rv := objc.Send[MTROvenCavityOperationalStateClusterOperationalCommandResponseParams](instance.ID, objc.Sel("initWithResponseValue:error:"), responseValue, error_)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenCavityOperationalStateClusterOperationalCommandResponseParams/commandResponseState
func (m_ MTROvenCavityOperationalStateClusterOperationalCommandResponseParams) CommandResponseState() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("commandResponseState"))
	return rv
}


// SetCommandResponseState sets the value of the commandResponseState property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenCavityOperationalStateClusterOperationalCommandResponseParams/commandResponseState
func (m_ MTROvenCavityOperationalStateClusterOperationalCommandResponseParams) SetCommandResponseState(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCommandResponseState:"), value)
}

