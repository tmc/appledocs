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
	// properties:
	NewVersion() objc.IObject /* cross-framework: NSNumber */
	SetNewVersion(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	UpdateToken() objc.IObject /* cross-framework: Data */
	SetUpdateToken(value objc.IObject /* cross-framework: Data */)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterapplyupdaterequestparams-1mlcr/newversion
func (m_ MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams) NewVersion() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("newVersion"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterapplyupdaterequestparams-1mlcr/newversion
func (m_ MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams) SetNewVersion(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNewVersion:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterapplyupdaterequestparams-1mlcr/serversideprocessingtimeout
func (m_ MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterapplyupdaterequestparams-1mlcr/serversideprocessingtimeout
func (m_ MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterapplyupdaterequestparams-1mlcr/timedinvoketimeoutms
func (m_ MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterapplyupdaterequestparams-1mlcr/timedinvoketimeoutms
func (m_ MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterapplyupdaterequestparams-1mlcr/updatetoken
func (m_ MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams) UpdateToken() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("updateToken"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterapplyupdaterequestparams-1mlcr/updatetoken
func (m_ MTROtaSoftwareUpdateProviderClusterApplyUpdateRequestParams) SetUpdateToken(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUpdateToken:"), value)
}



