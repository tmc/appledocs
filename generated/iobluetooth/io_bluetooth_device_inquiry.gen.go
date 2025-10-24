// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class IOBluetoothDeviceInquiry */


/* debug [class_header]: Header for IOBluetoothDeviceInquiry */
// The class instance for the [BluetoothDeviceInquiry] class.
var (
	BluetoothDeviceInquiryClass     _BluetoothDeviceInquiryClass
	BluetoothDeviceInquiryClassOnce sync.Once
)

func getBluetoothDeviceInquiryClass() _BluetoothDeviceInquiryClass {
	BluetoothDeviceInquiryClassOnce.Do(func() {
		BluetoothDeviceInquiryClass = _BluetoothDeviceInquiryClass{objc.GetClass("IOBluetoothDeviceInquiry")}
	})
	return BluetoothDeviceInquiryClass
}

type _BluetoothDeviceInquiryClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for BluetoothDeviceInquiry */
// An interface definition for the [BluetoothDeviceInquiry] class.
type IBluetoothDeviceInquiry interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for BluetoothDeviceInquiry */
	// properties:
	Delegate() objc.ID
	SetDelegate(value objc.ID)
	InquiryLength() uint8 /* not a class type */
	SetInquiryLength(value uint8 /* not a class type */)
	SearchType() BluetoothDeviceSearchTypes /* typedef */
	SetSearchType(value BluetoothDeviceSearchTypes /* typedef */)
	UpdateNewDeviceNames() bool
	SetUpdateNewDeviceNames(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for BluetoothDeviceInquiry */
	// methods:
	ClearFoundDevices()
	FoundDevices() foundation.Array
	SetSearchCriteriaMajorDeviceClassMinorDeviceClass(inServiceClassMajor BluetoothServiceClassMajor /* typedef */, inMajorDeviceClass BluetoothDeviceClassMajor /* typedef */, inMinorDeviceClass BluetoothDeviceClassMinor /* typedef */)
	Start() int
	Stop() int
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for BluetoothDeviceInquiry */
// Alloc allocates a new instance without initialization.
func (bc _BluetoothDeviceInquiryClass) Alloc() BluetoothDeviceInquiry {
	rv := objc.Send[BluetoothDeviceInquiry](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (bc _BluetoothDeviceInquiryClass) New() BluetoothDeviceInquiry {
	rv := objc.Send[BluetoothDeviceInquiry](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ BluetoothDeviceInquiry) Init() BluetoothDeviceInquiry {
	rv := objc.Send[BluetoothDeviceInquiry](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ BluetoothDeviceInquiry) Autorelease() BluetoothDeviceInquiry {
	rv := objc.Send[BluetoothDeviceInquiry](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBluetoothDeviceInquiry creates a new BluetoothDeviceInquiry instance.
func NewBluetoothDeviceInquiry() BluetoothDeviceInquiry {
	return getBluetoothDeviceInquiryClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for BluetoothDeviceInquiry */
// Object representing a device inquiry that finds Bluetooth devices in-range of the computer, and (optionally) retrieves name information for them.
//
// You should only use this object if your application needs to know about in-range devices and cannot use the GUI provided by the IOBluetoothUI framework. It will not let you perform unlimited back-to-back inquiries, but will instead throttle the number of attempted inquiries if too many are attempted within a small window of time. Important Note: DO NOT perform remote name requests on devices from delegate methods or while this object is in use. If you wish to do your own remote name requests on devices, do them after you have stopped this object. If you do not heed this warning, you could potentially deadlock your process.


// Object representing a device inquiry that finds Bluetooth devices in-range of the computer, and (optionally) retrieves name information for them.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDeviceInquiry
type BluetoothDeviceInquiry struct {
	objectivec.Object
}

// BluetoothDeviceInquiryFrom constructs a [BluetoothDeviceInquiry] from an unsafe.Pointer.
//
// Object representing a device inquiry that finds Bluetooth devices in-range of the computer, and (optionally) retrieves name information for them.
func BluetoothDeviceInquiryFrom(ptr unsafe.Pointer) BluetoothDeviceInquiry {
	return BluetoothDeviceInquiry{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for BluetoothDeviceInquiry */

// Initializes an alloc’d inquiry object, and sets the delegate object, as if -setDelegate: were called on it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDeviceInquiry/init(delegate:)
func NewBluetoothDeviceInquiryWithDelegate(delegate objc.IObject) BluetoothDeviceInquiry {
	instance := getBluetoothDeviceInquiryClass().Alloc()
	rv := objc.Send[BluetoothDeviceInquiry](instance.ID, objc.Sel("initWithDelegate:"), delegate)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewBluetoothDeviceInquiryWithDelegate */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for BluetoothDeviceInquiry */

// Class method to create an inquiry object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDeviceInquiry/inquiryWithDelegate:
func (bc _BluetoothDeviceInquiryClass) InquiryWithDelegate(delegate objc.IObject) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("inquiryWithDelegate:"), delegate)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=InquiryWithDelegate) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for BluetoothDeviceInquiry */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for BluetoothDeviceInquiry */

// Removes all found devices from the inquiry object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDeviceInquiry/clearFoundDevices()
func (b_ BluetoothDeviceInquiry) ClearFoundDevices() {
	objc.Send[objc.ID](b_.ID, objc.Sel("clearFoundDevices"))
}/* debug [instance_methods/method]: ClearFoundDevices */


// Returns found IOBluetoothDevice objects as an array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDeviceInquiry/foundDevices()
func (b_ BluetoothDeviceInquiry) FoundDevices() foundation.Array {
	rv := objc.Send[foundation.Array](b_.ID, objc.Sel("foundDevices"))
	return rv
}/* debug [instance_methods/method]: FoundDevices */


// Use this method to set the criteria for the device search.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDeviceInquiry/setSearchCriteria(_:majorDeviceClass:minorDeviceClass:)
func (b_ BluetoothDeviceInquiry) SetSearchCriteriaMajorDeviceClassMinorDeviceClass(inServiceClassMajor BluetoothServiceClassMajor /* typedef */, inMajorDeviceClass BluetoothDeviceClassMajor /* typedef */, inMinorDeviceClass BluetoothDeviceClassMinor /* typedef */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setSearchCriteria:majorDeviceClass:minorDeviceClass:"), inServiceClassMajor, inMajorDeviceClass, inMinorDeviceClass)
}/* debug [instance_methods/method]: SetSearchCriteriaMajorDeviceClassMinorDeviceClass */


// Tells inquiry object to begin the inquiry and name updating process, if specified.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDeviceInquiry/start()
func (b_ BluetoothDeviceInquiry) Start() int {
	rv := objc.Send[int](b_.ID, objc.Sel("start"))
	return rv
}/* debug [instance_methods/method]: Start */


// Halts the inquiry object. Could either stop the search for new devices, or the updating of found device names.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDeviceInquiry/stop()
func (b_ BluetoothDeviceInquiry) Stop() int {
	rv := objc.Send[int](b_.ID, objc.Sel("stop"))
	return rv
}/* debug [instance_methods/method]: Stop */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for BluetoothDeviceInquiry */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDeviceInquiry/delegate
func (b_ BluetoothDeviceInquiry) Delegate() objc.ID {
	rv := objc.Send[objc.ID](b_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDeviceInquiry/delegate
func (b_ BluetoothDeviceInquiry) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// Set the length of the inquiry that is performed each time -start is used on an inquiry object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDeviceInquiry/inquiryLength
func (b_ BluetoothDeviceInquiry) InquiryLength() uint8 /* not a class type */ {
	rv := objc.Send[uint8](b_.ID, objc.Sel("inquiryLength"))
	return rv
}/* debug [instance_properties/getter]: inquiryLength */


// Set the length of the inquiry that is performed each time -start is used on an inquiry object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDeviceInquiry/inquiryLength
func (b_ BluetoothDeviceInquiry) SetInquiryLength(value uint8 /* not a class type */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setInquiryLength:"), value)
}/* debug [instance_properties/setter]: inquiryLength */


// Set the devices that are found.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDeviceInquiry/searchType
func (b_ BluetoothDeviceInquiry) SearchType() BluetoothDeviceSearchTypes /* typedef */ {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("searchType"))
	return rv
}/* debug [instance_properties/getter]: searchType */


// Set the devices that are found.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDeviceInquiry/searchType
func (b_ BluetoothDeviceInquiry) SetSearchType(value BluetoothDeviceSearchTypes /* typedef */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setSearchType:"), value)
}/* debug [instance_properties/setter]: searchType */


// Sets whether or not the inquiry object will retrieve the names of devices found during the search.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDeviceInquiry/updateNewDeviceNames
func (b_ BluetoothDeviceInquiry) UpdateNewDeviceNames() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("updateNewDeviceNames"))
	return rv
}/* debug [instance_properties/getter]: updateNewDeviceNames */


// Sets whether or not the inquiry object will retrieve the names of devices found during the search.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDeviceInquiry/updateNewDeviceNames
func (b_ BluetoothDeviceInquiry) SetUpdateNewDeviceNames(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setUpdateNewDeviceNames:"), value)
}/* debug [instance_properties/setter]: updateNewDeviceNames */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class IOBluetoothDeviceInquiry */


