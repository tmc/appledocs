// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRGeneralDiagnosticsClusterNetworkInterfaceType] class.
var (
	MTRGeneralDiagnosticsClusterNetworkInterfaceTypeClass     _MTRGeneralDiagnosticsClusterNetworkInterfaceTypeClass
	MTRGeneralDiagnosticsClusterNetworkInterfaceTypeClassOnce sync.Once
)

func getMTRGeneralDiagnosticsClusterNetworkInterfaceTypeClass() _MTRGeneralDiagnosticsClusterNetworkInterfaceTypeClass {
	MTRGeneralDiagnosticsClusterNetworkInterfaceTypeClassOnce.Do(func() {
		MTRGeneralDiagnosticsClusterNetworkInterfaceTypeClass = _MTRGeneralDiagnosticsClusterNetworkInterfaceTypeClass{objc.GetClass("MTRGeneralDiagnosticsClusterNetworkInterfaceType")}
	})
	return MTRGeneralDiagnosticsClusterNetworkInterfaceTypeClass
}

type _MTRGeneralDiagnosticsClusterNetworkInterfaceTypeClass struct {
	class objc.Class
}

// An interface definition for the [MTRGeneralDiagnosticsClusterNetworkInterfaceType] class.
type IMTRGeneralDiagnosticsClusterNetworkInterfaceType interface {
	IMTRGeneralDiagnosticsClusterNetworkInterface
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGeneralDiagnosticsClusterNetworkInterfaceType
type MTRGeneralDiagnosticsClusterNetworkInterfaceType struct {
	MTRGeneralDiagnosticsClusterNetworkInterface
}

// MTRGeneralDiagnosticsClusterNetworkInterfaceTypeFrom constructs a [MTRGeneralDiagnosticsClusterNetworkInterfaceType] from an unsafe.Pointer.
func MTRGeneralDiagnosticsClusterNetworkInterfaceTypeFrom(ptr unsafe.Pointer) MTRGeneralDiagnosticsClusterNetworkInterfaceType {
	return MTRGeneralDiagnosticsClusterNetworkInterfaceType{
		MTRGeneralDiagnosticsClusterNetworkInterface: MTRGeneralDiagnosticsClusterNetworkInterfaceFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRGeneralDiagnosticsClusterNetworkInterfaceTypeClass) Alloc() MTRGeneralDiagnosticsClusterNetworkInterfaceType {
	rv := objc.Send[MTRGeneralDiagnosticsClusterNetworkInterfaceType](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRGeneralDiagnosticsClusterNetworkInterfaceTypeClass) New() MTRGeneralDiagnosticsClusterNetworkInterfaceType {
	rv := objc.Send[MTRGeneralDiagnosticsClusterNetworkInterfaceType](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRGeneralDiagnosticsClusterNetworkInterfaceType) Init() MTRGeneralDiagnosticsClusterNetworkInterfaceType {
	rv := objc.Send[MTRGeneralDiagnosticsClusterNetworkInterfaceType](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRGeneralDiagnosticsClusterNetworkInterfaceType) Autorelease() MTRGeneralDiagnosticsClusterNetworkInterfaceType {
	rv := objc.Send[MTRGeneralDiagnosticsClusterNetworkInterfaceType](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRGeneralDiagnosticsClusterNetworkInterfaceType creates a new MTRGeneralDiagnosticsClusterNetworkInterfaceType instance.
func NewMTRGeneralDiagnosticsClusterNetworkInterfaceType() MTRGeneralDiagnosticsClusterNetworkInterfaceType {
	return getMTRGeneralDiagnosticsClusterNetworkInterfaceTypeClass().New()
}




