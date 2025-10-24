// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRAccessControlClusterReviewFabricRestrictionsResponseParams] class.
var (
	MTRAccessControlClusterReviewFabricRestrictionsResponseParamsClass     _MTRAccessControlClusterReviewFabricRestrictionsResponseParamsClass
	MTRAccessControlClusterReviewFabricRestrictionsResponseParamsClassOnce sync.Once
)

func getMTRAccessControlClusterReviewFabricRestrictionsResponseParamsClass() _MTRAccessControlClusterReviewFabricRestrictionsResponseParamsClass {
	MTRAccessControlClusterReviewFabricRestrictionsResponseParamsClassOnce.Do(func() {
		MTRAccessControlClusterReviewFabricRestrictionsResponseParamsClass = _MTRAccessControlClusterReviewFabricRestrictionsResponseParamsClass{objc.GetClass("MTRAccessControlClusterReviewFabricRestrictionsResponseParams")}
	})
	return MTRAccessControlClusterReviewFabricRestrictionsResponseParamsClass
}

type _MTRAccessControlClusterReviewFabricRestrictionsResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRAccessControlClusterReviewFabricRestrictionsResponseParams] class.
type IMTRAccessControlClusterReviewFabricRestrictionsResponseParams interface {
	objectivec.IObject
	// properties:
	Token() objc.IObject /* cross-framework: NSNumber */
	SetToken(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccessControlClusterReviewFabricRestrictionsResponseParams
type MTRAccessControlClusterReviewFabricRestrictionsResponseParams struct {
	objectivec.Object
}

// MTRAccessControlClusterReviewFabricRestrictionsResponseParamsFrom constructs a [MTRAccessControlClusterReviewFabricRestrictionsResponseParams] from an unsafe.Pointer.
func MTRAccessControlClusterReviewFabricRestrictionsResponseParamsFrom(ptr unsafe.Pointer) MTRAccessControlClusterReviewFabricRestrictionsResponseParams {
	return MTRAccessControlClusterReviewFabricRestrictionsResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRAccessControlClusterReviewFabricRestrictionsResponseParamsClass) Alloc() MTRAccessControlClusterReviewFabricRestrictionsResponseParams {
	rv := objc.Send[MTRAccessControlClusterReviewFabricRestrictionsResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRAccessControlClusterReviewFabricRestrictionsResponseParamsClass) New() MTRAccessControlClusterReviewFabricRestrictionsResponseParams {
	rv := objc.Send[MTRAccessControlClusterReviewFabricRestrictionsResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRAccessControlClusterReviewFabricRestrictionsResponseParams) Init() MTRAccessControlClusterReviewFabricRestrictionsResponseParams {
	rv := objc.Send[MTRAccessControlClusterReviewFabricRestrictionsResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRAccessControlClusterReviewFabricRestrictionsResponseParams) Autorelease() MTRAccessControlClusterReviewFabricRestrictionsResponseParams {
	rv := objc.Send[MTRAccessControlClusterReviewFabricRestrictionsResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRAccessControlClusterReviewFabricRestrictionsResponseParams creates a new MTRAccessControlClusterReviewFabricRestrictionsResponseParams instance.
func NewMTRAccessControlClusterReviewFabricRestrictionsResponseParams() MTRAccessControlClusterReviewFabricRestrictionsResponseParams {
	return getMTRAccessControlClusterReviewFabricRestrictionsResponseParamsClass().New()
}



// Initialize an MTRAccessControlClusterReviewFabricRestrictionsResponseParams with a response-value dictionary of the sort that MTRDeviceResponseHandler would receive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccessControlClusterReviewFabricRestrictionsResponseParams/init(responseValue:)
func NewMTRAccessControlClusterReviewFabricRestrictionsResponseParamsWithResponseValueError(responseValue foundation.IDictionary, error_ unsafe.Pointer) MTRAccessControlClusterReviewFabricRestrictionsResponseParams {
	instance := getMTRAccessControlClusterReviewFabricRestrictionsResponseParamsClass().Alloc()
	rv := objc.Send[MTRAccessControlClusterReviewFabricRestrictionsResponseParams](instance.ID, objc.Sel("initWithResponseValue:error:"), responseValue, error_)
	rv.Autorelease()
	return rv
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccessControlClusterReviewFabricRestrictionsResponseParams/token
func (m_ MTRAccessControlClusterReviewFabricRestrictionsResponseParams) Token() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("token"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccessControlClusterReviewFabricRestrictionsResponseParams/token
func (m_ MTRAccessControlClusterReviewFabricRestrictionsResponseParams) SetToken(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setToken:"), value)
}


