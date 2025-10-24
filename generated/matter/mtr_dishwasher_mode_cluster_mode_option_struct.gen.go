// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	// properties:
	Label() objc.IObject /* cross-framework: NSString */
	SetLabel(value objc.IObject /* cross-framework: NSString */)
	Mode() objc.IObject /* cross-framework: NSNumber */
	SetMode(value objc.IObject /* cross-framework: NSNumber */)
	ModeTags() objc.IObject /* cross-framework: NSArray */
	SetModeTags(value objc.IObject /* cross-framework: NSArray */)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherModeClusterModeOptionStruct/label
func (m_ MTRDishwasherModeClusterModeOptionStruct) Label() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("label"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherModeClusterModeOptionStruct/label
func (m_ MTRDishwasherModeClusterModeOptionStruct) SetLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLabel:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherModeClusterModeOptionStruct/mode
func (m_ MTRDishwasherModeClusterModeOptionStruct) Mode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("mode"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherModeClusterModeOptionStruct/mode
func (m_ MTRDishwasherModeClusterModeOptionStruct) SetMode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMode:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherModeClusterModeOptionStruct/modeTags
func (m_ MTRDishwasherModeClusterModeOptionStruct) ModeTags() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("modeTags"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDishwasherModeClusterModeOptionStruct/modeTags
func (m_ MTRDishwasherModeClusterModeOptionStruct) SetModeTags(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setModeTags:"), value)
}



