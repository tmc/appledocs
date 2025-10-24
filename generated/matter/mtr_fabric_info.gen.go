// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRFabricInfo */


/* debug [class_header]: Header for MTRFabricInfo */
// The class instance for the [MTRFabricInfo] class.
var (
	MTRFabricInfoClass     _MTRFabricInfoClass
	MTRFabricInfoClassOnce sync.Once
)

func getMTRFabricInfoClass() _MTRFabricInfoClass {
	MTRFabricInfoClassOnce.Do(func() {
		MTRFabricInfoClass = _MTRFabricInfoClass{objc.GetClass("MTRFabricInfo")}
	})
	return MTRFabricInfoClass
}

type _MTRFabricInfoClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRFabricInfo */
// An interface definition for the [MTRFabricInfo] class.
type IMTRFabricInfo interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRFabricInfo */
	// properties:
	FabricID() objc.IObject /* cross-framework: NSNumber */
	FabricIndex() objc.IObject /* cross-framework: NSNumber */
	IntermediateCertificate() unsafe.Pointer
	IntermediateCertificateTLV() unsafe.Pointer
	Label() objc.IObject /* cross-framework: NSString */
	NodeID() objc.IObject /* cross-framework: NSNumber */
	OperationalCertificate() unsafe.Pointer
	OperationalCertificateTLV() unsafe.Pointer
	RootCertificate() unsafe.Pointer
	RootCertificateTLV() unsafe.Pointer
	RootPublicKey() objc.IObject /* cross-framework: NSData */
	VendorID() objc.IObject /* cross-framework: NSNumber */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRFabricInfo */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRFabricInfo */
// Alloc allocates a new instance without initialization.
func (mc _MTRFabricInfoClass) Alloc() MTRFabricInfo {
	rv := objc.Send[MTRFabricInfo](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRFabricInfoClass) New() MTRFabricInfo {
	rv := objc.Send[MTRFabricInfo](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRFabricInfo) Init() MTRFabricInfo {
	rv := objc.Send[MTRFabricInfo](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRFabricInfo) Autorelease() MTRFabricInfo {
	rv := objc.Send[MTRFabricInfo](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRFabricInfo creates a new MTRFabricInfo instance.
func NewMTRFabricInfo() MTRFabricInfo {
	return getMTRFabricInfoClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRFabricInfo */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRFabricInfo
type MTRFabricInfo struct {
	objectivec.Object
}

// MTRFabricInfoFrom constructs a [MTRFabricInfo] from an unsafe.Pointer.
func MTRFabricInfoFrom(ptr unsafe.Pointer) MTRFabricInfo {
	return MTRFabricInfo{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRFabricInfo *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRFabricInfo */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRFabricInfo */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRFabricInfo */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRFabricInfo */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRFabricInfo/fabricID
func (m_ MTRFabricInfo) FabricID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("fabricID"))
	return rv
}/* debug [instance_properties/getter]: fabricID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRFabricInfo/fabricIndex
func (m_ MTRFabricInfo) FabricIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("fabricIndex"))
	return rv
}/* debug [instance_properties/getter]: fabricIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRFabricInfo/intermediateCertificate
func (m_ MTRFabricInfo) IntermediateCertificate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("intermediateCertificate"))
	return rv
}/* debug [instance_properties/getter]: intermediateCertificate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRFabricInfo/intermediateCertificateTLV
func (m_ MTRFabricInfo) IntermediateCertificateTLV() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("intermediateCertificateTLV"))
	return rv
}/* debug [instance_properties/getter]: intermediateCertificateTLV */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRFabricInfo/label
func (m_ MTRFabricInfo) Label() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("label"))
	return rv
}/* debug [instance_properties/getter]: label */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRFabricInfo/nodeID
func (m_ MTRFabricInfo) NodeID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nodeID"))
	return rv
}/* debug [instance_properties/getter]: nodeID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRFabricInfo/operationalCertificate
func (m_ MTRFabricInfo) OperationalCertificate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("operationalCertificate"))
	return rv
}/* debug [instance_properties/getter]: operationalCertificate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRFabricInfo/operationalCertificateTLV
func (m_ MTRFabricInfo) OperationalCertificateTLV() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("operationalCertificateTLV"))
	return rv
}/* debug [instance_properties/getter]: operationalCertificateTLV */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRFabricInfo/rootCertificate
func (m_ MTRFabricInfo) RootCertificate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("rootCertificate"))
	return rv
}/* debug [instance_properties/getter]: rootCertificate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRFabricInfo/rootCertificateTLV
func (m_ MTRFabricInfo) RootCertificateTLV() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("rootCertificateTLV"))
	return rv
}/* debug [instance_properties/getter]: rootCertificateTLV */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRFabricInfo/rootPublicKey
func (m_ MTRFabricInfo) RootPublicKey() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("rootPublicKey"))
	return rv
}/* debug [instance_properties/getter]: rootPublicKey */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRFabricInfo/vendorID
func (m_ MTRFabricInfo) VendorID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("vendorID"))
	return rv
}/* debug [instance_properties/getter]: vendorID */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRFabricInfo */



