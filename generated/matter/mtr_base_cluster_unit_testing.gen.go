// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterUnitTesting] class.
var (
	MTRBaseClusterUnitTestingClass     _MTRBaseClusterUnitTestingClass
	MTRBaseClusterUnitTestingClassOnce sync.Once
)

func getMTRBaseClusterUnitTestingClass() _MTRBaseClusterUnitTestingClass {
	MTRBaseClusterUnitTestingClassOnce.Do(func() {
		MTRBaseClusterUnitTestingClass = _MTRBaseClusterUnitTestingClass{objc.GetClass("MTRBaseClusterUnitTesting")}
	})
	return MTRBaseClusterUnitTestingClass
}

type _MTRBaseClusterUnitTestingClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterUnitTesting] class.
type IMTRBaseClusterUnitTesting interface {
	IMTRGenericBaseCluster
	// properties:
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterUnitTesting
type MTRBaseClusterUnitTesting struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterUnitTestingFrom constructs a [MTRBaseClusterUnitTesting] from an unsafe.Pointer.
func MTRBaseClusterUnitTestingFrom(ptr unsafe.Pointer) MTRBaseClusterUnitTesting {
	return MTRBaseClusterUnitTesting{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterUnitTestingClass) Alloc() MTRBaseClusterUnitTesting {
	rv := objc.Send[MTRBaseClusterUnitTesting](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterUnitTestingClass) New() MTRBaseClusterUnitTesting {
	rv := objc.Send[MTRBaseClusterUnitTesting](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterUnitTesting) Init() MTRBaseClusterUnitTesting {
	rv := objc.Send[MTRBaseClusterUnitTesting](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterUnitTesting) Autorelease() MTRBaseClusterUnitTesting {
	rv := objc.Send[MTRBaseClusterUnitTesting](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterUnitTesting creates a new MTRBaseClusterUnitTesting instance.
func NewMTRBaseClusterUnitTesting() MTRBaseClusterUnitTesting {
	return getMTRBaseClusterUnitTestingClass().New()
}
