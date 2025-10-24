// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRThreadNetworkDirectoryClusterThreadNetworkStruct */


/* debug [class_header]: Header for MTRThreadNetworkDirectoryClusterThreadNetworkStruct */
// The class instance for the [MTRThreadNetworkDirectoryClusterThreadNetworkStruct] class.
var (
	MTRThreadNetworkDirectoryClusterThreadNetworkStructClass     _MTRThreadNetworkDirectoryClusterThreadNetworkStructClass
	MTRThreadNetworkDirectoryClusterThreadNetworkStructClassOnce sync.Once
)

func getMTRThreadNetworkDirectoryClusterThreadNetworkStructClass() _MTRThreadNetworkDirectoryClusterThreadNetworkStructClass {
	MTRThreadNetworkDirectoryClusterThreadNetworkStructClassOnce.Do(func() {
		MTRThreadNetworkDirectoryClusterThreadNetworkStructClass = _MTRThreadNetworkDirectoryClusterThreadNetworkStructClass{objc.GetClass("MTRThreadNetworkDirectoryClusterThreadNetworkStruct")}
	})
	return MTRThreadNetworkDirectoryClusterThreadNetworkStructClass
}

type _MTRThreadNetworkDirectoryClusterThreadNetworkStructClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRThreadNetworkDirectoryClusterThreadNetworkStruct */
// An interface definition for the [MTRThreadNetworkDirectoryClusterThreadNetworkStruct] class.
type IMTRThreadNetworkDirectoryClusterThreadNetworkStruct interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRThreadNetworkDirectoryClusterThreadNetworkStruct */
	// properties:
	ActiveTimestamp() objc.IObject /* cross-framework: NSNumber */
	SetActiveTimestamp(value objc.IObject /* cross-framework: NSNumber */)
	Channel() objc.IObject /* cross-framework: NSNumber */
	SetChannel(value objc.IObject /* cross-framework: NSNumber */)
	ExtendedPanID() foundation.Data
	SetExtendedPanID(value foundation.Data)
	NetworkName() objc.IObject /* cross-framework: NSString */
	SetNetworkName(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRThreadNetworkDirectoryClusterThreadNetworkStruct */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRThreadNetworkDirectoryClusterThreadNetworkStruct */
// Alloc allocates a new instance without initialization.
func (mc _MTRThreadNetworkDirectoryClusterThreadNetworkStructClass) Alloc() MTRThreadNetworkDirectoryClusterThreadNetworkStruct {
	rv := objc.Send[MTRThreadNetworkDirectoryClusterThreadNetworkStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRThreadNetworkDirectoryClusterThreadNetworkStructClass) New() MTRThreadNetworkDirectoryClusterThreadNetworkStruct {
	rv := objc.Send[MTRThreadNetworkDirectoryClusterThreadNetworkStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRThreadNetworkDirectoryClusterThreadNetworkStruct) Init() MTRThreadNetworkDirectoryClusterThreadNetworkStruct {
	rv := objc.Send[MTRThreadNetworkDirectoryClusterThreadNetworkStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRThreadNetworkDirectoryClusterThreadNetworkStruct) Autorelease() MTRThreadNetworkDirectoryClusterThreadNetworkStruct {
	rv := objc.Send[MTRThreadNetworkDirectoryClusterThreadNetworkStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRThreadNetworkDirectoryClusterThreadNetworkStruct creates a new MTRThreadNetworkDirectoryClusterThreadNetworkStruct instance.
func NewMTRThreadNetworkDirectoryClusterThreadNetworkStruct() MTRThreadNetworkDirectoryClusterThreadNetworkStruct {
	return getMTRThreadNetworkDirectoryClusterThreadNetworkStructClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRThreadNetworkDirectoryClusterThreadNetworkStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDirectoryClusterThreadNetworkStruct
type MTRThreadNetworkDirectoryClusterThreadNetworkStruct struct {
	objectivec.Object
}

// MTRThreadNetworkDirectoryClusterThreadNetworkStructFrom constructs a [MTRThreadNetworkDirectoryClusterThreadNetworkStruct] from an unsafe.Pointer.
func MTRThreadNetworkDirectoryClusterThreadNetworkStructFrom(ptr unsafe.Pointer) MTRThreadNetworkDirectoryClusterThreadNetworkStruct {
	return MTRThreadNetworkDirectoryClusterThreadNetworkStruct{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRThreadNetworkDirectoryClusterThreadNetworkStruct *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRThreadNetworkDirectoryClusterThreadNetworkStruct */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRThreadNetworkDirectoryClusterThreadNetworkStruct */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRThreadNetworkDirectoryClusterThreadNetworkStruct */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRThreadNetworkDirectoryClusterThreadNetworkStruct */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDirectoryClusterThreadNetworkStruct/activeTimestamp
func (m_ MTRThreadNetworkDirectoryClusterThreadNetworkStruct) ActiveTimestamp() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("activeTimestamp"))
	return rv
}/* debug [instance_properties/getter]: activeTimestamp */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDirectoryClusterThreadNetworkStruct/activeTimestamp
func (m_ MTRThreadNetworkDirectoryClusterThreadNetworkStruct) SetActiveTimestamp(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setActiveTimestamp:"), value)
}/* debug [instance_properties/setter]: activeTimestamp */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdirectoryclusterthreadnetworkstruct/channel
func (m_ MTRThreadNetworkDirectoryClusterThreadNetworkStruct) Channel() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("channel"))
	return rv
}/* debug [instance_properties/getter]: channel */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdirectoryclusterthreadnetworkstruct/channel
func (m_ MTRThreadNetworkDirectoryClusterThreadNetworkStruct) SetChannel(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setChannel:"), value)
}/* debug [instance_properties/setter]: channel */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdirectoryclusterthreadnetworkstruct/extendedpanid
func (m_ MTRThreadNetworkDirectoryClusterThreadNetworkStruct) ExtendedPanID() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("extendedPanID"))
	return rv
}/* debug [instance_properties/getter]: extendedPanID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdirectoryclusterthreadnetworkstruct/extendedpanid
func (m_ MTRThreadNetworkDirectoryClusterThreadNetworkStruct) SetExtendedPanID(value foundation.Data) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setExtendedPanID:"), value)
}/* debug [instance_properties/setter]: extendedPanID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdirectoryclusterthreadnetworkstruct/networkname
func (m_ MTRThreadNetworkDirectoryClusterThreadNetworkStruct) NetworkName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("networkName"))
	return rv
}/* debug [instance_properties/getter]: networkName */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdirectoryclusterthreadnetworkstruct/networkname
func (m_ MTRThreadNetworkDirectoryClusterThreadNetworkStruct) SetNetworkName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNetworkName:"), value)
}/* debug [instance_properties/setter]: networkName */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRThreadNetworkDirectoryClusterThreadNetworkStruct */



