// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRServiceAreaClusterProgressStruct */


/* debug [class_header]: Header for MTRServiceAreaClusterProgressStruct */
// The class instance for the [MTRServiceAreaClusterProgressStruct] class.
var (
	MTRServiceAreaClusterProgressStructClass     _MTRServiceAreaClusterProgressStructClass
	MTRServiceAreaClusterProgressStructClassOnce sync.Once
)

func getMTRServiceAreaClusterProgressStructClass() _MTRServiceAreaClusterProgressStructClass {
	MTRServiceAreaClusterProgressStructClassOnce.Do(func() {
		MTRServiceAreaClusterProgressStructClass = _MTRServiceAreaClusterProgressStructClass{objc.GetClass("MTRServiceAreaClusterProgressStruct")}
	})
	return MTRServiceAreaClusterProgressStructClass
}

type _MTRServiceAreaClusterProgressStructClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRServiceAreaClusterProgressStruct */
// An interface definition for the [MTRServiceAreaClusterProgressStruct] class.
type IMTRServiceAreaClusterProgressStruct interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRServiceAreaClusterProgressStruct */
	// properties:
	AreaID() objc.IObject /* cross-framework: NSNumber */
	SetAreaID(value objc.IObject /* cross-framework: NSNumber */)
	Status() objc.IObject /* cross-framework: NSNumber */
	SetStatus(value objc.IObject /* cross-framework: NSNumber */)
	TotalOperationalTime() objc.IObject /* cross-framework: NSNumber */
	SetTotalOperationalTime(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRServiceAreaClusterProgressStruct */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRServiceAreaClusterProgressStruct */
// Alloc allocates a new instance without initialization.
func (mc _MTRServiceAreaClusterProgressStructClass) Alloc() MTRServiceAreaClusterProgressStruct {
	rv := objc.Send[MTRServiceAreaClusterProgressStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRServiceAreaClusterProgressStructClass) New() MTRServiceAreaClusterProgressStruct {
	rv := objc.Send[MTRServiceAreaClusterProgressStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRServiceAreaClusterProgressStruct) Init() MTRServiceAreaClusterProgressStruct {
	rv := objc.Send[MTRServiceAreaClusterProgressStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRServiceAreaClusterProgressStruct) Autorelease() MTRServiceAreaClusterProgressStruct {
	rv := objc.Send[MTRServiceAreaClusterProgressStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRServiceAreaClusterProgressStruct creates a new MTRServiceAreaClusterProgressStruct instance.
func NewMTRServiceAreaClusterProgressStruct() MTRServiceAreaClusterProgressStruct {
	return getMTRServiceAreaClusterProgressStructClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRServiceAreaClusterProgressStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterProgressStruct
type MTRServiceAreaClusterProgressStruct struct {
	objectivec.Object
}

// MTRServiceAreaClusterProgressStructFrom constructs a [MTRServiceAreaClusterProgressStruct] from an unsafe.Pointer.
func MTRServiceAreaClusterProgressStructFrom(ptr unsafe.Pointer) MTRServiceAreaClusterProgressStruct {
	return MTRServiceAreaClusterProgressStruct{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRServiceAreaClusterProgressStruct *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRServiceAreaClusterProgressStruct */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRServiceAreaClusterProgressStruct */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRServiceAreaClusterProgressStruct */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRServiceAreaClusterProgressStruct */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterProgressStruct/areaID
func (m_ MTRServiceAreaClusterProgressStruct) AreaID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("areaID"))
	return rv
}/* debug [instance_properties/getter]: areaID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterProgressStruct/areaID
func (m_ MTRServiceAreaClusterProgressStruct) SetAreaID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAreaID:"), value)
}/* debug [instance_properties/setter]: areaID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrserviceareaclusterprogressstruct/status
func (m_ MTRServiceAreaClusterProgressStruct) Status() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("status"))
	return rv
}/* debug [instance_properties/getter]: status */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrserviceareaclusterprogressstruct/status
func (m_ MTRServiceAreaClusterProgressStruct) SetStatus(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStatus:"), value)
}/* debug [instance_properties/setter]: status */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrserviceareaclusterprogressstruct/totaloperationaltime
func (m_ MTRServiceAreaClusterProgressStruct) TotalOperationalTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("totalOperationalTime"))
	return rv
}/* debug [instance_properties/getter]: totalOperationalTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrserviceareaclusterprogressstruct/totaloperationaltime
func (m_ MTRServiceAreaClusterProgressStruct) SetTotalOperationalTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTotalOperationalTime:"), value)
}/* debug [instance_properties/setter]: totalOperationalTime */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRServiceAreaClusterProgressStruct */



