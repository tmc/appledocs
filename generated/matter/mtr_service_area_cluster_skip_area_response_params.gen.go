// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRServiceAreaClusterSkipAreaResponseParams] class.
var (
	MTRServiceAreaClusterSkipAreaResponseParamsClass     _MTRServiceAreaClusterSkipAreaResponseParamsClass
	MTRServiceAreaClusterSkipAreaResponseParamsClassOnce sync.Once
)

func getMTRServiceAreaClusterSkipAreaResponseParamsClass() _MTRServiceAreaClusterSkipAreaResponseParamsClass {
	MTRServiceAreaClusterSkipAreaResponseParamsClassOnce.Do(func() {
		MTRServiceAreaClusterSkipAreaResponseParamsClass = _MTRServiceAreaClusterSkipAreaResponseParamsClass{objc.GetClass("MTRServiceAreaClusterSkipAreaResponseParams")}
	})
	return MTRServiceAreaClusterSkipAreaResponseParamsClass
}

type _MTRServiceAreaClusterSkipAreaResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRServiceAreaClusterSkipAreaResponseParams] class.
type IMTRServiceAreaClusterSkipAreaResponseParams interface {
	objectivec.IObject
	// properties:
	Status() objc.IObject /* cross-framework: NSNumber */
	SetStatus(value objc.IObject /* cross-framework: NSNumber */)
	StatusText() objc.IObject /* cross-framework: NSString */
	SetStatusText(value objc.IObject /* cross-framework: NSString */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterSkipAreaResponseParams
type MTRServiceAreaClusterSkipAreaResponseParams struct {
	objectivec.Object
}

// MTRServiceAreaClusterSkipAreaResponseParamsFrom constructs a [MTRServiceAreaClusterSkipAreaResponseParams] from an unsafe.Pointer.
func MTRServiceAreaClusterSkipAreaResponseParamsFrom(ptr unsafe.Pointer) MTRServiceAreaClusterSkipAreaResponseParams {
	return MTRServiceAreaClusterSkipAreaResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRServiceAreaClusterSkipAreaResponseParamsClass) Alloc() MTRServiceAreaClusterSkipAreaResponseParams {
	rv := objc.Send[MTRServiceAreaClusterSkipAreaResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRServiceAreaClusterSkipAreaResponseParamsClass) New() MTRServiceAreaClusterSkipAreaResponseParams {
	rv := objc.Send[MTRServiceAreaClusterSkipAreaResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRServiceAreaClusterSkipAreaResponseParams) Init() MTRServiceAreaClusterSkipAreaResponseParams {
	rv := objc.Send[MTRServiceAreaClusterSkipAreaResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRServiceAreaClusterSkipAreaResponseParams) Autorelease() MTRServiceAreaClusterSkipAreaResponseParams {
	rv := objc.Send[MTRServiceAreaClusterSkipAreaResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRServiceAreaClusterSkipAreaResponseParams creates a new MTRServiceAreaClusterSkipAreaResponseParams instance.
func NewMTRServiceAreaClusterSkipAreaResponseParams() MTRServiceAreaClusterSkipAreaResponseParams {
	return getMTRServiceAreaClusterSkipAreaResponseParamsClass().New()
}



// Initialize an MTRServiceAreaClusterSkipAreaResponseParams with a response-value dictionary of the sort that MTRDeviceResponseHandler would receive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterSkipAreaResponseParams/init(responseValue:)
func NewMTRServiceAreaClusterSkipAreaResponseParamsWithResponseValueError(responseValue foundation.IDictionary, error_ unsafe.Pointer) MTRServiceAreaClusterSkipAreaResponseParams {
	instance := getMTRServiceAreaClusterSkipAreaResponseParamsClass().Alloc()
	rv := objc.Send[MTRServiceAreaClusterSkipAreaResponseParams](instance.ID, objc.Sel("initWithResponseValue:error:"), responseValue, error_)
	rv.Autorelease()
	return rv
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterSkipAreaResponseParams/status
func (m_ MTRServiceAreaClusterSkipAreaResponseParams) Status() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("status"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterSkipAreaResponseParams/status
func (m_ MTRServiceAreaClusterSkipAreaResponseParams) SetStatus(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatus:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterSkipAreaResponseParams/statusText
func (m_ MTRServiceAreaClusterSkipAreaResponseParams) StatusText() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("statusText"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterSkipAreaResponseParams/statusText
func (m_ MTRServiceAreaClusterSkipAreaResponseParams) SetStatusText(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatusText:"), value)
}


