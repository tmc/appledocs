// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRContentAppObserverClusterContentAppMessageResponseParams] class.
var (
	MTRContentAppObserverClusterContentAppMessageResponseParamsClass     _MTRContentAppObserverClusterContentAppMessageResponseParamsClass
	MTRContentAppObserverClusterContentAppMessageResponseParamsClassOnce sync.Once
)

func getMTRContentAppObserverClusterContentAppMessageResponseParamsClass() _MTRContentAppObserverClusterContentAppMessageResponseParamsClass {
	MTRContentAppObserverClusterContentAppMessageResponseParamsClassOnce.Do(func() {
		MTRContentAppObserverClusterContentAppMessageResponseParamsClass = _MTRContentAppObserverClusterContentAppMessageResponseParamsClass{objc.GetClass("MTRContentAppObserverClusterContentAppMessageResponseParams")}
	})
	return MTRContentAppObserverClusterContentAppMessageResponseParamsClass
}

type _MTRContentAppObserverClusterContentAppMessageResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRContentAppObserverClusterContentAppMessageResponseParams] class.
type IMTRContentAppObserverClusterContentAppMessageResponseParams interface {
	objectivec.IObject
	// properties:
	Data() objc.IObject /* cross-framework: NSString */
	SetData(value objc.IObject /* cross-framework: NSString */)
	EncodingHint() objc.IObject /* cross-framework: NSString */
	SetEncodingHint(value objc.IObject /* cross-framework: NSString */)
	Status() objc.IObject /* cross-framework: NSNumber */
	SetStatus(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentAppObserverClusterContentAppMessageResponseParams
type MTRContentAppObserverClusterContentAppMessageResponseParams struct {
	objectivec.Object
}

// MTRContentAppObserverClusterContentAppMessageResponseParamsFrom constructs a [MTRContentAppObserverClusterContentAppMessageResponseParams] from an unsafe.Pointer.
func MTRContentAppObserverClusterContentAppMessageResponseParamsFrom(ptr unsafe.Pointer) MTRContentAppObserverClusterContentAppMessageResponseParams {
	return MTRContentAppObserverClusterContentAppMessageResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRContentAppObserverClusterContentAppMessageResponseParamsClass) Alloc() MTRContentAppObserverClusterContentAppMessageResponseParams {
	rv := objc.Send[MTRContentAppObserverClusterContentAppMessageResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRContentAppObserverClusterContentAppMessageResponseParamsClass) New() MTRContentAppObserverClusterContentAppMessageResponseParams {
	rv := objc.Send[MTRContentAppObserverClusterContentAppMessageResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRContentAppObserverClusterContentAppMessageResponseParams) Init() MTRContentAppObserverClusterContentAppMessageResponseParams {
	rv := objc.Send[MTRContentAppObserverClusterContentAppMessageResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRContentAppObserverClusterContentAppMessageResponseParams) Autorelease() MTRContentAppObserverClusterContentAppMessageResponseParams {
	rv := objc.Send[MTRContentAppObserverClusterContentAppMessageResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRContentAppObserverClusterContentAppMessageResponseParams creates a new MTRContentAppObserverClusterContentAppMessageResponseParams instance.
func NewMTRContentAppObserverClusterContentAppMessageResponseParams() MTRContentAppObserverClusterContentAppMessageResponseParams {
	return getMTRContentAppObserverClusterContentAppMessageResponseParamsClass().New()
}



// Initialize an MTRContentAppObserverClusterContentAppMessageResponseParams with a response-value dictionary of the sort that MTRDeviceResponseHandler would receive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentAppObserverClusterContentAppMessageResponseParams/init(responseValue:)
func NewMTRContentAppObserverClusterContentAppMessageResponseParamsWithResponseValueError(responseValue foundation.IDictionary, error_ unsafe.Pointer) MTRContentAppObserverClusterContentAppMessageResponseParams {
	instance := getMTRContentAppObserverClusterContentAppMessageResponseParamsClass().Alloc()
	rv := objc.Send[MTRContentAppObserverClusterContentAppMessageResponseParams](instance.ID, objc.Sel("initWithResponseValue:error:"), responseValue, error_)
	rv.Autorelease()
	return rv
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentAppObserverClusterContentAppMessageResponseParams/data
func (m_ MTRContentAppObserverClusterContentAppMessageResponseParams) Data() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("data"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentAppObserverClusterContentAppMessageResponseParams/data
func (m_ MTRContentAppObserverClusterContentAppMessageResponseParams) SetData(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setData:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentAppObserverClusterContentAppMessageResponseParams/encodingHint
func (m_ MTRContentAppObserverClusterContentAppMessageResponseParams) EncodingHint() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("encodingHint"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentAppObserverClusterContentAppMessageResponseParams/encodingHint
func (m_ MTRContentAppObserverClusterContentAppMessageResponseParams) SetEncodingHint(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEncodingHint:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentAppObserverClusterContentAppMessageResponseParams/status
func (m_ MTRContentAppObserverClusterContentAppMessageResponseParams) Status() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("status"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRContentAppObserverClusterContentAppMessageResponseParams/status
func (m_ MTRContentAppObserverClusterContentAppMessageResponseParams) SetStatus(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatus:"), value)
}


