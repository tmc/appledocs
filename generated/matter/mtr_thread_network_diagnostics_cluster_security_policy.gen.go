// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRThreadNetworkDiagnosticsClusterSecurityPolicy] class.
var (
	MTRThreadNetworkDiagnosticsClusterSecurityPolicyClass     _MTRThreadNetworkDiagnosticsClusterSecurityPolicyClass
	MTRThreadNetworkDiagnosticsClusterSecurityPolicyClassOnce sync.Once
)

func getMTRThreadNetworkDiagnosticsClusterSecurityPolicyClass() _MTRThreadNetworkDiagnosticsClusterSecurityPolicyClass {
	MTRThreadNetworkDiagnosticsClusterSecurityPolicyClassOnce.Do(func() {
		MTRThreadNetworkDiagnosticsClusterSecurityPolicyClass = _MTRThreadNetworkDiagnosticsClusterSecurityPolicyClass{objc.GetClass("MTRThreadNetworkDiagnosticsClusterSecurityPolicy")}
	})
	return MTRThreadNetworkDiagnosticsClusterSecurityPolicyClass
}

type _MTRThreadNetworkDiagnosticsClusterSecurityPolicyClass struct {
	class objc.Class
}

// An interface definition for the [MTRThreadNetworkDiagnosticsClusterSecurityPolicy] class.
type IMTRThreadNetworkDiagnosticsClusterSecurityPolicy interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterSecurityPolicy
type MTRThreadNetworkDiagnosticsClusterSecurityPolicy struct {
	objectivec.Object
}

// MTRThreadNetworkDiagnosticsClusterSecurityPolicyFrom constructs a [MTRThreadNetworkDiagnosticsClusterSecurityPolicy] from an unsafe.Pointer.
func MTRThreadNetworkDiagnosticsClusterSecurityPolicyFrom(ptr unsafe.Pointer) MTRThreadNetworkDiagnosticsClusterSecurityPolicy {
	return MTRThreadNetworkDiagnosticsClusterSecurityPolicy{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRThreadNetworkDiagnosticsClusterSecurityPolicyClass) Alloc() MTRThreadNetworkDiagnosticsClusterSecurityPolicy {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterSecurityPolicy](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRThreadNetworkDiagnosticsClusterSecurityPolicyClass) New() MTRThreadNetworkDiagnosticsClusterSecurityPolicy {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterSecurityPolicy](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRThreadNetworkDiagnosticsClusterSecurityPolicy) Init() MTRThreadNetworkDiagnosticsClusterSecurityPolicy {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterSecurityPolicy](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRThreadNetworkDiagnosticsClusterSecurityPolicy) Autorelease() MTRThreadNetworkDiagnosticsClusterSecurityPolicy {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterSecurityPolicy](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRThreadNetworkDiagnosticsClusterSecurityPolicy creates a new MTRThreadNetworkDiagnosticsClusterSecurityPolicy instance.
func NewMTRThreadNetworkDiagnosticsClusterSecurityPolicy() MTRThreadNetworkDiagnosticsClusterSecurityPolicy {
	return getMTRThreadNetworkDiagnosticsClusterSecurityPolicyClass().New()
}




