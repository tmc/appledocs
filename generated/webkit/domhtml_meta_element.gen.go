// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class DOMHTMLMetaElement */


/* debug [class_header]: Header for DOMHTMLMetaElement */
// The class instance for the [DOMHTMLMetaElement] class.
var (
	DOMHTMLMetaElementClass     _DOMHTMLMetaElementClass
	DOMHTMLMetaElementClassOnce sync.Once
)

func getDOMHTMLMetaElementClass() _DOMHTMLMetaElementClass {
	DOMHTMLMetaElementClassOnce.Do(func() {
		DOMHTMLMetaElementClass = _DOMHTMLMetaElementClass{objc.GetClass("DOMHTMLMetaElement")}
	})
	return DOMHTMLMetaElementClass
}

type _DOMHTMLMetaElementClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DOMHTMLMetaElement */
// An interface definition for the [DOMHTMLMetaElement] class.
type IDOMHTMLMetaElement interface {
	IDOMHTMLElement
	
/* debug [class_interface_properties]: Properties for DOMHTMLMetaElement */
	// properties:
	Content() objc.IObject /* cross-framework: NSString */
	SetContent(value objc.IObject /* cross-framework: NSString */)
	HttpEquiv() objc.IObject /* cross-framework: NSString */
	SetHttpEquiv(value objc.IObject /* cross-framework: NSString */)
	Name() objc.IObject /* cross-framework: NSString */
	SetName(value objc.IObject /* cross-framework: NSString */)
	Scheme() objc.IObject /* cross-framework: NSString */
	SetScheme(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DOMHTMLMetaElement */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DOMHTMLMetaElement */
// Alloc allocates a new instance without initialization.
func (dc _DOMHTMLMetaElementClass) Alloc() DOMHTMLMetaElement {
	rv := objc.Send[DOMHTMLMetaElement](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMHTMLMetaElementClass) New() DOMHTMLMetaElement {
	rv := objc.Send[DOMHTMLMetaElement](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMHTMLMetaElement) Init() DOMHTMLMetaElement {
	rv := objc.Send[DOMHTMLMetaElement](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMHTMLMetaElement) Autorelease() DOMHTMLMetaElement {
	rv := objc.Send[DOMHTMLMetaElement](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMHTMLMetaElement creates a new DOMHTMLMetaElement instance.
func NewDOMHTMLMetaElement() DOMHTMLMetaElement {
	return getDOMHTMLMetaElementClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DOMHTMLMetaElement */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLMetaElement
type DOMHTMLMetaElement struct {
	DOMHTMLElement
}

// DOMHTMLMetaElementFrom constructs a [DOMHTMLMetaElement] from an unsafe.Pointer.
func DOMHTMLMetaElementFrom(ptr unsafe.Pointer) DOMHTMLMetaElement {
	return DOMHTMLMetaElement{
		DOMHTMLElement: DOMHTMLElementFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DOMHTMLMetaElement *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DOMHTMLMetaElement */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DOMHTMLMetaElement */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DOMHTMLMetaElement */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DOMHTMLMetaElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLMetaElement/content
func (d_ DOMHTMLMetaElement) Content() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("content"))
	return rv
}/* debug [instance_properties/getter]: content */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLMetaElement/content
func (d_ DOMHTMLMetaElement) SetContent(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setContent:"), value)
}/* debug [instance_properties/setter]: content */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLMetaElement/httpEquiv
func (d_ DOMHTMLMetaElement) HttpEquiv() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("httpEquiv"))
	return rv
}/* debug [instance_properties/getter]: httpEquiv */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLMetaElement/httpEquiv
func (d_ DOMHTMLMetaElement) SetHttpEquiv(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setHttpEquiv:"), value)
}/* debug [instance_properties/setter]: httpEquiv */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLMetaElement/name
func (d_ DOMHTMLMetaElement) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLMetaElement/name
func (d_ DOMHTMLMetaElement) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setName:"), value)
}/* debug [instance_properties/setter]: name */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLMetaElement/scheme
func (d_ DOMHTMLMetaElement) Scheme() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("scheme"))
	return rv
}/* debug [instance_properties/getter]: scheme */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLMetaElement/scheme
func (d_ DOMHTMLMetaElement) SetScheme(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setScheme:"), value)
}/* debug [instance_properties/setter]: scheme */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class DOMHTMLMetaElement */



