// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRTestClusterClusterTestFabricScoped] class.
var (
	MTRTestClusterClusterTestFabricScopedClass     _MTRTestClusterClusterTestFabricScopedClass
	MTRTestClusterClusterTestFabricScopedClassOnce sync.Once
)

func getMTRTestClusterClusterTestFabricScopedClass() _MTRTestClusterClusterTestFabricScopedClass {
	MTRTestClusterClusterTestFabricScopedClassOnce.Do(func() {
		MTRTestClusterClusterTestFabricScopedClass = _MTRTestClusterClusterTestFabricScopedClass{objc.GetClass("MTRTestClusterClusterTestFabricScoped")}
	})
	return MTRTestClusterClusterTestFabricScopedClass
}

type _MTRTestClusterClusterTestFabricScopedClass struct {
	class objc.Class
}

// An interface definition for the [MTRTestClusterClusterTestFabricScoped] class.
type IMTRTestClusterClusterTestFabricScoped interface {
	IMTRUnitTestingClusterTestFabricScoped
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTestClusterClusterTestFabricScoped
type MTRTestClusterClusterTestFabricScoped struct {
	MTRUnitTestingClusterTestFabricScoped
}

// MTRTestClusterClusterTestFabricScopedFrom constructs a [MTRTestClusterClusterTestFabricScoped] from an unsafe.Pointer.
func MTRTestClusterClusterTestFabricScopedFrom(ptr unsafe.Pointer) MTRTestClusterClusterTestFabricScoped {
	return MTRTestClusterClusterTestFabricScoped{
		MTRUnitTestingClusterTestFabricScoped: MTRUnitTestingClusterTestFabricScopedFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTestClusterClusterTestFabricScopedClass) Alloc() MTRTestClusterClusterTestFabricScoped {
	rv := objc.Send[MTRTestClusterClusterTestFabricScoped](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTestClusterClusterTestFabricScopedClass) New() MTRTestClusterClusterTestFabricScoped {
	rv := objc.Send[MTRTestClusterClusterTestFabricScoped](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTestClusterClusterTestFabricScoped) Init() MTRTestClusterClusterTestFabricScoped {
	rv := objc.Send[MTRTestClusterClusterTestFabricScoped](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTestClusterClusterTestFabricScoped) Autorelease() MTRTestClusterClusterTestFabricScoped {
	rv := objc.Send[MTRTestClusterClusterTestFabricScoped](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTestClusterClusterTestFabricScoped creates a new MTRTestClusterClusterTestFabricScoped instance.
func NewMTRTestClusterClusterTestFabricScoped() MTRTestClusterClusterTestFabricScoped {
	return getMTRTestClusterClusterTestFabricScopedClass().New()
}




