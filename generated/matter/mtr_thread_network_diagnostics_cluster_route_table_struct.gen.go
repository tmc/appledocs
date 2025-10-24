// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRThreadNetworkDiagnosticsClusterRouteTableStruct] class.
var (
	MTRThreadNetworkDiagnosticsClusterRouteTableStructClass     _MTRThreadNetworkDiagnosticsClusterRouteTableStructClass
	MTRThreadNetworkDiagnosticsClusterRouteTableStructClassOnce sync.Once
)

func getMTRThreadNetworkDiagnosticsClusterRouteTableStructClass() _MTRThreadNetworkDiagnosticsClusterRouteTableStructClass {
	MTRThreadNetworkDiagnosticsClusterRouteTableStructClassOnce.Do(func() {
		MTRThreadNetworkDiagnosticsClusterRouteTableStructClass = _MTRThreadNetworkDiagnosticsClusterRouteTableStructClass{objc.GetClass("MTRThreadNetworkDiagnosticsClusterRouteTableStruct")}
	})
	return MTRThreadNetworkDiagnosticsClusterRouteTableStructClass
}

type _MTRThreadNetworkDiagnosticsClusterRouteTableStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRThreadNetworkDiagnosticsClusterRouteTableStruct] class.
type IMTRThreadNetworkDiagnosticsClusterRouteTableStruct interface {
	objectivec.IObject
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
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterRouteTableStruct
type MTRThreadNetworkDiagnosticsClusterRouteTableStruct struct {
	objectivec.Object
}

// MTRThreadNetworkDiagnosticsClusterRouteTableStructFrom constructs a [MTRThreadNetworkDiagnosticsClusterRouteTableStruct] from an unsafe.Pointer.
func MTRThreadNetworkDiagnosticsClusterRouteTableStructFrom(ptr unsafe.Pointer) MTRThreadNetworkDiagnosticsClusterRouteTableStruct {
	return MTRThreadNetworkDiagnosticsClusterRouteTableStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRThreadNetworkDiagnosticsClusterRouteTableStructClass) Alloc() MTRThreadNetworkDiagnosticsClusterRouteTableStruct {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterRouteTableStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRThreadNetworkDiagnosticsClusterRouteTableStructClass) New() MTRThreadNetworkDiagnosticsClusterRouteTableStruct {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterRouteTableStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTableStruct) Init() MTRThreadNetworkDiagnosticsClusterRouteTableStruct {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterRouteTableStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTableStruct) Autorelease() MTRThreadNetworkDiagnosticsClusterRouteTableStruct {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterRouteTableStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRThreadNetworkDiagnosticsClusterRouteTableStruct creates a new MTRThreadNetworkDiagnosticsClusterRouteTableStruct instance.
func NewMTRThreadNetworkDiagnosticsClusterRouteTableStruct() MTRThreadNetworkDiagnosticsClusterRouteTableStruct {
	return getMTRThreadNetworkDiagnosticsClusterRouteTableStructClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetablestruct/age
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTableStruct) Age() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("age"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetablestruct/age
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTableStruct) SetAge(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAge:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetablestruct/allocated
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTableStruct) Allocated() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("allocated"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetablestruct/allocated
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTableStruct) SetAllocated(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAllocated:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetablestruct/extaddress
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTableStruct) ExtAddress() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("extAddress"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetablestruct/extaddress
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTableStruct) SetExtAddress(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setExtAddress:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetablestruct/linkestablished
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTableStruct) LinkEstablished() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("linkEstablished"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetablestruct/linkestablished
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTableStruct) SetLinkEstablished(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLinkEstablished:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetablestruct/lqiin
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTableStruct) LqiIn() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("lqiIn"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetablestruct/lqiin
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTableStruct) SetLqiIn(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLqiIn:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetablestruct/lqiout
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTableStruct) LqiOut() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("lqiOut"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetablestruct/lqiout
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTableStruct) SetLqiOut(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLqiOut:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetablestruct/nexthop
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTableStruct) NextHop() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nextHop"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetablestruct/nexthop
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTableStruct) SetNextHop(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNextHop:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetablestruct/pathcost
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTableStruct) PathCost() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("pathCost"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetablestruct/pathcost
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTableStruct) SetPathCost(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPathCost:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetablestruct/rloc16
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTableStruct) Rloc16() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("rloc16"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetablestruct/rloc16
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTableStruct) SetRloc16(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRloc16:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetablestruct/routerid
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTableStruct) RouterId() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("routerId"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetablestruct/routerid
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTableStruct) SetRouterId(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRouterId:"), value)
}



