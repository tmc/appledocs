// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterBridgedDeviceBasic] class.
var (
	MTRBaseClusterBridgedDeviceBasicClass     _MTRBaseClusterBridgedDeviceBasicClass
	MTRBaseClusterBridgedDeviceBasicClassOnce sync.Once
)

func getMTRBaseClusterBridgedDeviceBasicClass() _MTRBaseClusterBridgedDeviceBasicClass {
	MTRBaseClusterBridgedDeviceBasicClassOnce.Do(func() {
		MTRBaseClusterBridgedDeviceBasicClass = _MTRBaseClusterBridgedDeviceBasicClass{objc.GetClass("MTRBaseClusterBridgedDeviceBasic")}
	})
	return MTRBaseClusterBridgedDeviceBasicClass
}

type _MTRBaseClusterBridgedDeviceBasicClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterBridgedDeviceBasic] class.
type IMTRBaseClusterBridgedDeviceBasic interface {
	IMTRBaseClusterBridgedDeviceBasicInformation
	// properties:
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterBridgedDeviceBasic
type MTRBaseClusterBridgedDeviceBasic struct {
	MTRBaseClusterBridgedDeviceBasicInformation
}

// MTRBaseClusterBridgedDeviceBasicFrom constructs a [MTRBaseClusterBridgedDeviceBasic] from an unsafe.Pointer.
func MTRBaseClusterBridgedDeviceBasicFrom(ptr unsafe.Pointer) MTRBaseClusterBridgedDeviceBasic {
	return MTRBaseClusterBridgedDeviceBasic{
		MTRBaseClusterBridgedDeviceBasicInformation: MTRBaseClusterBridgedDeviceBasicInformationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterBridgedDeviceBasicClass) Alloc() MTRBaseClusterBridgedDeviceBasic {
	rv := objc.Send[MTRBaseClusterBridgedDeviceBasic](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterBridgedDeviceBasicClass) New() MTRBaseClusterBridgedDeviceBasic {
	rv := objc.Send[MTRBaseClusterBridgedDeviceBasic](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterBridgedDeviceBasic) Init() MTRBaseClusterBridgedDeviceBasic {
	rv := objc.Send[MTRBaseClusterBridgedDeviceBasic](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterBridgedDeviceBasic) Autorelease() MTRBaseClusterBridgedDeviceBasic {
	rv := objc.Send[MTRBaseClusterBridgedDeviceBasic](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterBridgedDeviceBasic creates a new MTRBaseClusterBridgedDeviceBasic instance.
func NewMTRBaseClusterBridgedDeviceBasic() MTRBaseClusterBridgedDeviceBasic {
	return getMTRBaseClusterBridgedDeviceBasicClass().New()
}
