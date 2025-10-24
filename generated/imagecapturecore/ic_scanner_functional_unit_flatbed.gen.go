// Code generated from Apple documentation for ImageCaptureCore. DO NOT EDIT.

package imagecapturecore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class ICScannerFunctionalUnitFlatbed */


/* debug [class_header]: Header for ICScannerFunctionalUnitFlatbed */
// The class instance for the [ICScannerFunctionalUnitFlatbed] class.
var (
	ICScannerFunctionalUnitFlatbedClass     _ICScannerFunctionalUnitFlatbedClass
	ICScannerFunctionalUnitFlatbedClassOnce sync.Once
)

func getICScannerFunctionalUnitFlatbedClass() _ICScannerFunctionalUnitFlatbedClass {
	ICScannerFunctionalUnitFlatbedClassOnce.Do(func() {
		ICScannerFunctionalUnitFlatbedClass = _ICScannerFunctionalUnitFlatbedClass{objc.GetClass("ICScannerFunctionalUnitFlatbed")}
	})
	return ICScannerFunctionalUnitFlatbedClass
}

type _ICScannerFunctionalUnitFlatbedClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ICScannerFunctionalUnitFlatbed */
// An interface definition for the [ICScannerFunctionalUnitFlatbed] class.
type IICScannerFunctionalUnitFlatbed interface {
	IICScannerFunctionalUnit
	
/* debug [class_interface_properties]: Properties for ICScannerFunctionalUnitFlatbed */
	// properties:
	DocumentType() unsafe.Pointer
	SetDocumentType(value unsafe.Pointer)
	DocumentSize() Size get /* not a class type */
	SetDocumentSize(value Size get /* not a class type */)
	SupportedDocumentTypes() unsafe.Pointer
	SetSupportedDocumentTypes(value unsafe.Pointer)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ICScannerFunctionalUnitFlatbed */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ICScannerFunctionalUnitFlatbed */
// Alloc allocates a new instance without initialization.
func (ic _ICScannerFunctionalUnitFlatbedClass) Alloc() ICScannerFunctionalUnitFlatbed {
	rv := objc.Send[ICScannerFunctionalUnitFlatbed](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _ICScannerFunctionalUnitFlatbedClass) New() ICScannerFunctionalUnitFlatbed {
	rv := objc.Send[ICScannerFunctionalUnitFlatbed](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ICScannerFunctionalUnitFlatbed) Init() ICScannerFunctionalUnitFlatbed {
	rv := objc.Send[ICScannerFunctionalUnitFlatbed](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ICScannerFunctionalUnitFlatbed) Autorelease() ICScannerFunctionalUnitFlatbed {
	rv := objc.Send[ICScannerFunctionalUnitFlatbed](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewICScannerFunctionalUnitFlatbed creates a new ICScannerFunctionalUnitFlatbed instance.
func NewICScannerFunctionalUnitFlatbed() ICScannerFunctionalUnitFlatbed {
	return getICScannerFunctionalUnitFlatbedClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ICScannerFunctionalUnitFlatbed */
// An object that represents the flatbed unit on a scanner.


// An object that represents the flatbed unit on a scanner.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore/ICScannerFunctionalUnitFlatbed
type ICScannerFunctionalUnitFlatbed struct {
	ICScannerFunctionalUnit
}

// ICScannerFunctionalUnitFlatbedFrom constructs a [ICScannerFunctionalUnitFlatbed] from an unsafe.Pointer.
//
// An object that represents the flatbed unit on a scanner.
func ICScannerFunctionalUnitFlatbedFrom(ptr unsafe.Pointer) ICScannerFunctionalUnitFlatbed {
	return ICScannerFunctionalUnitFlatbed{
		ICScannerFunctionalUnit: ICScannerFunctionalUnitFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ICScannerFunctionalUnitFlatbed *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ICScannerFunctionalUnitFlatbed */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ICScannerFunctionalUnitFlatbed */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ICScannerFunctionalUnitFlatbed */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ICScannerFunctionalUnitFlatbed */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunitflatbed/1507793-documenttype
func (i_ ICScannerFunctionalUnitFlatbed) DocumentType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("documentType"))
	return rv
}/* debug [instance_properties/getter]: documentType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunitflatbed/1507793-documenttype
func (i_ ICScannerFunctionalUnitFlatbed) SetDocumentType(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDocumentType:"), value)
}/* debug [instance_properties/setter]: documentType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunitflatbed/1507971-documentsize
func (i_ ICScannerFunctionalUnitFlatbed) DocumentSize() Size get /* not a class type */ {
	rv := objc.Send[objc.ID](i_.ID, objc.Sel("documentSize"))
	return rv
}/* debug [instance_properties/getter]: documentSize */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunitflatbed/1507971-documentsize
func (i_ ICScannerFunctionalUnitFlatbed) SetDocumentSize(value Size get /* not a class type */) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDocumentSize:"), value)
}/* debug [instance_properties/setter]: documentSize */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunitflatbed/1508067-supporteddocumenttypes
func (i_ ICScannerFunctionalUnitFlatbed) SupportedDocumentTypes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("supportedDocumentTypes"))
	return rv
}/* debug [instance_properties/getter]: supportedDocumentTypes */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/imagecapturecore/icscannerfunctionalunitflatbed/1508067-supporteddocumenttypes
func (i_ ICScannerFunctionalUnitFlatbed) SetSupportedDocumentTypes(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSupportedDocumentTypes:"), value)
}/* debug [instance_properties/setter]: supportedDocumentTypes */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ICScannerFunctionalUnitFlatbed */



