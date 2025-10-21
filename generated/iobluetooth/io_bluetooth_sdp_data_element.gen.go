// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [BluetoothSDPDataElement] class.
var (
	BluetoothSDPDataElementClass     _BluetoothSDPDataElementClass
	BluetoothSDPDataElementClassOnce sync.Once
)

func getBluetoothSDPDataElementClass() _BluetoothSDPDataElementClass {
	BluetoothSDPDataElementClassOnce.Do(func() {
		BluetoothSDPDataElementClass = _BluetoothSDPDataElementClass{objc.GetClass("IOBluetoothSDPDataElement")}
	})
	return BluetoothSDPDataElementClass
}

type _BluetoothSDPDataElementClass struct {
	class objc.Class
}

// An interface definition for the [BluetoothSDPDataElement] class.
type IBluetoothSDPDataElement interface {
	objectivec.IObject
	ContainsDataElement(dataElement unsafe.Pointer) bool
	ContainsValue(cmpValue unsafe.Pointer) bool
	GetArrayValue() unsafe.Pointer
	GetDataValue() unsafe.Pointer
	GetNumberValue() unsafe.Pointer
	GetSDPDataElementRef() unsafe.Pointer
	GetSize() unsafe.Pointer
	GetSizeDescriptor() unsafe.Pointer
	GetStringValue() unsafe.Pointer
	GetTypeDescriptor() unsafe.Pointer
	GetUUIDValue() unsafe.Pointer
	GetValue() unsafe.Pointer
}

// An instance of this class represents a single SDP data element as defined by the Bluetooth SDP spec.
//
// The data types described by the spec have been mapped onto the base Foundation classes NSNumber, NSArray, NSData as well as IOBluetoothSDPUUID. The number and boolean types (type descriptor 1, 2 and 5) are represented as NSNumber objects with the exception of 128-bit numbers which are represented as NSData objects in their raw format. The UUID type (type descriptor 3) is represented by IOBluetoothSDPUUID. The string and URL types (type descriptor 4 and 8) are represented by NSString. The sequence types (type descriptor 6 and 7) are represented by NSArray. Typically, you will not need to create an IOBluetoothSDPDataElement directly, the system will do that automatically for both client and server operations. However, the current API for adding SDP services to the system does allow the use of an NSDictionary based format for creating new services. The purpose for that is to allow a service to be built up completely in a text file (a plist for example) and then easily imported into an app and added to the system without a lot of tedious code to build up the entire SDP service record. The basis for that NSDictionary structure comes from the IOBluetoothSDPDataElement. At its simplest, a data element is made up of three parts: the type descriptor, the size (from which the size descriptor is generated) and the actual value. To provide a complete representation of a data element, an NSDictionary with three entries can be used. Each of the three entries has a key/value pair representing one of the three attributes of a data element. The first key/value pair has a key ‘DataElementType’ that contains a number value with the actual type descriptor for the data element. The second pair has a key ‘DataElementSize’ that contains the actual size of the element in bytes. The size descriptor will be calculated based on the size and type of the element. The third pair is the value itself whose key is ‘DataElementValue’ and whose type corresponds to the type mapping above. In addition to this complete description of a data element, their are some shortcuts that can be used for some of the common types and sizes. If the ‘DataElementType’ value is one of the numeric types (1, 2), the ‘DataElementValue’ can be an NSData instead of an NSNumber. In that case, the numeric data is taken in network byte order (MSB first). Additionally, the ‘DataElementSize’ parameter may be omitted and the size will be taken from the length of the data object. If the ‘DataElementType’ value is the nil type (0), no ‘DataElementSize’ or ‘DataElementValue’ entries are needed. If the ‘DataElementType’ value is any of the other types, the ‘DataElementSize’ entry is not needed since the size will be taken directly from the value (data, array, string). In the case where the element is an unsigned, 32-bit integer (type descriptor 1, size descriptor 4), the value itself may simply be a number (instead of a dictionary as in the previous examples). In the case where the element is a UUID (type descriptor 3), the value itself may be a data object. The UUID type will be inferred and the size taken from the length of the data object. In the case where the element is a text string (type descriptor 4), the value may be a string object. The text string type will be inferred and the size taken from the length of the string. In the case where the element is a data element sequence, the value may be an array object. The type will be inferred and the size taken from the length of the array. Additionally, the array must contain sub-elements that will be parsed out individually.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPDataElement
type BluetoothSDPDataElement struct {
	objectivec.Object
}

// BluetoothSDPDataElementFrom constructs a [BluetoothSDPDataElement] from an unsafe.Pointer.
//
// An instance of this class represents a single SDP data element as defined by the Bluetooth SDP spec.
func BluetoothSDPDataElementFrom(ptr unsafe.Pointer) BluetoothSDPDataElement {
	return BluetoothSDPDataElement{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (bc _BluetoothSDPDataElementClass) Alloc() BluetoothSDPDataElement {
	rv := objc.Send[BluetoothSDPDataElement](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (bc _BluetoothSDPDataElementClass) New() BluetoothSDPDataElement {
	rv := objc.Send[BluetoothSDPDataElement](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BluetoothSDPDataElement) Init() BluetoothSDPDataElement {
	rv := objc.Send[BluetoothSDPDataElement](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BluetoothSDPDataElement) Autorelease() BluetoothSDPDataElement {
	rv := objc.Send[BluetoothSDPDataElement](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBluetoothSDPDataElement creates a new BluetoothSDPDataElement instance.
func NewBluetoothSDPDataElement() BluetoothSDPDataElement {
	return getBluetoothSDPDataElementClass().New()
}


// Initializes a new IOBluetoothSDPDataElement with the given value.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPDataElement/init(elementValue:)
func NewBluetoothSDPDataElementWithElementValue(element unsafe.Pointer) BluetoothSDPDataElement {
	instance := getBluetoothSDPDataElementClass().Alloc()
	rv := objc.Send[BluetoothSDPDataElement](instance.ID, objc.Sel("initWithElementValue:"), element)
	rv.Autorelease()
	return rv
}

// Initializes a new IOBluetoothSDPDataElement with the given attributes.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPDataElement/init(type:sizeDescriptor:size:value:)
func NewBluetoothSDPDataElementWithTypeSizeDescriptorSizeValue(newType unsafe.Pointer, newSizeDescriptor unsafe.Pointer, newSize unsafe.Pointer, newValue unsafe.Pointer) BluetoothSDPDataElement {
	instance := getBluetoothSDPDataElementClass().Alloc()
	rv := objc.Send[BluetoothSDPDataElement](instance.ID, objc.Sel("initWithType:sizeDescriptor:size:value:"), newType, newSizeDescriptor, newSize, newValue)
	rv.Autorelease()
	return rv
}


// Creates a new IOBluetoothSDPDataElement with the given value.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPDataElement/withElementValue(_:)
func (bc _BluetoothSDPDataElementClass) WithElementValue(element unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("withElementValue:"), element)
	return rv
}

// Method call to convert an IOBluetoothSDPDataElementRef into an IOBluetoothSDPDataElement *.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPDataElement/withSDPDataElementRef(_:)
func (bc _BluetoothSDPDataElementClass) WithSDPDataElementRef(sdpDataElementRef unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("withSDPDataElementRef:"), sdpDataElementRef)
	return rv
}

// Creates a new IOBluetoothSDPDataElement with the given attributes.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPDataElement/withType(_:sizeDescriptor:size:value:)
func (bc _BluetoothSDPDataElementClass) WithTypeSizeDescriptorSizeValue(type_ unsafe.Pointer, newSizeDescriptor unsafe.Pointer, newSize unsafe.Pointer, newValue unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("withType:sizeDescriptor:size:value:"), type_, newSizeDescriptor, newSize, newValue)
	return rv
}

// Checks to see if the target data element is the same as the dataElement parameter or if it contains the dataElement parameter (if its a sequence type).
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPDataElement/contains(_:)
func (b_ BluetoothSDPDataElement) ContainsDataElement(dataElement unsafe.Pointer) bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("containsDataElement:"), dataElement)
	return rv
}

// Checks to see if the target data element’s value is the same as the value parameter or if it contains the value parameter.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPDataElement/containsValue(_:)
func (b_ BluetoothSDPDataElement) ContainsValue(cmpValue unsafe.Pointer) bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("containsValue:"), cmpValue)
	return rv
}

// If the data element is represented by an array object, it returns the value as an NSArray.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPDataElement/getArrayValue()
func (b_ BluetoothSDPDataElement) GetArrayValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("getArrayValue"))
	return rv
}

// If the data element is represented by a data object, it returns the value as an NSData.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPDataElement/getDataValue()
func (b_ BluetoothSDPDataElement) GetDataValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("getDataValue"))
	return rv
}

// If the data element is represented by a number, it returns the value as an NSNumber.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPDataElement/getNumberValue()
func (b_ BluetoothSDPDataElement) GetNumberValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("getNumberValue"))
	return rv
}

// Returns an IOBluetoothSDPDataElementRef representation of the target IOBluetoothSDPDataElement object.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPDataElement/getRef()
func (b_ BluetoothSDPDataElement) GetSDPDataElementRef() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("getSDPDataElementRef"))
	return rv
}

// Returns the size in bytes of the target data element.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPDataElement/getSize()
func (b_ BluetoothSDPDataElement) GetSize() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("getSize"))
	return rv
}

// Returns the SDP spec defined data element size descriptor for the target data element.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPDataElement/getSizeDescriptor()
func (b_ BluetoothSDPDataElement) GetSizeDescriptor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("getSizeDescriptor"))
	return rv
}

// If the data element is represented by a string object, it returns the value as an NSString.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPDataElement/getStringValue()
func (b_ BluetoothSDPDataElement) GetStringValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("getStringValue"))
	return rv
}

// Returns the SDP spec defined data element type descriptor for the target data element.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPDataElement/getTypeDescriptor()
func (b_ BluetoothSDPDataElement) GetTypeDescriptor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("getTypeDescriptor"))
	return rv
}

// If the data element is a UUID (type 3), it returns the value as an IOBluetoothSDPUUID.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPDataElement/getUUIDValue()
func (b_ BluetoothSDPDataElement) GetUUIDValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("getUUIDValue"))
	return rv
}

// Returns the object value of the data element.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPDataElement/getValue()
func (b_ BluetoothSDPDataElement) GetValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("getValue"))
	return rv
}


