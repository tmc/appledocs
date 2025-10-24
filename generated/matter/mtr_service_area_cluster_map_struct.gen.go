// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRServiceAreaClusterMapStruct */


/* debug [class_header]: Header for MTRServiceAreaClusterMapStruct */
// The class instance for the [MTRServiceAreaClusterMapStruct] class.
var (
	MTRServiceAreaClusterMapStructClass     _MTRServiceAreaClusterMapStructClass
	MTRServiceAreaClusterMapStructClassOnce sync.Once
)

func getMTRServiceAreaClusterMapStructClass() _MTRServiceAreaClusterMapStructClass {
	MTRServiceAreaClusterMapStructClassOnce.Do(func() {
		MTRServiceAreaClusterMapStructClass = _MTRServiceAreaClusterMapStructClass{objc.GetClass("MTRServiceAreaClusterMapStruct")}
	})
	return MTRServiceAreaClusterMapStructClass
}

type _MTRServiceAreaClusterMapStructClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRServiceAreaClusterMapStruct */
// An interface definition for the [MTRServiceAreaClusterMapStruct] class.
type IMTRServiceAreaClusterMapStruct interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRServiceAreaClusterMapStruct */
	// properties:
	MapID() objc.IObject /* cross-framework: NSNumber */
	SetMapID(value objc.IObject /* cross-framework: NSNumber */)
	Name() objc.IObject /* cross-framework: NSString */
	SetName(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRServiceAreaClusterMapStruct */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRServiceAreaClusterMapStruct */
// Alloc allocates a new instance without initialization.
func (mc _MTRServiceAreaClusterMapStructClass) Alloc() MTRServiceAreaClusterMapStruct {
	rv := objc.Send[MTRServiceAreaClusterMapStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRServiceAreaClusterMapStructClass) New() MTRServiceAreaClusterMapStruct {
	rv := objc.Send[MTRServiceAreaClusterMapStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRServiceAreaClusterMapStruct) Init() MTRServiceAreaClusterMapStruct {
	rv := objc.Send[MTRServiceAreaClusterMapStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRServiceAreaClusterMapStruct) Autorelease() MTRServiceAreaClusterMapStruct {
	rv := objc.Send[MTRServiceAreaClusterMapStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRServiceAreaClusterMapStruct creates a new MTRServiceAreaClusterMapStruct instance.
func NewMTRServiceAreaClusterMapStruct() MTRServiceAreaClusterMapStruct {
	return getMTRServiceAreaClusterMapStructClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRServiceAreaClusterMapStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterMapStruct
type MTRServiceAreaClusterMapStruct struct {
	objectivec.Object
}

// MTRServiceAreaClusterMapStructFrom constructs a [MTRServiceAreaClusterMapStruct] from an unsafe.Pointer.
func MTRServiceAreaClusterMapStructFrom(ptr unsafe.Pointer) MTRServiceAreaClusterMapStruct {
	return MTRServiceAreaClusterMapStruct{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRServiceAreaClusterMapStruct *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRServiceAreaClusterMapStruct */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRServiceAreaClusterMapStruct */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRServiceAreaClusterMapStruct */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRServiceAreaClusterMapStruct */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterMapStruct/mapID
func (m_ MTRServiceAreaClusterMapStruct) MapID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("mapID"))
	return rv
}/* debug [instance_properties/getter]: mapID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterMapStruct/mapID
func (m_ MTRServiceAreaClusterMapStruct) SetMapID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMapID:"), value)
}/* debug [instance_properties/setter]: mapID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrserviceareaclustermapstruct/name
func (m_ MTRServiceAreaClusterMapStruct) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrserviceareaclustermapstruct/name
func (m_ MTRServiceAreaClusterMapStruct) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setName:"), value)
}/* debug [instance_properties/setter]: name */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRServiceAreaClusterMapStruct */



