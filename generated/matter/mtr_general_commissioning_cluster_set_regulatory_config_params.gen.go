// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRGeneralCommissioningClusterSetRegulatoryConfigParams] class.
var (
	MTRGeneralCommissioningClusterSetRegulatoryConfigParamsClass     _MTRGeneralCommissioningClusterSetRegulatoryConfigParamsClass
	MTRGeneralCommissioningClusterSetRegulatoryConfigParamsClassOnce sync.Once
)

func getMTRGeneralCommissioningClusterSetRegulatoryConfigParamsClass() _MTRGeneralCommissioningClusterSetRegulatoryConfigParamsClass {
	MTRGeneralCommissioningClusterSetRegulatoryConfigParamsClassOnce.Do(func() {
		MTRGeneralCommissioningClusterSetRegulatoryConfigParamsClass = _MTRGeneralCommissioningClusterSetRegulatoryConfigParamsClass{objc.GetClass("MTRGeneralCommissioningClusterSetRegulatoryConfigParams")}
	})
	return MTRGeneralCommissioningClusterSetRegulatoryConfigParamsClass
}

type _MTRGeneralCommissioningClusterSetRegulatoryConfigParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTRGeneralCommissioningClusterSetRegulatoryConfigParams] class.
type IMTRGeneralCommissioningClusterSetRegulatoryConfigParams interface {
	objectivec.IObject
	// properties:
	Breadcrumb() objc.IObject /* cross-framework: NSNumber */
	SetBreadcrumb(value objc.IObject /* cross-framework: NSNumber */)
	CountryCode() objc.IObject /* cross-framework: NSString */
	SetCountryCode(value objc.IObject /* cross-framework: NSString */)
	NewRegulatoryConfig() objc.IObject /* cross-framework: NSNumber */
	SetNewRegulatoryConfig(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGeneralCommissioningClusterSetRegulatoryConfigParams
type MTRGeneralCommissioningClusterSetRegulatoryConfigParams struct {
	objectivec.Object
}

// MTRGeneralCommissioningClusterSetRegulatoryConfigParamsFrom constructs a [MTRGeneralCommissioningClusterSetRegulatoryConfigParams] from an unsafe.Pointer.
func MTRGeneralCommissioningClusterSetRegulatoryConfigParamsFrom(ptr unsafe.Pointer) MTRGeneralCommissioningClusterSetRegulatoryConfigParams {
	return MTRGeneralCommissioningClusterSetRegulatoryConfigParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRGeneralCommissioningClusterSetRegulatoryConfigParamsClass) Alloc() MTRGeneralCommissioningClusterSetRegulatoryConfigParams {
	rv := objc.Send[MTRGeneralCommissioningClusterSetRegulatoryConfigParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRGeneralCommissioningClusterSetRegulatoryConfigParamsClass) New() MTRGeneralCommissioningClusterSetRegulatoryConfigParams {
	rv := objc.Send[MTRGeneralCommissioningClusterSetRegulatoryConfigParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRGeneralCommissioningClusterSetRegulatoryConfigParams) Init() MTRGeneralCommissioningClusterSetRegulatoryConfigParams {
	rv := objc.Send[MTRGeneralCommissioningClusterSetRegulatoryConfigParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRGeneralCommissioningClusterSetRegulatoryConfigParams) Autorelease() MTRGeneralCommissioningClusterSetRegulatoryConfigParams {
	rv := objc.Send[MTRGeneralCommissioningClusterSetRegulatoryConfigParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRGeneralCommissioningClusterSetRegulatoryConfigParams creates a new MTRGeneralCommissioningClusterSetRegulatoryConfigParams instance.
func NewMTRGeneralCommissioningClusterSetRegulatoryConfigParams() MTRGeneralCommissioningClusterSetRegulatoryConfigParams {
	return getMTRGeneralCommissioningClusterSetRegulatoryConfigParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneralcommissioningclustersetregulatoryconfigparams/breadcrumb
func (m_ MTRGeneralCommissioningClusterSetRegulatoryConfigParams) Breadcrumb() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("breadcrumb"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneralcommissioningclustersetregulatoryconfigparams/breadcrumb
func (m_ MTRGeneralCommissioningClusterSetRegulatoryConfigParams) SetBreadcrumb(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBreadcrumb:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneralcommissioningclustersetregulatoryconfigparams/countrycode
func (m_ MTRGeneralCommissioningClusterSetRegulatoryConfigParams) CountryCode() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("countryCode"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneralcommissioningclustersetregulatoryconfigparams/countrycode
func (m_ MTRGeneralCommissioningClusterSetRegulatoryConfigParams) SetCountryCode(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCountryCode:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneralcommissioningclustersetregulatoryconfigparams/newregulatoryconfig
func (m_ MTRGeneralCommissioningClusterSetRegulatoryConfigParams) NewRegulatoryConfig() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("newRegulatoryConfig"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneralcommissioningclustersetregulatoryconfigparams/newregulatoryconfig
func (m_ MTRGeneralCommissioningClusterSetRegulatoryConfigParams) SetNewRegulatoryConfig(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNewRegulatoryConfig:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneralcommissioningclustersetregulatoryconfigparams/serversideprocessingtimeout
func (m_ MTRGeneralCommissioningClusterSetRegulatoryConfigParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneralcommissioningclustersetregulatoryconfigparams/serversideprocessingtimeout
func (m_ MTRGeneralCommissioningClusterSetRegulatoryConfigParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneralcommissioningclustersetregulatoryconfigparams/timedinvoketimeoutms
func (m_ MTRGeneralCommissioningClusterSetRegulatoryConfigParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgeneralcommissioningclustersetregulatoryconfigparams/timedinvoketimeoutms
func (m_ MTRGeneralCommissioningClusterSetRegulatoryConfigParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



