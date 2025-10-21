// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTROvenModeClusterModeTagStruct] class.
var (
	MTROvenModeClusterModeTagStructClass     _MTROvenModeClusterModeTagStructClass
	MTROvenModeClusterModeTagStructClassOnce sync.Once
)

func getMTROvenModeClusterModeTagStructClass() _MTROvenModeClusterModeTagStructClass {
	MTROvenModeClusterModeTagStructClassOnce.Do(func() {
		MTROvenModeClusterModeTagStructClass = _MTROvenModeClusterModeTagStructClass{objc.GetClass("MTROvenModeClusterModeTagStruct")}
	})
	return MTROvenModeClusterModeTagStructClass
}

type _MTROvenModeClusterModeTagStructClass struct {
	class objc.Class
}

// An interface definition for the [MTROvenModeClusterModeTagStruct] class.
type IMTROvenModeClusterModeTagStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenModeClusterModeTagStruct
type MTROvenModeClusterModeTagStruct struct {
	objectivec.Object
}

// MTROvenModeClusterModeTagStructFrom constructs a [MTROvenModeClusterModeTagStruct] from an unsafe.Pointer.
func MTROvenModeClusterModeTagStructFrom(ptr unsafe.Pointer) MTROvenModeClusterModeTagStruct {
	return MTROvenModeClusterModeTagStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROvenModeClusterModeTagStructClass) Alloc() MTROvenModeClusterModeTagStruct {
	rv := objc.Send[MTROvenModeClusterModeTagStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROvenModeClusterModeTagStructClass) New() MTROvenModeClusterModeTagStruct {
	rv := objc.Send[MTROvenModeClusterModeTagStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROvenModeClusterModeTagStruct) Init() MTROvenModeClusterModeTagStruct {
	rv := objc.Send[MTROvenModeClusterModeTagStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROvenModeClusterModeTagStruct) Autorelease() MTROvenModeClusterModeTagStruct {
	rv := objc.Send[MTROvenModeClusterModeTagStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROvenModeClusterModeTagStruct creates a new MTROvenModeClusterModeTagStruct instance.
func NewMTROvenModeClusterModeTagStruct() MTROvenModeClusterModeTagStruct {
	return getMTROvenModeClusterModeTagStructClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenModeClusterModeTagStruct/mfgCode
func (m_ MTROvenModeClusterModeTagStruct) MfgCode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("mfgCode"))
	return rv
}


// SetMfgCode sets the value of the mfgCode property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenModeClusterModeTagStruct/mfgCode
func (m_ MTROvenModeClusterModeTagStruct) SetMfgCode(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMfgCode:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenModeClusterModeTagStruct/value
func (m_ MTROvenModeClusterModeTagStruct) Value() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("value"))
	return rv
}


// SetValue sets the value of the value property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenModeClusterModeTagStruct/value
func (m_ MTROvenModeClusterModeTagStruct) SetValue(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValue:"), value)
}


