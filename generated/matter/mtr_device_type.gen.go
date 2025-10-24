// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRDeviceType */


/* debug [class_header]: Header for MTRDeviceType */
// The class instance for the [MTRDeviceType] class.
var (
	MTRDeviceTypeClass     _MTRDeviceTypeClass
	MTRDeviceTypeClassOnce sync.Once
)

func getMTRDeviceTypeClass() _MTRDeviceTypeClass {
	MTRDeviceTypeClassOnce.Do(func() {
		MTRDeviceTypeClass = _MTRDeviceTypeClass{objc.GetClass("MTRDeviceType")}
	})
	return MTRDeviceTypeClass
}

type _MTRDeviceTypeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRDeviceType */
// An interface definition for the [MTRDeviceType] class.
type IMTRDeviceType interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRDeviceType */
	// properties:
	Id() objc.IObject /* cross-framework: NSNumber */
	SetId(value objc.IObject /* cross-framework: NSNumber */)
	IsUtility() bool
	SetIsUtility(value bool)
	Name() objc.IObject /* cross-framework: NSString */
	SetName(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRDeviceType */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRDeviceType */
// Alloc allocates a new instance without initialization.
func (mc _MTRDeviceTypeClass) Alloc() MTRDeviceType {
	rv := objc.Send[MTRDeviceType](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRDeviceTypeClass) New() MTRDeviceType {
	rv := objc.Send[MTRDeviceType](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDeviceType) Init() MTRDeviceType {
	rv := objc.Send[MTRDeviceType](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDeviceType) Autorelease() MTRDeviceType {
	rv := objc.Send[MTRDeviceType](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDeviceType creates a new MTRDeviceType instance.
func NewMTRDeviceType() MTRDeviceType {
	return getMTRDeviceTypeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRDeviceType */
// Meta-data about a device type defined in the Matter specification.


// Meta-data about a device type defined in the Matter specification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceType
type MTRDeviceType struct {
	objectivec.Object
}

// MTRDeviceTypeFrom constructs a [MTRDeviceType] from an unsafe.Pointer.
//
// Meta-data about a device type defined in the Matter specification.
func MTRDeviceTypeFrom(ptr unsafe.Pointer) MTRDeviceType {
	return MTRDeviceType{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRDeviceType */

// Returns an MTRDeviceType for the given ID, if the ID is known. Returns nil for unknown IDs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceType/init(forID:)
func NewMTRDeviceTypeForID(deviceTypeID objc.IObject /* cross-framework: NSNumber */) MTRDeviceType {
	rv := objc.Send[MTRDeviceType](objc.ID(getMTRDeviceTypeClass().class), objc.Sel("deviceTypeForID:"), deviceTypeID)
	return rv
}/* debug [class_init_methods/constructor]: NewMTRDeviceTypeForID */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRDeviceType */

// Returns an MTRDeviceType for the given ID, if the ID is known. Returns nil for unknown IDs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceType/init(forID:)
func (mc _MTRDeviceTypeClass) DeviceTypeForID(deviceTypeID objc.IObject /* cross-framework: NSNumber */) MTRDeviceType {
	rv := objc.Send[MTRDeviceType](objc.ID(mc.class), objc.Sel("deviceTypeForID:"), deviceTypeID)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DeviceTypeForID) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRDeviceType */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRDeviceType */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRDeviceType */

// The identifier of the device type (32-bit unsigned integer).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicetype/id
func (m_ MTRDeviceType) Id() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("id"))
	return rv
}/* debug [instance_properties/getter]: id */


// The identifier of the device type (32-bit unsigned integer).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicetype/id
func (m_ MTRDeviceType) SetId(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setId:"), value)
}/* debug [instance_properties/setter]: id */


// Returns whether this is a utility device type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicetype/isutility
func (m_ MTRDeviceType) IsUtility() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isUtility"))
	return rv
}/* debug [instance_properties/getter]: isUtility */


// Returns whether this is a utility device type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicetype/isutility
func (m_ MTRDeviceType) SetIsUtility(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsUtility:"), value)
}/* debug [instance_properties/setter]: isUtility */


// Returns the name of the device type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicetype/name
func (m_ MTRDeviceType) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// Returns the name of the device type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicetype/name
func (m_ MTRDeviceType) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setName:"), value)
}/* debug [instance_properties/setter]: name */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRDeviceType */


