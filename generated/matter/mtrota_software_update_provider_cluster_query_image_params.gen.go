// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTROTASoftwareUpdateProviderClusterQueryImageParams */


/* debug [class_header]: Header for MTROTASoftwareUpdateProviderClusterQueryImageParams */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTROTASoftwareUpdateProviderClusterQueryImageParams */
// An interface definition for the [MTROTASoftwareUpdateProviderClusterQueryImageParams] class.
type IMTROTASoftwareUpdateProviderClusterQueryImageParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTROTASoftwareUpdateProviderClusterQueryImageParams */
	// properties:
	HardwareVersion() objc.IObject /* cross-framework: NSNumber */
	SetHardwareVersion(value objc.IObject /* cross-framework: NSNumber */)
	Location() objc.IObject /* cross-framework: NSString */
	SetLocation(value objc.IObject /* cross-framework: NSString */)
	MetadataForProvider() objc.IObject /* cross-framework: NSData */
	SetMetadataForProvider(value objc.IObject /* cross-framework: NSData */)
	ProductId() objc.IObject /* cross-framework: NSNumber */
	SetProductId(value objc.IObject /* cross-framework: NSNumber */)
	ProductID() objc.IObject /* cross-framework: NSNumber */
	SetProductID(value objc.IObject /* cross-framework: NSNumber */)
	ProtocolsSupported() objc.IObject /* cross-framework: NSArray */
	SetProtocolsSupported(value objc.IObject /* cross-framework: NSArray */)
	RequestorCanConsent() objc.IObject /* cross-framework: NSNumber */
	SetRequestorCanConsent(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	SoftwareVersion() objc.IObject /* cross-framework: NSNumber */
	SetSoftwareVersion(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	VendorId() objc.IObject /* cross-framework: NSNumber */
	SetVendorId(value objc.IObject /* cross-framework: NSNumber */)
	VendorID() objc.IObject /* cross-framework: NSNumber */
	SetVendorID(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTROTASoftwareUpdateProviderClusterQueryImageParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTROTASoftwareUpdateProviderClusterQueryImageParams */
// Alloc allocates a new instance without initialization.
func (mc _MTROTASoftwareUpdateProviderClusterQueryImageParamsClass) Alloc() MTROTASoftwareUpdateProviderClusterQueryImageParams {
	rv := objc.Send[MTROTASoftwareUpdateProviderClusterQueryImageParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTROTASoftwareUpdateProviderClusterQueryImageParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateProviderClusterQueryImageParams-8z02b
type MTROTASoftwareUpdateProviderClusterQueryImageParams struct {
	objectivec.Object
}

// MTROTASoftwareUpdateProviderClusterQueryImageParamsFrom constructs a [MTROTASoftwareUpdateProviderClusterQueryImageParams] from an unsafe.Pointer.
func MTROTASoftwareUpdateProviderClusterQueryImageParamsFrom(ptr unsafe.Pointer) MTROTASoftwareUpdateProviderClusterQueryImageParams {
	return MTROTASoftwareUpdateProviderClusterQueryImageParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTROTASoftwareUpdateProviderClusterQueryImageParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTROTASoftwareUpdateProviderClusterQueryImageParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTROTASoftwareUpdateProviderClusterQueryImageParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTROTASoftwareUpdateProviderClusterQueryImageParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTROTASoftwareUpdateProviderClusterQueryImageParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateProviderClusterQueryImageParams-8z02b/hardwareVersion
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageParams) HardwareVersion() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("hardwareVersion"))
	return rv
}/* debug [instance_properties/getter]: hardwareVersion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateProviderClusterQueryImageParams-8z02b/hardwareVersion
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageParams) SetHardwareVersion(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHardwareVersion:"), value)
}/* debug [instance_properties/setter]: hardwareVersion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateProviderClusterQueryImageParams-8z02b/location
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageParams) Location() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("location"))
	return rv
}/* debug [instance_properties/getter]: location */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateProviderClusterQueryImageParams-8z02b/location
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageParams) SetLocation(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLocation:"), value)
}/* debug [instance_properties/setter]: location */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateProviderClusterQueryImageParams-8z02b/metadataForProvider
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageParams) MetadataForProvider() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("metadataForProvider"))
	return rv
}/* debug [instance_properties/getter]: metadataForProvider */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateProviderClusterQueryImageParams-8z02b/metadataForProvider
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageParams) SetMetadataForProvider(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMetadataForProvider:"), value)
}/* debug [instance_properties/setter]: metadataForProvider */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateProviderClusterQueryImageParams-8z02b/productId-9dqxm
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageParams) ProductId() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("productId"))
	return rv
}/* debug [instance_properties/getter]: productId */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateProviderClusterQueryImageParams-8z02b/productId-9dqxm
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageParams) SetProductId(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProductId:"), value)
}/* debug [instance_properties/setter]: productId */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateProviderClusterQueryImageParams-8z02b/productID-9dqyi
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageParams) ProductID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("productID"))
	return rv
}/* debug [instance_properties/getter]: productID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateProviderClusterQueryImageParams-8z02b/productID-9dqyi
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageParams) SetProductID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProductID:"), value)
}/* debug [instance_properties/setter]: productID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateProviderClusterQueryImageParams-8z02b/protocolsSupported
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageParams) ProtocolsSupported() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("protocolsSupported"))
	return rv
}/* debug [instance_properties/getter]: protocolsSupported */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateProviderClusterQueryImageParams-8z02b/protocolsSupported
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageParams) SetProtocolsSupported(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProtocolsSupported:"), value)
}/* debug [instance_properties/setter]: protocolsSupported */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateProviderClusterQueryImageParams-8z02b/requestorCanConsent
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageParams) RequestorCanConsent() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("requestorCanConsent"))
	return rv
}/* debug [instance_properties/getter]: requestorCanConsent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateProviderClusterQueryImageParams-8z02b/requestorCanConsent
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageParams) SetRequestorCanConsent(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRequestorCanConsent:"), value)
}/* debug [instance_properties/setter]: requestorCanConsent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateProviderClusterQueryImageParams-8z02b/serverSideProcessingTimeout
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateProviderClusterQueryImageParams-8z02b/serverSideProcessingTimeout
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateProviderClusterQueryImageParams-8z02b/softwareVersion
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageParams) SoftwareVersion() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("softwareVersion"))
	return rv
}/* debug [instance_properties/getter]: softwareVersion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateProviderClusterQueryImageParams-8z02b/softwareVersion
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageParams) SetSoftwareVersion(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSoftwareVersion:"), value)
}/* debug [instance_properties/setter]: softwareVersion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateProviderClusterQueryImageParams-8z02b/timedInvokeTimeoutMs
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateProviderClusterQueryImageParams-8z02b/timedInvokeTimeoutMs
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateProviderClusterQueryImageParams-8z02b/vendorId-6cv49
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageParams) VendorId() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("vendorId"))
	return rv
}/* debug [instance_properties/getter]: vendorId */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateProviderClusterQueryImageParams-8z02b/vendorId-6cv49
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageParams) SetVendorId(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVendorId:"), value)
}/* debug [instance_properties/setter]: vendorId */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateProviderClusterQueryImageParams-8z02b/vendorID-6cv55
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageParams) VendorID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("vendorID"))
	return rv
}/* debug [instance_properties/getter]: vendorID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateProviderClusterQueryImageParams-8z02b/vendorID-6cv55
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageParams) SetVendorID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVendorID:"), value)
}/* debug [instance_properties/setter]: vendorID */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTROTASoftwareUpdateProviderClusterQueryImageParams */



