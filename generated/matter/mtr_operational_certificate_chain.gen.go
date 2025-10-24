// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTROperationalCertificateChain */


/* debug [class_header]: Header for MTROperationalCertificateChain */
// The class instance for the [MTROperationalCertificateChain] class.
var (
	MTROperationalCertificateChainClass     _MTROperationalCertificateChainClass
	MTROperationalCertificateChainClassOnce sync.Once
)

func getMTROperationalCertificateChainClass() _MTROperationalCertificateChainClass {
	MTROperationalCertificateChainClassOnce.Do(func() {
		MTROperationalCertificateChainClass = _MTROperationalCertificateChainClass{objc.GetClass("MTROperationalCertificateChain")}
	})
	return MTROperationalCertificateChainClass
}

type _MTROperationalCertificateChainClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTROperationalCertificateChain */
// An interface definition for the [MTROperationalCertificateChain] class.
type IMTROperationalCertificateChain interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTROperationalCertificateChain */
	// properties:
	AdminSubject() objc.IObject /* cross-framework: NSNumber */
	SetAdminSubject(value objc.IObject /* cross-framework: NSNumber */)
	IntermediateCertificate() unsafe.Pointer
	SetIntermediateCertificate(value unsafe.Pointer)
	OperationalCertificate() unsafe.Pointer
	SetOperationalCertificate(value unsafe.Pointer)
	RootCertificate() unsafe.Pointer
	SetRootCertificate(value unsafe.Pointer)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTROperationalCertificateChain */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTROperationalCertificateChain */
// Alloc allocates a new instance without initialization.
func (mc _MTROperationalCertificateChainClass) Alloc() MTROperationalCertificateChain {
	rv := objc.Send[MTROperationalCertificateChain](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTROperationalCertificateChainClass) New() MTROperationalCertificateChain {
	rv := objc.Send[MTROperationalCertificateChain](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROperationalCertificateChain) Init() MTROperationalCertificateChain {
	rv := objc.Send[MTROperationalCertificateChain](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROperationalCertificateChain) Autorelease() MTROperationalCertificateChain {
	rv := objc.Send[MTROperationalCertificateChain](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROperationalCertificateChain creates a new MTROperationalCertificateChain instance.
func NewMTROperationalCertificateChain() MTROperationalCertificateChain {
	return getMTROperationalCertificateChainClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTROperationalCertificateChain */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCertificateChain
type MTROperationalCertificateChain struct {
	objectivec.Object
}

// MTROperationalCertificateChainFrom constructs a [MTROperationalCertificateChain] from an unsafe.Pointer.
func MTROperationalCertificateChainFrom(ptr unsafe.Pointer) MTROperationalCertificateChain {
	return MTROperationalCertificateChain{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTROperationalCertificateChain */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCertificateChain/init(operationalCertificate:intermediateCertificate:rootCertificate:adminSubject:)
func NewMTROperationalCertificateChainWithOperationalCertificateIntermediateCertificateRootCertificateAdminSubject(operationalCertificate unsafe.Pointer, intermediateCertificate unsafe.Pointer, rootCertificate unsafe.Pointer, adminSubject objc.IObject /* cross-framework: NSNumber */) MTROperationalCertificateChain {
	instance := getMTROperationalCertificateChainClass().Alloc()
	rv := objc.Send[MTROperationalCertificateChain](instance.ID, objc.Sel("initWithOperationalCertificate:intermediateCertificate:rootCertificate:adminSubject:"), operationalCertificate, intermediateCertificate, rootCertificate, adminSubject)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTROperationalCertificateChainWithOperationalCertificateIntermediateCertificateRootCertificateAdminSubject */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTROperationalCertificateChain */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTROperationalCertificateChain */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTROperationalCertificateChain */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTROperationalCertificateChain */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCertificateChain/adminSubject
func (m_ MTROperationalCertificateChain) AdminSubject() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("adminSubject"))
	return rv
}/* debug [instance_properties/getter]: adminSubject */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCertificateChain/adminSubject
func (m_ MTROperationalCertificateChain) SetAdminSubject(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAdminSubject:"), value)
}/* debug [instance_properties/setter]: adminSubject */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCertificateChain/intermediateCertificate
func (m_ MTROperationalCertificateChain) IntermediateCertificate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("intermediateCertificate"))
	return rv
}/* debug [instance_properties/getter]: intermediateCertificate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCertificateChain/intermediateCertificate
func (m_ MTROperationalCertificateChain) SetIntermediateCertificate(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIntermediateCertificate:"), value)
}/* debug [instance_properties/setter]: intermediateCertificate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCertificateChain/operationalCertificate
func (m_ MTROperationalCertificateChain) OperationalCertificate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("operationalCertificate"))
	return rv
}/* debug [instance_properties/getter]: operationalCertificate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCertificateChain/operationalCertificate
func (m_ MTROperationalCertificateChain) SetOperationalCertificate(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOperationalCertificate:"), value)
}/* debug [instance_properties/setter]: operationalCertificate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCertificateChain/rootCertificate
func (m_ MTROperationalCertificateChain) RootCertificate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("rootCertificate"))
	return rv
}/* debug [instance_properties/getter]: rootCertificate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCertificateChain/rootCertificate
func (m_ MTROperationalCertificateChain) SetRootCertificate(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRootCertificate:"), value)
}/* debug [instance_properties/setter]: rootCertificate */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTROperationalCertificateChain */


