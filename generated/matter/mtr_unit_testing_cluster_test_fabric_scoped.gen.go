// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRUnitTestingClusterTestFabricScoped] class.
var (
	MTRUnitTestingClusterTestFabricScopedClass     _MTRUnitTestingClusterTestFabricScopedClass
	MTRUnitTestingClusterTestFabricScopedClassOnce sync.Once
)

func getMTRUnitTestingClusterTestFabricScopedClass() _MTRUnitTestingClusterTestFabricScopedClass {
	MTRUnitTestingClusterTestFabricScopedClassOnce.Do(func() {
		MTRUnitTestingClusterTestFabricScopedClass = _MTRUnitTestingClusterTestFabricScopedClass{objc.GetClass("MTRUnitTestingClusterTestFabricScoped")}
	})
	return MTRUnitTestingClusterTestFabricScopedClass
}

type _MTRUnitTestingClusterTestFabricScopedClass struct {
	class objc.Class
}

// An interface definition for the [MTRUnitTestingClusterTestFabricScoped] class.
type IMTRUnitTestingClusterTestFabricScoped interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestFabricScoped
type MTRUnitTestingClusterTestFabricScoped struct {
	objectivec.Object
}

// MTRUnitTestingClusterTestFabricScopedFrom constructs a [MTRUnitTestingClusterTestFabricScoped] from an unsafe.Pointer.
func MTRUnitTestingClusterTestFabricScopedFrom(ptr unsafe.Pointer) MTRUnitTestingClusterTestFabricScoped {
	return MTRUnitTestingClusterTestFabricScoped{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterTestFabricScopedClass) Alloc() MTRUnitTestingClusterTestFabricScoped {
	rv := objc.Send[MTRUnitTestingClusterTestFabricScoped](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRUnitTestingClusterTestFabricScopedClass) New() MTRUnitTestingClusterTestFabricScoped {
	rv := objc.Send[MTRUnitTestingClusterTestFabricScoped](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterTestFabricScoped) Init() MTRUnitTestingClusterTestFabricScoped {
	rv := objc.Send[MTRUnitTestingClusterTestFabricScoped](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterTestFabricScoped) Autorelease() MTRUnitTestingClusterTestFabricScoped {
	rv := objc.Send[MTRUnitTestingClusterTestFabricScoped](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterTestFabricScoped creates a new MTRUnitTestingClusterTestFabricScoped instance.
func NewMTRUnitTestingClusterTestFabricScoped() MTRUnitTestingClusterTestFabricScoped {
	return getMTRUnitTestingClusterTestFabricScopedClass().New()
}




