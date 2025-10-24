// Code generated from Apple documentation for OSLog. DO NOT EDIT.

package oslog

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [OSLogMessageComponent] class.
type IOSLogMessageComponent interface {
	objectivec.IObject
	// properties:
	Argument() unsafe.Pointer
	SetArgument(value unsafe.Pointer)
	ArgumentCategory() unsafe.Pointer
	SetArgumentCategory(value unsafe.Pointer)
	ArgumentDataValue() objc.IObject /* cross-framework: Data */
	SetArgumentDataValue(value objc.IObject /* cross-framework: Data */)
	ArgumentDoubleValue() float64
	SetArgumentDoubleValue(value float64)
	ArgumentInt64Value() unsafe.Pointer
	SetArgumentInt64Value(value unsafe.Pointer)
	ArgumentNumberValue() objc.IObject /* cross-framework: NSNumber */
	SetArgumentNumberValue(value objc.IObject /* cross-framework: NSNumber */)
	ArgumentStringValue() objc.IObject /* cross-framework: NSString */
	SetArgumentStringValue(value objc.IObject /* cross-framework: NSString */)
	ArgumentUInt64Value() uint64
	SetArgumentUInt64Value(value uint64)
	FormatSubstring() objc.IObject /* cross-framework: NSString */
	SetFormatSubstring(value objc.IObject /* cross-framework: NSString */)
	Placeholder() objc.IObject /* cross-framework: NSString */
	SetPlaceholder(value objc.IObject /* cross-framework: NSString */)
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (oc _OSLogMessageComponentClass) Alloc() OSLogMessageComponent {
	rv := objc.Send[OSLogMessageComponent](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// The argument passed into the message component.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/oslog/oslogmessagecomponent/argument-swift.property
func (o_ OSLogMessageComponent) Argument() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("argument"))
	return rv
}


// The argument passed into the message component.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/oslog/oslogmessagecomponent/argument-swift.property
func (o_ OSLogMessageComponent) SetArgument(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setArgument:"), value)
}


// The type of argument that corresponds to the placeholder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/oslog/oslogmessagecomponent/argumentcategory-swift.property
func (o_ OSLogMessageComponent) ArgumentCategory() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("argumentCategory"))
	return rv
}


// The type of argument that corresponds to the placeholder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/oslog/oslogmessagecomponent/argumentcategory-swift.property
func (o_ OSLogMessageComponent) SetArgumentCategory(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setArgumentCategory:"), value)
}


// The argument formatted as a sequence of bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/oslog/oslogmessagecomponent/argumentdatavalue
func (o_ OSLogMessageComponent) ArgumentDataValue() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](o_.ID, objc.Sel("argumentDataValue"))
	return rv
}


// The argument formatted as a sequence of bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/oslog/oslogmessagecomponent/argumentdatavalue
func (o_ OSLogMessageComponent) SetArgumentDataValue(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setArgumentDataValue:"), value)
}


// The argument formatted as a double.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/oslog/oslogmessagecomponent/argumentdoublevalue
func (o_ OSLogMessageComponent) ArgumentDoubleValue() float64 {
	rv := objc.Send[float64](o_.ID, objc.Sel("argumentDoubleValue"))
	return rv
}


// The argument formatted as a double.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/oslog/oslogmessagecomponent/argumentdoublevalue
func (o_ OSLogMessageComponent) SetArgumentDoubleValue(value float64) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setArgumentDoubleValue:"), value)
}


// The argument formatted as a signed 64-bit integer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/oslog/oslogmessagecomponent/argumentint64value
func (o_ OSLogMessageComponent) ArgumentInt64Value() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("argumentInt64Value"))
	return rv
}


// The argument formatted as a signed 64-bit integer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/oslog/oslogmessagecomponent/argumentint64value
func (o_ OSLogMessageComponent) SetArgumentInt64Value(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setArgumentInt64Value:"), value)
}


// The argument formatted as a number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/oslog/oslogmessagecomponent/argumentnumbervalue
func (o_ OSLogMessageComponent) ArgumentNumberValue() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](o_.ID, objc.Sel("argumentNumberValue"))
	return rv
}


// The argument formatted as a number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/oslog/oslogmessagecomponent/argumentnumbervalue
func (o_ OSLogMessageComponent) SetArgumentNumberValue(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setArgumentNumberValue:"), value)
}


// The argument formatted as a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/oslog/oslogmessagecomponent/argumentstringvalue
func (o_ OSLogMessageComponent) ArgumentStringValue() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](o_.ID, objc.Sel("argumentStringValue"))
	return rv
}


// The argument formatted as a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/oslog/oslogmessagecomponent/argumentstringvalue
func (o_ OSLogMessageComponent) SetArgumentStringValue(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setArgumentStringValue:"), value)
}


// The argument formatted as an unsigned 64-bit integer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/oslog/oslogmessagecomponent/argumentuint64value
func (o_ OSLogMessageComponent) ArgumentUInt64Value() uint64 {
	rv := objc.Send[uint64](o_.ID, objc.Sel("argumentUInt64Value"))
	return rv
}


// The argument formatted as an unsigned 64-bit integer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/oslog/oslogmessagecomponent/argumentuint64value
func (o_ OSLogMessageComponent) SetArgumentUInt64Value(value uint64) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setArgumentUInt64Value:"), value)
}


// The text immediately preceding a placeholder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/oslog/oslogmessagecomponent/formatsubstring
func (o_ OSLogMessageComponent) FormatSubstring() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](o_.ID, objc.Sel("formatSubstring"))
	return rv
}


// The text immediately preceding a placeholder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/oslog/oslogmessagecomponent/formatsubstring
func (o_ OSLogMessageComponent) SetFormatSubstring(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setFormatSubstring:"), value)
}


// The placeholder text for the message component.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/oslog/oslogmessagecomponent/placeholder
func (o_ OSLogMessageComponent) Placeholder() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](o_.ID, objc.Sel("placeholder"))
	return rv
}


// The placeholder text for the message component.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/oslog/oslogmessagecomponent/placeholder
func (o_ OSLogMessageComponent) SetPlaceholder(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setPlaceholder:"), value)
}



