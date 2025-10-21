// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRRVCRunModeClusterModeOptionStruct] class.
var (
	MTRRVCRunModeClusterModeOptionStructClass     _MTRRVCRunModeClusterModeOptionStructClass
	MTRRVCRunModeClusterModeOptionStructClassOnce sync.Once
)

func getMTRRVCRunModeClusterModeOptionStructClass() _MTRRVCRunModeClusterModeOptionStructClass {
	MTRRVCRunModeClusterModeOptionStructClassOnce.Do(func() {
		MTRRVCRunModeClusterModeOptionStructClass = _MTRRVCRunModeClusterModeOptionStructClass{objc.GetClass("MTRRVCRunModeClusterModeOptionStruct")}
	})
	return MTRRVCRunModeClusterModeOptionStructClass
}

type _MTRRVCRunModeClusterModeOptionStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRRVCRunModeClusterModeOptionStruct] class.
type IMTRRVCRunModeClusterModeOptionStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRRVCRunModeClusterModeOptionStruct
type MTRRVCRunModeClusterModeOptionStruct struct {
	objectivec.Object
}

// MTRRVCRunModeClusterModeOptionStructFrom constructs a [MTRRVCRunModeClusterModeOptionStruct] from an unsafe.Pointer.
func MTRRVCRunModeClusterModeOptionStructFrom(ptr unsafe.Pointer) MTRRVCRunModeClusterModeOptionStruct {
	return MTRRVCRunModeClusterModeOptionStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRRVCRunModeClusterModeOptionStructClass) Alloc() MTRRVCRunModeClusterModeOptionStruct {
	rv := objc.Send[MTRRVCRunModeClusterModeOptionStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRRVCRunModeClusterModeOptionStructClass) New() MTRRVCRunModeClusterModeOptionStruct {
	rv := objc.Send[MTRRVCRunModeClusterModeOptionStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRRVCRunModeClusterModeOptionStruct) Init() MTRRVCRunModeClusterModeOptionStruct {
	rv := objc.Send[MTRRVCRunModeClusterModeOptionStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRRVCRunModeClusterModeOptionStruct) Autorelease() MTRRVCRunModeClusterModeOptionStruct {
	rv := objc.Send[MTRRVCRunModeClusterModeOptionStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRRVCRunModeClusterModeOptionStruct creates a new MTRRVCRunModeClusterModeOptionStruct instance.
func NewMTRRVCRunModeClusterModeOptionStruct() MTRRVCRunModeClusterModeOptionStruct {
	return getMTRRVCRunModeClusterModeOptionStructClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrrvcrunmodeclustermodeoptionstruct/label
func (m_ MTRRVCRunModeClusterModeOptionStruct) Label() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("label"))
	return rv
}


// SetLabel sets the value of the label property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrrvcrunmodeclustermodeoptionstruct/label
func (m_ MTRRVCRunModeClusterModeOptionStruct) SetLabel(value appkit.string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLabel:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrrvcrunmodeclustermodeoptionstruct/mode
func (m_ MTRRVCRunModeClusterModeOptionStruct) Mode() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("mode"))
	return rv
}


// SetMode sets the value of the mode property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrrvcrunmodeclustermodeoptionstruct/mode
func (m_ MTRRVCRunModeClusterModeOptionStruct) SetMode(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMode:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrrvcrunmodeclustermodeoptionstruct/modetags
func (m_ MTRRVCRunModeClusterModeOptionStruct) ModeTags() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("modeTags"))
	return rv
}


// SetModeTags sets the value of the modeTags property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrrvcrunmodeclustermodeoptionstruct/modetags
func (m_ MTRRVCRunModeClusterModeOptionStruct) SetModeTags(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setModeTags:"), value)
}



