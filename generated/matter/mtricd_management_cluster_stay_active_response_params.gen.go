// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRICDManagementClusterStayActiveResponseParams] class.
var (
	MTRICDManagementClusterStayActiveResponseParamsClass     _MTRICDManagementClusterStayActiveResponseParamsClass
	MTRICDManagementClusterStayActiveResponseParamsClassOnce sync.Once
)

func getMTRICDManagementClusterStayActiveResponseParamsClass() _MTRICDManagementClusterStayActiveResponseParamsClass {
	MTRICDManagementClusterStayActiveResponseParamsClassOnce.Do(func() {
		MTRICDManagementClusterStayActiveResponseParamsClass = _MTRICDManagementClusterStayActiveResponseParamsClass{objc.GetClass("MTRICDManagementClusterStayActiveResponseParams")}
	})
	return MTRICDManagementClusterStayActiveResponseParamsClass
}

type _MTRICDManagementClusterStayActiveResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRICDManagementClusterStayActiveResponseParams] class.
type IMTRICDManagementClusterStayActiveResponseParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterStayActiveResponseParams
type MTRICDManagementClusterStayActiveResponseParams struct {
	objectivec.Object
}

// MTRICDManagementClusterStayActiveResponseParamsFrom constructs a [MTRICDManagementClusterStayActiveResponseParams] from an unsafe.Pointer.
func MTRICDManagementClusterStayActiveResponseParamsFrom(ptr unsafe.Pointer) MTRICDManagementClusterStayActiveResponseParams {
	return MTRICDManagementClusterStayActiveResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRICDManagementClusterStayActiveResponseParamsClass) Alloc() MTRICDManagementClusterStayActiveResponseParams {
	rv := objc.Send[MTRICDManagementClusterStayActiveResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRICDManagementClusterStayActiveResponseParamsClass) New() MTRICDManagementClusterStayActiveResponseParams {
	rv := objc.Send[MTRICDManagementClusterStayActiveResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRICDManagementClusterStayActiveResponseParams) Init() MTRICDManagementClusterStayActiveResponseParams {
	rv := objc.Send[MTRICDManagementClusterStayActiveResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRICDManagementClusterStayActiveResponseParams) Autorelease() MTRICDManagementClusterStayActiveResponseParams {
	rv := objc.Send[MTRICDManagementClusterStayActiveResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRICDManagementClusterStayActiveResponseParams creates a new MTRICDManagementClusterStayActiveResponseParams instance.
func NewMTRICDManagementClusterStayActiveResponseParams() MTRICDManagementClusterStayActiveResponseParams {
	return getMTRICDManagementClusterStayActiveResponseParamsClass().New()
}


// Initialize an MTRICDManagementClusterStayActiveResponseParams with a response-value dictionary of the sort that MTRDeviceResponseHandler would receive.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterStayActiveResponseParams/init(responseValue:)
func NewMTRICDManagementClusterStayActiveResponseParamsWithResponseValueError(responseValue unsafe.Pointer, error_ unsafe.Pointer) MTRICDManagementClusterStayActiveResponseParams {
	instance := getMTRICDManagementClusterStayActiveResponseParamsClass().Alloc()
	rv := objc.Send[MTRICDManagementClusterStayActiveResponseParams](instance.ID, objc.Sel("initWithResponseValue:error:"), responseValue, error_)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterStayActiveResponseParams/promisedActiveDuration
func (m_ MTRICDManagementClusterStayActiveResponseParams) PromisedActiveDuration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("promisedActiveDuration"))
	return rv
}


// SetPromisedActiveDuration sets the value of the promisedActiveDuration property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterStayActiveResponseParams/promisedActiveDuration
func (m_ MTRICDManagementClusterStayActiveResponseParams) SetPromisedActiveDuration(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPromisedActiveDuration:"), value)
}

