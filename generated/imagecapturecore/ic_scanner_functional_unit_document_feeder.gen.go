// Code generated from Apple documentation for ImageCaptureCore. DO NOT EDIT.

package imagecapturecore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class ICScannerFunctionalUnitDocumentFeeder */


/* debug [class_header]: Header for ICScannerFunctionalUnitDocumentFeeder */
// The class instance for the [ICScannerFunctionalUnitDocumentFeeder] class.
var (
	ICScannerFunctionalUnitDocumentFeederClass     _ICScannerFunctionalUnitDocumentFeederClass
	ICScannerFunctionalUnitDocumentFeederClassOnce sync.Once
)

func getICScannerFunctionalUnitDocumentFeederClass() _ICScannerFunctionalUnitDocumentFeederClass {
	ICScannerFunctionalUnitDocumentFeederClassOnce.Do(func() {
		ICScannerFunctionalUnitDocumentFeederClass = _ICScannerFunctionalUnitDocumentFeederClass{objc.GetClass("ICScannerFunctionalUnitDocumentFeeder")}
	})
	return ICScannerFunctionalUnitDocumentFeederClass
}

type _ICScannerFunctionalUnitDocumentFeederClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ICScannerFunctionalUnitDocumentFeeder */
// An interface definition for the [ICScannerFunctionalUnitDocumentFeeder] class.
type IICScannerFunctionalUnitDocumentFeeder interface {
	IICScannerFunctionalUnit
	
/* debug [class_interface_properties]: Properties for ICScannerFunctionalUnitDocumentFeeder */
	// properties:
	SupportedDocumentTypes() unsafe.Pointer
	SetSupportedDocumentTypes(value unsafe.Pointer)
	SupportsDuplexScanning() unsafe.Pointer
	SetSupportsDuplexScanning(value unsafe.Pointer)
	DuplexScanningEnabled() unsafe.Pointer
	SetDuplexScanningEnabled(value unsafe.Pointer)
	OddPageOrientation() unsafe.Pointer
	SetOddPageOrientation(value unsafe.Pointer)
	DocumentType() unsafe.Pointer
	SetDocumentType(value unsafe.Pointer)
	DocumentLoaded() unsafe.Pointer
	SetDocumentLoaded(value unsafe.Pointer)
	ReverseFeederPageOrder() unsafe.Pointer
	SetReverseFeederPageOrder(value unsafe.Pointer)
	DocumentSize() Size get /* not a class type */
	SetDocumentSize(value Size get /* not a class type */)
	EvenPageOrientation() unsafe.Pointer
	SetEvenPageOrientation(value unsafe.Pointer)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ICScannerFunctionalUnitDocumentFeeder */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ICScannerFunctionalUnitDocumentFeeder */
// Alloc allocates a new instance without initialization.
func (ic _ICScannerFunctionalUnitDocumentFeederClass) Alloc() ICScannerFunctionalUnitDocumentFeeder {
	rv := objc.Send[ICScannerFunctionalUnitDocumentFeeder](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ICScannerFunctionalUnitDocumentFeederClass) New() ICScannerFunctionalUnitDocumentFeeder {
	rv := objc.Send[ICScannerFunctionalUnitDocumentFeeder](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ICScannerFunctionalUnitDocumentFeeder) Init() ICScannerFunctionalUnitDocumentFeeder {
	rv := objc.Send[ICScannerFunctionalUnitDocumentFeeder](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ICScannerFunctionalUnitDocumentFeeder) Autorelease() ICScannerFunctionalUnitDocumentFeeder {
	rv := objc.Send[ICScannerFunctionalUnitDocumentFeeder](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewICScannerFunctionalUnitDocumentFeeder creates a new ICScannerFunctionalUnitDocumentFeeder instance.
func NewICScannerFunctionalUnitDocumentFeeder() ICScannerFunctionalUnitDocumentFeeder {
	return getICScannerFunctionalUnitDocumentFeederClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ICScannerFunctionalUnitDocumentFeeder */
// An object that represents the document feeder unit on a scanner.


// An object that represents the document feeder unit on a scanner.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFunctionalUnitDocumentFeeder
type ICScannerFunctionalUnitDocumentFeeder struct {
	ICScannerFunctionalUnit
}

// ICScannerFunctionalUnitDocumentFeederFrom constructs a [ICScannerFunctionalUnitDocumentFeeder] from an unsafe.Pointer.
//
// An object that represents the document feeder unit on a scanner.
func ICScannerFunctionalUnitDocumentFeederFrom(ptr unsafe.Pointer) ICScannerFunctionalUnitDocumentFeeder {
	return ICScannerFunctionalUnitDocumentFeeder{
		ICScannerFunctionalUnit: ICScannerFunctionalUnitFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ICScannerFunctionalUnitDocumentFeeder *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ICScannerFunctionalUnitDocumentFeeder */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ICScannerFunctionalUnitDocumentFeeder */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ICScannerFunctionalUnitDocumentFeeder */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ICScannerFunctionalUnitDocumentFeeder */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunitdocumentfeeder/1507584-supporteddocumenttypes
func (i_ ICScannerFunctionalUnitDocumentFeeder) SupportedDocumentTypes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("supportedDocumentTypes"))
	return rv
}/* debug [instance_properties/getter]: supportedDocumentTypes */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunitdocumentfeeder/1507584-supporteddocumenttypes
func (i_ ICScannerFunctionalUnitDocumentFeeder) SetSupportedDocumentTypes(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSupportedDocumentTypes:"), value)
}/* debug [instance_properties/setter]: supportedDocumentTypes */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunitdocumentfeeder/1507588-supportsduplexscanning
func (i_ ICScannerFunctionalUnitDocumentFeeder) SupportsDuplexScanning() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("supportsDuplexScanning"))
	return rv
}/* debug [instance_properties/getter]: supportsDuplexScanning */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunitdocumentfeeder/1507588-supportsduplexscanning
func (i_ ICScannerFunctionalUnitDocumentFeeder) SetSupportsDuplexScanning(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSupportsDuplexScanning:"), value)
}/* debug [instance_properties/setter]: supportsDuplexScanning */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunitdocumentfeeder/1507592-duplexscanningenabled
func (i_ ICScannerFunctionalUnitDocumentFeeder) DuplexScanningEnabled() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("duplexScanningEnabled"))
	return rv
}/* debug [instance_properties/getter]: duplexScanningEnabled */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunitdocumentfeeder/1507592-duplexscanningenabled
func (i_ ICScannerFunctionalUnitDocumentFeeder) SetDuplexScanningEnabled(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDuplexScanningEnabled:"), value)
}/* debug [instance_properties/setter]: duplexScanningEnabled */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunitdocumentfeeder/1507739-oddpageorientation
func (i_ ICScannerFunctionalUnitDocumentFeeder) OddPageOrientation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("oddPageOrientation"))
	return rv
}/* debug [instance_properties/getter]: oddPageOrientation */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunitdocumentfeeder/1507739-oddpageorientation
func (i_ ICScannerFunctionalUnitDocumentFeeder) SetOddPageOrientation(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setOddPageOrientation:"), value)
}/* debug [instance_properties/setter]: oddPageOrientation */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunitdocumentfeeder/1507783-documenttype
func (i_ ICScannerFunctionalUnitDocumentFeeder) DocumentType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("documentType"))
	return rv
}/* debug [instance_properties/getter]: documentType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunitdocumentfeeder/1507783-documenttype
func (i_ ICScannerFunctionalUnitDocumentFeeder) SetDocumentType(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDocumentType:"), value)
}/* debug [instance_properties/setter]: documentType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunitdocumentfeeder/1507825-documentloaded
func (i_ ICScannerFunctionalUnitDocumentFeeder) DocumentLoaded() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("documentLoaded"))
	return rv
}/* debug [instance_properties/getter]: documentLoaded */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunitdocumentfeeder/1507825-documentloaded
func (i_ ICScannerFunctionalUnitDocumentFeeder) SetDocumentLoaded(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDocumentLoaded:"), value)
}/* debug [instance_properties/setter]: documentLoaded */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunitdocumentfeeder/1507908-reversefeederpageorder
func (i_ ICScannerFunctionalUnitDocumentFeeder) ReverseFeederPageOrder() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("reverseFeederPageOrder"))
	return rv
}/* debug [instance_properties/getter]: reverseFeederPageOrder */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunitdocumentfeeder/1507908-reversefeederpageorder
func (i_ ICScannerFunctionalUnitDocumentFeeder) SetReverseFeederPageOrder(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setReverseFeederPageOrder:"), value)
}/* debug [instance_properties/setter]: reverseFeederPageOrder */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunitdocumentfeeder/1508074-documentsize
func (i_ ICScannerFunctionalUnitDocumentFeeder) DocumentSize() Size get /* not a class type */ {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("documentSize"))
	return rv
}/* debug [instance_properties/getter]: documentSize */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunitdocumentfeeder/1508074-documentsize
func (i_ ICScannerFunctionalUnitDocumentFeeder) SetDocumentSize(value Size get /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDocumentSize:"), value)
}/* debug [instance_properties/setter]: documentSize */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunitdocumentfeeder/1508090-evenpageorientation
func (i_ ICScannerFunctionalUnitDocumentFeeder) EvenPageOrientation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("evenPageOrientation"))
	return rv
}/* debug [instance_properties/getter]: evenPageOrientation */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunitdocumentfeeder/1508090-evenpageorientation
func (i_ ICScannerFunctionalUnitDocumentFeeder) SetEvenPageOrientation(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setEvenPageOrientation:"), value)
}/* debug [instance_properties/setter]: evenPageOrientation */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ICScannerFunctionalUnitDocumentFeeder */



