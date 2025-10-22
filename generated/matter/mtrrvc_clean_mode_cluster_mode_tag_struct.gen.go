// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRRVCCleanModeClusterModeTagStruct] class.
var (
	MTRRVCCleanModeClusterModeTagStructClass     _MTRRVCCleanModeClusterModeTagStructClass
	MTRRVCCleanModeClusterModeTagStructClassOnce sync.Once
)

func getMTRRVCCleanModeClusterModeTagStructClass() _MTRRVCCleanModeClusterModeTagStructClass {
	MTRRVCCleanModeClusterModeTagStructClassOnce.Do(func() {
		MTRRVCCleanModeClusterModeTagStructClass = _MTRRVCCleanModeClusterModeTagStructClass{objc.GetClass("MTRRVCCleanModeClusterModeTagStruct")}
	})
	return MTRRVCCleanModeClusterModeTagStructClass
}

type _MTRRVCCleanModeClusterModeTagStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRRVCCleanModeClusterModeTagStruct] class.
type IMTRRVCCleanModeClusterModeTagStruct interface {
	objectivec.IObject
	MfgCode() foundation.Number
	SetMfgCode(value foundation.INumber)
	Value() foundation.Number
	SetValue(value foundation.INumber)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRRVCCleanModeClusterModeTagStruct
type MTRRVCCleanModeClusterModeTagStruct struct {
	objectivec.Object
}

// MTRRVCCleanModeClusterModeTagStructFrom constructs a [MTRRVCCleanModeClusterModeTagStruct] from an unsafe.Pointer.
func MTRRVCCleanModeClusterModeTagStructFrom(ptr unsafe.Pointer) MTRRVCCleanModeClusterModeTagStruct {
	return MTRRVCCleanModeClusterModeTagStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRRVCCleanModeClusterModeTagStructClass) Alloc() MTRRVCCleanModeClusterModeTagStruct {
	rv := objc.Send[MTRRVCCleanModeClusterModeTagStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRRVCCleanModeClusterModeTagStructClass) New() MTRRVCCleanModeClusterModeTagStruct {
	rv := objc.Send[MTRRVCCleanModeClusterModeTagStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRRVCCleanModeClusterModeTagStruct) Init() MTRRVCCleanModeClusterModeTagStruct {
	rv := objc.Send[MTRRVCCleanModeClusterModeTagStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRRVCCleanModeClusterModeTagStruct) Autorelease() MTRRVCCleanModeClusterModeTagStruct {
	rv := objc.Send[MTRRVCCleanModeClusterModeTagStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRRVCCleanModeClusterModeTagStruct creates a new MTRRVCCleanModeClusterModeTagStruct instance.
func NewMTRRVCCleanModeClusterModeTagStruct() MTRRVCCleanModeClusterModeTagStruct {
	return getMTRRVCCleanModeClusterModeTagStructClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrrvccleanmodeclustermodetagstruct/mfgcode
func (m_ MTRRVCCleanModeClusterModeTagStruct) MfgCode() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("mfgCode"))
	return rv
}


// SetMfgCode sets the value of the mfgCode property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrrvccleanmodeclustermodetagstruct/mfgcode
func (m_ MTRRVCCleanModeClusterModeTagStruct) SetMfgCode(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMfgCode:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrrvccleanmodeclustermodetagstruct/value
func (m_ MTRRVCCleanModeClusterModeTagStruct) Value() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("value"))
	return rv
}


// SetValue sets the value of the value property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrrvccleanmodeclustermodetagstruct/value
func (m_ MTRRVCCleanModeClusterModeTagStruct) SetValue(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValue:"), value)
}



