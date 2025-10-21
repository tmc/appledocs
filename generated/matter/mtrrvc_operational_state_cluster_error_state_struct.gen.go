// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRRVCOperationalStateClusterErrorStateStruct] class.
var (
	MTRRVCOperationalStateClusterErrorStateStructClass     _MTRRVCOperationalStateClusterErrorStateStructClass
	MTRRVCOperationalStateClusterErrorStateStructClassOnce sync.Once
)

func getMTRRVCOperationalStateClusterErrorStateStructClass() _MTRRVCOperationalStateClusterErrorStateStructClass {
	MTRRVCOperationalStateClusterErrorStateStructClassOnce.Do(func() {
		MTRRVCOperationalStateClusterErrorStateStructClass = _MTRRVCOperationalStateClusterErrorStateStructClass{objc.GetClass("MTRRVCOperationalStateClusterErrorStateStruct")}
	})
	return MTRRVCOperationalStateClusterErrorStateStructClass
}

type _MTRRVCOperationalStateClusterErrorStateStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRRVCOperationalStateClusterErrorStateStruct] class.
type IMTRRVCOperationalStateClusterErrorStateStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRRVCOperationalStateClusterErrorStateStruct
type MTRRVCOperationalStateClusterErrorStateStruct struct {
	objectivec.Object
}

// MTRRVCOperationalStateClusterErrorStateStructFrom constructs a [MTRRVCOperationalStateClusterErrorStateStruct] from an unsafe.Pointer.
func MTRRVCOperationalStateClusterErrorStateStructFrom(ptr unsafe.Pointer) MTRRVCOperationalStateClusterErrorStateStruct {
	return MTRRVCOperationalStateClusterErrorStateStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRRVCOperationalStateClusterErrorStateStructClass) Alloc() MTRRVCOperationalStateClusterErrorStateStruct {
	rv := objc.Send[MTRRVCOperationalStateClusterErrorStateStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRRVCOperationalStateClusterErrorStateStructClass) New() MTRRVCOperationalStateClusterErrorStateStruct {
	rv := objc.Send[MTRRVCOperationalStateClusterErrorStateStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRRVCOperationalStateClusterErrorStateStruct) Init() MTRRVCOperationalStateClusterErrorStateStruct {
	rv := objc.Send[MTRRVCOperationalStateClusterErrorStateStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRRVCOperationalStateClusterErrorStateStruct) Autorelease() MTRRVCOperationalStateClusterErrorStateStruct {
	rv := objc.Send[MTRRVCOperationalStateClusterErrorStateStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRRVCOperationalStateClusterErrorStateStruct creates a new MTRRVCOperationalStateClusterErrorStateStruct instance.
func NewMTRRVCOperationalStateClusterErrorStateStruct() MTRRVCOperationalStateClusterErrorStateStruct {
	return getMTRRVCOperationalStateClusterErrorStateStructClass().New()
}




