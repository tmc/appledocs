// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRAccessControlClusterAccessControlExtensionStruct] class.
var (
	MTRAccessControlClusterAccessControlExtensionStructClass     _MTRAccessControlClusterAccessControlExtensionStructClass
	MTRAccessControlClusterAccessControlExtensionStructClassOnce sync.Once
)

func getMTRAccessControlClusterAccessControlExtensionStructClass() _MTRAccessControlClusterAccessControlExtensionStructClass {
	MTRAccessControlClusterAccessControlExtensionStructClassOnce.Do(func() {
		MTRAccessControlClusterAccessControlExtensionStructClass = _MTRAccessControlClusterAccessControlExtensionStructClass{objc.GetClass("MTRAccessControlClusterAccessControlExtensionStruct")}
	})
	return MTRAccessControlClusterAccessControlExtensionStructClass
}

type _MTRAccessControlClusterAccessControlExtensionStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRAccessControlClusterAccessControlExtensionStruct] class.
type IMTRAccessControlClusterAccessControlExtensionStruct interface {
	objectivec.IObject
	// properties:
	Data() objc.IObject /* cross-framework: Data */
	SetData(value objc.IObject /* cross-framework: Data */)
	FabricIndex() objc.IObject /* cross-framework: NSNumber */
	SetFabricIndex(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccessControlClusterAccessControlExtensionStruct
type MTRAccessControlClusterAccessControlExtensionStruct struct {
	objectivec.Object
}

// MTRAccessControlClusterAccessControlExtensionStructFrom constructs a [MTRAccessControlClusterAccessControlExtensionStruct] from an unsafe.Pointer.
func MTRAccessControlClusterAccessControlExtensionStructFrom(ptr unsafe.Pointer) MTRAccessControlClusterAccessControlExtensionStruct {
	return MTRAccessControlClusterAccessControlExtensionStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRAccessControlClusterAccessControlExtensionStructClass) Alloc() MTRAccessControlClusterAccessControlExtensionStruct {
	rv := objc.Send[MTRAccessControlClusterAccessControlExtensionStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRAccessControlClusterAccessControlExtensionStructClass) New() MTRAccessControlClusterAccessControlExtensionStruct {
	rv := objc.Send[MTRAccessControlClusterAccessControlExtensionStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRAccessControlClusterAccessControlExtensionStruct) Init() MTRAccessControlClusterAccessControlExtensionStruct {
	rv := objc.Send[MTRAccessControlClusterAccessControlExtensionStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRAccessControlClusterAccessControlExtensionStruct) Autorelease() MTRAccessControlClusterAccessControlExtensionStruct {
	rv := objc.Send[MTRAccessControlClusterAccessControlExtensionStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRAccessControlClusterAccessControlExtensionStruct creates a new MTRAccessControlClusterAccessControlExtensionStruct instance.
func NewMTRAccessControlClusterAccessControlExtensionStruct() MTRAccessControlClusterAccessControlExtensionStruct {
	return getMTRAccessControlClusterAccessControlExtensionStructClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusteraccesscontrolextensionstruct/data
func (m_ MTRAccessControlClusterAccessControlExtensionStruct) Data() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("data"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusteraccesscontrolextensionstruct/data
func (m_ MTRAccessControlClusterAccessControlExtensionStruct) SetData(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setData:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusteraccesscontrolextensionstruct/fabricindex
func (m_ MTRAccessControlClusterAccessControlExtensionStruct) FabricIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("fabricIndex"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusteraccesscontrolextensionstruct/fabricindex
func (m_ MTRAccessControlClusterAccessControlExtensionStruct) SetFabricIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricIndex:"), value)
}



