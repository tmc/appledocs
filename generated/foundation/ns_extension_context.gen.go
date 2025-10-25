// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSExtensionContext */


/* debug [class_header]: Header for NSExtensionContext */
// The class instance for the [ExtensionContext] class.
var (
	ExtensionContextClass     _ExtensionContextClass
	ExtensionContextClassOnce sync.Once
)

func getExtensionContextClass() _ExtensionContextClass {
	ExtensionContextClassOnce.Do(func() {
		ExtensionContextClass = _ExtensionContextClass{objc.GetClass("NSExtensionContext")}
	})
	return ExtensionContextClass
}

type _ExtensionContextClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ExtensionContext */
// An interface definition for the [ExtensionContext] class.
type IExtensionContext interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ExtensionContext */
	// properties:
	HostedViewMaximumAllowedSize() corefoundation.CGSize
	SetHostedViewMaximumAllowedSize(value corefoundation.CGSize)
	Intent() objectivec.IObject
	SetIntent(value objectivec.IObject)
	WidgetActiveDisplayMode() objectivec.IObject
	SetWidgetActiveDisplayMode(value objectivec.IObject)
	WidgetLargestAvailableDisplayMode() objectivec.IObject
	SetWidgetLargestAvailableDisplayMode(value objectivec.IObject)
	NSExtensionItemsAndErrorsKey() IString
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ExtensionContext */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ExtensionContext */
// Alloc allocates a new instance without initialization.
func (ec _ExtensionContextClass) Alloc() ExtensionContext {
	rv := objc.Send[ExtensionContext](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ec _ExtensionContextClass) New() ExtensionContext {
	rv := objc.Send[ExtensionContext](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ ExtensionContext) Init() ExtensionContext {
	rv := objc.Send[ExtensionContext](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ ExtensionContext) Autorelease() ExtensionContext {
	rv := objc.Send[ExtensionContext](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewExtensionContext creates a new ExtensionContext instance.
func NewExtensionContext() ExtensionContext {
	return getExtensionContextClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ExtensionContext */
// The host app context from which an app extension is invoked.
//
// When a host app sends a request to an app extension, it provides an extension context. For many app extensions, the most important part of the context is the data the user wants to work with, which is contained in the property.


// The host app context from which an app extension is invoked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExtensionContext
type ExtensionContext struct {
	objectivec.Object
}

// ExtensionContextFrom constructs a [ExtensionContext] from an unsafe.Pointer.
//
// The host app context from which an app extension is invoked.
func ExtensionContextFrom(ptr unsafe.Pointer) ExtensionContext {
	return ExtensionContext{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ExtensionContext *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ExtensionContext */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ExtensionContext */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ExtensionContext */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ExtensionContext */

// The maximum size for a Siri hosted view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsextensioncontext/hostedviewmaximumallowedsize
func (e_ ExtensionContext) HostedViewMaximumAllowedSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](e_.ID, objc.Sel("hostedViewMaximumAllowedSize"))
	return rv
}/* debug [instance_properties/getter]: hostedViewMaximumAllowedSize */


// The maximum size for a Siri hosted view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsextensioncontext/hostedviewmaximumallowedsize
func (e_ ExtensionContext) SetHostedViewMaximumAllowedSize(value corefoundation.CGSize) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setHostedViewMaximumAllowedSize:"), value)
}/* debug [instance_properties/setter]: hostedViewMaximumAllowedSize */


// Metadata for populating your share extensions interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsextensioncontext/intent
func (e_ ExtensionContext) Intent() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](e_.ID, objc.Sel("intent"))
	return rv
}/* debug [instance_properties/getter]: intent */


// Metadata for populating your share extensions interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsextensioncontext/intent
func (e_ ExtensionContext) SetIntent(value objectivec.IObject) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setIntent:"), value)
}/* debug [instance_properties/setter]: intent */


// The active display mode of the widget.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsextensioncontext/widgetactivedisplaymode
func (e_ ExtensionContext) WidgetActiveDisplayMode() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](e_.ID, objc.Sel("widgetActiveDisplayMode"))
	return rv
}/* debug [instance_properties/getter]: widgetActiveDisplayMode */


// The active display mode of the widget.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsextensioncontext/widgetactivedisplaymode
func (e_ ExtensionContext) SetWidgetActiveDisplayMode(value objectivec.IObject) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setWidgetActiveDisplayMode:"), value)
}/* debug [instance_properties/setter]: widgetActiveDisplayMode */


// The largest display mode the widget supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsextensioncontext/widgetlargestavailabledisplaymode
func (e_ ExtensionContext) WidgetLargestAvailableDisplayMode() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](e_.ID, objc.Sel("widgetLargestAvailableDisplayMode"))
	return rv
}/* debug [instance_properties/getter]: widgetLargestAvailableDisplayMode */


// The largest display mode the widget supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsextensioncontext/widgetlargestavailabledisplaymode
func (e_ ExtensionContext) SetWidgetLargestAvailableDisplayMode(value objectivec.IObject) {
	objc.Send[objc.ID](e_.ID, objc.Sel("setWidgetLargestAvailableDisplayMode:"), value)
}/* debug [instance_properties/setter]: widgetLargestAvailableDisplayMode */


// The extension items and errors key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsextensionitemsanderrorskey
func (e_ ExtensionContext) NSExtensionItemsAndErrorsKey() IString {
	rv := objc.Send[String](e_.ID, objc.Sel("NSExtensionItemsAndErrorsKey"))
	return rv
}/* debug [instance_properties/getter]: NSExtensionItemsAndErrorsKey */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSExtensionContext */


