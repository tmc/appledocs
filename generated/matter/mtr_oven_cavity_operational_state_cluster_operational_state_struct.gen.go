// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTROvenCavityOperationalStateClusterOperationalStateStruct */


/* debug [class_header]: Header for MTROvenCavityOperationalStateClusterOperationalStateStruct */
// The class instance for the [MTROvenCavityOperationalStateClusterOperationalStateStruct] class.
var (
	MTROvenCavityOperationalStateClusterOperationalStateStructClass     _MTROvenCavityOperationalStateClusterOperationalStateStructClass
	MTROvenCavityOperationalStateClusterOperationalStateStructClassOnce sync.Once
)

func getMTROvenCavityOperationalStateClusterOperationalStateStructClass() _MTROvenCavityOperationalStateClusterOperationalStateStructClass {
	MTROvenCavityOperationalStateClusterOperationalStateStructClassOnce.Do(func() {
		MTROvenCavityOperationalStateClusterOperationalStateStructClass = _MTROvenCavityOperationalStateClusterOperationalStateStructClass{objc.GetClass("MTROvenCavityOperationalStateClusterOperationalStateStruct")}
	})
	return MTROvenCavityOperationalStateClusterOperationalStateStructClass
}

type _MTROvenCavityOperationalStateClusterOperationalStateStructClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTROvenCavityOperationalStateClusterOperationalStateStruct */
// An interface definition for the [MTROvenCavityOperationalStateClusterOperationalStateStruct] class.
type IMTROvenCavityOperationalStateClusterOperationalStateStruct interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTROvenCavityOperationalStateClusterOperationalStateStruct */
	// properties:
	OperationalStateID() objc.IObject /* cross-framework: NSNumber */
	SetOperationalStateID(value objc.IObject /* cross-framework: NSNumber */)
	OperationalStateLabel() objc.IObject /* cross-framework: NSString */
	SetOperationalStateLabel(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTROvenCavityOperationalStateClusterOperationalStateStruct */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTROvenCavityOperationalStateClusterOperationalStateStruct */
// Alloc allocates a new instance without initialization.
func (mc _MTROvenCavityOperationalStateClusterOperationalStateStructClass) Alloc() MTROvenCavityOperationalStateClusterOperationalStateStruct {
	rv := objc.Send[MTROvenCavityOperationalStateClusterOperationalStateStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTROvenCavityOperationalStateClusterOperationalStateStructClass) New() MTROvenCavityOperationalStateClusterOperationalStateStruct {
	rv := objc.Send[MTROvenCavityOperationalStateClusterOperationalStateStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROvenCavityOperationalStateClusterOperationalStateStruct) Init() MTROvenCavityOperationalStateClusterOperationalStateStruct {
	rv := objc.Send[MTROvenCavityOperationalStateClusterOperationalStateStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROvenCavityOperationalStateClusterOperationalStateStruct) Autorelease() MTROvenCavityOperationalStateClusterOperationalStateStruct {
	rv := objc.Send[MTROvenCavityOperationalStateClusterOperationalStateStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROvenCavityOperationalStateClusterOperationalStateStruct creates a new MTROvenCavityOperationalStateClusterOperationalStateStruct instance.
func NewMTROvenCavityOperationalStateClusterOperationalStateStruct() MTROvenCavityOperationalStateClusterOperationalStateStruct {
	return getMTROvenCavityOperationalStateClusterOperationalStateStructClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTROvenCavityOperationalStateClusterOperationalStateStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenCavityOperationalStateClusterOperationalStateStruct
type MTROvenCavityOperationalStateClusterOperationalStateStruct struct {
	objectivec.Object
}

// MTROvenCavityOperationalStateClusterOperationalStateStructFrom constructs a [MTROvenCavityOperationalStateClusterOperationalStateStruct] from an unsafe.Pointer.
func MTROvenCavityOperationalStateClusterOperationalStateStructFrom(ptr unsafe.Pointer) MTROvenCavityOperationalStateClusterOperationalStateStruct {
	return MTROvenCavityOperationalStateClusterOperationalStateStruct{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTROvenCavityOperationalStateClusterOperationalStateStruct *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTROvenCavityOperationalStateClusterOperationalStateStruct */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTROvenCavityOperationalStateClusterOperationalStateStruct */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTROvenCavityOperationalStateClusterOperationalStateStruct */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTROvenCavityOperationalStateClusterOperationalStateStruct */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenCavityOperationalStateClusterOperationalStateStruct/operationalStateID
func (m_ MTROvenCavityOperationalStateClusterOperationalStateStruct) OperationalStateID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("operationalStateID"))
	return rv
}/* debug [instance_properties/getter]: operationalStateID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenCavityOperationalStateClusterOperationalStateStruct/operationalStateID
func (m_ MTROvenCavityOperationalStateClusterOperationalStateStruct) SetOperationalStateID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOperationalStateID:"), value)
}/* debug [instance_properties/setter]: operationalStateID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrovencavityoperationalstateclusteroperationalstatestruct/operationalstatelabel
func (m_ MTROvenCavityOperationalStateClusterOperationalStateStruct) OperationalStateLabel() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("operationalStateLabel"))
	return rv
}/* debug [instance_properties/getter]: operationalStateLabel */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrovencavityoperationalstateclusteroperationalstatestruct/operationalstatelabel
func (m_ MTROvenCavityOperationalStateClusterOperationalStateStruct) SetOperationalStateLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOperationalStateLabel:"), value)
}/* debug [instance_properties/setter]: operationalStateLabel */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTROvenCavityOperationalStateClusterOperationalStateStruct */



