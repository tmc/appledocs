// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRDishwasherModeClusterModeOptionStruct] class.
var (
	MTRDishwasherModeClusterModeOptionStructClass     _MTRDishwasherModeClusterModeOptionStructClass
	MTRDishwasherModeClusterModeOptionStructClassOnce sync.Once
)

func getMTRDishwasherModeClusterModeOptionStructClass() _MTRDishwasherModeClusterModeOptionStructClass {
	MTRDishwasherModeClusterModeOptionStructClassOnce.Do(func() {
		MTRDishwasherModeClusterModeOptionStructClass = _MTRDishwasherModeClusterModeOptionStructClass{objc.GetClass("MTRDishwasherModeClusterModeOptionStruct")}
	})
	return MTRDishwasherModeClusterModeOptionStructClass
}

type _MTRDishwasherModeClusterModeOptionStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRDishwasherModeClusterModeOptionStruct] class.
type IMTRDishwasherModeClusterModeOptionStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherModeClusterModeOptionStruct
type MTRDishwasherModeClusterModeOptionStruct struct {
	objectivec.Object
}

// MTRDishwasherModeClusterModeOptionStructFrom constructs a [MTRDishwasherModeClusterModeOptionStruct] from an unsafe.Pointer.
func MTRDishwasherModeClusterModeOptionStructFrom(ptr unsafe.Pointer) MTRDishwasherModeClusterModeOptionStruct {
	return MTRDishwasherModeClusterModeOptionStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDishwasherModeClusterModeOptionStructClass) Alloc() MTRDishwasherModeClusterModeOptionStruct {
	rv := objc.Send[MTRDishwasherModeClusterModeOptionStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRDishwasherModeClusterModeOptionStructClass) New() MTRDishwasherModeClusterModeOptionStruct {
	rv := objc.Send[MTRDishwasherModeClusterModeOptionStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDishwasherModeClusterModeOptionStruct) Init() MTRDishwasherModeClusterModeOptionStruct {
	rv := objc.Send[MTRDishwasherModeClusterModeOptionStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDishwasherModeClusterModeOptionStruct) Autorelease() MTRDishwasherModeClusterModeOptionStruct {
	rv := objc.Send[MTRDishwasherModeClusterModeOptionStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDishwasherModeClusterModeOptionStruct creates a new MTRDishwasherModeClusterModeOptionStruct instance.
func NewMTRDishwasherModeClusterModeOptionStruct() MTRDishwasherModeClusterModeOptionStruct {
	return getMTRDishwasherModeClusterModeOptionStructClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherModeClusterModeOptionStruct/label
func (m_ MTRDishwasherModeClusterModeOptionStruct) Label() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("label"))
	return rv
}


// SetLabel sets the value of the label property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherModeClusterModeOptionStruct/label
func (m_ MTRDishwasherModeClusterModeOptionStruct) SetLabel(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLabel:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherModeClusterModeOptionStruct/mode
func (m_ MTRDishwasherModeClusterModeOptionStruct) Mode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("mode"))
	return rv
}


// SetMode sets the value of the mode property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherModeClusterModeOptionStruct/mode
func (m_ MTRDishwasherModeClusterModeOptionStruct) SetMode(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMode:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherModeClusterModeOptionStruct/modeTags
func (m_ MTRDishwasherModeClusterModeOptionStruct) ModeTags() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("modeTags"))
	return rv
}


// SetModeTags sets the value of the modeTags property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherModeClusterModeOptionStruct/modeTags
func (m_ MTRDishwasherModeClusterModeOptionStruct) SetModeTags(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setModeTags:"), value)
}


