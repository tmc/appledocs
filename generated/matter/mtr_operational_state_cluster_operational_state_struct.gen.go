// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTROperationalStateClusterOperationalStateStruct] class.
var (
	MTROperationalStateClusterOperationalStateStructClass     _MTROperationalStateClusterOperationalStateStructClass
	MTROperationalStateClusterOperationalStateStructClassOnce sync.Once
)

func getMTROperationalStateClusterOperationalStateStructClass() _MTROperationalStateClusterOperationalStateStructClass {
	MTROperationalStateClusterOperationalStateStructClassOnce.Do(func() {
		MTROperationalStateClusterOperationalStateStructClass = _MTROperationalStateClusterOperationalStateStructClass{objc.GetClass("MTROperationalStateClusterOperationalStateStruct")}
	})
	return MTROperationalStateClusterOperationalStateStructClass
}

type _MTROperationalStateClusterOperationalStateStructClass struct {
	class objc.Class
}

// An interface definition for the [MTROperationalStateClusterOperationalStateStruct] class.
type IMTROperationalStateClusterOperationalStateStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROperationalStateClusterOperationalStateStruct
type MTROperationalStateClusterOperationalStateStruct struct {
	objectivec.Object
}

// MTROperationalStateClusterOperationalStateStructFrom constructs a [MTROperationalStateClusterOperationalStateStruct] from an unsafe.Pointer.
func MTROperationalStateClusterOperationalStateStructFrom(ptr unsafe.Pointer) MTROperationalStateClusterOperationalStateStruct {
	return MTROperationalStateClusterOperationalStateStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTROperationalStateClusterOperationalStateStructClass) Alloc() MTROperationalStateClusterOperationalStateStruct {
	rv := objc.Send[MTROperationalStateClusterOperationalStateStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTROperationalStateClusterOperationalStateStructClass) New() MTROperationalStateClusterOperationalStateStruct {
	rv := objc.Send[MTROperationalStateClusterOperationalStateStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROperationalStateClusterOperationalStateStruct) Init() MTROperationalStateClusterOperationalStateStruct {
	rv := objc.Send[MTROperationalStateClusterOperationalStateStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROperationalStateClusterOperationalStateStruct) Autorelease() MTROperationalStateClusterOperationalStateStruct {
	rv := objc.Send[MTROperationalStateClusterOperationalStateStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROperationalStateClusterOperationalStateStruct creates a new MTROperationalStateClusterOperationalStateStruct instance.
func NewMTROperationalStateClusterOperationalStateStruct() MTROperationalStateClusterOperationalStateStruct {
	return getMTROperationalStateClusterOperationalStateStructClass().New()
}




