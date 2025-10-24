// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterBridgedDeviceBasicInformation] class.
var (
	MTRBaseClusterBridgedDeviceBasicInformationClass     _MTRBaseClusterBridgedDeviceBasicInformationClass
	MTRBaseClusterBridgedDeviceBasicInformationClassOnce sync.Once
)

func getMTRBaseClusterBridgedDeviceBasicInformationClass() _MTRBaseClusterBridgedDeviceBasicInformationClass {
	MTRBaseClusterBridgedDeviceBasicInformationClassOnce.Do(func() {
		MTRBaseClusterBridgedDeviceBasicInformationClass = _MTRBaseClusterBridgedDeviceBasicInformationClass{objc.GetClass("MTRBaseClusterBridgedDeviceBasicInformation")}
	})
	return MTRBaseClusterBridgedDeviceBasicInformationClass
}

type _MTRBaseClusterBridgedDeviceBasicInformationClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterBridgedDeviceBasicInformation] class.
type IMTRBaseClusterBridgedDeviceBasicInformation interface {
	IMTRGenericBaseCluster
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterBridgedDeviceBasicInformation
type MTRBaseClusterBridgedDeviceBasicInformation struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterBridgedDeviceBasicInformationFrom constructs a [MTRBaseClusterBridgedDeviceBasicInformation] from an unsafe.Pointer.
func MTRBaseClusterBridgedDeviceBasicInformationFrom(ptr unsafe.Pointer) MTRBaseClusterBridgedDeviceBasicInformation {
	return MTRBaseClusterBridgedDeviceBasicInformation{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterBridgedDeviceBasicInformationClass) Alloc() MTRBaseClusterBridgedDeviceBasicInformation {
	rv := objc.Send[MTRBaseClusterBridgedDeviceBasicInformation](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterBridgedDeviceBasicInformationClass) New() MTRBaseClusterBridgedDeviceBasicInformation {
	rv := objc.Send[MTRBaseClusterBridgedDeviceBasicInformation](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterBridgedDeviceBasicInformation) Init() MTRBaseClusterBridgedDeviceBasicInformation {
	rv := objc.Send[MTRBaseClusterBridgedDeviceBasicInformation](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterBridgedDeviceBasicInformation) Autorelease() MTRBaseClusterBridgedDeviceBasicInformation {
	rv := objc.Send[MTRBaseClusterBridgedDeviceBasicInformation](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterBridgedDeviceBasicInformation creates a new MTRBaseClusterBridgedDeviceBasicInformation instance.
func NewMTRBaseClusterBridgedDeviceBasicInformation() MTRBaseClusterBridgedDeviceBasicInformation {
	return getMTRBaseClusterBridgedDeviceBasicInformationClass().New()
}




