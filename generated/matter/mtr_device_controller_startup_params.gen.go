// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRDeviceControllerStartupParams */


/* debug [class_header]: Header for MTRDeviceControllerStartupParams */
// The class instance for the [MTRDeviceControllerStartupParams] class.
var (
	MTRDeviceControllerStartupParamsClass     _MTRDeviceControllerStartupParamsClass
	MTRDeviceControllerStartupParamsClassOnce sync.Once
)

func getMTRDeviceControllerStartupParamsClass() _MTRDeviceControllerStartupParamsClass {
	MTRDeviceControllerStartupParamsClassOnce.Do(func() {
		MTRDeviceControllerStartupParamsClass = _MTRDeviceControllerStartupParamsClass{objc.GetClass("MTRDeviceControllerStartupParams")}
	})
	return MTRDeviceControllerStartupParamsClass
}

type _MTRDeviceControllerStartupParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRDeviceControllerStartupParams */
// An interface definition for the [MTRDeviceControllerStartupParams] class.
type IMTRDeviceControllerStartupParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRDeviceControllerStartupParams */
	// properties:
	CaseAuthenticatedTags() unsafe.Pointer
	SetCaseAuthenticatedTags(value unsafe.Pointer)
	FabricID() objc.IObject /* cross-framework: NSNumber */
	FabricId() uint64
	IntermediateCertificate() unsafe.Pointer
	SetIntermediateCertificate(value unsafe.Pointer)
	Ipk() objc.IObject /* cross-framework: NSData */
	NocSigner() unsafe.Pointer
	NodeID() objc.IObject /* cross-framework: NSNumber */
	SetNodeID(value objc.IObject /* cross-framework: NSNumber */)
	NodeId() objc.IObject /* cross-framework: NSNumber */
	SetNodeId(value objc.IObject /* cross-framework: NSNumber */)
	OperationalCertificate() unsafe.Pointer
	OperationalCertificateIssuer() unsafe.Pointer
	SetOperationalCertificateIssuer(value unsafe.Pointer)
	OperationalCertificateIssuerQueue() unsafe.Pointer
	SetOperationalCertificateIssuerQueue(value unsafe.Pointer)
	OperationalKeypair() unsafe.Pointer
	SetOperationalKeypair(value unsafe.Pointer)
	RootCertificate() unsafe.Pointer
	SetRootCertificate(value unsafe.Pointer)
	VendorId() objc.IObject /* cross-framework: NSNumber */
	SetVendorId(value objc.IObject /* cross-framework: NSNumber */)
	VendorID() objc.IObject /* cross-framework: NSNumber */
	SetVendorID(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRDeviceControllerStartupParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRDeviceControllerStartupParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRDeviceControllerStartupParamsClass) Alloc() MTRDeviceControllerStartupParams {
	rv := objc.Send[MTRDeviceControllerStartupParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRDeviceControllerStartupParamsClass) New() MTRDeviceControllerStartupParams {
	rv := objc.Send[MTRDeviceControllerStartupParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDeviceControllerStartupParams) Init() MTRDeviceControllerStartupParams {
	rv := objc.Send[MTRDeviceControllerStartupParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDeviceControllerStartupParams) Autorelease() MTRDeviceControllerStartupParams {
	rv := objc.Send[MTRDeviceControllerStartupParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDeviceControllerStartupParams creates a new MTRDeviceControllerStartupParams instance.
func NewMTRDeviceControllerStartupParams() MTRDeviceControllerStartupParams {
	return getMTRDeviceControllerStartupParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRDeviceControllerStartupParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceControllerStartupParams
type MTRDeviceControllerStartupParams struct {
	objectivec.Object
}

// MTRDeviceControllerStartupParamsFrom constructs a [MTRDeviceControllerStartupParams] from an unsafe.Pointer.
func MTRDeviceControllerStartupParamsFrom(ptr unsafe.Pointer) MTRDeviceControllerStartupParams {
	return MTRDeviceControllerStartupParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRDeviceControllerStartupParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceControllerStartupParams/init(ipk:fabricID:nocSigner:)
func NewMTRDeviceControllerStartupParamsWithIPKFabricIDNocSigner(ipk objc.IObject /* cross-framework: NSData */, fabricID objc.IObject /* cross-framework: NSNumber */, nocSigner unsafe.Pointer) MTRDeviceControllerStartupParams {
	instance := getMTRDeviceControllerStartupParamsClass().Alloc()
	rv := objc.Send[MTRDeviceControllerStartupParams](instance.ID, objc.Sel("initWithIPK:fabricID:nocSigner:"), ipk, fabricID, nocSigner)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRDeviceControllerStartupParamsWithIPKFabricIDNocSigner */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceControllerStartupParams/init(ipk:operationalKeypair:operationalCertificate:intermediateCertificate:rootCertificate:)
func NewMTRDeviceControllerStartupParamsWithIPKOperationalKeypairOperationalCertificateIntermediateCertificateRootCertificate(ipk objc.IObject /* cross-framework: NSData */, operationalKeypair unsafe.Pointer, operationalCertificate unsafe.Pointer, intermediateCertificate unsafe.Pointer, rootCertificate unsafe.Pointer) MTRDeviceControllerStartupParams {
	instance := getMTRDeviceControllerStartupParamsClass().Alloc()
	rv := objc.Send[MTRDeviceControllerStartupParams](instance.ID, objc.Sel("initWithIPK:operationalKeypair:operationalCertificate:intermediateCertificate:rootCertificate:"), ipk, operationalKeypair, operationalCertificate, intermediateCertificate, rootCertificate)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRDeviceControllerStartupParamsWithIPKOperationalKeypairOperationalCertificateIntermediateCertificateRootCertificate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceControllerStartupParams/init(operationalKeypair:operationalCertificate:intermediateCertificate:rootCertificate:ipk:)
func NewMTRDeviceControllerStartupParamsWithOperationalKeypairOperationalCertificateIntermediateCertificateRootCertificateIpk(operationalKeypair unsafe.Pointer, operationalCertificate unsafe.Pointer, intermediateCertificate unsafe.Pointer, rootCertificate unsafe.Pointer, ipk objc.IObject /* cross-framework: NSData */) MTRDeviceControllerStartupParams {
	instance := getMTRDeviceControllerStartupParamsClass().Alloc()
	rv := objc.Send[MTRDeviceControllerStartupParams](instance.ID, objc.Sel("initWithOperationalKeypair:operationalCertificate:intermediateCertificate:rootCertificate:ipk:"), operationalKeypair, operationalCertificate, intermediateCertificate, rootCertificate, ipk)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRDeviceControllerStartupParamsWithOperationalKeypairOperationalCertificateIntermediateCertificateRootCertificateIpk */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceControllerStartupParams/init(signingKeypair:fabricId:ipk:)
func NewMTRDeviceControllerStartupParamsWithSigningKeypairFabricIdIpk(nocSigner unsafe.Pointer, fabricId uint64, ipk objc.IObject /* cross-framework: NSData */) MTRDeviceControllerStartupParams {
	instance := getMTRDeviceControllerStartupParamsClass().Alloc()
	rv := objc.Send[MTRDeviceControllerStartupParams](instance.ID, objc.Sel("initWithSigningKeypair:fabricId:ipk:"), nocSigner, fabricId, ipk)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRDeviceControllerStartupParamsWithSigningKeypairFabricIdIpk */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRDeviceControllerStartupParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRDeviceControllerStartupParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRDeviceControllerStartupParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRDeviceControllerStartupParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceControllerStartupParams/caseAuthenticatedTags
func (m_ MTRDeviceControllerStartupParams) CaseAuthenticatedTags() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("caseAuthenticatedTags"))
	return rv
}/* debug [instance_properties/getter]: caseAuthenticatedTags */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceControllerStartupParams/caseAuthenticatedTags
func (m_ MTRDeviceControllerStartupParams) SetCaseAuthenticatedTags(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCaseAuthenticatedTags:"), value)
}/* debug [instance_properties/setter]: caseAuthenticatedTags */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceControllerStartupParams/fabricID-1cm6z
func (m_ MTRDeviceControllerStartupParams) FabricID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("fabricID"))
	return rv
}/* debug [instance_properties/getter]: fabricID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceControllerStartupParams/fabricId-1cm7v
func (m_ MTRDeviceControllerStartupParams) FabricId() uint64 {
	rv := objc.Send[uint64](m_.ID, objc.Sel("fabricId"))
	return rv
}/* debug [instance_properties/getter]: fabricId */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceControllerStartupParams/intermediateCertificate
func (m_ MTRDeviceControllerStartupParams) IntermediateCertificate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("intermediateCertificate"))
	return rv
}/* debug [instance_properties/getter]: intermediateCertificate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceControllerStartupParams/intermediateCertificate
func (m_ MTRDeviceControllerStartupParams) SetIntermediateCertificate(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIntermediateCertificate:"), value)
}/* debug [instance_properties/setter]: intermediateCertificate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceControllerStartupParams/ipk
func (m_ MTRDeviceControllerStartupParams) Ipk() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("ipk"))
	return rv
}/* debug [instance_properties/getter]: ipk */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceControllerStartupParams/nocSigner
func (m_ MTRDeviceControllerStartupParams) NocSigner() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("nocSigner"))
	return rv
}/* debug [instance_properties/getter]: nocSigner */


// Node id for this controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceControllerStartupParams/nodeID-9iwwv
func (m_ MTRDeviceControllerStartupParams) NodeID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nodeID"))
	return rv
}/* debug [instance_properties/getter]: nodeID */


// Node id for this controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceControllerStartupParams/nodeID-9iwwv
func (m_ MTRDeviceControllerStartupParams) SetNodeID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNodeID:"), value)
}/* debug [instance_properties/setter]: nodeID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceControllerStartupParams/nodeId-9iwxr
func (m_ MTRDeviceControllerStartupParams) NodeId() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nodeId"))
	return rv
}/* debug [instance_properties/getter]: nodeId */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceControllerStartupParams/nodeId-9iwxr
func (m_ MTRDeviceControllerStartupParams) SetNodeId(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNodeId:"), value)
}/* debug [instance_properties/setter]: nodeId */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceControllerStartupParams/operationalCertificate
func (m_ MTRDeviceControllerStartupParams) OperationalCertificate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("operationalCertificate"))
	return rv
}/* debug [instance_properties/getter]: operationalCertificate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceControllerStartupParams/operationalCertificateIssuer
func (m_ MTRDeviceControllerStartupParams) OperationalCertificateIssuer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("operationalCertificateIssuer"))
	return rv
}/* debug [instance_properties/getter]: operationalCertificateIssuer */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceControllerStartupParams/operationalCertificateIssuer
func (m_ MTRDeviceControllerStartupParams) SetOperationalCertificateIssuer(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOperationalCertificateIssuer:"), value)
}/* debug [instance_properties/setter]: operationalCertificateIssuer */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceControllerStartupParams/operationalCertificateIssuerQueue
func (m_ MTRDeviceControllerStartupParams) OperationalCertificateIssuerQueue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("operationalCertificateIssuerQueue"))
	return rv
}/* debug [instance_properties/getter]: operationalCertificateIssuerQueue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceControllerStartupParams/operationalCertificateIssuerQueue
func (m_ MTRDeviceControllerStartupParams) SetOperationalCertificateIssuerQueue(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOperationalCertificateIssuerQueue:"), value)
}/* debug [instance_properties/setter]: operationalCertificateIssuerQueue */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceControllerStartupParams/operationalKeypair
func (m_ MTRDeviceControllerStartupParams) OperationalKeypair() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("operationalKeypair"))
	return rv
}/* debug [instance_properties/getter]: operationalKeypair */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceControllerStartupParams/operationalKeypair
func (m_ MTRDeviceControllerStartupParams) SetOperationalKeypair(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOperationalKeypair:"), value)
}/* debug [instance_properties/setter]: operationalKeypair */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceControllerStartupParams/rootCertificate
func (m_ MTRDeviceControllerStartupParams) RootCertificate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("rootCertificate"))
	return rv
}/* debug [instance_properties/getter]: rootCertificate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceControllerStartupParams/rootCertificate
func (m_ MTRDeviceControllerStartupParams) SetRootCertificate(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRootCertificate:"), value)
}/* debug [instance_properties/setter]: rootCertificate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceControllerStartupParams/vendorId-8ru0w
func (m_ MTRDeviceControllerStartupParams) VendorId() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("vendorId"))
	return rv
}/* debug [instance_properties/getter]: vendorId */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceControllerStartupParams/vendorId-8ru0w
func (m_ MTRDeviceControllerStartupParams) SetVendorId(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVendorId:"), value)
}/* debug [instance_properties/setter]: vendorId */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceControllerStartupParams/vendorID-8ru1s
func (m_ MTRDeviceControllerStartupParams) VendorID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("vendorID"))
	return rv
}/* debug [instance_properties/getter]: vendorID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceControllerStartupParams/vendorID-8ru1s
func (m_ MTRDeviceControllerStartupParams) SetVendorID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVendorID:"), value)
}/* debug [instance_properties/setter]: vendorID */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRDeviceControllerStartupParams */


