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
	// properties:
	DelayedActionTime() objc.IObject /* cross-framework: NSNumber */
	SetDelayedActionTime(value objc.IObject /* cross-framework: NSNumber */)
	ImageURI() objc.IObject /* cross-framework: NSString */
	SetImageURI(value objc.IObject /* cross-framework: NSString */)
	MetadataForRequestor() objc.IObject /* cross-framework: Data */
	SetMetadataForRequestor(value objc.IObject /* cross-framework: Data */)
	SoftwareVersion() objc.IObject /* cross-framework: NSNumber */
	SetSoftwareVersion(value objc.IObject /* cross-framework: NSNumber */)
	SoftwareVersionString() objc.IObject /* cross-framework: NSString */
	SetSoftwareVersionString(value objc.IObject /* cross-framework: NSString */)
	Status() objc.IObject /* cross-framework: NSNumber */
	SetStatus(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	UpdateToken() objc.IObject /* cross-framework: Data */
	SetUpdateToken(value objc.IObject /* cross-framework: Data */)
	UserConsentNeeded() objc.IObject /* cross-framework: NSNumber */
	SetUserConsentNeeded(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageresponseparams-48oao/delayedactiontime
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams) DelayedActionTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("delayedActionTime"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageresponseparams-48oao/delayedactiontime
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams) SetDelayedActionTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDelayedActionTime:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageresponseparams-48oao/imageuri
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams) ImageURI() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("imageURI"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageresponseparams-48oao/imageuri
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams) SetImageURI(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setImageURI:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageresponseparams-48oao/metadataforrequestor
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams) MetadataForRequestor() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("metadataForRequestor"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageresponseparams-48oao/metadataforrequestor
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams) SetMetadataForRequestor(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMetadataForRequestor:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageresponseparams-48oao/softwareversion
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams) SoftwareVersion() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("softwareVersion"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageresponseparams-48oao/softwareversion
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams) SetSoftwareVersion(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSoftwareVersion:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageresponseparams-48oao/softwareversionstring
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams) SoftwareVersionString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("softwareVersionString"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageresponseparams-48oao/softwareversionstring
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams) SetSoftwareVersionString(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSoftwareVersionString:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageresponseparams-48oao/status
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams) Status() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("status"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageresponseparams-48oao/status
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams) SetStatus(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatus:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageresponseparams-48oao/timedinvoketimeoutms
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageresponseparams-48oao/timedinvoketimeoutms
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageresponseparams-48oao/updatetoken
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams) UpdateToken() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("updateToken"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageresponseparams-48oao/updatetoken
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams) SetUpdateToken(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUpdateToken:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageresponseparams-48oao/userconsentneeded
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams) UserConsentNeeded() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("userConsentNeeded"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageresponseparams-48oao/userconsentneeded
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams) SetUserConsentNeeded(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserConsentNeeded:"), value)
}



