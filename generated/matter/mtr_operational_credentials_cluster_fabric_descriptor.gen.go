// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTROperationalCredentialsClusterFabricDescriptor */


/* debug [class_header]: Header for MTROperationalCredentialsClusterFabricDescriptor */
// The class instance for the [MTROperationalCredentialsClusterFabricDescriptor] class.
var (
	MTROperationalCredentialsClusterFabricDescriptorClass     _MTROperationalCredentialsClusterFabricDescriptorClass
	MTROperationalCredentialsClusterFabricDescriptorClassOnce sync.Once
)

func getMTROperationalCredentialsClusterFabricDescriptorClass() _MTROperationalCredentialsClusterFabricDescriptorClass {
	MTROperationalCredentialsClusterFabricDescriptorClassOnce.Do(func() {
		MTROperationalCredentialsClusterFabricDescriptorClass = _MTROperationalCredentialsClusterFabricDescriptorClass{objc.GetClass("MTROperationalCredentialsClusterFabricDescriptor")}
	})
	return MTROperationalCredentialsClusterFabricDescriptorClass
}

type _MTROperationalCredentialsClusterFabricDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTROperationalCredentialsClusterFabricDescriptor */
// An interface definition for the [MTROperationalCredentialsClusterFabricDescriptor] class.
type IMTROperationalCredentialsClusterFabricDescriptor interface {
	IMTROperationalCredentialsClusterFabricDescriptorStruct
	
/* debug [class_interface_properties]: Properties for MTROperationalCredentialsClusterFabricDescriptor */
	// properties:
	FabricIndex() objc.IObject /* cross-framework: NSNumber */
	SetFabricIndex(value objc.IObject /* cross-framework: NSNumber */)
	Label() objc.IObject /* cross-framework: NSString */
	SetLabel(value objc.IObject /* cross-framework: NSString */)
	RootPublicKey() objc.IObject /* cross-framework: NSData */
	SetRootPublicKey(value objc.IObject /* cross-framework: NSData */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTROperationalCredentialsClusterFabricDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTROperationalCredentialsClusterFabricDescriptor */
// Alloc allocates a new instance without initialization.
func (mc _MTROperationalCredentialsClusterFabricDescriptorClass) Alloc() MTROperationalCredentialsClusterFabricDescriptor {
	rv := objc.Send[MTROperationalCredentialsClusterFabricDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTROperationalCredentialsClusterFabricDescriptorClass) New() MTROperationalCredentialsClusterFabricDescriptor {
	rv := objc.Send[MTROperationalCredentialsClusterFabricDescriptor](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROperationalCredentialsClusterFabricDescriptor) Init() MTROperationalCredentialsClusterFabricDescriptor {
	rv := objc.Send[MTROperationalCredentialsClusterFabricDescriptor](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROperationalCredentialsClusterFabricDescriptor) Autorelease() MTROperationalCredentialsClusterFabricDescriptor {
	rv := objc.Send[MTROperationalCredentialsClusterFabricDescriptor](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROperationalCredentialsClusterFabricDescriptor creates a new MTROperationalCredentialsClusterFabricDescriptor instance.
func NewMTROperationalCredentialsClusterFabricDescriptor() MTROperationalCredentialsClusterFabricDescriptor {
	return getMTROperationalCredentialsClusterFabricDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTROperationalCredentialsClusterFabricDescriptor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterFabricDescriptor
type MTROperationalCredentialsClusterFabricDescriptor struct {
	MTROperationalCredentialsClusterFabricDescriptorStruct
}

// MTROperationalCredentialsClusterFabricDescriptorFrom constructs a [MTROperationalCredentialsClusterFabricDescriptor] from an unsafe.Pointer.
func MTROperationalCredentialsClusterFabricDescriptorFrom(ptr unsafe.Pointer) MTROperationalCredentialsClusterFabricDescriptor {
	return MTROperationalCredentialsClusterFabricDescriptor{
		MTROperationalCredentialsClusterFabricDescriptorStruct: MTROperationalCredentialsClusterFabricDescriptorStructFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTROperationalCredentialsClusterFabricDescriptor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTROperationalCredentialsClusterFabricDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTROperationalCredentialsClusterFabricDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTROperationalCredentialsClusterFabricDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTROperationalCredentialsClusterFabricDescriptor */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterFabricDescriptor/fabricIndex
func (m_ MTROperationalCredentialsClusterFabricDescriptor) FabricIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("fabricIndex"))
	return rv
}/* debug [instance_properties/getter]: fabricIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterFabricDescriptor/fabricIndex
func (m_ MTROperationalCredentialsClusterFabricDescriptor) SetFabricIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricIndex:"), value)
}/* debug [instance_properties/setter]: fabricIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterFabricDescriptor/label
func (m_ MTROperationalCredentialsClusterFabricDescriptor) Label() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("label"))
	return rv
}/* debug [instance_properties/getter]: label */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterFabricDescriptor/label
func (m_ MTROperationalCredentialsClusterFabricDescriptor) SetLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLabel:"), value)
}/* debug [instance_properties/setter]: label */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterFabricDescriptor/rootPublicKey
func (m_ MTROperationalCredentialsClusterFabricDescriptor) RootPublicKey() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("rootPublicKey"))
	return rv
}/* debug [instance_properties/getter]: rootPublicKey */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterFabricDescriptor/rootPublicKey
func (m_ MTROperationalCredentialsClusterFabricDescriptor) SetRootPublicKey(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRootPublicKey:"), value)
}/* debug [instance_properties/setter]: rootPublicKey */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTROperationalCredentialsClusterFabricDescriptor */



