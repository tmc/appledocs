// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTREnergyEVSEModeClusterModeTagStruct] class.
var (
	MTREnergyEVSEModeClusterModeTagStructClass     _MTREnergyEVSEModeClusterModeTagStructClass
	MTREnergyEVSEModeClusterModeTagStructClassOnce sync.Once
)

func getMTREnergyEVSEModeClusterModeTagStructClass() _MTREnergyEVSEModeClusterModeTagStructClass {
	MTREnergyEVSEModeClusterModeTagStructClassOnce.Do(func() {
		MTREnergyEVSEModeClusterModeTagStructClass = _MTREnergyEVSEModeClusterModeTagStructClass{objc.GetClass("MTREnergyEVSEModeClusterModeTagStruct")}
	})
	return MTREnergyEVSEModeClusterModeTagStructClass
}

type _MTREnergyEVSEModeClusterModeTagStructClass struct {
	class objc.Class
}

// An interface definition for the [MTREnergyEVSEModeClusterModeTagStruct] class.
type IMTREnergyEVSEModeClusterModeTagStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEModeClusterModeTagStruct
type MTREnergyEVSEModeClusterModeTagStruct struct {
	objectivec.Object
}

// MTREnergyEVSEModeClusterModeTagStructFrom constructs a [MTREnergyEVSEModeClusterModeTagStruct] from an unsafe.Pointer.
func MTREnergyEVSEModeClusterModeTagStructFrom(ptr unsafe.Pointer) MTREnergyEVSEModeClusterModeTagStruct {
	return MTREnergyEVSEModeClusterModeTagStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTREnergyEVSEModeClusterModeTagStructClass) Alloc() MTREnergyEVSEModeClusterModeTagStruct {
	rv := objc.Send[MTREnergyEVSEModeClusterModeTagStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTREnergyEVSEModeClusterModeTagStructClass) New() MTREnergyEVSEModeClusterModeTagStruct {
	rv := objc.Send[MTREnergyEVSEModeClusterModeTagStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTREnergyEVSEModeClusterModeTagStruct) Init() MTREnergyEVSEModeClusterModeTagStruct {
	rv := objc.Send[MTREnergyEVSEModeClusterModeTagStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTREnergyEVSEModeClusterModeTagStruct) Autorelease() MTREnergyEVSEModeClusterModeTagStruct {
	rv := objc.Send[MTREnergyEVSEModeClusterModeTagStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTREnergyEVSEModeClusterModeTagStruct creates a new MTREnergyEVSEModeClusterModeTagStruct instance.
func NewMTREnergyEVSEModeClusterModeTagStruct() MTREnergyEVSEModeClusterModeTagStruct {
	return getMTREnergyEVSEModeClusterModeTagStructClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEModeClusterModeTagStruct/mfgCode
func (m_ MTREnergyEVSEModeClusterModeTagStruct) MfgCode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("mfgCode"))
	return rv
}


// SetMfgCode sets the value of the mfgCode property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEModeClusterModeTagStruct/mfgCode
func (m_ MTREnergyEVSEModeClusterModeTagStruct) SetMfgCode(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMfgCode:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEModeClusterModeTagStruct/value
func (m_ MTREnergyEVSEModeClusterModeTagStruct) Value() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("value"))
	return rv
}


// SetValue sets the value of the value property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEModeClusterModeTagStruct/value
func (m_ MTREnergyEVSEModeClusterModeTagStruct) SetValue(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValue:"), value)
}



