// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRTestClusterClusterTestUnknownCommandParams] class.
var (
	MTRTestClusterClusterTestUnknownCommandParamsClass     _MTRTestClusterClusterTestUnknownCommandParamsClass
	MTRTestClusterClusterTestUnknownCommandParamsClassOnce sync.Once
)

func getMTRTestClusterClusterTestUnknownCommandParamsClass() _MTRTestClusterClusterTestUnknownCommandParamsClass {
	MTRTestClusterClusterTestUnknownCommandParamsClassOnce.Do(func() {
		MTRTestClusterClusterTestUnknownCommandParamsClass = _MTRTestClusterClusterTestUnknownCommandParamsClass{objc.GetClass("MTRTestClusterClusterTestUnknownCommandParams")}
	})
	return MTRTestClusterClusterTestUnknownCommandParamsClass
}

type _MTRTestClusterClusterTestUnknownCommandParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRTestClusterClusterTestUnknownCommandParams] class.
type IMTRTestClusterClusterTestUnknownCommandParams interface {
	IMTRUnitTestingClusterTestUnknownCommandParams
	ServerSideProcessingTimeout() foundation.Number
	SetServerSideProcessingTimeout(value foundation.INumber)
	TimedInvokeTimeoutMs() foundation.Number
	SetTimedInvokeTimeoutMs(value foundation.INumber)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTestClusterClusterTestUnknownCommandParams
type MTRTestClusterClusterTestUnknownCommandParams struct {
	MTRUnitTestingClusterTestUnknownCommandParams
}

// MTRTestClusterClusterTestUnknownCommandParamsFrom constructs a [MTRTestClusterClusterTestUnknownCommandParams] from an unsafe.Pointer.
func MTRTestClusterClusterTestUnknownCommandParamsFrom(ptr unsafe.Pointer) MTRTestClusterClusterTestUnknownCommandParams {
	return MTRTestClusterClusterTestUnknownCommandParams{
		MTRUnitTestingClusterTestUnknownCommandParams: MTRUnitTestingClusterTestUnknownCommandParamsFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTestClusterClusterTestUnknownCommandParamsClass) Alloc() MTRTestClusterClusterTestUnknownCommandParams {
	rv := objc.Send[MTRTestClusterClusterTestUnknownCommandParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTestClusterClusterTestUnknownCommandParamsClass) New() MTRTestClusterClusterTestUnknownCommandParams {
	rv := objc.Send[MTRTestClusterClusterTestUnknownCommandParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTestClusterClusterTestUnknownCommandParams) Init() MTRTestClusterClusterTestUnknownCommandParams {
	rv := objc.Send[MTRTestClusterClusterTestUnknownCommandParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTestClusterClusterTestUnknownCommandParams) Autorelease() MTRTestClusterClusterTestUnknownCommandParams {
	rv := objc.Send[MTRTestClusterClusterTestUnknownCommandParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTestClusterClusterTestUnknownCommandParams creates a new MTRTestClusterClusterTestUnknownCommandParams instance.
func NewMTRTestClusterClusterTestUnknownCommandParams() MTRTestClusterClusterTestUnknownCommandParams {
	return getMTRTestClusterClusterTestUnknownCommandParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestunknowncommandparams/serversideprocessingtimeout
func (m_ MTRTestClusterClusterTestUnknownCommandParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestunknowncommandparams/serversideprocessingtimeout
func (m_ MTRTestClusterClusterTestUnknownCommandParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestunknowncommandparams/timedinvoketimeoutms
func (m_ MTRTestClusterClusterTestUnknownCommandParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestunknowncommandparams/timedinvoketimeoutms
func (m_ MTRTestClusterClusterTestUnknownCommandParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



