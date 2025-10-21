// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRUserLabelClusterLabelStruct] class.
var (
	MTRUserLabelClusterLabelStructClass     _MTRUserLabelClusterLabelStructClass
	MTRUserLabelClusterLabelStructClassOnce sync.Once
)

func getMTRUserLabelClusterLabelStructClass() _MTRUserLabelClusterLabelStructClass {
	MTRUserLabelClusterLabelStructClassOnce.Do(func() {
		MTRUserLabelClusterLabelStructClass = _MTRUserLabelClusterLabelStructClass{objc.GetClass("MTRUserLabelClusterLabelStruct")}
	})
	return MTRUserLabelClusterLabelStructClass
}

type _MTRUserLabelClusterLabelStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRUserLabelClusterLabelStruct] class.
type IMTRUserLabelClusterLabelStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUserLabelClusterLabelStruct
type MTRUserLabelClusterLabelStruct struct {
	objectivec.Object
}

// MTRUserLabelClusterLabelStructFrom constructs a [MTRUserLabelClusterLabelStruct] from an unsafe.Pointer.
func MTRUserLabelClusterLabelStructFrom(ptr unsafe.Pointer) MTRUserLabelClusterLabelStruct {
	return MTRUserLabelClusterLabelStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRUserLabelClusterLabelStructClass) Alloc() MTRUserLabelClusterLabelStruct {
	rv := objc.Send[MTRUserLabelClusterLabelStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRUserLabelClusterLabelStructClass) New() MTRUserLabelClusterLabelStruct {
	rv := objc.Send[MTRUserLabelClusterLabelStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUserLabelClusterLabelStruct) Init() MTRUserLabelClusterLabelStruct {
	rv := objc.Send[MTRUserLabelClusterLabelStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUserLabelClusterLabelStruct) Autorelease() MTRUserLabelClusterLabelStruct {
	rv := objc.Send[MTRUserLabelClusterLabelStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUserLabelClusterLabelStruct creates a new MTRUserLabelClusterLabelStruct instance.
func NewMTRUserLabelClusterLabelStruct() MTRUserLabelClusterLabelStruct {
	return getMTRUserLabelClusterLabelStructClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtruserlabelclusterlabelstruct/label
func (m_ MTRUserLabelClusterLabelStruct) Label() string {
	rv := objc.Send[string](m_.ID, objc.Sel("label"))
	return rv
}


// SetLabel sets the value of the label property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtruserlabelclusterlabelstruct/label
func (m_ MTRUserLabelClusterLabelStruct) SetLabel(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLabel:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtruserlabelclusterlabelstruct/value
func (m_ MTRUserLabelClusterLabelStruct) Value() string {
	rv := objc.Send[string](m_.ID, objc.Sel("value"))
	return rv
}


// SetValue sets the value of the value property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtruserlabelclusterlabelstruct/value
func (m_ MTRUserLabelClusterLabelStruct) SetValue(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValue:"), objc.String(value))
}



