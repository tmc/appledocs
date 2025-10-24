// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRDeviceControllerFactoryParams */


/* debug [class_header]: Header for MTRDeviceControllerFactoryParams */
// The class instance for the [MTRDeviceControllerFactoryParams] class.
var (
	MTRDeviceControllerFactoryParamsClass     _MTRDeviceControllerFactoryParamsClass
	MTRDeviceControllerFactoryParamsClassOnce sync.Once
)

func getMTRDeviceControllerFactoryParamsClass() _MTRDeviceControllerFactoryParamsClass {
	MTRDeviceControllerFactoryParamsClassOnce.Do(func() {
		MTRDeviceControllerFactoryParamsClass = _MTRDeviceControllerFactoryParamsClass{objc.GetClass("MTRDeviceControllerFactoryParams")}
	})
	return MTRDeviceControllerFactoryParamsClass
}

type _MTRDeviceControllerFactoryParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRDeviceControllerFactoryParams */
// An interface definition for the [MTRDeviceControllerFactoryParams] class.
type IMTRDeviceControllerFactoryParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRDeviceControllerFactoryParams */
	// properties:
	CertificationDeclarationCertificates() []foundation.Data
	SetCertificationDeclarationCertificates(value []foundation.Data)
	OtaProviderDelegate() unsafe.Pointer
	SetOtaProviderDelegate(value unsafe.Pointer)
	Port() objc.IObject /* cross-framework: NSNumber */
	SetPort(value objc.IObject /* cross-framework: NSNumber */)
	ProductAttestationAuthorityCertificates() []foundation.Data
	SetProductAttestationAuthorityCertificates(value []foundation.Data)
	ShouldStartServer() bool
	SetShouldStartServer(value bool)
	Storage() unsafe.Pointer
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRDeviceControllerFactoryParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRDeviceControllerFactoryParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRDeviceControllerFactoryParamsClass) Alloc() MTRDeviceControllerFactoryParams {
	rv := objc.Send[MTRDeviceControllerFactoryParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRDeviceControllerFactoryParamsClass) New() MTRDeviceControllerFactoryParams {
	rv := objc.Send[MTRDeviceControllerFactoryParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDeviceControllerFactoryParams) Init() MTRDeviceControllerFactoryParams {
	rv := objc.Send[MTRDeviceControllerFactoryParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDeviceControllerFactoryParams) Autorelease() MTRDeviceControllerFactoryParams {
	rv := objc.Send[MTRDeviceControllerFactoryParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDeviceControllerFactoryParams creates a new MTRDeviceControllerFactoryParams instance.
func NewMTRDeviceControllerFactoryParams() MTRDeviceControllerFactoryParams {
	return getMTRDeviceControllerFactoryParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRDeviceControllerFactoryParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceControllerFactoryParams
type MTRDeviceControllerFactoryParams struct {
	objectivec.Object
}

// MTRDeviceControllerFactoryParamsFrom constructs a [MTRDeviceControllerFactoryParams] from an unsafe.Pointer.
func MTRDeviceControllerFactoryParamsFrom(ptr unsafe.Pointer) MTRDeviceControllerFactoryParams {
	return MTRDeviceControllerFactoryParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRDeviceControllerFactoryParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceControllerFactoryParams/init(storage:)
func NewMTRDeviceControllerFactoryParamsWithStorage(storage unsafe.Pointer) MTRDeviceControllerFactoryParams {
	instance := getMTRDeviceControllerFactoryParamsClass().Alloc()
	rv := objc.Send[MTRDeviceControllerFactoryParams](instance.ID, objc.Sel("initWithStorage:"), storage)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRDeviceControllerFactoryParamsWithStorage */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRDeviceControllerFactoryParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRDeviceControllerFactoryParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRDeviceControllerFactoryParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRDeviceControllerFactoryParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceControllerFactoryParams/certificationDeclarationCertificates
func (m_ MTRDeviceControllerFactoryParams) CertificationDeclarationCertificates() []foundation.Data {
	rv := objc.Send[[]foundation.Data](m_.ID, objc.Sel("certificationDeclarationCertificates"))
	return rv
}/* debug [instance_properties/getter]: certificationDeclarationCertificates */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceControllerFactoryParams/certificationDeclarationCertificates
func (m_ MTRDeviceControllerFactoryParams) SetCertificationDeclarationCertificates(value []foundation.Data) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](m_.ID, objc.Sel("setCertificationDeclarationCertificates:"), nsArray)
}/* debug [instance_properties/setter]: certificationDeclarationCertificates */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceControllerFactoryParams/otaProviderDelegate
func (m_ MTRDeviceControllerFactoryParams) OtaProviderDelegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("otaProviderDelegate"))
	return rv
}/* debug [instance_properties/getter]: otaProviderDelegate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceControllerFactoryParams/otaProviderDelegate
func (m_ MTRDeviceControllerFactoryParams) SetOtaProviderDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOtaProviderDelegate:"), value)
}/* debug [instance_properties/setter]: otaProviderDelegate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceControllerFactoryParams/port
func (m_ MTRDeviceControllerFactoryParams) Port() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("port"))
	return rv
}/* debug [instance_properties/getter]: port */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceControllerFactoryParams/port
func (m_ MTRDeviceControllerFactoryParams) SetPort(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPort:"), value)
}/* debug [instance_properties/setter]: port */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceControllerFactoryParams/productAttestationAuthorityCertificates
func (m_ MTRDeviceControllerFactoryParams) ProductAttestationAuthorityCertificates() []foundation.Data {
	rv := objc.Send[[]foundation.Data](m_.ID, objc.Sel("productAttestationAuthorityCertificates"))
	return rv
}/* debug [instance_properties/getter]: productAttestationAuthorityCertificates */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceControllerFactoryParams/productAttestationAuthorityCertificates
func (m_ MTRDeviceControllerFactoryParams) SetProductAttestationAuthorityCertificates(value []foundation.Data) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](m_.ID, objc.Sel("setProductAttestationAuthorityCertificates:"), nsArray)
}/* debug [instance_properties/setter]: productAttestationAuthorityCertificates */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceControllerFactoryParams/shouldStartServer
func (m_ MTRDeviceControllerFactoryParams) ShouldStartServer() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("shouldStartServer"))
	return rv
}/* debug [instance_properties/getter]: shouldStartServer */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceControllerFactoryParams/shouldStartServer
func (m_ MTRDeviceControllerFactoryParams) SetShouldStartServer(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShouldStartServer:"), value)
}/* debug [instance_properties/setter]: shouldStartServer */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceControllerFactoryParams/storage
func (m_ MTRDeviceControllerFactoryParams) Storage() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("storage"))
	return rv
}/* debug [instance_properties/getter]: storage */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRDeviceControllerFactoryParams */


