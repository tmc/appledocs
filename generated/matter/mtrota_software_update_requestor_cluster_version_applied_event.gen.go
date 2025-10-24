// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTROTASoftwareUpdateRequestorClusterVersionAppliedEvent */


/* debug [class_header]: Header for MTROTASoftwareUpdateRequestorClusterVersionAppliedEvent */
// The class instance for the [MTROTASoftwareUpdateRequestorClusterVersionAppliedEvent] class.
var (
	MTROTASoftwareUpdateRequestorClusterVersionAppliedEventClass     _MTROTASoftwareUpdateRequestorClusterVersionAppliedEventClass
	MTROTASoftwareUpdateRequestorClusterVersionAppliedEventClassOnce sync.Once
)

func getMTROTASoftwareUpdateRequestorClusterVersionAppliedEventClass() _MTROTASoftwareUpdateRequestorClusterVersionAppliedEventClass {
	MTROTASoftwareUpdateRequestorClusterVersionAppliedEventClassOnce.Do(func() {
		MTROTASoftwareUpdateRequestorClusterVersionAppliedEventClass = _MTROTASoftwareUpdateRequestorClusterVersionAppliedEventClass{objc.GetClass("MTROTASoftwareUpdateRequestorClusterVersionAppliedEvent")}
	})
	return MTROTASoftwareUpdateRequestorClusterVersionAppliedEventClass
}

type _MTROTASoftwareUpdateRequestorClusterVersionAppliedEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTROTASoftwareUpdateRequestorClusterVersionAppliedEvent */
// An interface definition for the [MTROTASoftwareUpdateRequestorClusterVersionAppliedEvent] class.
type IMTROTASoftwareUpdateRequestorClusterVersionAppliedEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTROTASoftwareUpdateRequestorClusterVersionAppliedEvent */
	// properties:
	ProductID() objc.IObject /* cross-framework: NSNumber */
	SetProductID(value objc.IObject /* cross-framework: NSNumber */)
	SoftwareVersion() objc.IObject /* cross-framework: NSNumber */
	SetSoftwareVersion(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTROTASoftwareUpdateRequestorClusterVersionAppliedEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTROTASoftwareUpdateRequestorClusterVersionAppliedEvent */
// Alloc allocates a new instance without initialization.
func (mc _MTROTASoftwareUpdateRequestorClusterVersionAppliedEventClass) Alloc() MTROTASoftwareUpdateRequestorClusterVersionAppliedEvent {
	rv := objc.Send[MTROTASoftwareUpdateRequestorClusterVersionAppliedEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTROTASoftwareUpdateRequestorClusterVersionAppliedEventClass) New() MTROTASoftwareUpdateRequestorClusterVersionAppliedEvent {
	rv := objc.Send[MTROTASoftwareUpdateRequestorClusterVersionAppliedEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROTASoftwareUpdateRequestorClusterVersionAppliedEvent) Init() MTROTASoftwareUpdateRequestorClusterVersionAppliedEvent {
	rv := objc.Send[MTROTASoftwareUpdateRequestorClusterVersionAppliedEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROTASoftwareUpdateRequestorClusterVersionAppliedEvent) Autorelease() MTROTASoftwareUpdateRequestorClusterVersionAppliedEvent {
	rv := objc.Send[MTROTASoftwareUpdateRequestorClusterVersionAppliedEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROTASoftwareUpdateRequestorClusterVersionAppliedEvent creates a new MTROTASoftwareUpdateRequestorClusterVersionAppliedEvent instance.
func NewMTROTASoftwareUpdateRequestorClusterVersionAppliedEvent() MTROTASoftwareUpdateRequestorClusterVersionAppliedEvent {
	return getMTROTASoftwareUpdateRequestorClusterVersionAppliedEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTROTASoftwareUpdateRequestorClusterVersionAppliedEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateRequestorClusterVersionAppliedEvent-94prr
type MTROTASoftwareUpdateRequestorClusterVersionAppliedEvent struct {
	objectivec.Object
}

// MTROTASoftwareUpdateRequestorClusterVersionAppliedEventFrom constructs a [MTROTASoftwareUpdateRequestorClusterVersionAppliedEvent] from an unsafe.Pointer.
func MTROTASoftwareUpdateRequestorClusterVersionAppliedEventFrom(ptr unsafe.Pointer) MTROTASoftwareUpdateRequestorClusterVersionAppliedEvent {
	return MTROTASoftwareUpdateRequestorClusterVersionAppliedEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTROTASoftwareUpdateRequestorClusterVersionAppliedEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTROTASoftwareUpdateRequestorClusterVersionAppliedEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTROTASoftwareUpdateRequestorClusterVersionAppliedEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTROTASoftwareUpdateRequestorClusterVersionAppliedEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTROTASoftwareUpdateRequestorClusterVersionAppliedEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateRequestorClusterVersionAppliedEvent-94prr/productID
func (m_ MTROTASoftwareUpdateRequestorClusterVersionAppliedEvent) ProductID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("productID"))
	return rv
}/* debug [instance_properties/getter]: productID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateRequestorClusterVersionAppliedEvent-94prr/productID
func (m_ MTROTASoftwareUpdateRequestorClusterVersionAppliedEvent) SetProductID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProductID:"), value)
}/* debug [instance_properties/setter]: productID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateRequestorClusterVersionAppliedEvent-94prr/softwareVersion
func (m_ MTROTASoftwareUpdateRequestorClusterVersionAppliedEvent) SoftwareVersion() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("softwareVersion"))
	return rv
}/* debug [instance_properties/getter]: softwareVersion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTASoftwareUpdateRequestorClusterVersionAppliedEvent-94prr/softwareVersion
func (m_ MTROTASoftwareUpdateRequestorClusterVersionAppliedEvent) SetSoftwareVersion(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSoftwareVersion:"), value)
}/* debug [instance_properties/setter]: softwareVersion */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTROTASoftwareUpdateRequestorClusterVersionAppliedEvent */



