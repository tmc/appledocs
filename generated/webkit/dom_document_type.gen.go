// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class DOMDocumentType */


/* debug [class_header]: Header for DOMDocumentType */
// The class instance for the [DOMDocumentType] class.
var (
	DOMDocumentTypeClass     _DOMDocumentTypeClass
	DOMDocumentTypeClassOnce sync.Once
)

func getDOMDocumentTypeClass() _DOMDocumentTypeClass {
	DOMDocumentTypeClassOnce.Do(func() {
		DOMDocumentTypeClass = _DOMDocumentTypeClass{objc.GetClass("DOMDocumentType")}
	})
	return DOMDocumentTypeClass
}

type _DOMDocumentTypeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DOMDocumentType */
// An interface definition for the [DOMDocumentType] class.
type IDOMDocumentType interface {
	IDOMNode
	
/* debug [class_interface_properties]: Properties for DOMDocumentType */
	// properties:
	Entities() IDOMNamedNodeMap
	InternalSubset() objc.IObject /* cross-framework: NSString */
	Name() objc.IObject /* cross-framework: NSString */
	Notations() IDOMNamedNodeMap
	PublicId() objc.IObject /* cross-framework: NSString */
	SystemId() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DOMDocumentType */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DOMDocumentType */
// Alloc allocates a new instance without initialization.
func (dc _DOMDocumentTypeClass) Alloc() DOMDocumentType {
	rv := objc.Send[DOMDocumentType](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMDocumentTypeClass) New() DOMDocumentType {
	rv := objc.Send[DOMDocumentType](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMDocumentType) Init() DOMDocumentType {
	rv := objc.Send[DOMDocumentType](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMDocumentType) Autorelease() DOMDocumentType {
	rv := objc.Send[DOMDocumentType](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMDocumentType creates a new DOMDocumentType instance.
func NewDOMDocumentType() DOMDocumentType {
	return getDOMDocumentTypeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DOMDocumentType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMDocumentType
type DOMDocumentType struct {
	DOMNode
}

// DOMDocumentTypeFrom constructs a [DOMDocumentType] from an unsafe.Pointer.
func DOMDocumentTypeFrom(ptr unsafe.Pointer) DOMDocumentType {
	return DOMDocumentType{
		DOMNode: DOMNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DOMDocumentType *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DOMDocumentType */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DOMDocumentType */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DOMDocumentType */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DOMDocumentType */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMDocumentType/entities
func (d_ DOMDocumentType) Entities() IDOMNamedNodeMap {
	rv := objc.Send[DOMNamedNodeMap](d_.ID, objc.Sel("entities"))
	return rv
}/* debug [instance_properties/getter]: entities */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMDocumentType/internalSubset
func (d_ DOMDocumentType) InternalSubset() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("internalSubset"))
	return rv
}/* debug [instance_properties/getter]: internalSubset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMDocumentType/name
func (d_ DOMDocumentType) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMDocumentType/notations
func (d_ DOMDocumentType) Notations() IDOMNamedNodeMap {
	rv := objc.Send[DOMNamedNodeMap](d_.ID, objc.Sel("notations"))
	return rv
}/* debug [instance_properties/getter]: notations */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMDocumentType/publicId
func (d_ DOMDocumentType) PublicId() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("publicId"))
	return rv
}/* debug [instance_properties/getter]: publicId */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMDocumentType/systemId
func (d_ DOMDocumentType) SystemId() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("systemId"))
	return rv
}/* debug [instance_properties/getter]: systemId */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class DOMDocumentType */



