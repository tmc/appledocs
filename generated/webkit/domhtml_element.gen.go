// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class DOMHTMLElement */


/* debug [class_header]: Header for DOMHTMLElement */
// The class instance for the [DOMHTMLElement] class.
var (
	DOMHTMLElementClass     _DOMHTMLElementClass
	DOMHTMLElementClassOnce sync.Once
)

func getDOMHTMLElementClass() _DOMHTMLElementClass {
	DOMHTMLElementClassOnce.Do(func() {
		DOMHTMLElementClass = _DOMHTMLElementClass{objc.GetClass("DOMHTMLElement")}
	})
	return DOMHTMLElementClass
}

type _DOMHTMLElementClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DOMHTMLElement */
// An interface definition for the [DOMHTMLElement] class.
type IDOMHTMLElement interface {
	IDOMElement
	
/* debug [class_interface_properties]: Properties for DOMHTMLElement */
	// properties:
	AccessKey() objc.IObject /* cross-framework: NSString */
	SetAccessKey(value objc.IObject /* cross-framework: NSString */)
	Children() IDOMHTMLCollection
	ContentEditable() objc.IObject /* cross-framework: NSString */
	SetContentEditable(value objc.IObject /* cross-framework: NSString */)
	Dir() objc.IObject /* cross-framework: NSString */
	SetDir(value objc.IObject /* cross-framework: NSString */)
	IdName() objc.IObject /* cross-framework: NSString */
	SetIdName(value objc.IObject /* cross-framework: NSString */)
	InnerText() objc.IObject /* cross-framework: NSString */
	SetInnerText(value objc.IObject /* cross-framework: NSString */)
	IsContentEditable() bool
	Lang() objc.IObject /* cross-framework: NSString */
	SetLang(value objc.IObject /* cross-framework: NSString */)
	OuterText() objc.IObject /* cross-framework: NSString */
	SetOuterText(value objc.IObject /* cross-framework: NSString */)
	TabIndex() int
	SetTabIndex(value int)
	Title() objc.IObject /* cross-framework: NSString */
	SetTitle(value objc.IObject /* cross-framework: NSString */)
	TitleDisplayString() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DOMHTMLElement */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DOMHTMLElement */
// Alloc allocates a new instance without initialization.
func (dc _DOMHTMLElementClass) Alloc() DOMHTMLElement {
	rv := objc.Send[DOMHTMLElement](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMHTMLElementClass) New() DOMHTMLElement {
	rv := objc.Send[DOMHTMLElement](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMHTMLElement) Init() DOMHTMLElement {
	rv := objc.Send[DOMHTMLElement](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMHTMLElement) Autorelease() DOMHTMLElement {
	rv := objc.Send[DOMHTMLElement](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMHTMLElement creates a new DOMHTMLElement instance.
func NewDOMHTMLElement() DOMHTMLElement {
	return getDOMHTMLElementClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DOMHTMLElement */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLElement
type DOMHTMLElement struct {
	DOMElement
}

// DOMHTMLElementFrom constructs a [DOMHTMLElement] from an unsafe.Pointer.
func DOMHTMLElementFrom(ptr unsafe.Pointer) DOMHTMLElement {
	return DOMHTMLElement{
		DOMElement: DOMElementFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DOMHTMLElement *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DOMHTMLElement */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DOMHTMLElement */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DOMHTMLElement */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DOMHTMLElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLElement/accessKey
func (d_ DOMHTMLElement) AccessKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("accessKey"))
	return rv
}/* debug [instance_properties/getter]: accessKey */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLElement/accessKey
func (d_ DOMHTMLElement) SetAccessKey(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setAccessKey:"), value)
}/* debug [instance_properties/setter]: accessKey */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLElement/children
func (d_ DOMHTMLElement) Children() IDOMHTMLCollection {
	rv := objc.Send[DOMHTMLCollection](d_.ID, objc.Sel("children"))
	return rv
}/* debug [instance_properties/getter]: children */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLElement/contentEditable
func (d_ DOMHTMLElement) ContentEditable() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("contentEditable"))
	return rv
}/* debug [instance_properties/getter]: contentEditable */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLElement/contentEditable
func (d_ DOMHTMLElement) SetContentEditable(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setContentEditable:"), value)
}/* debug [instance_properties/setter]: contentEditable */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLElement/dir
func (d_ DOMHTMLElement) Dir() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("dir"))
	return rv
}/* debug [instance_properties/getter]: dir */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLElement/dir
func (d_ DOMHTMLElement) SetDir(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDir:"), value)
}/* debug [instance_properties/setter]: dir */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLElement/idName
func (d_ DOMHTMLElement) IdName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("idName"))
	return rv
}/* debug [instance_properties/getter]: idName */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLElement/idName
func (d_ DOMHTMLElement) SetIdName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setIdName:"), value)
}/* debug [instance_properties/setter]: idName */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLElement/innerText
func (d_ DOMHTMLElement) InnerText() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("innerText"))
	return rv
}/* debug [instance_properties/getter]: innerText */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLElement/innerText
func (d_ DOMHTMLElement) SetInnerText(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setInnerText:"), value)
}/* debug [instance_properties/setter]: innerText */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLElement/isContentEditable
func (d_ DOMHTMLElement) IsContentEditable() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("isContentEditable"))
	return rv
}/* debug [instance_properties/getter]: isContentEditable */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLElement/lang
func (d_ DOMHTMLElement) Lang() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("lang"))
	return rv
}/* debug [instance_properties/getter]: lang */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLElement/lang
func (d_ DOMHTMLElement) SetLang(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setLang:"), value)
}/* debug [instance_properties/setter]: lang */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLElement/outerText
func (d_ DOMHTMLElement) OuterText() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("outerText"))
	return rv
}/* debug [instance_properties/getter]: outerText */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLElement/outerText
func (d_ DOMHTMLElement) SetOuterText(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setOuterText:"), value)
}/* debug [instance_properties/setter]: outerText */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLElement/tabIndex
func (d_ DOMHTMLElement) TabIndex() int {
	rv := objc.Send[int](d_.ID, objc.Sel("tabIndex"))
	return rv
}/* debug [instance_properties/getter]: tabIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLElement/tabIndex
func (d_ DOMHTMLElement) SetTabIndex(value int) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setTabIndex:"), value)
}/* debug [instance_properties/setter]: tabIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLElement/title
func (d_ DOMHTMLElement) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("title"))
	return rv
}/* debug [instance_properties/getter]: title */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLElement/title
func (d_ DOMHTMLElement) SetTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setTitle:"), value)
}/* debug [instance_properties/setter]: title */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMHTMLElement/titleDisplayString
func (d_ DOMHTMLElement) TitleDisplayString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("titleDisplayString"))
	return rv
}/* debug [instance_properties/getter]: titleDisplayString */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class DOMHTMLElement */



