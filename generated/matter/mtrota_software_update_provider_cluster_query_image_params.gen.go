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
	HardwareVersion() foundation.Number
	SetHardwareVersion(value foundation.INumber)
	Location() string
	SetLocation(value string)
	MetadataForProvider() foundation.Data
	SetMetadataForProvider(value foundation.IData)
	ProductID() foundation.Number
	SetProductID(value foundation.INumber)
	ProductId() foundation.Number
	SetProductId(value foundation.INumber)
	ProtocolsSupported() unsafe.Pointer
	SetProtocolsSupported(value unsafe.Pointer)
	RequestorCanConsent() foundation.Number
	SetRequestorCanConsent(value foundation.INumber)
	ServerSideProcessingTimeout() foundation.Number
	SetServerSideProcessingTimeout(value foundation.INumber)
	SoftwareVersion() foundation.Number
	SetSoftwareVersion(value foundation.INumber)
	TimedInvokeTimeoutMs() foundation.Number
	SetTimedInvokeTimeoutMs(value foundation.INumber)
	VendorID() foundation.Number
	SetVendorID(value foundation.INumber)
	VendorId() foundation.Number
	SetVendorId(value foundation.INumber)
}

//
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


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-8z02b/hardwareversion
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageParams) HardwareVersion() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("hardwareVersion"))
	return rv
}


// SetHardwareVersion sets the value of the hardwareVersion property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-8z02b/hardwareversion
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageParams) SetHardwareVersion(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHardwareVersion:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-8z02b/location
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageParams) Location() string {
	rv := objc.Send[string](m_.ID, objc.Sel("location"))
	return rv
}


// SetLocation sets the value of the location property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-8z02b/location
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageParams) SetLocation(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLocation:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-8z02b/metadataforprovider
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageParams) MetadataForProvider() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("metadataForProvider"))
	return rv
}


// SetMetadataForProvider sets the value of the metadataForProvider property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-8z02b/metadataforprovider
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageParams) SetMetadataForProvider(value foundation.IData) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMetadataForProvider:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-8z02b/productid-9dqyi
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageParams) ProductID() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("productID"))
	return rv
}


// SetProductID sets the value of the productID property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-8z02b/productid-9dqyi
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageParams) SetProductID(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProductID:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-8z02b/productid-9dqxm
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageParams) ProductId() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("productId"))
	return rv
}


// SetProductId sets the value of the productId property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-8z02b/productid-9dqxm
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageParams) SetProductId(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProductId:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-8z02b/protocolssupported
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageParams) ProtocolsSupported() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("protocolsSupported"))
	return rv
}


// SetProtocolsSupported sets the value of the protocolsSupported property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-8z02b/protocolssupported
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageParams) SetProtocolsSupported(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProtocolsSupported:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-8z02b/requestorcanconsent
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageParams) RequestorCanConsent() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("requestorCanConsent"))
	return rv
}


// SetRequestorCanConsent sets the value of the requestorCanConsent property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-8z02b/requestorcanconsent
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageParams) SetRequestorCanConsent(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRequestorCanConsent:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-8z02b/serversideprocessingtimeout
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageParams) ServerSideProcessingTimeout() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}


// SetServerSideProcessingTimeout sets the value of the serverSideProcessingTimeout property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-8z02b/serversideprocessingtimeout
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageParams) SetServerSideProcessingTimeout(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-8z02b/softwareversion
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageParams) SoftwareVersion() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("softwareVersion"))
	return rv
}


// SetSoftwareVersion sets the value of the softwareVersion property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-8z02b/softwareversion
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageParams) SetSoftwareVersion(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSoftwareVersion:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-8z02b/timedinvoketimeoutms
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageParams) TimedInvokeTimeoutMs() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}


// SetTimedInvokeTimeoutMs sets the value of the timedInvokeTimeoutMs property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-8z02b/timedinvoketimeoutms
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageParams) SetTimedInvokeTimeoutMs(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-8z02b/vendorid-6cv55
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageParams) VendorID() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("vendorID"))
	return rv
}


// SetVendorID sets the value of the vendorID property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-8z02b/vendorid-6cv55
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageParams) SetVendorID(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVendorID:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-8z02b/vendorid-6cv49
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageParams) VendorId() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("vendorId"))
	return rv
}


// SetVendorId sets the value of the vendorId property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrotasoftwareupdateproviderclusterqueryimageparams-8z02b/vendorid-6cv49
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageParams) SetVendorId(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVendorId:"), value)
}



