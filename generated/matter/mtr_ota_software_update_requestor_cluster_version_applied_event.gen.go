// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTROtaSoftwareUpdateRequestorClusterVersionAppliedEvent */


/* debug [class_header]: Header for MTROtaSoftwareUpdateRequestorClusterVersionAppliedEvent */
// The class instance for the [MTROtaSoftwareUpdateRequestorClusterVersionAppliedEvent] class.
var (
	MTROtaSoftwareUpdateRequestorClusterVersionAppliedEventClass     _MTROtaSoftwareUpdateRequestorClusterVersionAppliedEventClass
	MTROtaSoftwareUpdateRequestorClusterVersionAppliedEventClassOnce sync.Once
)

func getMTROtaSoftwareUpdateRequestorClusterVersionAppliedEventClass() _MTROtaSoftwareUpdateRequestorClusterVersionAppliedEventClass {
	MTROtaSoftwareUpdateRequestorClusterVersionAppliedEventClassOnce.Do(func() {
		MTROtaSoftwareUpdateRequestorClusterVersionAppliedEventClass = _MTROtaSoftwareUpdateRequestorClusterVersionAppliedEventClass{objc.GetClass("MTROtaSoftwareUpdateRequestorClusterVersionAppliedEvent")}
	})
	return MTROtaSoftwareUpdateRequestorClusterVersionAppliedEventClass
}

type _MTROtaSoftwareUpdateRequestorClusterVersionAppliedEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTROtaSoftwareUpdateRequestorClusterVersionAppliedEvent */
// An interface definition for the [MTROtaSoftwareUpdateRequestorClusterVersionAppliedEvent] class.
type IMTROtaSoftwareUpdateRequestorClusterVersionAppliedEvent interface {
	IMTROTASoftwareUpdateRequestorClusterVersionAppliedEvent
	
/* debug [class_interface_properties]: Properties for MTROtaSoftwareUpdateRequestorClusterVersionAppliedEvent */
	// properties:
	ProductID() objc.IObject /* cross-framework: NSNumber */
	SetProductID(value objc.IObject /* cross-framework: NSNumber */)
	SoftwareVersion() objc.IObject /* cross-framework: NSNumber */
	SetSoftwareVersion(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTROtaSoftwareUpdateRequestorClusterVersionAppliedEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTROtaSoftwareUpdateRequestorClusterVersionAppliedEvent */
// Alloc allocates a new instance without initialization.
func (mc _MTROtaSoftwareUpdateRequestorClusterVersionAppliedEventClass) Alloc() MTROtaSoftwareUpdateRequestorClusterVersionAppliedEvent {
	rv := objc.Send[MTROtaSoftwareUpdateRequestorClusterVersionAppliedEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTROtaSoftwareUpdateRequestorClusterVersionAppliedEventClass) New() MTROtaSoftwareUpdateRequestorClusterVersionAppliedEvent {
	rv := objc.Send[MTROtaSoftwareUpdateRequestorClusterVersionAppliedEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROtaSoftwareUpdateRequestorClusterVersionAppliedEvent) Init() MTROtaSoftwareUpdateRequestorClusterVersionAppliedEvent {
	rv := objc.Send[MTROtaSoftwareUpdateRequestorClusterVersionAppliedEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROtaSoftwareUpdateRequestorClusterVersionAppliedEvent) Autorelease() MTROtaSoftwareUpdateRequestorClusterVersionAppliedEvent {
	rv := objc.Send[MTROtaSoftwareUpdateRequestorClusterVersionAppliedEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROtaSoftwareUpdateRequestorClusterVersionAppliedEvent creates a new MTROtaSoftwareUpdateRequestorClusterVersionAppliedEvent instance.
func NewMTROtaSoftwareUpdateRequestorClusterVersionAppliedEvent() MTROtaSoftwareUpdateRequestorClusterVersionAppliedEvent {
	return getMTROtaSoftwareUpdateRequestorClusterVersionAppliedEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTROtaSoftwareUpdateRequestorClusterVersionAppliedEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateRequestorClusterVersionAppliedEvent-970fn
type MTROtaSoftwareUpdateRequestorClusterVersionAppliedEvent struct {
	MTROTASoftwareUpdateRequestorClusterVersionAppliedEvent
}

// MTROtaSoftwareUpdateRequestorClusterVersionAppliedEventFrom constructs a [MTROtaSoftwareUpdateRequestorClusterVersionAppliedEvent] from an unsafe.Pointer.
func MTROtaSoftwareUpdateRequestorClusterVersionAppliedEventFrom(ptr unsafe.Pointer) MTROtaSoftwareUpdateRequestorClusterVersionAppliedEvent {
	return MTROtaSoftwareUpdateRequestorClusterVersionAppliedEvent{
		MTROTASoftwareUpdateRequestorClusterVersionAppliedEvent: MTROTASoftwareUpdateRequestorClusterVersionAppliedEventFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTROtaSoftwareUpdateRequestorClusterVersionAppliedEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTROtaSoftwareUpdateRequestorClusterVersionAppliedEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTROtaSoftwareUpdateRequestorClusterVersionAppliedEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTROtaSoftwareUpdateRequestorClusterVersionAppliedEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTROtaSoftwareUpdateRequestorClusterVersionAppliedEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateRequestorClusterVersionAppliedEvent-970fn/productID
func (m_ MTROtaSoftwareUpdateRequestorClusterVersionAppliedEvent) ProductID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("productID"))
	return rv
}/* debug [instance_properties/getter]: productID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateRequestorClusterVersionAppliedEvent-970fn/productID
func (m_ MTROtaSoftwareUpdateRequestorClusterVersionAppliedEvent) SetProductID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProductID:"), value)
}/* debug [instance_properties/setter]: productID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateRequestorClusterVersionAppliedEvent-970fn/softwareVersion
func (m_ MTROtaSoftwareUpdateRequestorClusterVersionAppliedEvent) SoftwareVersion() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("softwareVersion"))
	return rv
}/* debug [instance_properties/getter]: softwareVersion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROtaSoftwareUpdateRequestorClusterVersionAppliedEvent-970fn/softwareVersion
func (m_ MTROtaSoftwareUpdateRequestorClusterVersionAppliedEvent) SetSoftwareVersion(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSoftwareVersion:"), value)
}/* debug [instance_properties/setter]: softwareVersion */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTROtaSoftwareUpdateRequestorClusterVersionAppliedEvent */



