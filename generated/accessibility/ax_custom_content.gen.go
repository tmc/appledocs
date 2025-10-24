// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [AXCustomContent] class.
type IAXCustomContent interface {
	objectivec.IObject
	// properties:
	AttributedLabel() objc.IObject /* cross-framework: AttributedString */
	SetAttributedLabel(value objc.IObject /* cross-framework: AttributedString */)
	AttributedValue() objc.IObject /* cross-framework: AttributedString */
	SetAttributedValue(value objc.IObject /* cross-framework: AttributedString */)
	Importance() unsafe.Pointer
	SetImportance(value unsafe.Pointer)
	Label() objc.IObject /* cross-framework: NSString */
	SetLabel(value objc.IObject /* cross-framework: NSString */)
	Value() objc.IObject /* cross-framework: NSString */
	SetValue(value objc.IObject /* cross-framework: NSString */)
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (ac _AXCustomContentClass) Alloc() AXCustomContent {
	rv := objc.Send[AXCustomContent](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// A localized attributed string that identifies the label for this content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/accessibility/axcustomcontent/attributedlabel
func (a_ AXCustomContent) AttributedLabel() objc.IObject /* cross-framework: AttributedString */ {
	rv := objc.Send[foundation.AttributedString](a_.ID, objc.Sel("attributedLabel"))
	return rv
}


// A localized attributed string that identifies the label for this content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/accessibility/axcustomcontent/attributedlabel
func (a_ AXCustomContent) SetAttributedLabel(value objc.IObject /* cross-framework: AttributedString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAttributedLabel:"), value)
}


// A localized attributed string that provides a value for the label.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/accessibility/axcustomcontent/attributedvalue
func (a_ AXCustomContent) AttributedValue() objc.IObject /* cross-framework: AttributedString */ {
	rv := objc.Send[foundation.AttributedString](a_.ID, objc.Sel("attributedValue"))
	return rv
}


// A localized attributed string that provides a value for the label.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/accessibility/axcustomcontent/attributedvalue
func (a_ AXCustomContent) SetAttributedValue(value objc.IObject /* cross-framework: AttributedString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAttributedValue:"), value)
}


// An object that determines when to output custom accessibility content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/accessibility/axcustomcontent/importance-swift.property
func (a_ AXCustomContent) Importance() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("importance"))
	return rv
}


// An object that determines when to output custom accessibility content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/accessibility/axcustomcontent/importance-swift.property
func (a_ AXCustomContent) SetImportance(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setImportance:"), value)
}


// A localized string that identifies the label for this content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/accessibility/axcustomcontent/label
func (a_ AXCustomContent) Label() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("label"))
	return rv
}


// A localized string that identifies the label for this content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/accessibility/axcustomcontent/label
func (a_ AXCustomContent) SetLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setLabel:"), value)
}


// A localized string that provides a value for the label.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/accessibility/axcustomcontent/value
func (a_ AXCustomContent) Value() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("value"))
	return rv
}


// A localized string that provides a value for the label.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/accessibility/axcustomcontent/value
func (a_ AXCustomContent) SetValue(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setValue:"), value)
}



