// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTROvenCavityOperationalStateClusterErrorStateStruct */


/* debug [class_header]: Header for MTROvenCavityOperationalStateClusterErrorStateStruct */
// The class instance for the [MTROvenCavityOperationalStateClusterErrorStateStruct] class.
var (
	MTROvenCavityOperationalStateClusterErrorStateStructClass     _MTROvenCavityOperationalStateClusterErrorStateStructClass
	MTROvenCavityOperationalStateClusterErrorStateStructClassOnce sync.Once
)

func getMTROvenCavityOperationalStateClusterErrorStateStructClass() _MTROvenCavityOperationalStateClusterErrorStateStructClass {
	MTROvenCavityOperationalStateClusterErrorStateStructClassOnce.Do(func() {
		MTROvenCavityOperationalStateClusterErrorStateStructClass = _MTROvenCavityOperationalStateClusterErrorStateStructClass{objc.GetClass("MTROvenCavityOperationalStateClusterErrorStateStruct")}
	})
	return MTROvenCavityOperationalStateClusterErrorStateStructClass
}

type _MTROvenCavityOperationalStateClusterErrorStateStructClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTROvenCavityOperationalStateClusterErrorStateStruct */
// An interface definition for the [MTROvenCavityOperationalStateClusterErrorStateStruct] class.
type IMTROvenCavityOperationalStateClusterErrorStateStruct interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTROvenCavityOperationalStateClusterErrorStateStruct */
	// properties:
	ErrorStateDetails() objc.IObject /* cross-framework: NSString */
	SetErrorStateDetails(value objc.IObject /* cross-framework: NSString */)
	ErrorStateID() objc.IObject /* cross-framework: NSNumber */
	SetErrorStateID(value objc.IObject /* cross-framework: NSNumber */)
	ErrorStateLabel() objc.IObject /* cross-framework: NSString */
	SetErrorStateLabel(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTROvenCavityOperationalStateClusterErrorStateStruct */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTROvenCavityOperationalStateClusterErrorStateStruct */
// Alloc allocates a new instance without initialization.
func (mc _MTROvenCavityOperationalStateClusterErrorStateStructClass) Alloc() MTROvenCavityOperationalStateClusterErrorStateStruct {
	rv := objc.Send[MTROvenCavityOperationalStateClusterErrorStateStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTROvenCavityOperationalStateClusterErrorStateStructClass) New() MTROvenCavityOperationalStateClusterErrorStateStruct {
	rv := objc.Send[MTROvenCavityOperationalStateClusterErrorStateStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROvenCavityOperationalStateClusterErrorStateStruct) Init() MTROvenCavityOperationalStateClusterErrorStateStruct {
	rv := objc.Send[MTROvenCavityOperationalStateClusterErrorStateStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROvenCavityOperationalStateClusterErrorStateStruct) Autorelease() MTROvenCavityOperationalStateClusterErrorStateStruct {
	rv := objc.Send[MTROvenCavityOperationalStateClusterErrorStateStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROvenCavityOperationalStateClusterErrorStateStruct creates a new MTROvenCavityOperationalStateClusterErrorStateStruct instance.
func NewMTROvenCavityOperationalStateClusterErrorStateStruct() MTROvenCavityOperationalStateClusterErrorStateStruct {
	return getMTROvenCavityOperationalStateClusterErrorStateStructClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTROvenCavityOperationalStateClusterErrorStateStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenCavityOperationalStateClusterErrorStateStruct
type MTROvenCavityOperationalStateClusterErrorStateStruct struct {
	objectivec.Object
}

// MTROvenCavityOperationalStateClusterErrorStateStructFrom constructs a [MTROvenCavityOperationalStateClusterErrorStateStruct] from an unsafe.Pointer.
func MTROvenCavityOperationalStateClusterErrorStateStructFrom(ptr unsafe.Pointer) MTROvenCavityOperationalStateClusterErrorStateStruct {
	return MTROvenCavityOperationalStateClusterErrorStateStruct{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTROvenCavityOperationalStateClusterErrorStateStruct *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTROvenCavityOperationalStateClusterErrorStateStruct */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTROvenCavityOperationalStateClusterErrorStateStruct */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTROvenCavityOperationalStateClusterErrorStateStruct */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTROvenCavityOperationalStateClusterErrorStateStruct */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenCavityOperationalStateClusterErrorStateStruct/errorStateDetails
func (m_ MTROvenCavityOperationalStateClusterErrorStateStruct) ErrorStateDetails() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("errorStateDetails"))
	return rv
}/* debug [instance_properties/getter]: errorStateDetails */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROvenCavityOperationalStateClusterErrorStateStruct/errorStateDetails
func (m_ MTROvenCavityOperationalStateClusterErrorStateStruct) SetErrorStateDetails(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setErrorStateDetails:"), value)
}/* debug [instance_properties/setter]: errorStateDetails */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrovencavityoperationalstateclustererrorstatestruct/errorstateid
func (m_ MTROvenCavityOperationalStateClusterErrorStateStruct) ErrorStateID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("errorStateID"))
	return rv
}/* debug [instance_properties/getter]: errorStateID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrovencavityoperationalstateclustererrorstatestruct/errorstateid
func (m_ MTROvenCavityOperationalStateClusterErrorStateStruct) SetErrorStateID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setErrorStateID:"), value)
}/* debug [instance_properties/setter]: errorStateID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrovencavityoperationalstateclustererrorstatestruct/errorstatelabel
func (m_ MTROvenCavityOperationalStateClusterErrorStateStruct) ErrorStateLabel() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("errorStateLabel"))
	return rv
}/* debug [instance_properties/getter]: errorStateLabel */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrovencavityoperationalstateclustererrorstatestruct/errorstatelabel
func (m_ MTROvenCavityOperationalStateClusterErrorStateStruct) SetErrorStateLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setErrorStateLabel:"), value)
}/* debug [instance_properties/setter]: errorStateLabel */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTROvenCavityOperationalStateClusterErrorStateStruct */



