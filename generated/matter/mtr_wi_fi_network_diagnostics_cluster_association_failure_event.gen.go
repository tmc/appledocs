// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRWiFiNetworkDiagnosticsClusterAssociationFailureEvent */


/* debug [class_header]: Header for MTRWiFiNetworkDiagnosticsClusterAssociationFailureEvent */
// The class instance for the [MTRWiFiNetworkDiagnosticsClusterAssociationFailureEvent] class.
var (
	MTRWiFiNetworkDiagnosticsClusterAssociationFailureEventClass     _MTRWiFiNetworkDiagnosticsClusterAssociationFailureEventClass
	MTRWiFiNetworkDiagnosticsClusterAssociationFailureEventClassOnce sync.Once
)

func getMTRWiFiNetworkDiagnosticsClusterAssociationFailureEventClass() _MTRWiFiNetworkDiagnosticsClusterAssociationFailureEventClass {
	MTRWiFiNetworkDiagnosticsClusterAssociationFailureEventClassOnce.Do(func() {
		MTRWiFiNetworkDiagnosticsClusterAssociationFailureEventClass = _MTRWiFiNetworkDiagnosticsClusterAssociationFailureEventClass{objc.GetClass("MTRWiFiNetworkDiagnosticsClusterAssociationFailureEvent")}
	})
	return MTRWiFiNetworkDiagnosticsClusterAssociationFailureEventClass
}

type _MTRWiFiNetworkDiagnosticsClusterAssociationFailureEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRWiFiNetworkDiagnosticsClusterAssociationFailureEvent */
// An interface definition for the [MTRWiFiNetworkDiagnosticsClusterAssociationFailureEvent] class.
type IMTRWiFiNetworkDiagnosticsClusterAssociationFailureEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRWiFiNetworkDiagnosticsClusterAssociationFailureEvent */
	// properties:
	AssociationFailure() objc.IObject /* cross-framework: NSNumber */
	SetAssociationFailure(value objc.IObject /* cross-framework: NSNumber */)
	AssociationFailureCause() objc.IObject /* cross-framework: NSNumber */
	SetAssociationFailureCause(value objc.IObject /* cross-framework: NSNumber */)
	Status() objc.IObject /* cross-framework: NSNumber */
	SetStatus(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRWiFiNetworkDiagnosticsClusterAssociationFailureEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRWiFiNetworkDiagnosticsClusterAssociationFailureEvent */
// Alloc allocates a new instance without initialization.
func (mc _MTRWiFiNetworkDiagnosticsClusterAssociationFailureEventClass) Alloc() MTRWiFiNetworkDiagnosticsClusterAssociationFailureEvent {
	rv := objc.Send[MTRWiFiNetworkDiagnosticsClusterAssociationFailureEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRWiFiNetworkDiagnosticsClusterAssociationFailureEventClass) New() MTRWiFiNetworkDiagnosticsClusterAssociationFailureEvent {
	rv := objc.Send[MTRWiFiNetworkDiagnosticsClusterAssociationFailureEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRWiFiNetworkDiagnosticsClusterAssociationFailureEvent) Init() MTRWiFiNetworkDiagnosticsClusterAssociationFailureEvent {
	rv := objc.Send[MTRWiFiNetworkDiagnosticsClusterAssociationFailureEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRWiFiNetworkDiagnosticsClusterAssociationFailureEvent) Autorelease() MTRWiFiNetworkDiagnosticsClusterAssociationFailureEvent {
	rv := objc.Send[MTRWiFiNetworkDiagnosticsClusterAssociationFailureEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRWiFiNetworkDiagnosticsClusterAssociationFailureEvent creates a new MTRWiFiNetworkDiagnosticsClusterAssociationFailureEvent instance.
func NewMTRWiFiNetworkDiagnosticsClusterAssociationFailureEvent() MTRWiFiNetworkDiagnosticsClusterAssociationFailureEvent {
	return getMTRWiFiNetworkDiagnosticsClusterAssociationFailureEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRWiFiNetworkDiagnosticsClusterAssociationFailureEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWiFiNetworkDiagnosticsClusterAssociationFailureEvent
type MTRWiFiNetworkDiagnosticsClusterAssociationFailureEvent struct {
	objectivec.Object
}

// MTRWiFiNetworkDiagnosticsClusterAssociationFailureEventFrom constructs a [MTRWiFiNetworkDiagnosticsClusterAssociationFailureEvent] from an unsafe.Pointer.
func MTRWiFiNetworkDiagnosticsClusterAssociationFailureEventFrom(ptr unsafe.Pointer) MTRWiFiNetworkDiagnosticsClusterAssociationFailureEvent {
	return MTRWiFiNetworkDiagnosticsClusterAssociationFailureEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRWiFiNetworkDiagnosticsClusterAssociationFailureEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRWiFiNetworkDiagnosticsClusterAssociationFailureEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRWiFiNetworkDiagnosticsClusterAssociationFailureEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRWiFiNetworkDiagnosticsClusterAssociationFailureEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRWiFiNetworkDiagnosticsClusterAssociationFailureEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWiFiNetworkDiagnosticsClusterAssociationFailureEvent/associationFailure
func (m_ MTRWiFiNetworkDiagnosticsClusterAssociationFailureEvent) AssociationFailure() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("associationFailure"))
	return rv
}/* debug [instance_properties/getter]: associationFailure */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWiFiNetworkDiagnosticsClusterAssociationFailureEvent/associationFailure
func (m_ MTRWiFiNetworkDiagnosticsClusterAssociationFailureEvent) SetAssociationFailure(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAssociationFailure:"), value)
}/* debug [instance_properties/setter]: associationFailure */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWiFiNetworkDiagnosticsClusterAssociationFailureEvent/associationFailureCause
func (m_ MTRWiFiNetworkDiagnosticsClusterAssociationFailureEvent) AssociationFailureCause() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("associationFailureCause"))
	return rv
}/* debug [instance_properties/getter]: associationFailureCause */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWiFiNetworkDiagnosticsClusterAssociationFailureEvent/associationFailureCause
func (m_ MTRWiFiNetworkDiagnosticsClusterAssociationFailureEvent) SetAssociationFailureCause(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAssociationFailureCause:"), value)
}/* debug [instance_properties/setter]: associationFailureCause */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWiFiNetworkDiagnosticsClusterAssociationFailureEvent/status
func (m_ MTRWiFiNetworkDiagnosticsClusterAssociationFailureEvent) Status() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("status"))
	return rv
}/* debug [instance_properties/getter]: status */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRWiFiNetworkDiagnosticsClusterAssociationFailureEvent/status
func (m_ MTRWiFiNetworkDiagnosticsClusterAssociationFailureEvent) SetStatus(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatus:"), value)
}/* debug [instance_properties/setter]: status */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRWiFiNetworkDiagnosticsClusterAssociationFailureEvent */



