// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRRVCOperationalStateClusterGoHomeParams] class.
var (
	MTRRVCOperationalStateClusterGoHomeParamsClass     _MTRRVCOperationalStateClusterGoHomeParamsClass
	MTRRVCOperationalStateClusterGoHomeParamsClassOnce sync.Once
)

func getMTRRVCOperationalStateClusterGoHomeParamsClass() _MTRRVCOperationalStateClusterGoHomeParamsClass {
	MTRRVCOperationalStateClusterGoHomeParamsClassOnce.Do(func() {
		MTRRVCOperationalStateClusterGoHomeParamsClass = _MTRRVCOperationalStateClusterGoHomeParamsClass{objc.GetClass("MTRRVCOperationalStateClusterGoHomeParams")}
	})
	return MTRRVCOperationalStateClusterGoHomeParamsClass
}

type _MTRRVCOperationalStateClusterGoHomeParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRRVCOperationalStateClusterGoHomeParams] class.
type IMTRRVCOperationalStateClusterGoHomeParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRRVCOperationalStateClusterGoHomeParams
type MTRRVCOperationalStateClusterGoHomeParams struct {
	objectivec.Object
}

// MTRRVCOperationalStateClusterGoHomeParamsFrom constructs a [MTRRVCOperationalStateClusterGoHomeParams] from an unsafe.Pointer.
func MTRRVCOperationalStateClusterGoHomeParamsFrom(ptr unsafe.Pointer) MTRRVCOperationalStateClusterGoHomeParams {
	return MTRRVCOperationalStateClusterGoHomeParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRRVCOperationalStateClusterGoHomeParamsClass) Alloc() MTRRVCOperationalStateClusterGoHomeParams {
	rv := objc.Send[MTRRVCOperationalStateClusterGoHomeParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRRVCOperationalStateClusterGoHomeParamsClass) New() MTRRVCOperationalStateClusterGoHomeParams {
	rv := objc.Send[MTRRVCOperationalStateClusterGoHomeParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRRVCOperationalStateClusterGoHomeParams) Init() MTRRVCOperationalStateClusterGoHomeParams {
	rv := objc.Send[MTRRVCOperationalStateClusterGoHomeParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRRVCOperationalStateClusterGoHomeParams) Autorelease() MTRRVCOperationalStateClusterGoHomeParams {
	rv := objc.Send[MTRRVCOperationalStateClusterGoHomeParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRRVCOperationalStateClusterGoHomeParams creates a new MTRRVCOperationalStateClusterGoHomeParams instance.
func NewMTRRVCOperationalStateClusterGoHomeParams() MTRRVCOperationalStateClusterGoHomeParams {
	return getMTRRVCOperationalStateClusterGoHomeParamsClass().New()
}


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRRVCOperationalStateClusterGoHomeParams/serverSideProcessingTimeout
func (m_ MTRRVCOperationalStateClusterGoHomeParams) ServerSideProcessingTimeout() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
// Controls how much time, in seconds, we will allow for the server to process the command.

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRRVCOperationalStateClusterGoHomeParams/serverSideProcessingTimeout
func (m_ MTRRVCOperationalStateClusterGoHomeParams) SetServerSideProcessingTimeout(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}
// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRRVCOperationalStateClusterGoHomeParams/timedInvokeTimeoutMs
func (m_ MTRRVCOperationalStateClusterGoHomeParams) TimedInvokeTimeoutMs() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
// Controls whether the command is a timed command (using Timed Invoke).

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRRVCOperationalStateClusterGoHomeParams/timedInvokeTimeoutMs
func (m_ MTRRVCOperationalStateClusterGoHomeParams) SetTimedInvokeTimeoutMs(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}


