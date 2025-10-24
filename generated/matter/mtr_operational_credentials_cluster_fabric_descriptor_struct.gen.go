// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTROperationalCredentialsClusterFabricDescriptorStruct */


/* debug [class_header]: Header for MTROperationalCredentialsClusterFabricDescriptorStruct */
// The class instance for the [MTROperationalCredentialsClusterFabricDescriptorStruct] class.
var (
	MTROperationalCredentialsClusterFabricDescriptorStructClass     _MTROperationalCredentialsClusterFabricDescriptorStructClass
	MTROperationalCredentialsClusterFabricDescriptorStructClassOnce sync.Once
)

func getMTROperationalCredentialsClusterFabricDescriptorStructClass() _MTROperationalCredentialsClusterFabricDescriptorStructClass {
	MTROperationalCredentialsClusterFabricDescriptorStructClassOnce.Do(func() {
		MTROperationalCredentialsClusterFabricDescriptorStructClass = _MTROperationalCredentialsClusterFabricDescriptorStructClass{objc.GetClass("MTROperationalCredentialsClusterFabricDescriptorStruct")}
	})
	return MTROperationalCredentialsClusterFabricDescriptorStructClass
}

type _MTROperationalCredentialsClusterFabricDescriptorStructClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTROperationalCredentialsClusterFabricDescriptorStruct */
// An interface definition for the [MTROperationalCredentialsClusterFabricDescriptorStruct] class.
type IMTROperationalCredentialsClusterFabricDescriptorStruct interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTROperationalCredentialsClusterFabricDescriptorStruct */
	// properties:
	FabricId() objc.IObject /* cross-framework: NSNumber */
	SetFabricId(value objc.IObject /* cross-framework: NSNumber */)
	FabricID() objc.IObject /* cross-framework: NSNumber */
	SetFabricID(value objc.IObject /* cross-framework: NSNumber */)
	FabricIndex() objc.IObject /* cross-framework: NSNumber */
	SetFabricIndex(value objc.IObject /* cross-framework: NSNumber */)
	Label() objc.IObject /* cross-framework: NSString */
	SetLabel(value objc.IObject /* cross-framework: NSString */)
	NodeID() objc.IObject /* cross-framework: NSNumber */
	SetNodeID(value objc.IObject /* cross-framework: NSNumber */)
	NodeId() objc.IObject /* cross-framework: NSNumber */
	SetNodeId(value objc.IObject /* cross-framework: NSNumber */)
	RootPublicKey() objc.IObject /* cross-framework: NSData */
	SetRootPublicKey(value objc.IObject /* cross-framework: NSData */)
	VendorId() objc.IObject /* cross-framework: NSNumber */
	SetVendorId(value objc.IObject /* cross-framework: NSNumber */)
	VendorID() objc.IObject /* cross-framework: NSNumber */
	SetVendorID(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTROperationalCredentialsClusterFabricDescriptorStruct */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTROperationalCredentialsClusterFabricDescriptorStruct */
// Alloc allocates a new instance without initialization.
func (mc _MTROperationalCredentialsClusterFabricDescriptorStructClass) Alloc() MTROperationalCredentialsClusterFabricDescriptorStruct {
	rv := objc.Send[MTROperationalCredentialsClusterFabricDescriptorStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTROperationalCredentialsClusterFabricDescriptorStructClass) New() MTROperationalCredentialsClusterFabricDescriptorStruct {
	rv := objc.Send[MTROperationalCredentialsClusterFabricDescriptorStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROperationalCredentialsClusterFabricDescriptorStruct) Init() MTROperationalCredentialsClusterFabricDescriptorStruct {
	rv := objc.Send[MTROperationalCredentialsClusterFabricDescriptorStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROperationalCredentialsClusterFabricDescriptorStruct) Autorelease() MTROperationalCredentialsClusterFabricDescriptorStruct {
	rv := objc.Send[MTROperationalCredentialsClusterFabricDescriptorStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROperationalCredentialsClusterFabricDescriptorStruct creates a new MTROperationalCredentialsClusterFabricDescriptorStruct instance.
func NewMTROperationalCredentialsClusterFabricDescriptorStruct() MTROperationalCredentialsClusterFabricDescriptorStruct {
	return getMTROperationalCredentialsClusterFabricDescriptorStructClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTROperationalCredentialsClusterFabricDescriptorStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterFabricDescriptorStruct
type MTROperationalCredentialsClusterFabricDescriptorStruct struct {
	objectivec.Object
}

// MTROperationalCredentialsClusterFabricDescriptorStructFrom constructs a [MTROperationalCredentialsClusterFabricDescriptorStruct] from an unsafe.Pointer.
func MTROperationalCredentialsClusterFabricDescriptorStructFrom(ptr unsafe.Pointer) MTROperationalCredentialsClusterFabricDescriptorStruct {
	return MTROperationalCredentialsClusterFabricDescriptorStruct{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTROperationalCredentialsClusterFabricDescriptorStruct *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTROperationalCredentialsClusterFabricDescriptorStruct */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTROperationalCredentialsClusterFabricDescriptorStruct */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTROperationalCredentialsClusterFabricDescriptorStruct */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTROperationalCredentialsClusterFabricDescriptorStruct */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterFabricDescriptorStruct/fabricId-5tetp
func (m_ MTROperationalCredentialsClusterFabricDescriptorStruct) FabricId() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("fabricId"))
	return rv
}/* debug [instance_properties/getter]: fabricId */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterFabricDescriptorStruct/fabricId-5tetp
func (m_ MTROperationalCredentialsClusterFabricDescriptorStruct) SetFabricId(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricId:"), value)
}/* debug [instance_properties/setter]: fabricId */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterFabricDescriptorStruct/fabricID-5teul
func (m_ MTROperationalCredentialsClusterFabricDescriptorStruct) FabricID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("fabricID"))
	return rv
}/* debug [instance_properties/getter]: fabricID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterFabricDescriptorStruct/fabricID-5teul
func (m_ MTROperationalCredentialsClusterFabricDescriptorStruct) SetFabricID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricID:"), value)
}/* debug [instance_properties/setter]: fabricID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterFabricDescriptorStruct/fabricIndex
func (m_ MTROperationalCredentialsClusterFabricDescriptorStruct) FabricIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("fabricIndex"))
	return rv
}/* debug [instance_properties/getter]: fabricIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterFabricDescriptorStruct/fabricIndex
func (m_ MTROperationalCredentialsClusterFabricDescriptorStruct) SetFabricIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricIndex:"), value)
}/* debug [instance_properties/setter]: fabricIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterFabricDescriptorStruct/label
func (m_ MTROperationalCredentialsClusterFabricDescriptorStruct) Label() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("label"))
	return rv
}/* debug [instance_properties/getter]: label */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterFabricDescriptorStruct/label
func (m_ MTROperationalCredentialsClusterFabricDescriptorStruct) SetLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLabel:"), value)
}/* debug [instance_properties/setter]: label */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterFabricDescriptorStruct/nodeID-3hsb0
func (m_ MTROperationalCredentialsClusterFabricDescriptorStruct) NodeID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nodeID"))
	return rv
}/* debug [instance_properties/getter]: nodeID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterFabricDescriptorStruct/nodeID-3hsb0
func (m_ MTROperationalCredentialsClusterFabricDescriptorStruct) SetNodeID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNodeID:"), value)
}/* debug [instance_properties/setter]: nodeID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterFabricDescriptorStruct/nodeId-3hsbw
func (m_ MTROperationalCredentialsClusterFabricDescriptorStruct) NodeId() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nodeId"))
	return rv
}/* debug [instance_properties/getter]: nodeId */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterFabricDescriptorStruct/nodeId-3hsbw
func (m_ MTROperationalCredentialsClusterFabricDescriptorStruct) SetNodeId(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNodeId:"), value)
}/* debug [instance_properties/setter]: nodeId */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterFabricDescriptorStruct/rootPublicKey
func (m_ MTROperationalCredentialsClusterFabricDescriptorStruct) RootPublicKey() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("rootPublicKey"))
	return rv
}/* debug [instance_properties/getter]: rootPublicKey */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterFabricDescriptorStruct/rootPublicKey
func (m_ MTROperationalCredentialsClusterFabricDescriptorStruct) SetRootPublicKey(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRootPublicKey:"), value)
}/* debug [instance_properties/setter]: rootPublicKey */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterFabricDescriptorStruct/vendorId-3iaxd
func (m_ MTROperationalCredentialsClusterFabricDescriptorStruct) VendorId() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("vendorId"))
	return rv
}/* debug [instance_properties/getter]: vendorId */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterFabricDescriptorStruct/vendorId-3iaxd
func (m_ MTROperationalCredentialsClusterFabricDescriptorStruct) SetVendorId(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVendorId:"), value)
}/* debug [instance_properties/setter]: vendorId */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterFabricDescriptorStruct/vendorID-3iay9
func (m_ MTROperationalCredentialsClusterFabricDescriptorStruct) VendorID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("vendorID"))
	return rv
}/* debug [instance_properties/getter]: vendorID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalCredentialsClusterFabricDescriptorStruct/vendorID-3iay9
func (m_ MTROperationalCredentialsClusterFabricDescriptorStruct) SetVendorID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setVendorID:"), value)
}/* debug [instance_properties/setter]: vendorID */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTROperationalCredentialsClusterFabricDescriptorStruct */



