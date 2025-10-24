// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRServiceAreaClusterSelectAreasResponseParams] class.
var (
	MTRServiceAreaClusterSelectAreasResponseParamsClass     _MTRServiceAreaClusterSelectAreasResponseParamsClass
	MTRServiceAreaClusterSelectAreasResponseParamsClassOnce sync.Once
)

func getMTRServiceAreaClusterSelectAreasResponseParamsClass() _MTRServiceAreaClusterSelectAreasResponseParamsClass {
	MTRServiceAreaClusterSelectAreasResponseParamsClassOnce.Do(func() {
		MTRServiceAreaClusterSelectAreasResponseParamsClass = _MTRServiceAreaClusterSelectAreasResponseParamsClass{objc.GetClass("MTRServiceAreaClusterSelectAreasResponseParams")}
	})
	return MTRServiceAreaClusterSelectAreasResponseParamsClass
}

type _MTRServiceAreaClusterSelectAreasResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRServiceAreaClusterSelectAreasResponseParams] class.
type IMTRServiceAreaClusterSelectAreasResponseParams interface {
	objectivec.IObject
	// properties:
	Status() objc.IObject /* cross-framework: NSNumber */
	SetStatus(value objc.IObject /* cross-framework: NSNumber */)
	StatusText() objc.IObject /* cross-framework: NSString */
	SetStatusText(value objc.IObject /* cross-framework: NSString */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterSelectAreasResponseParams
type MTRServiceAreaClusterSelectAreasResponseParams struct {
	objectivec.Object
}

// MTRServiceAreaClusterSelectAreasResponseParamsFrom constructs a [MTRServiceAreaClusterSelectAreasResponseParams] from an unsafe.Pointer.
func MTRServiceAreaClusterSelectAreasResponseParamsFrom(ptr unsafe.Pointer) MTRServiceAreaClusterSelectAreasResponseParams {
	return MTRServiceAreaClusterSelectAreasResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRServiceAreaClusterSelectAreasResponseParamsClass) Alloc() MTRServiceAreaClusterSelectAreasResponseParams {
	rv := objc.Send[MTRServiceAreaClusterSelectAreasResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRServiceAreaClusterSelectAreasResponseParamsClass) New() MTRServiceAreaClusterSelectAreasResponseParams {
	rv := objc.Send[MTRServiceAreaClusterSelectAreasResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRServiceAreaClusterSelectAreasResponseParams) Init() MTRServiceAreaClusterSelectAreasResponseParams {
	rv := objc.Send[MTRServiceAreaClusterSelectAreasResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRServiceAreaClusterSelectAreasResponseParams) Autorelease() MTRServiceAreaClusterSelectAreasResponseParams {
	rv := objc.Send[MTRServiceAreaClusterSelectAreasResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRServiceAreaClusterSelectAreasResponseParams creates a new MTRServiceAreaClusterSelectAreasResponseParams instance.
func NewMTRServiceAreaClusterSelectAreasResponseParams() MTRServiceAreaClusterSelectAreasResponseParams {
	return getMTRServiceAreaClusterSelectAreasResponseParamsClass().New()
}



// Initialize an MTRServiceAreaClusterSelectAreasResponseParams with a response-value dictionary of the sort that MTRDeviceResponseHandler would receive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterSelectAreasResponseParams/init(responseValue:)
func NewMTRServiceAreaClusterSelectAreasResponseParamsWithResponseValueError(responseValue foundation.IDictionary, error_ unsafe.Pointer) MTRServiceAreaClusterSelectAreasResponseParams {
	instance := getMTRServiceAreaClusterSelectAreasResponseParamsClass().Alloc()
	rv := objc.Send[MTRServiceAreaClusterSelectAreasResponseParams](instance.ID, objc.Sel("initWithResponseValue:error:"), responseValue, error_)
	rv.Autorelease()
	return rv
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterSelectAreasResponseParams/status
func (m_ MTRServiceAreaClusterSelectAreasResponseParams) Status() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("status"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterSelectAreasResponseParams/status
func (m_ MTRServiceAreaClusterSelectAreasResponseParams) SetStatus(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatus:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterSelectAreasResponseParams/statusText
func (m_ MTRServiceAreaClusterSelectAreasResponseParams) StatusText() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("statusText"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterSelectAreasResponseParams/statusText
func (m_ MTRServiceAreaClusterSelectAreasResponseParams) SetStatusText(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatusText:"), value)
}


