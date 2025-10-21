// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTREnergyEVSEModeClusterModeOptionStruct] class.
var (
	MTREnergyEVSEModeClusterModeOptionStructClass     _MTREnergyEVSEModeClusterModeOptionStructClass
	MTREnergyEVSEModeClusterModeOptionStructClassOnce sync.Once
)

func getMTREnergyEVSEModeClusterModeOptionStructClass() _MTREnergyEVSEModeClusterModeOptionStructClass {
	MTREnergyEVSEModeClusterModeOptionStructClassOnce.Do(func() {
		MTREnergyEVSEModeClusterModeOptionStructClass = _MTREnergyEVSEModeClusterModeOptionStructClass{objc.GetClass("MTREnergyEVSEModeClusterModeOptionStruct")}
	})
	return MTREnergyEVSEModeClusterModeOptionStructClass
}

type _MTREnergyEVSEModeClusterModeOptionStructClass struct {
	class objc.Class
}

// An interface definition for the [MTREnergyEVSEModeClusterModeOptionStruct] class.
type IMTREnergyEVSEModeClusterModeOptionStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEModeClusterModeOptionStruct
type MTREnergyEVSEModeClusterModeOptionStruct struct {
	objectivec.Object
}

// MTREnergyEVSEModeClusterModeOptionStructFrom constructs a [MTREnergyEVSEModeClusterModeOptionStruct] from an unsafe.Pointer.
func MTREnergyEVSEModeClusterModeOptionStructFrom(ptr unsafe.Pointer) MTREnergyEVSEModeClusterModeOptionStruct {
	return MTREnergyEVSEModeClusterModeOptionStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTREnergyEVSEModeClusterModeOptionStructClass) Alloc() MTREnergyEVSEModeClusterModeOptionStruct {
	rv := objc.Send[MTREnergyEVSEModeClusterModeOptionStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTREnergyEVSEModeClusterModeOptionStructClass) New() MTREnergyEVSEModeClusterModeOptionStruct {
	rv := objc.Send[MTREnergyEVSEModeClusterModeOptionStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTREnergyEVSEModeClusterModeOptionStruct) Init() MTREnergyEVSEModeClusterModeOptionStruct {
	rv := objc.Send[MTREnergyEVSEModeClusterModeOptionStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTREnergyEVSEModeClusterModeOptionStruct) Autorelease() MTREnergyEVSEModeClusterModeOptionStruct {
	rv := objc.Send[MTREnergyEVSEModeClusterModeOptionStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTREnergyEVSEModeClusterModeOptionStruct creates a new MTREnergyEVSEModeClusterModeOptionStruct instance.
func NewMTREnergyEVSEModeClusterModeOptionStruct() MTREnergyEVSEModeClusterModeOptionStruct {
	return getMTREnergyEVSEModeClusterModeOptionStructClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEModeClusterModeOptionStruct/label
func (m_ MTREnergyEVSEModeClusterModeOptionStruct) Label() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("label"))
	return rv
}


// SetLabel sets the value of the label property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEModeClusterModeOptionStruct/label
func (m_ MTREnergyEVSEModeClusterModeOptionStruct) SetLabel(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLabel:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEModeClusterModeOptionStruct/mode
func (m_ MTREnergyEVSEModeClusterModeOptionStruct) Mode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("mode"))
	return rv
}


// SetMode sets the value of the mode property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEModeClusterModeOptionStruct/mode
func (m_ MTREnergyEVSEModeClusterModeOptionStruct) SetMode(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMode:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEModeClusterModeOptionStruct/modeTags
func (m_ MTREnergyEVSEModeClusterModeOptionStruct) ModeTags() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("modeTags"))
	return rv
}


// SetModeTags sets the value of the modeTags property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEModeClusterModeOptionStruct/modeTags
func (m_ MTREnergyEVSEModeClusterModeOptionStruct) SetModeTags(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setModeTags:"), value)
}


