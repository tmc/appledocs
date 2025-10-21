// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterBridgedDeviceBasic] class.
var (
	MTRClusterBridgedDeviceBasicClass     _MTRClusterBridgedDeviceBasicClass
	MTRClusterBridgedDeviceBasicClassOnce sync.Once
)

func getMTRClusterBridgedDeviceBasicClass() _MTRClusterBridgedDeviceBasicClass {
	MTRClusterBridgedDeviceBasicClassOnce.Do(func() {
		MTRClusterBridgedDeviceBasicClass = _MTRClusterBridgedDeviceBasicClass{objc.GetClass("MTRClusterBridgedDeviceBasic")}
	})
	return MTRClusterBridgedDeviceBasicClass
}

type _MTRClusterBridgedDeviceBasicClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterBridgedDeviceBasic] class.
type IMTRClusterBridgedDeviceBasic interface {
	IMTRClusterBridgedDeviceBasicInformation
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterBridgedDeviceBasic
type MTRClusterBridgedDeviceBasic struct {
	MTRClusterBridgedDeviceBasicInformation
}

// MTRClusterBridgedDeviceBasicFrom constructs a [MTRClusterBridgedDeviceBasic] from an unsafe.Pointer.
func MTRClusterBridgedDeviceBasicFrom(ptr unsafe.Pointer) MTRClusterBridgedDeviceBasic {
	return MTRClusterBridgedDeviceBasic{
		MTRClusterBridgedDeviceBasicInformation: MTRClusterBridgedDeviceBasicInformationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterBridgedDeviceBasicClass) Alloc() MTRClusterBridgedDeviceBasic {
	rv := objc.Send[MTRClusterBridgedDeviceBasic](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterBridgedDeviceBasicClass) New() MTRClusterBridgedDeviceBasic {
	rv := objc.Send[MTRClusterBridgedDeviceBasic](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterBridgedDeviceBasic) Init() MTRClusterBridgedDeviceBasic {
	rv := objc.Send[MTRClusterBridgedDeviceBasic](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterBridgedDeviceBasic) Autorelease() MTRClusterBridgedDeviceBasic {
	rv := objc.Send[MTRClusterBridgedDeviceBasic](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterBridgedDeviceBasic creates a new MTRClusterBridgedDeviceBasic instance.
func NewMTRClusterBridgedDeviceBasic() MTRClusterBridgedDeviceBasic {
	return getMTRClusterBridgedDeviceBasicClass().New()
}




