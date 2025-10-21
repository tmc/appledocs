// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRMicrowaveOvenModeClusterModeOptionStruct] class.
var (
	MTRMicrowaveOvenModeClusterModeOptionStructClass     _MTRMicrowaveOvenModeClusterModeOptionStructClass
	MTRMicrowaveOvenModeClusterModeOptionStructClassOnce sync.Once
)

func getMTRMicrowaveOvenModeClusterModeOptionStructClass() _MTRMicrowaveOvenModeClusterModeOptionStructClass {
	MTRMicrowaveOvenModeClusterModeOptionStructClassOnce.Do(func() {
		MTRMicrowaveOvenModeClusterModeOptionStructClass = _MTRMicrowaveOvenModeClusterModeOptionStructClass{objc.GetClass("MTRMicrowaveOvenModeClusterModeOptionStruct")}
	})
	return MTRMicrowaveOvenModeClusterModeOptionStructClass
}

type _MTRMicrowaveOvenModeClusterModeOptionStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRMicrowaveOvenModeClusterModeOptionStruct] class.
type IMTRMicrowaveOvenModeClusterModeOptionStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMicrowaveOvenModeClusterModeOptionStruct
type MTRMicrowaveOvenModeClusterModeOptionStruct struct {
	objectivec.Object
}

// MTRMicrowaveOvenModeClusterModeOptionStructFrom constructs a [MTRMicrowaveOvenModeClusterModeOptionStruct] from an unsafe.Pointer.
func MTRMicrowaveOvenModeClusterModeOptionStructFrom(ptr unsafe.Pointer) MTRMicrowaveOvenModeClusterModeOptionStruct {
	return MTRMicrowaveOvenModeClusterModeOptionStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRMicrowaveOvenModeClusterModeOptionStructClass) Alloc() MTRMicrowaveOvenModeClusterModeOptionStruct {
	rv := objc.Send[MTRMicrowaveOvenModeClusterModeOptionStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRMicrowaveOvenModeClusterModeOptionStructClass) New() MTRMicrowaveOvenModeClusterModeOptionStruct {
	rv := objc.Send[MTRMicrowaveOvenModeClusterModeOptionStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRMicrowaveOvenModeClusterModeOptionStruct) Init() MTRMicrowaveOvenModeClusterModeOptionStruct {
	rv := objc.Send[MTRMicrowaveOvenModeClusterModeOptionStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRMicrowaveOvenModeClusterModeOptionStruct) Autorelease() MTRMicrowaveOvenModeClusterModeOptionStruct {
	rv := objc.Send[MTRMicrowaveOvenModeClusterModeOptionStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRMicrowaveOvenModeClusterModeOptionStruct creates a new MTRMicrowaveOvenModeClusterModeOptionStruct instance.
func NewMTRMicrowaveOvenModeClusterModeOptionStruct() MTRMicrowaveOvenModeClusterModeOptionStruct {
	return getMTRMicrowaveOvenModeClusterModeOptionStructClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMicrowaveOvenModeClusterModeOptionStruct/label
func (m_ MTRMicrowaveOvenModeClusterModeOptionStruct) Label() string {
	rv := objc.Send[string](m_.ID, objc.Sel("label"))
	return rv
}


// SetLabel sets the value of the label property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMicrowaveOvenModeClusterModeOptionStruct/label
func (m_ MTRMicrowaveOvenModeClusterModeOptionStruct) SetLabel(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLabel:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMicrowaveOvenModeClusterModeOptionStruct/mode
func (m_ MTRMicrowaveOvenModeClusterModeOptionStruct) Mode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("mode"))
	return rv
}


// SetMode sets the value of the mode property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMicrowaveOvenModeClusterModeOptionStruct/mode
func (m_ MTRMicrowaveOvenModeClusterModeOptionStruct) SetMode(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMode:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMicrowaveOvenModeClusterModeOptionStruct/modeTags
func (m_ MTRMicrowaveOvenModeClusterModeOptionStruct) ModeTags() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("modeTags"))
	return rv
}


// SetModeTags sets the value of the modeTags property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMicrowaveOvenModeClusterModeOptionStruct/modeTags
func (m_ MTRMicrowaveOvenModeClusterModeOptionStruct) SetModeTags(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setModeTags:"), value)
}



