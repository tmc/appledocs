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
}

//
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


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-zidv/hardwareversion
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageParams) HardwareVersion() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("hardwareVersion"))
	return rv
}


// SetHardwareVersion sets the value of the hardwareVersion property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-zidv/hardwareversion
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageParams) SetHardwareVersion(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHardwareVersion:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-zidv/location
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageParams) Location() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("location"))
	return rv
}


// SetLocation sets the value of the location property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-zidv/location
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageParams) SetLocation(value appkit.string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLocation:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-zidv/metadataforprovider
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageParams) MetadataForProvider() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("metadataForProvider"))
	return rv
}


// SetMetadataForProvider sets the value of the metadataForProvider property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-zidv/metadataforprovider
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageParams) SetMetadataForProvider(value foundation.IData) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMetadataForProvider:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-zidv/protocolssupported
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageParams) ProtocolsSupported() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("protocolsSupported"))
	return rv
}


// SetProtocolsSupported sets the value of the protocolsSupported property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-zidv/protocolssupported
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageParams) SetProtocolsSupported(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProtocolsSupported:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-zidv/requestorcanconsent
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageParams) RequestorCanConsent() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("requestorCanConsent"))
	return rv
}


// SetRequestorCanConsent sets the value of the requestorCanConsent property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-zidv/requestorcanconsent
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageParams) SetRequestorCanConsent(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRequestorCanConsent:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-zidv/serversideprocessingtimeout
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-zidv/serversideprocessingtimeout
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-zidv/softwareversion
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageParams) SoftwareVersion() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("softwareVersion"))
	return rv
}


// SetSoftwareVersion sets the value of the softwareVersion property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-zidv/softwareversion
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageParams) SetSoftwareVersion(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSoftwareVersion:"), value)
}

// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-zidv/timedinvoketimeoutms
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
// Controls whether the command is a timed command (using Timed Invoke).

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-zidv/timedinvoketimeoutms
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}



