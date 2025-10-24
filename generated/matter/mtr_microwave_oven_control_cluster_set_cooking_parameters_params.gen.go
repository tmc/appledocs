// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRMicrowaveOvenControlClusterSetCookingParametersParams] class.
var (
	MTRMicrowaveOvenControlClusterSetCookingParametersParamsClass     _MTRMicrowaveOvenControlClusterSetCookingParametersParamsClass
	MTRMicrowaveOvenControlClusterSetCookingParametersParamsClassOnce sync.Once
)

func getMTRMicrowaveOvenControlClusterSetCookingParametersParamsClass() _MTRMicrowaveOvenControlClusterSetCookingParametersParamsClass {
	MTRMicrowaveOvenControlClusterSetCookingParametersParamsClassOnce.Do(func() {
		MTRMicrowaveOvenControlClusterSetCookingParametersParamsClass = _MTRMicrowaveOvenControlClusterSetCookingParametersParamsClass{objc.GetClass("MTRMicrowaveOvenControlClusterSetCookingParametersParams")}
	})
	return MTRMicrowaveOvenControlClusterSetCookingParametersParamsClass
}

type _MTRMicrowaveOvenControlClusterSetCookingParametersParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRMicrowaveOvenControlClusterSetCookingParametersParams] class.
type IMTRMicrowaveOvenControlClusterSetCookingParametersParams interface {
	objectivec.IObject
	// properties:
	CookMode() objc.IObject /* cross-framework: NSNumber */
	SetCookMode(value objc.IObject /* cross-framework: NSNumber */)
	CookTime() objc.IObject /* cross-framework: NSNumber */
	SetCookTime(value objc.IObject /* cross-framework: NSNumber */)
	PowerSetting() objc.IObject /* cross-framework: NSNumber */
	SetPowerSetting(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	StartAfterSetting() objc.IObject /* cross-framework: NSNumber */
	SetStartAfterSetting(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMicrowaveOvenControlClusterSetCookingParametersParams
type MTRMicrowaveOvenControlClusterSetCookingParametersParams struct {
	objectivec.Object
}

// MTRMicrowaveOvenControlClusterSetCookingParametersParamsFrom constructs a [MTRMicrowaveOvenControlClusterSetCookingParametersParams] from an unsafe.Pointer.
func MTRMicrowaveOvenControlClusterSetCookingParametersParamsFrom(ptr unsafe.Pointer) MTRMicrowaveOvenControlClusterSetCookingParametersParams {
	return MTRMicrowaveOvenControlClusterSetCookingParametersParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRMicrowaveOvenControlClusterSetCookingParametersParamsClass) Alloc() MTRMicrowaveOvenControlClusterSetCookingParametersParams {
	rv := objc.Send[MTRMicrowaveOvenControlClusterSetCookingParametersParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRMicrowaveOvenControlClusterSetCookingParametersParamsClass) New() MTRMicrowaveOvenControlClusterSetCookingParametersParams {
	rv := objc.Send[MTRMicrowaveOvenControlClusterSetCookingParametersParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRMicrowaveOvenControlClusterSetCookingParametersParams) Init() MTRMicrowaveOvenControlClusterSetCookingParametersParams {
	rv := objc.Send[MTRMicrowaveOvenControlClusterSetCookingParametersParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRMicrowaveOvenControlClusterSetCookingParametersParams) Autorelease() MTRMicrowaveOvenControlClusterSetCookingParametersParams {
	rv := objc.Send[MTRMicrowaveOvenControlClusterSetCookingParametersParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRMicrowaveOvenControlClusterSetCookingParametersParams creates a new MTRMicrowaveOvenControlClusterSetCookingParametersParams instance.
func NewMTRMicrowaveOvenControlClusterSetCookingParametersParams() MTRMicrowaveOvenControlClusterSetCookingParametersParams {
	return getMTRMicrowaveOvenControlClusterSetCookingParametersParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMicrowaveOvenControlClusterSetCookingParametersParams/cookMode
func (m_ MTRMicrowaveOvenControlClusterSetCookingParametersParams) CookMode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("cookMode"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMicrowaveOvenControlClusterSetCookingParametersParams/cookMode
func (m_ MTRMicrowaveOvenControlClusterSetCookingParametersParams) SetCookMode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCookMode:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMicrowaveOvenControlClusterSetCookingParametersParams/cookTime
func (m_ MTRMicrowaveOvenControlClusterSetCookingParametersParams) CookTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("cookTime"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMicrowaveOvenControlClusterSetCookingParametersParams/cookTime
func (m_ MTRMicrowaveOvenControlClusterSetCookingParametersParams) SetCookTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCookTime:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMicrowaveOvenControlClusterSetCookingParametersParams/powerSetting
func (m_ MTRMicrowaveOvenControlClusterSetCookingParametersParams) PowerSetting() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("powerSetting"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMicrowaveOvenControlClusterSetCookingParametersParams/powerSetting
func (m_ MTRMicrowaveOvenControlClusterSetCookingParametersParams) SetPowerSetting(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPowerSetting:"), value)
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMicrowaveOvenControlClusterSetCookingParametersParams/serverSideProcessingTimeout
func (m_ MTRMicrowaveOvenControlClusterSetCookingParametersParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMicrowaveOvenControlClusterSetCookingParametersParams/serverSideProcessingTimeout
func (m_ MTRMicrowaveOvenControlClusterSetCookingParametersParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMicrowaveOvenControlClusterSetCookingParametersParams/startAfterSetting
func (m_ MTRMicrowaveOvenControlClusterSetCookingParametersParams) StartAfterSetting() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("startAfterSetting"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMicrowaveOvenControlClusterSetCookingParametersParams/startAfterSetting
func (m_ MTRMicrowaveOvenControlClusterSetCookingParametersParams) SetStartAfterSetting(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStartAfterSetting:"), value)
}


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMicrowaveOvenControlClusterSetCookingParametersParams/timedInvokeTimeoutMs
func (m_ MTRMicrowaveOvenControlClusterSetCookingParametersParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMicrowaveOvenControlClusterSetCookingParametersParams/timedInvokeTimeoutMs
func (m_ MTRMicrowaveOvenControlClusterSetCookingParametersParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



