// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams] class.
var (
	MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParamsClass     _MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParamsClass
	MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParamsClassOnce sync.Once
)

func getMTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParamsClass() _MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParamsClass {
	MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParamsClassOnce.Do(func() {
		MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParamsClass = _MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParamsClass{objc.GetClass("MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams")}
	})
	return MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParamsClass
}

type _MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams] class.
type IMTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams interface {
	IMTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams
	// properties:
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	SoftwareVersion() objc.IObject /* cross-framework: NSNumber */
	SetSoftwareVersion(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	UpdateToken() objc.IObject /* cross-framework: Data */
	SetUpdateToken(value objc.IObject /* cross-framework: Data */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams-6s1jv
type MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams struct {
	MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams
}

// MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParamsFrom constructs a [MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams] from an unsafe.Pointer.
func MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParamsFrom(ptr unsafe.Pointer) MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams {
	return MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams{
		MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams: MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParamsFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParamsClass) Alloc() MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams {
	rv := objc.Send[MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParamsClass) New() MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams {
	rv := objc.Send[MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams) Init() MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams {
	rv := objc.Send[MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams) Autorelease() MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams {
	rv := objc.Send[MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams creates a new MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams instance.
func NewMTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams() MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams {
	return getMTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusternotifyupdateappliedparams-5eau8/serversideprocessingtimeout
func (m_ MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusternotifyupdateappliedparams-5eau8/serversideprocessingtimeout
func (m_ MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusternotifyupdateappliedparams-5eau8/softwareversion
func (m_ MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams) SoftwareVersion() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("softwareVersion"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusternotifyupdateappliedparams-5eau8/softwareversion
func (m_ MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams) SetSoftwareVersion(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSoftwareVersion:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusternotifyupdateappliedparams-5eau8/timedinvoketimeoutms
func (m_ MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusternotifyupdateappliedparams-5eau8/timedinvoketimeoutms
func (m_ MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusternotifyupdateappliedparams-5eau8/updatetoken
func (m_ MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams) UpdateToken() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("updateToken"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusternotifyupdateappliedparams-5eau8/updatetoken
func (m_ MTROtaSoftwareUpdateProviderClusterNotifyUpdateAppliedParams) SetUpdateToken(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUpdateToken:"), value)
}



