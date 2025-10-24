// Code generated from Apple documentation for ImageCaptureCore. DO NOT EDIT.

package imagecapturecore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class ICScannerFunctionalUnitNegativeTransparency */


/* debug [class_header]: Header for ICScannerFunctionalUnitNegativeTransparency */
// The class instance for the [ICScannerFunctionalUnitNegativeTransparency] class.
var (
	ICScannerFunctionalUnitNegativeTransparencyClass     _ICScannerFunctionalUnitNegativeTransparencyClass
	ICScannerFunctionalUnitNegativeTransparencyClassOnce sync.Once
)

func getICScannerFunctionalUnitNegativeTransparencyClass() _ICScannerFunctionalUnitNegativeTransparencyClass {
	ICScannerFunctionalUnitNegativeTransparencyClassOnce.Do(func() {
		ICScannerFunctionalUnitNegativeTransparencyClass = _ICScannerFunctionalUnitNegativeTransparencyClass{objc.GetClass("ICScannerFunctionalUnitNegativeTransparency")}
	})
	return ICScannerFunctionalUnitNegativeTransparencyClass
}

type _ICScannerFunctionalUnitNegativeTransparencyClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ICScannerFunctionalUnitNegativeTransparency */
// An interface definition for the [ICScannerFunctionalUnitNegativeTransparency] class.
type IICScannerFunctionalUnitNegativeTransparency interface {
	IICScannerFunctionalUnit
	
/* debug [class_interface_properties]: Properties for ICScannerFunctionalUnitNegativeTransparency */
	// properties:
	DocumentSize() Size get /* not a class type */
	SetDocumentSize(value Size get /* not a class type */)
	SupportedDocumentTypes() unsafe.Pointer
	SetSupportedDocumentTypes(value unsafe.Pointer)
	DocumentType() unsafe.Pointer
	SetDocumentType(value unsafe.Pointer)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ICScannerFunctionalUnitNegativeTransparency */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ICScannerFunctionalUnitNegativeTransparency */
// Alloc allocates a new instance without initialization.
func (ic _ICScannerFunctionalUnitNegativeTransparencyClass) Alloc() ICScannerFunctionalUnitNegativeTransparency {
	rv := objc.Send[ICScannerFunctionalUnitNegativeTransparency](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ICScannerFunctionalUnitNegativeTransparencyClass) New() ICScannerFunctionalUnitNegativeTransparency {
	rv := objc.Send[ICScannerFunctionalUnitNegativeTransparency](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ICScannerFunctionalUnitNegativeTransparency) Init() ICScannerFunctionalUnitNegativeTransparency {
	rv := objc.Send[ICScannerFunctionalUnitNegativeTransparency](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ICScannerFunctionalUnitNegativeTransparency) Autorelease() ICScannerFunctionalUnitNegativeTransparency {
	rv := objc.Send[ICScannerFunctionalUnitNegativeTransparency](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewICScannerFunctionalUnitNegativeTransparency creates a new ICScannerFunctionalUnitNegativeTransparency instance.
func NewICScannerFunctionalUnitNegativeTransparency() ICScannerFunctionalUnitNegativeTransparency {
	return getICScannerFunctionalUnitNegativeTransparencyClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ICScannerFunctionalUnitNegativeTransparency */
// An object that represents the transparency unit for scanning negatives on the scanner.


// An object that represents the transparency unit for scanning negatives on the scanner.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFunctionalUnitNegativeTransparency
type ICScannerFunctionalUnitNegativeTransparency struct {
	ICScannerFunctionalUnit
}

// ICScannerFunctionalUnitNegativeTransparencyFrom constructs a [ICScannerFunctionalUnitNegativeTransparency] from an unsafe.Pointer.
//
// An object that represents the transparency unit for scanning negatives on the scanner.
func ICScannerFunctionalUnitNegativeTransparencyFrom(ptr unsafe.Pointer) ICScannerFunctionalUnitNegativeTransparency {
	return ICScannerFunctionalUnitNegativeTransparency{
		ICScannerFunctionalUnit: ICScannerFunctionalUnitFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ICScannerFunctionalUnitNegativeTransparency *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ICScannerFunctionalUnitNegativeTransparency */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ICScannerFunctionalUnitNegativeTransparency */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ICScannerFunctionalUnitNegativeTransparency */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ICScannerFunctionalUnitNegativeTransparency */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunitnegativetransparency/1507816-documentsize
func (i_ ICScannerFunctionalUnitNegativeTransparency) DocumentSize() Size get /* not a class type */ {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("documentSize"))
	return rv
}/* debug [instance_properties/getter]: documentSize */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunitnegativetransparency/1507816-documentsize
func (i_ ICScannerFunctionalUnitNegativeTransparency) SetDocumentSize(value Size get /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDocumentSize:"), value)
}/* debug [instance_properties/setter]: documentSize */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunitnegativetransparency/1507868-supporteddocumenttypes
func (i_ ICScannerFunctionalUnitNegativeTransparency) SupportedDocumentTypes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("supportedDocumentTypes"))
	return rv
}/* debug [instance_properties/getter]: supportedDocumentTypes */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunitnegativetransparency/1507868-supporteddocumenttypes
func (i_ ICScannerFunctionalUnitNegativeTransparency) SetSupportedDocumentTypes(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSupportedDocumentTypes:"), value)
}/* debug [instance_properties/setter]: supportedDocumentTypes */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunitnegativetransparency/1508012-documenttype
func (i_ ICScannerFunctionalUnitNegativeTransparency) DocumentType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("documentType"))
	return rv
}/* debug [instance_properties/getter]: documentType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunitnegativetransparency/1508012-documenttype
func (i_ ICScannerFunctionalUnitNegativeTransparency) SetDocumentType(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDocumentType:"), value)
}/* debug [instance_properties/setter]: documentType */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ICScannerFunctionalUnitNegativeTransparency */



