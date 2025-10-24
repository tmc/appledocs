// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTRThreadNetworkDiagnosticsClusterRouteTable */


/* debug [class_header]: Header for MTRThreadNetworkDiagnosticsClusterRouteTable */
// The class instance for the [MTRThreadNetworkDiagnosticsClusterRouteTable] class.
var (
	MTRThreadNetworkDiagnosticsClusterRouteTableClass     _MTRThreadNetworkDiagnosticsClusterRouteTableClass
	MTRThreadNetworkDiagnosticsClusterRouteTableClassOnce sync.Once
)

func getMTRThreadNetworkDiagnosticsClusterRouteTableClass() _MTRThreadNetworkDiagnosticsClusterRouteTableClass {
	MTRThreadNetworkDiagnosticsClusterRouteTableClassOnce.Do(func() {
		MTRThreadNetworkDiagnosticsClusterRouteTableClass = _MTRThreadNetworkDiagnosticsClusterRouteTableClass{objc.GetClass("MTRThreadNetworkDiagnosticsClusterRouteTable")}
	})
	return MTRThreadNetworkDiagnosticsClusterRouteTableClass
}

type _MTRThreadNetworkDiagnosticsClusterRouteTableClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRThreadNetworkDiagnosticsClusterRouteTable */
// An interface definition for the [MTRThreadNetworkDiagnosticsClusterRouteTable] class.
type IMTRThreadNetworkDiagnosticsClusterRouteTable interface {
	IMTRThreadNetworkDiagnosticsClusterRouteTableStruct
	
/* debug [class_interface_properties]: Properties for MTRThreadNetworkDiagnosticsClusterRouteTable */
	// properties:
	Age() objc.IObject /* cross-framework: NSNumber */
	SetAge(value objc.IObject /* cross-framework: NSNumber */)
	Allocated() objc.IObject /* cross-framework: NSNumber */
	SetAllocated(value objc.IObject /* cross-framework: NSNumber */)
	ExtAddress() objc.IObject /* cross-framework: NSNumber */
	SetExtAddress(value objc.IObject /* cross-framework: NSNumber */)
	LinkEstablished() objc.IObject /* cross-framework: NSNumber */
	SetLinkEstablished(value objc.IObject /* cross-framework: NSNumber */)
	LqiIn() objc.IObject /* cross-framework: NSNumber */
	SetLqiIn(value objc.IObject /* cross-framework: NSNumber */)
	LqiOut() objc.IObject /* cross-framework: NSNumber */
	SetLqiOut(value objc.IObject /* cross-framework: NSNumber */)
	NextHop() objc.IObject /* cross-framework: NSNumber */
	SetNextHop(value objc.IObject /* cross-framework: NSNumber */)
	PathCost() objc.IObject /* cross-framework: NSNumber */
	SetPathCost(value objc.IObject /* cross-framework: NSNumber */)
	Rloc16() objc.IObject /* cross-framework: NSNumber */
	SetRloc16(value objc.IObject /* cross-framework: NSNumber */)
	RouterId() objc.IObject /* cross-framework: NSNumber */
	SetRouterId(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRThreadNetworkDiagnosticsClusterRouteTable */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRThreadNetworkDiagnosticsClusterRouteTable */
// Alloc allocates a new instance without initialization.
func (mc _MTRThreadNetworkDiagnosticsClusterRouteTableClass) Alloc() MTRThreadNetworkDiagnosticsClusterRouteTable {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterRouteTable](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRThreadNetworkDiagnosticsClusterRouteTableClass) New() MTRThreadNetworkDiagnosticsClusterRouteTable {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterRouteTable](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTable) Init() MTRThreadNetworkDiagnosticsClusterRouteTable {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterRouteTable](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTable) Autorelease() MTRThreadNetworkDiagnosticsClusterRouteTable {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterRouteTable](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRThreadNetworkDiagnosticsClusterRouteTable creates a new MTRThreadNetworkDiagnosticsClusterRouteTable instance.
func NewMTRThreadNetworkDiagnosticsClusterRouteTable() MTRThreadNetworkDiagnosticsClusterRouteTable {
	return getMTRThreadNetworkDiagnosticsClusterRouteTableClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRThreadNetworkDiagnosticsClusterRouteTable */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterRouteTable
type MTRThreadNetworkDiagnosticsClusterRouteTable struct {
	MTRThreadNetworkDiagnosticsClusterRouteTableStruct
}

// MTRThreadNetworkDiagnosticsClusterRouteTableFrom constructs a [MTRThreadNetworkDiagnosticsClusterRouteTable] from an unsafe.Pointer.
func MTRThreadNetworkDiagnosticsClusterRouteTableFrom(ptr unsafe.Pointer) MTRThreadNetworkDiagnosticsClusterRouteTable {
	return MTRThreadNetworkDiagnosticsClusterRouteTable{
		MTRThreadNetworkDiagnosticsClusterRouteTableStruct: MTRThreadNetworkDiagnosticsClusterRouteTableStructFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRThreadNetworkDiagnosticsClusterRouteTable *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRThreadNetworkDiagnosticsClusterRouteTable */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRThreadNetworkDiagnosticsClusterRouteTable */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRThreadNetworkDiagnosticsClusterRouteTable */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRThreadNetworkDiagnosticsClusterRouteTable */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterRouteTable/age
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTable) Age() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("age"))
	return rv
}/* debug [instance_properties/getter]: age */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterRouteTable/age
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTable) SetAge(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAge:"), value)
}/* debug [instance_properties/setter]: age */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterRouteTable/allocated
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTable) Allocated() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("allocated"))
	return rv
}/* debug [instance_properties/getter]: allocated */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterRouteTable/allocated
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTable) SetAllocated(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAllocated:"), value)
}/* debug [instance_properties/setter]: allocated */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterRouteTable/extAddress
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTable) ExtAddress() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("extAddress"))
	return rv
}/* debug [instance_properties/getter]: extAddress */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterRouteTable/extAddress
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTable) SetExtAddress(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setExtAddress:"), value)
}/* debug [instance_properties/setter]: extAddress */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterRouteTable/linkEstablished
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTable) LinkEstablished() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("linkEstablished"))
	return rv
}/* debug [instance_properties/getter]: linkEstablished */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterRouteTable/linkEstablished
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTable) SetLinkEstablished(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLinkEstablished:"), value)
}/* debug [instance_properties/setter]: linkEstablished */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterRouteTable/lqiIn
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTable) LqiIn() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("lqiIn"))
	return rv
}/* debug [instance_properties/getter]: lqiIn */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterRouteTable/lqiIn
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTable) SetLqiIn(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLqiIn:"), value)
}/* debug [instance_properties/setter]: lqiIn */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterRouteTable/lqiOut
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTable) LqiOut() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("lqiOut"))
	return rv
}/* debug [instance_properties/getter]: lqiOut */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterRouteTable/lqiOut
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTable) SetLqiOut(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLqiOut:"), value)
}/* debug [instance_properties/setter]: lqiOut */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterRouteTable/nextHop
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTable) NextHop() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nextHop"))
	return rv
}/* debug [instance_properties/getter]: nextHop */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterRouteTable/nextHop
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTable) SetNextHop(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNextHop:"), value)
}/* debug [instance_properties/setter]: nextHop */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterRouteTable/pathCost
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTable) PathCost() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("pathCost"))
	return rv
}/* debug [instance_properties/getter]: pathCost */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterRouteTable/pathCost
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTable) SetPathCost(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPathCost:"), value)
}/* debug [instance_properties/setter]: pathCost */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterRouteTable/rloc16
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTable) Rloc16() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("rloc16"))
	return rv
}/* debug [instance_properties/getter]: rloc16 */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterRouteTable/rloc16
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTable) SetRloc16(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRloc16:"), value)
}/* debug [instance_properties/setter]: rloc16 */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterRouteTable/routerId
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTable) RouterId() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("routerId"))
	return rv
}/* debug [instance_properties/getter]: routerId */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterRouteTable/routerId
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTable) SetRouterId(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRouterId:"), value)
}/* debug [instance_properties/setter]: routerId */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRThreadNetworkDiagnosticsClusterRouteTable */



