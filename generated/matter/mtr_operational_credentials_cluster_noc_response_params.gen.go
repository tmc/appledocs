// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTROperationalCredentialsClusterNOCResponseParams] class.
var (
	MTROperationalCredentialsClusterNOCResponseParamsClass     _MTROperationalCredentialsClusterNOCResponseParamsClass
	MTROperationalCredentialsClusterNOCResponseParamsClassOnce sync.Once
)

func getMTROperationalCredentialsClusterNOCResponseParamsClass() _MTROperationalCredentialsClusterNOCResponseParamsClass {
	MTROperationalCredentialsClusterNOCResponseParamsClassOnce.Do(func() {
		MTROperationalCredentialsClusterNOCResponseParamsClass = _MTROperationalCredentialsClusterNOCResponseParamsClass{objc.GetClass("MTROperationalCredentialsClusterNOCResponseParams")}
	})
	return MTROperationalCredentialsClusterNOCResponseParamsClass
}

type _MTROperationalCredentialsClusterNOCResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTROperationalCredentialsClusterNOCResponseParams] class.
type IMTROperationalCredentialsClusterNOCResponseParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterNOCResponseParams
type MTROperationalCredentialsClusterNOCResponseParams struct {
	objectivec.Object
}

// MTROperationalCredentialsClusterNOCResponseParamsFrom constructs a [MTROperationalCredentialsClusterNOCResponseParams] from an unsafe.Pointer.
func MTROperationalCredentialsClusterNOCResponseParamsFrom(ptr unsafe.Pointer) MTROperationalCredentialsClusterNOCResponseParams {
	return MTROperationalCredentialsClusterNOCResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROperationalCredentialsClusterNOCResponseParamsClass) Alloc() MTROperationalCredentialsClusterNOCResponseParams {
	rv := objc.Send[MTROperationalCredentialsClusterNOCResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROperationalCredentialsClusterNOCResponseParamsClass) New() MTROperationalCredentialsClusterNOCResponseParams {
	rv := objc.Send[MTROperationalCredentialsClusterNOCResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROperationalCredentialsClusterNOCResponseParams) Init() MTROperationalCredentialsClusterNOCResponseParams {
	rv := objc.Send[MTROperationalCredentialsClusterNOCResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROperationalCredentialsClusterNOCResponseParams) Autorelease() MTROperationalCredentialsClusterNOCResponseParams {
	rv := objc.Send[MTROperationalCredentialsClusterNOCResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROperationalCredentialsClusterNOCResponseParams creates a new MTROperationalCredentialsClusterNOCResponseParams instance.
func NewMTROperationalCredentialsClusterNOCResponseParams() MTROperationalCredentialsClusterNOCResponseParams {
	return getMTROperationalCredentialsClusterNOCResponseParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusternocresponseparams/debugtext
func (m_ MTROperationalCredentialsClusterNOCResponseParams) DebugText() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("debugText"))
	return rv
}


// SetDebugText sets the value of the debugText property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusternocresponseparams/debugtext
func (m_ MTROperationalCredentialsClusterNOCResponseParams) SetDebugText(value appkit.string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDebugText:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusternocresponseparams/fabricindex
func (m_ MTROperationalCredentialsClusterNOCResponseParams) FabricIndex() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("fabricIndex"))
	return rv
}


// SetFabricIndex sets the value of the fabricIndex property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusternocresponseparams/fabricindex
func (m_ MTROperationalCredentialsClusterNOCResponseParams) SetFabricIndex(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricIndex:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusternocresponseparams/statuscode
func (m_ MTROperationalCredentialsClusterNOCResponseParams) StatusCode() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("statusCode"))
	return rv
}


// SetStatusCode sets the value of the statusCode property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusternocresponseparams/statuscode
func (m_ MTROperationalCredentialsClusterNOCResponseParams) SetStatusCode(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatusCode:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusternocresponseparams/timedinvoketimeoutms
func (m_ MTROperationalCredentialsClusterNOCResponseParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusternocresponseparams/timedinvoketimeoutms
func (m_ MTROperationalCredentialsClusterNOCResponseParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



