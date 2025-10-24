// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTRThreadNetworkDiagnosticsClusterNeighborTable */


/* debug [class_header]: Header for MTRThreadNetworkDiagnosticsClusterNeighborTable */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRThreadNetworkDiagnosticsClusterNeighborTable */
// An interface definition for the [MTRThreadNetworkDiagnosticsClusterNeighborTable] class.
type IMTRThreadNetworkDiagnosticsClusterNeighborTable interface {
	IMTRThreadNetworkDiagnosticsClusterNeighborTableStruct
	
/* debug [class_interface_properties]: Properties for MTRThreadNetworkDiagnosticsClusterNeighborTable */
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
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRThreadNetworkDiagnosticsClusterNeighborTable */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRThreadNetworkDiagnosticsClusterNeighborTable */
// Alloc allocates a new instance without initialization.
func (mc _MTRThreadNetworkDiagnosticsClusterNeighborTableClass) Alloc() MTRThreadNetworkDiagnosticsClusterNeighborTable {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterNeighborTable](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRThreadNetworkDiagnosticsClusterNeighborTable */


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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRThreadNetworkDiagnosticsClusterNeighborTable *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRThreadNetworkDiagnosticsClusterNeighborTable */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRThreadNetworkDiagnosticsClusterNeighborTable */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRThreadNetworkDiagnosticsClusterNeighborTable */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRThreadNetworkDiagnosticsClusterNeighborTable */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterNeighborTable/age
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) Age() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("age"))
	return rv
}/* debug [instance_properties/getter]: age */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterNeighborTable/age
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) SetAge(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAge:"), value)
}/* debug [instance_properties/setter]: age */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterNeighborTable/averageRssi
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) AverageRssi() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("averageRssi"))
	return rv
}/* debug [instance_properties/getter]: averageRssi */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterNeighborTable/averageRssi
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) SetAverageRssi(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAverageRssi:"), value)
}/* debug [instance_properties/setter]: averageRssi */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterNeighborTable/extAddress
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) ExtAddress() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("extAddress"))
	return rv
}/* debug [instance_properties/getter]: extAddress */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterNeighborTable/extAddress
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) SetExtAddress(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setExtAddress:"), value)
}/* debug [instance_properties/setter]: extAddress */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterNeighborTable/frameErrorRate
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) FrameErrorRate() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("frameErrorRate"))
	return rv
}/* debug [instance_properties/getter]: frameErrorRate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterNeighborTable/frameErrorRate
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) SetFrameErrorRate(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFrameErrorRate:"), value)
}/* debug [instance_properties/setter]: frameErrorRate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterNeighborTable/fullNetworkData
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) FullNetworkData() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("fullNetworkData"))
	return rv
}/* debug [instance_properties/getter]: fullNetworkData */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterNeighborTable/fullNetworkData
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) SetFullNetworkData(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFullNetworkData:"), value)
}/* debug [instance_properties/setter]: fullNetworkData */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterNeighborTable/fullThreadDevice
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) FullThreadDevice() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("fullThreadDevice"))
	return rv
}/* debug [instance_properties/getter]: fullThreadDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterNeighborTable/fullThreadDevice
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) SetFullThreadDevice(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFullThreadDevice:"), value)
}/* debug [instance_properties/setter]: fullThreadDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterNeighborTable/isChild
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) IsChild() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("isChild"))
	return rv
}/* debug [instance_properties/getter]: isChild */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterNeighborTable/isChild
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) SetIsChild(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsChild:"), value)
}/* debug [instance_properties/setter]: isChild */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterNeighborTable/lastRssi
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) LastRssi() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("lastRssi"))
	return rv
}/* debug [instance_properties/getter]: lastRssi */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterNeighborTable/lastRssi
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) SetLastRssi(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLastRssi:"), value)
}/* debug [instance_properties/setter]: lastRssi */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterNeighborTable/linkFrameCounter
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) LinkFrameCounter() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("linkFrameCounter"))
	return rv
}/* debug [instance_properties/getter]: linkFrameCounter */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterNeighborTable/linkFrameCounter
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) SetLinkFrameCounter(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLinkFrameCounter:"), value)
}/* debug [instance_properties/setter]: linkFrameCounter */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterNeighborTable/lqi
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) Lqi() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("lqi"))
	return rv
}/* debug [instance_properties/getter]: lqi */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterNeighborTable/lqi
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) SetLqi(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLqi:"), value)
}/* debug [instance_properties/setter]: lqi */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterNeighborTable/messageErrorRate
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) MessageErrorRate() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("messageErrorRate"))
	return rv
}/* debug [instance_properties/getter]: messageErrorRate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterNeighborTable/messageErrorRate
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) SetMessageErrorRate(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMessageErrorRate:"), value)
}/* debug [instance_properties/setter]: messageErrorRate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterNeighborTable/mleFrameCounter
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) MleFrameCounter() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("mleFrameCounter"))
	return rv
}/* debug [instance_properties/getter]: mleFrameCounter */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterNeighborTable/mleFrameCounter
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) SetMleFrameCounter(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMleFrameCounter:"), value)
}/* debug [instance_properties/setter]: mleFrameCounter */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterNeighborTable/rloc16
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) Rloc16() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("rloc16"))
	return rv
}/* debug [instance_properties/getter]: rloc16 */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterNeighborTable/rloc16
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) SetRloc16(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRloc16:"), value)
}/* debug [instance_properties/setter]: rloc16 */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterNeighborTable/rxOnWhenIdle
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) RxOnWhenIdle() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("rxOnWhenIdle"))
	return rv
}/* debug [instance_properties/getter]: rxOnWhenIdle */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterNeighborTable/rxOnWhenIdle
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTable) SetRxOnWhenIdle(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRxOnWhenIdle:"), value)
}/* debug [instance_properties/setter]: rxOnWhenIdle */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRThreadNetworkDiagnosticsClusterNeighborTable */



