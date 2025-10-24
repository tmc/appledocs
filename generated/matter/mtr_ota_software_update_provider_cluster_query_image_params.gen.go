// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTROtaSoftwareUpdateProviderClusterQueryImageParams */


/* debug [class_header]: Header for MTROtaSoftwareUpdateProviderClusterQueryImageParams */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTROtaSoftwareUpdateProviderClusterQueryImageParams */
// An interface definition for the [MTROtaSoftwareUpdateProviderClusterQueryImageParams] class.
type IMTROtaSoftwareUpdateProviderClusterQueryImageParams interface {
	IMTROTASoftwareUpdateProviderClusterQueryImageParams
	
/* debug [class_interface_properties]: Properties for MTROtaSoftwareUpdateProviderClusterQueryImageParams */
	// properties:
	HardwareVersion() objc.IObject /* cross-framework: NSNumber */
	SetHardwareVersion(value objc.IObject /* cross-framework: NSNumber */)
	Location() objc.IObject /* cross-framework: NSString */
	SetLocation(value objc.IObject /* cross-framework: NSString */)
	MetadataForProvider() objc.IObject /* cross-framework: NSData */
	SetMetadataForProvider(value objc.IObject /* cross-framework: NSData */)
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
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTROtaSoftwareUpdateProviderClusterQueryImageParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTROtaSoftwareUpdateProviderClusterQueryImageParams */
// Alloc allocates a new instance without initialization.
func (mc _MTROtaSoftwareUpdateProviderClusterQueryImageParamsClass) Alloc() MTROtaSoftwareUpdateProviderClusterQueryImageParams {
	rv := objc.Send[MTROtaSoftwareUpdateProviderClusterQueryImageParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTROtaSoftwareUpdateProviderClusterQueryImageParams */


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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTROtaSoftwareUpdateProviderClusterQueryImageParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTROtaSoftwareUpdateProviderClusterQueryImageParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTROtaSoftwareUpdateProviderClusterQueryImageParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTROtaSoftwareUpdateProviderClusterQueryImageParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTROtaSoftwareUpdateProviderClusterQueryImageParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateProviderClusterQueryImageParams-zidv/hardwareVersion
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageParams) HardwareVersion() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("hardwareVersion"))
	return rv
}/* debug [instance_properties/getter]: hardwareVersion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateProviderClusterQueryImageParams-zidv/hardwareVersion
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageParams) SetHardwareVersion(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHardwareVersion:"), value)
}/* debug [instance_properties/setter]: hardwareVersion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateProviderClusterQueryImageParams-zidv/location
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageParams) Location() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("location"))
	return rv
}/* debug [instance_properties/getter]: location */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateProviderClusterQueryImageParams-zidv/location
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageParams) SetLocation(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLocation:"), value)
}/* debug [instance_properties/setter]: location */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateProviderClusterQueryImageParams-zidv/metadataForProvider
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageParams) MetadataForProvider() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("metadataForProvider"))
	return rv
}/* debug [instance_properties/getter]: metadataForProvider */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateProviderClusterQueryImageParams-zidv/metadataForProvider
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageParams) SetMetadataForProvider(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMetadataForProvider:"), value)
}/* debug [instance_properties/setter]: metadataForProvider */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateProviderClusterQueryImageParams-zidv/protocolsSupported
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageParams) ProtocolsSupported() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("protocolsSupported"))
	return rv
}/* debug [instance_properties/getter]: protocolsSupported */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateProviderClusterQueryImageParams-zidv/protocolsSupported
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageParams) SetProtocolsSupported(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProtocolsSupported:"), value)
}/* debug [instance_properties/setter]: protocolsSupported */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateProviderClusterQueryImageParams-zidv/requestorCanConsent
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageParams) RequestorCanConsent() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("requestorCanConsent"))
	return rv
}/* debug [instance_properties/getter]: requestorCanConsent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateProviderClusterQueryImageParams-zidv/requestorCanConsent
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageParams) SetRequestorCanConsent(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRequestorCanConsent:"), value)
}/* debug [instance_properties/setter]: requestorCanConsent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateProviderClusterQueryImageParams-zidv/serverSideProcessingTimeout
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateProviderClusterQueryImageParams-zidv/serverSideProcessingTimeout
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateProviderClusterQueryImageParams-zidv/softwareVersion
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageParams) SoftwareVersion() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("softwareVersion"))
	return rv
}/* debug [instance_properties/getter]: softwareVersion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateProviderClusterQueryImageParams-zidv/softwareVersion
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageParams) SetSoftwareVersion(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSoftwareVersion:"), value)
}/* debug [instance_properties/setter]: softwareVersion */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateProviderClusterQueryImageParams-zidv/timedInvokeTimeoutMs
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateProviderClusterQueryImageParams-zidv/timedInvokeTimeoutMs
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTROtaSoftwareUpdateProviderClusterQueryImageParams */



