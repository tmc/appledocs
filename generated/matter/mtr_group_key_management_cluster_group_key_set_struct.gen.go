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
	// properties:
	EpochKey0() objc.IObject /* cross-framework: Data */
	SetEpochKey0(value objc.IObject /* cross-framework: Data */)
	EpochKey1() objc.IObject /* cross-framework: Data */
	SetEpochKey1(value objc.IObject /* cross-framework: Data */)
	EpochKey2() objc.IObject /* cross-framework: Data */
	SetEpochKey2(value objc.IObject /* cross-framework: Data */)
	EpochStartTime0() objc.IObject /* cross-framework: NSNumber */
	SetEpochStartTime0(value objc.IObject /* cross-framework: NSNumber */)
	EpochStartTime1() objc.IObject /* cross-framework: NSNumber */
	SetEpochStartTime1(value objc.IObject /* cross-framework: NSNumber */)
	EpochStartTime2() objc.IObject /* cross-framework: NSNumber */
	SetEpochStartTime2(value objc.IObject /* cross-framework: NSNumber */)
	GroupKeySecurityPolicy() objc.IObject /* cross-framework: NSNumber */
	SetGroupKeySecurityPolicy(value objc.IObject /* cross-framework: NSNumber */)
	GroupKeySetID() objc.IObject /* cross-framework: NSNumber */
	SetGroupKeySetID(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupkeymanagementclustergroupkeysetstruct/epochkey0
func (m_ MTRGroupKeyManagementClusterGroupKeySetStruct) EpochKey0() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("epochKey0"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupkeymanagementclustergroupkeysetstruct/epochkey0
func (m_ MTRGroupKeyManagementClusterGroupKeySetStruct) SetEpochKey0(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEpochKey0:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupkeymanagementclustergroupkeysetstruct/epochkey1
func (m_ MTRGroupKeyManagementClusterGroupKeySetStruct) EpochKey1() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("epochKey1"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupkeymanagementclustergroupkeysetstruct/epochkey1
func (m_ MTRGroupKeyManagementClusterGroupKeySetStruct) SetEpochKey1(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEpochKey1:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupkeymanagementclustergroupkeysetstruct/epochkey2
func (m_ MTRGroupKeyManagementClusterGroupKeySetStruct) EpochKey2() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("epochKey2"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupkeymanagementclustergroupkeysetstruct/epochkey2
func (m_ MTRGroupKeyManagementClusterGroupKeySetStruct) SetEpochKey2(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEpochKey2:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupkeymanagementclustergroupkeysetstruct/epochstarttime0
func (m_ MTRGroupKeyManagementClusterGroupKeySetStruct) EpochStartTime0() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("epochStartTime0"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupkeymanagementclustergroupkeysetstruct/epochstarttime0
func (m_ MTRGroupKeyManagementClusterGroupKeySetStruct) SetEpochStartTime0(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEpochStartTime0:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupkeymanagementclustergroupkeysetstruct/epochstarttime1
func (m_ MTRGroupKeyManagementClusterGroupKeySetStruct) EpochStartTime1() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("epochStartTime1"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupkeymanagementclustergroupkeysetstruct/epochstarttime1
func (m_ MTRGroupKeyManagementClusterGroupKeySetStruct) SetEpochStartTime1(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEpochStartTime1:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupkeymanagementclustergroupkeysetstruct/epochstarttime2
func (m_ MTRGroupKeyManagementClusterGroupKeySetStruct) EpochStartTime2() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("epochStartTime2"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupkeymanagementclustergroupkeysetstruct/epochstarttime2
func (m_ MTRGroupKeyManagementClusterGroupKeySetStruct) SetEpochStartTime2(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEpochStartTime2:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupkeymanagementclustergroupkeysetstruct/groupkeysecuritypolicy
func (m_ MTRGroupKeyManagementClusterGroupKeySetStruct) GroupKeySecurityPolicy() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("groupKeySecurityPolicy"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupkeymanagementclustergroupkeysetstruct/groupkeysecuritypolicy
func (m_ MTRGroupKeyManagementClusterGroupKeySetStruct) SetGroupKeySecurityPolicy(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGroupKeySecurityPolicy:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupkeymanagementclustergroupkeysetstruct/groupkeysetid
func (m_ MTRGroupKeyManagementClusterGroupKeySetStruct) GroupKeySetID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("groupKeySetID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrgroupkeymanagementclustergroupkeysetstruct/groupkeysetid
func (m_ MTRGroupKeyManagementClusterGroupKeySetStruct) SetGroupKeySetID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGroupKeySetID:"), value)
}



