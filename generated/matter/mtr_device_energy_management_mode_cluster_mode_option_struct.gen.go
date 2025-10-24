// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
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
	// properties:
	Label() objc.IObject /* cross-framework: NSString */
	SetLabel(value objc.IObject /* cross-framework: NSString */)
	Mode() objc.IObject /* cross-framework: NSNumber */
	SetMode(value objc.IObject /* cross-framework: NSNumber */)
	ModeTags() objc.IObject /* cross-framework: NSArray */
	SetModeTags(value objc.IObject /* cross-framework: NSArray */)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementModeClusterModeOptionStruct/label
func (m_ MTRDeviceEnergyManagementModeClusterModeOptionStruct) Label() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("label"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementModeClusterModeOptionStruct/label
func (m_ MTRDeviceEnergyManagementModeClusterModeOptionStruct) SetLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLabel:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementModeClusterModeOptionStruct/mode
func (m_ MTRDeviceEnergyManagementModeClusterModeOptionStruct) Mode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("mode"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementModeClusterModeOptionStruct/mode
func (m_ MTRDeviceEnergyManagementModeClusterModeOptionStruct) SetMode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMode:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementModeClusterModeOptionStruct/modeTags
func (m_ MTRDeviceEnergyManagementModeClusterModeOptionStruct) ModeTags() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("modeTags"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceEnergyManagementModeClusterModeOptionStruct/modeTags
func (m_ MTRDeviceEnergyManagementModeClusterModeOptionStruct) SetModeTags(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setModeTags:"), value)
}



