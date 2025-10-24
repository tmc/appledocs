// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRICDManagementClusterRegisterClientResponseParams] class.
var (
	MTRICDManagementClusterRegisterClientResponseParamsClass     _MTRICDManagementClusterRegisterClientResponseParamsClass
	MTRICDManagementClusterRegisterClientResponseParamsClassOnce sync.Once
)

func getMTRICDManagementClusterRegisterClientResponseParamsClass() _MTRICDManagementClusterRegisterClientResponseParamsClass {
	MTRICDManagementClusterRegisterClientResponseParamsClassOnce.Do(func() {
		MTRICDManagementClusterRegisterClientResponseParamsClass = _MTRICDManagementClusterRegisterClientResponseParamsClass{objc.GetClass("MTRICDManagementClusterRegisterClientResponseParams")}
	})
	return MTRICDManagementClusterRegisterClientResponseParamsClass
}

type _MTRICDManagementClusterRegisterClientResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRICDManagementClusterRegisterClientResponseParams] class.
type IMTRICDManagementClusterRegisterClientResponseParams interface {
	objectivec.IObject
	// properties:
	IcdCounter() objc.IObject /* cross-framework: NSNumber */
	SetIcdCounter(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterRegisterClientResponseParams
type MTRICDManagementClusterRegisterClientResponseParams struct {
	objectivec.Object
}

// MTRICDManagementClusterRegisterClientResponseParamsFrom constructs a [MTRICDManagementClusterRegisterClientResponseParams] from an unsafe.Pointer.
func MTRICDManagementClusterRegisterClientResponseParamsFrom(ptr unsafe.Pointer) MTRICDManagementClusterRegisterClientResponseParams {
	return MTRICDManagementClusterRegisterClientResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRICDManagementClusterRegisterClientResponseParamsClass) Alloc() MTRICDManagementClusterRegisterClientResponseParams {
	rv := objc.Send[MTRICDManagementClusterRegisterClientResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRICDManagementClusterRegisterClientResponseParamsClass) New() MTRICDManagementClusterRegisterClientResponseParams {
	rv := objc.Send[MTRICDManagementClusterRegisterClientResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRICDManagementClusterRegisterClientResponseParams) Init() MTRICDManagementClusterRegisterClientResponseParams {
	rv := objc.Send[MTRICDManagementClusterRegisterClientResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRICDManagementClusterRegisterClientResponseParams) Autorelease() MTRICDManagementClusterRegisterClientResponseParams {
	rv := objc.Send[MTRICDManagementClusterRegisterClientResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRICDManagementClusterRegisterClientResponseParams creates a new MTRICDManagementClusterRegisterClientResponseParams instance.
func NewMTRICDManagementClusterRegisterClientResponseParams() MTRICDManagementClusterRegisterClientResponseParams {
	return getMTRICDManagementClusterRegisterClientResponseParamsClass().New()
}



// Initialize an MTRICDManagementClusterRegisterClientResponseParams with a response-value dictionary of the sort that MTRDeviceResponseHandler would receive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterRegisterClientResponseParams/init(responseValue:)
func NewMTRICDManagementClusterRegisterClientResponseParamsWithResponseValueError(responseValue foundation.IDictionary, error_ unsafe.Pointer) MTRICDManagementClusterRegisterClientResponseParams {
	instance := getMTRICDManagementClusterRegisterClientResponseParamsClass().Alloc()
	rv := objc.Send[MTRICDManagementClusterRegisterClientResponseParams](instance.ID, objc.Sel("initWithResponseValue:error:"), responseValue, error_)
	rv.Autorelease()
	return rv
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterRegisterClientResponseParams/icdCounter
func (m_ MTRICDManagementClusterRegisterClientResponseParams) IcdCounter() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("icdCounter"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterRegisterClientResponseParams/icdCounter
func (m_ MTRICDManagementClusterRegisterClientResponseParams) SetIcdCounter(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIcdCounter:"), value)
}


