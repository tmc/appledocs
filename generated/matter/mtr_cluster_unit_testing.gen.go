// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterUnitTesting] class.
var (
	MTRClusterUnitTestingClass     _MTRClusterUnitTestingClass
	MTRClusterUnitTestingClassOnce sync.Once
)

func getMTRClusterUnitTestingClass() _MTRClusterUnitTestingClass {
	MTRClusterUnitTestingClassOnce.Do(func() {
		MTRClusterUnitTestingClass = _MTRClusterUnitTestingClass{objc.GetClass("MTRClusterUnitTesting")}
	})
	return MTRClusterUnitTestingClass
}

type _MTRClusterUnitTestingClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterUnitTesting] class.
type IMTRClusterUnitTesting interface {
	IMTRGenericCluster
	// properties:
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterUnitTesting
type MTRClusterUnitTesting struct {
	MTRGenericCluster
}

// MTRClusterUnitTestingFrom constructs a [MTRClusterUnitTesting] from an unsafe.Pointer.
func MTRClusterUnitTestingFrom(ptr unsafe.Pointer) MTRClusterUnitTesting {
	return MTRClusterUnitTesting{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterUnitTestingClass) Alloc() MTRClusterUnitTesting {
	rv := objc.Send[MTRClusterUnitTesting](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterUnitTestingClass) New() MTRClusterUnitTesting {
	rv := objc.Send[MTRClusterUnitTesting](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterUnitTesting) Init() MTRClusterUnitTesting {
	rv := objc.Send[MTRClusterUnitTesting](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterUnitTesting) Autorelease() MTRClusterUnitTesting {
	rv := objc.Send[MTRClusterUnitTesting](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterUnitTesting creates a new MTRClusterUnitTesting instance.
func NewMTRClusterUnitTesting() MTRClusterUnitTesting {
	return getMTRClusterUnitTestingClass().New()
}
