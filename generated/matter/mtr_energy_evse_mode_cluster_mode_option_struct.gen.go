// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTREnergyEVSEModeClusterModeOptionStruct] class.
var (
	MTREnergyEVSEModeClusterModeOptionStructClass     _MTREnergyEVSEModeClusterModeOptionStructClass
	MTREnergyEVSEModeClusterModeOptionStructClassOnce sync.Once
)

func getMTREnergyEVSEModeClusterModeOptionStructClass() _MTREnergyEVSEModeClusterModeOptionStructClass {
	MTREnergyEVSEModeClusterModeOptionStructClassOnce.Do(func() {
		MTREnergyEVSEModeClusterModeOptionStructClass = _MTREnergyEVSEModeClusterModeOptionStructClass{objc.GetClass("MTREnergyEVSEModeClusterModeOptionStruct")}
	})
	return MTREnergyEVSEModeClusterModeOptionStructClass
}

type _MTREnergyEVSEModeClusterModeOptionStructClass struct {
	class objc.Class
}

// An interface definition for the [MTREnergyEVSEModeClusterModeOptionStruct] class.
type IMTREnergyEVSEModeClusterModeOptionStruct interface {
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
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEModeClusterModeOptionStruct
type MTREnergyEVSEModeClusterModeOptionStruct struct {
	objectivec.Object
}

// MTREnergyEVSEModeClusterModeOptionStructFrom constructs a [MTREnergyEVSEModeClusterModeOptionStruct] from an unsafe.Pointer.
func MTREnergyEVSEModeClusterModeOptionStructFrom(ptr unsafe.Pointer) MTREnergyEVSEModeClusterModeOptionStruct {
	return MTREnergyEVSEModeClusterModeOptionStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTREnergyEVSEModeClusterModeOptionStructClass) Alloc() MTREnergyEVSEModeClusterModeOptionStruct {
	rv := objc.Send[MTREnergyEVSEModeClusterModeOptionStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTREnergyEVSEModeClusterModeOptionStructClass) New() MTREnergyEVSEModeClusterModeOptionStruct {
	rv := objc.Send[MTREnergyEVSEModeClusterModeOptionStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTREnergyEVSEModeClusterModeOptionStruct) Init() MTREnergyEVSEModeClusterModeOptionStruct {
	rv := objc.Send[MTREnergyEVSEModeClusterModeOptionStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTREnergyEVSEModeClusterModeOptionStruct) Autorelease() MTREnergyEVSEModeClusterModeOptionStruct {
	rv := objc.Send[MTREnergyEVSEModeClusterModeOptionStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTREnergyEVSEModeClusterModeOptionStruct creates a new MTREnergyEVSEModeClusterModeOptionStruct instance.
func NewMTREnergyEVSEModeClusterModeOptionStruct() MTREnergyEVSEModeClusterModeOptionStruct {
	return getMTREnergyEVSEModeClusterModeOptionStructClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEModeClusterModeOptionStruct/label
func (m_ MTREnergyEVSEModeClusterModeOptionStruct) Label() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("label"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEModeClusterModeOptionStruct/label
func (m_ MTREnergyEVSEModeClusterModeOptionStruct) SetLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLabel:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEModeClusterModeOptionStruct/mode
func (m_ MTREnergyEVSEModeClusterModeOptionStruct) Mode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("mode"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEModeClusterModeOptionStruct/mode
func (m_ MTREnergyEVSEModeClusterModeOptionStruct) SetMode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMode:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEModeClusterModeOptionStruct/modeTags
func (m_ MTREnergyEVSEModeClusterModeOptionStruct) ModeTags() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("modeTags"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTREnergyEVSEModeClusterModeOptionStruct/modeTags
func (m_ MTREnergyEVSEModeClusterModeOptionStruct) SetModeTags(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setModeTags:"), value)
}



