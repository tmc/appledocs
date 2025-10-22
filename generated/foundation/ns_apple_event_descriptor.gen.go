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
	DescriptorAtIndex(index int) AppleEventDescriptor
	AttributeDescriptorForKeyword(keyword unsafe.Pointer) AppleEventDescriptor
	InsertDescriptorAtIndex(descriptor IAppleEventDescriptor, index int)
	RemoveDescriptorWithKeyword(keyword unsafe.Pointer)
	AeDesc() unsafe.Pointer
	DoubleValue() float64
	BooleanValue() bool
	SetBooleanValue(value bool)
	Data() Data
	SetData(value IData)
	DateValue() Date
	SetDateValue(value IDate)
	DescriptorType() unsafe.Pointer
	SetDescriptorType(value unsafe.Pointer)
	EnumCodeValue() unsafe.Pointer
	SetEnumCodeValue(value unsafe.Pointer)
	EventClass() unsafe.Pointer
	SetEventClass(value unsafe.Pointer)
	EventID() unsafe.Pointer
	SetEventID(value unsafe.Pointer)
	FileURLValue() URL
	SetFileURLValue(value IURL)
	Int32Value() unsafe.Pointer
	SetInt32Value(value unsafe.Pointer)
	IsRecordDescriptor() bool
	SetIsRecordDescriptor(value bool)
	NumberOfItems() int
	SetNumberOfItems(value int)
	ReturnID() unsafe.Pointer
	SetReturnID(value unsafe.Pointer)
	StringValue() string
	SetStringValue(value string)
	TransactionID() unsafe.Pointer
	SetTransactionID(value unsafe.Pointer)
	TypeCodeValue() unsafe.Pointer
	SetTypeCodeValue(value unsafe.Pointer)
}

// A wrapper for the Apple event descriptor data type.
//
// An instance of represents a descriptor—the basic building block for Apple events. This class is a wrapper for the underlying Apple event descriptor data type, . Scriptable Cocoa applications frequently work with instances of , but should rarely need to work directly with the data structure. A is a data structure that stores data and an accompanying four-character code. A descriptor can store a value, or it can store a list of other descriptors (which may also be lists). All the information in an Apple event is stored in descriptors and lists of descriptors, and every Apple event is itself a descriptor list that matches certain criteria. Descriptors can be used to build arbitrarily complex containers, so that one Apple event can represent a script statement such as . In working with Apple event descriptors, it can be useful to understand some of the underlying data types. You’ll find terms such as descriptor, descriptor list, Apple event record, and Apple event defined in Building an Apple Event in Apple Events Programming Guide. You’ll also find information on the four-character codes used to identify information within a descriptor. Apple event data types are defined in . The values of many four-character codes used by Apple (and in some cases reused by developers) can be found in . The most common reason to construct an Apple event with an instance of is to supply information in a return Apple event. The most common situation where you might need to extract information from an Apple event (as an instance of ) is when an Apple event handler installed by your application is invoked, as described in “Installing an Apple Event Handler” in . In addition, if you execute an AppleScript script using the class, you get an instance of as the return value, from which you can extract any required information. When you work with an instance of , you can access the underlying descriptor directly, if necessary, with the method. Other methods, including make it possible to create and initialize instances of without creating temporary instances of . The designated initializer for is . However, it is unlikely that you will need to create a subclass of . Cocoa doesn’t currently provide a mechanism for applications to directly send raw Apple events (though compiling and executing an AppleScript script with may result in Apple events being sent). However, Cocoa applications have full access to the Apple Event Manager C APIs for working with Apple events. So, for example, you might use an instance of to assemble an Apple event and call the Apple Event Manager function to send it. If you need to send Apple events, or if you need more information on some of the Apple event concepts described here, see Apple Events Programming Guide and .
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





// Initializes a newly allocated instance as a descriptor for the specified Carbon structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/init(aeDescNoCopy:)

func NewAppleEventDescriptorWithAEDescNoCopy(aeDesc unsafe.Pointer) AppleEventDescriptor {
	instance := getAppleEventDescriptorClass().Alloc()
	rv := objc.Send[AppleEventDescriptor](instance.ID, objc.Sel("initWithAEDescNoCopy:"), aeDesc)
	rv.Autorelease()
	return rv
}



// Creates a descriptor that represents an Apple event, initialized according to the specified information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/appleEvent(withEventClass:eventID:targetDescriptor:returnID:transactionID:)

func (ac _AppleEventDescriptorClass) AppleEventWithEventClassEventIDTargetDescriptorReturnIDTransactionID(eventClass unsafe.Pointer, eventID unsafe.Pointer, targetDescriptor IAppleEventDescriptor, returnID unsafe.Pointer, transactionID unsafe.Pointer) AppleEventDescriptor {
	rv := objc.Send[AppleEventDescriptor](objc.ID(ac.class), objc.Sel("appleEventWithEventClass:eventID:targetDescriptor:returnID:transactionID:"), eventClass, eventID, targetDescriptor, returnID, transactionID)
	return rv
}


// Creates a descriptor initialized with the specified event type that stores the specified data (from a series of bytes).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/descriptorWithDescriptorType:bytes:length:

func (ac _AppleEventDescriptorClass) DescriptorWithDescriptorTypeBytesLength(descriptorType unsafe.Pointer, bytes unsafe.Pointer, byteCount uint) AppleEventDescriptor {
	rv := objc.Send[AppleEventDescriptor](objc.ID(ac.class), objc.Sel("descriptorWithDescriptorType:bytes:length:"), descriptorType, bytes, byteCount)
	return rv
}


// Creates and initializes a descriptor with no parameter or attribute values set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/null()

func (ac _AppleEventDescriptorClass) NullDescriptor() AppleEventDescriptor {
	rv := objc.Send[AppleEventDescriptor](objc.ID(ac.class), objc.Sel("nullDescriptor"))
	return rv
}

// Returns the descriptor at the specified (one-based) position in the receiving descriptor list.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/atIndex(_:)
func (a_ AppleEventDescriptor) DescriptorAtIndex(index int) AppleEventDescriptor {
	rv := objc.Send[AppleEventDescriptor](a_.ID, objc.Sel("descriptorAtIndex:"), index)
	return rv
}

// Returns a descriptor for the receiver’s Apple event attribute identified by the specified keyword.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/attributeDescriptor(forKeyword:)
func (a_ AppleEventDescriptor) AttributeDescriptorForKeyword(keyword unsafe.Pointer) AppleEventDescriptor {
	rv := objc.Send[AppleEventDescriptor](a_.ID, objc.Sel("attributeDescriptorForKeyword:"), keyword)
	return rv
}

// Inserts a descriptor at the specified (one-based) position in the receiving descriptor list, replacing the existing descriptor, if any, at that position.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/insert(_:at:)
func (a_ AppleEventDescriptor) InsertDescriptorAtIndex(descriptor IAppleEventDescriptor, index int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("insertDescriptor:atIndex:"), descriptor, index)
}

// Removes the receiver’s descriptor identified by the specified keyword.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/remove(withKeyword:)
func (a_ AppleEventDescriptor) RemoveDescriptorWithKeyword(keyword unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("removeDescriptorWithKeyword:"), keyword)
}

// The structure encapsulated by the receiver, if it has one.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/aeDesc
func (a_ AppleEventDescriptor) AeDesc() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("aeDesc"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/doubleValue
func (a_ AppleEventDescriptor) DoubleValue() float64 {
	rv := objc.Send[float64](a_.ID, objc.Sel("doubleValue"))
	return rv
}

// The contents of the receiver as a Boolean value, coercing (to
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/booleanvalue
func (a_ AppleEventDescriptor) BooleanValue() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("booleanValue"))
	return rv
}


// SetBooleanValue sets the value of the booleanValue property.
// The contents of the receiver as a Boolean value, coercing (to

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/booleanvalue
func (a_ AppleEventDescriptor) SetBooleanValue(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setBooleanValue:"), value)
}

// The receiver’s data.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/data
func (a_ AppleEventDescriptor) Data() Data {
	rv := objc.Send[Data](a_.ID, objc.Sel("data"))
	return rv
}


// SetData sets the value of the data property.
// The receiver’s data.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/data
func (a_ AppleEventDescriptor) SetData(value IData) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setData:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/datevalue
func (a_ AppleEventDescriptor) DateValue() Date {
	rv := objc.Send[Date](a_.ID, objc.Sel("dateValue"))
	return rv
}


// SetDateValue sets the value of the dateValue property.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/datevalue
func (a_ AppleEventDescriptor) SetDateValue(value IDate) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDateValue:"), value)
}

// The descriptor type of the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/descriptortype
func (a_ AppleEventDescriptor) DescriptorType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("descriptorType"))
	return rv
}


// SetDescriptorType sets the value of the descriptorType property.
// The descriptor type of the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/descriptortype
func (a_ AppleEventDescriptor) SetDescriptorType(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDescriptorType:"), value)
}

// The contents of the receiver as an enumeration type, coercing to
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/enumcodevalue
func (a_ AppleEventDescriptor) EnumCodeValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("enumCodeValue"))
	return rv
}


// SetEnumCodeValue sets the value of the enumCodeValue property.
// The contents of the receiver as an enumeration type, coercing to

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/enumcodevalue
func (a_ AppleEventDescriptor) SetEnumCodeValue(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setEnumCodeValue:"), value)
}

// The event class for the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/eventclass
func (a_ AppleEventDescriptor) EventClass() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("eventClass"))
	return rv
}


// SetEventClass sets the value of the eventClass property.
// The event class for the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/eventclass
func (a_ AppleEventDescriptor) SetEventClass(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setEventClass:"), value)
}

// The event ID for the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/eventid
func (a_ AppleEventDescriptor) EventID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("eventID"))
	return rv
}


// SetEventID sets the value of the eventID property.
// The event ID for the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/eventid
func (a_ AppleEventDescriptor) SetEventID(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setEventID:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/fileurlvalue
func (a_ AppleEventDescriptor) FileURLValue() URL {
	rv := objc.Send[URL](a_.ID, objc.Sel("fileURLValue"))
	return rv
}


// SetFileURLValue sets the value of the fileURLValue property.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/fileurlvalue
func (a_ AppleEventDescriptor) SetFileURLValue(value IURL) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setFileURLValue:"), value)
}

// The contents of the receiver as an integer, coercing (to
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/int32value
func (a_ AppleEventDescriptor) Int32Value() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("int32Value"))
	return rv
}


// SetInt32Value sets the value of the int32Value property.
// The contents of the receiver as an integer, coercing (to

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/int32value
func (a_ AppleEventDescriptor) SetInt32Value(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setInt32Value:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/isrecorddescriptor
func (a_ AppleEventDescriptor) IsRecordDescriptor() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isRecordDescriptor"))
	return rv
}


// SetIsRecordDescriptor sets the value of the isRecordDescriptor property.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/isrecorddescriptor
func (a_ AppleEventDescriptor) SetIsRecordDescriptor(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsRecordDescriptor:"), value)
}

// The number of descriptors in the receiver’s descriptor list.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/numberofitems
func (a_ AppleEventDescriptor) NumberOfItems() int {
	rv := objc.Send[int](a_.ID, objc.Sel("numberOfItems"))
	return rv
}


// SetNumberOfItems sets the value of the numberOfItems property.
// The number of descriptors in the receiver’s descriptor list.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/numberofitems
func (a_ AppleEventDescriptor) SetNumberOfItems(value int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setNumberOfItems:"), value)
}

// The receiver’s return ID (the ID for a reply Apple event).
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/returnid
func (a_ AppleEventDescriptor) ReturnID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("returnID"))
	return rv
}


// SetReturnID sets the value of the returnID property.
// The receiver’s return ID (the ID for a reply Apple event).

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/returnid
func (a_ AppleEventDescriptor) SetReturnID(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setReturnID:"), value)
}

// The contents of the receiver as a Unicode text string, coercing to
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/stringvalue
func (a_ AppleEventDescriptor) StringValue() string {
	rv := objc.Send[string](a_.ID, objc.Sel("stringValue"))
	return rv
}


// SetStringValue sets the value of the stringValue property.
// The contents of the receiver as a Unicode text string, coercing to

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/stringvalue
func (a_ AppleEventDescriptor) SetStringValue(value string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setStringValue:"), objc.String(value))
}

// The receiver’s transaction ID, if any.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/transactionid
func (a_ AppleEventDescriptor) TransactionID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("transactionID"))
	return rv
}


// SetTransactionID sets the value of the transactionID property.
// The receiver’s transaction ID, if any.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/transactionid
func (a_ AppleEventDescriptor) SetTransactionID(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setTransactionID:"), value)
}

// The contents of the receiver as a type, coercing to
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/typecodevalue
func (a_ AppleEventDescriptor) TypeCodeValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("typeCodeValue"))
	return rv
}


// SetTypeCodeValue sets the value of the typeCodeValue property.
// The contents of the receiver as a type, coercing to

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsappleeventdescriptor/typecodevalue
func (a_ AppleEventDescriptor) SetTypeCodeValue(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setTypeCodeValue:"), value)
}


