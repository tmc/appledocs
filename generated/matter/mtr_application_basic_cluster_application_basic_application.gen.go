// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRApplicationBasicClusterApplicationBasicApplication] class.
var (
	MTRApplicationBasicClusterApplicationBasicApplicationClass     _MTRApplicationBasicClusterApplicationBasicApplicationClass
	MTRApplicationBasicClusterApplicationBasicApplicationClassOnce sync.Once
)

func getMTRApplicationBasicClusterApplicationBasicApplicationClass() _MTRApplicationBasicClusterApplicationBasicApplicationClass {
	MTRApplicationBasicClusterApplicationBasicApplicationClassOnce.Do(func() {
		MTRApplicationBasicClusterApplicationBasicApplicationClass = _MTRApplicationBasicClusterApplicationBasicApplicationClass{objc.GetClass("MTRApplicationBasicClusterApplicationBasicApplication")}
	})
	return MTRApplicationBasicClusterApplicationBasicApplicationClass
}

type _MTRApplicationBasicClusterApplicationBasicApplicationClass struct {
	class objc.Class
}

// An interface definition for the [MTRApplicationBasicClusterApplicationBasicApplication] class.
type IMTRApplicationBasicClusterApplicationBasicApplication interface {
	IMTRApplicationBasicClusterApplicationStruct
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRApplicationBasicClusterApplicationBasicApplication
type MTRApplicationBasicClusterApplicationBasicApplication struct {
	MTRApplicationBasicClusterApplicationStruct
}

// MTRApplicationBasicClusterApplicationBasicApplicationFrom constructs a [MTRApplicationBasicClusterApplicationBasicApplication] from an unsafe.Pointer.
func MTRApplicationBasicClusterApplicationBasicApplicationFrom(ptr unsafe.Pointer) MTRApplicationBasicClusterApplicationBasicApplication {
	return MTRApplicationBasicClusterApplicationBasicApplication{
		MTRApplicationBasicClusterApplicationStruct: MTRApplicationBasicClusterApplicationStructFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRApplicationBasicClusterApplicationBasicApplicationClass) Alloc() MTRApplicationBasicClusterApplicationBasicApplication {
	rv := objc.Send[MTRApplicationBasicClusterApplicationBasicApplication](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRApplicationBasicClusterApplicationBasicApplicationClass) New() MTRApplicationBasicClusterApplicationBasicApplication {
	rv := objc.Send[MTRApplicationBasicClusterApplicationBasicApplication](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRApplicationBasicClusterApplicationBasicApplication) Init() MTRApplicationBasicClusterApplicationBasicApplication {
	rv := objc.Send[MTRApplicationBasicClusterApplicationBasicApplication](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRApplicationBasicClusterApplicationBasicApplication) Autorelease() MTRApplicationBasicClusterApplicationBasicApplication {
	rv := objc.Send[MTRApplicationBasicClusterApplicationBasicApplication](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRApplicationBasicClusterApplicationBasicApplication creates a new MTRApplicationBasicClusterApplicationBasicApplication instance.
func NewMTRApplicationBasicClusterApplicationBasicApplication() MTRApplicationBasicClusterApplicationBasicApplication {
	return getMTRApplicationBasicClusterApplicationBasicApplicationClass().New()
}




