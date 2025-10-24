// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRServiceAreaClusterSkipAreaParams] class.
var (
	MTRServiceAreaClusterSkipAreaParamsClass     _MTRServiceAreaClusterSkipAreaParamsClass
	MTRServiceAreaClusterSkipAreaParamsClassOnce sync.Once
)

func getMTRServiceAreaClusterSkipAreaParamsClass() _MTRServiceAreaClusterSkipAreaParamsClass {
	MTRServiceAreaClusterSkipAreaParamsClassOnce.Do(func() {
		MTRServiceAreaClusterSkipAreaParamsClass = _MTRServiceAreaClusterSkipAreaParamsClass{objc.GetClass("MTRServiceAreaClusterSkipAreaParams")}
	})
	return MTRServiceAreaClusterSkipAreaParamsClass
}

type _MTRServiceAreaClusterSkipAreaParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRServiceAreaClusterSkipAreaParams] class.
type IMTRServiceAreaClusterSkipAreaParams interface {
	objectivec.IObject
	// properties:
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	SkippedArea() objc.IObject /* cross-framework: NSNumber */
	SetSkippedArea(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterSkipAreaParams
type MTRServiceAreaClusterSkipAreaParams struct {
	objectivec.Object
}

// MTRServiceAreaClusterSkipAreaParamsFrom constructs a [MTRServiceAreaClusterSkipAreaParams] from an unsafe.Pointer.
func MTRServiceAreaClusterSkipAreaParamsFrom(ptr unsafe.Pointer) MTRServiceAreaClusterSkipAreaParams {
	return MTRServiceAreaClusterSkipAreaParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRServiceAreaClusterSkipAreaParamsClass) Alloc() MTRServiceAreaClusterSkipAreaParams {
	rv := objc.Send[MTRServiceAreaClusterSkipAreaParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRServiceAreaClusterSkipAreaParamsClass) New() MTRServiceAreaClusterSkipAreaParams {
	rv := objc.Send[MTRServiceAreaClusterSkipAreaParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRServiceAreaClusterSkipAreaParams) Init() MTRServiceAreaClusterSkipAreaParams {
	rv := objc.Send[MTRServiceAreaClusterSkipAreaParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRServiceAreaClusterSkipAreaParams) Autorelease() MTRServiceAreaClusterSkipAreaParams {
	rv := objc.Send[MTRServiceAreaClusterSkipAreaParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRServiceAreaClusterSkipAreaParams creates a new MTRServiceAreaClusterSkipAreaParams instance.
func NewMTRServiceAreaClusterSkipAreaParams() MTRServiceAreaClusterSkipAreaParams {
	return getMTRServiceAreaClusterSkipAreaParamsClass().New()
}



// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterSkipAreaParams/serverSideProcessingTimeout
func (m_ MTRServiceAreaClusterSkipAreaParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterSkipAreaParams/serverSideProcessingTimeout
func (m_ MTRServiceAreaClusterSkipAreaParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterSkipAreaParams/skippedArea
func (m_ MTRServiceAreaClusterSkipAreaParams) SkippedArea() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("skippedArea"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterSkipAreaParams/skippedArea
func (m_ MTRServiceAreaClusterSkipAreaParams) SetSkippedArea(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSkippedArea:"), value)
}


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterSkipAreaParams/timedInvokeTimeoutMs
func (m_ MTRServiceAreaClusterSkipAreaParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterSkipAreaParams/timedInvokeTimeoutMs
func (m_ MTRServiceAreaClusterSkipAreaParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



