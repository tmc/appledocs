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
	CoerceToDescriptorType(descriptorType unsafe.Pointer) AppleEventDescriptor
	DescriptorForKeyword(keyword unsafe.Pointer) AppleEventDescriptor
	InsertDescriptorAtIndex(descriptor IAppleEventDescriptor, index int)
	KeywordForDescriptorAtIndex(index int) unsafe.Pointer
	ParamDescriptorForKeyword(keyword unsafe.Pointer) AppleEventDescriptor
	RemoveDescriptorAtIndex(index int)
	RemoveDescriptorWithKeyword(keyword unsafe.Pointer)
	RemoveParamDescriptorWithKeyword(keyword unsafe.Pointer)
	SendEventWithOptionsTimeoutError(sendOptions NSAppleEventSendOptions, timeoutInSeconds TimeInterval, error_ IError) AppleEventDescriptor
	SetAttributeDescriptorForKeyword(descriptor IAppleEventDescriptor, keyword unsafe.Pointer)
	SetDescriptorForKeyword(descriptor IAppleEventDescriptor, keyword unsafe.Pointer)
	SetParamDescriptorForKeyword(descriptor IAppleEventDescriptor, keyword unsafe.Pointer)
	AeDesc() unsafe.Pointer
	BooleanValue() unsafe.Pointer
	Data() NSData
	DateValue() NSDate
	DescriptorType() unsafe.Pointer
	DoubleValue() float64
	EnumCodeValue() unsafe.Pointer
	EventClass() unsafe.Pointer
	EventID() unsafe.Pointer
	FileURLValue() URL
	Int32Value() unsafe.Pointer
	IsRecordDescriptor() bool
	NumberOfItems() int
	ReturnID() unsafe.Pointer
	StringValue() string
	TransactionID() unsafe.Pointer
	TypeCodeValue() unsafe.Pointer
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



// Initializes a newly allocated instance as an empty list descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/init(listDescriptor:)
func NewAppleEventDescriptorListDescriptor() AppleEventDescriptor {
	instance := getAppleEventDescriptorClass().Alloc()
	rv := objc.Send[AppleEventDescriptor](instance.ID, objc.Sel("initListDescriptor"))
	rv.Autorelease()
	return rv
}


// Initializes a newly allocated instance as a descriptor that is an Apple event record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/init(recordDescriptor:)
func NewAppleEventDescriptorRecordDescriptor() AppleEventDescriptor {
	instance := getAppleEventDescriptorClass().Alloc()
	rv := objc.Send[AppleEventDescriptor](instance.ID, objc.Sel("initRecordDescriptor"))
	rv.Autorelease()
	return rv
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


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/init(applicationURL:)
func NewAppleEventDescriptorWithApplicationURL(applicationURL IURL) AppleEventDescriptor {
	rv := objc.Send[AppleEventDescriptor](objc.ID(getAppleEventDescriptorClass().class), objc.Sel("descriptorWithApplicationURL:"), applicationURL)
	return rv
}


// Creates a descriptor initialized with type that stores the specified Boolean value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/init(boolean:)
func NewAppleEventDescriptorWithBoolean(boolean unsafe.Pointer) AppleEventDescriptor {
	rv := objc.Send[AppleEventDescriptor](objc.ID(getAppleEventDescriptorClass().class), objc.Sel("descriptorWithBoolean:"), boolean)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/init(bundleIdentifier:)
func NewAppleEventDescriptorWithBundleIdentifier(bundleIdentifier string) AppleEventDescriptor {
	rv := objc.Send[AppleEventDescriptor](objc.ID(getAppleEventDescriptorClass().class), objc.Sel("descriptorWithBundleIdentifier:"), objc.String(bundleIdentifier))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/init(date:)
func NewAppleEventDescriptorWithDate(date IDate) AppleEventDescriptor {
	rv := objc.Send[AppleEventDescriptor](objc.ID(getAppleEventDescriptorClass().class), objc.Sel("descriptorWithDate:"), date)
	return rv
}


// Initializes a newly allocated instance as a descriptor with the specified descriptor type and data (from an arbitrary sequence of bytes and a length count).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/init(descriptorType:bytes:length:)
func NewAppleEventDescriptorWithDescriptorTypeBytesLength(descriptorType unsafe.Pointer, bytes unsafe.Pointer, byteCount uint) AppleEventDescriptor {
	instance := getAppleEventDescriptorClass().Alloc()
	rv := objc.Send[AppleEventDescriptor](instance.ID, objc.Sel("initWithDescriptorType:bytes:length:"), descriptorType, bytes, byteCount)
	rv.Autorelease()
	return rv
}


// Initializes a newly allocated instance as a descriptor with the specified descriptor type and data (from an instance of ).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/init(descriptorType:data:)
func NewAppleEventDescriptorWithDescriptorTypeData(descriptorType unsafe.Pointer, data IData) AppleEventDescriptor {
	instance := getAppleEventDescriptorClass().Alloc()
	rv := objc.Send[AppleEventDescriptor](instance.ID, objc.Sel("initWithDescriptorType:data:"), descriptorType, data)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/init(double:)
func NewAppleEventDescriptorWithDouble(doubleValue float64) AppleEventDescriptor {
	rv := objc.Send[AppleEventDescriptor](objc.ID(getAppleEventDescriptorClass().class), objc.Sel("descriptorWithDouble:"), doubleValue)
	return rv
}


// Creates a descriptor initialized with type that stores the specified enumerator data type value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/init(enumCode:)
func NewAppleEventDescriptorWithEnumCode(enumerator unsafe.Pointer) AppleEventDescriptor {
	rv := objc.Send[AppleEventDescriptor](objc.ID(getAppleEventDescriptorClass().class), objc.Sel("descriptorWithEnumCode:"), enumerator)
	return rv
}


// Initializes a newly allocated instance as a descriptor for an Apple event, initialized with the specified values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/init(eventClass:eventID:targetDescriptor:returnID:transactionID:)
func NewAppleEventDescriptorWithEventClassEventIDTargetDescriptorReturnIDTransactionID(eventClass unsafe.Pointer, eventID unsafe.Pointer, targetDescriptor IAppleEventDescriptor, returnID unsafe.Pointer, transactionID unsafe.Pointer) AppleEventDescriptor {
	instance := getAppleEventDescriptorClass().Alloc()
	rv := objc.Send[AppleEventDescriptor](instance.ID, objc.Sel("initWithEventClass:eventID:targetDescriptor:returnID:transactionID:"), eventClass, eventID, targetDescriptor, returnID, transactionID)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/init(fileURL:)
func NewAppleEventDescriptorWithFileURL(fileURL IURL) AppleEventDescriptor {
	rv := objc.Send[AppleEventDescriptor](objc.ID(getAppleEventDescriptorClass().class), objc.Sel("descriptorWithFileURL:"), fileURL)
	return rv
}


// Creates a descriptor initialized with Apple event type that stores the specified integer value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/init(int32:)
func NewAppleEventDescriptorWithInt32(signedInt unsafe.Pointer) AppleEventDescriptor {
	rv := objc.Send[AppleEventDescriptor](objc.ID(getAppleEventDescriptorClass().class), objc.Sel("descriptorWithInt32:"), signedInt)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/init(processIdentifier:)
func NewAppleEventDescriptorWithProcessIdentifier(processIdentifier unsafe.Pointer) AppleEventDescriptor {
	rv := objc.Send[AppleEventDescriptor](objc.ID(getAppleEventDescriptorClass().class), objc.Sel("descriptorWithProcessIdentifier:"), processIdentifier)
	return rv
}


// Creates a descriptor initialized with type that stores the text from the specified string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/init(string:)
func NewAppleEventDescriptorWithString(string_ string) AppleEventDescriptor {
	rv := objc.Send[AppleEventDescriptor](objc.ID(getAppleEventDescriptorClass().class), objc.Sel("descriptorWithString:"), objc.String(string_))
	return rv
}


// Creates a descriptor initialized with type that stores the specified type value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/init(typeCode:)
func NewAppleEventDescriptorWithTypeCode(typeCode unsafe.Pointer) AppleEventDescriptor {
	rv := objc.Send[AppleEventDescriptor](objc.ID(getAppleEventDescriptorClass().class), objc.Sel("descriptorWithTypeCode:"), typeCode)
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


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/currentProcess()
func (ac _AppleEventDescriptorClass) CurrentProcessDescriptor() AppleEventDescriptor {
	rv := objc.Send[AppleEventDescriptor](objc.ID(ac.class), objc.Sel("currentProcessDescriptor"))
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


// Creates a descriptor initialized with the specified event type that stores the specified data (from an instance of ).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/descriptorWithDescriptorType:data:
func (ac _AppleEventDescriptorClass) DescriptorWithDescriptorTypeData(descriptorType unsafe.Pointer, data IData) AppleEventDescriptor {
	rv := objc.Send[AppleEventDescriptor](objc.ID(ac.class), objc.Sel("descriptorWithDescriptorType:data:"), descriptorType, data)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/init(applicationURL:)
func (ac _AppleEventDescriptorClass) DescriptorWithApplicationURL(applicationURL IURL) AppleEventDescriptor {
	rv := objc.Send[AppleEventDescriptor](objc.ID(ac.class), objc.Sel("descriptorWithApplicationURL:"), applicationURL)
	return rv
}


// Creates a descriptor initialized with type that stores the specified Boolean value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/init(boolean:)
func (ac _AppleEventDescriptorClass) DescriptorWithBoolean(boolean unsafe.Pointer) AppleEventDescriptor {
	rv := objc.Send[AppleEventDescriptor](objc.ID(ac.class), objc.Sel("descriptorWithBoolean:"), boolean)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/init(bundleIdentifier:)
func (ac _AppleEventDescriptorClass) DescriptorWithBundleIdentifier(bundleIdentifier string) AppleEventDescriptor {
	rv := objc.Send[AppleEventDescriptor](objc.ID(ac.class), objc.Sel("descriptorWithBundleIdentifier:"), objc.String(bundleIdentifier))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/init(date:)
func (ac _AppleEventDescriptorClass) DescriptorWithDate(date IDate) AppleEventDescriptor {
	rv := objc.Send[AppleEventDescriptor](objc.ID(ac.class), objc.Sel("descriptorWithDate:"), date)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/init(double:)
func (ac _AppleEventDescriptorClass) DescriptorWithDouble(doubleValue float64) AppleEventDescriptor {
	rv := objc.Send[AppleEventDescriptor](objc.ID(ac.class), objc.Sel("descriptorWithDouble:"), doubleValue)
	return rv
}


// Creates a descriptor initialized with type that stores the specified enumerator data type value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/init(enumCode:)
func (ac _AppleEventDescriptorClass) DescriptorWithEnumCode(enumerator unsafe.Pointer) AppleEventDescriptor {
	rv := objc.Send[AppleEventDescriptor](objc.ID(ac.class), objc.Sel("descriptorWithEnumCode:"), enumerator)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/init(fileURL:)
func (ac _AppleEventDescriptorClass) DescriptorWithFileURL(fileURL IURL) AppleEventDescriptor {
	rv := objc.Send[AppleEventDescriptor](objc.ID(ac.class), objc.Sel("descriptorWithFileURL:"), fileURL)
	return rv
}


// Creates a descriptor initialized with Apple event type that stores the specified integer value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/init(int32:)
func (ac _AppleEventDescriptorClass) DescriptorWithInt32(signedInt unsafe.Pointer) AppleEventDescriptor {
	rv := objc.Send[AppleEventDescriptor](objc.ID(ac.class), objc.Sel("descriptorWithInt32:"), signedInt)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/init(processIdentifier:)
func (ac _AppleEventDescriptorClass) DescriptorWithProcessIdentifier(processIdentifier unsafe.Pointer) AppleEventDescriptor {
	rv := objc.Send[AppleEventDescriptor](objc.ID(ac.class), objc.Sel("descriptorWithProcessIdentifier:"), processIdentifier)
	return rv
}


// Creates a descriptor initialized with type that stores the text from the specified string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/init(string:)
func (ac _AppleEventDescriptorClass) DescriptorWithString(string_ string) AppleEventDescriptor {
	rv := objc.Send[AppleEventDescriptor](objc.ID(ac.class), objc.Sel("descriptorWithString:"), objc.String(string_))
	return rv
}


// Creates a descriptor initialized with type that stores the specified type value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/init(typeCode:)
func (ac _AppleEventDescriptorClass) DescriptorWithTypeCode(typeCode unsafe.Pointer) AppleEventDescriptor {
	rv := objc.Send[AppleEventDescriptor](objc.ID(ac.class), objc.Sel("descriptorWithTypeCode:"), typeCode)
	return rv
}


// Creates and initializes an empty list descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/list()
func (ac _AppleEventDescriptorClass) ListDescriptor() AppleEventDescriptor {
	rv := objc.Send[AppleEventDescriptor](objc.ID(ac.class), objc.Sel("listDescriptor"))
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


// Creates and initializes a descriptor for an Apple event record whose data has yet to be set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/record()
func (ac _AppleEventDescriptorClass) RecordDescriptor() AppleEventDescriptor {
	rv := objc.Send[AppleEventDescriptor](objc.ID(ac.class), objc.Sel("recordDescriptor"))
	return rv
}


// Returns the descriptor at the specified (one-based) position in the receiving descriptor list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/atIndex(_:)
func (a_ AppleEventDescriptor) DescriptorAtIndex(index int) AppleEventDescriptor {
	rv := objc.Send[AppleEventDescriptor](a_.ID, objc.Sel("descriptorAtIndex:"), index)
	return rv
}


// Returns a descriptor for the receiver’s Apple event attribute identified by the specified keyword.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/attributeDescriptor(forKeyword:)
func (a_ AppleEventDescriptor) AttributeDescriptorForKeyword(keyword unsafe.Pointer) AppleEventDescriptor {
	rv := objc.Send[AppleEventDescriptor](a_.ID, objc.Sel("attributeDescriptorForKeyword:"), keyword)
	return rv
}


// Returns a descriptor obtained by coercing the receiver to the specified type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/coerce(toDescriptorType:)
func (a_ AppleEventDescriptor) CoerceToDescriptorType(descriptorType unsafe.Pointer) AppleEventDescriptor {
	rv := objc.Send[AppleEventDescriptor](a_.ID, objc.Sel("coerceToDescriptorType:"), descriptorType)
	return rv
}


// Returns the receiver’s descriptor for the specified keyword.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/forKeyword(_:)
func (a_ AppleEventDescriptor) DescriptorForKeyword(keyword unsafe.Pointer) AppleEventDescriptor {
	rv := objc.Send[AppleEventDescriptor](a_.ID, objc.Sel("descriptorForKeyword:"), keyword)
	return rv
}


// Inserts a descriptor at the specified (one-based) position in the receiving descriptor list, replacing the existing descriptor, if any, at that position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/insert(_:at:)
func (a_ AppleEventDescriptor) InsertDescriptorAtIndex(descriptor IAppleEventDescriptor, index int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("insertDescriptor:atIndex:"), descriptor, index)
}


// Returns the keyword for the descriptor at the specified (one-based) position in the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/keywordForDescriptor(at:)
func (a_ AppleEventDescriptor) KeywordForDescriptorAtIndex(index int) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("keywordForDescriptorAtIndex:"), index)
	return rv
}


// Returns a descriptor for the receiver’s Apple event parameter identified by the specified keyword.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/paramDescriptor(forKeyword:)
func (a_ AppleEventDescriptor) ParamDescriptorForKeyword(keyword unsafe.Pointer) AppleEventDescriptor {
	rv := objc.Send[AppleEventDescriptor](a_.ID, objc.Sel("paramDescriptorForKeyword:"), keyword)
	return rv
}


// Removes the descriptor at the specified (one-based) position in the receiving descriptor list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/remove(at:)
func (a_ AppleEventDescriptor) RemoveDescriptorAtIndex(index int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("removeDescriptorAtIndex:"), index)
}


// Removes the receiver’s descriptor identified by the specified keyword.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/remove(withKeyword:)
func (a_ AppleEventDescriptor) RemoveDescriptorWithKeyword(keyword unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("removeDescriptorWithKeyword:"), keyword)
}


// Removes the receiver’s parameter descriptor identified by the specified keyword.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/removeParamDescriptor(withKeyword:)
func (a_ AppleEventDescriptor) RemoveParamDescriptorWithKeyword(keyword unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("removeParamDescriptorWithKeyword:"), keyword)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/sendEvent(options:timeout:)
func (a_ AppleEventDescriptor) SendEventWithOptionsTimeoutError(sendOptions NSAppleEventSendOptions, timeoutInSeconds TimeInterval, error_ IError) AppleEventDescriptor {
	rv := objc.Send[AppleEventDescriptor](a_.ID, objc.Sel("sendEventWithOptions:timeout:error:"), sendOptions, timeoutInSeconds, error_)
	return rv
}


// Adds a descriptor to the receiver as an attribute identified by the specified keyword.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/setAttribute(_:forKeyword:)
func (a_ AppleEventDescriptor) SetAttributeDescriptorForKeyword(descriptor IAppleEventDescriptor, keyword unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAttributeDescriptor:forKeyword:"), descriptor, keyword)
}


// Adds a descriptor, identified by a keyword, to the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/setDescriptor(_:forKeyword:)
func (a_ AppleEventDescriptor) SetDescriptorForKeyword(descriptor IAppleEventDescriptor, keyword unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDescriptor:forKeyword:"), descriptor, keyword)
}


// Adds a descriptor to the receiver as an Apple event parameter identified by the specified keyword.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/setParam(_:forKeyword:)
func (a_ AppleEventDescriptor) SetParamDescriptorForKeyword(descriptor IAppleEventDescriptor, keyword unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setParamDescriptor:forKeyword:"), descriptor, keyword)
}


// The structure encapsulated by the receiver, if it has one.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/aeDesc
func (a_ AppleEventDescriptor) AeDesc() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("aeDesc"))
	return rv
}


// The contents of the receiver as a Boolean value, coercing (to ) if necessary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/booleanValue
func (a_ AppleEventDescriptor) BooleanValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("booleanValue"))
	return rv
}


// The receiver’s data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/data
func (a_ AppleEventDescriptor) Data() NSData {
	rv := objc.Send[NSData](a_.ID, objc.Sel("data"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/dateValue
func (a_ AppleEventDescriptor) DateValue() NSDate {
	rv := objc.Send[NSDate](a_.ID, objc.Sel("dateValue"))
	return rv
}


// The descriptor type of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/descriptorType
func (a_ AppleEventDescriptor) DescriptorType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("descriptorType"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/doubleValue
func (a_ AppleEventDescriptor) DoubleValue() float64 {
	rv := objc.Send[float64](a_.ID, objc.Sel("doubleValue"))
	return rv
}


// The contents of the receiver as an enumeration type, coercing to if necessary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/enumCodeValue
func (a_ AppleEventDescriptor) EnumCodeValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("enumCodeValue"))
	return rv
}


// The event class for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/eventClass
func (a_ AppleEventDescriptor) EventClass() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("eventClass"))
	return rv
}


// The event ID for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/eventID
func (a_ AppleEventDescriptor) EventID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("eventID"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/fileURLValue
func (a_ AppleEventDescriptor) FileURLValue() URL {
	rv := objc.Send[URL](a_.ID, objc.Sel("fileURLValue"))
	return rv
}


// The contents of the receiver as an integer, coercing (to ) if necessary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/int32Value
func (a_ AppleEventDescriptor) Int32Value() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("int32Value"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/isRecordDescriptor
func (a_ AppleEventDescriptor) IsRecordDescriptor() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isRecordDescriptor"))
	return rv
}


// The number of descriptors in the receiver’s descriptor list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/numberOfItems
func (a_ AppleEventDescriptor) NumberOfItems() int {
	rv := objc.Send[int](a_.ID, objc.Sel("numberOfItems"))
	return rv
}


// The receiver’s return ID (the ID for a reply Apple event).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/returnID
func (a_ AppleEventDescriptor) ReturnID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("returnID"))
	return rv
}


// The contents of the receiver as a Unicode text string, coercing to if necessary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/stringValue
func (a_ AppleEventDescriptor) StringValue() string {
	rv := objc.Send[string](a_.ID, objc.Sel("stringValue"))
	return rv
}


// The receiver’s transaction ID, if any.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/transactionID
func (a_ AppleEventDescriptor) TransactionID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("transactionID"))
	return rv
}


// The contents of the receiver as a type, coercing to if necessary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/typeCodeValue
func (a_ AppleEventDescriptor) TypeCodeValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("typeCodeValue"))
	return rv
}


