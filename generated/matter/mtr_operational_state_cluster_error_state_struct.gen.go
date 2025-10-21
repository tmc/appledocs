// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTROperationalStateClusterErrorStateStruct] class.
var (
	MTROperationalStateClusterErrorStateStructClass     _MTROperationalStateClusterErrorStateStructClass
	MTROperationalStateClusterErrorStateStructClassOnce sync.Once
)

func getMTROperationalStateClusterErrorStateStructClass() _MTROperationalStateClusterErrorStateStructClass {
	MTROperationalStateClusterErrorStateStructClassOnce.Do(func() {
		MTROperationalStateClusterErrorStateStructClass = _MTROperationalStateClusterErrorStateStructClass{objc.GetClass("MTROperationalStateClusterErrorStateStruct")}
	})
	return MTROperationalStateClusterErrorStateStructClass
}

type _MTROperationalStateClusterErrorStateStructClass struct {
	class objc.Class
}

// An interface definition for the [MTROperationalStateClusterErrorStateStruct] class.
type IMTROperationalStateClusterErrorStateStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalStateClusterErrorStateStruct
type MTROperationalStateClusterErrorStateStruct struct {
	objectivec.Object
}

// MTROperationalStateClusterErrorStateStructFrom constructs a [MTROperationalStateClusterErrorStateStruct] from an unsafe.Pointer.
func MTROperationalStateClusterErrorStateStructFrom(ptr unsafe.Pointer) MTROperationalStateClusterErrorStateStruct {
	return MTROperationalStateClusterErrorStateStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROperationalStateClusterErrorStateStructClass) Alloc() MTROperationalStateClusterErrorStateStruct {
	rv := objc.Send[MTROperationalStateClusterErrorStateStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROperationalStateClusterErrorStateStructClass) New() MTROperationalStateClusterErrorStateStruct {
	rv := objc.Send[MTROperationalStateClusterErrorStateStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROperationalStateClusterErrorStateStruct) Init() MTROperationalStateClusterErrorStateStruct {
	rv := objc.Send[MTROperationalStateClusterErrorStateStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROperationalStateClusterErrorStateStruct) Autorelease() MTROperationalStateClusterErrorStateStruct {
	rv := objc.Send[MTROperationalStateClusterErrorStateStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROperationalStateClusterErrorStateStruct creates a new MTROperationalStateClusterErrorStateStruct instance.
func NewMTROperationalStateClusterErrorStateStruct() MTROperationalStateClusterErrorStateStruct {
	return getMTROperationalStateClusterErrorStateStructClass().New()
}




