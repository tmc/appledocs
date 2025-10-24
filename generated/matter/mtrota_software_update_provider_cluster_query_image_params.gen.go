// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTROTASoftwareUpdateProviderClusterQueryImageParams] class.
var (
	MTROTASoftwareUpdateProviderClusterQueryImageParamsClass     _MTROTASoftwareUpdateProviderClusterQueryImageParamsClass
	MTROTASoftwareUpdateProviderClusterQueryImageParamsClassOnce sync.Once
)

func getMTROTASoftwareUpdateProviderClusterQueryImageParamsClass() _MTROTASoftwareUpdateProviderClusterQueryImageParamsClass {
	MTROTASoftwareUpdateProviderClusterQueryImageParamsClassOnce.Do(func() {
		MTROTASoftwareUpdateProviderClusterQueryImageParamsClass = _MTROTASoftwareUpdateProviderClusterQueryImageParamsClass{objc.GetClass("MTROTASoftwareUpdateProviderClusterQueryImageParams")}
	})
	return MTROTASoftwareUpdateProviderClusterQueryImageParamsClass
}

type _MTROTASoftwareUpdateProviderClusterQueryImageParamsClass struct {
	class objc.Class
}

// An interface definition for the [MTROTASoftwareUpdateProviderClusterQueryImageParams] class.
type IMTROTASoftwareUpdateProviderClusterQueryImageParams interface {
	objectivec.IObject
	// properties:
	HardwareVersion() objc.IObject /* cross-framework: NSNumber */
	SetHardwareVersion(value objc.IObject /* cross-framework: NSNumber */)
	Location() objc.IObject /* cross-framework: NSString */
	SetLocation(value objc.IObject /* cross-framework: NSString */)
	MetadataForProvider() objc.IObject /* cross-framework: Data */
	SetMetadataForProvider(value objc.IObject /* cross-framework: Data */)
	ProductID() objc.IObject /* cross-framework: NSNumber */
	SetProductID(value objc.IObject /* cross-framework: NSNumber */)
	ProductId() objc.IObject /* cross-framework: NSNumber */
	SetProductId(value objc.IObject /* cross-framework: NSNumber */)
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
	VendorID() objc.IObject /* cross-framework: NSNumber */
	SetVendorID(value objc.IObject /* cross-framework: NSNumber */)
	VendorId() objc.IObject /* cross-framework: NSNumber */
	SetVendorId(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateProviderClusterQueryImageParams-8z02b
type MTROTASoftwareUpdateProviderClusterQueryImageParams struct {
	objectivec.Object
}

// MTROTASoftwareUpdateProviderClusterQueryImageParamsFrom constructs a [MTROTASoftwareUpdateProviderClusterQueryImageParams] from an unsafe.Pointer.
func MTROTASoftwareUpdateProviderClusterQueryImageParamsFrom(ptr unsafe.Pointer) MTROTASoftwareUpdateProviderClusterQueryImageParams {
	return MTROTASoftwareUpdateProviderClusterQueryImageParams{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROTASoftwareUpdateProviderClusterQueryImageParamsClass) Alloc() MTROTASoftwareUpdateProviderClusterQueryImageParams {
	rv := objc.Send[MTROTASoftwareUpdateProviderClusterQueryImageParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROTASoftwareUpdateProviderClusterQueryImageParamsClass) New() MTROTASoftwareUpdateProviderClusterQueryImageParams {
	rv := objc.Send[MTROTASoftwareUpdateProviderClusterQueryImageParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageParams) Init() MTROTASoftwareUpdateProviderClusterQueryImageParams {
	rv := objc.Send[MTROTASoftwareUpdateProviderClusterQueryImageParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageParams) Autorelease() MTROTASoftwareUpdateProviderClusterQueryImageParams {
	rv := objc.Send[MTROTASoftwareUpdateProviderClusterQueryImageParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROTASoftwareUpdateProviderClusterQueryImageParams creates a new MTROTASoftwareUpdateProviderClusterQueryImageParams instance.
func NewMTROTASoftwareUpdateProviderClusterQueryImageParams() MTROTASoftwareUpdateProviderClusterQueryImageParams {
	return getMTROTASoftwareUpdateProviderClusterQueryImageParamsClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-8z02b/hardwareversion
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageParams) HardwareVersion() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("hardwareVersion"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-8z02b/hardwareversion
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageParams) SetHardwareVersion(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHardwareVersion:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-8z02b/location
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageParams) Location() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("location"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-8z02b/location
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageParams) SetLocation(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLocation:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-8z02b/metadataforprovider
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageParams) MetadataForProvider() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("metadataForProvider"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-8z02b/metadataforprovider
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageParams) SetMetadataForProvider(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMetadataForProvider:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-8z02b/productid-9dqyi
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageParams) ProductID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("productID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-8z02b/productid-9dqyi
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageParams) SetProductID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProductID:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-8z02b/productid-9dqxm
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageParams) ProductId() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("productId"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-8z02b/productid-9dqxm
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageParams) SetProductId(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProductId:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-8z02b/protocolssupported
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageParams) ProtocolsSupported() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("protocolsSupported"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-8z02b/protocolssupported
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageParams) SetProtocolsSupported(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProtocolsSupported:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-8z02b/requestorcanconsent
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageParams) RequestorCanConsent() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("requestorCanConsent"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-8z02b/requestorcanconsent
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageParams) SetRequestorCanConsent(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRequestorCanConsent:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-8z02b/serversideprocessingtimeout
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-8z02b/serversideprocessingtimeout
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-8z02b/softwareversion
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageParams) SoftwareVersion() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("softwareVersion"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-8z02b/softwareversion
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageParams) SetSoftwareVersion(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSoftwareVersion:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-8z02b/timedinvoketimeoutms
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-8z02b/timedinvoketimeoutms
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-8z02b/vendorid-6cv55
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageParams) VendorID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("vendorID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-8z02b/vendorid-6cv55
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageParams) SetVendorID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVendorID:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-8z02b/vendorid-6cv49
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageParams) VendorId() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("vendorId"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-8z02b/vendorid-6cv49
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageParams) SetVendorId(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVendorId:"), value)
}



