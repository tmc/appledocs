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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageresponseparams-6wuvt/delayedactiontime
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageResponseParams) DelayedActionTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("delayedActionTime"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageresponseparams-6wuvt/delayedactiontime
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageResponseParams) SetDelayedActionTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDelayedActionTime:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageresponseparams-6wuvt/imageuri
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageResponseParams) ImageURI() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("imageURI"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageresponseparams-6wuvt/imageuri
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageResponseParams) SetImageURI(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setImageURI:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageresponseparams-6wuvt/metadataforrequestor
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageResponseParams) MetadataForRequestor() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("metadataForRequestor"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageresponseparams-6wuvt/metadataforrequestor
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageResponseParams) SetMetadataForRequestor(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMetadataForRequestor:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageresponseparams-6wuvt/softwareversion
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageResponseParams) SoftwareVersion() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("softwareVersion"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageresponseparams-6wuvt/softwareversion
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageResponseParams) SetSoftwareVersion(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSoftwareVersion:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageresponseparams-6wuvt/softwareversionstring
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageResponseParams) SoftwareVersionString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("softwareVersionString"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageresponseparams-6wuvt/softwareversionstring
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageResponseParams) SetSoftwareVersionString(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSoftwareVersionString:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageresponseparams-6wuvt/status
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageResponseParams) Status() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("status"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageresponseparams-6wuvt/status
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageResponseParams) SetStatus(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatus:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageresponseparams-6wuvt/timedinvoketimeoutms
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageResponseParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageresponseparams-6wuvt/timedinvoketimeoutms
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageResponseParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageresponseparams-6wuvt/updatetoken
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageResponseParams) UpdateToken() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("updateToken"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageresponseparams-6wuvt/updatetoken
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageResponseParams) SetUpdateToken(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUpdateToken:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageresponseparams-6wuvt/userconsentneeded
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageResponseParams) UserConsentNeeded() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("userConsentNeeded"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageresponseparams-6wuvt/userconsentneeded
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageResponseParams) SetUserConsentNeeded(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserConsentNeeded:"), value)
}



