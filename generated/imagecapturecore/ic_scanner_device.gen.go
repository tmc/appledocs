// Code generated from Apple documentation for ImageCaptureCore. DO NOT EDIT.

package imagecapturecore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class ICScannerDevice */


/* debug [class_header]: Header for ICScannerDevice */
// The class instance for the [ICScannerDevice] class.
var (
	ICScannerDeviceClass     _ICScannerDeviceClass
	ICScannerDeviceClassOnce sync.Once
)

func getICScannerDeviceClass() _ICScannerDeviceClass {
	ICScannerDeviceClassOnce.Do(func() {
		ICScannerDeviceClass = _ICScannerDeviceClass{objc.GetClass("ICScannerDevice")}
	})
	return ICScannerDeviceClass
}

type _ICScannerDeviceClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ICScannerDevice */
// An interface definition for the [ICScannerDevice] class.
type IICScannerDevice interface {
	IICDevice
	
/* debug [class_interface_properties]: Properties for ICScannerDevice */
	// properties:
	AvailableFunctionalUnitTypes() objc.IObject /* cross-framework: NSNumber */
	SetAvailableFunctionalUnitTypes(value objc.IObject /* cross-framework: NSNumber */)
	TransferMode() unsafe.Pointer
	SetTransferMode(value unsafe.Pointer)
	DocumentName() unsafe.Pointer
	SetDocumentName(value unsafe.Pointer)
	DocumentUTI() unsafe.Pointer
	SetDocumentUTI(value unsafe.Pointer)
	SelectedFunctionalUnit() ICScannerFunctionalUnit
	SetSelectedFunctionalUnit(value ICScannerFunctionalUnit)
	DownloadsDirectory() unsafe.Pointer
	SetDownloadsDirectory(value unsafe.Pointer)
	MaxMemoryBandSize() unsafe.Pointer
	SetMaxMemoryBandSize(value unsafe.Pointer)
	DefaultUsername() unsafe.Pointer
	SetDefaultUsername(value unsafe.Pointer)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ICScannerDevice */
	// methods:
	RequestOverviewScan()
	RequestSelect()
	RequestSelectFunctionalUnit(type_ ICScannerFunctionalUnitType)
	CancelScan()
	RequestScan()
	RequestOpenSession()
	RequestOpenSessionWithCredentialsPassword(username string, password string)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ICScannerDevice */
// Alloc allocates a new instance without initialization.
func (ic _ICScannerDeviceClass) Alloc() ICScannerDevice {
	rv := objc.Send[ICScannerDevice](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ICScannerDeviceClass) New() ICScannerDevice {
	rv := objc.Send[ICScannerDevice](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ICScannerDevice) Init() ICScannerDevice {
	rv := objc.Send[ICScannerDevice](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ICScannerDevice) Autorelease() ICScannerDevice {
	rv := objc.Send[ICScannerDevice](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewICScannerDevice creates a new ICScannerDevice instance.
func NewICScannerDevice() ICScannerDevice {
	return getICScannerDeviceClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ICScannerDevice */
// An object that represents a scanner.
//
// An instance of ICScannerDevice class is intended to be used by the ICScannerDeviceView object. The ICScannerDeviceView class encapsulates the complexities of setting scan parameters, performing scans and saving the result. The developer should consider using ICScannerDeviceView instead of building their own views using the ICScannerDevice object.


// An object that represents a scanner.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerDevice
type ICScannerDevice struct {
	ICDevice
}

// ICScannerDeviceFrom constructs a [ICScannerDevice] from an unsafe.Pointer.
//
// An object that represents a scanner.
func ICScannerDeviceFrom(ptr unsafe.Pointer) ICScannerDevice {
	return ICScannerDevice{
		ICDevice: ICDeviceFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ICScannerDevice *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ICScannerDevice */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ICScannerDevice */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ICScannerDevice */

// Starts an overview scan on the selected functional unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdevice/1507646-requestoverviewscan
func (i_ ICScannerDevice) RequestOverviewScan() {
	objc.Send[objc.ID](i_.ID, objc.Sel("requestOverviewScan"))
}/* debug [instance_methods/method]: RequestOverviewScan */


// Requests to select a functional unit on the scanner.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdevice/1507654-requestselect
func (i_ ICScannerDevice) RequestSelect() {
	objc.Send[objc.ID](i_.ID, objc.Sel("requestSelect"))
}/* debug [instance_methods/method]: RequestSelect */


// Requests to select a functional unit on the scanner.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdevice/1507654-requestselectfunctionalunit
func (i_ ICScannerDevice) RequestSelectFunctionalUnit(type_ ICScannerFunctionalUnitType) {
	objc.Send[objc.ID](i_.ID, objc.Sel("requestSelectFunctionalUnit:"), type_)
}/* debug [instance_methods/method]: RequestSelectFunctionalUnit */


// Cancels the current scan.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdevice/1507771-cancelscan
func (i_ ICScannerDevice) CancelScan() {
	objc.Send[objc.ID](i_.ID, objc.Sel("cancelScan"))
}/* debug [instance_methods/method]: CancelScan */


// Starts a scan on the selected functional unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdevice/1508117-requestscan
func (i_ ICScannerDevice) RequestScan() {
	objc.Send[objc.ID](i_.ID, objc.Sel("requestScan"))
}/* debug [instance_methods/method]: RequestScan */


// Opens a session on the protected device with the authorized username and passcode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdevice/2881931-requestopensession
func (i_ ICScannerDevice) RequestOpenSession() {
	objc.Send[objc.ID](i_.ID, objc.Sel("requestOpenSession"))
}/* debug [instance_methods/method]: RequestOpenSession */


// Opens a session on the protected device with the authorized username and passcode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdevice/2881931-requestopensessionwithcredential
func (i_ ICScannerDevice) RequestOpenSessionWithCredentialsPassword(username string, password string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("requestOpenSessionWithCredentials:password:"), objc.String(username), objc.String(password))
}/* debug [instance_methods/method]: RequestOpenSessionWithCredentialsPassword */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ICScannerDevice */

// An array of functional unit types available on this scanner.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdevice/1507679-availablefunctionalunittypes
func (i_ ICScannerDevice) AvailableFunctionalUnitTypes() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](i_.ID, objc.Sel("availableFunctionalUnitTypes"))
	return rv
}/* debug [instance_properties/getter]: availableFunctionalUnitTypes */


// An array of functional unit types available on this scanner.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdevice/1507679-availablefunctionalunittypes
func (i_ ICScannerDevice) SetAvailableFunctionalUnitTypes(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setAvailableFunctionalUnitTypes:"), value)
}/* debug [instance_properties/setter]: availableFunctionalUnitTypes */


// The transfer mode for the scanned document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdevice/1507698-transfermode
func (i_ ICScannerDevice) TransferMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("transferMode"))
	return rv
}/* debug [instance_properties/getter]: transferMode */


// The transfer mode for the scanned document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdevice/1507698-transfermode
func (i_ ICScannerDevice) SetTransferMode(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setTransferMode:"), value)
}/* debug [instance_properties/setter]: transferMode */


// The document’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdevice/1507879-documentname
func (i_ ICScannerDevice) DocumentName() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("documentName"))
	return rv
}/* debug [instance_properties/getter]: documentName */


// The document’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdevice/1507879-documentname
func (i_ ICScannerDevice) SetDocumentName(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDocumentName:"), value)
}/* debug [instance_properties/setter]: documentName */


// The document’s uniform type identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdevice/1507955-documentuti
func (i_ ICScannerDevice) DocumentUTI() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("documentUTI"))
	return rv
}/* debug [instance_properties/getter]: documentUTI */


// The document’s uniform type identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdevice/1507955-documentuti
func (i_ ICScannerDevice) SetDocumentUTI(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDocumentUTI:"), value)
}/* debug [instance_properties/setter]: documentUTI */


// The currently selected functional unit on the scanner.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdevice/1508000-selectedfunctionalunit
func (i_ ICScannerDevice) SelectedFunctionalUnit() ICScannerFunctionalUnit {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("selectedFunctionalUnit"))
	return rv
}/* debug [instance_properties/getter]: selectedFunctionalUnit */


// The currently selected functional unit on the scanner.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdevice/1508000-selectedfunctionalunit
func (i_ ICScannerDevice) SetSelectedFunctionalUnit(value ICScannerFunctionalUnit) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSelectedFunctionalUnit:"), value)
}/* debug [instance_properties/setter]: selectedFunctionalUnit */


// The downloads directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdevice/1508025-downloadsdirectory
func (i_ ICScannerDevice) DownloadsDirectory() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("downloadsDirectory"))
	return rv
}/* debug [instance_properties/getter]: downloadsDirectory */


// The downloads directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdevice/1508025-downloadsdirectory
func (i_ ICScannerDevice) SetDownloadsDirectory(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDownloadsDirectory:"), value)
}/* debug [instance_properties/setter]: downloadsDirectory */


// The total maximum band size requested when performing a memory-based transfer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdevice/1508156-maxmemorybandsize
func (i_ ICScannerDevice) MaxMemoryBandSize() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("maxMemoryBandSize"))
	return rv
}/* debug [instance_properties/getter]: maxMemoryBandSize */


// The total maximum band size requested when performing a memory-based transfer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdevice/1508156-maxmemorybandsize
func (i_ ICScannerDevice) SetMaxMemoryBandSize(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMaxMemoryBandSize:"), value)
}/* debug [instance_properties/setter]: maxMemoryBandSize */


// A default username on protected scanners.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdevice/2881932-defaultusername
func (i_ ICScannerDevice) DefaultUsername() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("defaultUsername"))
	return rv
}/* debug [instance_properties/getter]: defaultUsername */


// A default username on protected scanners.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdevice/2881932-defaultusername
func (i_ ICScannerDevice) SetDefaultUsername(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDefaultUsername:"), value)
}/* debug [instance_properties/setter]: defaultUsername */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ICScannerDevice */



