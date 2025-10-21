// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRMicrowaveOvenModeClusterModeTagStruct] class.
var (
	MTRMicrowaveOvenModeClusterModeTagStructClass     _MTRMicrowaveOvenModeClusterModeTagStructClass
	MTRMicrowaveOvenModeClusterModeTagStructClassOnce sync.Once
)

func getMTRMicrowaveOvenModeClusterModeTagStructClass() _MTRMicrowaveOvenModeClusterModeTagStructClass {
	MTRMicrowaveOvenModeClusterModeTagStructClassOnce.Do(func() {
		MTRMicrowaveOvenModeClusterModeTagStructClass = _MTRMicrowaveOvenModeClusterModeTagStructClass{objc.GetClass("MTRMicrowaveOvenModeClusterModeTagStruct")}
	})
	return MTRMicrowaveOvenModeClusterModeTagStructClass
}

type _MTRMicrowaveOvenModeClusterModeTagStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRMicrowaveOvenModeClusterModeTagStruct] class.
type IMTRMicrowaveOvenModeClusterModeTagStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMicrowaveOvenModeClusterModeTagStruct
type MTRMicrowaveOvenModeClusterModeTagStruct struct {
	objectivec.Object
}

// MTRMicrowaveOvenModeClusterModeTagStructFrom constructs a [MTRMicrowaveOvenModeClusterModeTagStruct] from an unsafe.Pointer.
func MTRMicrowaveOvenModeClusterModeTagStructFrom(ptr unsafe.Pointer) MTRMicrowaveOvenModeClusterModeTagStruct {
	return MTRMicrowaveOvenModeClusterModeTagStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRMicrowaveOvenModeClusterModeTagStructClass) Alloc() MTRMicrowaveOvenModeClusterModeTagStruct {
	rv := objc.Send[MTRMicrowaveOvenModeClusterModeTagStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRMicrowaveOvenModeClusterModeTagStructClass) New() MTRMicrowaveOvenModeClusterModeTagStruct {
	rv := objc.Send[MTRMicrowaveOvenModeClusterModeTagStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRMicrowaveOvenModeClusterModeTagStruct) Init() MTRMicrowaveOvenModeClusterModeTagStruct {
	rv := objc.Send[MTRMicrowaveOvenModeClusterModeTagStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRMicrowaveOvenModeClusterModeTagStruct) Autorelease() MTRMicrowaveOvenModeClusterModeTagStruct {
	rv := objc.Send[MTRMicrowaveOvenModeClusterModeTagStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRMicrowaveOvenModeClusterModeTagStruct creates a new MTRMicrowaveOvenModeClusterModeTagStruct instance.
func NewMTRMicrowaveOvenModeClusterModeTagStruct() MTRMicrowaveOvenModeClusterModeTagStruct {
	return getMTRMicrowaveOvenModeClusterModeTagStructClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMicrowaveOvenModeClusterModeTagStruct/mfgCode
func (m_ MTRMicrowaveOvenModeClusterModeTagStruct) MfgCode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("mfgCode"))
	return rv
}


// SetMfgCode sets the value of the mfgCode property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMicrowaveOvenModeClusterModeTagStruct/mfgCode
func (m_ MTRMicrowaveOvenModeClusterModeTagStruct) SetMfgCode(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMfgCode:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMicrowaveOvenModeClusterModeTagStruct/value
func (m_ MTRMicrowaveOvenModeClusterModeTagStruct) Value() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("value"))
	return rv
}


// SetValue sets the value of the value property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMicrowaveOvenModeClusterModeTagStruct/value
func (m_ MTRMicrowaveOvenModeClusterModeTagStruct) SetValue(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValue:"), value)
}



