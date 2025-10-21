// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRLaundryWasherModeClusterModeTagStruct] class.
var (
	MTRLaundryWasherModeClusterModeTagStructClass     _MTRLaundryWasherModeClusterModeTagStructClass
	MTRLaundryWasherModeClusterModeTagStructClassOnce sync.Once
)

func getMTRLaundryWasherModeClusterModeTagStructClass() _MTRLaundryWasherModeClusterModeTagStructClass {
	MTRLaundryWasherModeClusterModeTagStructClassOnce.Do(func() {
		MTRLaundryWasherModeClusterModeTagStructClass = _MTRLaundryWasherModeClusterModeTagStructClass{objc.GetClass("MTRLaundryWasherModeClusterModeTagStruct")}
	})
	return MTRLaundryWasherModeClusterModeTagStructClass
}

type _MTRLaundryWasherModeClusterModeTagStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRLaundryWasherModeClusterModeTagStruct] class.
type IMTRLaundryWasherModeClusterModeTagStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLaundryWasherModeClusterModeTagStruct
type MTRLaundryWasherModeClusterModeTagStruct struct {
	objectivec.Object
}

// MTRLaundryWasherModeClusterModeTagStructFrom constructs a [MTRLaundryWasherModeClusterModeTagStruct] from an unsafe.Pointer.
func MTRLaundryWasherModeClusterModeTagStructFrom(ptr unsafe.Pointer) MTRLaundryWasherModeClusterModeTagStruct {
	return MTRLaundryWasherModeClusterModeTagStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRLaundryWasherModeClusterModeTagStructClass) Alloc() MTRLaundryWasherModeClusterModeTagStruct {
	rv := objc.Send[MTRLaundryWasherModeClusterModeTagStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRLaundryWasherModeClusterModeTagStructClass) New() MTRLaundryWasherModeClusterModeTagStruct {
	rv := objc.Send[MTRLaundryWasherModeClusterModeTagStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRLaundryWasherModeClusterModeTagStruct) Init() MTRLaundryWasherModeClusterModeTagStruct {
	rv := objc.Send[MTRLaundryWasherModeClusterModeTagStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRLaundryWasherModeClusterModeTagStruct) Autorelease() MTRLaundryWasherModeClusterModeTagStruct {
	rv := objc.Send[MTRLaundryWasherModeClusterModeTagStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRLaundryWasherModeClusterModeTagStruct creates a new MTRLaundryWasherModeClusterModeTagStruct instance.
func NewMTRLaundryWasherModeClusterModeTagStruct() MTRLaundryWasherModeClusterModeTagStruct {
	return getMTRLaundryWasherModeClusterModeTagStructClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLaundryWasherModeClusterModeTagStruct/mfgCode
func (m_ MTRLaundryWasherModeClusterModeTagStruct) MfgCode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("mfgCode"))
	return rv
}


// SetMfgCode sets the value of the mfgCode property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLaundryWasherModeClusterModeTagStruct/mfgCode
func (m_ MTRLaundryWasherModeClusterModeTagStruct) SetMfgCode(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMfgCode:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLaundryWasherModeClusterModeTagStruct/value
func (m_ MTRLaundryWasherModeClusterModeTagStruct) Value() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("value"))
	return rv
}


// SetValue sets the value of the value property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLaundryWasherModeClusterModeTagStruct/value
func (m_ MTRLaundryWasherModeClusterModeTagStruct) SetValue(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValue:"), value)
}


