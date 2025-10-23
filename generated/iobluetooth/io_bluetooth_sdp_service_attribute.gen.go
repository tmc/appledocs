// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [BluetoothSDPServiceAttribute] class.
var (
	BluetoothSDPServiceAttributeClass     _BluetoothSDPServiceAttributeClass
	BluetoothSDPServiceAttributeClassOnce sync.Once
)

func getBluetoothSDPServiceAttributeClass() _BluetoothSDPServiceAttributeClass {
	BluetoothSDPServiceAttributeClassOnce.Do(func() {
		BluetoothSDPServiceAttributeClass = _BluetoothSDPServiceAttributeClass{objc.GetClass("IOBluetoothSDPServiceAttribute")}
	})
	return BluetoothSDPServiceAttributeClass
}

type _BluetoothSDPServiceAttributeClass struct {
	class objc.Class
}

// An interface definition for the [BluetoothSDPServiceAttribute] class.
type IBluetoothSDPServiceAttribute interface {
	objectivec.IObject
	// properties:
	// methods:
	GetDataElement() IBluetoothSDPDataElement
	GetAttributeID() BluetoothSDPServiceAttributeID /* typedef */
	GetIDDataElement() IBluetoothSDPDataElement
}

// IOBluetoothSDPServiceAttribute represents a single SDP service attribute.
//
// A service attribute contains two components: an attribute ID and a data element.


// IOBluetoothSDPServiceAttribute represents a single SDP service attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceAttribute
type BluetoothSDPServiceAttribute struct {
	objectivec.Object
}

// BluetoothSDPServiceAttributeFrom constructs a [BluetoothSDPServiceAttribute] from an unsafe.Pointer.
//
// IOBluetoothSDPServiceAttribute represents a single SDP service attribute.
func BluetoothSDPServiceAttributeFrom(ptr unsafe.Pointer) BluetoothSDPServiceAttribute {
	return BluetoothSDPServiceAttribute{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (bc _BluetoothSDPServiceAttributeClass) Alloc() BluetoothSDPServiceAttribute {
	rv := objc.Send[BluetoothSDPServiceAttribute](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (bc _BluetoothSDPServiceAttributeClass) New() BluetoothSDPServiceAttribute {
	rv := objc.Send[BluetoothSDPServiceAttribute](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BluetoothSDPServiceAttribute) Init() BluetoothSDPServiceAttribute {
	rv := objc.Send[BluetoothSDPServiceAttribute](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BluetoothSDPServiceAttribute) Autorelease() BluetoothSDPServiceAttribute {
	rv := objc.Send[BluetoothSDPServiceAttribute](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBluetoothSDPServiceAttribute creates a new BluetoothSDPServiceAttribute instance.
func NewBluetoothSDPServiceAttribute() BluetoothSDPServiceAttribute {
	return getBluetoothSDPServiceAttributeClass().New()
}



// Initializes a new service attribute with the given ID and data element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceAttribute/init(id:attributeElement:)
func NewBluetoothSDPServiceAttributeWithIDAttributeElement(newAttributeID BluetoothSDPServiceAttributeID /* typedef */, attributeElement BluetoothSDPDataElement /* already interface */) BluetoothSDPServiceAttribute {
	instance := getBluetoothSDPServiceAttributeClass().Alloc()
	rv := objc.Send[BluetoothSDPServiceAttribute](instance.ID, objc.Sel("initWithID:attributeElement:"), newAttributeID, attributeElement)
	rv.Autorelease()
	return rv
}


// Initializes a new service attribute with the given ID and element value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceAttribute/init(id:attributeElementValue:)
func NewBluetoothSDPServiceAttributeWithIDAttributeElementValue(newAttributeID BluetoothSDPServiceAttributeID /* typedef */, attributeElementValue objectivec.IObject) BluetoothSDPServiceAttribute {
	instance := getBluetoothSDPServiceAttributeClass().Alloc()
	rv := objc.Send[BluetoothSDPServiceAttribute](instance.ID, objc.Sel("initWithID:attributeElementValue:"), newAttributeID, attributeElementValue)
	rv.Autorelease()
	return rv
}



// Creates a new service attribute with the given ID and data element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceAttribute/withID(_:attributeElement:)
func (bc _BluetoothSDPServiceAttributeClass) WithIDAttributeElement(newAttributeID BluetoothSDPServiceAttributeID /* typedef */, attributeElement BluetoothSDPDataElement /* already interface */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("withID:attributeElement:"), newAttributeID, attributeElement)
	return rv
}


// Creates a new service attribute with the given ID and element value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceAttribute/withID(_:attributeElementValue:)
func (bc _BluetoothSDPServiceAttributeClass) WithIDAttributeElementValue(newAttributeID BluetoothSDPServiceAttributeID /* typedef */, attributeElementValue objectivec.IObject) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("withID:attributeElementValue:"), newAttributeID, attributeElementValue)
	return rv
}


// Returns the data element for the target service attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceAttribute/getDataElement()
func (b_ BluetoothSDPServiceAttribute) GetDataElement() IBluetoothSDPDataElement {
	rv := objc.Send[BluetoothSDPDataElement](b_.ID, objc.Sel("getDataElement"))
	return rv
}


// Returns the attribute ID for the target service attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceAttribute/getID()
func (b_ BluetoothSDPServiceAttribute) GetAttributeID() BluetoothSDPServiceAttributeID /* typedef */ {
	rv := objc.Send[BluetoothSDPServiceAttributeID](b_.ID, objc.Sel("getAttributeID"))
	return rv
}


// Returns the data element representing the attribute ID for the target service attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceAttribute/getIDDataElement()
func (b_ BluetoothSDPServiceAttribute) GetIDDataElement() IBluetoothSDPDataElement {
	rv := objc.Send[BluetoothSDPDataElement](b_.ID, objc.Sel("getIDDataElement"))
	return rv
}


