// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRSoftwareDiagnosticsClusterSoftwareFaultEvent */


/* debug [class_header]: Header for MTRSoftwareDiagnosticsClusterSoftwareFaultEvent */
// The class instance for the [MTRSoftwareDiagnosticsClusterSoftwareFaultEvent] class.
var (
	MTRSoftwareDiagnosticsClusterSoftwareFaultEventClass     _MTRSoftwareDiagnosticsClusterSoftwareFaultEventClass
	MTRSoftwareDiagnosticsClusterSoftwareFaultEventClassOnce sync.Once
)

func getMTRSoftwareDiagnosticsClusterSoftwareFaultEventClass() _MTRSoftwareDiagnosticsClusterSoftwareFaultEventClass {
	MTRSoftwareDiagnosticsClusterSoftwareFaultEventClassOnce.Do(func() {
		MTRSoftwareDiagnosticsClusterSoftwareFaultEventClass = _MTRSoftwareDiagnosticsClusterSoftwareFaultEventClass{objc.GetClass("MTRSoftwareDiagnosticsClusterSoftwareFaultEvent")}
	})
	return MTRSoftwareDiagnosticsClusterSoftwareFaultEventClass
}

type _MTRSoftwareDiagnosticsClusterSoftwareFaultEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRSoftwareDiagnosticsClusterSoftwareFaultEvent */
// An interface definition for the [MTRSoftwareDiagnosticsClusterSoftwareFaultEvent] class.
type IMTRSoftwareDiagnosticsClusterSoftwareFaultEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRSoftwareDiagnosticsClusterSoftwareFaultEvent */
	// properties:
	FaultRecording() objc.IObject /* cross-framework: NSData */
	SetFaultRecording(value objc.IObject /* cross-framework: NSData */)
	Id() objc.IObject /* cross-framework: NSNumber */
	SetId(value objc.IObject /* cross-framework: NSNumber */)
	Name() objc.IObject /* cross-framework: NSString */
	SetName(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRSoftwareDiagnosticsClusterSoftwareFaultEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRSoftwareDiagnosticsClusterSoftwareFaultEvent */
// Alloc allocates a new instance without initialization.
func (mc _MTRSoftwareDiagnosticsClusterSoftwareFaultEventClass) Alloc() MTRSoftwareDiagnosticsClusterSoftwareFaultEvent {
	rv := objc.Send[MTRSoftwareDiagnosticsClusterSoftwareFaultEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRSoftwareDiagnosticsClusterSoftwareFaultEventClass) New() MTRSoftwareDiagnosticsClusterSoftwareFaultEvent {
	rv := objc.Send[MTRSoftwareDiagnosticsClusterSoftwareFaultEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRSoftwareDiagnosticsClusterSoftwareFaultEvent) Init() MTRSoftwareDiagnosticsClusterSoftwareFaultEvent {
	rv := objc.Send[MTRSoftwareDiagnosticsClusterSoftwareFaultEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRSoftwareDiagnosticsClusterSoftwareFaultEvent) Autorelease() MTRSoftwareDiagnosticsClusterSoftwareFaultEvent {
	rv := objc.Send[MTRSoftwareDiagnosticsClusterSoftwareFaultEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRSoftwareDiagnosticsClusterSoftwareFaultEvent creates a new MTRSoftwareDiagnosticsClusterSoftwareFaultEvent instance.
func NewMTRSoftwareDiagnosticsClusterSoftwareFaultEvent() MTRSoftwareDiagnosticsClusterSoftwareFaultEvent {
	return getMTRSoftwareDiagnosticsClusterSoftwareFaultEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRSoftwareDiagnosticsClusterSoftwareFaultEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSoftwareDiagnosticsClusterSoftwareFaultEvent
type MTRSoftwareDiagnosticsClusterSoftwareFaultEvent struct {
	objectivec.Object
}

// MTRSoftwareDiagnosticsClusterSoftwareFaultEventFrom constructs a [MTRSoftwareDiagnosticsClusterSoftwareFaultEvent] from an unsafe.Pointer.
func MTRSoftwareDiagnosticsClusterSoftwareFaultEventFrom(ptr unsafe.Pointer) MTRSoftwareDiagnosticsClusterSoftwareFaultEvent {
	return MTRSoftwareDiagnosticsClusterSoftwareFaultEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRSoftwareDiagnosticsClusterSoftwareFaultEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRSoftwareDiagnosticsClusterSoftwareFaultEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRSoftwareDiagnosticsClusterSoftwareFaultEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRSoftwareDiagnosticsClusterSoftwareFaultEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRSoftwareDiagnosticsClusterSoftwareFaultEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSoftwareDiagnosticsClusterSoftwareFaultEvent/faultRecording
func (m_ MTRSoftwareDiagnosticsClusterSoftwareFaultEvent) FaultRecording() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("faultRecording"))
	return rv
}/* debug [instance_properties/getter]: faultRecording */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSoftwareDiagnosticsClusterSoftwareFaultEvent/faultRecording
func (m_ MTRSoftwareDiagnosticsClusterSoftwareFaultEvent) SetFaultRecording(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFaultRecording:"), value)
}/* debug [instance_properties/setter]: faultRecording */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSoftwareDiagnosticsClusterSoftwareFaultEvent/id
func (m_ MTRSoftwareDiagnosticsClusterSoftwareFaultEvent) Id() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("id"))
	return rv
}/* debug [instance_properties/getter]: id */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSoftwareDiagnosticsClusterSoftwareFaultEvent/id
func (m_ MTRSoftwareDiagnosticsClusterSoftwareFaultEvent) SetId(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setId:"), value)
}/* debug [instance_properties/setter]: id */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSoftwareDiagnosticsClusterSoftwareFaultEvent/name
func (m_ MTRSoftwareDiagnosticsClusterSoftwareFaultEvent) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSoftwareDiagnosticsClusterSoftwareFaultEvent/name
func (m_ MTRSoftwareDiagnosticsClusterSoftwareFaultEvent) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setName:"), value)
}/* debug [instance_properties/setter]: name */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRSoftwareDiagnosticsClusterSoftwareFaultEvent */



