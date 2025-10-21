// Code generated from Apple documentation for ImageCaptureCore. DO NOT EDIT.

package imagecapturecore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
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
	objectivec.IObject
	RequestOpenSessionWithCredentialsPassword(username string, password string)
	RequestOverviewScan()
}

// An object that represents a scanner.
//
// An instance of ICScannerDevice class is intended to be used by the ICScannerDeviceView object. The ICScannerDeviceView class encapsulates the complexities of setting scan parameters, performing scans and saving the result. The developer should consider using ICScannerDeviceView instead of building their own views using the ICScannerDevice object.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerDevice
type ICScannerDevice struct {
	objectivec.Object
}

// ICScannerDeviceFrom constructs a [ICScannerDevice] from an unsafe.Pointer.
//
// An object that represents a scanner.
func ICScannerDeviceFrom(ptr unsafe.Pointer) ICScannerDevice {
	return ICScannerDevice{objectivec.Object{objc.ID(ptr)}}
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



