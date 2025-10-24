// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams */


/* debug [class_header]: Header for MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams */
// An interface definition for the [MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams] class.
type IMTROtaSoftwareUpdateProviderClusterQueryImageResponseParams interface {
	IMTROTASoftwareUpdateProviderClusterQueryImageResponseParams
	
/* debug [class_interface_properties]: Properties for MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams */
	// properties:
	DelayedActionTime() objc.IObject /* cross-framework: NSNumber */
	SetDelayedActionTime(value objc.IObject /* cross-framework: NSNumber */)
	ImageURI() objc.IObject /* cross-framework: NSString */
	SetImageURI(value objc.IObject /* cross-framework: NSString */)
	MetadataForRequestor() objc.IObject /* cross-framework: NSData */
	SetMetadataForRequestor(value objc.IObject /* cross-framework: NSData */)
	SoftwareVersion() objc.IObject /* cross-framework: NSNumber */
	SetSoftwareVersion(value objc.IObject /* cross-framework: NSNumber */)
	SoftwareVersionString() objc.IObject /* cross-framework: NSString */
	SetSoftwareVersionString(value objc.IObject /* cross-framework: NSString */)
	Status() objc.IObject /* cross-framework: NSNumber */
	SetStatus(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	UpdateToken() objc.IObject /* cross-framework: NSData */
	SetUpdateToken(value objc.IObject /* cross-framework: NSData */)
	UserConsentNeeded() objc.IObject /* cross-framework: NSNumber */
	SetUserConsentNeeded(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams */
// Alloc allocates a new instance without initialization.
func (mc _MTROtaSoftwareUpdateProviderClusterQueryImageResponseParamsClass) Alloc() MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams {
	rv := objc.Send[MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams */


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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams-48oao/delayedActionTime
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams) DelayedActionTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("delayedActionTime"))
	return rv
}/* debug [instance_properties/getter]: delayedActionTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams-48oao/delayedActionTime
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams) SetDelayedActionTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDelayedActionTime:"), value)
}/* debug [instance_properties/setter]: delayedActionTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams-48oao/imageURI
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams) ImageURI() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("imageURI"))
	return rv
}/* debug [instance_properties/getter]: imageURI */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams-48oao/imageURI
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams) SetImageURI(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setImageURI:"), value)
}/* debug [instance_properties/setter]: imageURI */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams-48oao/metadataForRequestor
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams) MetadataForRequestor() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("metadataForRequestor"))
	return rv
}/* debug [instance_properties/getter]: metadataForRequestor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams-48oao/metadataForRequestor
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams) SetMetadataForRequestor(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMetadataForRequestor:"), value)
}/* debug [instance_properties/setter]: metadataForRequestor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams-48oao/softwareVersion
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams) SoftwareVersion() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("softwareVersion"))
	return rv
}/* debug [instance_properties/getter]: softwareVersion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams-48oao/softwareVersion
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams) SetSoftwareVersion(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSoftwareVersion:"), value)
}/* debug [instance_properties/setter]: softwareVersion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams-48oao/softwareVersionString
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams) SoftwareVersionString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("softwareVersionString"))
	return rv
}/* debug [instance_properties/getter]: softwareVersionString */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams-48oao/softwareVersionString
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams) SetSoftwareVersionString(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSoftwareVersionString:"), value)
}/* debug [instance_properties/setter]: softwareVersionString */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams-48oao/status
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams) Status() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("status"))
	return rv
}/* debug [instance_properties/getter]: status */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams-48oao/status
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams) SetStatus(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatus:"), value)
}/* debug [instance_properties/setter]: status */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams-48oao/timedInvokeTimeoutMs
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams-48oao/timedInvokeTimeoutMs
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams-48oao/updateToken
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams) UpdateToken() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("updateToken"))
	return rv
}/* debug [instance_properties/getter]: updateToken */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams-48oao/updateToken
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams) SetUpdateToken(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUpdateToken:"), value)
}/* debug [instance_properties/setter]: updateToken */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams-48oao/userConsentNeeded
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams) UserConsentNeeded() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("userConsentNeeded"))
	return rv
}/* debug [instance_properties/getter]: userConsentNeeded */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams-48oao/userConsentNeeded
func (m_ MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams) SetUserConsentNeeded(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserConsentNeeded:"), value)
}/* debug [instance_properties/setter]: userConsentNeeded */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTROtaSoftwareUpdateProviderClusterQueryImageResponseParams */



