// Code generated from Apple documentation for OSLog. DO NOT EDIT.

package oslog

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/metal"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class OSLogMessageComponent */


/* debug [class_header]: Header for OSLogMessageComponent */
// The class instance for the [OSLogMessageComponent] class.
var (
	OSLogMessageComponentClass     _OSLogMessageComponentClass
	OSLogMessageComponentClassOnce sync.Once
)

func getOSLogMessageComponentClass() _OSLogMessageComponentClass {
	OSLogMessageComponentClassOnce.Do(func() {
		OSLogMessageComponentClass = _OSLogMessageComponentClass{objc.GetClass("OSLogMessageComponent")}
	})
	return OSLogMessageComponentClass
}

type _OSLogMessageComponentClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for OSLogMessageComponent */
// An interface definition for the [OSLogMessageComponent] class.
type IOSLogMessageComponent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for OSLogMessageComponent */
	// properties:
	ArgumentCategory() OSLogMessageComponentArgumentCategory
	ArgumentDataValue() objc.IObject /* cross-framework: NSData */
	ArgumentDoubleValue() float64
	ArgumentInt64Value() int64
	ArgumentNumberValue() objc.IObject /* cross-framework: NSNumber */
	ArgumentStringValue() objc.IObject /* cross-framework: NSString */
	ArgumentUInt64Value() uint64
	FormatSubstring() objc.IObject /* cross-framework: NSString */
	Placeholder() objc.IObject /* cross-framework: NSString */
	Argument() metal.Argument
	SetArgument(value metal.Argument)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for OSLogMessageComponent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for OSLogMessageComponent */
// Alloc allocates a new instance without initialization.
func (oc _OSLogMessageComponentClass) Alloc() OSLogMessageComponent {
	rv := objc.Send[OSLogMessageComponent](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (oc _OSLogMessageComponentClass) New() OSLogMessageComponent {
	rv := objc.Send[OSLogMessageComponent](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ OSLogMessageComponent) Init() OSLogMessageComponent {
	rv := objc.Send[OSLogMessageComponent](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ OSLogMessageComponent) Autorelease() OSLogMessageComponent {
	rv := objc.Send[OSLogMessageComponent](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewOSLogMessageComponent creates a new OSLogMessageComponent instance.
func NewOSLogMessageComponent() OSLogMessageComponent {
	return getOSLogMessageComponentClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for OSLogMessageComponent */
// The message arguments for a particular entry.
//
// There is one component for each placeholder in the formatString plus one component for any text after the last placeholder.


// The message arguments for a particular entry.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogMessageComponent
type OSLogMessageComponent struct {
	objectivec.Object
}

// OSLogMessageComponentFrom constructs a [OSLogMessageComponent] from an unsafe.Pointer.
//
// The message arguments for a particular entry.
func OSLogMessageComponentFrom(ptr unsafe.Pointer) OSLogMessageComponent {
	return OSLogMessageComponent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for OSLogMessageComponent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for OSLogMessageComponent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for OSLogMessageComponent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for OSLogMessageComponent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for OSLogMessageComponent */

// The type of argument that corresponds to the placeholder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogMessageComponent/argumentCategory-swift.property
func (o_ OSLogMessageComponent) ArgumentCategory() OSLogMessageComponentArgumentCategory {
	rv := objc.Send[OSLogMessageComponentArgumentCategory](o_.ID, objc.Sel("argumentCategory"))
	return rv
}/* debug [instance_properties/getter]: argumentCategory */


// The argument formatted as a sequence of bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogMessageComponent/argumentDataValue
func (o_ OSLogMessageComponent) ArgumentDataValue() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](o_.ID, objc.Sel("argumentDataValue"))
	return rv
}/* debug [instance_properties/getter]: argumentDataValue */


// The argument formatted as a double.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogMessageComponent/argumentDoubleValue
func (o_ OSLogMessageComponent) ArgumentDoubleValue() float64 {
	rv := objc.Send[float64](o_.ID, objc.Sel("argumentDoubleValue"))
	return rv
}/* debug [instance_properties/getter]: argumentDoubleValue */


// The argument formatted as a signed 64-bit integer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogMessageComponent/argumentInt64Value
func (o_ OSLogMessageComponent) ArgumentInt64Value() int64 {
	rv := objc.Send[int64](o_.ID, objc.Sel("argumentInt64Value"))
	return rv
}/* debug [instance_properties/getter]: argumentInt64Value */


// The argument formatted as a number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogMessageComponent/argumentNumberValue
func (o_ OSLogMessageComponent) ArgumentNumberValue() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](o_.ID, objc.Sel("argumentNumberValue"))
	return rv
}/* debug [instance_properties/getter]: argumentNumberValue */


// The argument formatted as a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogMessageComponent/argumentStringValue
func (o_ OSLogMessageComponent) ArgumentStringValue() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](o_.ID, objc.Sel("argumentStringValue"))
	return rv
}/* debug [instance_properties/getter]: argumentStringValue */


// The argument formatted as an unsigned 64-bit integer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogMessageComponent/argumentUInt64Value
func (o_ OSLogMessageComponent) ArgumentUInt64Value() uint64 {
	rv := objc.Send[uint64](o_.ID, objc.Sel("argumentUInt64Value"))
	return rv
}/* debug [instance_properties/getter]: argumentUInt64Value */


// The text immediately preceding a placeholder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogMessageComponent/formatSubstring
func (o_ OSLogMessageComponent) FormatSubstring() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](o_.ID, objc.Sel("formatSubstring"))
	return rv
}/* debug [instance_properties/getter]: formatSubstring */


// The placeholder text for the message component.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogMessageComponent/placeholder
func (o_ OSLogMessageComponent) Placeholder() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](o_.ID, objc.Sel("placeholder"))
	return rv
}/* debug [instance_properties/getter]: placeholder */


// The argument passed into the message component.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/oslog/oslogmessagecomponent/argument-swift.property
func (o_ OSLogMessageComponent) Argument() metal.Argument {
	rv := objc.Send[metal.Argument](o_.ID, objc.Sel("argument"))
	return rv
}/* debug [instance_properties/getter]: argument */


// The argument passed into the message component.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/oslog/oslogmessagecomponent/argument-swift.property
func (o_ OSLogMessageComponent) SetArgument(value metal.Argument) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setArgument:"), value)
}/* debug [instance_properties/setter]: argument */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class OSLogMessageComponent */



