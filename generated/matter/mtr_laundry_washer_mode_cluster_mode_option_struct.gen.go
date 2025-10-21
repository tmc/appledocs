// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRLaundryWasherModeClusterModeOptionStruct] class.
var (
	MTRLaundryWasherModeClusterModeOptionStructClass     _MTRLaundryWasherModeClusterModeOptionStructClass
	MTRLaundryWasherModeClusterModeOptionStructClassOnce sync.Once
)

func getMTRLaundryWasherModeClusterModeOptionStructClass() _MTRLaundryWasherModeClusterModeOptionStructClass {
	MTRLaundryWasherModeClusterModeOptionStructClassOnce.Do(func() {
		MTRLaundryWasherModeClusterModeOptionStructClass = _MTRLaundryWasherModeClusterModeOptionStructClass{objc.GetClass("MTRLaundryWasherModeClusterModeOptionStruct")}
	})
	return MTRLaundryWasherModeClusterModeOptionStructClass
}

type _MTRLaundryWasherModeClusterModeOptionStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRLaundryWasherModeClusterModeOptionStruct] class.
type IMTRLaundryWasherModeClusterModeOptionStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLaundryWasherModeClusterModeOptionStruct
type MTRLaundryWasherModeClusterModeOptionStruct struct {
	objectivec.Object
}

// MTRLaundryWasherModeClusterModeOptionStructFrom constructs a [MTRLaundryWasherModeClusterModeOptionStruct] from an unsafe.Pointer.
func MTRLaundryWasherModeClusterModeOptionStructFrom(ptr unsafe.Pointer) MTRLaundryWasherModeClusterModeOptionStruct {
	return MTRLaundryWasherModeClusterModeOptionStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRLaundryWasherModeClusterModeOptionStructClass) Alloc() MTRLaundryWasherModeClusterModeOptionStruct {
	rv := objc.Send[MTRLaundryWasherModeClusterModeOptionStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRLaundryWasherModeClusterModeOptionStructClass) New() MTRLaundryWasherModeClusterModeOptionStruct {
	rv := objc.Send[MTRLaundryWasherModeClusterModeOptionStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRLaundryWasherModeClusterModeOptionStruct) Init() MTRLaundryWasherModeClusterModeOptionStruct {
	rv := objc.Send[MTRLaundryWasherModeClusterModeOptionStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRLaundryWasherModeClusterModeOptionStruct) Autorelease() MTRLaundryWasherModeClusterModeOptionStruct {
	rv := objc.Send[MTRLaundryWasherModeClusterModeOptionStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRLaundryWasherModeClusterModeOptionStruct creates a new MTRLaundryWasherModeClusterModeOptionStruct instance.
func NewMTRLaundryWasherModeClusterModeOptionStruct() MTRLaundryWasherModeClusterModeOptionStruct {
	return getMTRLaundryWasherModeClusterModeOptionStructClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLaundryWasherModeClusterModeOptionStruct/label
func (m_ MTRLaundryWasherModeClusterModeOptionStruct) Label() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("label"))
	return rv
}


// SetLabel sets the value of the label property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLaundryWasherModeClusterModeOptionStruct/label
func (m_ MTRLaundryWasherModeClusterModeOptionStruct) SetLabel(value appkit.string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLabel:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLaundryWasherModeClusterModeOptionStruct/mode
func (m_ MTRLaundryWasherModeClusterModeOptionStruct) Mode() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("mode"))
	return rv
}


// SetMode sets the value of the mode property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLaundryWasherModeClusterModeOptionStruct/mode
func (m_ MTRLaundryWasherModeClusterModeOptionStruct) SetMode(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMode:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLaundryWasherModeClusterModeOptionStruct/modeTags
func (m_ MTRLaundryWasherModeClusterModeOptionStruct) ModeTags() objc.ID {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("modeTags"))
	return rv
}


// SetModeTags sets the value of the modeTags property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRLaundryWasherModeClusterModeOptionStruct/modeTags
func (m_ MTRLaundryWasherModeClusterModeOptionStruct) SetModeTags(value objc.ID) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setModeTags:"), value)
}



