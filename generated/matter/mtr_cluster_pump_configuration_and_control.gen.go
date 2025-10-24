// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterPumpConfigurationAndControl] class.
var (
	MTRClusterPumpConfigurationAndControlClass     _MTRClusterPumpConfigurationAndControlClass
	MTRClusterPumpConfigurationAndControlClassOnce sync.Once
)

func getMTRClusterPumpConfigurationAndControlClass() _MTRClusterPumpConfigurationAndControlClass {
	MTRClusterPumpConfigurationAndControlClassOnce.Do(func() {
		MTRClusterPumpConfigurationAndControlClass = _MTRClusterPumpConfigurationAndControlClass{objc.GetClass("MTRClusterPumpConfigurationAndControl")}
	})
	return MTRClusterPumpConfigurationAndControlClass
}

type _MTRClusterPumpConfigurationAndControlClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterPumpConfigurationAndControl] class.
type IMTRClusterPumpConfigurationAndControl interface {
	IMTRGenericCluster
	// properties:
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterPumpConfigurationAndControl
type MTRClusterPumpConfigurationAndControl struct {
	MTRGenericCluster
}

// MTRClusterPumpConfigurationAndControlFrom constructs a [MTRClusterPumpConfigurationAndControl] from an unsafe.Pointer.
func MTRClusterPumpConfigurationAndControlFrom(ptr unsafe.Pointer) MTRClusterPumpConfigurationAndControl {
	return MTRClusterPumpConfigurationAndControl{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterPumpConfigurationAndControlClass) Alloc() MTRClusterPumpConfigurationAndControl {
	rv := objc.Send[MTRClusterPumpConfigurationAndControl](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterPumpConfigurationAndControlClass) New() MTRClusterPumpConfigurationAndControl {
	rv := objc.Send[MTRClusterPumpConfigurationAndControl](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterPumpConfigurationAndControl) Init() MTRClusterPumpConfigurationAndControl {
	rv := objc.Send[MTRClusterPumpConfigurationAndControl](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterPumpConfigurationAndControl) Autorelease() MTRClusterPumpConfigurationAndControl {
	rv := objc.Send[MTRClusterPumpConfigurationAndControl](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterPumpConfigurationAndControl creates a new MTRClusterPumpConfigurationAndControl instance.
func NewMTRClusterPumpConfigurationAndControl() MTRClusterPumpConfigurationAndControl {
	return getMTRClusterPumpConfigurationAndControlClass().New()
}
