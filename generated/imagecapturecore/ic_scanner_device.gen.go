// Code generated from Apple documentation for ImageCaptureCore. DO NOT EDIT.

package imagecapturecore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/mlcompute"
)

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

// An interface definition for the [ICScannerDevice] class.
type IICScannerDevice interface {
	IICDevice
	RequestOpenSessionWithCredentialsPassword(username string, password string)
	RequestOverviewScan()
	DocumentName() string
	SetDocumentName(value string)
	TransferMode() unsafe.Pointer
	SetTransferMode(value unsafe.Pointer)
	AvailableFunctionalUnitTypes() foundation.Number
	SetAvailableFunctionalUnitTypes(value foundation.INumber)
	DefaultUsername() string
	SetDefaultUsername(value string)
	DocumentUTI() string
	SetDocumentUTI(value string)
	DownloadsDirectory() foundation.URL
	SetDownloadsDirectory(value foundation.IURL)
	MaxMemoryBandSize() unsafe.Pointer
	SetMaxMemoryBandSize(value unsafe.Pointer)
	SelectedFunctionalUnit() unsafe.Pointer
	SetSelectedFunctionalUnit(value unsafe.Pointer)
}

// An object that represents a scanner.
//
// An instance of ICScannerDevice class is intended to be used by the ICScannerDeviceView object. The ICScannerDeviceView class encapsulates the complexities of setting scan parameters, performing scans and saving the result. The developer should consider using ICScannerDeviceView instead of building their own views using the ICScannerDevice object.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerDevice
type ICScannerDevice struct {
	mlcompute.ICDevice
}

// ICScannerDeviceFrom constructs a [ICScannerDevice] from an unsafe.Pointer.
//
// An object that represents a scanner.
func ICScannerDeviceFrom(ptr unsafe.Pointer) ICScannerDevice {
	return ICScannerDevice{
		ICDevice: mlcompute.ICDeviceFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _ICScannerDeviceClass) Alloc() ICScannerDevice {
	rv := objc.Send[ICScannerDevice](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Opens a session on the protected device with the authorized username and passcode.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerDevice/requestOpenSession(withCredentials:password:)
func (i_ ICScannerDevice) RequestOpenSessionWithCredentialsPassword(username string, password string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("requestOpenSessionWithCredentials:password:"), objc.String(username), objc.String(password))
}

// Starts an overview scan on the selected functional unit.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerDevice/requestOverviewScan()
func (i_ ICScannerDevice) RequestOverviewScan() {
	objc.Send[objc.ID](i_.ID, objc.Sel("requestOverviewScan"))
}

// The document’s name.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerDevice/documentName
func (i_ ICScannerDevice) DocumentName() string {
	rv := objc.Send[string](i_.ID, objc.Sel("documentName"))
	return rv
}


// SetDocumentName sets the value of the documentName property.
// The document’s name.

//
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerDevice/documentName
func (i_ ICScannerDevice) SetDocumentName(value string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDocumentName:"), objc.String(value))
}

// The transfer mode for the scanned document.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerDevice/transferMode
func (i_ ICScannerDevice) TransferMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("transferMode"))
	return rv
}


// SetTransferMode sets the value of the transferMode property.
// The transfer mode for the scanned document.

//
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerDevice/transferMode
func (i_ ICScannerDevice) SetTransferMode(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setTransferMode:"), value)
}

// An array of functional unit types available on this scanner.
//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdevice/availablefunctionalunittypes
func (i_ ICScannerDevice) AvailableFunctionalUnitTypes() foundation.Number {
	rv := objc.Send[foundation.Number](i_.ID, objc.Sel("availableFunctionalUnitTypes"))
	return rv
}


// SetAvailableFunctionalUnitTypes sets the value of the availableFunctionalUnitTypes property.
// An array of functional unit types available on this scanner.

//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdevice/availablefunctionalunittypes
func (i_ ICScannerDevice) SetAvailableFunctionalUnitTypes(value foundation.INumber) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setAvailableFunctionalUnitTypes:"), value)
}

// A default username on protected scanners.
//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdevice/defaultusername
func (i_ ICScannerDevice) DefaultUsername() string {
	rv := objc.Send[string](i_.ID, objc.Sel("defaultUsername"))
	return rv
}


// SetDefaultUsername sets the value of the defaultUsername property.
// A default username on protected scanners.

//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdevice/defaultusername
func (i_ ICScannerDevice) SetDefaultUsername(value string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDefaultUsername:"), objc.String(value))
}

// The document’s uniform type identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdevice/documentuti
func (i_ ICScannerDevice) DocumentUTI() string {
	rv := objc.Send[string](i_.ID, objc.Sel("documentUTI"))
	return rv
}


// SetDocumentUTI sets the value of the documentUTI property.
// The document’s uniform type identifier.

//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdevice/documentuti
func (i_ ICScannerDevice) SetDocumentUTI(value string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDocumentUTI:"), objc.String(value))
}

// The downloads directory.
//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdevice/downloadsdirectory
func (i_ ICScannerDevice) DownloadsDirectory() foundation.URL {
	rv := objc.Send[foundation.URL](i_.ID, objc.Sel("downloadsDirectory"))
	return rv
}


// SetDownloadsDirectory sets the value of the downloadsDirectory property.
// The downloads directory.

//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdevice/downloadsdirectory
func (i_ ICScannerDevice) SetDownloadsDirectory(value foundation.IURL) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDownloadsDirectory:"), value)
}

// The total maximum band size requested when performing a memory-based transfer.
//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdevice/maxmemorybandsize
func (i_ ICScannerDevice) MaxMemoryBandSize() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("maxMemoryBandSize"))
	return rv
}


// SetMaxMemoryBandSize sets the value of the maxMemoryBandSize property.
// The total maximum band size requested when performing a memory-based transfer.

//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdevice/maxmemorybandsize
func (i_ ICScannerDevice) SetMaxMemoryBandSize(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMaxMemoryBandSize:"), value)
}

// The currently selected functional unit on the scanner.
//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdevice/selectedfunctionalunit
func (i_ ICScannerDevice) SelectedFunctionalUnit() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("selectedFunctionalUnit"))
	return rv
}


// SetSelectedFunctionalUnit sets the value of the selectedFunctionalUnit property.
// The currently selected functional unit on the scanner.

//
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerdevice/selectedfunctionalunit
func (i_ ICScannerDevice) SetSelectedFunctionalUnit(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSelectedFunctionalUnit:"), value)
}



