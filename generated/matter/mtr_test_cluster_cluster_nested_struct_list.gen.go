// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRTestClusterClusterNestedStructList] class.
var (
	MTRTestClusterClusterNestedStructListClass     _MTRTestClusterClusterNestedStructListClass
	MTRTestClusterClusterNestedStructListClassOnce sync.Once
)

func getMTRTestClusterClusterNestedStructListClass() _MTRTestClusterClusterNestedStructListClass {
	MTRTestClusterClusterNestedStructListClassOnce.Do(func() {
		MTRTestClusterClusterNestedStructListClass = _MTRTestClusterClusterNestedStructListClass{objc.GetClass("MTRTestClusterClusterNestedStructList")}
	})
	return MTRTestClusterClusterNestedStructListClass
}

type _MTRTestClusterClusterNestedStructListClass struct {
	class objc.Class
}

// An interface definition for the [MTRTestClusterClusterNestedStructList] class.
type IMTRTestClusterClusterNestedStructList interface {
	IMTRUnitTestingClusterNestedStructList
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTestClusterClusterNestedStructList
type MTRTestClusterClusterNestedStructList struct {
	MTRUnitTestingClusterNestedStructList
}

// MTRTestClusterClusterNestedStructListFrom constructs a [MTRTestClusterClusterNestedStructList] from an unsafe.Pointer.
func MTRTestClusterClusterNestedStructListFrom(ptr unsafe.Pointer) MTRTestClusterClusterNestedStructList {
	return MTRTestClusterClusterNestedStructList{
		MTRUnitTestingClusterNestedStructList: MTRUnitTestingClusterNestedStructListFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTestClusterClusterNestedStructListClass) Alloc() MTRTestClusterClusterNestedStructList {
	rv := objc.Send[MTRTestClusterClusterNestedStructList](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTestClusterClusterNestedStructListClass) New() MTRTestClusterClusterNestedStructList {
	rv := objc.Send[MTRTestClusterClusterNestedStructList](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTestClusterClusterNestedStructList) Init() MTRTestClusterClusterNestedStructList {
	rv := objc.Send[MTRTestClusterClusterNestedStructList](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTestClusterClusterNestedStructList) Autorelease() MTRTestClusterClusterNestedStructList {
	rv := objc.Send[MTRTestClusterClusterNestedStructList](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTestClusterClusterNestedStructList creates a new MTRTestClusterClusterNestedStructList instance.
func NewMTRTestClusterClusterNestedStructList() MTRTestClusterClusterNestedStructList {
	return getMTRTestClusterClusterNestedStructListClass().New()
}




