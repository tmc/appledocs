// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRDataTypeLocationDescriptorStruct */


/* debug [class_header]: Header for MTRDataTypeLocationDescriptorStruct */
// The class instance for the [MTRDataTypeLocationDescriptorStruct] class.
var (
	MTRDataTypeLocationDescriptorStructClass     _MTRDataTypeLocationDescriptorStructClass
	MTRDataTypeLocationDescriptorStructClassOnce sync.Once
)

func getMTRDataTypeLocationDescriptorStructClass() _MTRDataTypeLocationDescriptorStructClass {
	MTRDataTypeLocationDescriptorStructClassOnce.Do(func() {
		MTRDataTypeLocationDescriptorStructClass = _MTRDataTypeLocationDescriptorStructClass{objc.GetClass("MTRDataTypeLocationDescriptorStruct")}
	})
	return MTRDataTypeLocationDescriptorStructClass
}

type _MTRDataTypeLocationDescriptorStructClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRDataTypeLocationDescriptorStruct */
// An interface definition for the [MTRDataTypeLocationDescriptorStruct] class.
type IMTRDataTypeLocationDescriptorStruct interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRDataTypeLocationDescriptorStruct */
	// properties:
	AreaType() objc.IObject /* cross-framework: NSNumber */
	SetAreaType(value objc.IObject /* cross-framework: NSNumber */)
	FloorNumber() objc.IObject /* cross-framework: NSNumber */
	SetFloorNumber(value objc.IObject /* cross-framework: NSNumber */)
	LocationName() objc.IObject /* cross-framework: NSString */
	SetLocationName(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRDataTypeLocationDescriptorStruct */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRDataTypeLocationDescriptorStruct */
// Alloc allocates a new instance without initialization.
func (mc _MTRDataTypeLocationDescriptorStructClass) Alloc() MTRDataTypeLocationDescriptorStruct {
	rv := objc.Send[MTRDataTypeLocationDescriptorStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRDataTypeLocationDescriptorStructClass) New() MTRDataTypeLocationDescriptorStruct {
	rv := objc.Send[MTRDataTypeLocationDescriptorStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDataTypeLocationDescriptorStruct) Init() MTRDataTypeLocationDescriptorStruct {
	rv := objc.Send[MTRDataTypeLocationDescriptorStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDataTypeLocationDescriptorStruct) Autorelease() MTRDataTypeLocationDescriptorStruct {
	rv := objc.Send[MTRDataTypeLocationDescriptorStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDataTypeLocationDescriptorStruct creates a new MTRDataTypeLocationDescriptorStruct instance.
func NewMTRDataTypeLocationDescriptorStruct() MTRDataTypeLocationDescriptorStruct {
	return getMTRDataTypeLocationDescriptorStructClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRDataTypeLocationDescriptorStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDataTypeLocationDescriptorStruct
type MTRDataTypeLocationDescriptorStruct struct {
	objectivec.Object
}

// MTRDataTypeLocationDescriptorStructFrom constructs a [MTRDataTypeLocationDescriptorStruct] from an unsafe.Pointer.
func MTRDataTypeLocationDescriptorStructFrom(ptr unsafe.Pointer) MTRDataTypeLocationDescriptorStruct {
	return MTRDataTypeLocationDescriptorStruct{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRDataTypeLocationDescriptorStruct *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRDataTypeLocationDescriptorStruct */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRDataTypeLocationDescriptorStruct */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRDataTypeLocationDescriptorStruct */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRDataTypeLocationDescriptorStruct */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDataTypeLocationDescriptorStruct/areaType
func (m_ MTRDataTypeLocationDescriptorStruct) AreaType() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("areaType"))
	return rv
}/* debug [instance_properties/getter]: areaType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDataTypeLocationDescriptorStruct/areaType
func (m_ MTRDataTypeLocationDescriptorStruct) SetAreaType(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAreaType:"), value)
}/* debug [instance_properties/setter]: areaType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdatatypelocationdescriptorstruct/floornumber
func (m_ MTRDataTypeLocationDescriptorStruct) FloorNumber() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("floorNumber"))
	return rv
}/* debug [instance_properties/getter]: floorNumber */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdatatypelocationdescriptorstruct/floornumber
func (m_ MTRDataTypeLocationDescriptorStruct) SetFloorNumber(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFloorNumber:"), value)
}/* debug [instance_properties/setter]: floorNumber */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdatatypelocationdescriptorstruct/locationname
func (m_ MTRDataTypeLocationDescriptorStruct) LocationName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("locationName"))
	return rv
}/* debug [instance_properties/getter]: locationName */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdatatypelocationdescriptorstruct/locationname
func (m_ MTRDataTypeLocationDescriptorStruct) SetLocationName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLocationName:"), value)
}/* debug [instance_properties/setter]: locationName */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRDataTypeLocationDescriptorStruct */



