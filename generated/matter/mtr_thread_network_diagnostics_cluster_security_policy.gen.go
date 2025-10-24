// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRThreadNetworkDiagnosticsClusterSecurityPolicy */


/* debug [class_header]: Header for MTRThreadNetworkDiagnosticsClusterSecurityPolicy */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRThreadNetworkDiagnosticsClusterSecurityPolicy */
// An interface definition for the [MTRThreadNetworkDiagnosticsClusterSecurityPolicy] class.
type IMTRThreadNetworkDiagnosticsClusterSecurityPolicy interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRThreadNetworkDiagnosticsClusterSecurityPolicy */
	// properties:
	Flags() objc.IObject /* cross-framework: NSNumber */
	SetFlags(value objc.IObject /* cross-framework: NSNumber */)
	RotationTime() objc.IObject /* cross-framework: NSNumber */
	SetRotationTime(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRThreadNetworkDiagnosticsClusterSecurityPolicy */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRThreadNetworkDiagnosticsClusterSecurityPolicy */
// Alloc allocates a new instance without initialization.
func (mc _MTRThreadNetworkDiagnosticsClusterSecurityPolicyClass) Alloc() MTRThreadNetworkDiagnosticsClusterSecurityPolicy {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterSecurityPolicy](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRThreadNetworkDiagnosticsClusterSecurityPolicy */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterSecurityPolicy
type MTRThreadNetworkDiagnosticsClusterSecurityPolicy struct {
	objectivec.Object
}

// MTRThreadNetworkDiagnosticsClusterSecurityPolicyFrom constructs a [MTRThreadNetworkDiagnosticsClusterSecurityPolicy] from an unsafe.Pointer.
func MTRThreadNetworkDiagnosticsClusterSecurityPolicyFrom(ptr unsafe.Pointer) MTRThreadNetworkDiagnosticsClusterSecurityPolicy {
	return MTRThreadNetworkDiagnosticsClusterSecurityPolicy{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRThreadNetworkDiagnosticsClusterSecurityPolicy *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRThreadNetworkDiagnosticsClusterSecurityPolicy */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRThreadNetworkDiagnosticsClusterSecurityPolicy */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRThreadNetworkDiagnosticsClusterSecurityPolicy */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRThreadNetworkDiagnosticsClusterSecurityPolicy */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterSecurityPolicy/flags
func (m_ MTRThreadNetworkDiagnosticsClusterSecurityPolicy) Flags() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("flags"))
	return rv
}/* debug [instance_properties/getter]: flags */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterSecurityPolicy/flags
func (m_ MTRThreadNetworkDiagnosticsClusterSecurityPolicy) SetFlags(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFlags:"), value)
}/* debug [instance_properties/setter]: flags */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterSecurityPolicy/rotationTime
func (m_ MTRThreadNetworkDiagnosticsClusterSecurityPolicy) RotationTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("rotationTime"))
	return rv
}/* debug [instance_properties/getter]: rotationTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRThreadNetworkDiagnosticsClusterSecurityPolicy/rotationTime
func (m_ MTRThreadNetworkDiagnosticsClusterSecurityPolicy) SetRotationTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRotationTime:"), value)
}/* debug [instance_properties/setter]: rotationTime */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRThreadNetworkDiagnosticsClusterSecurityPolicy */



