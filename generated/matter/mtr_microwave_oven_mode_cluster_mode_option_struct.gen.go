// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMicrowaveOvenModeClusterModeOptionStruct/label
func (m_ MTRMicrowaveOvenModeClusterModeOptionStruct) Label() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("label"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMicrowaveOvenModeClusterModeOptionStruct/label
func (m_ MTRMicrowaveOvenModeClusterModeOptionStruct) SetLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLabel:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMicrowaveOvenModeClusterModeOptionStruct/mode
func (m_ MTRMicrowaveOvenModeClusterModeOptionStruct) Mode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("mode"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMicrowaveOvenModeClusterModeOptionStruct/mode
func (m_ MTRMicrowaveOvenModeClusterModeOptionStruct) SetMode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMode:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMicrowaveOvenModeClusterModeOptionStruct/modeTags
func (m_ MTRMicrowaveOvenModeClusterModeOptionStruct) ModeTags() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("modeTags"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMicrowaveOvenModeClusterModeOptionStruct/modeTags
func (m_ MTRMicrowaveOvenModeClusterModeOptionStruct) SetModeTags(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setModeTags:"), value)
}



