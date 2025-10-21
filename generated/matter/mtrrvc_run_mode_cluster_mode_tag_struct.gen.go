// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRRVCRunModeClusterModeTagStruct] class.
var (
	MTRRVCRunModeClusterModeTagStructClass     _MTRRVCRunModeClusterModeTagStructClass
	MTRRVCRunModeClusterModeTagStructClassOnce sync.Once
)

func getMTRRVCRunModeClusterModeTagStructClass() _MTRRVCRunModeClusterModeTagStructClass {
	MTRRVCRunModeClusterModeTagStructClassOnce.Do(func() {
		MTRRVCRunModeClusterModeTagStructClass = _MTRRVCRunModeClusterModeTagStructClass{objc.GetClass("MTRRVCRunModeClusterModeTagStruct")}
	})
	return MTRRVCRunModeClusterModeTagStructClass
}

type _MTRRVCRunModeClusterModeTagStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRRVCRunModeClusterModeTagStruct] class.
type IMTRRVCRunModeClusterModeTagStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRRVCRunModeClusterModeTagStruct
type MTRRVCRunModeClusterModeTagStruct struct {
	objectivec.Object
}

// MTRRVCRunModeClusterModeTagStructFrom constructs a [MTRRVCRunModeClusterModeTagStruct] from an unsafe.Pointer.
func MTRRVCRunModeClusterModeTagStructFrom(ptr unsafe.Pointer) MTRRVCRunModeClusterModeTagStruct {
	return MTRRVCRunModeClusterModeTagStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRRVCRunModeClusterModeTagStructClass) Alloc() MTRRVCRunModeClusterModeTagStruct {
	rv := objc.Send[MTRRVCRunModeClusterModeTagStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRRVCRunModeClusterModeTagStructClass) New() MTRRVCRunModeClusterModeTagStruct {
	rv := objc.Send[MTRRVCRunModeClusterModeTagStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRRVCRunModeClusterModeTagStruct) Init() MTRRVCRunModeClusterModeTagStruct {
	rv := objc.Send[MTRRVCRunModeClusterModeTagStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRRVCRunModeClusterModeTagStruct) Autorelease() MTRRVCRunModeClusterModeTagStruct {
	rv := objc.Send[MTRRVCRunModeClusterModeTagStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRRVCRunModeClusterModeTagStruct creates a new MTRRVCRunModeClusterModeTagStruct instance.
func NewMTRRVCRunModeClusterModeTagStruct() MTRRVCRunModeClusterModeTagStruct {
	return getMTRRVCRunModeClusterModeTagStructClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrrvcrunmodeclustermodetagstruct/mfgcode
func (m_ MTRRVCRunModeClusterModeTagStruct) MfgCode() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("mfgCode"))
	return rv
}


// SetMfgCode sets the value of the mfgCode property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrrvcrunmodeclustermodetagstruct/mfgcode
func (m_ MTRRVCRunModeClusterModeTagStruct) SetMfgCode(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMfgCode:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrrvcrunmodeclustermodetagstruct/value
func (m_ MTRRVCRunModeClusterModeTagStruct) Value() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("value"))
	return rv
}


// SetValue sets the value of the value property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrrvcrunmodeclustermodetagstruct/value
func (m_ MTRRVCRunModeClusterModeTagStruct) SetValue(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValue:"), value)
}



