// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRValveConfigurationAndControlClusterCloseParams] class.
var (
	MTRValveConfigurationAndControlClusterCloseParamsClass     _MTRValveConfigurationAndControlClusterCloseParamsClass
	MTRValveConfigurationAndControlClusterCloseParamsClassOnce sync.Once
)

func getMTRValveConfigurationAndControlClusterCloseParamsClass() _MTRValveConfigurationAndControlClusterCloseParamsClass {
	MTRValveConfigurationAndControlClusterCloseParamsClassOnce.Do(func() {
		MTRValveConfigurationAndControlClusterCloseParamsClass = _MTRValveConfigurationAndControlClusterCloseParamsClass{objc.GetClass("MTRValveConfigurationAndControlClusterCloseParams")}
	})
	return MTRValveConfigurationAndControlClusterCloseParamsClass
}

type _MTRValveConfigurationAndControlClusterCloseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRValveConfigurationAndControlClusterCloseParams] class.
type IMTRValveConfigurationAndControlClusterCloseParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRValveConfigurationAndControlClusterCloseParams
type MTRValveConfigurationAndControlClusterCloseParams struct {
	objectivec.Object
}

// MTRValveConfigurationAndControlClusterCloseParamsFrom constructs a [MTRValveConfigurationAndControlClusterCloseParams] from an unsafe.Pointer.
func MTRValveConfigurationAndControlClusterCloseParamsFrom(ptr unsafe.Pointer) MTRValveConfigurationAndControlClusterCloseParams {
	return MTRValveConfigurationAndControlClusterCloseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRValveConfigurationAndControlClusterCloseParamsClass) Alloc() MTRValveConfigurationAndControlClusterCloseParams {
	rv := objc.Send[MTRValveConfigurationAndControlClusterCloseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRValveConfigurationAndControlClusterCloseParamsClass) New() MTRValveConfigurationAndControlClusterCloseParams {
	rv := objc.Send[MTRValveConfigurationAndControlClusterCloseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRValveConfigurationAndControlClusterCloseParams) Init() MTRValveConfigurationAndControlClusterCloseParams {
	rv := objc.Send[MTRValveConfigurationAndControlClusterCloseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRValveConfigurationAndControlClusterCloseParams) Autorelease() MTRValveConfigurationAndControlClusterCloseParams {
	rv := objc.Send[MTRValveConfigurationAndControlClusterCloseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRValveConfigurationAndControlClusterCloseParams creates a new MTRValveConfigurationAndControlClusterCloseParams instance.
func NewMTRValveConfigurationAndControlClusterCloseParams() MTRValveConfigurationAndControlClusterCloseParams {
	return getMTRValveConfigurationAndControlClusterCloseParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrvalveconfigurationandcontrolclustercloseparams/serversideprocessingtimeout
func (m_ MTRValveConfigurationAndControlClusterCloseParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrvalveconfigurationandcontrolclustercloseparams/serversideprocessingtimeout
func (m_ MTRValveConfigurationAndControlClusterCloseParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrvalveconfigurationandcontrolclustercloseparams/timedinvoketimeoutms
func (m_ MTRValveConfigurationAndControlClusterCloseParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrvalveconfigurationandcontrolclustercloseparams/timedinvoketimeoutms
func (m_ MTRValveConfigurationAndControlClusterCloseParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



