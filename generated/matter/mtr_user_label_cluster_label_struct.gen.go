// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	// properties:
	Label() objc.IObject /* cross-framework: NSString */
	SetLabel(value objc.IObject /* cross-framework: NSString */)
	Value() objc.IObject /* cross-framework: NSString */
	SetValue(value objc.IObject /* cross-framework: NSString */)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtruserlabelclusterlabelstruct/label
func (m_ MTRUserLabelClusterLabelStruct) Label() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("label"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtruserlabelclusterlabelstruct/label
func (m_ MTRUserLabelClusterLabelStruct) SetLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLabel:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtruserlabelclusterlabelstruct/value
func (m_ MTRUserLabelClusterLabelStruct) Value() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("value"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtruserlabelclusterlabelstruct/value
func (m_ MTRUserLabelClusterLabelStruct) SetValue(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValue:"), value)
}



