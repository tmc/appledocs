// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams] class.
var (
	MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParamsClass     _MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParamsClass
	MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParamsClassOnce sync.Once
)

func getMTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParamsClass() _MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParamsClass {
	MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParamsClassOnce.Do(func() {
		MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParamsClass = _MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParamsClass{objc.GetClass("MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams")}
	})
	return MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParamsClass
}

type _MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams] class.
type IMTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams-5eau8
type MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams struct {
	objectivec.Object
}

// MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParamsFrom constructs a [MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams] from an unsafe.Pointer.
func MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParamsFrom(ptr unsafe.Pointer) MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams {
	return MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParamsClass) Alloc() MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams {
	rv := objc.Send[MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParamsClass) New() MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams {
	rv := objc.Send[MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams) Init() MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams {
	rv := objc.Send[MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams) Autorelease() MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams {
	rv := objc.Send[MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams creates a new MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams instance.
func NewMTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams() MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams {
	return getMTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusternotifyupdateappliedparams-5eau8/serversideprocessingtimeout
func (m_ MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusternotifyupdateappliedparams-5eau8/serversideprocessingtimeout
func (m_ MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusternotifyupdateappliedparams-5eau8/softwareversion
func (m_ MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams) SoftwareVersion() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("softwareVersion"))
	return rv
}


// SetSoftwareVersion sets the value of the softwareVersion property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusternotifyupdateappliedparams-5eau8/softwareversion
func (m_ MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams) SetSoftwareVersion(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSoftwareVersion:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusternotifyupdateappliedparams-5eau8/timedinvoketimeoutms
func (m_ MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusternotifyupdateappliedparams-5eau8/timedinvoketimeoutms
func (m_ MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusternotifyupdateappliedparams-5eau8/updatetoken
func (m_ MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams) UpdateToken() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("updateToken"))
	return rv
}


// SetUpdateToken sets the value of the updateToken property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusternotifyupdateappliedparams-5eau8/updatetoken
func (m_ MTROTASoftwareUpdateProviderClusterNotifyUpdateAppliedParams) SetUpdateToken(value foundation.IData) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUpdateToken:"), value)
}



