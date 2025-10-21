// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
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




