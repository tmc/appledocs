// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterBridgedDeviceBasicInformation] class.
var (
	MTRClusterBridgedDeviceBasicInformationClass     _MTRClusterBridgedDeviceBasicInformationClass
	MTRClusterBridgedDeviceBasicInformationClassOnce sync.Once
)

func getMTRClusterBridgedDeviceBasicInformationClass() _MTRClusterBridgedDeviceBasicInformationClass {
	MTRClusterBridgedDeviceBasicInformationClassOnce.Do(func() {
		MTRClusterBridgedDeviceBasicInformationClass = _MTRClusterBridgedDeviceBasicInformationClass{objc.GetClass("MTRClusterBridgedDeviceBasicInformation")}
	})
	return MTRClusterBridgedDeviceBasicInformationClass
}

type _MTRClusterBridgedDeviceBasicInformationClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterBridgedDeviceBasicInformation] class.
type IMTRClusterBridgedDeviceBasicInformation interface {
	IMTRGenericCluster
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterBridgedDeviceBasicInformation
type MTRClusterBridgedDeviceBasicInformation struct {
	MTRGenericCluster
}

// MTRClusterBridgedDeviceBasicInformationFrom constructs a [MTRClusterBridgedDeviceBasicInformation] from an unsafe.Pointer.
func MTRClusterBridgedDeviceBasicInformationFrom(ptr unsafe.Pointer) MTRClusterBridgedDeviceBasicInformation {
	return MTRClusterBridgedDeviceBasicInformation{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterBridgedDeviceBasicInformationClass) Alloc() MTRClusterBridgedDeviceBasicInformation {
	rv := objc.Send[MTRClusterBridgedDeviceBasicInformation](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterBridgedDeviceBasicInformationClass) New() MTRClusterBridgedDeviceBasicInformation {
	rv := objc.Send[MTRClusterBridgedDeviceBasicInformation](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterBridgedDeviceBasicInformation) Init() MTRClusterBridgedDeviceBasicInformation {
	rv := objc.Send[MTRClusterBridgedDeviceBasicInformation](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterBridgedDeviceBasicInformation) Autorelease() MTRClusterBridgedDeviceBasicInformation {
	rv := objc.Send[MTRClusterBridgedDeviceBasicInformation](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterBridgedDeviceBasicInformation creates a new MTRClusterBridgedDeviceBasicInformation instance.
func NewMTRClusterBridgedDeviceBasicInformation() MTRClusterBridgedDeviceBasicInformation {
	return getMTRClusterBridgedDeviceBasicInformationClass().New()
}




