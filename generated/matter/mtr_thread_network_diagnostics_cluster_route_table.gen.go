// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [MTRThreadNetworkDiagnosticsClusterRouteTable] class.
type IMTRThreadNetworkDiagnosticsClusterRouteTable interface {
	IMTRThreadNetworkDiagnosticsClusterRouteTableStruct
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
	// methods:
}



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

// Alloc allocates a new instance without initialization.
func (mc _MTRThreadNetworkDiagnosticsClusterRouteTableClass) Alloc() MTRThreadNetworkDiagnosticsClusterRouteTable {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterRouteTable](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetable/age
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTable) Age() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("age"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetable/age
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTable) SetAge(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAge:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetable/allocated
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTable) Allocated() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("allocated"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetable/allocated
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTable) SetAllocated(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAllocated:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetable/extaddress
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTable) ExtAddress() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("extAddress"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetable/extaddress
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTable) SetExtAddress(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setExtAddress:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetable/linkestablished
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTable) LinkEstablished() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("linkEstablished"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetable/linkestablished
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTable) SetLinkEstablished(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLinkEstablished:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetable/lqiin
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTable) LqiIn() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("lqiIn"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetable/lqiin
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTable) SetLqiIn(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLqiIn:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetable/lqiout
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTable) LqiOut() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("lqiOut"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetable/lqiout
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTable) SetLqiOut(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLqiOut:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetable/nexthop
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTable) NextHop() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nextHop"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetable/nexthop
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTable) SetNextHop(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNextHop:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetable/pathcost
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTable) PathCost() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("pathCost"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetable/pathcost
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTable) SetPathCost(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPathCost:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetable/rloc16
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTable) Rloc16() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("rloc16"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetable/rloc16
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTable) SetRloc16(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRloc16:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetable/routerid
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTable) RouterId() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("routerId"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetable/routerid
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTable) SetRouterId(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRouterId:"), value)
}



