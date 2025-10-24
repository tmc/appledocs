// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTREnergyEVSEModeClusterChangeToModeResponseParams] class.
var (
	MTREnergyEVSEModeClusterChangeToModeResponseParamsClass     _MTREnergyEVSEModeClusterChangeToModeResponseParamsClass
	MTREnergyEVSEModeClusterChangeToModeResponseParamsClassOnce sync.Once
)

func getMTREnergyEVSEModeClusterChangeToModeResponseParamsClass() _MTREnergyEVSEModeClusterChangeToModeResponseParamsClass {
	MTREnergyEVSEModeClusterChangeToModeResponseParamsClassOnce.Do(func() {
		MTREnergyEVSEModeClusterChangeToModeResponseParamsClass = _MTREnergyEVSEModeClusterChangeToModeResponseParamsClass{objc.GetClass("MTREnergyEVSEModeClusterChangeToModeResponseParams")}
	})
	return MTREnergyEVSEModeClusterChangeToModeResponseParamsClass
}

type _MTREnergyEVSEModeClusterChangeToModeResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTREnergyEVSEModeClusterChangeToModeResponseParams] class.
type IMTREnergyEVSEModeClusterChangeToModeResponseParams interface {
	objectivec.IObject
	// properties:
	Status() objc.IObject /* cross-framework: NSNumber */
	SetStatus(value objc.IObject /* cross-framework: NSNumber */)
	StatusText() objc.IObject /* cross-framework: NSString */
	SetStatusText(value objc.IObject /* cross-framework: NSString */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEModeClusterChangeToModeResponseParams
type MTREnergyEVSEModeClusterChangeToModeResponseParams struct {
	objectivec.Object
}

// MTREnergyEVSEModeClusterChangeToModeResponseParamsFrom constructs a [MTREnergyEVSEModeClusterChangeToModeResponseParams] from an unsafe.Pointer.
func MTREnergyEVSEModeClusterChangeToModeResponseParamsFrom(ptr unsafe.Pointer) MTREnergyEVSEModeClusterChangeToModeResponseParams {
	return MTREnergyEVSEModeClusterChangeToModeResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTREnergyEVSEModeClusterChangeToModeResponseParamsClass) Alloc() MTREnergyEVSEModeClusterChangeToModeResponseParams {
	rv := objc.Send[MTREnergyEVSEModeClusterChangeToModeResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTREnergyEVSEModeClusterChangeToModeResponseParamsClass) New() MTREnergyEVSEModeClusterChangeToModeResponseParams {
	rv := objc.Send[MTREnergyEVSEModeClusterChangeToModeResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTREnergyEVSEModeClusterChangeToModeResponseParams) Init() MTREnergyEVSEModeClusterChangeToModeResponseParams {
	rv := objc.Send[MTREnergyEVSEModeClusterChangeToModeResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTREnergyEVSEModeClusterChangeToModeResponseParams) Autorelease() MTREnergyEVSEModeClusterChangeToModeResponseParams {
	rv := objc.Send[MTREnergyEVSEModeClusterChangeToModeResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTREnergyEVSEModeClusterChangeToModeResponseParams creates a new MTREnergyEVSEModeClusterChangeToModeResponseParams instance.
func NewMTREnergyEVSEModeClusterChangeToModeResponseParams() MTREnergyEVSEModeClusterChangeToModeResponseParams {
	return getMTREnergyEVSEModeClusterChangeToModeResponseParamsClass().New()
}



// Initialize an MTREnergyEVSEModeClusterChangeToModeResponseParams with a response-value dictionary of the sort that MTRDeviceResponseHandler would receive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEModeClusterChangeToModeResponseParams/init(responseValue:)
func NewMTREnergyEVSEModeClusterChangeToModeResponseParamsWithResponseValueError(responseValue foundation.IDictionary, error_ unsafe.Pointer) MTREnergyEVSEModeClusterChangeToModeResponseParams {
	instance := getMTREnergyEVSEModeClusterChangeToModeResponseParamsClass().Alloc()
	rv := objc.Send[MTREnergyEVSEModeClusterChangeToModeResponseParams](instance.ID, objc.Sel("initWithResponseValue:error:"), responseValue, error_)
	rv.Autorelease()
	return rv
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEModeClusterChangeToModeResponseParams/status
func (m_ MTREnergyEVSEModeClusterChangeToModeResponseParams) Status() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("status"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEModeClusterChangeToModeResponseParams/status
func (m_ MTREnergyEVSEModeClusterChangeToModeResponseParams) SetStatus(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatus:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEModeClusterChangeToModeResponseParams/statusText
func (m_ MTREnergyEVSEModeClusterChangeToModeResponseParams) StatusText() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("statusText"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEModeClusterChangeToModeResponseParams/statusText
func (m_ MTREnergyEVSEModeClusterChangeToModeResponseParams) SetStatusText(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatusText:"), value)
}


