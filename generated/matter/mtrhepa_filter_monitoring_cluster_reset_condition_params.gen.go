// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRHEPAFilterMonitoringClusterResetConditionParams] class.
var (
	MTRHEPAFilterMonitoringClusterResetConditionParamsClass     _MTRHEPAFilterMonitoringClusterResetConditionParamsClass
	MTRHEPAFilterMonitoringClusterResetConditionParamsClassOnce sync.Once
)

func getMTRHEPAFilterMonitoringClusterResetConditionParamsClass() _MTRHEPAFilterMonitoringClusterResetConditionParamsClass {
	MTRHEPAFilterMonitoringClusterResetConditionParamsClassOnce.Do(func() {
		MTRHEPAFilterMonitoringClusterResetConditionParamsClass = _MTRHEPAFilterMonitoringClusterResetConditionParamsClass{objc.GetClass("MTRHEPAFilterMonitoringClusterResetConditionParams")}
	})
	return MTRHEPAFilterMonitoringClusterResetConditionParamsClass
}

type _MTRHEPAFilterMonitoringClusterResetConditionParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRHEPAFilterMonitoringClusterResetConditionParams] class.
type IMTRHEPAFilterMonitoringClusterResetConditionParams interface {
	objectivec.IObject
	// properties:
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRHEPAFilterMonitoringClusterResetConditionParams
type MTRHEPAFilterMonitoringClusterResetConditionParams struct {
	objectivec.Object
}

// MTRHEPAFilterMonitoringClusterResetConditionParamsFrom constructs a [MTRHEPAFilterMonitoringClusterResetConditionParams] from an unsafe.Pointer.
func MTRHEPAFilterMonitoringClusterResetConditionParamsFrom(ptr unsafe.Pointer) MTRHEPAFilterMonitoringClusterResetConditionParams {
	return MTRHEPAFilterMonitoringClusterResetConditionParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRHEPAFilterMonitoringClusterResetConditionParamsClass) Alloc() MTRHEPAFilterMonitoringClusterResetConditionParams {
	rv := objc.Send[MTRHEPAFilterMonitoringClusterResetConditionParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRHEPAFilterMonitoringClusterResetConditionParamsClass) New() MTRHEPAFilterMonitoringClusterResetConditionParams {
	rv := objc.Send[MTRHEPAFilterMonitoringClusterResetConditionParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRHEPAFilterMonitoringClusterResetConditionParams) Init() MTRHEPAFilterMonitoringClusterResetConditionParams {
	rv := objc.Send[MTRHEPAFilterMonitoringClusterResetConditionParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRHEPAFilterMonitoringClusterResetConditionParams) Autorelease() MTRHEPAFilterMonitoringClusterResetConditionParams {
	rv := objc.Send[MTRHEPAFilterMonitoringClusterResetConditionParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRHEPAFilterMonitoringClusterResetConditionParams creates a new MTRHEPAFilterMonitoringClusterResetConditionParams instance.
func NewMTRHEPAFilterMonitoringClusterResetConditionParams() MTRHEPAFilterMonitoringClusterResetConditionParams {
	return getMTRHEPAFilterMonitoringClusterResetConditionParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrhepafiltermonitoringclusterresetconditionparams/serversideprocessingtimeout
func (m_ MTRHEPAFilterMonitoringClusterResetConditionParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrhepafiltermonitoringclusterresetconditionparams/serversideprocessingtimeout
func (m_ MTRHEPAFilterMonitoringClusterResetConditionParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrhepafiltermonitoringclusterresetconditionparams/timedinvoketimeoutms
func (m_ MTRHEPAFilterMonitoringClusterResetConditionParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrhepafiltermonitoringclusterresetconditionparams/timedinvoketimeoutms
func (m_ MTRHEPAFilterMonitoringClusterResetConditionParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



