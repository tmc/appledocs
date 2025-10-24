// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRTimeSynchronizationClusterSetDefaultNTPParams] class.
var (
	MTRTimeSynchronizationClusterSetDefaultNTPParamsClass     _MTRTimeSynchronizationClusterSetDefaultNTPParamsClass
	MTRTimeSynchronizationClusterSetDefaultNTPParamsClassOnce sync.Once
)

func getMTRTimeSynchronizationClusterSetDefaultNTPParamsClass() _MTRTimeSynchronizationClusterSetDefaultNTPParamsClass {
	MTRTimeSynchronizationClusterSetDefaultNTPParamsClassOnce.Do(func() {
		MTRTimeSynchronizationClusterSetDefaultNTPParamsClass = _MTRTimeSynchronizationClusterSetDefaultNTPParamsClass{objc.GetClass("MTRTimeSynchronizationClusterSetDefaultNTPParams")}
	})
	return MTRTimeSynchronizationClusterSetDefaultNTPParamsClass
}

type _MTRTimeSynchronizationClusterSetDefaultNTPParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRTimeSynchronizationClusterSetDefaultNTPParams] class.
type IMTRTimeSynchronizationClusterSetDefaultNTPParams interface {
	objectivec.IObject
	// properties:
	DefaultNTP() objc.IObject /* cross-framework: NSString */
	SetDefaultNTP(value objc.IObject /* cross-framework: NSString */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterSetDefaultNTPParams
type MTRTimeSynchronizationClusterSetDefaultNTPParams struct {
	objectivec.Object
}

// MTRTimeSynchronizationClusterSetDefaultNTPParamsFrom constructs a [MTRTimeSynchronizationClusterSetDefaultNTPParams] from an unsafe.Pointer.
func MTRTimeSynchronizationClusterSetDefaultNTPParamsFrom(ptr unsafe.Pointer) MTRTimeSynchronizationClusterSetDefaultNTPParams {
	return MTRTimeSynchronizationClusterSetDefaultNTPParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTimeSynchronizationClusterSetDefaultNTPParamsClass) Alloc() MTRTimeSynchronizationClusterSetDefaultNTPParams {
	rv := objc.Send[MTRTimeSynchronizationClusterSetDefaultNTPParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTimeSynchronizationClusterSetDefaultNTPParamsClass) New() MTRTimeSynchronizationClusterSetDefaultNTPParams {
	rv := objc.Send[MTRTimeSynchronizationClusterSetDefaultNTPParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTimeSynchronizationClusterSetDefaultNTPParams) Init() MTRTimeSynchronizationClusterSetDefaultNTPParams {
	rv := objc.Send[MTRTimeSynchronizationClusterSetDefaultNTPParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTimeSynchronizationClusterSetDefaultNTPParams) Autorelease() MTRTimeSynchronizationClusterSetDefaultNTPParams {
	rv := objc.Send[MTRTimeSynchronizationClusterSetDefaultNTPParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTimeSynchronizationClusterSetDefaultNTPParams creates a new MTRTimeSynchronizationClusterSetDefaultNTPParams instance.
func NewMTRTimeSynchronizationClusterSetDefaultNTPParams() MTRTimeSynchronizationClusterSetDefaultNTPParams {
	return getMTRTimeSynchronizationClusterSetDefaultNTPParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterSetDefaultNTPParams/defaultNTP
func (m_ MTRTimeSynchronizationClusterSetDefaultNTPParams) DefaultNTP() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("defaultNTP"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterSetDefaultNTPParams/defaultNTP
func (m_ MTRTimeSynchronizationClusterSetDefaultNTPParams) SetDefaultNTP(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDefaultNTP:"), value)
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterSetDefaultNTPParams/serverSideProcessingTimeout
func (m_ MTRTimeSynchronizationClusterSetDefaultNTPParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterSetDefaultNTPParams/serverSideProcessingTimeout
func (m_ MTRTimeSynchronizationClusterSetDefaultNTPParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterSetDefaultNTPParams/timedInvokeTimeoutMs
func (m_ MTRTimeSynchronizationClusterSetDefaultNTPParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterSetDefaultNTPParams/timedInvokeTimeoutMs
func (m_ MTRTimeSynchronizationClusterSetDefaultNTPParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



