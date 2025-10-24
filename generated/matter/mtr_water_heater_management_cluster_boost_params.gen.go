// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRWaterHeaterManagementClusterBoostParams] class.
var (
	MTRWaterHeaterManagementClusterBoostParamsClass     _MTRWaterHeaterManagementClusterBoostParamsClass
	MTRWaterHeaterManagementClusterBoostParamsClassOnce sync.Once
)

func getMTRWaterHeaterManagementClusterBoostParamsClass() _MTRWaterHeaterManagementClusterBoostParamsClass {
	MTRWaterHeaterManagementClusterBoostParamsClassOnce.Do(func() {
		MTRWaterHeaterManagementClusterBoostParamsClass = _MTRWaterHeaterManagementClusterBoostParamsClass{objc.GetClass("MTRWaterHeaterManagementClusterBoostParams")}
	})
	return MTRWaterHeaterManagementClusterBoostParamsClass
}

type _MTRWaterHeaterManagementClusterBoostParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRWaterHeaterManagementClusterBoostParams] class.
type IMTRWaterHeaterManagementClusterBoostParams interface {
	objectivec.IObject
	// properties:
	BoostInfo() IMTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct
	SetBoostInfo(value IMTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterManagementClusterBoostParams
type MTRWaterHeaterManagementClusterBoostParams struct {
	objectivec.Object
}

// MTRWaterHeaterManagementClusterBoostParamsFrom constructs a [MTRWaterHeaterManagementClusterBoostParams] from an unsafe.Pointer.
func MTRWaterHeaterManagementClusterBoostParamsFrom(ptr unsafe.Pointer) MTRWaterHeaterManagementClusterBoostParams {
	return MTRWaterHeaterManagementClusterBoostParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRWaterHeaterManagementClusterBoostParamsClass) Alloc() MTRWaterHeaterManagementClusterBoostParams {
	rv := objc.Send[MTRWaterHeaterManagementClusterBoostParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRWaterHeaterManagementClusterBoostParamsClass) New() MTRWaterHeaterManagementClusterBoostParams {
	rv := objc.Send[MTRWaterHeaterManagementClusterBoostParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRWaterHeaterManagementClusterBoostParams) Init() MTRWaterHeaterManagementClusterBoostParams {
	rv := objc.Send[MTRWaterHeaterManagementClusterBoostParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRWaterHeaterManagementClusterBoostParams) Autorelease() MTRWaterHeaterManagementClusterBoostParams {
	rv := objc.Send[MTRWaterHeaterManagementClusterBoostParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRWaterHeaterManagementClusterBoostParams creates a new MTRWaterHeaterManagementClusterBoostParams instance.
func NewMTRWaterHeaterManagementClusterBoostParams() MTRWaterHeaterManagementClusterBoostParams {
	return getMTRWaterHeaterManagementClusterBoostParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterManagementClusterBoostParams/boostInfo
func (m_ MTRWaterHeaterManagementClusterBoostParams) BoostInfo() IMTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct {
	rv := objc.Send[MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct](m_.ID, objc.Sel("boostInfo"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterManagementClusterBoostParams/boostInfo
func (m_ MTRWaterHeaterManagementClusterBoostParams) SetBoostInfo(value IMTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBoostInfo:"), value)
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterManagementClusterBoostParams/serverSideProcessingTimeout
func (m_ MTRWaterHeaterManagementClusterBoostParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterManagementClusterBoostParams/serverSideProcessingTimeout
func (m_ MTRWaterHeaterManagementClusterBoostParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterManagementClusterBoostParams/timedInvokeTimeoutMs
func (m_ MTRWaterHeaterManagementClusterBoostParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterManagementClusterBoostParams/timedInvokeTimeoutMs
func (m_ MTRWaterHeaterManagementClusterBoostParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



