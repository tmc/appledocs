// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRNetworkCommissioningClusterNetworkInfoStruct] class.
var (
	MTRNetworkCommissioningClusterNetworkInfoStructClass     _MTRNetworkCommissioningClusterNetworkInfoStructClass
	MTRNetworkCommissioningClusterNetworkInfoStructClassOnce sync.Once
)

func getMTRNetworkCommissioningClusterNetworkInfoStructClass() _MTRNetworkCommissioningClusterNetworkInfoStructClass {
	MTRNetworkCommissioningClusterNetworkInfoStructClassOnce.Do(func() {
		MTRNetworkCommissioningClusterNetworkInfoStructClass = _MTRNetworkCommissioningClusterNetworkInfoStructClass{objc.GetClass("MTRNetworkCommissioningClusterNetworkInfoStruct")}
	})
	return MTRNetworkCommissioningClusterNetworkInfoStructClass
}

type _MTRNetworkCommissioningClusterNetworkInfoStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRNetworkCommissioningClusterNetworkInfoStruct] class.
type IMTRNetworkCommissioningClusterNetworkInfoStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRNetworkCommissioningClusterNetworkInfoStruct
type MTRNetworkCommissioningClusterNetworkInfoStruct struct {
	objectivec.Object
}

// MTRNetworkCommissioningClusterNetworkInfoStructFrom constructs a [MTRNetworkCommissioningClusterNetworkInfoStruct] from an unsafe.Pointer.
func MTRNetworkCommissioningClusterNetworkInfoStructFrom(ptr unsafe.Pointer) MTRNetworkCommissioningClusterNetworkInfoStruct {
	return MTRNetworkCommissioningClusterNetworkInfoStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRNetworkCommissioningClusterNetworkInfoStructClass) Alloc() MTRNetworkCommissioningClusterNetworkInfoStruct {
	rv := objc.Send[MTRNetworkCommissioningClusterNetworkInfoStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRNetworkCommissioningClusterNetworkInfoStructClass) New() MTRNetworkCommissioningClusterNetworkInfoStruct {
	rv := objc.Send[MTRNetworkCommissioningClusterNetworkInfoStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRNetworkCommissioningClusterNetworkInfoStruct) Init() MTRNetworkCommissioningClusterNetworkInfoStruct {
	rv := objc.Send[MTRNetworkCommissioningClusterNetworkInfoStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRNetworkCommissioningClusterNetworkInfoStruct) Autorelease() MTRNetworkCommissioningClusterNetworkInfoStruct {
	rv := objc.Send[MTRNetworkCommissioningClusterNetworkInfoStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRNetworkCommissioningClusterNetworkInfoStruct creates a new MTRNetworkCommissioningClusterNetworkInfoStruct instance.
func NewMTRNetworkCommissioningClusterNetworkInfoStruct() MTRNetworkCommissioningClusterNetworkInfoStruct {
	return getMTRNetworkCommissioningClusterNetworkInfoStructClass().New()
}




