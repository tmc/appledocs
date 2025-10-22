// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRValveConfigurationAndControlClusterOpenParams] class.
var (
	MTRValveConfigurationAndControlClusterOpenParamsClass     _MTRValveConfigurationAndControlClusterOpenParamsClass
	MTRValveConfigurationAndControlClusterOpenParamsClassOnce sync.Once
)

func getMTRValveConfigurationAndControlClusterOpenParamsClass() _MTRValveConfigurationAndControlClusterOpenParamsClass {
	MTRValveConfigurationAndControlClusterOpenParamsClassOnce.Do(func() {
		MTRValveConfigurationAndControlClusterOpenParamsClass = _MTRValveConfigurationAndControlClusterOpenParamsClass{objc.GetClass("MTRValveConfigurationAndControlClusterOpenParams")}
	})
	return MTRValveConfigurationAndControlClusterOpenParamsClass
}

type _MTRValveConfigurationAndControlClusterOpenParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRValveConfigurationAndControlClusterOpenParams] class.
type IMTRValveConfigurationAndControlClusterOpenParams interface {
	objectivec.IObject
	OpenDuration() foundation.Number
	SetOpenDuration(value foundation.INumber)
	ServerSideProcessingTimeout() foundation.Number
	SetServerSideProcessingTimeout(value foundation.INumber)
	TargetLevel() foundation.Number
	SetTargetLevel(value foundation.INumber)
	TimedInvokeTimeoutMs() foundation.Number
	SetTimedInvokeTimeoutMs(value foundation.INumber)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRValveConfigurationAndControlClusterOpenParams
type MTRValveConfigurationAndControlClusterOpenParams struct {
	objectivec.Object
}

// MTRValveConfigurationAndControlClusterOpenParamsFrom constructs a [MTRValveConfigurationAndControlClusterOpenParams] from an unsafe.Pointer.
func MTRValveConfigurationAndControlClusterOpenParamsFrom(ptr unsafe.Pointer) MTRValveConfigurationAndControlClusterOpenParams {
	return MTRValveConfigurationAndControlClusterOpenParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRValveConfigurationAndControlClusterOpenParamsClass) Alloc() MTRValveConfigurationAndControlClusterOpenParams {
	rv := objc.Send[MTRValveConfigurationAndControlClusterOpenParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRValveConfigurationAndControlClusterOpenParamsClass) New() MTRValveConfigurationAndControlClusterOpenParams {
	rv := objc.Send[MTRValveConfigurationAndControlClusterOpenParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRValveConfigurationAndControlClusterOpenParams) Init() MTRValveConfigurationAndControlClusterOpenParams {
	rv := objc.Send[MTRValveConfigurationAndControlClusterOpenParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRValveConfigurationAndControlClusterOpenParams) Autorelease() MTRValveConfigurationAndControlClusterOpenParams {
	rv := objc.Send[MTRValveConfigurationAndControlClusterOpenParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRValveConfigurationAndControlClusterOpenParams creates a new MTRValveConfigurationAndControlClusterOpenParams instance.
func NewMTRValveConfigurationAndControlClusterOpenParams() MTRValveConfigurationAndControlClusterOpenParams {
	return getMTRValveConfigurationAndControlClusterOpenParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrvalveconfigurationandcontrolclusteropenparams/openduration
func (m_ MTRValveConfigurationAndControlClusterOpenParams) OpenDuration() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("openDuration"))
	return rv
}


// SetOpenDuration sets the value of the openDuration property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrvalveconfigurationandcontrolclusteropenparams/openduration
func (m_ MTRValveConfigurationAndControlClusterOpenParams) SetOpenDuration(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOpenDuration:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrvalveconfigurationandcontrolclusteropenparams/serversideprocessingtimeout
func (m_ MTRValveConfigurationAndControlClusterOpenParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrvalveconfigurationandcontrolclusteropenparams/serversideprocessingtimeout
func (m_ MTRValveConfigurationAndControlClusterOpenParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrvalveconfigurationandcontrolclusteropenparams/targetlevel
func (m_ MTRValveConfigurationAndControlClusterOpenParams) TargetLevel() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("targetLevel"))
	return rv
}


// SetTargetLevel sets the value of the targetLevel property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrvalveconfigurationandcontrolclusteropenparams/targetlevel
func (m_ MTRValveConfigurationAndControlClusterOpenParams) SetTargetLevel(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTargetLevel:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrvalveconfigurationandcontrolclusteropenparams/timedinvoketimeoutms
func (m_ MTRValveConfigurationAndControlClusterOpenParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrvalveconfigurationandcontrolclusteropenparams/timedinvoketimeoutms
func (m_ MTRValveConfigurationAndControlClusterOpenParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



