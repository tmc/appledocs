// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRWaterHeaterModeClusterModeTagStruct] class.
var (
	MTRWaterHeaterModeClusterModeTagStructClass     _MTRWaterHeaterModeClusterModeTagStructClass
	MTRWaterHeaterModeClusterModeTagStructClassOnce sync.Once
)

func getMTRWaterHeaterModeClusterModeTagStructClass() _MTRWaterHeaterModeClusterModeTagStructClass {
	MTRWaterHeaterModeClusterModeTagStructClassOnce.Do(func() {
		MTRWaterHeaterModeClusterModeTagStructClass = _MTRWaterHeaterModeClusterModeTagStructClass{objc.GetClass("MTRWaterHeaterModeClusterModeTagStruct")}
	})
	return MTRWaterHeaterModeClusterModeTagStructClass
}

type _MTRWaterHeaterModeClusterModeTagStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRWaterHeaterModeClusterModeTagStruct] class.
type IMTRWaterHeaterModeClusterModeTagStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterModeClusterModeTagStruct
type MTRWaterHeaterModeClusterModeTagStruct struct {
	objectivec.Object
}

// MTRWaterHeaterModeClusterModeTagStructFrom constructs a [MTRWaterHeaterModeClusterModeTagStruct] from an unsafe.Pointer.
func MTRWaterHeaterModeClusterModeTagStructFrom(ptr unsafe.Pointer) MTRWaterHeaterModeClusterModeTagStruct {
	return MTRWaterHeaterModeClusterModeTagStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRWaterHeaterModeClusterModeTagStructClass) Alloc() MTRWaterHeaterModeClusterModeTagStruct {
	rv := objc.Send[MTRWaterHeaterModeClusterModeTagStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRWaterHeaterModeClusterModeTagStructClass) New() MTRWaterHeaterModeClusterModeTagStruct {
	rv := objc.Send[MTRWaterHeaterModeClusterModeTagStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRWaterHeaterModeClusterModeTagStruct) Init() MTRWaterHeaterModeClusterModeTagStruct {
	rv := objc.Send[MTRWaterHeaterModeClusterModeTagStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRWaterHeaterModeClusterModeTagStruct) Autorelease() MTRWaterHeaterModeClusterModeTagStruct {
	rv := objc.Send[MTRWaterHeaterModeClusterModeTagStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRWaterHeaterModeClusterModeTagStruct creates a new MTRWaterHeaterModeClusterModeTagStruct instance.
func NewMTRWaterHeaterModeClusterModeTagStruct() MTRWaterHeaterModeClusterModeTagStruct {
	return getMTRWaterHeaterModeClusterModeTagStructClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterModeClusterModeTagStruct/mfgCode
func (m_ MTRWaterHeaterModeClusterModeTagStruct) MfgCode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("mfgCode"))
	return rv
}


// SetMfgCode sets the value of the mfgCode property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterModeClusterModeTagStruct/mfgCode
func (m_ MTRWaterHeaterModeClusterModeTagStruct) SetMfgCode(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMfgCode:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterModeClusterModeTagStruct/value
func (m_ MTRWaterHeaterModeClusterModeTagStruct) Value() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("value"))
	return rv
}


// SetValue sets the value of the value property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterModeClusterModeTagStruct/value
func (m_ MTRWaterHeaterModeClusterModeTagStruct) SetValue(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValue:"), value)
}


