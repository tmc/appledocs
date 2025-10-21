// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRWaterHeaterModeClusterModeOptionStruct] class.
var (
	MTRWaterHeaterModeClusterModeOptionStructClass     _MTRWaterHeaterModeClusterModeOptionStructClass
	MTRWaterHeaterModeClusterModeOptionStructClassOnce sync.Once
)

func getMTRWaterHeaterModeClusterModeOptionStructClass() _MTRWaterHeaterModeClusterModeOptionStructClass {
	MTRWaterHeaterModeClusterModeOptionStructClassOnce.Do(func() {
		MTRWaterHeaterModeClusterModeOptionStructClass = _MTRWaterHeaterModeClusterModeOptionStructClass{objc.GetClass("MTRWaterHeaterModeClusterModeOptionStruct")}
	})
	return MTRWaterHeaterModeClusterModeOptionStructClass
}

type _MTRWaterHeaterModeClusterModeOptionStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRWaterHeaterModeClusterModeOptionStruct] class.
type IMTRWaterHeaterModeClusterModeOptionStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterModeClusterModeOptionStruct
type MTRWaterHeaterModeClusterModeOptionStruct struct {
	objectivec.Object
}

// MTRWaterHeaterModeClusterModeOptionStructFrom constructs a [MTRWaterHeaterModeClusterModeOptionStruct] from an unsafe.Pointer.
func MTRWaterHeaterModeClusterModeOptionStructFrom(ptr unsafe.Pointer) MTRWaterHeaterModeClusterModeOptionStruct {
	return MTRWaterHeaterModeClusterModeOptionStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRWaterHeaterModeClusterModeOptionStructClass) Alloc() MTRWaterHeaterModeClusterModeOptionStruct {
	rv := objc.Send[MTRWaterHeaterModeClusterModeOptionStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRWaterHeaterModeClusterModeOptionStructClass) New() MTRWaterHeaterModeClusterModeOptionStruct {
	rv := objc.Send[MTRWaterHeaterModeClusterModeOptionStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRWaterHeaterModeClusterModeOptionStruct) Init() MTRWaterHeaterModeClusterModeOptionStruct {
	rv := objc.Send[MTRWaterHeaterModeClusterModeOptionStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRWaterHeaterModeClusterModeOptionStruct) Autorelease() MTRWaterHeaterModeClusterModeOptionStruct {
	rv := objc.Send[MTRWaterHeaterModeClusterModeOptionStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRWaterHeaterModeClusterModeOptionStruct creates a new MTRWaterHeaterModeClusterModeOptionStruct instance.
func NewMTRWaterHeaterModeClusterModeOptionStruct() MTRWaterHeaterModeClusterModeOptionStruct {
	return getMTRWaterHeaterModeClusterModeOptionStructClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterModeClusterModeOptionStruct/label
func (m_ MTRWaterHeaterModeClusterModeOptionStruct) Label() string {
	rv := objc.Send[string](m_.ID, objc.Sel("label"))
	return rv
}


// SetLabel sets the value of the label property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterModeClusterModeOptionStruct/label
func (m_ MTRWaterHeaterModeClusterModeOptionStruct) SetLabel(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLabel:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterModeClusterModeOptionStruct/mode
func (m_ MTRWaterHeaterModeClusterModeOptionStruct) Mode() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("mode"))
	return rv
}


// SetMode sets the value of the mode property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterModeClusterModeOptionStruct/mode
func (m_ MTRWaterHeaterModeClusterModeOptionStruct) SetMode(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMode:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterModeClusterModeOptionStruct/modeTags
func (m_ MTRWaterHeaterModeClusterModeOptionStruct) ModeTags() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("modeTags"))
	return rv
}


// SetModeTags sets the value of the modeTags property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWaterHeaterModeClusterModeOptionStruct/modeTags
func (m_ MTRWaterHeaterModeClusterModeOptionStruct) SetModeTags(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setModeTags:"), value)
}



