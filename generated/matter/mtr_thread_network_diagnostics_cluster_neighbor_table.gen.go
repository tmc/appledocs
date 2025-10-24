// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRThreadNetworkDiagnosticsClusterNeighborTable] class.
var (
	MTRThreadNetworkDiagnosticsClusterNeighborTableClass     _MTRThreadNetworkDiagnosticsClusterNeighborTableClass
	MTRThreadNetworkDiagnosticsClusterNeighborTableClassOnce sync.Once
)

func getMTRThreadNetworkDiagnosticsClusterNeighborTableClass() _MTRThreadNetworkDiagnosticsClusterNeighborTableClass {
	MTRThreadNetworkDiagnosticsClusterNeighborTableClassOnce.Do(func() {
		MTRThreadNetworkDiagnosticsClusterNeighborTableClass = _MTRThreadNetworkDiagnosticsClusterNeighborTableClass{objc.GetClass("MTRThreadNetworkDiagnosticsClusterNeighborTable")}
	})
	return MTRThreadNetworkDiagnosticsClusterNeighborTableClass
}

type _MTRThreadNetworkDiagnosticsClusterNeighborTableClass struct {
	class objc.Class
}

// An interface definition for the [MTRThreadNetworkDiagnosticsClusterNeighborTable] class.
type IMTRThreadNetworkDiagnosticsClusterNeighborTable interface {
	IMTRThreadNetworkDiagnosticsClusterNeighborTableStruct
	// properties:
	Age() objc.IObject /* cross-framework: NSNumber */
	SetAge(value objc.IObject /* cross-framework: NSNumber */)
	AverageRssi() objc.IObject /* cross-framework: NSNumber */
	SetAverageRssi(value objc.IObject /* cross-framework: NSNumber */)
	ExtAddress() objc.IObject /* cross-framework: NSNumber */
	SetExtAddress(value objc.IObject /* cross-framework: NSNumber */)
	FrameErrorRate() objc.IObject /* cross-framework: NSNumber */
	SetFrameErrorRate(value objc.IObject /* cross-framework: NSNumber */)
	FullNetworkData() objc.IObject /* cross-framework: NSNumber */
	SetFullNetworkData(value objc.IObject /* cross-framework: NSNumber */)
	FullThreadDevice() objc.IObject /* cross-framework: NSNumber */
	SetFullThreadDevice(value objc.IObject /* cross-framework: NSNumber */)
	IsChild() objc.IObject /* cross-framework: NSNumber */
	SetIsChild(value objc.IObject /* cross-framework: NSNumber */)
	LastRssi() objc.IObject /* cross-framework: NSNumber */
	SetLastRssi(value objc.IObject /* cross-framework: NSNumber */)
	LinkFrameCounter() objc.IObject /* cross-framework: NSNumber */
	SetLinkFrameCounter(value objc.IObject /* cross-framework: NSNumber */)
	Lqi() objc.IObject /* cross-framework: NSNumber */
	SetLqi(value objc.IObject /* cross-framework: NSNumber */)
	MessageErrorRate() objc.IObject /* cross-framework: NSNumber */
	SetMessageErrorRate(value objc.IObject /* cross-framework: NSNumber */)
	MleFrameCounter() objc.IObject /* cross-framework: NSNumber */
	SetMleFrameCounter(value objc.IObject /* cross-framework: NSNumber */)
	Rloc16() objc.IObject /* cross-framework: NSNumber */
	SetRloc16(value objc.IObject /* cross-framework: NSNumber */)
	RxOnWhenIdle() objc.IObject /* cross-framework: NSNumber */
	SetRxOnWhenIdle(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterNeighborTable
type MTRThreadNetworkDiagnosticsClusterNeighborTable struct {
	MTRThreadNetworkDiagnosticsClusterNeighborTableStruct
}

// MTRThreadNetworkDiagnosticsClusterNeighborTableFrom constructs a [MTRThreadNetworkDiagnosticsClusterNeighborTable] from an unsafe.Pointer.
func MTRThreadNetworkDiagnosticsClusterNeighborTableFrom(ptr unsafe.Pointer) MTRThreadNetworkDiagnosticsClusterNeighborTable {
	return MTRThreadNetworkDiagnosticsClusterNeighborTable{
		MTRThreadNetworkDiagnosticsClusterNeighborTableStruct: MTRThreadNetworkDiagnosticsClusterNeighborTableStructFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRThreadNetworkDiagnosticsClusterNeighborTableClass) Alloc() MTRThreadNetworkDiagnosticsClusterNeighborTable {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterNeighborTable](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRThreadNetworkDiagnosticsClusterNeighborTableClass) New() MTRThreadNetworkDiagnosticsClusterNeighborTable {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterNeighborTable](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) Init() MTRThreadNetworkDiagnosticsClusterNeighborTable {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterNeighborTable](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) Autorelease() MTRThreadNetworkDiagnosticsClusterNeighborTable {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterNeighborTable](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRThreadNetworkDiagnosticsClusterNeighborTable creates a new MTRThreadNetworkDiagnosticsClusterNeighborTable instance.
func NewMTRThreadNetworkDiagnosticsClusterNeighborTable() MTRThreadNetworkDiagnosticsClusterNeighborTable {
	return getMTRThreadNetworkDiagnosticsClusterNeighborTableClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortable/age
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) Age() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("age"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortable/age
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) SetAge(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAge:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortable/averagerssi
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) AverageRssi() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("averageRssi"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortable/averagerssi
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) SetAverageRssi(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAverageRssi:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortable/extaddress
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) ExtAddress() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("extAddress"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortable/extaddress
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) SetExtAddress(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setExtAddress:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortable/frameerrorrate
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) FrameErrorRate() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("frameErrorRate"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortable/frameerrorrate
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) SetFrameErrorRate(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFrameErrorRate:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortable/fullnetworkdata
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) FullNetworkData() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("fullNetworkData"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortable/fullnetworkdata
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) SetFullNetworkData(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFullNetworkData:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortable/fullthreaddevice
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) FullThreadDevice() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("fullThreadDevice"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortable/fullthreaddevice
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) SetFullThreadDevice(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFullThreadDevice:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortable/ischild
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) IsChild() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("isChild"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortable/ischild
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) SetIsChild(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsChild:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortable/lastrssi
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) LastRssi() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("lastRssi"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortable/lastrssi
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) SetLastRssi(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLastRssi:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortable/linkframecounter
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) LinkFrameCounter() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("linkFrameCounter"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortable/linkframecounter
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) SetLinkFrameCounter(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLinkFrameCounter:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortable/lqi
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) Lqi() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("lqi"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortable/lqi
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) SetLqi(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLqi:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortable/messageerrorrate
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) MessageErrorRate() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("messageErrorRate"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortable/messageerrorrate
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) SetMessageErrorRate(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMessageErrorRate:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortable/mleframecounter
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) MleFrameCounter() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("mleFrameCounter"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortable/mleframecounter
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) SetMleFrameCounter(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMleFrameCounter:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortable/rloc16
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) Rloc16() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("rloc16"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortable/rloc16
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) SetRloc16(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRloc16:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortable/rxonwhenidle
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) RxOnWhenIdle() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("rxOnWhenIdle"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrthreadnetworkdiagnosticsclusterneighbortable/rxonwhenidle
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) SetRxOnWhenIdle(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRxOnWhenIdle:"), value)
}



