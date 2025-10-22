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
	Age() foundation.Number
	SetAge(value foundation.INumber)
	Allocated() foundation.Number
	SetAllocated(value foundation.INumber)
	ExtAddress() foundation.Number
	SetExtAddress(value foundation.INumber)
	LinkEstablished() foundation.Number
	SetLinkEstablished(value foundation.INumber)
	LqiIn() foundation.Number
	SetLqiIn(value foundation.INumber)
	LqiOut() foundation.Number
	SetLqiOut(value foundation.INumber)
	NextHop() foundation.Number
	SetNextHop(value foundation.INumber)
	PathCost() foundation.Number
	SetPathCost(value foundation.INumber)
	Rloc16() foundation.Number
	SetRloc16(value foundation.INumber)
	RouterId() foundation.Number
	SetRouterId(value foundation.INumber)
}

//
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


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetable/age
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTable) Age() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("age"))
	return rv
}


// SetAge sets the value of the age property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetable/age
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTable) SetAge(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAge:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetable/allocated
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTable) Allocated() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("allocated"))
	return rv
}


// SetAllocated sets the value of the allocated property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetable/allocated
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTable) SetAllocated(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAllocated:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetable/extaddress
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTable) ExtAddress() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("extAddress"))
	return rv
}


// SetExtAddress sets the value of the extAddress property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetable/extaddress
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTable) SetExtAddress(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setExtAddress:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetable/linkestablished
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTable) LinkEstablished() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("linkEstablished"))
	return rv
}


// SetLinkEstablished sets the value of the linkEstablished property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetable/linkestablished
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTable) SetLinkEstablished(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLinkEstablished:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetable/lqiin
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTable) LqiIn() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("lqiIn"))
	return rv
}


// SetLqiIn sets the value of the lqiIn property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetable/lqiin
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTable) SetLqiIn(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLqiIn:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetable/lqiout
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTable) LqiOut() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("lqiOut"))
	return rv
}


// SetLqiOut sets the value of the lqiOut property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetable/lqiout
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTable) SetLqiOut(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLqiOut:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetable/nexthop
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTable) NextHop() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("nextHop"))
	return rv
}


// SetNextHop sets the value of the nextHop property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetable/nexthop
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTable) SetNextHop(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNextHop:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetable/pathcost
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTable) PathCost() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("pathCost"))
	return rv
}


// SetPathCost sets the value of the pathCost property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetable/pathcost
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTable) SetPathCost(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPathCost:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetable/rloc16
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTable) Rloc16() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("rloc16"))
	return rv
}


// SetRloc16 sets the value of the rloc16 property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetable/rloc16
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTable) SetRloc16(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRloc16:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetable/routerid
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTable) RouterId() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("routerId"))
	return rv
}


// SetRouterId sets the value of the routerId property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterroutetable/routerid
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTable) SetRouterId(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRouterId:"), value)
}



