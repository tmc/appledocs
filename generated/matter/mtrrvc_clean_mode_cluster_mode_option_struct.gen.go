// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRRVCCleanModeClusterModeOptionStruct] class.
var (
	MTRRVCCleanModeClusterModeOptionStructClass     _MTRRVCCleanModeClusterModeOptionStructClass
	MTRRVCCleanModeClusterModeOptionStructClassOnce sync.Once
)

func getMTRRVCCleanModeClusterModeOptionStructClass() _MTRRVCCleanModeClusterModeOptionStructClass {
	MTRRVCCleanModeClusterModeOptionStructClassOnce.Do(func() {
		MTRRVCCleanModeClusterModeOptionStructClass = _MTRRVCCleanModeClusterModeOptionStructClass{objc.GetClass("MTRRVCCleanModeClusterModeOptionStruct")}
	})
	return MTRRVCCleanModeClusterModeOptionStructClass
}

type _MTRRVCCleanModeClusterModeOptionStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRRVCCleanModeClusterModeOptionStruct] class.
type IMTRRVCCleanModeClusterModeOptionStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRRVCCleanModeClusterModeOptionStruct
type MTRRVCCleanModeClusterModeOptionStruct struct {
	objectivec.Object
}

// MTRRVCCleanModeClusterModeOptionStructFrom constructs a [MTRRVCCleanModeClusterModeOptionStruct] from an unsafe.Pointer.
func MTRRVCCleanModeClusterModeOptionStructFrom(ptr unsafe.Pointer) MTRRVCCleanModeClusterModeOptionStruct {
	return MTRRVCCleanModeClusterModeOptionStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRRVCCleanModeClusterModeOptionStructClass) Alloc() MTRRVCCleanModeClusterModeOptionStruct {
	rv := objc.Send[MTRRVCCleanModeClusterModeOptionStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRRVCCleanModeClusterModeOptionStructClass) New() MTRRVCCleanModeClusterModeOptionStruct {
	rv := objc.Send[MTRRVCCleanModeClusterModeOptionStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRRVCCleanModeClusterModeOptionStruct) Init() MTRRVCCleanModeClusterModeOptionStruct {
	rv := objc.Send[MTRRVCCleanModeClusterModeOptionStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRRVCCleanModeClusterModeOptionStruct) Autorelease() MTRRVCCleanModeClusterModeOptionStruct {
	rv := objc.Send[MTRRVCCleanModeClusterModeOptionStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRRVCCleanModeClusterModeOptionStruct creates a new MTRRVCCleanModeClusterModeOptionStruct instance.
func NewMTRRVCCleanModeClusterModeOptionStruct() MTRRVCCleanModeClusterModeOptionStruct {
	return getMTRRVCCleanModeClusterModeOptionStructClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrrvccleanmodeclustermodeoptionstruct/label
func (m_ MTRRVCCleanModeClusterModeOptionStruct) Label() string {
	rv := objc.Send[string](m_.ID, objc.Sel("label"))
	return rv
}


// SetLabel sets the value of the label property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrrvccleanmodeclustermodeoptionstruct/label
func (m_ MTRRVCCleanModeClusterModeOptionStruct) SetLabel(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLabel:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrrvccleanmodeclustermodeoptionstruct/mode
func (m_ MTRRVCCleanModeClusterModeOptionStruct) Mode() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("mode"))
	return rv
}


// SetMode sets the value of the mode property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrrvccleanmodeclustermodeoptionstruct/mode
func (m_ MTRRVCCleanModeClusterModeOptionStruct) SetMode(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMode:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrrvccleanmodeclustermodeoptionstruct/modetags
func (m_ MTRRVCCleanModeClusterModeOptionStruct) ModeTags() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("modeTags"))
	return rv
}


// SetModeTags sets the value of the modeTags property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrrvccleanmodeclustermodeoptionstruct/modetags
func (m_ MTRRVCCleanModeClusterModeOptionStruct) SetModeTags(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setModeTags:"), value)
}



