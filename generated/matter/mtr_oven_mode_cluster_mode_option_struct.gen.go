// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTROvenModeClusterModeOptionStruct] class.
var (
	MTROvenModeClusterModeOptionStructClass     _MTROvenModeClusterModeOptionStructClass
	MTROvenModeClusterModeOptionStructClassOnce sync.Once
)

func getMTROvenModeClusterModeOptionStructClass() _MTROvenModeClusterModeOptionStructClass {
	MTROvenModeClusterModeOptionStructClassOnce.Do(func() {
		MTROvenModeClusterModeOptionStructClass = _MTROvenModeClusterModeOptionStructClass{objc.GetClass("MTROvenModeClusterModeOptionStruct")}
	})
	return MTROvenModeClusterModeOptionStructClass
}

type _MTROvenModeClusterModeOptionStructClass struct {
	class objc.Class
}

// An interface definition for the [MTROvenModeClusterModeOptionStruct] class.
type IMTROvenModeClusterModeOptionStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenModeClusterModeOptionStruct
type MTROvenModeClusterModeOptionStruct struct {
	objectivec.Object
}

// MTROvenModeClusterModeOptionStructFrom constructs a [MTROvenModeClusterModeOptionStruct] from an unsafe.Pointer.
func MTROvenModeClusterModeOptionStructFrom(ptr unsafe.Pointer) MTROvenModeClusterModeOptionStruct {
	return MTROvenModeClusterModeOptionStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROvenModeClusterModeOptionStructClass) Alloc() MTROvenModeClusterModeOptionStruct {
	rv := objc.Send[MTROvenModeClusterModeOptionStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROvenModeClusterModeOptionStructClass) New() MTROvenModeClusterModeOptionStruct {
	rv := objc.Send[MTROvenModeClusterModeOptionStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROvenModeClusterModeOptionStruct) Init() MTROvenModeClusterModeOptionStruct {
	rv := objc.Send[MTROvenModeClusterModeOptionStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROvenModeClusterModeOptionStruct) Autorelease() MTROvenModeClusterModeOptionStruct {
	rv := objc.Send[MTROvenModeClusterModeOptionStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROvenModeClusterModeOptionStruct creates a new MTROvenModeClusterModeOptionStruct instance.
func NewMTROvenModeClusterModeOptionStruct() MTROvenModeClusterModeOptionStruct {
	return getMTROvenModeClusterModeOptionStructClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenModeClusterModeOptionStruct/label
func (m_ MTROvenModeClusterModeOptionStruct) Label() string {
	rv := objc.Send[string](m_.ID, objc.Sel("label"))
	return rv
}


// SetLabel sets the value of the label property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenModeClusterModeOptionStruct/label
func (m_ MTROvenModeClusterModeOptionStruct) SetLabel(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLabel:"), objc.String(value))
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenModeClusterModeOptionStruct/mode
func (m_ MTROvenModeClusterModeOptionStruct) Mode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("mode"))
	return rv
}


// SetMode sets the value of the mode property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenModeClusterModeOptionStruct/mode
func (m_ MTROvenModeClusterModeOptionStruct) SetMode(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMode:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenModeClusterModeOptionStruct/modeTags
func (m_ MTROvenModeClusterModeOptionStruct) ModeTags() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("modeTags"))
	return rv
}


// SetModeTags sets the value of the modeTags property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenModeClusterModeOptionStruct/modeTags
func (m_ MTROvenModeClusterModeOptionStruct) SetModeTags(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setModeTags:"), value)
}


