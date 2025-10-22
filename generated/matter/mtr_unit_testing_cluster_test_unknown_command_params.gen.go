// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRUnitTestingClusterTestUnknownCommandParams] class.
var (
	MTRUnitTestingClusterTestUnknownCommandParamsClass     _MTRUnitTestingClusterTestUnknownCommandParamsClass
	MTRUnitTestingClusterTestUnknownCommandParamsClassOnce sync.Once
)

func getMTRUnitTestingClusterTestUnknownCommandParamsClass() _MTRUnitTestingClusterTestUnknownCommandParamsClass {
	MTRUnitTestingClusterTestUnknownCommandParamsClassOnce.Do(func() {
		MTRUnitTestingClusterTestUnknownCommandParamsClass = _MTRUnitTestingClusterTestUnknownCommandParamsClass{objc.GetClass("MTRUnitTestingClusterTestUnknownCommandParams")}
	})
	return MTRUnitTestingClusterTestUnknownCommandParamsClass
}

type _MTRUnitTestingClusterTestUnknownCommandParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRUnitTestingClusterTestUnknownCommandParams] class.
type IMTRUnitTestingClusterTestUnknownCommandParams interface {
	objectivec.IObject
	ServerSideProcessingTimeout() foundation.Number
	SetServerSideProcessingTimeout(value foundation.INumber)
	TimedInvokeTimeoutMs() foundation.Number
	SetTimedInvokeTimeoutMs(value foundation.INumber)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestUnknownCommandParams
type MTRUnitTestingClusterTestUnknownCommandParams struct {
	objectivec.Object
}

// MTRUnitTestingClusterTestUnknownCommandParamsFrom constructs a [MTRUnitTestingClusterTestUnknownCommandParams] from an unsafe.Pointer.
func MTRUnitTestingClusterTestUnknownCommandParamsFrom(ptr unsafe.Pointer) MTRUnitTestingClusterTestUnknownCommandParams {
	return MTRUnitTestingClusterTestUnknownCommandParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterTestUnknownCommandParamsClass) Alloc() MTRUnitTestingClusterTestUnknownCommandParams {
	rv := objc.Send[MTRUnitTestingClusterTestUnknownCommandParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRUnitTestingClusterTestUnknownCommandParamsClass) New() MTRUnitTestingClusterTestUnknownCommandParams {
	rv := objc.Send[MTRUnitTestingClusterTestUnknownCommandParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterTestUnknownCommandParams) Init() MTRUnitTestingClusterTestUnknownCommandParams {
	rv := objc.Send[MTRUnitTestingClusterTestUnknownCommandParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterTestUnknownCommandParams) Autorelease() MTRUnitTestingClusterTestUnknownCommandParams {
	rv := objc.Send[MTRUnitTestingClusterTestUnknownCommandParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterTestUnknownCommandParams creates a new MTRUnitTestingClusterTestUnknownCommandParams instance.
func NewMTRUnitTestingClusterTestUnknownCommandParams() MTRUnitTestingClusterTestUnknownCommandParams {
	return getMTRUnitTestingClusterTestUnknownCommandParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestunknowncommandparams/serversideprocessingtimeout
func (m_ MTRUnitTestingClusterTestUnknownCommandParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestunknowncommandparams/serversideprocessingtimeout
func (m_ MTRUnitTestingClusterTestUnknownCommandParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestunknowncommandparams/timedinvoketimeoutms
func (m_ MTRUnitTestingClusterTestUnknownCommandParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestunknowncommandparams/timedinvoketimeoutms
func (m_ MTRUnitTestingClusterTestUnknownCommandParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



