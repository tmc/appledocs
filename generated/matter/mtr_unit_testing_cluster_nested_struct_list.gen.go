// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRUnitTestingClusterNestedStructList] class.
var (
	MTRUnitTestingClusterNestedStructListClass     _MTRUnitTestingClusterNestedStructListClass
	MTRUnitTestingClusterNestedStructListClassOnce sync.Once
)

func getMTRUnitTestingClusterNestedStructListClass() _MTRUnitTestingClusterNestedStructListClass {
	MTRUnitTestingClusterNestedStructListClassOnce.Do(func() {
		MTRUnitTestingClusterNestedStructListClass = _MTRUnitTestingClusterNestedStructListClass{objc.GetClass("MTRUnitTestingClusterNestedStructList")}
	})
	return MTRUnitTestingClusterNestedStructListClass
}

type _MTRUnitTestingClusterNestedStructListClass struct {
	class objc.Class
}

// An interface definition for the [MTRUnitTestingClusterNestedStructList] class.
type IMTRUnitTestingClusterNestedStructList interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterNestedStructList
type MTRUnitTestingClusterNestedStructList struct {
	objectivec.Object
}

// MTRUnitTestingClusterNestedStructListFrom constructs a [MTRUnitTestingClusterNestedStructList] from an unsafe.Pointer.
func MTRUnitTestingClusterNestedStructListFrom(ptr unsafe.Pointer) MTRUnitTestingClusterNestedStructList {
	return MTRUnitTestingClusterNestedStructList{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterNestedStructListClass) Alloc() MTRUnitTestingClusterNestedStructList {
	rv := objc.Send[MTRUnitTestingClusterNestedStructList](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRUnitTestingClusterNestedStructListClass) New() MTRUnitTestingClusterNestedStructList {
	rv := objc.Send[MTRUnitTestingClusterNestedStructList](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterNestedStructList) Init() MTRUnitTestingClusterNestedStructList {
	rv := objc.Send[MTRUnitTestingClusterNestedStructList](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterNestedStructList) Autorelease() MTRUnitTestingClusterNestedStructList {
	rv := objc.Send[MTRUnitTestingClusterNestedStructList](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterNestedStructList creates a new MTRUnitTestingClusterNestedStructList instance.
func NewMTRUnitTestingClusterNestedStructList() MTRUnitTestingClusterNestedStructList {
	return getMTRUnitTestingClusterNestedStructListClass().New()
}




