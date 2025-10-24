// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRICDManagementClusterStayActiveRequestParams] class.
var (
	MTRICDManagementClusterStayActiveRequestParamsClass     _MTRICDManagementClusterStayActiveRequestParamsClass
	MTRICDManagementClusterStayActiveRequestParamsClassOnce sync.Once
)

func getMTRICDManagementClusterStayActiveRequestParamsClass() _MTRICDManagementClusterStayActiveRequestParamsClass {
	MTRICDManagementClusterStayActiveRequestParamsClassOnce.Do(func() {
		MTRICDManagementClusterStayActiveRequestParamsClass = _MTRICDManagementClusterStayActiveRequestParamsClass{objc.GetClass("MTRICDManagementClusterStayActiveRequestParams")}
	})
	return MTRICDManagementClusterStayActiveRequestParamsClass
}

type _MTRICDManagementClusterStayActiveRequestParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRICDManagementClusterStayActiveRequestParams] class.
type IMTRICDManagementClusterStayActiveRequestParams interface {
	objectivec.IObject
	// properties:
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	StayActiveDuration() objc.IObject /* cross-framework: NSNumber */
	SetStayActiveDuration(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterStayActiveRequestParams
type MTRICDManagementClusterStayActiveRequestParams struct {
	objectivec.Object
}

// MTRICDManagementClusterStayActiveRequestParamsFrom constructs a [MTRICDManagementClusterStayActiveRequestParams] from an unsafe.Pointer.
func MTRICDManagementClusterStayActiveRequestParamsFrom(ptr unsafe.Pointer) MTRICDManagementClusterStayActiveRequestParams {
	return MTRICDManagementClusterStayActiveRequestParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRICDManagementClusterStayActiveRequestParamsClass) Alloc() MTRICDManagementClusterStayActiveRequestParams {
	rv := objc.Send[MTRICDManagementClusterStayActiveRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRICDManagementClusterStayActiveRequestParamsClass) New() MTRICDManagementClusterStayActiveRequestParams {
	rv := objc.Send[MTRICDManagementClusterStayActiveRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRICDManagementClusterStayActiveRequestParams) Init() MTRICDManagementClusterStayActiveRequestParams {
	rv := objc.Send[MTRICDManagementClusterStayActiveRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRICDManagementClusterStayActiveRequestParams) Autorelease() MTRICDManagementClusterStayActiveRequestParams {
	rv := objc.Send[MTRICDManagementClusterStayActiveRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRICDManagementClusterStayActiveRequestParams creates a new MTRICDManagementClusterStayActiveRequestParams instance.
func NewMTRICDManagementClusterStayActiveRequestParams() MTRICDManagementClusterStayActiveRequestParams {
	return getMTRICDManagementClusterStayActiveRequestParamsClass().New()
}



// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterStayActiveRequestParams/serverSideProcessingTimeout
func (m_ MTRICDManagementClusterStayActiveRequestParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterStayActiveRequestParams/serverSideProcessingTimeout
func (m_ MTRICDManagementClusterStayActiveRequestParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterStayActiveRequestParams/stayActiveDuration
func (m_ MTRICDManagementClusterStayActiveRequestParams) StayActiveDuration() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("stayActiveDuration"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterStayActiveRequestParams/stayActiveDuration
func (m_ MTRICDManagementClusterStayActiveRequestParams) SetStayActiveDuration(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStayActiveDuration:"), value)
}


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterStayActiveRequestParams/timedInvokeTimeoutMs
func (m_ MTRICDManagementClusterStayActiveRequestParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRICDManagementClusterStayActiveRequestParams/timedInvokeTimeoutMs
func (m_ MTRICDManagementClusterStayActiveRequestParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



