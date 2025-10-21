// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRDeviceEnergyManagementModeClusterModeOptionStruct] class.
var (
	MTRDeviceEnergyManagementModeClusterModeOptionStructClass     _MTRDeviceEnergyManagementModeClusterModeOptionStructClass
	MTRDeviceEnergyManagementModeClusterModeOptionStructClassOnce sync.Once
)

func getMTRDeviceEnergyManagementModeClusterModeOptionStructClass() _MTRDeviceEnergyManagementModeClusterModeOptionStructClass {
	MTRDeviceEnergyManagementModeClusterModeOptionStructClassOnce.Do(func() {
		MTRDeviceEnergyManagementModeClusterModeOptionStructClass = _MTRDeviceEnergyManagementModeClusterModeOptionStructClass{objc.GetClass("MTRDeviceEnergyManagementModeClusterModeOptionStruct")}
	})
	return MTRDeviceEnergyManagementModeClusterModeOptionStructClass
}

type _MTRDeviceEnergyManagementModeClusterModeOptionStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRDeviceEnergyManagementModeClusterModeOptionStruct] class.
type IMTRDeviceEnergyManagementModeClusterModeOptionStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementModeClusterModeOptionStruct
type MTRDeviceEnergyManagementModeClusterModeOptionStruct struct {
	objectivec.Object
}

// MTRDeviceEnergyManagementModeClusterModeOptionStructFrom constructs a [MTRDeviceEnergyManagementModeClusterModeOptionStruct] from an unsafe.Pointer.
func MTRDeviceEnergyManagementModeClusterModeOptionStructFrom(ptr unsafe.Pointer) MTRDeviceEnergyManagementModeClusterModeOptionStruct {
	return MTRDeviceEnergyManagementModeClusterModeOptionStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDeviceEnergyManagementModeClusterModeOptionStructClass) Alloc() MTRDeviceEnergyManagementModeClusterModeOptionStruct {
	rv := objc.Send[MTRDeviceEnergyManagementModeClusterModeOptionStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDeviceEnergyManagementModeClusterModeOptionStructClass) New() MTRDeviceEnergyManagementModeClusterModeOptionStruct {
	rv := objc.Send[MTRDeviceEnergyManagementModeClusterModeOptionStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDeviceEnergyManagementModeClusterModeOptionStruct) Init() MTRDeviceEnergyManagementModeClusterModeOptionStruct {
	rv := objc.Send[MTRDeviceEnergyManagementModeClusterModeOptionStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDeviceEnergyManagementModeClusterModeOptionStruct) Autorelease() MTRDeviceEnergyManagementModeClusterModeOptionStruct {
	rv := objc.Send[MTRDeviceEnergyManagementModeClusterModeOptionStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDeviceEnergyManagementModeClusterModeOptionStruct creates a new MTRDeviceEnergyManagementModeClusterModeOptionStruct instance.
func NewMTRDeviceEnergyManagementModeClusterModeOptionStruct() MTRDeviceEnergyManagementModeClusterModeOptionStruct {
	return getMTRDeviceEnergyManagementModeClusterModeOptionStructClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementModeClusterModeOptionStruct/label
func (m_ MTRDeviceEnergyManagementModeClusterModeOptionStruct) Label() string {
	rv := objc.Send[string](m_.ID, objc.Sel("label"))
	return rv
}


// SetLabel sets the value of the label property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementModeClusterModeOptionStruct/label
func (m_ MTRDeviceEnergyManagementModeClusterModeOptionStruct) SetLabel(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLabel:"), objc.String(value))
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementModeClusterModeOptionStruct/mode
func (m_ MTRDeviceEnergyManagementModeClusterModeOptionStruct) Mode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("mode"))
	return rv
}


// SetMode sets the value of the mode property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementModeClusterModeOptionStruct/mode
func (m_ MTRDeviceEnergyManagementModeClusterModeOptionStruct) SetMode(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMode:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementModeClusterModeOptionStruct/modeTags
func (m_ MTRDeviceEnergyManagementModeClusterModeOptionStruct) ModeTags() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("modeTags"))
	return rv
}


// SetModeTags sets the value of the modeTags property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementModeClusterModeOptionStruct/modeTags
func (m_ MTRDeviceEnergyManagementModeClusterModeOptionStruct) SetModeTags(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setModeTags:"), value)
}


