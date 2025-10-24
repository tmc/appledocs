// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterPumpConfigurationAndControl] class.
var (
	MTRBaseClusterPumpConfigurationAndControlClass     _MTRBaseClusterPumpConfigurationAndControlClass
	MTRBaseClusterPumpConfigurationAndControlClassOnce sync.Once
)

func getMTRBaseClusterPumpConfigurationAndControlClass() _MTRBaseClusterPumpConfigurationAndControlClass {
	MTRBaseClusterPumpConfigurationAndControlClassOnce.Do(func() {
		MTRBaseClusterPumpConfigurationAndControlClass = _MTRBaseClusterPumpConfigurationAndControlClass{objc.GetClass("MTRBaseClusterPumpConfigurationAndControl")}
	})
	return MTRBaseClusterPumpConfigurationAndControlClass
}

type _MTRBaseClusterPumpConfigurationAndControlClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterPumpConfigurationAndControl] class.
type IMTRBaseClusterPumpConfigurationAndControl interface {
	IMTRGenericBaseCluster
	// properties:
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterPumpConfigurationAndControl
type MTRBaseClusterPumpConfigurationAndControl struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterPumpConfigurationAndControlFrom constructs a [MTRBaseClusterPumpConfigurationAndControl] from an unsafe.Pointer.
func MTRBaseClusterPumpConfigurationAndControlFrom(ptr unsafe.Pointer) MTRBaseClusterPumpConfigurationAndControl {
	return MTRBaseClusterPumpConfigurationAndControl{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterPumpConfigurationAndControlClass) Alloc() MTRBaseClusterPumpConfigurationAndControl {
	rv := objc.Send[MTRBaseClusterPumpConfigurationAndControl](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterPumpConfigurationAndControlClass) New() MTRBaseClusterPumpConfigurationAndControl {
	rv := objc.Send[MTRBaseClusterPumpConfigurationAndControl](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterPumpConfigurationAndControl) Init() MTRBaseClusterPumpConfigurationAndControl {
	rv := objc.Send[MTRBaseClusterPumpConfigurationAndControl](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterPumpConfigurationAndControl) Autorelease() MTRBaseClusterPumpConfigurationAndControl {
	rv := objc.Send[MTRBaseClusterPumpConfigurationAndControl](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterPumpConfigurationAndControl creates a new MTRBaseClusterPumpConfigurationAndControl instance.
func NewMTRBaseClusterPumpConfigurationAndControl() MTRBaseClusterPumpConfigurationAndControl {
	return getMTRBaseClusterPumpConfigurationAndControlClass().New()
}
