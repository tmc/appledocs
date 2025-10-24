// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRProductIdentity */


/* debug [class_header]: Header for MTRProductIdentity */
// The class instance for the [MTRProductIdentity] class.
var (
	MTRProductIdentityClass     _MTRProductIdentityClass
	MTRProductIdentityClassOnce sync.Once
)

func getMTRProductIdentityClass() _MTRProductIdentityClass {
	MTRProductIdentityClassOnce.Do(func() {
		MTRProductIdentityClass = _MTRProductIdentityClass{objc.GetClass("MTRProductIdentity")}
	})
	return MTRProductIdentityClass
}

type _MTRProductIdentityClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRProductIdentity */
// An interface definition for the [MTRProductIdentity] class.
type IMTRProductIdentity interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRProductIdentity */
	// properties:
	ProductID() objc.IObject /* cross-framework: NSNumber */
	VendorID() objc.IObject /* cross-framework: NSNumber */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRProductIdentity */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRProductIdentity */
// Alloc allocates a new instance without initialization.
func (mc _MTRProductIdentityClass) Alloc() MTRProductIdentity {
	rv := objc.Send[MTRProductIdentity](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRProductIdentityClass) New() MTRProductIdentity {
	rv := objc.Send[MTRProductIdentity](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRProductIdentity) Init() MTRProductIdentity {
	rv := objc.Send[MTRProductIdentity](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRProductIdentity) Autorelease() MTRProductIdentity {
	rv := objc.Send[MTRProductIdentity](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRProductIdentity creates a new MTRProductIdentity instance.
func NewMTRProductIdentity() MTRProductIdentity {
	return getMTRProductIdentityClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRProductIdentity */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRProductIdentity
type MTRProductIdentity struct {
	objectivec.Object
}

// MTRProductIdentityFrom constructs a [MTRProductIdentity] from an unsafe.Pointer.
func MTRProductIdentityFrom(ptr unsafe.Pointer) MTRProductIdentity {
	return MTRProductIdentity{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRProductIdentity */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRProductIdentity/init(vendorID:productID:)
func NewMTRProductIdentityWithVendorIDProductID(vendorID objc.IObject /* cross-framework: NSNumber */, productID objc.IObject /* cross-framework: NSNumber */) MTRProductIdentity {
	instance := getMTRProductIdentityClass().Alloc()
	rv := objc.Send[MTRProductIdentity](instance.ID, objc.Sel("initWithVendorID:productID:"), vendorID, productID)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRProductIdentityWithVendorIDProductID */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRProductIdentity */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRProductIdentity */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRProductIdentity */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRProductIdentity */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRProductIdentity/productID
func (m_ MTRProductIdentity) ProductID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("productID"))
	return rv
}/* debug [instance_properties/getter]: productID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRProductIdentity/vendorID
func (m_ MTRProductIdentity) VendorID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("vendorID"))
	return rv
}/* debug [instance_properties/getter]: vendorID */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRProductIdentity */


