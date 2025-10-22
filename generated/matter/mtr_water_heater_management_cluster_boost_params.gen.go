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
	BoostInfo() MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct
	SetBoostInfo(value IMTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct)
	ServerSideProcessingTimeout() foundation.Number
	SetServerSideProcessingTimeout(value foundation.INumber)
	TimedInvokeTimeoutMs() foundation.Number
	SetTimedInvokeTimeoutMs(value foundation.INumber)
}

//
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


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterManagementClusterBoostParams/boostInfo
func (m_ MTRWaterHeaterManagementClusterBoostParams) BoostInfo() MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct {
	rv := objc.Send[MTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct](m_.ID, objc.Sel("boostInfo"))
	return rv
}


// SetBoostInfo sets the value of the boostInfo property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterManagementClusterBoostParams/boostInfo
func (m_ MTRWaterHeaterManagementClusterBoostParams) SetBoostInfo(value IMTRWaterHeaterManagementClusterWaterHeaterBoostInfoStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBoostInfo:"), value)
}

// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterManagementClusterBoostParams/serverSideProcessingTimeout
func (m_ MTRWaterHeaterManagementClusterBoostParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
// Controls how much time, in seconds, we will allow for the server to process the command.

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterManagementClusterBoostParams/serverSideProcessingTimeout
func (m_ MTRWaterHeaterManagementClusterBoostParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterManagementClusterBoostParams/timedInvokeTimeoutMs
func (m_ MTRWaterHeaterManagementClusterBoostParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
// Controls whether the command is a timed command (using Timed Invoke).

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterManagementClusterBoostParams/timedInvokeTimeoutMs
func (m_ MTRWaterHeaterManagementClusterBoostParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



