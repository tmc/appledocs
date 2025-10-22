// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRAccessControlClusterExtensionEntry] class.
var (
	MTRAccessControlClusterExtensionEntryClass     _MTRAccessControlClusterExtensionEntryClass
	MTRAccessControlClusterExtensionEntryClassOnce sync.Once
)

func getMTRAccessControlClusterExtensionEntryClass() _MTRAccessControlClusterExtensionEntryClass {
	MTRAccessControlClusterExtensionEntryClassOnce.Do(func() {
		MTRAccessControlClusterExtensionEntryClass = _MTRAccessControlClusterExtensionEntryClass{objc.GetClass("MTRAccessControlClusterExtensionEntry")}
	})
	return MTRAccessControlClusterExtensionEntryClass
}

type _MTRAccessControlClusterExtensionEntryClass struct {
	class objc.Class
}

// An interface definition for the [MTRAccessControlClusterExtensionEntry] class.
type IMTRAccessControlClusterExtensionEntry interface {
	IMTRAccessControlClusterAccessControlExtensionStruct
	Data() foundation.Data
	SetData(value foundation.IData)
	FabricIndex() foundation.Number
	SetFabricIndex(value foundation.INumber)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccessControlClusterExtensionEntry
type MTRAccessControlClusterExtensionEntry struct {
	MTRAccessControlClusterAccessControlExtensionStruct
}

// MTRAccessControlClusterExtensionEntryFrom constructs a [MTRAccessControlClusterExtensionEntry] from an unsafe.Pointer.
func MTRAccessControlClusterExtensionEntryFrom(ptr unsafe.Pointer) MTRAccessControlClusterExtensionEntry {
	return MTRAccessControlClusterExtensionEntry{
		MTRAccessControlClusterAccessControlExtensionStruct: MTRAccessControlClusterAccessControlExtensionStructFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRAccessControlClusterExtensionEntryClass) Alloc() MTRAccessControlClusterExtensionEntry {
	rv := objc.Send[MTRAccessControlClusterExtensionEntry](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRAccessControlClusterExtensionEntryClass) New() MTRAccessControlClusterExtensionEntry {
	rv := objc.Send[MTRAccessControlClusterExtensionEntry](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRAccessControlClusterExtensionEntry) Init() MTRAccessControlClusterExtensionEntry {
	rv := objc.Send[MTRAccessControlClusterExtensionEntry](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRAccessControlClusterExtensionEntry) Autorelease() MTRAccessControlClusterExtensionEntry {
	rv := objc.Send[MTRAccessControlClusterExtensionEntry](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRAccessControlClusterExtensionEntry creates a new MTRAccessControlClusterExtensionEntry instance.
func NewMTRAccessControlClusterExtensionEntry() MTRAccessControlClusterExtensionEntry {
	return getMTRAccessControlClusterExtensionEntryClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusterextensionentry/data
func (m_ MTRAccessControlClusterExtensionEntry) Data() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("data"))
	return rv
}


// SetData sets the value of the data property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusterextensionentry/data
func (m_ MTRAccessControlClusterExtensionEntry) SetData(value foundation.IData) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setData:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusterextensionentry/fabricindex
func (m_ MTRAccessControlClusterExtensionEntry) FabricIndex() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("fabricIndex"))
	return rv
}


// SetFabricIndex sets the value of the fabricIndex property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusterextensionentry/fabricindex
func (m_ MTRAccessControlClusterExtensionEntry) SetFabricIndex(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricIndex:"), value)
}



