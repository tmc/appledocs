// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRServiceAreaClusterSelectAreasParams] class.
var (
	MTRServiceAreaClusterSelectAreasParamsClass     _MTRServiceAreaClusterSelectAreasParamsClass
	MTRServiceAreaClusterSelectAreasParamsClassOnce sync.Once
)

func getMTRServiceAreaClusterSelectAreasParamsClass() _MTRServiceAreaClusterSelectAreasParamsClass {
	MTRServiceAreaClusterSelectAreasParamsClassOnce.Do(func() {
		MTRServiceAreaClusterSelectAreasParamsClass = _MTRServiceAreaClusterSelectAreasParamsClass{objc.GetClass("MTRServiceAreaClusterSelectAreasParams")}
	})
	return MTRServiceAreaClusterSelectAreasParamsClass
}

type _MTRServiceAreaClusterSelectAreasParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRServiceAreaClusterSelectAreasParams] class.
type IMTRServiceAreaClusterSelectAreasParams interface {
	objectivec.IObject
	NewAreas() objc.ID
	SetNewAreas(value objc.ID)
	ServerSideProcessingTimeout() foundation.Number
	SetServerSideProcessingTimeout(value foundation.INumber)
	TimedInvokeTimeoutMs() foundation.Number
	SetTimedInvokeTimeoutMs(value foundation.INumber)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterSelectAreasParams
type MTRServiceAreaClusterSelectAreasParams struct {
	objectivec.Object
}

// MTRServiceAreaClusterSelectAreasParamsFrom constructs a [MTRServiceAreaClusterSelectAreasParams] from an unsafe.Pointer.
func MTRServiceAreaClusterSelectAreasParamsFrom(ptr unsafe.Pointer) MTRServiceAreaClusterSelectAreasParams {
	return MTRServiceAreaClusterSelectAreasParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRServiceAreaClusterSelectAreasParamsClass) Alloc() MTRServiceAreaClusterSelectAreasParams {
	rv := objc.Send[MTRServiceAreaClusterSelectAreasParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRServiceAreaClusterSelectAreasParamsClass) New() MTRServiceAreaClusterSelectAreasParams {
	rv := objc.Send[MTRServiceAreaClusterSelectAreasParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRServiceAreaClusterSelectAreasParams) Init() MTRServiceAreaClusterSelectAreasParams {
	rv := objc.Send[MTRServiceAreaClusterSelectAreasParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRServiceAreaClusterSelectAreasParams) Autorelease() MTRServiceAreaClusterSelectAreasParams {
	rv := objc.Send[MTRServiceAreaClusterSelectAreasParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRServiceAreaClusterSelectAreasParams creates a new MTRServiceAreaClusterSelectAreasParams instance.
func NewMTRServiceAreaClusterSelectAreasParams() MTRServiceAreaClusterSelectAreasParams {
	return getMTRServiceAreaClusterSelectAreasParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterSelectAreasParams/newAreas
func (m_ MTRServiceAreaClusterSelectAreasParams) NewAreas() objc.ID {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("newAreas"))
	return rv
}


// SetNewAreas sets the value of the newAreas property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterSelectAreasParams/newAreas
func (m_ MTRServiceAreaClusterSelectAreasParams) SetNewAreas(value objc.ID) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNewAreas:"), value)
}

// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterSelectAreasParams/serverSideProcessingTimeout
func (m_ MTRServiceAreaClusterSelectAreasParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
// Controls how much time, in seconds, we will allow for the server to process the command.

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterSelectAreasParams/serverSideProcessingTimeout
func (m_ MTRServiceAreaClusterSelectAreasParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterSelectAreasParams/timedInvokeTimeoutMs
func (m_ MTRServiceAreaClusterSelectAreasParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
// Controls whether the command is a timed command (using Timed Invoke).

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterSelectAreasParams/timedInvokeTimeoutMs
func (m_ MTRServiceAreaClusterSelectAreasParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



