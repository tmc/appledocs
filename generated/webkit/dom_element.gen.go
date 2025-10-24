// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class DOMElement */


/* debug [class_header]: Header for DOMElement */
// The class instance for the [DOMElement] class.
var (
	DOMElementClass     _DOMElementClass
	DOMElementClassOnce sync.Once
)

func getDOMElementClass() _DOMElementClass {
	DOMElementClassOnce.Do(func() {
		DOMElementClass = _DOMElementClass{objc.GetClass("DOMElement")}
	})
	return DOMElementClass
}

type _DOMElementClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DOMElement */
// An interface definition for the [DOMElement] class.
type IDOMElement interface {
	IDOMNode
	
/* debug [class_interface_properties]: Properties for DOMElement */
	// properties:
	ChildElementCount() objectivec.IObject
	ClassName() objc.IObject /* cross-framework: NSString */
	SetClassName(value objc.IObject /* cross-framework: NSString */)
	ClientHeight() int
	ClientLeft() int
	ClientTop() int
	ClientWidth() int
	FirstElementChild() IDOMElement
	InnerHTML() objc.IObject /* cross-framework: NSString */
	SetInnerHTML(value objc.IObject /* cross-framework: NSString */)
	InnerText() objc.IObject /* cross-framework: NSString */
	LastElementChild() IDOMElement
	NextElementSibling() IDOMElement
	OffsetHeight() int
	OffsetLeft() int
	OffsetParent() IDOMElement
	OffsetTop() int
	OffsetWidth() int
	OuterHTML() objc.IObject /* cross-framework: NSString */
	SetOuterHTML(value objc.IObject /* cross-framework: NSString */)
	PreviousElementSibling() IDOMElement
	ScrollHeight() int
	ScrollLeft() int
	SetScrollLeft(value int)
	ScrollTop() int
	SetScrollTop(value int)
	ScrollWidth() int
	Style() IDOMCSSStyleDeclaration
	TagName() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DOMElement */
	// methods:
	Image() appkit.Image
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DOMElement */
// Alloc allocates a new instance without initialization.
func (dc _DOMElementClass) Alloc() DOMElement {
	rv := objc.Send[DOMElement](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DOMElementClass) New() DOMElement {
	rv := objc.Send[DOMElement](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DOMElement) Init() DOMElement {
	rv := objc.Send[DOMElement](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DOMElement) Autorelease() DOMElement {
	rv := objc.Send[DOMElement](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDOMElement creates a new DOMElement instance.
func NewDOMElement() DOMElement {
	return getDOMElementClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DOMElement */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMElement
type DOMElement struct {
	DOMNode
}

// DOMElementFrom constructs a [DOMElement] from an unsafe.Pointer.
func DOMElementFrom(ptr unsafe.Pointer) DOMElement {
	return DOMElement{
		DOMNode: DOMNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DOMElement *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DOMElement */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DOMElement */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DOMElement */

// Returns an image associated with the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMElement/image()
func (d_ DOMElement) Image() appkit.Image {
	rv := objc.Send[appkit.Image](d_.ID, objc.Sel("image"))
	return rv
}/* debug [instance_methods/method]: Image */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DOMElement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMElement/childElementCount
func (d_ DOMElement) ChildElementCount() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](d_.ID, objc.Sel("childElementCount"))
	return rv
}/* debug [instance_properties/getter]: childElementCount */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMElement/className
func (d_ DOMElement) ClassName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("className"))
	return rv
}/* debug [instance_properties/getter]: className */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMElement/className
func (d_ DOMElement) SetClassName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setClassName:"), value)
}/* debug [instance_properties/setter]: className */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMElement/clientHeight
func (d_ DOMElement) ClientHeight() int {
	rv := objc.Send[int](d_.ID, objc.Sel("clientHeight"))
	return rv
}/* debug [instance_properties/getter]: clientHeight */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMElement/clientLeft
func (d_ DOMElement) ClientLeft() int {
	rv := objc.Send[int](d_.ID, objc.Sel("clientLeft"))
	return rv
}/* debug [instance_properties/getter]: clientLeft */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMElement/clientTop
func (d_ DOMElement) ClientTop() int {
	rv := objc.Send[int](d_.ID, objc.Sel("clientTop"))
	return rv
}/* debug [instance_properties/getter]: clientTop */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMElement/clientWidth
func (d_ DOMElement) ClientWidth() int {
	rv := objc.Send[int](d_.ID, objc.Sel("clientWidth"))
	return rv
}/* debug [instance_properties/getter]: clientWidth */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMElement/firstElementChild
func (d_ DOMElement) FirstElementChild() IDOMElement {
	rv := objc.Send[DOMElement](d_.ID, objc.Sel("firstElementChild"))
	return rv
}/* debug [instance_properties/getter]: firstElementChild */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMElement/innerHTML
func (d_ DOMElement) InnerHTML() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("innerHTML"))
	return rv
}/* debug [instance_properties/getter]: innerHTML */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMElement/innerHTML
func (d_ DOMElement) SetInnerHTML(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setInnerHTML:"), value)
}/* debug [instance_properties/setter]: innerHTML */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMElement/innerText
func (d_ DOMElement) InnerText() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("innerText"))
	return rv
}/* debug [instance_properties/getter]: innerText */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMElement/lastElementChild
func (d_ DOMElement) LastElementChild() IDOMElement {
	rv := objc.Send[DOMElement](d_.ID, objc.Sel("lastElementChild"))
	return rv
}/* debug [instance_properties/getter]: lastElementChild */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMElement/nextElementSibling
func (d_ DOMElement) NextElementSibling() IDOMElement {
	rv := objc.Send[DOMElement](d_.ID, objc.Sel("nextElementSibling"))
	return rv
}/* debug [instance_properties/getter]: nextElementSibling */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMElement/offsetHeight
func (d_ DOMElement) OffsetHeight() int {
	rv := objc.Send[int](d_.ID, objc.Sel("offsetHeight"))
	return rv
}/* debug [instance_properties/getter]: offsetHeight */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMElement/offsetLeft
func (d_ DOMElement) OffsetLeft() int {
	rv := objc.Send[int](d_.ID, objc.Sel("offsetLeft"))
	return rv
}/* debug [instance_properties/getter]: offsetLeft */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMElement/offsetParent
func (d_ DOMElement) OffsetParent() IDOMElement {
	rv := objc.Send[DOMElement](d_.ID, objc.Sel("offsetParent"))
	return rv
}/* debug [instance_properties/getter]: offsetParent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMElement/offsetTop
func (d_ DOMElement) OffsetTop() int {
	rv := objc.Send[int](d_.ID, objc.Sel("offsetTop"))
	return rv
}/* debug [instance_properties/getter]: offsetTop */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMElement/offsetWidth
func (d_ DOMElement) OffsetWidth() int {
	rv := objc.Send[int](d_.ID, objc.Sel("offsetWidth"))
	return rv
}/* debug [instance_properties/getter]: offsetWidth */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMElement/outerHTML
func (d_ DOMElement) OuterHTML() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("outerHTML"))
	return rv
}/* debug [instance_properties/getter]: outerHTML */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMElement/outerHTML
func (d_ DOMElement) SetOuterHTML(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setOuterHTML:"), value)
}/* debug [instance_properties/setter]: outerHTML */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMElement/previousElementSibling
func (d_ DOMElement) PreviousElementSibling() IDOMElement {
	rv := objc.Send[DOMElement](d_.ID, objc.Sel("previousElementSibling"))
	return rv
}/* debug [instance_properties/getter]: previousElementSibling */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMElement/scrollHeight
func (d_ DOMElement) ScrollHeight() int {
	rv := objc.Send[int](d_.ID, objc.Sel("scrollHeight"))
	return rv
}/* debug [instance_properties/getter]: scrollHeight */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMElement/scrollLeft
func (d_ DOMElement) ScrollLeft() int {
	rv := objc.Send[int](d_.ID, objc.Sel("scrollLeft"))
	return rv
}/* debug [instance_properties/getter]: scrollLeft */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMElement/scrollLeft
func (d_ DOMElement) SetScrollLeft(value int) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setScrollLeft:"), value)
}/* debug [instance_properties/setter]: scrollLeft */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMElement/scrollTop
func (d_ DOMElement) ScrollTop() int {
	rv := objc.Send[int](d_.ID, objc.Sel("scrollTop"))
	return rv
}/* debug [instance_properties/getter]: scrollTop */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMElement/scrollTop
func (d_ DOMElement) SetScrollTop(value int) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setScrollTop:"), value)
}/* debug [instance_properties/setter]: scrollTop */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMElement/scrollWidth
func (d_ DOMElement) ScrollWidth() int {
	rv := objc.Send[int](d_.ID, objc.Sel("scrollWidth"))
	return rv
}/* debug [instance_properties/getter]: scrollWidth */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMElement/style
func (d_ DOMElement) Style() IDOMCSSStyleDeclaration {
	rv := objc.Send[DOMCSSStyleDeclaration](d_.ID, objc.Sel("style"))
	return rv
}/* debug [instance_properties/getter]: style */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/DOMElement/tagName
func (d_ DOMElement) TagName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("tagName"))
	return rv
}/* debug [instance_properties/getter]: tagName */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class DOMElement */



