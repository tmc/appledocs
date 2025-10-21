// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterValveConfigurationAndControl] class.
var (
	MTRClusterValveConfigurationAndControlClass     _MTRClusterValveConfigurationAndControlClass
	MTRClusterValveConfigurationAndControlClassOnce sync.Once
)

func getMTRClusterValveConfigurationAndControlClass() _MTRClusterValveConfigurationAndControlClass {
	MTRClusterValveConfigurationAndControlClassOnce.Do(func() {
		MTRClusterValveConfigurationAndControlClass = _MTRClusterValveConfigurationAndControlClass{objc.GetClass("MTRClusterValveConfigurationAndControl")}
	})
	return MTRClusterValveConfigurationAndControlClass
}

type _MTRClusterValveConfigurationAndControlClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterValveConfigurationAndControl] class.
type IMTRClusterValveConfigurationAndControl interface {
	IMTRGenericCluster
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterValveConfigurationAndControl
type MTRClusterValveConfigurationAndControl struct {
	MTRGenericCluster
}

// MTRClusterValveConfigurationAndControlFrom constructs a [MTRClusterValveConfigurationAndControl] from an unsafe.Pointer.
func MTRClusterValveConfigurationAndControlFrom(ptr unsafe.Pointer) MTRClusterValveConfigurationAndControl {
	return MTRClusterValveConfigurationAndControl{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterValveConfigurationAndControlClass) Alloc() MTRClusterValveConfigurationAndControl {
	rv := objc.Send[MTRClusterValveConfigurationAndControl](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterValveConfigurationAndControlClass) New() MTRClusterValveConfigurationAndControl {
	rv := objc.Send[MTRClusterValveConfigurationAndControl](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterValveConfigurationAndControl) Init() MTRClusterValveConfigurationAndControl {
	rv := objc.Send[MTRClusterValveConfigurationAndControl](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterValveConfigurationAndControl) Autorelease() MTRClusterValveConfigurationAndControl {
	rv := objc.Send[MTRClusterValveConfigurationAndControl](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterValveConfigurationAndControl creates a new MTRClusterValveConfigurationAndControl instance.
func NewMTRClusterValveConfigurationAndControl() MTRClusterValveConfigurationAndControl {
	return getMTRClusterValveConfigurationAndControlClass().New()
}




