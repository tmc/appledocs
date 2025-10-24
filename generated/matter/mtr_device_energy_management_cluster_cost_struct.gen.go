// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRDeviceEnergyManagementClusterCostStruct] class.
var (
	MTRDeviceEnergyManagementClusterCostStructClass     _MTRDeviceEnergyManagementClusterCostStructClass
	MTRDeviceEnergyManagementClusterCostStructClassOnce sync.Once
)

func getMTRDeviceEnergyManagementClusterCostStructClass() _MTRDeviceEnergyManagementClusterCostStructClass {
	MTRDeviceEnergyManagementClusterCostStructClassOnce.Do(func() {
		MTRDeviceEnergyManagementClusterCostStructClass = _MTRDeviceEnergyManagementClusterCostStructClass{objc.GetClass("MTRDeviceEnergyManagementClusterCostStruct")}
	})
	return MTRDeviceEnergyManagementClusterCostStructClass
}

type _MTRDeviceEnergyManagementClusterCostStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRDeviceEnergyManagementClusterCostStruct] class.
type IMTRDeviceEnergyManagementClusterCostStruct interface {
	objectivec.IObject
	// properties:
	CostType() objc.IObject /* cross-framework: NSNumber */
	SetCostType(value objc.IObject /* cross-framework: NSNumber */)
	Currency() objc.IObject /* cross-framework: NSNumber */
	SetCurrency(value objc.IObject /* cross-framework: NSNumber */)
	DecimalPoints() objc.IObject /* cross-framework: NSNumber */
	SetDecimalPoints(value objc.IObject /* cross-framework: NSNumber */)
	Value() objc.IObject /* cross-framework: NSNumber */
	SetValue(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterCostStruct
type MTRDeviceEnergyManagementClusterCostStruct struct {
	objectivec.Object
}

// MTRDeviceEnergyManagementClusterCostStructFrom constructs a [MTRDeviceEnergyManagementClusterCostStruct] from an unsafe.Pointer.
func MTRDeviceEnergyManagementClusterCostStructFrom(ptr unsafe.Pointer) MTRDeviceEnergyManagementClusterCostStruct {
	return MTRDeviceEnergyManagementClusterCostStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDeviceEnergyManagementClusterCostStructClass) Alloc() MTRDeviceEnergyManagementClusterCostStruct {
	rv := objc.Send[MTRDeviceEnergyManagementClusterCostStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDeviceEnergyManagementClusterCostStructClass) New() MTRDeviceEnergyManagementClusterCostStruct {
	rv := objc.Send[MTRDeviceEnergyManagementClusterCostStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDeviceEnergyManagementClusterCostStruct) Init() MTRDeviceEnergyManagementClusterCostStruct {
	rv := objc.Send[MTRDeviceEnergyManagementClusterCostStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDeviceEnergyManagementClusterCostStruct) Autorelease() MTRDeviceEnergyManagementClusterCostStruct {
	rv := objc.Send[MTRDeviceEnergyManagementClusterCostStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDeviceEnergyManagementClusterCostStruct creates a new MTRDeviceEnergyManagementClusterCostStruct instance.
func NewMTRDeviceEnergyManagementClusterCostStruct() MTRDeviceEnergyManagementClusterCostStruct {
	return getMTRDeviceEnergyManagementClusterCostStructClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterCostStruct/costType
func (m_ MTRDeviceEnergyManagementClusterCostStruct) CostType() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("costType"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterCostStruct/costType
func (m_ MTRDeviceEnergyManagementClusterCostStruct) SetCostType(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCostType:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterCostStruct/currency
func (m_ MTRDeviceEnergyManagementClusterCostStruct) Currency() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("currency"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterCostStruct/currency
func (m_ MTRDeviceEnergyManagementClusterCostStruct) SetCurrency(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCurrency:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterCostStruct/decimalPoints
func (m_ MTRDeviceEnergyManagementClusterCostStruct) DecimalPoints() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("decimalPoints"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterCostStruct/decimalPoints
func (m_ MTRDeviceEnergyManagementClusterCostStruct) SetDecimalPoints(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDecimalPoints:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterCostStruct/value
func (m_ MTRDeviceEnergyManagementClusterCostStruct) Value() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("value"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterCostStruct/value
func (m_ MTRDeviceEnergyManagementClusterCostStruct) SetValue(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValue:"), value)
}



