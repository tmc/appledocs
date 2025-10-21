// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRAccessControlClusterAccessControlEntry] class.
var (
	MTRAccessControlClusterAccessControlEntryClass     _MTRAccessControlClusterAccessControlEntryClass
	MTRAccessControlClusterAccessControlEntryClassOnce sync.Once
)

func getMTRAccessControlClusterAccessControlEntryClass() _MTRAccessControlClusterAccessControlEntryClass {
	MTRAccessControlClusterAccessControlEntryClassOnce.Do(func() {
		MTRAccessControlClusterAccessControlEntryClass = _MTRAccessControlClusterAccessControlEntryClass{objc.GetClass("MTRAccessControlClusterAccessControlEntry")}
	})
	return MTRAccessControlClusterAccessControlEntryClass
}

type _MTRAccessControlClusterAccessControlEntryClass struct {
	class objc.Class
}

// An interface definition for the [MTRAccessControlClusterAccessControlEntry] class.
type IMTRAccessControlClusterAccessControlEntry interface {
	IMTRAccessControlClusterAccessControlEntryStruct
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccessControlClusterAccessControlEntry
type MTRAccessControlClusterAccessControlEntry struct {
	MTRAccessControlClusterAccessControlEntryStruct
}

// MTRAccessControlClusterAccessControlEntryFrom constructs a [MTRAccessControlClusterAccessControlEntry] from an unsafe.Pointer.
func MTRAccessControlClusterAccessControlEntryFrom(ptr unsafe.Pointer) MTRAccessControlClusterAccessControlEntry {
	return MTRAccessControlClusterAccessControlEntry{
		MTRAccessControlClusterAccessControlEntryStruct: MTRAccessControlClusterAccessControlEntryStructFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRAccessControlClusterAccessControlEntryClass) Alloc() MTRAccessControlClusterAccessControlEntry {
	rv := objc.Send[MTRAccessControlClusterAccessControlEntry](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRAccessControlClusterAccessControlEntryClass) New() MTRAccessControlClusterAccessControlEntry {
	rv := objc.Send[MTRAccessControlClusterAccessControlEntry](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRAccessControlClusterAccessControlEntry) Init() MTRAccessControlClusterAccessControlEntry {
	rv := objc.Send[MTRAccessControlClusterAccessControlEntry](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRAccessControlClusterAccessControlEntry) Autorelease() MTRAccessControlClusterAccessControlEntry {
	rv := objc.Send[MTRAccessControlClusterAccessControlEntry](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRAccessControlClusterAccessControlEntry creates a new MTRAccessControlClusterAccessControlEntry instance.
func NewMTRAccessControlClusterAccessControlEntry() MTRAccessControlClusterAccessControlEntry {
	return getMTRAccessControlClusterAccessControlEntryClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusteraccesscontrolentry/authmode
func (m_ MTRAccessControlClusterAccessControlEntry) AuthMode() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("authMode"))
	return rv
}


// SetAuthMode sets the value of the authMode property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusteraccesscontrolentry/authmode
func (m_ MTRAccessControlClusterAccessControlEntry) SetAuthMode(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAuthMode:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusteraccesscontrolentry/fabricindex
func (m_ MTRAccessControlClusterAccessControlEntry) FabricIndex() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("fabricIndex"))
	return rv
}


// SetFabricIndex sets the value of the fabricIndex property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusteraccesscontrolentry/fabricindex
func (m_ MTRAccessControlClusterAccessControlEntry) SetFabricIndex(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricIndex:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusteraccesscontrolentry/privilege
func (m_ MTRAccessControlClusterAccessControlEntry) Privilege() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("privilege"))
	return rv
}


// SetPrivilege sets the value of the privilege property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusteraccesscontrolentry/privilege
func (m_ MTRAccessControlClusterAccessControlEntry) SetPrivilege(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPrivilege:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusteraccesscontrolentry/subjects
func (m_ MTRAccessControlClusterAccessControlEntry) Subjects() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("subjects"))
	return rv
}


// SetSubjects sets the value of the subjects property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusteraccesscontrolentry/subjects
func (m_ MTRAccessControlClusterAccessControlEntry) SetSubjects(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSubjects:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusteraccesscontrolentry/targets
func (m_ MTRAccessControlClusterAccessControlEntry) Targets() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("targets"))
	return rv
}


// SetTargets sets the value of the targets property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusteraccesscontrolentry/targets
func (m_ MTRAccessControlClusterAccessControlEntry) SetTargets(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTargets:"), value)
}



