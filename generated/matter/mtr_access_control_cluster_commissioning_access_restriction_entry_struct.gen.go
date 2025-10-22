// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRAccessControlClusterCommissioningAccessRestrictionEntryStruct] class.
var (
	MTRAccessControlClusterCommissioningAccessRestrictionEntryStructClass     _MTRAccessControlClusterCommissioningAccessRestrictionEntryStructClass
	MTRAccessControlClusterCommissioningAccessRestrictionEntryStructClassOnce sync.Once
)

func getMTRAccessControlClusterCommissioningAccessRestrictionEntryStructClass() _MTRAccessControlClusterCommissioningAccessRestrictionEntryStructClass {
	MTRAccessControlClusterCommissioningAccessRestrictionEntryStructClassOnce.Do(func() {
		MTRAccessControlClusterCommissioningAccessRestrictionEntryStructClass = _MTRAccessControlClusterCommissioningAccessRestrictionEntryStructClass{objc.GetClass("MTRAccessControlClusterCommissioningAccessRestrictionEntryStruct")}
	})
	return MTRAccessControlClusterCommissioningAccessRestrictionEntryStructClass
}

type _MTRAccessControlClusterCommissioningAccessRestrictionEntryStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRAccessControlClusterCommissioningAccessRestrictionEntryStruct] class.
type IMTRAccessControlClusterCommissioningAccessRestrictionEntryStruct interface {
	objectivec.IObject
	Cluster() foundation.Number
	SetCluster(value foundation.INumber)
	Endpoint() foundation.Number
	SetEndpoint(value foundation.INumber)
	Restrictions() objc.ID
	SetRestrictions(value objc.ID)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccessControlClusterCommissioningAccessRestrictionEntryStruct
type MTRAccessControlClusterCommissioningAccessRestrictionEntryStruct struct {
	objectivec.Object
}

// MTRAccessControlClusterCommissioningAccessRestrictionEntryStructFrom constructs a [MTRAccessControlClusterCommissioningAccessRestrictionEntryStruct] from an unsafe.Pointer.
func MTRAccessControlClusterCommissioningAccessRestrictionEntryStructFrom(ptr unsafe.Pointer) MTRAccessControlClusterCommissioningAccessRestrictionEntryStruct {
	return MTRAccessControlClusterCommissioningAccessRestrictionEntryStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRAccessControlClusterCommissioningAccessRestrictionEntryStructClass) Alloc() MTRAccessControlClusterCommissioningAccessRestrictionEntryStruct {
	rv := objc.Send[MTRAccessControlClusterCommissioningAccessRestrictionEntryStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRAccessControlClusterCommissioningAccessRestrictionEntryStructClass) New() MTRAccessControlClusterCommissioningAccessRestrictionEntryStruct {
	rv := objc.Send[MTRAccessControlClusterCommissioningAccessRestrictionEntryStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRAccessControlClusterCommissioningAccessRestrictionEntryStruct) Init() MTRAccessControlClusterCommissioningAccessRestrictionEntryStruct {
	rv := objc.Send[MTRAccessControlClusterCommissioningAccessRestrictionEntryStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRAccessControlClusterCommissioningAccessRestrictionEntryStruct) Autorelease() MTRAccessControlClusterCommissioningAccessRestrictionEntryStruct {
	rv := objc.Send[MTRAccessControlClusterCommissioningAccessRestrictionEntryStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRAccessControlClusterCommissioningAccessRestrictionEntryStruct creates a new MTRAccessControlClusterCommissioningAccessRestrictionEntryStruct instance.
func NewMTRAccessControlClusterCommissioningAccessRestrictionEntryStruct() MTRAccessControlClusterCommissioningAccessRestrictionEntryStruct {
	return getMTRAccessControlClusterCommissioningAccessRestrictionEntryStructClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccessControlClusterCommissioningAccessRestrictionEntryStruct/cluster
func (m_ MTRAccessControlClusterCommissioningAccessRestrictionEntryStruct) Cluster() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("cluster"))
	return rv
}


// SetCluster sets the value of the cluster property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccessControlClusterCommissioningAccessRestrictionEntryStruct/cluster
func (m_ MTRAccessControlClusterCommissioningAccessRestrictionEntryStruct) SetCluster(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCluster:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccessControlClusterCommissioningAccessRestrictionEntryStruct/endpoint
func (m_ MTRAccessControlClusterCommissioningAccessRestrictionEntryStruct) Endpoint() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("endpoint"))
	return rv
}


// SetEndpoint sets the value of the endpoint property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccessControlClusterCommissioningAccessRestrictionEntryStruct/endpoint
func (m_ MTRAccessControlClusterCommissioningAccessRestrictionEntryStruct) SetEndpoint(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEndpoint:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccessControlClusterCommissioningAccessRestrictionEntryStruct/restrictions
func (m_ MTRAccessControlClusterCommissioningAccessRestrictionEntryStruct) Restrictions() objc.ID {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("restrictions"))
	return rv
}


// SetRestrictions sets the value of the restrictions property.
//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAccessControlClusterCommissioningAccessRestrictionEntryStruct/restrictions
func (m_ MTRAccessControlClusterCommissioningAccessRestrictionEntryStruct) SetRestrictions(value objc.ID) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRestrictions:"), value)
}



