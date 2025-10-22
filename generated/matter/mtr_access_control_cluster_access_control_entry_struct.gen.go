// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRAccessControlClusterAccessControlEntryStruct] class.
var (
	MTRAccessControlClusterAccessControlEntryStructClass     _MTRAccessControlClusterAccessControlEntryStructClass
	MTRAccessControlClusterAccessControlEntryStructClassOnce sync.Once
)

func getMTRAccessControlClusterAccessControlEntryStructClass() _MTRAccessControlClusterAccessControlEntryStructClass {
	MTRAccessControlClusterAccessControlEntryStructClassOnce.Do(func() {
		MTRAccessControlClusterAccessControlEntryStructClass = _MTRAccessControlClusterAccessControlEntryStructClass{objc.GetClass("MTRAccessControlClusterAccessControlEntryStruct")}
	})
	return MTRAccessControlClusterAccessControlEntryStructClass
}

type _MTRAccessControlClusterAccessControlEntryStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRAccessControlClusterAccessControlEntryStruct] class.
type IMTRAccessControlClusterAccessControlEntryStruct interface {
	objectivec.IObject
	AuthMode() foundation.Number
	SetAuthMode(value foundation.INumber)
	FabricIndex() foundation.Number
	SetFabricIndex(value foundation.INumber)
	Privilege() foundation.Number
	SetPrivilege(value foundation.INumber)
	Subjects() unsafe.Pointer
	SetSubjects(value unsafe.Pointer)
	Targets() unsafe.Pointer
	SetTargets(value unsafe.Pointer)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccessControlClusterAccessControlEntryStruct
type MTRAccessControlClusterAccessControlEntryStruct struct {
	objectivec.Object
}

// MTRAccessControlClusterAccessControlEntryStructFrom constructs a [MTRAccessControlClusterAccessControlEntryStruct] from an unsafe.Pointer.
func MTRAccessControlClusterAccessControlEntryStructFrom(ptr unsafe.Pointer) MTRAccessControlClusterAccessControlEntryStruct {
	return MTRAccessControlClusterAccessControlEntryStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRAccessControlClusterAccessControlEntryStructClass) Alloc() MTRAccessControlClusterAccessControlEntryStruct {
	rv := objc.Send[MTRAccessControlClusterAccessControlEntryStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRAccessControlClusterAccessControlEntryStructClass) New() MTRAccessControlClusterAccessControlEntryStruct {
	rv := objc.Send[MTRAccessControlClusterAccessControlEntryStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRAccessControlClusterAccessControlEntryStruct) Init() MTRAccessControlClusterAccessControlEntryStruct {
	rv := objc.Send[MTRAccessControlClusterAccessControlEntryStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRAccessControlClusterAccessControlEntryStruct) Autorelease() MTRAccessControlClusterAccessControlEntryStruct {
	rv := objc.Send[MTRAccessControlClusterAccessControlEntryStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRAccessControlClusterAccessControlEntryStruct creates a new MTRAccessControlClusterAccessControlEntryStruct instance.
func NewMTRAccessControlClusterAccessControlEntryStruct() MTRAccessControlClusterAccessControlEntryStruct {
	return getMTRAccessControlClusterAccessControlEntryStructClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusteraccesscontrolentrystruct/authmode
func (m_ MTRAccessControlClusterAccessControlEntryStruct) AuthMode() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("authMode"))
	return rv
}


// SetAuthMode sets the value of the authMode property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusteraccesscontrolentrystruct/authmode
func (m_ MTRAccessControlClusterAccessControlEntryStruct) SetAuthMode(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAuthMode:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusteraccesscontrolentrystruct/fabricindex
func (m_ MTRAccessControlClusterAccessControlEntryStruct) FabricIndex() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("fabricIndex"))
	return rv
}


// SetFabricIndex sets the value of the fabricIndex property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusteraccesscontrolentrystruct/fabricindex
func (m_ MTRAccessControlClusterAccessControlEntryStruct) SetFabricIndex(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricIndex:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusteraccesscontrolentrystruct/privilege
func (m_ MTRAccessControlClusterAccessControlEntryStruct) Privilege() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("privilege"))
	return rv
}


// SetPrivilege sets the value of the privilege property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusteraccesscontrolentrystruct/privilege
func (m_ MTRAccessControlClusterAccessControlEntryStruct) SetPrivilege(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPrivilege:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusteraccesscontrolentrystruct/subjects
func (m_ MTRAccessControlClusterAccessControlEntryStruct) Subjects() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("subjects"))
	return rv
}


// SetSubjects sets the value of the subjects property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusteraccesscontrolentrystruct/subjects
func (m_ MTRAccessControlClusterAccessControlEntryStruct) SetSubjects(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSubjects:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusteraccesscontrolentrystruct/targets
func (m_ MTRAccessControlClusterAccessControlEntryStruct) Targets() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("targets"))
	return rv
}


// SetTargets sets the value of the targets property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtraccesscontrolclusteraccesscontrolentrystruct/targets
func (m_ MTRAccessControlClusterAccessControlEntryStruct) SetTargets(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTargets:"), value)
}



