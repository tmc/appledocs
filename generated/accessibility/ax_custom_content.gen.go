// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AXCustomContent */


/* debug [class_header]: Header for AXCustomContent */
// The class instance for the [AXCustomContent] class.
var (
	AXCustomContentClass     _AXCustomContentClass
	AXCustomContentClassOnce sync.Once
)

func getAXCustomContentClass() _AXCustomContentClass {
	AXCustomContentClassOnce.Do(func() {
		AXCustomContentClass = _AXCustomContentClass{objc.GetClass("AXCustomContent")}
	})
	return AXCustomContentClass
}

type _AXCustomContentClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AXCustomContent */
// An interface definition for the [AXCustomContent] class.
type IAXCustomContent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AXCustomContent */
	// properties:
	AttributedLabel() foundation.AttributedString
	AttributedValue() foundation.AttributedString
	Importance() AXCustomContentImportance
	SetImportance(value AXCustomContentImportance)
	Label() objc.IObject /* cross-framework: NSString */
	Value() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AXCustomContent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AXCustomContent */
// Alloc allocates a new instance without initialization.
func (ac _AXCustomContentClass) Alloc() AXCustomContent {
	rv := objc.Send[AXCustomContent](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AXCustomContentClass) New() AXCustomContent {
	rv := objc.Send[AXCustomContent](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AXCustomContent) Init() AXCustomContent {
	rv := objc.Send[AXCustomContent](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AXCustomContent) Autorelease() AXCustomContent {
	rv := objc.Send[AXCustomContent](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAXCustomContent creates a new AXCustomContent instance.
func NewAXCustomContent() AXCustomContent {
	return getAXCustomContentClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AXCustomContent */
// Objects that define custom content and the timing of its output.
//
// An object contains the accessibility strings for the labels you apply to your accessibility content. Combine them with the protocol to allow your users to experience the content in a more appropriate manner for each assistive technology.


// Objects that define custom content and the timing of its output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXCustomContent
type AXCustomContent struct {
	objectivec.Object
}

// AXCustomContentFrom constructs a [AXCustomContent] from an unsafe.Pointer.
//
// Objects that define custom content and the timing of its output.
func AXCustomContentFrom(ptr unsafe.Pointer) AXCustomContent {
	return AXCustomContent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AXCustomContent */

// Creates new custom content with an attributed string and attributed value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXCustomContent/init(attributedLabel:attributedValue:)
func NewAXCustomContentWithAttributedLabelAttributedValue(label foundation.AttributedString, value foundation.AttributedString) AXCustomContent {
	rv := objc.Send[AXCustomContent](objc.ID(getAXCustomContentClass().class), objc.Sel("customContentWithAttributedLabel:attributedValue:"), label, value)
	return rv
}/* debug [class_init_methods/constructor]: NewAXCustomContentWithAttributedLabelAttributedValue */


// Creates new custom content with a label and value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXCustomContent/init(label:value:)
func NewAXCustomContentWithLabelValue(label objc.IObject /* cross-framework: NSString */, value objc.IObject /* cross-framework: NSString */) AXCustomContent {
	rv := objc.Send[AXCustomContent](objc.ID(getAXCustomContentClass().class), objc.Sel("customContentWithLabel:value:"), label, value)
	return rv
}/* debug [class_init_methods/constructor]: NewAXCustomContentWithLabelValue */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AXCustomContent */

// Creates new custom content with an attributed string and attributed value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXCustomContent/init(attributedLabel:attributedValue:)
func (ac _AXCustomContentClass) CustomContentWithAttributedLabelAttributedValue(label foundation.AttributedString, value foundation.AttributedString) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ac.class), objc.Sel("customContentWithAttributedLabel:attributedValue:"), label, value)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CustomContentWithAttributedLabelAttributedValue) */


// Creates new custom content with a label and value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXCustomContent/init(label:value:)
func (ac _AXCustomContentClass) CustomContentWithLabelValue(label objc.IObject /* cross-framework: NSString */, value objc.IObject /* cross-framework: NSString */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ac.class), objc.Sel("customContentWithLabel:value:"), label, value)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CustomContentWithLabelValue) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AXCustomContent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AXCustomContent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AXCustomContent */

// A localized attributed string that identifies the label for this content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXCustomContent/attributedLabel
func (a_ AXCustomContent) AttributedLabel() foundation.AttributedString {
	rv := objc.Send[foundation.AttributedString](a_.ID, objc.Sel("attributedLabel"))
	return rv
}/* debug [instance_properties/getter]: attributedLabel */


// A localized attributed string that provides a value for the label.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXCustomContent/attributedValue
func (a_ AXCustomContent) AttributedValue() foundation.AttributedString {
	rv := objc.Send[foundation.AttributedString](a_.ID, objc.Sel("attributedValue"))
	return rv
}/* debug [instance_properties/getter]: attributedValue */


// An object that determines when to output custom accessibility content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXCustomContent/importance-swift.property
func (a_ AXCustomContent) Importance() AXCustomContentImportance {
	rv := objc.Send[AXCustomContentImportance](a_.ID, objc.Sel("importance"))
	return rv
}/* debug [instance_properties/getter]: importance */


// An object that determines when to output custom accessibility content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXCustomContent/importance-swift.property
func (a_ AXCustomContent) SetImportance(value AXCustomContentImportance) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setImportance:"), value)
}/* debug [instance_properties/setter]: importance */


// A localized string that identifies the label for this content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXCustomContent/label
func (a_ AXCustomContent) Label() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("label"))
	return rv
}/* debug [instance_properties/getter]: label */


// A localized string that provides a value for the label.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXCustomContent/value
func (a_ AXCustomContent) Value() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("value"))
	return rv
}/* debug [instance_properties/getter]: value */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AXCustomContent */


