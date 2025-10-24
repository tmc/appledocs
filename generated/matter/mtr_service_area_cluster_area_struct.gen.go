// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRServiceAreaClusterAreaStruct */


/* debug [class_header]: Header for MTRServiceAreaClusterAreaStruct */
// The class instance for the [MTRServiceAreaClusterAreaStruct] class.
var (
	MTRServiceAreaClusterAreaStructClass     _MTRServiceAreaClusterAreaStructClass
	MTRServiceAreaClusterAreaStructClassOnce sync.Once
)

func getMTRServiceAreaClusterAreaStructClass() _MTRServiceAreaClusterAreaStructClass {
	MTRServiceAreaClusterAreaStructClassOnce.Do(func() {
		MTRServiceAreaClusterAreaStructClass = _MTRServiceAreaClusterAreaStructClass{objc.GetClass("MTRServiceAreaClusterAreaStruct")}
	})
	return MTRServiceAreaClusterAreaStructClass
}

type _MTRServiceAreaClusterAreaStructClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRServiceAreaClusterAreaStruct */
// An interface definition for the [MTRServiceAreaClusterAreaStruct] class.
type IMTRServiceAreaClusterAreaStruct interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRServiceAreaClusterAreaStruct */
	// properties:
	AreaID() objc.IObject /* cross-framework: NSNumber */
	SetAreaID(value objc.IObject /* cross-framework: NSNumber */)
	AreaInfo() IMTRServiceAreaClusterAreaInfoStruct
	SetAreaInfo(value IMTRServiceAreaClusterAreaInfoStruct)
	MapID() objc.IObject /* cross-framework: NSNumber */
	SetMapID(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRServiceAreaClusterAreaStruct */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRServiceAreaClusterAreaStruct */
// Alloc allocates a new instance without initialization.
func (mc _MTRServiceAreaClusterAreaStructClass) Alloc() MTRServiceAreaClusterAreaStruct {
	rv := objc.Send[MTRServiceAreaClusterAreaStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRServiceAreaClusterAreaStructClass) New() MTRServiceAreaClusterAreaStruct {
	rv := objc.Send[MTRServiceAreaClusterAreaStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRServiceAreaClusterAreaStruct) Init() MTRServiceAreaClusterAreaStruct {
	rv := objc.Send[MTRServiceAreaClusterAreaStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRServiceAreaClusterAreaStruct) Autorelease() MTRServiceAreaClusterAreaStruct {
	rv := objc.Send[MTRServiceAreaClusterAreaStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRServiceAreaClusterAreaStruct creates a new MTRServiceAreaClusterAreaStruct instance.
func NewMTRServiceAreaClusterAreaStruct() MTRServiceAreaClusterAreaStruct {
	return getMTRServiceAreaClusterAreaStructClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRServiceAreaClusterAreaStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterAreaStruct
type MTRServiceAreaClusterAreaStruct struct {
	objectivec.Object
}

// MTRServiceAreaClusterAreaStructFrom constructs a [MTRServiceAreaClusterAreaStruct] from an unsafe.Pointer.
func MTRServiceAreaClusterAreaStructFrom(ptr unsafe.Pointer) MTRServiceAreaClusterAreaStruct {
	return MTRServiceAreaClusterAreaStruct{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRServiceAreaClusterAreaStruct *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRServiceAreaClusterAreaStruct */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRServiceAreaClusterAreaStruct */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRServiceAreaClusterAreaStruct */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRServiceAreaClusterAreaStruct */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterAreaStruct/areaID
func (m_ MTRServiceAreaClusterAreaStruct) AreaID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("areaID"))
	return rv
}/* debug [instance_properties/getter]: areaID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRServiceAreaClusterAreaStruct/areaID
func (m_ MTRServiceAreaClusterAreaStruct) SetAreaID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAreaID:"), value)
}/* debug [instance_properties/setter]: areaID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrserviceareaclusterareastruct/areainfo
func (m_ MTRServiceAreaClusterAreaStruct) AreaInfo() IMTRServiceAreaClusterAreaInfoStruct {
	rv := objc.Send[MTRServiceAreaClusterAreaInfoStruct](m_.ID, objc.Sel("areaInfo"))
	return rv
}/* debug [instance_properties/getter]: areaInfo */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrserviceareaclusterareastruct/areainfo
func (m_ MTRServiceAreaClusterAreaStruct) SetAreaInfo(value IMTRServiceAreaClusterAreaInfoStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAreaInfo:"), value)
}/* debug [instance_properties/setter]: areaInfo */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrserviceareaclusterareastruct/mapid
func (m_ MTRServiceAreaClusterAreaStruct) MapID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("mapID"))
	return rv
}/* debug [instance_properties/getter]: mapID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrserviceareaclusterareastruct/mapid
func (m_ MTRServiceAreaClusterAreaStruct) SetMapID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMapID:"), value)
}/* debug [instance_properties/setter]: mapID */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRServiceAreaClusterAreaStruct */



