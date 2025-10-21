// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRModeSelectClusterSemanticTagStruct] class.
var (
	MTRModeSelectClusterSemanticTagStructClass     _MTRModeSelectClusterSemanticTagStructClass
	MTRModeSelectClusterSemanticTagStructClassOnce sync.Once
)

func getMTRModeSelectClusterSemanticTagStructClass() _MTRModeSelectClusterSemanticTagStructClass {
	MTRModeSelectClusterSemanticTagStructClassOnce.Do(func() {
		MTRModeSelectClusterSemanticTagStructClass = _MTRModeSelectClusterSemanticTagStructClass{objc.GetClass("MTRModeSelectClusterSemanticTagStruct")}
	})
	return MTRModeSelectClusterSemanticTagStructClass
}

type _MTRModeSelectClusterSemanticTagStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRModeSelectClusterSemanticTagStruct] class.
type IMTRModeSelectClusterSemanticTagStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRModeSelectClusterSemanticTagStruct
type MTRModeSelectClusterSemanticTagStruct struct {
	objectivec.Object
}

// MTRModeSelectClusterSemanticTagStructFrom constructs a [MTRModeSelectClusterSemanticTagStruct] from an unsafe.Pointer.
func MTRModeSelectClusterSemanticTagStructFrom(ptr unsafe.Pointer) MTRModeSelectClusterSemanticTagStruct {
	return MTRModeSelectClusterSemanticTagStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRModeSelectClusterSemanticTagStructClass) Alloc() MTRModeSelectClusterSemanticTagStruct {
	rv := objc.Send[MTRModeSelectClusterSemanticTagStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRModeSelectClusterSemanticTagStructClass) New() MTRModeSelectClusterSemanticTagStruct {
	rv := objc.Send[MTRModeSelectClusterSemanticTagStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRModeSelectClusterSemanticTagStruct) Init() MTRModeSelectClusterSemanticTagStruct {
	rv := objc.Send[MTRModeSelectClusterSemanticTagStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRModeSelectClusterSemanticTagStruct) Autorelease() MTRModeSelectClusterSemanticTagStruct {
	rv := objc.Send[MTRModeSelectClusterSemanticTagStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRModeSelectClusterSemanticTagStruct creates a new MTRModeSelectClusterSemanticTagStruct instance.
func NewMTRModeSelectClusterSemanticTagStruct() MTRModeSelectClusterSemanticTagStruct {
	return getMTRModeSelectClusterSemanticTagStructClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmodeselectclustersemantictagstruct/mfgcode
func (m_ MTRModeSelectClusterSemanticTagStruct) MfgCode() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("mfgCode"))
	return rv
}


// SetMfgCode sets the value of the mfgCode property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmodeselectclustersemantictagstruct/mfgcode
func (m_ MTRModeSelectClusterSemanticTagStruct) SetMfgCode(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMfgCode:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmodeselectclustersemantictagstruct/value
func (m_ MTRModeSelectClusterSemanticTagStruct) Value() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("value"))
	return rv
}


// SetValue sets the value of the value property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmodeselectclustersemantictagstruct/value
func (m_ MTRModeSelectClusterSemanticTagStruct) SetValue(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValue:"), value)
}



