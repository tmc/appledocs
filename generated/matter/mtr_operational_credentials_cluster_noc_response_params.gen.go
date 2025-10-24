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
	// properties:
	DebugText() objc.IObject /* cross-framework: NSString */
	SetDebugText(value objc.IObject /* cross-framework: NSString */)
	FabricIndex() objc.IObject /* cross-framework: NSNumber */
	SetFabricIndex(value objc.IObject /* cross-framework: NSNumber */)
	StatusCode() objc.IObject /* cross-framework: NSNumber */
	SetStatusCode(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusternocresponseparams/debugtext
func (m_ MTROperationalCredentialsClusterNOCResponseParams) DebugText() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("debugText"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusternocresponseparams/debugtext
func (m_ MTROperationalCredentialsClusterNOCResponseParams) SetDebugText(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDebugText:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusternocresponseparams/fabricindex
func (m_ MTROperationalCredentialsClusterNOCResponseParams) FabricIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("fabricIndex"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusternocresponseparams/fabricindex
func (m_ MTROperationalCredentialsClusterNOCResponseParams) SetFabricIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricIndex:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusternocresponseparams/statuscode
func (m_ MTROperationalCredentialsClusterNOCResponseParams) StatusCode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("statusCode"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusternocresponseparams/statuscode
func (m_ MTROperationalCredentialsClusterNOCResponseParams) SetStatusCode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatusCode:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusternocresponseparams/timedinvoketimeoutms
func (m_ MTROperationalCredentialsClusterNOCResponseParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtroperationalcredentialsclusternocresponseparams/timedinvoketimeoutms
func (m_ MTROperationalCredentialsClusterNOCResponseParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



