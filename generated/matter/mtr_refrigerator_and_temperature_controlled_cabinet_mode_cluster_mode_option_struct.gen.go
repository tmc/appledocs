// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStruct] class.
var (
	MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStructClass     _MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStructClass
	MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStructClassOnce sync.Once
)

func getMTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStructClass() _MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStructClass {
	MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStructClassOnce.Do(func() {
		MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStructClass = _MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStructClass{objc.GetClass("MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStruct")}
	})
	return MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStructClass
}

type _MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStruct] class.
type IMTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStruct interface {
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
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStruct
type MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStruct struct {
	objectivec.Object
}

// MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStructFrom constructs a [MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStruct] from an unsafe.Pointer.
func MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStructFrom(ptr unsafe.Pointer) MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStruct {
	return MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStructClass) Alloc() MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStruct {
	rv := objc.Send[MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStructClass) New() MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStruct {
	rv := objc.Send[MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStruct) Init() MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStruct {
	rv := objc.Send[MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStruct) Autorelease() MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStruct {
	rv := objc.Send[MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStruct creates a new MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStruct instance.
func NewMTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStruct() MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStruct {
	return getMTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStructClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStruct/label
func (m_ MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStruct) Label() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("label"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStruct/label
func (m_ MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStruct) SetLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLabel:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStruct/mode
func (m_ MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStruct) Mode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("mode"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStruct/mode
func (m_ MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStruct) SetMode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMode:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStruct/modeTags
func (m_ MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStruct) ModeTags() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("modeTags"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStruct/modeTags
func (m_ MTRRefrigeratorAndTemperatureControlledCabinetModeClusterModeOptionStruct) SetModeTags(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setModeTags:"), value)
}



