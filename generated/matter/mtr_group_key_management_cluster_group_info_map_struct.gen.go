// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRGroupKeyManagementClusterGroupInfoMapStruct] class.
var (
	MTRGroupKeyManagementClusterGroupInfoMapStructClass     _MTRGroupKeyManagementClusterGroupInfoMapStructClass
	MTRGroupKeyManagementClusterGroupInfoMapStructClassOnce sync.Once
)

func getMTRGroupKeyManagementClusterGroupInfoMapStructClass() _MTRGroupKeyManagementClusterGroupInfoMapStructClass {
	MTRGroupKeyManagementClusterGroupInfoMapStructClassOnce.Do(func() {
		MTRGroupKeyManagementClusterGroupInfoMapStructClass = _MTRGroupKeyManagementClusterGroupInfoMapStructClass{objc.GetClass("MTRGroupKeyManagementClusterGroupInfoMapStruct")}
	})
	return MTRGroupKeyManagementClusterGroupInfoMapStructClass
}

type _MTRGroupKeyManagementClusterGroupInfoMapStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRGroupKeyManagementClusterGroupInfoMapStruct] class.
type IMTRGroupKeyManagementClusterGroupInfoMapStruct interface {
	objectivec.IObject
	Endpoints() unsafe.Pointer
	SetEndpoints(value unsafe.Pointer)
	FabricIndex() foundation.Number
	SetFabricIndex(value foundation.INumber)
	GroupId() foundation.Number
	SetGroupId(value foundation.INumber)
	GroupName() string
	SetGroupName(value string)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGroupKeyManagementClusterGroupInfoMapStruct
type MTRGroupKeyManagementClusterGroupInfoMapStruct struct {
	objectivec.Object
}

// MTRGroupKeyManagementClusterGroupInfoMapStructFrom constructs a [MTRGroupKeyManagementClusterGroupInfoMapStruct] from an unsafe.Pointer.
func MTRGroupKeyManagementClusterGroupInfoMapStructFrom(ptr unsafe.Pointer) MTRGroupKeyManagementClusterGroupInfoMapStruct {
	return MTRGroupKeyManagementClusterGroupInfoMapStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRGroupKeyManagementClusterGroupInfoMapStructClass) Alloc() MTRGroupKeyManagementClusterGroupInfoMapStruct {
	rv := objc.Send[MTRGroupKeyManagementClusterGroupInfoMapStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRGroupKeyManagementClusterGroupInfoMapStructClass) New() MTRGroupKeyManagementClusterGroupInfoMapStruct {
	rv := objc.Send[MTRGroupKeyManagementClusterGroupInfoMapStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRGroupKeyManagementClusterGroupInfoMapStruct) Init() MTRGroupKeyManagementClusterGroupInfoMapStruct {
	rv := objc.Send[MTRGroupKeyManagementClusterGroupInfoMapStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRGroupKeyManagementClusterGroupInfoMapStruct) Autorelease() MTRGroupKeyManagementClusterGroupInfoMapStruct {
	rv := objc.Send[MTRGroupKeyManagementClusterGroupInfoMapStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRGroupKeyManagementClusterGroupInfoMapStruct creates a new MTRGroupKeyManagementClusterGroupInfoMapStruct instance.
func NewMTRGroupKeyManagementClusterGroupInfoMapStruct() MTRGroupKeyManagementClusterGroupInfoMapStruct {
	return getMTRGroupKeyManagementClusterGroupInfoMapStructClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupkeymanagementclustergroupinfomapstruct/endpoints
func (m_ MTRGroupKeyManagementClusterGroupInfoMapStruct) Endpoints() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("endpoints"))
	return rv
}


// SetEndpoints sets the value of the endpoints property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupkeymanagementclustergroupinfomapstruct/endpoints
func (m_ MTRGroupKeyManagementClusterGroupInfoMapStruct) SetEndpoints(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEndpoints:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupkeymanagementclustergroupinfomapstruct/fabricindex
func (m_ MTRGroupKeyManagementClusterGroupInfoMapStruct) FabricIndex() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("fabricIndex"))
	return rv
}


// SetFabricIndex sets the value of the fabricIndex property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupkeymanagementclustergroupinfomapstruct/fabricindex
func (m_ MTRGroupKeyManagementClusterGroupInfoMapStruct) SetFabricIndex(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricIndex:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupkeymanagementclustergroupinfomapstruct/groupid
func (m_ MTRGroupKeyManagementClusterGroupInfoMapStruct) GroupId() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("groupId"))
	return rv
}


// SetGroupId sets the value of the groupId property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupkeymanagementclustergroupinfomapstruct/groupid
func (m_ MTRGroupKeyManagementClusterGroupInfoMapStruct) SetGroupId(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGroupId:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupkeymanagementclustergroupinfomapstruct/groupname
func (m_ MTRGroupKeyManagementClusterGroupInfoMapStruct) GroupName() string {
	rv := objc.Send[string](m_.ID, objc.Sel("groupName"))
	return rv
}


// SetGroupName sets the value of the groupName property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupkeymanagementclustergroupinfomapstruct/groupname
func (m_ MTRGroupKeyManagementClusterGroupInfoMapStruct) SetGroupName(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGroupName:"), objc.String(value))
}



