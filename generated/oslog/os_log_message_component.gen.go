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
}

// The message arguments for a particular entry.
//
// There is one component for each placeholder in the formatString plus one component for any text after the last placeholder.
//
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


// The type of argument that corresponds to the placeholder.
//
// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogMessageComponent/argumentCategory-swift.property
func (o_ OSLogMessageComponent) ArgumentCategory() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("argumentCategory"))
	return rv
}

// The argument formatted as a sequence of bytes.
//
// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogMessageComponent/argumentDataValue
func (o_ OSLogMessageComponent) ArgumentDataValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("argumentDataValue"))
	return rv
}

// The argument formatted as a double.
//
// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogMessageComponent/argumentDoubleValue
func (o_ OSLogMessageComponent) ArgumentDoubleValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("argumentDoubleValue"))
	return rv
}

// The argument formatted as a signed 64-bit integer.
//
// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogMessageComponent/argumentInt64Value
func (o_ OSLogMessageComponent) ArgumentInt64Value() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("argumentInt64Value"))
	return rv
}

// The argument formatted as a number.
//
// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogMessageComponent/argumentNumberValue
func (o_ OSLogMessageComponent) ArgumentNumberValue() foundation.Number {
	rv := objc.Send[foundation.Number](o_.ID, objc.Sel("argumentNumberValue"))
	return rv
}

// The argument formatted as a string.
//
// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogMessageComponent/argumentStringValue
func (o_ OSLogMessageComponent) ArgumentStringValue() string {
	rv := objc.Send[string](o_.ID, objc.Sel("argumentStringValue"))
	return rv
}

// The argument formatted as an unsigned 64-bit integer.
//
// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogMessageComponent/argumentUInt64Value
func (o_ OSLogMessageComponent) ArgumentUInt64Value() uint64 {
	rv := objc.Send[uint64](o_.ID, objc.Sel("argumentUInt64Value"))
	return rv
}

// The text immediately preceding a placeholder.
//
// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogMessageComponent/formatSubstring
func (o_ OSLogMessageComponent) FormatSubstring() string {
	rv := objc.Send[string](o_.ID, objc.Sel("formatSubstring"))
	return rv
}

// The placeholder text for the message component.
//
// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogMessageComponent/placeholder
func (o_ OSLogMessageComponent) Placeholder() string {
	rv := objc.Send[string](o_.ID, objc.Sel("placeholder"))
	return rv
}

// The argument passed into the message component.
//
// [Full Topic]: https://developer.apple.com/documentation/oslog/oslogmessagecomponent/argument-swift.property
func (o_ OSLogMessageComponent) Argument() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("argument"))
	return rv
}


// SetArgument sets the value of the argument property.
// The argument passed into the message component.

//
// [Full Topic]: https://developer.apple.com/documentation/oslog/oslogmessagecomponent/argument-swift.property
func (o_ OSLogMessageComponent) SetArgument(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setArgument:"), value)
}



