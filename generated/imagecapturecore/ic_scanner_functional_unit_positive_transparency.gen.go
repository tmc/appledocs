// Code generated from Apple documentation for ImageCaptureCore. DO NOT EDIT.

package imagecapturecore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class ICScannerFunctionalUnitPositiveTransparency */


/* debug [class_header]: Header for ICScannerFunctionalUnitPositiveTransparency */
// The class instance for the [ICScannerFunctionalUnitPositiveTransparency] class.
var (
	ICScannerFunctionalUnitPositiveTransparencyClass     _ICScannerFunctionalUnitPositiveTransparencyClass
	ICScannerFunctionalUnitPositiveTransparencyClassOnce sync.Once
)

func getICScannerFunctionalUnitPositiveTransparencyClass() _ICScannerFunctionalUnitPositiveTransparencyClass {
	ICScannerFunctionalUnitPositiveTransparencyClassOnce.Do(func() {
		ICScannerFunctionalUnitPositiveTransparencyClass = _ICScannerFunctionalUnitPositiveTransparencyClass{objc.GetClass("ICScannerFunctionalUnitPositiveTransparency")}
	})
	return ICScannerFunctionalUnitPositiveTransparencyClass
}

type _ICScannerFunctionalUnitPositiveTransparencyClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ICScannerFunctionalUnitPositiveTransparency */
// An interface definition for the [ICScannerFunctionalUnitPositiveTransparency] class.
type IICScannerFunctionalUnitPositiveTransparency interface {
	IICScannerFunctionalUnit
	
/* debug [class_interface_properties]: Properties for ICScannerFunctionalUnitPositiveTransparency */
	// properties:
	DocumentType() unsafe.Pointer
	SetDocumentType(value unsafe.Pointer)
	DocumentSize() Size get /* not a class type */
	SetDocumentSize(value Size get /* not a class type */)
	SupportedDocumentTypes() unsafe.Pointer
	SetSupportedDocumentTypes(value unsafe.Pointer)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ICScannerFunctionalUnitPositiveTransparency */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ICScannerFunctionalUnitPositiveTransparency */
// Alloc allocates a new instance without initialization.
func (ic _ICScannerFunctionalUnitPositiveTransparencyClass) Alloc() ICScannerFunctionalUnitPositiveTransparency {
	rv := objc.Send[ICScannerFunctionalUnitPositiveTransparency](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ICScannerFunctionalUnitPositiveTransparencyClass) New() ICScannerFunctionalUnitPositiveTransparency {
	rv := objc.Send[ICScannerFunctionalUnitPositiveTransparency](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ICScannerFunctionalUnitPositiveTransparency) Init() ICScannerFunctionalUnitPositiveTransparency {
	rv := objc.Send[ICScannerFunctionalUnitPositiveTransparency](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ICScannerFunctionalUnitPositiveTransparency) Autorelease() ICScannerFunctionalUnitPositiveTransparency {
	rv := objc.Send[ICScannerFunctionalUnitPositiveTransparency](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewICScannerFunctionalUnitPositiveTransparency creates a new ICScannerFunctionalUnitPositiveTransparency instance.
func NewICScannerFunctionalUnitPositiveTransparency() ICScannerFunctionalUnitPositiveTransparency {
	return getICScannerFunctionalUnitPositiveTransparencyClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ICScannerFunctionalUnitPositiveTransparency */
// An object that represents the transparency unit for scanning positives on the scanner.


// An object that represents the transparency unit for scanning positives on the scanner.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFunctionalUnitPositiveTransparency
type ICScannerFunctionalUnitPositiveTransparency struct {
	ICScannerFunctionalUnit
}

// ICScannerFunctionalUnitPositiveTransparencyFrom constructs a [ICScannerFunctionalUnitPositiveTransparency] from an unsafe.Pointer.
//
// An object that represents the transparency unit for scanning positives on the scanner.
func ICScannerFunctionalUnitPositiveTransparencyFrom(ptr unsafe.Pointer) ICScannerFunctionalUnitPositiveTransparency {
	return ICScannerFunctionalUnitPositiveTransparency{
		ICScannerFunctionalUnit: ICScannerFunctionalUnitFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ICScannerFunctionalUnitPositiveTransparency *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ICScannerFunctionalUnitPositiveTransparency */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ICScannerFunctionalUnitPositiveTransparency */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ICScannerFunctionalUnitPositiveTransparency */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ICScannerFunctionalUnitPositiveTransparency */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunitpositivetransparency/1507591-documenttype
func (i_ ICScannerFunctionalUnitPositiveTransparency) DocumentType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("documentType"))
	return rv
}/* debug [instance_properties/getter]: documentType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunitpositivetransparency/1507591-documenttype
func (i_ ICScannerFunctionalUnitPositiveTransparency) SetDocumentType(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDocumentType:"), value)
}/* debug [instance_properties/setter]: documentType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunitpositivetransparency/1507652-documentsize
func (i_ ICScannerFunctionalUnitPositiveTransparency) DocumentSize() Size get /* not a class type */ {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("documentSize"))
	return rv
}/* debug [instance_properties/getter]: documentSize */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunitpositivetransparency/1507652-documentsize
func (i_ ICScannerFunctionalUnitPositiveTransparency) SetDocumentSize(value Size get /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDocumentSize:"), value)
}/* debug [instance_properties/setter]: documentSize */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunitpositivetransparency/1508152-supporteddocumenttypes
func (i_ ICScannerFunctionalUnitPositiveTransparency) SupportedDocumentTypes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("supportedDocumentTypes"))
	return rv
}/* debug [instance_properties/getter]: supportedDocumentTypes */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunitpositivetransparency/1508152-supporteddocumenttypes
func (i_ ICScannerFunctionalUnitPositiveTransparency) SetSupportedDocumentTypes(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSupportedDocumentTypes:"), value)
}/* debug [instance_properties/setter]: supportedDocumentTypes */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ICScannerFunctionalUnitPositiveTransparency */





