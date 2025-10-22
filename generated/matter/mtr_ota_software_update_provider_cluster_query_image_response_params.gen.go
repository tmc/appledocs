// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams] class.
var (
	MTROtaSoftwareUpdateProviderClusterQueryImageResponseParamsClass     _MTROtaSoftwareUpdateProviderClusterQueryImageResponseParamsClass
	MTROtaSoftwareUpdateProviderClusterQueryImageResponseParamsClassOnce sync.Once
)

func getMTROtaSoftwareUpdateProviderClusterQueryImageResponseParamsClass() _MTROtaSoftwareUpdateProviderClusterQueryImageResponseParamsClass {
	MTROtaSoftwareUpdateProviderClusterQueryImageResponseParamsClassOnce.Do(func() {
		MTROtaSoftwareUpdateProviderClusterQueryImageResponseParamsClass = _MTROtaSoftwareUpdateProviderClusterQueryImageResponseParamsClass{objc.GetClass("MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams")}
	})
	return MTROtaSoftwareUpdateProviderClusterQueryImageResponseParamsClass
}

type _MTROtaSoftwareUpdateProviderClusterQueryImageResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams] class.
type IMTROtaSoftwareUpdateProviderClusterQueryImageResponseParams interface {
	IMTROTASoftwareUpdateProviderClusterQueryImageResponseParams
	DelayedActionTime() foundation.Number
	SetDelayedActionTime(value foundation.INumber)
	ImageURI() string
	SetImageURI(value string)
	MetadataForRequestor() foundation.Data
	SetMetadataForRequestor(value foundation.IData)
	SoftwareVersion() foundation.Number
	SetSoftwareVersion(value foundation.INumber)
	SoftwareVersionString() string
	SetSoftwareVersionString(value string)
	Status() foundation.Number
	SetStatus(value foundation.INumber)
	TimedInvokeTimeoutMs() foundation.Number
	SetTimedInvokeTimeoutMs(value foundation.INumber)
	UpdateToken() foundation.Data
	SetUpdateToken(value foundation.IData)
	UserConsentNeeded() foundation.Number
	SetUserConsentNeeded(value foundation.INumber)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams-48oao
type MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams struct {
	MTROTASoftwareUpdateProviderClusterQueryImageResponseParams
}

// MTROtaSoftwareUpdateProviderClusterQueryImageResponseParamsFrom constructs a [MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams] from an unsafe.Pointer.
func MTROtaSoftwareUpdateProviderClusterQueryImageResponseParamsFrom(ptr unsafe.Pointer) MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams {
	return MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams{
		MTROTASoftwareUpdateProviderClusterQueryImageResponseParams: MTROTASoftwareUpdateProviderClusterQueryImageResponseParamsFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROtaSoftwareUpdateProviderClusterQueryImageResponseParamsClass) Alloc() MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams {
	rv := objc.Send[MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROtaSoftwareUpdateProviderClusterQueryImageResponseParamsClass) New() MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams {
	rv := objc.Send[MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams) Init() MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams {
	rv := objc.Send[MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams) Autorelease() MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams {
	rv := objc.Send[MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROtaSoftwareUpdateProviderClusterQueryImageResponseParams creates a new MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams instance.
func NewMTROtaSoftwareUpdateProviderClusterQueryImageResponseParams() MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams {
	return getMTROtaSoftwareUpdateProviderClusterQueryImageResponseParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageresponseparams-48oao/delayedactiontime
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams) DelayedActionTime() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("delayedActionTime"))
	return rv
}


// SetDelayedActionTime sets the value of the delayedActionTime property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageresponseparams-48oao/delayedactiontime
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams) SetDelayedActionTime(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDelayedActionTime:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageresponseparams-48oao/imageuri
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams) ImageURI() string {
	rv := objc.Send[string](m_.ID, objc.Sel("imageURI"))
	return rv
}


// SetImageURI sets the value of the imageURI property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageresponseparams-48oao/imageuri
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams) SetImageURI(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setImageURI:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageresponseparams-48oao/metadataforrequestor
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams) MetadataForRequestor() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("metadataForRequestor"))
	return rv
}


// SetMetadataForRequestor sets the value of the metadataForRequestor property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageresponseparams-48oao/metadataforrequestor
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams) SetMetadataForRequestor(value foundation.IData) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMetadataForRequestor:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageresponseparams-48oao/softwareversion
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams) SoftwareVersion() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("softwareVersion"))
	return rv
}


// SetSoftwareVersion sets the value of the softwareVersion property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageresponseparams-48oao/softwareversion
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams) SetSoftwareVersion(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSoftwareVersion:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageresponseparams-48oao/softwareversionstring
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams) SoftwareVersionString() string {
	rv := objc.Send[string](m_.ID, objc.Sel("softwareVersionString"))
	return rv
}


// SetSoftwareVersionString sets the value of the softwareVersionString property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageresponseparams-48oao/softwareversionstring
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams) SetSoftwareVersionString(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSoftwareVersionString:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageresponseparams-48oao/status
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams) Status() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("status"))
	return rv
}


// SetStatus sets the value of the status property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageresponseparams-48oao/status
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams) SetStatus(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatus:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageresponseparams-48oao/timedinvoketimeoutms
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageresponseparams-48oao/timedinvoketimeoutms
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageresponseparams-48oao/updatetoken
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams) UpdateToken() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("updateToken"))
	return rv
}


// SetUpdateToken sets the value of the updateToken property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageresponseparams-48oao/updatetoken
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams) SetUpdateToken(value foundation.IData) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUpdateToken:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageresponseparams-48oao/userconsentneeded
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams) UserConsentNeeded() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("userConsentNeeded"))
	return rv
}


// SetUserConsentNeeded sets the value of the userConsentNeeded property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageresponseparams-48oao/userconsentneeded
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams) SetUserConsentNeeded(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserConsentNeeded:"), value)
}



