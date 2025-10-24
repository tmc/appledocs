// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AppleEventDescriptor] class.
var (
	AppleEventDescriptorClass     _AppleEventDescriptorClass
	AppleEventDescriptorClassOnce sync.Once
)

func getAppleEventDescriptorClass() _AppleEventDescriptorClass {
	AppleEventDescriptorClassOnce.Do(func() {
		AppleEventDescriptorClass = _AppleEventDescriptorClass{objc.GetClass("NSAppleEventDescriptor")}
	})
	return AppleEventDescriptorClass
}

type _AppleEventDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [AppleEventDescriptor] class.
type IAppleEventDescriptor interface {
	objectivec.IObject
	// properties:
	AeDesc() unsafe.Pointer
	BooleanValue() bool
	SetBooleanValue(value bool)
	Data() IData
	SetData(value IData)
	DateValue() IDate
	SetDateValue(value IDate)
	DescriptorType() unsafe.Pointer
	SetDescriptorType(value unsafe.Pointer)
	DoubleValue() float64
	SetDoubleValue(value float64)
	EnumCodeValue() uint32 /* not a class type */
	SetEnumCodeValue(value uint32 /* not a class type */)
	EventClass() unsafe.Pointer
	SetEventClass(value unsafe.Pointer)
	EventID() unsafe.Pointer
	SetEventID(value unsafe.Pointer)
	FileURLValue() IURL
	SetFileURLValue(value IURL)
	Int32Value() unsafe.Pointer
	SetInt32Value(value unsafe.Pointer)
	IsRecordDescriptor() bool
	SetIsRecordDescriptor(value bool)
	NumberOfItems() int
	SetNumberOfItems(value int)
	ReturnID() unsafe.Pointer
	SetReturnID(value unsafe.Pointer)
	StringValue() IString
	SetStringValue(value IString)
	TransactionID() unsafe.Pointer
	SetTransactionID(value unsafe.Pointer)
	TypeCodeValue() uint32 /* not a class type */
	SetTypeCodeValue(value uint32 /* not a class type */)
	// methods:
}

// A wrapper for the Apple event descriptor data type.
//
// An instance of represents a descriptor—the basic building block for Apple events. This class is a wrapper for the underlying Apple event descriptor data type, . Scriptable Cocoa applications frequently work with instances of , but should rarely need to work directly with the data structure. A is a data structure that stores data and an accompanying four-character code. A descriptor can store a value, or it can store a list of other descriptors (which may also be lists). All the information in an Apple event is stored in descriptors and lists of descriptors, and every Apple event is itself a descriptor list that matches certain criteria. Descriptors can be used to build arbitrarily complex containers, so that one Apple event can represent a script statement such as . In working with Apple event descriptors, it can be useful to understand some of the underlying data types. You’ll find terms such as descriptor, descriptor list, Apple event record, and Apple event defined in Building an Apple Event in Apple Events Programming Guide. You’ll also find information on the four-character codes used to identify information within a descriptor. Apple event data types are defined in . The values of many four-character codes used by Apple (and in some cases reused by developers) can be found in . The most common reason to construct an Apple event with an instance of is to supply information in a return Apple event. The most common situation where you might need to extract information from an Apple event (as an instance of ) is when an Apple event handler installed by your application is invoked, as described in “Installing an Apple Event Handler” in . In addition, if you execute an AppleScript script using the class, you get an instance of as the return value, from which you can extract any required information. When you work with an instance of , you can access the underlying descriptor directly, if necessary, with the method. Other methods, including make it possible to create and initialize instances of without creating temporary instances of . The designated initializer for is . However, it is unlikely that you will need to create a subclass of . Cocoa doesn’t currently provide a mechanism for applications to directly send raw Apple events (though compiling and executing an AppleScript script with may result in Apple events being sent). However, Cocoa applications have full access to the Apple Event Manager C APIs for working with Apple events. So, for example, you might use an instance of to assemble an Apple event and call the Apple Event Manager function to send it. If you need to send Apple events, or if you need more information on some of the Apple event concepts described here, see Apple Events Programming Guide and .


// A wrapper for the Apple event descriptor data type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor
type AppleEventDescriptor struct {
	objectivec.Object
}

// AppleEventDescriptorFrom constructs a [AppleEventDescriptor] from an unsafe.Pointer.
//
// A wrapper for the Apple event descriptor data type.
func AppleEventDescriptorFrom(ptr unsafe.Pointer) AppleEventDescriptor {
	return AppleEventDescriptor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AppleEventDescriptorClass) Alloc() AppleEventDescriptor {
	rv := objc.Send[AppleEventDescriptor](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AppleEventDescriptorClass) New() AppleEventDescriptor {
	rv := objc.Send[AppleEventDescriptor](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AppleEventDescriptor) Init() AppleEventDescriptor {
	rv := objc.Send[AppleEventDescriptor](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AppleEventDescriptor) Autorelease() AppleEventDescriptor {
	rv := objc.Send[AppleEventDescriptor](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAppleEventDescriptor creates a new AppleEventDescriptor instance.
func NewAppleEventDescriptor() AppleEventDescriptor {
	return getAppleEventDescriptorClass().New()
}



// The structure encapsulated by the receiver, if it has one.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/aeDesc
func (a_ AppleEventDescriptor) AeDesc() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("aeDesc"))
	return rv
}


// The contents of the receiver as a Boolean value, coercing (to
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/booleanvalue
func (a_ AppleEventDescriptor) BooleanValue() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("booleanValue"))
	return rv
}


// The contents of the receiver as a Boolean value, coercing (to
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/booleanvalue
func (a_ AppleEventDescriptor) SetBooleanValue(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setBooleanValue:"), value)
}


// The receiver’s data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/data
func (a_ AppleEventDescriptor) Data() IData {
	rv := objc.Send[Data](a_.ID, objc.Sel("data"))
	return rv
}


// The receiver’s data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/data
func (a_ AppleEventDescriptor) SetData(value IData) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setData:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/datevalue
func (a_ AppleEventDescriptor) DateValue() IDate {
	rv := objc.Send[Date](a_.ID, objc.Sel("dateValue"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/datevalue
func (a_ AppleEventDescriptor) SetDateValue(value IDate) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDateValue:"), value)
}


// The descriptor type of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/descriptortype
func (a_ AppleEventDescriptor) DescriptorType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("descriptorType"))
	return rv
}


// The descriptor type of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/descriptortype
func (a_ AppleEventDescriptor) SetDescriptorType(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDescriptorType:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/doublevalue
func (a_ AppleEventDescriptor) DoubleValue() float64 {
	rv := objc.Send[float64](a_.ID, objc.Sel("doubleValue"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/doublevalue
func (a_ AppleEventDescriptor) SetDoubleValue(value float64) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDoubleValue:"), value)
}


// The contents of the receiver as an enumeration type, coercing to
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/enumcodevalue
func (a_ AppleEventDescriptor) EnumCodeValue() uint32 /* not a class type */ {
	rv := objc.Send[uint32](a_.ID, objc.Sel("enumCodeValue"))
	return rv
}


// The contents of the receiver as an enumeration type, coercing to
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/enumcodevalue
func (a_ AppleEventDescriptor) SetEnumCodeValue(value uint32 /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setEnumCodeValue:"), value)
}


// The event class for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/eventclass
func (a_ AppleEventDescriptor) EventClass() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("eventClass"))
	return rv
}


// The event class for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/eventclass
func (a_ AppleEventDescriptor) SetEventClass(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setEventClass:"), value)
}


// The event ID for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/eventid
func (a_ AppleEventDescriptor) EventID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("eventID"))
	return rv
}


// The event ID for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/eventid
func (a_ AppleEventDescriptor) SetEventID(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setEventID:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/fileurlvalue
func (a_ AppleEventDescriptor) FileURLValue() IURL {
	rv := objc.Send[URL](a_.ID, objc.Sel("fileURLValue"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/fileurlvalue
func (a_ AppleEventDescriptor) SetFileURLValue(value IURL) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setFileURLValue:"), value)
}


// The contents of the receiver as an integer, coercing (to
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/int32value
func (a_ AppleEventDescriptor) Int32Value() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("int32Value"))
	return rv
}


// The contents of the receiver as an integer, coercing (to
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/int32value
func (a_ AppleEventDescriptor) SetInt32Value(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setInt32Value:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/isrecorddescriptor
func (a_ AppleEventDescriptor) IsRecordDescriptor() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isRecordDescriptor"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/isrecorddescriptor
func (a_ AppleEventDescriptor) SetIsRecordDescriptor(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsRecordDescriptor:"), value)
}


// The number of descriptors in the receiver’s descriptor list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/numberofitems
func (a_ AppleEventDescriptor) NumberOfItems() int {
	rv := objc.Send[int](a_.ID, objc.Sel("numberOfItems"))
	return rv
}


// The number of descriptors in the receiver’s descriptor list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/numberofitems
func (a_ AppleEventDescriptor) SetNumberOfItems(value int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setNumberOfItems:"), value)
}


// The receiver’s return ID (the ID for a reply Apple event).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/returnid
func (a_ AppleEventDescriptor) ReturnID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("returnID"))
	return rv
}


// The receiver’s return ID (the ID for a reply Apple event).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/returnid
func (a_ AppleEventDescriptor) SetReturnID(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setReturnID:"), value)
}


// The contents of the receiver as a Unicode text string, coercing to
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/stringvalue
func (a_ AppleEventDescriptor) StringValue() IString {
	rv := objc.Send[String](a_.ID, objc.Sel("stringValue"))
	return rv
}


// The contents of the receiver as a Unicode text string, coercing to
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/stringvalue
func (a_ AppleEventDescriptor) SetStringValue(value IString) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setStringValue:"), value)
}


// The receiver’s transaction ID, if any.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/transactionid
func (a_ AppleEventDescriptor) TransactionID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("transactionID"))
	return rv
}


// The receiver’s transaction ID, if any.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/transactionid
func (a_ AppleEventDescriptor) SetTransactionID(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setTransactionID:"), value)
}


// The contents of the receiver as a type, coercing to
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/typecodevalue
func (a_ AppleEventDescriptor) TypeCodeValue() uint32 /* not a class type */ {
	rv := objc.Send[uint32](a_.ID, objc.Sel("typeCodeValue"))
	return rv
}


// The contents of the receiver as a type, coercing to
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/typecodevalue
func (a_ AppleEventDescriptor) SetTypeCodeValue(value uint32 /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setTypeCodeValue:"), value)
}



