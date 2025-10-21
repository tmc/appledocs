// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [BluetoothDeviceInquiry] class.
type IBluetoothDeviceInquiry interface {
	objectivec.IObject
	ClearFoundDevices()
	FoundDevices() unsafe.Pointer
	SetSearchCriteriaMajorDeviceClassMinorDeviceClass(inServiceClassMajor unsafe.Pointer, inMajorDeviceClass unsafe.Pointer, inMinorDeviceClass unsafe.Pointer)
	Start() unsafe.Pointer
	Stop() unsafe.Pointer
}

// Object representing a device inquiry that finds Bluetooth devices in-range of the computer, and (optionally) retrieves name information for them.
//
// You should only use this object if your application needs to know about in-range devices and cannot use the GUI provided by the IOBluetoothUI framework. It will not let you perform unlimited back-to-back inquiries, but will instead throttle the number of attempted inquiries if too many are attempted within a small window of time. Important Note: DO NOT perform remote name requests on devices from delegate methods or while this object is in use. If you wish to do your own remote name requests on devices, do them after you have stopped this object. If you do not heed this warning, you could potentially deadlock your process.
//
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

// Alloc allocates a new instance without initialization.
func (bc _BluetoothDeviceInquiryClass) Alloc() BluetoothDeviceInquiry {
	rv := objc.Send[BluetoothDeviceInquiry](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Initializes an alloc’d inquiry object, and sets the delegate object, as if -setDelegate: were called on it.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDeviceInquiry/init(delegate:)
func NewBluetoothDeviceInquiryWithDelegate(delegate objc.ID) BluetoothDeviceInquiry {
	instance := getBluetoothDeviceInquiryClass().Alloc()
	rv := objc.Send[BluetoothDeviceInquiry](instance.ID, objc.Sel("initWithDelegate:"), delegate)
	rv.Autorelease()
	return rv
}


// Class method to create an inquiry object.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDeviceInquiry/inquiryWithDelegate:
func (bc _BluetoothDeviceInquiryClass) InquiryWithDelegate(delegate objc.ID) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(bc.class), objc.Sel("inquiryWithDelegate:"), delegate)
	return rv
}

// Removes all found devices from the inquiry object.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDeviceInquiry/clearFoundDevices()
func (b_ BluetoothDeviceInquiry) ClearFoundDevices() {
	objc.Send[objc.ID](b_.ID, objc.Sel("clearFoundDevices"))
}

// Returns found IOBluetoothDevice objects as an array.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDeviceInquiry/foundDevices()
func (b_ BluetoothDeviceInquiry) FoundDevices() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("foundDevices"))
	return rv
}

// Use this method to set the criteria for the device search.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDeviceInquiry/setSearchCriteria(_:majorDeviceClass:minorDeviceClass:)
func (b_ BluetoothDeviceInquiry) SetSearchCriteriaMajorDeviceClassMinorDeviceClass(inServiceClassMajor unsafe.Pointer, inMajorDeviceClass unsafe.Pointer, inMinorDeviceClass unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setSearchCriteria:majorDeviceClass:minorDeviceClass:"), inServiceClassMajor, inMajorDeviceClass, inMinorDeviceClass)
}

// Tells inquiry object to begin the inquiry and name updating process, if specified.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDeviceInquiry/start()
func (b_ BluetoothDeviceInquiry) Start() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("start"))
	return rv
}

// Halts the inquiry object. Could either stop the search for new devices, or the updating of found device names.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDeviceInquiry/stop()
func (b_ BluetoothDeviceInquiry) Stop() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("stop"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDeviceInquiry/delegate
func (b_ BluetoothDeviceInquiry) Delegate() objc.ID {
	rv := objc.Send[objc.ID](b_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDeviceInquiry/delegate
func (b_ BluetoothDeviceInquiry) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setDelegate:"), value)
}

// Set the length of the inquiry that is performed each time -start is used on an inquiry object.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDeviceInquiry/inquiryLength
func (b_ BluetoothDeviceInquiry) InquiryLength() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("inquiryLength"))
	return rv
}


// SetInquiryLength sets the value of the inquiryLength property.
// Set the length of the inquiry that is performed each time -start is used on an inquiry object.

//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDeviceInquiry/inquiryLength
func (b_ BluetoothDeviceInquiry) SetInquiryLength(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setInquiryLength:"), value)
}

// Set the devices that are found.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDeviceInquiry/searchType
func (b_ BluetoothDeviceInquiry) SearchType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("searchType"))
	return rv
}


// SetSearchType sets the value of the searchType property.
// Set the devices that are found.

//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDeviceInquiry/searchType
func (b_ BluetoothDeviceInquiry) SetSearchType(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setSearchType:"), value)
}

// Sets whether or not the inquiry object will retrieve the names of devices found during the search.
//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDeviceInquiry/updateNewDeviceNames
func (b_ BluetoothDeviceInquiry) UpdateNewDeviceNames() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("updateNewDeviceNames"))
	return rv
}


// SetUpdateNewDeviceNames sets the value of the updateNewDeviceNames property.
// Sets whether or not the inquiry object will retrieve the names of devices found during the search.

//
// [Full Topic]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDeviceInquiry/updateNewDeviceNames
func (b_ BluetoothDeviceInquiry) SetUpdateNewDeviceNames(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setUpdateNewDeviceNames:"), value)
}


