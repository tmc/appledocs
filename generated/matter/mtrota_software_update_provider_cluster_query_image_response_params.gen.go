// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTROTASoftwareUpdateProviderClusterQueryImageResponseParams */


/* debug [class_header]: Header for MTROTASoftwareUpdateProviderClusterQueryImageResponseParams */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTROTASoftwareUpdateProviderClusterQueryImageResponseParams */
// An interface definition for the [MTROTASoftwareUpdateProviderClusterQueryImageResponseParams] class.
type IMTROTASoftwareUpdateProviderClusterQueryImageResponseParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTROTASoftwareUpdateProviderClusterQueryImageResponseParams */
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

	
/* debug [class_interface_methods]: Methods for MTROTASoftwareUpdateProviderClusterQueryImageResponseParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTROTASoftwareUpdateProviderClusterQueryImageResponseParams */
// Alloc allocates a new instance without initialization.
func (mc _MTROTASoftwareUpdateProviderClusterQueryImageResponseParamsClass) Alloc() MTROTASoftwareUpdateProviderClusterQueryImageResponseParams {
	rv := objc.Send[MTROTASoftwareUpdateProviderClusterQueryImageResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTROTASoftwareUpdateProviderClusterQueryImageResponseParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateProviderClusterQueryImageResponseParams-6wuvt
type MTROTASoftwareUpdateProviderClusterQueryImageResponseParams struct {
	objectivec.Object
}

// MTROTASoftwareUpdateProviderClusterQueryImageResponseParamsFrom constructs a [MTROTASoftwareUpdateProviderClusterQueryImageResponseParams] from an unsafe.Pointer.
func MTROTASoftwareUpdateProviderClusterQueryImageResponseParamsFrom(ptr unsafe.Pointer) MTROTASoftwareUpdateProviderClusterQueryImageResponseParams {
	return MTROTASoftwareUpdateProviderClusterQueryImageResponseParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTROTASoftwareUpdateProviderClusterQueryImageResponseParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateProviderClusterQueryImageResponseParams-6wuvt/init(responseValue:)
func NewMTROTASoftwareUpdateProviderClusterQueryImageResponseParamsWithResponseValueError(responseValue foundation.IDictionary, error_ unsafe.Pointer) MTROTASoftwareUpdateProviderClusterQueryImageResponseParams {
	instance := getMTROTASoftwareUpdateProviderClusterQueryImageResponseParamsClass().Alloc()
	rv := objc.Send[MTROTASoftwareUpdateProviderClusterQueryImageResponseParams](instance.ID, objc.Sel("initWithResponseValue:error:"), responseValue, error_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTROTASoftwareUpdateProviderClusterQueryImageResponseParamsWithResponseValueError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTROTASoftwareUpdateProviderClusterQueryImageResponseParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTROTASoftwareUpdateProviderClusterQueryImageResponseParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTROTASoftwareUpdateProviderClusterQueryImageResponseParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTROTASoftwareUpdateProviderClusterQueryImageResponseParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateProviderClusterQueryImageResponseParams-6wuvt/delayedActionTime
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageResponseParams) DelayedActionTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("delayedActionTime"))
	return rv
}/* debug [instance_properties/getter]: delayedActionTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateProviderClusterQueryImageResponseParams-6wuvt/delayedActionTime
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageResponseParams) SetDelayedActionTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDelayedActionTime:"), value)
}/* debug [instance_properties/setter]: delayedActionTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateProviderClusterQueryImageResponseParams-6wuvt/imageURI
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageResponseParams) ImageURI() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("imageURI"))
	return rv
}/* debug [instance_properties/getter]: imageURI */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateProviderClusterQueryImageResponseParams-6wuvt/imageURI
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageResponseParams) SetImageURI(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setImageURI:"), value)
}/* debug [instance_properties/setter]: imageURI */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateProviderClusterQueryImageResponseParams-6wuvt/metadataForRequestor
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageResponseParams) MetadataForRequestor() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("metadataForRequestor"))
	return rv
}/* debug [instance_properties/getter]: metadataForRequestor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateProviderClusterQueryImageResponseParams-6wuvt/metadataForRequestor
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageResponseParams) SetMetadataForRequestor(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMetadataForRequestor:"), value)
}/* debug [instance_properties/setter]: metadataForRequestor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateProviderClusterQueryImageResponseParams-6wuvt/softwareVersion
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageResponseParams) SoftwareVersion() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("softwareVersion"))
	return rv
}/* debug [instance_properties/getter]: softwareVersion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateProviderClusterQueryImageResponseParams-6wuvt/softwareVersion
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageResponseParams) SetSoftwareVersion(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSoftwareVersion:"), value)
}/* debug [instance_properties/setter]: softwareVersion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateProviderClusterQueryImageResponseParams-6wuvt/softwareVersionString
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageResponseParams) SoftwareVersionString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("softwareVersionString"))
	return rv
}/* debug [instance_properties/getter]: softwareVersionString */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateProviderClusterQueryImageResponseParams-6wuvt/softwareVersionString
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageResponseParams) SetSoftwareVersionString(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSoftwareVersionString:"), value)
}/* debug [instance_properties/setter]: softwareVersionString */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateProviderClusterQueryImageResponseParams-6wuvt/status
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageResponseParams) Status() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("status"))
	return rv
}/* debug [instance_properties/getter]: status */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateProviderClusterQueryImageResponseParams-6wuvt/status
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageResponseParams) SetStatus(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatus:"), value)
}/* debug [instance_properties/setter]: status */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateProviderClusterQueryImageResponseParams-6wuvt/timedInvokeTimeoutMs
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageResponseParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateProviderClusterQueryImageResponseParams-6wuvt/timedInvokeTimeoutMs
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageResponseParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateProviderClusterQueryImageResponseParams-6wuvt/updateToken
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageResponseParams) UpdateToken() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("updateToken"))
	return rv
}/* debug [instance_properties/getter]: updateToken */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateProviderClusterQueryImageResponseParams-6wuvt/updateToken
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageResponseParams) SetUpdateToken(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUpdateToken:"), value)
}/* debug [instance_properties/setter]: updateToken */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateProviderClusterQueryImageResponseParams-6wuvt/userConsentNeeded
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageResponseParams) UserConsentNeeded() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("userConsentNeeded"))
	return rv
}/* debug [instance_properties/getter]: userConsentNeeded */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateProviderClusterQueryImageResponseParams-6wuvt/userConsentNeeded
func (m_ MTROTASoftwareUpdateProviderClusterQueryImageResponseParams) SetUserConsentNeeded(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setUserConsentNeeded:"), value)
}/* debug [instance_properties/setter]: userConsentNeeded */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTROTASoftwareUpdateProviderClusterQueryImageResponseParams */


