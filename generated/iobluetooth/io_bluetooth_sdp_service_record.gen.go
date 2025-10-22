// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [BluetoothSDPServiceRecord] class.
var (
	BluetoothSDPServiceRecordClass     _BluetoothSDPServiceRecordClass
	BluetoothSDPServiceRecordClassOnce sync.Once
)

func getBluetoothSDPServiceRecordClass() _BluetoothSDPServiceRecordClass {
	BluetoothSDPServiceRecordClassOnce.Do(func() {
		BluetoothSDPServiceRecordClass = _BluetoothSDPServiceRecordClass{objc.GetClass("IOBluetoothSDPServiceRecord")}
	})
	return BluetoothSDPServiceRecordClass
}

type _BluetoothSDPServiceRecordClass struct {
	class objc.Class
}

// An interface definition for the [BluetoothSDPServiceRecord] class.
type IBluetoothSDPServiceRecord interface {
	objectivec.IObject
	GetAttributeDataElement(attributeID IBluetoothSDPServiceAttributeID) BluetoothSDPDataElement
	GetAttributes() foundation.Dictionary
	GetDevice() BluetoothDevice
	GetServiceRecordHandle(outServiceRecordHandle IBluetoothSDPServiceRecordHandle) unsafe.Pointer
	GetL2CAPPSM(outPSM IBluetoothL2CAPPSM) unsafe.Pointer
	GetRFCOMMChannelID(rfcommChannelID IBluetoothRFCOMMChannelID) unsafe.Pointer
	GetSDPServiceRecordRef() BluetoothSDPServiceRecordRef
	GetServiceName() foundation.String
	HandsFreeSupportedFeatures() unsafe.Pointer
	HasServiceFromArray(array objectivec.IObject) bool
	MatchesSearchArray(searchArray objectivec.IObject) bool
	MatchesUUID16(uuid16 IBluetoothSDPUUID16) bool
	MatchesUUIDArray(uuidArray objectivec.IObject) bool
	RemoveServiceRecord() unsafe.Pointer
	Attributes() objc.ID
	Device() IOBluetoothDevice
	SortedAttributes() objc.ID
}

// An instance of this class represents a single SDP service record.
//
// As a service record, an instance of this class has an NSDictionary of service attributes. It also has a link to the IOBluetoothDevice that the service belongs to. The service dictionary is keyed off of the attribute ID of each attribute represented as an NSNumber.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceRecord
type BluetoothSDPServiceRecord struct {
	objectivec.Object
}

// BluetoothSDPServiceRecordFrom constructs a [BluetoothSDPServiceRecord] from an unsafe.Pointer.
//
// An instance of this class represents a single SDP service record.
func BluetoothSDPServiceRecordFrom(ptr unsafe.Pointer) BluetoothSDPServiceRecord {
	return BluetoothSDPServiceRecord{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (bc _BluetoothSDPServiceRecordClass) Alloc() BluetoothSDPServiceRecord {
	rv := objc.Send[BluetoothSDPServiceRecord](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (bc _BluetoothSDPServiceRecordClass) New() BluetoothSDPServiceRecord {
	rv := objc.Send[BluetoothSDPServiceRecord](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BluetoothSDPServiceRecord) Init() BluetoothSDPServiceRecord {
	rv := objc.Send[BluetoothSDPServiceRecord](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BluetoothSDPServiceRecord) Autorelease() BluetoothSDPServiceRecord {
	rv := objc.Send[BluetoothSDPServiceRecord](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBluetoothSDPServiceRecord creates a new BluetoothSDPServiceRecord instance.
func NewBluetoothSDPServiceRecord() BluetoothSDPServiceRecord {
	return getBluetoothSDPServiceRecordClass().New()
}




// Returns an initialized IOBluetoothSDPServiceRecord * with the attributes specified in the provided service dictionary. Provide a pointer to an IOBlueotothDevice if you wish to associate the record to a specific IOBluetoothDevice.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceRecord/init(serviceDictionary:device:)
func NewBluetoothSDPServiceRecordWithServiceDictionaryDevice(serviceDict objectivec.IObject, device IOBluetoothDevice) BluetoothSDPServiceRecord {
	instance := getBluetoothSDPServiceRecordClass().Alloc()
	rv := objc.Send[BluetoothSDPServiceRecord](instance.ID, objc.Sel("initWithServiceDictionary:device:"), serviceDict, device)
	rv.Autorelease()
	return rv
}


// Adds a service to the local SDP server.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceRecord/publishedServiceRecord(with:)
func (bc _BluetoothSDPServiceRecordClass) PublishedServiceRecordWithDictionary(serviceDict objectivec.IObject) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("publishedServiceRecordWithDictionary:"), serviceDict)
	return rv
}

// Method call to convert an IOBluetoothSDPServiceRecordRef into an IOBluetoothSDPServiceRecord *.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceRecord/withSDPServiceRecordRef(_:)
func (bc _BluetoothSDPServiceRecordClass) WithSDPServiceRecordRef(sdpServiceRecordRef IBluetoothSDPServiceRecordRef) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("withSDPServiceRecordRef:"), sdpServiceRecordRef)
	return rv
}

// Returns an IOBluetoothSDPServiceRecord * with the attributes specified in the provided service dictionary. Provide a pointer to an IOBlueotothDevice if you wish to associate the record to a specific IOBluetoothDevice.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceRecord/withServiceDictionary(_:device:)
func (bc _BluetoothSDPServiceRecordClass) WithServiceDictionaryDevice(serviceDict objectivec.IObject, device IOBluetoothDevice) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("withServiceDictionary:device:"), serviceDict, device)
	return rv
}

// Returns the data element for the given attribute ID in the target service.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceRecord/getAttributeDataElement(_:)
func (b_ BluetoothSDPServiceRecord) GetAttributeDataElement(attributeID IBluetoothSDPServiceAttributeID) BluetoothSDPDataElement {
	rv := objc.Send[BluetoothSDPDataElement](b_.ID, objc.Sel("getAttributeDataElement:"), attributeID)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceRecord/getAttributes
func (b_ BluetoothSDPServiceRecord) GetAttributes() foundation.Dictionary {
	rv := objc.Send[foundation.Dictionary](b_.ID, objc.Sel("getAttributes"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceRecord/getDevice
func (b_ BluetoothSDPServiceRecord) GetDevice() BluetoothDevice {
	rv := objc.Send[BluetoothDevice](b_.ID, objc.Sel("getDevice"))
	return rv
}

// Allows the discovery of the service record handle assigned to the service.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceRecord/getHandle(_:)
func (b_ BluetoothSDPServiceRecord) GetServiceRecordHandle(outServiceRecordHandle IBluetoothSDPServiceRecordHandle) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("getServiceRecordHandle:"), outServiceRecordHandle)
	return rv
}

// Allows the discovery of the L2CAP PSM assigned to the service.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceRecord/getL2CAPPSM(_:)
func (b_ BluetoothSDPServiceRecord) GetL2CAPPSM(outPSM IBluetoothL2CAPPSM) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("getL2CAPPSM:"), outPSM)
	return rv
}

// Allows the discovery of the RFCOMM channel ID assigned to the service.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceRecord/getRFCOMMChannelID(_:)
func (b_ BluetoothSDPServiceRecord) GetRFCOMMChannelID(rfcommChannelID IBluetoothRFCOMMChannelID) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("getRFCOMMChannelID:"), rfcommChannelID)
	return rv
}

// Returns an IOBluetoothSDPServiceRecordRef representation of the target IOBluetoothSDPServiceRecord object.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceRecord/getRef()
func (b_ BluetoothSDPServiceRecord) GetSDPServiceRecordRef() BluetoothSDPServiceRecordRef {
	rv := objc.Send[BluetoothSDPServiceRecordRef](b_.ID, objc.Sel("getSDPServiceRecordRef"))
	return rv
}

// Returns the name of the service.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceRecord/getServiceName()
func (b_ BluetoothSDPServiceRecord) GetServiceName() foundation.String {
	rv := objc.Send[foundation.String](b_.ID, objc.Sel("getServiceName"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceRecord/handsFreeSupportedFeatures()
func (b_ BluetoothSDPServiceRecord) HandsFreeSupportedFeatures() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("handsFreeSupportedFeatures"))
	return rv
}

// Returns TRUE if any one of the UUIDs in the given array is found in the target service.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceRecord/hasService(from:)
func (b_ BluetoothSDPServiceRecord) HasServiceFromArray(array objectivec.IObject) bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("hasServiceFromArray:"), array)
	return rv
}

// Returns TRUE any of the UUID arrays in the search array match the target service.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceRecord/matchesSearch(_:)
func (b_ BluetoothSDPServiceRecord) MatchesSearchArray(searchArray objectivec.IObject) bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("matchesSearchArray:"), searchArray)
	return rv
}

// Returns TRUE the UUID16 is found in the target service.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceRecord/matchesUUID16(_:)
func (b_ BluetoothSDPServiceRecord) MatchesUUID16(uuid16 IBluetoothSDPUUID16) bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("matchesUUID16:"), uuid16)
	return rv
}

// Returns TRUE if ALL of the UUIDs in the given array is found in the target service.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceRecord/matchesUUIDArray(_:)
func (b_ BluetoothSDPServiceRecord) MatchesUUIDArray(uuidArray objectivec.IObject) bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("matchesUUIDArray:"), uuidArray)
	return rv
}

// Removes the service from the local SDP server.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceRecord/remove()
func (b_ BluetoothSDPServiceRecord) RemoveServiceRecord() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("removeServiceRecord"))
	return rv
}

// Returns an NSDictionary containing the attributes for the service.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceRecord/attributes
func (b_ BluetoothSDPServiceRecord) Attributes() objc.ID {
	rv := objc.Send[objc.ID](b_.ID, objc.Sel("attributes"))
	return rv
}

// Returns the IOBluetoothDevice that the target service belongs to.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceRecord/device
func (b_ BluetoothSDPServiceRecord) Device() IOBluetoothDevice {
	rv := objc.Send[IOBluetoothDevice](b_.ID, objc.Sel("device"))
	return rv
}

// Returns a sorted array of SDP attributes
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceRecord/sortedAttributes-swift.property
func (b_ BluetoothSDPServiceRecord) SortedAttributes() objc.ID {
	rv := objc.Send[objc.ID](b_.ID, objc.Sel("sortedAttributes"))
	return rv
}


