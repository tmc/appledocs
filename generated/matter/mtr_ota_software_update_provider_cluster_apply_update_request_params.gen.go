// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams] class.
var (
	MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParamsClass     _MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParamsClass
	MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParamsClassOnce sync.Once
)

func getMTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParamsClass() _MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParamsClass {
	MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParamsClassOnce.Do(func() {
		MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParamsClass = _MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParamsClass{objc.GetClass("MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams")}
	})
	return MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParamsClass
}

type _MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams] class.
type IMTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams interface {
	IMTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams-5k4nj
type MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams struct {
	MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams
}

// MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParamsFrom constructs a [MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams] from an unsafe.Pointer.
func MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParamsFrom(ptr unsafe.Pointer) MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams {
	return MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams{
		MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParams: MTROTASoftwareUpdateProviderClusterApplyUpdateRequestParamsFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParamsClass) Alloc() MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams {
	rv := objc.Send[MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParamsClass) New() MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams {
	rv := objc.Send[MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams) Init() MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams {
	rv := objc.Send[MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams) Autorelease() MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams {
	rv := objc.Send[MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams creates a new MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams instance.
func NewMTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams() MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams {
	return getMTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterapplyupdaterequestparams-1mlcr/newversion
func (m_ MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams) NewVersion() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("newVersion"))
	return rv
}


// SetNewVersion sets the value of the newVersion property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterapplyupdaterequestparams-1mlcr/newversion
func (m_ MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams) SetNewVersion(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNewVersion:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterapplyupdaterequestparams-1mlcr/serversideprocessingtimeout
func (m_ MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterapplyupdaterequestparams-1mlcr/serversideprocessingtimeout
func (m_ MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterapplyupdaterequestparams-1mlcr/timedinvoketimeoutms
func (m_ MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterapplyupdaterequestparams-1mlcr/timedinvoketimeoutms
func (m_ MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterapplyupdaterequestparams-1mlcr/updatetoken
func (m_ MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams) UpdateToken() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("updateToken"))
	return rv
}


// SetUpdateToken sets the value of the updateToken property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterapplyupdaterequestparams-1mlcr/updatetoken
func (m_ MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams) SetUpdateToken(value foundation.IData) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUpdateToken:"), value)
}



