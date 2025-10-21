// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRAccessControlClusterReviewFabricRestrictionsParams] class.
var (
	MTRAccessControlClusterReviewFabricRestrictionsParamsClass     _MTRAccessControlClusterReviewFabricRestrictionsParamsClass
	MTRAccessControlClusterReviewFabricRestrictionsParamsClassOnce sync.Once
)

func getMTRAccessControlClusterReviewFabricRestrictionsParamsClass() _MTRAccessControlClusterReviewFabricRestrictionsParamsClass {
	MTRAccessControlClusterReviewFabricRestrictionsParamsClassOnce.Do(func() {
		MTRAccessControlClusterReviewFabricRestrictionsParamsClass = _MTRAccessControlClusterReviewFabricRestrictionsParamsClass{objc.GetClass("MTRAccessControlClusterReviewFabricRestrictionsParams")}
	})
	return MTRAccessControlClusterReviewFabricRestrictionsParamsClass
}

type _MTRAccessControlClusterReviewFabricRestrictionsParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRAccessControlClusterReviewFabricRestrictionsParams] class.
type IMTRAccessControlClusterReviewFabricRestrictionsParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccessControlClusterReviewFabricRestrictionsParams
type MTRAccessControlClusterReviewFabricRestrictionsParams struct {
	objectivec.Object
}

// MTRAccessControlClusterReviewFabricRestrictionsParamsFrom constructs a [MTRAccessControlClusterReviewFabricRestrictionsParams] from an unsafe.Pointer.
func MTRAccessControlClusterReviewFabricRestrictionsParamsFrom(ptr unsafe.Pointer) MTRAccessControlClusterReviewFabricRestrictionsParams {
	return MTRAccessControlClusterReviewFabricRestrictionsParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRAccessControlClusterReviewFabricRestrictionsParamsClass) Alloc() MTRAccessControlClusterReviewFabricRestrictionsParams {
	rv := objc.Send[MTRAccessControlClusterReviewFabricRestrictionsParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRAccessControlClusterReviewFabricRestrictionsParamsClass) New() MTRAccessControlClusterReviewFabricRestrictionsParams {
	rv := objc.Send[MTRAccessControlClusterReviewFabricRestrictionsParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRAccessControlClusterReviewFabricRestrictionsParams) Init() MTRAccessControlClusterReviewFabricRestrictionsParams {
	rv := objc.Send[MTRAccessControlClusterReviewFabricRestrictionsParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRAccessControlClusterReviewFabricRestrictionsParams) Autorelease() MTRAccessControlClusterReviewFabricRestrictionsParams {
	rv := objc.Send[MTRAccessControlClusterReviewFabricRestrictionsParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRAccessControlClusterReviewFabricRestrictionsParams creates a new MTRAccessControlClusterReviewFabricRestrictionsParams instance.
func NewMTRAccessControlClusterReviewFabricRestrictionsParams() MTRAccessControlClusterReviewFabricRestrictionsParams {
	return getMTRAccessControlClusterReviewFabricRestrictionsParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccessControlClusterReviewFabricRestrictionsParams/arl
func (m_ MTRAccessControlClusterReviewFabricRestrictionsParams) Arl() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("arl"))
	return rv
}


// SetArl sets the value of the arl property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccessControlClusterReviewFabricRestrictionsParams/arl
func (m_ MTRAccessControlClusterReviewFabricRestrictionsParams) SetArl(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setArl:"), value)
}

// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccessControlClusterReviewFabricRestrictionsParams/serverSideProcessingTimeout
func (m_ MTRAccessControlClusterReviewFabricRestrictionsParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
// Controls how much time, in seconds, we will allow for the server to process the command.

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccessControlClusterReviewFabricRestrictionsParams/serverSideProcessingTimeout
func (m_ MTRAccessControlClusterReviewFabricRestrictionsParams) SetServerSideProcessingTimeout(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccessControlClusterReviewFabricRestrictionsParams/timedInvokeTimeoutMs
func (m_ MTRAccessControlClusterReviewFabricRestrictionsParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
// Controls whether the command is a timed command (using Timed Invoke).

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccessControlClusterReviewFabricRestrictionsParams/timedInvokeTimeoutMs
func (m_ MTRAccessControlClusterReviewFabricRestrictionsParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



