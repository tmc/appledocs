// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTROTASoftwareUpdateProviderClusterQueryImageResponseParams] class.
var (
	MTROTASoftwareUpdateProviderClusterQueryImageResponseParamsClass     _MTROTASoftwareUpdateProviderClusterQueryImageResponseParamsClass
	MTROTASoftwareUpdateProviderClusterQueryImageResponseParamsClassOnce sync.Once
)

func getMTROTASoftwareUpdateProviderClusterQueryImageResponseParamsClass() _MTROTASoftwareUpdateProviderClusterQueryImageResponseParamsClass {
	MTROTASoftwareUpdateProviderClusterQueryImageResponseParamsClassOnce.Do(func() {
		MTROTASoftwareUpdateProviderClusterQueryImageResponseParamsClass = _MTROTASoftwareUpdateProviderClusterQueryImageResponseParamsClass{objc.GetClass("MTROTASoftwareUpdateProviderClusterQueryImageResponseParams")}
	})
	return MTROTASoftwareUpdateProviderClusterQueryImageResponseParamsClass
}

type _MTROTASoftwareUpdateProviderClusterQueryImageResponseParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTROTASoftwareUpdateProviderClusterQueryImageResponseParams] class.
type IMTROTASoftwareUpdateProviderClusterQueryImageResponseParams interface {
	objectivec.IObject
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
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateProviderClusterQueryImageResponseParams-6wuvt
type MTROTASoftwareUpdateProviderClusterQueryImageResponseParams struct {
	objectivec.Object
}

// MTROTASoftwareUpdateProviderClusterQueryImageResponseParamsFrom constructs a [MTROTASoftwareUpdateProviderClusterQueryImageResponseParams] from an unsafe.Pointer.
func MTROTASoftwareUpdateProviderClusterQueryImageResponseParamsFrom(ptr unsafe.Pointer) MTROTASoftwareUpdateProviderClusterQueryImageResponseParams {
	return MTROTASoftwareUpdateProviderClusterQueryImageResponseParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROTASoftwareUpdateProviderClusterQueryImageResponseParamsClass) Alloc() MTROTASoftwareUpdateProviderClusterQueryImageResponseParams {
	rv := objc.Send[MTROTASoftwareUpdateProviderClusterQueryImageResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROTASoftwareUpdateProviderClusterQueryImageResponseParamsClass) New() MTROTASoftwareUpdateProviderClusterQueryImageResponseParams {
	rv := objc.Send[MTROTASoftwareUpdateProviderClusterQueryImageResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageResponseParams) Init() MTROTASoftwareUpdateProviderClusterQueryImageResponseParams {
	rv := objc.Send[MTROTASoftwareUpdateProviderClusterQueryImageResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageResponseParams) Autorelease() MTROTASoftwareUpdateProviderClusterQueryImageResponseParams {
	rv := objc.Send[MTROTASoftwareUpdateProviderClusterQueryImageResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROTASoftwareUpdateProviderClusterQueryImageResponseParams creates a new MTROTASoftwareUpdateProviderClusterQueryImageResponseParams instance.
func NewMTROTASoftwareUpdateProviderClusterQueryImageResponseParams() MTROTASoftwareUpdateProviderClusterQueryImageResponseParams {
	return getMTROTASoftwareUpdateProviderClusterQueryImageResponseParamsClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageresponseparams-6wuvt/delayedactiontime
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageResponseParams) DelayedActionTime() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("delayedActionTime"))
	return rv
}


// SetDelayedActionTime sets the value of the delayedActionTime property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageresponseparams-6wuvt/delayedactiontime
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageResponseParams) SetDelayedActionTime(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDelayedActionTime:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageresponseparams-6wuvt/imageuri
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageResponseParams) ImageURI() string {
	rv := objc.Send[string](m_.ID, objc.Sel("imageURI"))
	return rv
}


// SetImageURI sets the value of the imageURI property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageresponseparams-6wuvt/imageuri
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageResponseParams) SetImageURI(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setImageURI:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageresponseparams-6wuvt/metadataforrequestor
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageResponseParams) MetadataForRequestor() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("metadataForRequestor"))
	return rv
}


// SetMetadataForRequestor sets the value of the metadataForRequestor property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageresponseparams-6wuvt/metadataforrequestor
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageResponseParams) SetMetadataForRequestor(value foundation.IData) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMetadataForRequestor:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageresponseparams-6wuvt/softwareversion
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageResponseParams) SoftwareVersion() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("softwareVersion"))
	return rv
}


// SetSoftwareVersion sets the value of the softwareVersion property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageresponseparams-6wuvt/softwareversion
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageResponseParams) SetSoftwareVersion(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSoftwareVersion:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageresponseparams-6wuvt/softwareversionstring
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageResponseParams) SoftwareVersionString() string {
	rv := objc.Send[string](m_.ID, objc.Sel("softwareVersionString"))
	return rv
}


// SetSoftwareVersionString sets the value of the softwareVersionString property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageresponseparams-6wuvt/softwareversionstring
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageResponseParams) SetSoftwareVersionString(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSoftwareVersionString:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageresponseparams-6wuvt/status
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageResponseParams) Status() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("status"))
	return rv
}


// SetStatus sets the value of the status property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageresponseparams-6wuvt/status
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageResponseParams) SetStatus(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatus:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageresponseparams-6wuvt/timedinvoketimeoutms
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageResponseParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageresponseparams-6wuvt/timedinvoketimeoutms
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageResponseParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageresponseparams-6wuvt/updatetoken
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageResponseParams) UpdateToken() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("updateToken"))
	return rv
}


// SetUpdateToken sets the value of the updateToken property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageresponseparams-6wuvt/updatetoken
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageResponseParams) SetUpdateToken(value foundation.IData) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUpdateToken:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageresponseparams-6wuvt/userconsentneeded
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageResponseParams) UserConsentNeeded() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("userConsentNeeded"))
	return rv
}


// SetUserConsentNeeded sets the value of the userConsentNeeded property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageresponseparams-6wuvt/userconsentneeded
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageResponseParams) SetUserConsentNeeded(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserConsentNeeded:"), value)
}



