// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterValveConfigurationAndControl] class.
var (
	MTRBaseClusterValveConfigurationAndControlClass     _MTRBaseClusterValveConfigurationAndControlClass
	MTRBaseClusterValveConfigurationAndControlClassOnce sync.Once
)

func getMTRBaseClusterValveConfigurationAndControlClass() _MTRBaseClusterValveConfigurationAndControlClass {
	MTRBaseClusterValveConfigurationAndControlClassOnce.Do(func() {
		MTRBaseClusterValveConfigurationAndControlClass = _MTRBaseClusterValveConfigurationAndControlClass{objc.GetClass("MTRBaseClusterValveConfigurationAndControl")}
	})
	return MTRBaseClusterValveConfigurationAndControlClass
}

type _MTRBaseClusterValveConfigurationAndControlClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterValveConfigurationAndControl] class.
type IMTRBaseClusterValveConfigurationAndControl interface {
	IMTRGenericBaseCluster
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterValveConfigurationAndControl
type MTRBaseClusterValveConfigurationAndControl struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterValveConfigurationAndControlFrom constructs a [MTRBaseClusterValveConfigurationAndControl] from an unsafe.Pointer.
func MTRBaseClusterValveConfigurationAndControlFrom(ptr unsafe.Pointer) MTRBaseClusterValveConfigurationAndControl {
	return MTRBaseClusterValveConfigurationAndControl{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterValveConfigurationAndControlClass) Alloc() MTRBaseClusterValveConfigurationAndControl {
	rv := objc.Send[MTRBaseClusterValveConfigurationAndControl](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterValveConfigurationAndControlClass) New() MTRBaseClusterValveConfigurationAndControl {
	rv := objc.Send[MTRBaseClusterValveConfigurationAndControl](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterValveConfigurationAndControl) Init() MTRBaseClusterValveConfigurationAndControl {
	rv := objc.Send[MTRBaseClusterValveConfigurationAndControl](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterValveConfigurationAndControl) Autorelease() MTRBaseClusterValveConfigurationAndControl {
	rv := objc.Send[MTRBaseClusterValveConfigurationAndControl](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterValveConfigurationAndControl creates a new MTRBaseClusterValveConfigurationAndControl instance.
func NewMTRBaseClusterValveConfigurationAndControl() MTRBaseClusterValveConfigurationAndControl {
	return getMTRBaseClusterValveConfigurationAndControlClass().New()
}




