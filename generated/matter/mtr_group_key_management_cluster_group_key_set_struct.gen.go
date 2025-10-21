// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRGroupKeyManagementClusterGroupKeySetStruct] class.
var (
	MTRGroupKeyManagementClusterGroupKeySetStructClass     _MTRGroupKeyManagementClusterGroupKeySetStructClass
	MTRGroupKeyManagementClusterGroupKeySetStructClassOnce sync.Once
)

func getMTRGroupKeyManagementClusterGroupKeySetStructClass() _MTRGroupKeyManagementClusterGroupKeySetStructClass {
	MTRGroupKeyManagementClusterGroupKeySetStructClassOnce.Do(func() {
		MTRGroupKeyManagementClusterGroupKeySetStructClass = _MTRGroupKeyManagementClusterGroupKeySetStructClass{objc.GetClass("MTRGroupKeyManagementClusterGroupKeySetStruct")}
	})
	return MTRGroupKeyManagementClusterGroupKeySetStructClass
}

type _MTRGroupKeyManagementClusterGroupKeySetStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRGroupKeyManagementClusterGroupKeySetStruct] class.
type IMTRGroupKeyManagementClusterGroupKeySetStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGroupKeyManagementClusterGroupKeySetStruct
type MTRGroupKeyManagementClusterGroupKeySetStruct struct {
	objectivec.Object
}

// MTRGroupKeyManagementClusterGroupKeySetStructFrom constructs a [MTRGroupKeyManagementClusterGroupKeySetStruct] from an unsafe.Pointer.
func MTRGroupKeyManagementClusterGroupKeySetStructFrom(ptr unsafe.Pointer) MTRGroupKeyManagementClusterGroupKeySetStruct {
	return MTRGroupKeyManagementClusterGroupKeySetStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRGroupKeyManagementClusterGroupKeySetStructClass) Alloc() MTRGroupKeyManagementClusterGroupKeySetStruct {
	rv := objc.Send[MTRGroupKeyManagementClusterGroupKeySetStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRGroupKeyManagementClusterGroupKeySetStructClass) New() MTRGroupKeyManagementClusterGroupKeySetStruct {
	rv := objc.Send[MTRGroupKeyManagementClusterGroupKeySetStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRGroupKeyManagementClusterGroupKeySetStruct) Init() MTRGroupKeyManagementClusterGroupKeySetStruct {
	rv := objc.Send[MTRGroupKeyManagementClusterGroupKeySetStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRGroupKeyManagementClusterGroupKeySetStruct) Autorelease() MTRGroupKeyManagementClusterGroupKeySetStruct {
	rv := objc.Send[MTRGroupKeyManagementClusterGroupKeySetStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRGroupKeyManagementClusterGroupKeySetStruct creates a new MTRGroupKeyManagementClusterGroupKeySetStruct instance.
func NewMTRGroupKeyManagementClusterGroupKeySetStruct() MTRGroupKeyManagementClusterGroupKeySetStruct {
	return getMTRGroupKeyManagementClusterGroupKeySetStructClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupkeymanagementclustergroupkeysetstruct/epochstarttime0
func (m_ MTRGroupKeyManagementClusterGroupKeySetStruct) EpochStartTime0() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("epochStartTime0"))
	return rv
}


// SetEpochStartTime0 sets the value of the epochStartTime0 property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupkeymanagementclustergroupkeysetstruct/epochstarttime0
func (m_ MTRGroupKeyManagementClusterGroupKeySetStruct) SetEpochStartTime0(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEpochStartTime0:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupkeymanagementclustergroupkeysetstruct/groupkeysecuritypolicy
func (m_ MTRGroupKeyManagementClusterGroupKeySetStruct) GroupKeySecurityPolicy() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("groupKeySecurityPolicy"))
	return rv
}


// SetGroupKeySecurityPolicy sets the value of the groupKeySecurityPolicy property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupkeymanagementclustergroupkeysetstruct/groupkeysecuritypolicy
func (m_ MTRGroupKeyManagementClusterGroupKeySetStruct) SetGroupKeySecurityPolicy(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGroupKeySecurityPolicy:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupkeymanagementclustergroupkeysetstruct/epochkey0
func (m_ MTRGroupKeyManagementClusterGroupKeySetStruct) EpochKey0() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("epochKey0"))
	return rv
}


// SetEpochKey0 sets the value of the epochKey0 property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupkeymanagementclustergroupkeysetstruct/epochkey0
func (m_ MTRGroupKeyManagementClusterGroupKeySetStruct) SetEpochKey0(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEpochKey0:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupkeymanagementclustergroupkeysetstruct/epochkey1
func (m_ MTRGroupKeyManagementClusterGroupKeySetStruct) EpochKey1() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("epochKey1"))
	return rv
}


// SetEpochKey1 sets the value of the epochKey1 property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupkeymanagementclustergroupkeysetstruct/epochkey1
func (m_ MTRGroupKeyManagementClusterGroupKeySetStruct) SetEpochKey1(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEpochKey1:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupkeymanagementclustergroupkeysetstruct/epochstarttime1
func (m_ MTRGroupKeyManagementClusterGroupKeySetStruct) EpochStartTime1() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("epochStartTime1"))
	return rv
}


// SetEpochStartTime1 sets the value of the epochStartTime1 property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupkeymanagementclustergroupkeysetstruct/epochstarttime1
func (m_ MTRGroupKeyManagementClusterGroupKeySetStruct) SetEpochStartTime1(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEpochStartTime1:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupkeymanagementclustergroupkeysetstruct/epochkey2
func (m_ MTRGroupKeyManagementClusterGroupKeySetStruct) EpochKey2() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("epochKey2"))
	return rv
}


// SetEpochKey2 sets the value of the epochKey2 property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupkeymanagementclustergroupkeysetstruct/epochkey2
func (m_ MTRGroupKeyManagementClusterGroupKeySetStruct) SetEpochKey2(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEpochKey2:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupkeymanagementclustergroupkeysetstruct/groupkeysetid
func (m_ MTRGroupKeyManagementClusterGroupKeySetStruct) GroupKeySetID() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("groupKeySetID"))
	return rv
}


// SetGroupKeySetID sets the value of the groupKeySetID property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupkeymanagementclustergroupkeysetstruct/groupkeysetid
func (m_ MTRGroupKeyManagementClusterGroupKeySetStruct) SetGroupKeySetID(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGroupKeySetID:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupkeymanagementclustergroupkeysetstruct/epochstarttime2
func (m_ MTRGroupKeyManagementClusterGroupKeySetStruct) EpochStartTime2() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("epochStartTime2"))
	return rv
}


// SetEpochStartTime2 sets the value of the epochStartTime2 property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupkeymanagementclustergroupkeysetstruct/epochstarttime2
func (m_ MTRGroupKeyManagementClusterGroupKeySetStruct) SetEpochStartTime2(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEpochStartTime2:"), value)
}



