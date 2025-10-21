// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRGeneralDiagnosticsClusterNetworkInterface] class.
var (
	MTRGeneralDiagnosticsClusterNetworkInterfaceClass     _MTRGeneralDiagnosticsClusterNetworkInterfaceClass
	MTRGeneralDiagnosticsClusterNetworkInterfaceClassOnce sync.Once
)

func getMTRGeneralDiagnosticsClusterNetworkInterfaceClass() _MTRGeneralDiagnosticsClusterNetworkInterfaceClass {
	MTRGeneralDiagnosticsClusterNetworkInterfaceClassOnce.Do(func() {
		MTRGeneralDiagnosticsClusterNetworkInterfaceClass = _MTRGeneralDiagnosticsClusterNetworkInterfaceClass{objc.GetClass("MTRGeneralDiagnosticsClusterNetworkInterface")}
	})
	return MTRGeneralDiagnosticsClusterNetworkInterfaceClass
}

type _MTRGeneralDiagnosticsClusterNetworkInterfaceClass struct {
	class objc.Class
}

// An interface definition for the [MTRGeneralDiagnosticsClusterNetworkInterface] class.
type IMTRGeneralDiagnosticsClusterNetworkInterface interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGeneralDiagnosticsClusterNetworkInterface
type MTRGeneralDiagnosticsClusterNetworkInterface struct {
	objectivec.Object
}

// MTRGeneralDiagnosticsClusterNetworkInterfaceFrom constructs a [MTRGeneralDiagnosticsClusterNetworkInterface] from an unsafe.Pointer.
func MTRGeneralDiagnosticsClusterNetworkInterfaceFrom(ptr unsafe.Pointer) MTRGeneralDiagnosticsClusterNetworkInterface {
	return MTRGeneralDiagnosticsClusterNetworkInterface{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRGeneralDiagnosticsClusterNetworkInterfaceClass) Alloc() MTRGeneralDiagnosticsClusterNetworkInterface {
	rv := objc.Send[MTRGeneralDiagnosticsClusterNetworkInterface](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRGeneralDiagnosticsClusterNetworkInterfaceClass) New() MTRGeneralDiagnosticsClusterNetworkInterface {
	rv := objc.Send[MTRGeneralDiagnosticsClusterNetworkInterface](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRGeneralDiagnosticsClusterNetworkInterface) Init() MTRGeneralDiagnosticsClusterNetworkInterface {
	rv := objc.Send[MTRGeneralDiagnosticsClusterNetworkInterface](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRGeneralDiagnosticsClusterNetworkInterface) Autorelease() MTRGeneralDiagnosticsClusterNetworkInterface {
	rv := objc.Send[MTRGeneralDiagnosticsClusterNetworkInterface](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRGeneralDiagnosticsClusterNetworkInterface creates a new MTRGeneralDiagnosticsClusterNetworkInterface instance.
func NewMTRGeneralDiagnosticsClusterNetworkInterface() MTRGeneralDiagnosticsClusterNetworkInterface {
	return getMTRGeneralDiagnosticsClusterNetworkInterfaceClass().New()
}




