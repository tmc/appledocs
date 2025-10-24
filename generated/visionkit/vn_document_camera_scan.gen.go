// Code generated from Apple documentation for VisionKit. DO NOT EDIT.

package visionkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VNDocumentCameraScan */

/* debug [class_header]: Header for VNDocumentCameraScan */
// The class instance for the [DocumentCameraScan] class.
var (
	DocumentCameraScanClass     _DocumentCameraScanClass
	DocumentCameraScanClassOnce sync.Once
)

func getDocumentCameraScanClass() _DocumentCameraScanClass {
	DocumentCameraScanClassOnce.Do(func() {
		DocumentCameraScanClass = _DocumentCameraScanClass{objc.GetClass("VNDocumentCameraScan")}
	})
	return DocumentCameraScanClass
}

type _DocumentCameraScanClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for DocumentCameraScan */
// An interface definition for the [DocumentCameraScan] class.
type IDocumentCameraScan interface {
	objectivec.IObject

	/* debug [class_interface_properties]: Properties for DocumentCameraScan */
	// properties:
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for DocumentCameraScan */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for DocumentCameraScan */
// Alloc allocates a new instance without initialization.
func (dc _DocumentCameraScanClass) Alloc() DocumentCameraScan {
	rv := objc.Send[DocumentCameraScan](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DocumentCameraScanClass) New() DocumentCameraScan {
	rv := objc.Send[DocumentCameraScan](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DocumentCameraScan) Init() DocumentCameraScan {
	rv := objc.Send[DocumentCameraScan](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DocumentCameraScan) Autorelease() DocumentCameraScan {
	rv := objc.Send[DocumentCameraScan](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDocumentCameraScan creates a new DocumentCameraScan instance.
func NewDocumentCameraScan() DocumentCameraScan {
	return getDocumentCameraScanClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for DocumentCameraScan */
// A single document scanned in the document camera.

// A single document scanned in the document camera.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/VisionKit/VNDocumentCameraScan
type DocumentCameraScan struct {
	objectivec.Object
}

// DocumentCameraScanFrom constructs a [DocumentCameraScan] from an unsafe.Pointer.
//
// A single document scanned in the document camera.
func DocumentCameraScanFrom(ptr unsafe.Pointer) DocumentCameraScan {
	return DocumentCameraScan{objectivec.Object{objc.ID(ptr)}}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for DocumentCameraScan */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for DocumentCameraScan */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for DocumentCameraScan */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for DocumentCameraScan */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for DocumentCameraScan */
/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class VNDocumentCameraScan */
