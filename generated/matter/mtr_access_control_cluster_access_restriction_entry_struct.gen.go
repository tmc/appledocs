// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRAccessControlClusterAccessRestrictionEntryStruct] class.
var (
	MTRAccessControlClusterAccessRestrictionEntryStructClass     _MTRAccessControlClusterAccessRestrictionEntryStructClass
	MTRAccessControlClusterAccessRestrictionEntryStructClassOnce sync.Once
)

func getMTRAccessControlClusterAccessRestrictionEntryStructClass() _MTRAccessControlClusterAccessRestrictionEntryStructClass {
	MTRAccessControlClusterAccessRestrictionEntryStructClassOnce.Do(func() {
		MTRAccessControlClusterAccessRestrictionEntryStructClass = _MTRAccessControlClusterAccessRestrictionEntryStructClass{objc.GetClass("MTRAccessControlClusterAccessRestrictionEntryStruct")}
	})
	return MTRAccessControlClusterAccessRestrictionEntryStructClass
}

type _MTRAccessControlClusterAccessRestrictionEntryStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRAccessControlClusterAccessRestrictionEntryStruct] class.
type IMTRAccessControlClusterAccessRestrictionEntryStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccessControlClusterAccessRestrictionEntryStruct
type MTRAccessControlClusterAccessRestrictionEntryStruct struct {
	objectivec.Object
}

// MTRAccessControlClusterAccessRestrictionEntryStructFrom constructs a [MTRAccessControlClusterAccessRestrictionEntryStruct] from an unsafe.Pointer.
func MTRAccessControlClusterAccessRestrictionEntryStructFrom(ptr unsafe.Pointer) MTRAccessControlClusterAccessRestrictionEntryStruct {
	return MTRAccessControlClusterAccessRestrictionEntryStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRAccessControlClusterAccessRestrictionEntryStructClass) Alloc() MTRAccessControlClusterAccessRestrictionEntryStruct {
	rv := objc.Send[MTRAccessControlClusterAccessRestrictionEntryStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRAccessControlClusterAccessRestrictionEntryStructClass) New() MTRAccessControlClusterAccessRestrictionEntryStruct {
	rv := objc.Send[MTRAccessControlClusterAccessRestrictionEntryStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRAccessControlClusterAccessRestrictionEntryStruct) Init() MTRAccessControlClusterAccessRestrictionEntryStruct {
	rv := objc.Send[MTRAccessControlClusterAccessRestrictionEntryStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRAccessControlClusterAccessRestrictionEntryStruct) Autorelease() MTRAccessControlClusterAccessRestrictionEntryStruct {
	rv := objc.Send[MTRAccessControlClusterAccessRestrictionEntryStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRAccessControlClusterAccessRestrictionEntryStruct creates a new MTRAccessControlClusterAccessRestrictionEntryStruct instance.
func NewMTRAccessControlClusterAccessRestrictionEntryStruct() MTRAccessControlClusterAccessRestrictionEntryStruct {
	return getMTRAccessControlClusterAccessRestrictionEntryStructClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccessControlClusterAccessRestrictionEntryStruct/cluster
func (m_ MTRAccessControlClusterAccessRestrictionEntryStruct) Cluster() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("cluster"))
	return rv
}


// SetCluster sets the value of the cluster property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccessControlClusterAccessRestrictionEntryStruct/cluster
func (m_ MTRAccessControlClusterAccessRestrictionEntryStruct) SetCluster(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCluster:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccessControlClusterAccessRestrictionEntryStruct/endpoint
func (m_ MTRAccessControlClusterAccessRestrictionEntryStruct) Endpoint() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("endpoint"))
	return rv
}


// SetEndpoint sets the value of the endpoint property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccessControlClusterAccessRestrictionEntryStruct/endpoint
func (m_ MTRAccessControlClusterAccessRestrictionEntryStruct) SetEndpoint(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEndpoint:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccessControlClusterAccessRestrictionEntryStruct/fabricIndex
func (m_ MTRAccessControlClusterAccessRestrictionEntryStruct) FabricIndex() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("fabricIndex"))
	return rv
}


// SetFabricIndex sets the value of the fabricIndex property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccessControlClusterAccessRestrictionEntryStruct/fabricIndex
func (m_ MTRAccessControlClusterAccessRestrictionEntryStruct) SetFabricIndex(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricIndex:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccessControlClusterAccessRestrictionEntryStruct/restrictions
func (m_ MTRAccessControlClusterAccessRestrictionEntryStruct) Restrictions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("restrictions"))
	return rv
}


// SetRestrictions sets the value of the restrictions property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccessControlClusterAccessRestrictionEntryStruct/restrictions
func (m_ MTRAccessControlClusterAccessRestrictionEntryStruct) SetRestrictions(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRestrictions:"), value)
}



