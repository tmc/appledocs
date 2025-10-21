// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRDishwasherModeClusterModeTagStruct] class.
var (
	MTRDishwasherModeClusterModeTagStructClass     _MTRDishwasherModeClusterModeTagStructClass
	MTRDishwasherModeClusterModeTagStructClassOnce sync.Once
)

func getMTRDishwasherModeClusterModeTagStructClass() _MTRDishwasherModeClusterModeTagStructClass {
	MTRDishwasherModeClusterModeTagStructClassOnce.Do(func() {
		MTRDishwasherModeClusterModeTagStructClass = _MTRDishwasherModeClusterModeTagStructClass{objc.GetClass("MTRDishwasherModeClusterModeTagStruct")}
	})
	return MTRDishwasherModeClusterModeTagStructClass
}

type _MTRDishwasherModeClusterModeTagStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRDishwasherModeClusterModeTagStruct] class.
type IMTRDishwasherModeClusterModeTagStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherModeClusterModeTagStruct
type MTRDishwasherModeClusterModeTagStruct struct {
	objectivec.Object
}

// MTRDishwasherModeClusterModeTagStructFrom constructs a [MTRDishwasherModeClusterModeTagStruct] from an unsafe.Pointer.
func MTRDishwasherModeClusterModeTagStructFrom(ptr unsafe.Pointer) MTRDishwasherModeClusterModeTagStruct {
	return MTRDishwasherModeClusterModeTagStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDishwasherModeClusterModeTagStructClass) Alloc() MTRDishwasherModeClusterModeTagStruct {
	rv := objc.Send[MTRDishwasherModeClusterModeTagStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDishwasherModeClusterModeTagStructClass) New() MTRDishwasherModeClusterModeTagStruct {
	rv := objc.Send[MTRDishwasherModeClusterModeTagStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDishwasherModeClusterModeTagStruct) Init() MTRDishwasherModeClusterModeTagStruct {
	rv := objc.Send[MTRDishwasherModeClusterModeTagStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDishwasherModeClusterModeTagStruct) Autorelease() MTRDishwasherModeClusterModeTagStruct {
	rv := objc.Send[MTRDishwasherModeClusterModeTagStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDishwasherModeClusterModeTagStruct creates a new MTRDishwasherModeClusterModeTagStruct instance.
func NewMTRDishwasherModeClusterModeTagStruct() MTRDishwasherModeClusterModeTagStruct {
	return getMTRDishwasherModeClusterModeTagStructClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherModeClusterModeTagStruct/mfgCode
func (m_ MTRDishwasherModeClusterModeTagStruct) MfgCode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("mfgCode"))
	return rv
}


// SetMfgCode sets the value of the mfgCode property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherModeClusterModeTagStruct/mfgCode
func (m_ MTRDishwasherModeClusterModeTagStruct) SetMfgCode(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMfgCode:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherModeClusterModeTagStruct/value
func (m_ MTRDishwasherModeClusterModeTagStruct) Value() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("value"))
	return rv
}


// SetValue sets the value of the value property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherModeClusterModeTagStruct/value
func (m_ MTRDishwasherModeClusterModeTagStruct) SetValue(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValue:"), value)
}


