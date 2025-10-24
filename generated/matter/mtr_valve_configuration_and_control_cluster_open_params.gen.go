// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
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
	// properties:
	OpenDuration() objc.IObject /* cross-framework: NSNumber */
	SetOpenDuration(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TargetLevel() objc.IObject /* cross-framework: NSNumber */
	SetTargetLevel(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}

// [Full Topic]
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

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrvalveconfigurationandcontrolclusteropenparams/openduration
func (m_ MTRValveConfigurationAndControlClusterOpenParams) OpenDuration() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("openDuration"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrvalveconfigurationandcontrolclusteropenparams/openduration
func (m_ MTRValveConfigurationAndControlClusterOpenParams) SetOpenDuration(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOpenDuration:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrvalveconfigurationandcontrolclusteropenparams/serversideprocessingtimeout
func (m_ MTRValveConfigurationAndControlClusterOpenParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrvalveconfigurationandcontrolclusteropenparams/serversideprocessingtimeout
func (m_ MTRValveConfigurationAndControlClusterOpenParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrvalveconfigurationandcontrolclusteropenparams/targetlevel
func (m_ MTRValveConfigurationAndControlClusterOpenParams) TargetLevel() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("targetLevel"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrvalveconfigurationandcontrolclusteropenparams/targetlevel
func (m_ MTRValveConfigurationAndControlClusterOpenParams) SetTargetLevel(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTargetLevel:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrvalveconfigurationandcontrolclusteropenparams/timedinvoketimeoutms
func (m_ MTRValveConfigurationAndControlClusterOpenParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrvalveconfigurationandcontrolclusteropenparams/timedinvoketimeoutms
func (m_ MTRValveConfigurationAndControlClusterOpenParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}
