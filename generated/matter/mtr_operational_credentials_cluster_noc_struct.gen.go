// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTROperationalCredentialsClusterNOCStruct */


/* debug [class_header]: Header for MTROperationalCredentialsClusterNOCStruct */
// The class instance for the [MTROperationalCredentialsClusterNOCStruct] class.
var (
	MTROperationalCredentialsClusterNOCStructClass     _MTROperationalCredentialsClusterNOCStructClass
	MTROperationalCredentialsClusterNOCStructClassOnce sync.Once
)

func getMTROperationalCredentialsClusterNOCStructClass() _MTROperationalCredentialsClusterNOCStructClass {
	MTROperationalCredentialsClusterNOCStructClassOnce.Do(func() {
		MTROperationalCredentialsClusterNOCStructClass = _MTROperationalCredentialsClusterNOCStructClass{objc.GetClass("MTROperationalCredentialsClusterNOCStruct")}
	})
	return MTROperationalCredentialsClusterNOCStructClass
}

type _MTROperationalCredentialsClusterNOCStructClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTROperationalCredentialsClusterNOCStruct */
// An interface definition for the [MTROperationalCredentialsClusterNOCStruct] class.
type IMTROperationalCredentialsClusterNOCStruct interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTROperationalCredentialsClusterNOCStruct */
	// properties:
	FabricIndex() objc.IObject /* cross-framework: NSNumber */
	SetFabricIndex(value objc.IObject /* cross-framework: NSNumber */)
	Icac() objc.IObject /* cross-framework: NSData */
	SetIcac(value objc.IObject /* cross-framework: NSData */)
	Noc() objc.IObject /* cross-framework: NSData */
	SetNoc(value objc.IObject /* cross-framework: NSData */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTROperationalCredentialsClusterNOCStruct */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTROperationalCredentialsClusterNOCStruct */
// Alloc allocates a new instance without initialization.
func (mc _MTROperationalCredentialsClusterNOCStructClass) Alloc() MTROperationalCredentialsClusterNOCStruct {
	rv := objc.Send[MTROperationalCredentialsClusterNOCStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTROperationalCredentialsClusterNOCStructClass) New() MTROperationalCredentialsClusterNOCStruct {
	rv := objc.Send[MTROperationalCredentialsClusterNOCStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROperationalCredentialsClusterNOCStruct) Init() MTROperationalCredentialsClusterNOCStruct {
	rv := objc.Send[MTROperationalCredentialsClusterNOCStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROperationalCredentialsClusterNOCStruct) Autorelease() MTROperationalCredentialsClusterNOCStruct {
	rv := objc.Send[MTROperationalCredentialsClusterNOCStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROperationalCredentialsClusterNOCStruct creates a new MTROperationalCredentialsClusterNOCStruct instance.
func NewMTROperationalCredentialsClusterNOCStruct() MTROperationalCredentialsClusterNOCStruct {
	return getMTROperationalCredentialsClusterNOCStructClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTROperationalCredentialsClusterNOCStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterNOCStruct
type MTROperationalCredentialsClusterNOCStruct struct {
	objectivec.Object
}

// MTROperationalCredentialsClusterNOCStructFrom constructs a [MTROperationalCredentialsClusterNOCStruct] from an unsafe.Pointer.
func MTROperationalCredentialsClusterNOCStructFrom(ptr unsafe.Pointer) MTROperationalCredentialsClusterNOCStruct {
	return MTROperationalCredentialsClusterNOCStruct{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTROperationalCredentialsClusterNOCStruct *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTROperationalCredentialsClusterNOCStruct */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTROperationalCredentialsClusterNOCStruct */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTROperationalCredentialsClusterNOCStruct */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTROperationalCredentialsClusterNOCStruct */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterNOCStruct/fabricIndex
func (m_ MTROperationalCredentialsClusterNOCStruct) FabricIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("fabricIndex"))
	return rv
}/* debug [instance_properties/getter]: fabricIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterNOCStruct/fabricIndex
func (m_ MTROperationalCredentialsClusterNOCStruct) SetFabricIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricIndex:"), value)
}/* debug [instance_properties/setter]: fabricIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterNOCStruct/icac
func (m_ MTROperationalCredentialsClusterNOCStruct) Icac() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("icac"))
	return rv
}/* debug [instance_properties/getter]: icac */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterNOCStruct/icac
func (m_ MTROperationalCredentialsClusterNOCStruct) SetIcac(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIcac:"), value)
}/* debug [instance_properties/setter]: icac */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterNOCStruct/noc
func (m_ MTROperationalCredentialsClusterNOCStruct) Noc() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("noc"))
	return rv
}/* debug [instance_properties/getter]: noc */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterNOCStruct/noc
func (m_ MTROperationalCredentialsClusterNOCStruct) SetNoc(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNoc:"), value)
}/* debug [instance_properties/setter]: noc */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTROperationalCredentialsClusterNOCStruct */



