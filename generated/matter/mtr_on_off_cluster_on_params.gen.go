// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTROnOffClusterOnParams] class.
var (
	MTROnOffClusterOnParamsClass     _MTROnOffClusterOnParamsClass
	MTROnOffClusterOnParamsClassOnce sync.Once
)

func getMTROnOffClusterOnParamsClass() _MTROnOffClusterOnParamsClass {
	MTROnOffClusterOnParamsClassOnce.Do(func() {
		MTROnOffClusterOnParamsClass = _MTROnOffClusterOnParamsClass{objc.GetClass("MTROnOffClusterOnParams")}
	})
	return MTROnOffClusterOnParamsClass
}

type _MTROnOffClusterOnParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTROnOffClusterOnParams] class.
type IMTROnOffClusterOnParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROnOffClusterOnParams
type MTROnOffClusterOnParams struct {
	objectivec.Object
}

// MTROnOffClusterOnParamsFrom constructs a [MTROnOffClusterOnParams] from an unsafe.Pointer.
func MTROnOffClusterOnParamsFrom(ptr unsafe.Pointer) MTROnOffClusterOnParams {
	return MTROnOffClusterOnParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROnOffClusterOnParamsClass) Alloc() MTROnOffClusterOnParams {
	rv := objc.Send[MTROnOffClusterOnParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROnOffClusterOnParamsClass) New() MTROnOffClusterOnParams {
	rv := objc.Send[MTROnOffClusterOnParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROnOffClusterOnParams) Init() MTROnOffClusterOnParams {
	rv := objc.Send[MTROnOffClusterOnParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROnOffClusterOnParams) Autorelease() MTROnOffClusterOnParams {
	rv := objc.Send[MTROnOffClusterOnParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROnOffClusterOnParams creates a new MTROnOffClusterOnParams instance.
func NewMTROnOffClusterOnParams() MTROnOffClusterOnParams {
	return getMTROnOffClusterOnParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtronoffclusteronparams/serversideprocessingtimeout
func (m_ MTROnOffClusterOnParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtronoffclusteronparams/serversideprocessingtimeout
func (m_ MTROnOffClusterOnParams) SetServerSideProcessingTimeout(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtronoffclusteronparams/timedinvoketimeoutms
func (m_ MTROnOffClusterOnParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtronoffclusteronparams/timedinvoketimeoutms
func (m_ MTROnOffClusterOnParams) SetTimedInvokeTimeoutMs(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



