// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class IOBluetoothSDPServiceRecord */


/* debug [class_header]: Header for IOBluetoothSDPServiceRecord */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for BluetoothSDPServiceRecord */
// An interface definition for the [BluetoothSDPServiceRecord] class.
type IBluetoothSDPServiceRecord interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for BluetoothSDPServiceRecord */
	// properties:
	Attributes() objc.IObject /* cross-framework: NSDictionary */
	Device() IOBluetoothDevice
	SortedAttributes() objc.IObject /* cross-framework: NSArray */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for BluetoothSDPServiceRecord */
	// methods:
	GetAttributeDataElement(attributeID BluetoothSDPServiceAttributeID /* typedef */) IBluetoothSDPDataElement
	GetServiceRecordHandle(outServiceRecordHandle BluetoothSDPServiceRecordHandle /* typedef */) int
	GetL2CAPPSM(outPSM BluetoothL2CAPPSM /* typedef */) int
	GetSDPServiceRecordRef() BluetoothSDPServiceRecordRef /* typedef */
	GetRFCOMMChannelID(rfcommChannelID BluetoothRFCOMMChannelID /* typedef */) int
	GetServiceName() foundation.String
	HandsFreeSupportedFeatures() uint16 /* not a class type */
	HasServiceFromArray(array objc.IObject /* cross-framework: NSArray */) bool
	MatchesSearchArray(searchArray objc.IObject /* cross-framework: NSArray */) bool
	MatchesUUID16(uuid16 BluetoothSDPUUID16 /* typedef */) bool
	MatchesUUIDArray(uuidArray objc.IObject /* cross-framework: NSArray */) bool
	RemoveServiceRecord() int
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for BluetoothSDPServiceRecord */
// Alloc allocates a new instance without initialization.
func (bc _BluetoothSDPServiceRecordClass) Alloc() BluetoothSDPServiceRecord {
	rv := objc.Send[BluetoothSDPServiceRecord](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for BluetoothSDPServiceRecord */
// An instance of this class represents a single SDP service record.
//
// As a service record, an instance of this class has an NSDictionary of service attributes. It also has a link to the IOBluetoothDevice that the service belongs to. The service dictionary is keyed off of the attribute ID of each attribute represented as an NSNumber.


// An instance of this class represents a single SDP service record.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for BluetoothSDPServiceRecord */

// Returns an initialized IOBluetoothSDPServiceRecord * with the attributes specified in the provided service dictionary. Provide a pointer to an IOBlueotothDevice if you wish to associate the record to a specific IOBluetoothDevice.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceRecord/init(serviceDictionary:device:)
func NewBluetoothSDPServiceRecordWithServiceDictionaryDevice(serviceDict objc.IObject /* cross-framework: NSDictionary */, device IOBluetoothDevice) BluetoothSDPServiceRecord {
	instance := getBluetoothSDPServiceRecordClass().Alloc()
	rv := objc.Send[BluetoothSDPServiceRecord](instance.ID, objc.Sel("initWithServiceDictionary:device:"), serviceDict, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewBluetoothSDPServiceRecordWithServiceDictionaryDevice */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for BluetoothSDPServiceRecord */

// Adds a service to the local SDP server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceRecord/publishedServiceRecord(with:)
func (bc _BluetoothSDPServiceRecordClass) PublishedServiceRecordWithDictionary(serviceDict objc.IObject /* cross-framework: NSDictionary */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("publishedServiceRecordWithDictionary:"), serviceDict)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PublishedServiceRecordWithDictionary) */


// Method call to convert an IOBluetoothSDPServiceRecordRef into an IOBluetoothSDPServiceRecord *.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceRecord/withSDPServiceRecordRef(_:)
func (bc _BluetoothSDPServiceRecordClass) WithSDPServiceRecordRef(sdpServiceRecordRef BluetoothSDPServiceRecordRef /* typedef */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("withSDPServiceRecordRef:"), sdpServiceRecordRef)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=WithSDPServiceRecordRef) */


// Returns an IOBluetoothSDPServiceRecord * with the attributes specified in the provided service dictionary. Provide a pointer to an IOBlueotothDevice if you wish to associate the record to a specific IOBluetoothDevice.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceRecord/withServiceDictionary(_:device:)
func (bc _BluetoothSDPServiceRecordClass) WithServiceDictionaryDevice(serviceDict objc.IObject /* cross-framework: NSDictionary */, device IOBluetoothDevice) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("withServiceDictionary:device:"), serviceDict, device)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=WithServiceDictionaryDevice) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for BluetoothSDPServiceRecord */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for BluetoothSDPServiceRecord */

// Returns the data element for the given attribute ID in the target service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceRecord/getAttributeDataElement(_:)
func (b_ BluetoothSDPServiceRecord) GetAttributeDataElement(attributeID BluetoothSDPServiceAttributeID /* typedef */) IBluetoothSDPDataElement {
	rv := objc.Send[BluetoothSDPDataElement](b_.ID, objc.Sel("getAttributeDataElement:"), attributeID)
	return rv
}/* debug [instance_methods/method]: GetAttributeDataElement */


// Allows the discovery of the service record handle assigned to the service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceRecord/getHandle(_:)
func (b_ BluetoothSDPServiceRecord) GetServiceRecordHandle(outServiceRecordHandle BluetoothSDPServiceRecordHandle /* typedef */) int {
	rv := objc.Send[int](b_.ID, objc.Sel("getServiceRecordHandle:"), outServiceRecordHandle)
	return rv
}/* debug [instance_methods/method]: GetServiceRecordHandle */


// Allows the discovery of the L2CAP PSM assigned to the service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceRecord/getL2CAPPSM(_:)
func (b_ BluetoothSDPServiceRecord) GetL2CAPPSM(outPSM BluetoothL2CAPPSM /* typedef */) int {
	rv := objc.Send[int](b_.ID, objc.Sel("getL2CAPPSM:"), outPSM)
	return rv
}/* debug [instance_methods/method]: GetL2CAPPSM */


// Returns an IOBluetoothSDPServiceRecordRef representation of the target IOBluetoothSDPServiceRecord object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceRecord/getRef()
func (b_ BluetoothSDPServiceRecord) GetSDPServiceRecordRef() BluetoothSDPServiceRecordRef /* typedef */ {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("getSDPServiceRecordRef"))
	return rv
}/* debug [instance_methods/method]: GetSDPServiceRecordRef */


// Allows the discovery of the RFCOMM channel ID assigned to the service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceRecord/getRFCOMMChannelID(_:)
func (b_ BluetoothSDPServiceRecord) GetRFCOMMChannelID(rfcommChannelID BluetoothRFCOMMChannelID /* typedef */) int {
	rv := objc.Send[int](b_.ID, objc.Sel("getRFCOMMChannelID:"), rfcommChannelID)
	return rv
}/* debug [instance_methods/method]: GetRFCOMMChannelID */


// Returns the name of the service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceRecord/getServiceName()
func (b_ BluetoothSDPServiceRecord) GetServiceName() foundation.String {
	rv := objc.Send[foundation.String](b_.ID, objc.Sel("getServiceName"))
	return rv
}/* debug [instance_methods/method]: GetServiceName */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceRecord/handsFreeSupportedFeatures()
func (b_ BluetoothSDPServiceRecord) HandsFreeSupportedFeatures() uint16 /* not a class type */ {
	rv := objc.Send[uint16](b_.ID, objc.Sel("handsFreeSupportedFeatures"))
	return rv
}/* debug [instance_methods/method]: HandsFreeSupportedFeatures */


// Returns TRUE if any one of the UUIDs in the given array is found in the target service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceRecord/hasService(from:)
func (b_ BluetoothSDPServiceRecord) HasServiceFromArray(array objc.IObject /* cross-framework: NSArray */) bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("hasServiceFromArray:"), array)
	return rv
}/* debug [instance_methods/method]: HasServiceFromArray */


// Returns TRUE any of the UUID arrays in the search array match the target service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceRecord/matchesSearch(_:)
func (b_ BluetoothSDPServiceRecord) MatchesSearchArray(searchArray objc.IObject /* cross-framework: NSArray */) bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("matchesSearchArray:"), searchArray)
	return rv
}/* debug [instance_methods/method]: MatchesSearchArray */


// Returns TRUE the UUID16 is found in the target service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceRecord/matchesUUID16(_:)
func (b_ BluetoothSDPServiceRecord) MatchesUUID16(uuid16 BluetoothSDPUUID16 /* typedef */) bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("matchesUUID16:"), uuid16)
	return rv
}/* debug [instance_methods/method]: MatchesUUID16 */


// Returns TRUE if ALL of the UUIDs in the given array is found in the target service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceRecord/matchesUUIDArray(_:)
func (b_ BluetoothSDPServiceRecord) MatchesUUIDArray(uuidArray objc.IObject /* cross-framework: NSArray */) bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("matchesUUIDArray:"), uuidArray)
	return rv
}/* debug [instance_methods/method]: MatchesUUIDArray */


// Removes the service from the local SDP server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceRecord/remove()
func (b_ BluetoothSDPServiceRecord) RemoveServiceRecord() int {
	rv := objc.Send[int](b_.ID, objc.Sel("removeServiceRecord"))
	return rv
}/* debug [instance_methods/method]: RemoveServiceRecord */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for BluetoothSDPServiceRecord */

// Returns an NSDictionary containing the attributes for the service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceRecord/attributes
func (b_ BluetoothSDPServiceRecord) Attributes() objc.IObject /* cross-framework: NSDictionary */ {
	rv := objc.Send[foundation.NSDictionary](b_.ID, objc.Sel("attributes"))
	return rv
}/* debug [instance_properties/getter]: attributes */


// Returns the IOBluetoothDevice that the target service belongs to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceRecord/device
func (b_ BluetoothSDPServiceRecord) Device() IOBluetoothDevice {
	rv := objc.Send[objc.ID](b_.ID, objc.Sel("device"))
	return rv
}/* debug [instance_properties/getter]: device */


// Returns a sorted array of SDP attributes
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothSDPServiceRecord/sortedAttributes-swift.property
func (b_ BluetoothSDPServiceRecord) SortedAttributes() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](b_.ID, objc.Sel("sortedAttributes"))
	return rv
}/* debug [instance_properties/getter]: sortedAttributes */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class IOBluetoothSDPServiceRecord */


