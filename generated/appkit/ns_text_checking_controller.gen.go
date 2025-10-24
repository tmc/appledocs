// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSTextCheckingController */


/* debug [class_header]: Header for NSTextCheckingController */
// The class instance for the [TextCheckingController] class.
var (
	TextCheckingControllerClass     _TextCheckingControllerClass
	TextCheckingControllerClassOnce sync.Once
)

func getTextCheckingControllerClass() _TextCheckingControllerClass {
	TextCheckingControllerClassOnce.Do(func() {
		TextCheckingControllerClass = _TextCheckingControllerClass{objc.GetClass("NSTextCheckingController")}
	})
	return TextCheckingControllerClass
}

type _TextCheckingControllerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TextCheckingController */
// An interface definition for the [TextCheckingController] class.
type ITextCheckingController interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for TextCheckingController */
	// properties:
	Client() TextCheckingClient /* not a class type */
	SetClient(value TextCheckingClient /* not a class type */)
	SpellCheckerDocumentTag() int
	SetSpellCheckerDocumentTag(value int)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TextCheckingController */
	// methods:
	DidChangeTextInRange(range_ corefoundation.Range)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TextCheckingController */
// Alloc allocates a new instance without initialization.
func (tc _TextCheckingControllerClass) Alloc() TextCheckingController {
	rv := objc.Send[TextCheckingController](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TextCheckingControllerClass) New() TextCheckingController {
	rv := objc.Send[TextCheckingController](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TextCheckingController) Init() TextCheckingController {
	rv := objc.Send[TextCheckingController](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TextCheckingController) Autorelease() TextCheckingController {
	rv := objc.Send[TextCheckingController](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextCheckingController creates a new TextCheckingController instance.
func NewTextCheckingController() TextCheckingController {
	return getTextCheckingControllerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TextCheckingController */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextCheckingController
type TextCheckingController struct {
	objectivec.Object
}

// TextCheckingControllerFrom constructs a [TextCheckingController] from an unsafe.Pointer.
func TextCheckingControllerFrom(ptr unsafe.Pointer) TextCheckingController {
	return TextCheckingController{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TextCheckingController *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TextCheckingController */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TextCheckingController */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TextCheckingController */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextCheckingController/didChangeText(in:)
func (t_ TextCheckingController) DidChangeTextInRange(range_ corefoundation.Range) {
	objc.Send[objc.ID](t_.ID, objc.Sel("didChangeTextInRange:"), range_)
}/* debug [instance_methods/method]: DidChangeTextInRange */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TextCheckingController */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextcheckingcontroller/client
func (t_ TextCheckingController) Client() TextCheckingClient /* not a class type */ {
	rv := objc.Send[TextCheckingClient](t_.ID, objc.Sel("client"))
	return rv
}/* debug [instance_properties/getter]: client */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextcheckingcontroller/client
func (t_ TextCheckingController) SetClient(value TextCheckingClient /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setClient:"), value)
}/* debug [instance_properties/setter]: client */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextcheckingcontroller/spellcheckerdocumenttag
func (t_ TextCheckingController) SpellCheckerDocumentTag() int {
	rv := objc.Send[int](t_.ID, objc.Sel("spellCheckerDocumentTag"))
	return rv
}/* debug [instance_properties/getter]: spellCheckerDocumentTag */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstextcheckingcontroller/spellcheckerdocumenttag
func (t_ TextCheckingController) SetSpellCheckerDocumentTag(value int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSpellCheckerDocumentTag:"), value)
}/* debug [instance_properties/setter]: spellCheckerDocumentTag */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSTextCheckingController */



