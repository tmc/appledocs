// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTROtaSoftwareUpdateProviderClusterQueryImageParams] class.
var (
	MTROtaSoftwareUpdateProviderClusterQueryImageParamsClass     _MTROtaSoftwareUpdateProviderClusterQueryImageParamsClass
	MTROtaSoftwareUpdateProviderClusterQueryImageParamsClassOnce sync.Once
)

func getMTROtaSoftwareUpdateProviderClusterQueryImageParamsClass() _MTROtaSoftwareUpdateProviderClusterQueryImageParamsClass {
	MTROtaSoftwareUpdateProviderClusterQueryImageParamsClassOnce.Do(func() {
		MTROtaSoftwareUpdateProviderClusterQueryImageParamsClass = _MTROtaSoftwareUpdateProviderClusterQueryImageParamsClass{objc.GetClass("MTROtaSoftwareUpdateProviderClusterQueryImageParams")}
	})
	return MTROtaSoftwareUpdateProviderClusterQueryImageParamsClass
}

type _MTROtaSoftwareUpdateProviderClusterQueryImageParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTROtaSoftwareUpdateProviderClusterQueryImageParams] class.
type IMTROtaSoftwareUpdateProviderClusterQueryImageParams interface {
	IMTROTASoftwareUpdateProviderClusterQueryImageParams
	// properties:
	HardwareVersion() objc.IObject /* cross-framework: NSNumber */
	SetHardwareVersion(value objc.IObject /* cross-framework: NSNumber */)
	Location() objc.IObject /* cross-framework: NSString */
	SetLocation(value objc.IObject /* cross-framework: NSString */)
	MetadataForProvider() objc.IObject /* cross-framework: Data */
	SetMetadataForProvider(value objc.IObject /* cross-framework: Data */)
	ProtocolsSupported() unsafe.Pointer
	SetProtocolsSupported(value unsafe.Pointer)
	RequestorCanConsent() objc.IObject /* cross-framework: NSNumber */
	SetRequestorCanConsent(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	SoftwareVersion() objc.IObject /* cross-framework: NSNumber */
	SetSoftwareVersion(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateProviderClusterQueryImageParams-zidv
type MTROtaSoftwareUpdateProviderClusterQueryImageParams struct {
	MTROTASoftwareUpdateProviderClusterQueryImageParams
}

// MTROtaSoftwareUpdateProviderClusterQueryImageParamsFrom constructs a [MTROtaSoftwareUpdateProviderClusterQueryImageParams] from an unsafe.Pointer.
func MTROtaSoftwareUpdateProviderClusterQueryImageParamsFrom(ptr unsafe.Pointer) MTROtaSoftwareUpdateProviderClusterQueryImageParams {
	return MTROtaSoftwareUpdateProviderClusterQueryImageParams{
		MTROTASoftwareUpdateProviderClusterQueryImageParams: MTROTASoftwareUpdateProviderClusterQueryImageParamsFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROtaSoftwareUpdateProviderClusterQueryImageParamsClass) Alloc() MTROtaSoftwareUpdateProviderClusterQueryImageParams {
	rv := objc.Send[MTROtaSoftwareUpdateProviderClusterQueryImageParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROtaSoftwareUpdateProviderClusterQueryImageParamsClass) New() MTROtaSoftwareUpdateProviderClusterQueryImageParams {
	rv := objc.Send[MTROtaSoftwareUpdateProviderClusterQueryImageParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageParams) Init() MTROtaSoftwareUpdateProviderClusterQueryImageParams {
	rv := objc.Send[MTROtaSoftwareUpdateProviderClusterQueryImageParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageParams) Autorelease() MTROtaSoftwareUpdateProviderClusterQueryImageParams {
	rv := objc.Send[MTROtaSoftwareUpdateProviderClusterQueryImageParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROtaSoftwareUpdateProviderClusterQueryImageParams creates a new MTROtaSoftwareUpdateProviderClusterQueryImageParams instance.
func NewMTROtaSoftwareUpdateProviderClusterQueryImageParams() MTROtaSoftwareUpdateProviderClusterQueryImageParams {
	return getMTROtaSoftwareUpdateProviderClusterQueryImageParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-zidv/hardwareversion
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageParams) HardwareVersion() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("hardwareVersion"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-zidv/hardwareversion
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageParams) SetHardwareVersion(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHardwareVersion:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-zidv/location
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageParams) Location() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("location"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-zidv/location
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageParams) SetLocation(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLocation:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-zidv/metadataforprovider
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageParams) MetadataForProvider() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("metadataForProvider"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-zidv/metadataforprovider
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageParams) SetMetadataForProvider(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMetadataForProvider:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-zidv/protocolssupported
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageParams) ProtocolsSupported() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("protocolsSupported"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-zidv/protocolssupported
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageParams) SetProtocolsSupported(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProtocolsSupported:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-zidv/requestorcanconsent
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageParams) RequestorCanConsent() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("requestorCanConsent"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-zidv/requestorcanconsent
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageParams) SetRequestorCanConsent(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRequestorCanConsent:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-zidv/serversideprocessingtimeout
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-zidv/serversideprocessingtimeout
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-zidv/softwareversion
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageParams) SoftwareVersion() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("softwareVersion"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-zidv/softwareversion
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageParams) SetSoftwareVersion(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSoftwareVersion:"), value)
}


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-zidv/timedinvoketimeoutms
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-zidv/timedinvoketimeoutms
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



