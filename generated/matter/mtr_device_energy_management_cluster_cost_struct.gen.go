// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
}

//
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


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterCostStruct/costType
func (m_ MTRDeviceEnergyManagementClusterCostStruct) CostType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("costType"))
	return rv
}


// SetCostType sets the value of the costType property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterCostStruct/costType
func (m_ MTRDeviceEnergyManagementClusterCostStruct) SetCostType(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCostType:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterCostStruct/currency
func (m_ MTRDeviceEnergyManagementClusterCostStruct) Currency() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("currency"))
	return rv
}


// SetCurrency sets the value of the currency property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterCostStruct/currency
func (m_ MTRDeviceEnergyManagementClusterCostStruct) SetCurrency(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCurrency:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterCostStruct/decimalPoints
func (m_ MTRDeviceEnergyManagementClusterCostStruct) DecimalPoints() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("decimalPoints"))
	return rv
}


// SetDecimalPoints sets the value of the decimalPoints property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterCostStruct/decimalPoints
func (m_ MTRDeviceEnergyManagementClusterCostStruct) SetDecimalPoints(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDecimalPoints:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterCostStruct/value
func (m_ MTRDeviceEnergyManagementClusterCostStruct) Value() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("value"))
	return rv
}


// SetValue sets the value of the value property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementClusterCostStruct/value
func (m_ MTRDeviceEnergyManagementClusterCostStruct) SetValue(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValue:"), value)
}


