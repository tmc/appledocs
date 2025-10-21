// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRApplicationBasicClusterApplicationStruct] class.
var (
	MTRApplicationBasicClusterApplicationStructClass     _MTRApplicationBasicClusterApplicationStructClass
	MTRApplicationBasicClusterApplicationStructClassOnce sync.Once
)

func getMTRApplicationBasicClusterApplicationStructClass() _MTRApplicationBasicClusterApplicationStructClass {
	MTRApplicationBasicClusterApplicationStructClassOnce.Do(func() {
		MTRApplicationBasicClusterApplicationStructClass = _MTRApplicationBasicClusterApplicationStructClass{objc.GetClass("MTRApplicationBasicClusterApplicationStruct")}
	})
	return MTRApplicationBasicClusterApplicationStructClass
}

type _MTRApplicationBasicClusterApplicationStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRApplicationBasicClusterApplicationStruct] class.
type IMTRApplicationBasicClusterApplicationStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRApplicationBasicClusterApplicationStruct
type MTRApplicationBasicClusterApplicationStruct struct {
	objectivec.Object
}

// MTRApplicationBasicClusterApplicationStructFrom constructs a [MTRApplicationBasicClusterApplicationStruct] from an unsafe.Pointer.
func MTRApplicationBasicClusterApplicationStructFrom(ptr unsafe.Pointer) MTRApplicationBasicClusterApplicationStruct {
	return MTRApplicationBasicClusterApplicationStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRApplicationBasicClusterApplicationStructClass) Alloc() MTRApplicationBasicClusterApplicationStruct {
	rv := objc.Send[MTRApplicationBasicClusterApplicationStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRApplicationBasicClusterApplicationStructClass) New() MTRApplicationBasicClusterApplicationStruct {
	rv := objc.Send[MTRApplicationBasicClusterApplicationStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRApplicationBasicClusterApplicationStruct) Init() MTRApplicationBasicClusterApplicationStruct {
	rv := objc.Send[MTRApplicationBasicClusterApplicationStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRApplicationBasicClusterApplicationStruct) Autorelease() MTRApplicationBasicClusterApplicationStruct {
	rv := objc.Send[MTRApplicationBasicClusterApplicationStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRApplicationBasicClusterApplicationStruct creates a new MTRApplicationBasicClusterApplicationStruct instance.
func NewMTRApplicationBasicClusterApplicationStruct() MTRApplicationBasicClusterApplicationStruct {
	return getMTRApplicationBasicClusterApplicationStructClass().New()
}




