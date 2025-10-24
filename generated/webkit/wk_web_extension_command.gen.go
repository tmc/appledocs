// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class WKWebExtensionCommand */

/* debug [class_header]: Header for WKWebExtensionCommand */
// The class instance for the [WebExtensionCommand] class.
var (
	WebExtensionCommandClass     _WebExtensionCommandClass
	WebExtensionCommandClassOnce sync.Once
)

func getWebExtensionCommandClass() _WebExtensionCommandClass {
	WebExtensionCommandClassOnce.Do(func() {
		WebExtensionCommandClass = _WebExtensionCommandClass{objc.GetClass("WKWebExtensionCommand")}
	})
	return WebExtensionCommandClass
}

type _WebExtensionCommandClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for WebExtensionCommand */
// An interface definition for the [WebExtensionCommand] class.
type IWebExtensionCommand interface {
	objectivec.IObject

	/* debug [class_interface_properties]: Properties for WebExtensionCommand */
	// properties:
	ActivationKey() objc.IObject /* cross-framework: NSString */
	SetActivationKey(value objc.IObject /* cross-framework: NSString */)
	Identifier() objc.IObject        /* cross-framework: NSString */
	MenuItem() MenuElement           /* not a class type */
	ModifierFlags() KeyModifierFlags /* not a class type */
	SetModifierFlags(value KeyModifierFlags /* not a class type */)
	Title() objc.IObject /* cross-framework: NSString */
	WebExtensionContext() IWKWebExtensionContext
	Id() objc.IObject /* cross-framework: NSString */
	SetId(value objc.IObject /* cross-framework: NSString */)
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for WebExtensionCommand */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for WebExtensionCommand */
// Alloc allocates a new instance without initialization.
func (wc _WebExtensionCommandClass) Alloc() WebExtensionCommand {
	rv := objc.Send[WebExtensionCommand](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (wc _WebExtensionCommandClass) New() WebExtensionCommand {
	rv := objc.Send[WebExtensionCommand](objc.ID(wc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (w_ WebExtensionCommand) Init() WebExtensionCommand {
	rv := objc.Send[WebExtensionCommand](w_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w_ WebExtensionCommand) Autorelease() WebExtensionCommand {
	rv := objc.Send[WebExtensionCommand](w_.ID, objc.Sel("autorelease"))
	return rv
}

// NewWebExtensionCommand creates a new WebExtensionCommand instance.
func NewWebExtensionCommand() WebExtensionCommand {
	return getWebExtensionCommandClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for WebExtensionCommand */
// An object that encapsulates the properties for an individual web extension command.
//
// Provides access to command properties such as a unique identifier, a descriptive title, and shortcut keys. Commands can be used by a web extension to perform specific actions within a web extension context, such toggling features, or interacting with web content. These commands enhance the functionality of the extension by allowing users to invoke actions quickly.

// An object that encapsulates the properties for an individual web extension command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/Command
type WebExtensionCommand struct {
	objectivec.Object
}

// WebExtensionCommandFrom constructs a [WebExtensionCommand] from an unsafe.Pointer.
//
// An object that encapsulates the properties for an individual web extension command.
func WebExtensionCommandFrom(ptr unsafe.Pointer) WebExtensionCommand {
	return WebExtensionCommand{objectivec.Object{objc.ID(ptr)}}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for WebExtensionCommand */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for WebExtensionCommand */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for WebExtensionCommand */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for WebExtensionCommand */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for WebExtensionCommand */

// The primary key used to trigger the command, distinct from any modifier flags.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/Command/activationKey
func (w_ WebExtensionCommand) ActivationKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("activationKey"))
	return rv
} /* debug [instance_properties/getter]: activationKey */

// The primary key used to trigger the command, distinct from any modifier flags.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/Command/activationKey
func (w_ WebExtensionCommand) SetActivationKey(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setActivationKey:"), value)
} /* debug [instance_properties/setter]: activationKey */

// A unique identifier for the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/Command/id
func (w_ WebExtensionCommand) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("identifier"))
	return rv
} /* debug [instance_properties/getter]: identifier */

// A menu item representation of the web extension command for use in menus.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/Command/menuItem
func (w_ WebExtensionCommand) MenuItem() MenuElement /* not a class type */ {
	rv := objc.Send[MenuElement](w_.ID, objc.Sel("menuItem"))
	return rv
} /* debug [instance_properties/getter]: menuItem */

// The modifier flags used with the activation key to trigger the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/Command/modifierFlags
func (w_ WebExtensionCommand) ModifierFlags() KeyModifierFlags /* not a class type */ {
	rv := objc.Send[KeyModifierFlags](w_.ID, objc.Sel("modifierFlags"))
	return rv
} /* debug [instance_properties/getter]: modifierFlags */

// The modifier flags used with the activation key to trigger the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/Command/modifierFlags
func (w_ WebExtensionCommand) SetModifierFlags(value KeyModifierFlags /* not a class type */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setModifierFlags:"), value)
} /* debug [instance_properties/setter]: modifierFlags */

// A descriptive title for the command to help discoverability.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/Command/title
func (w_ WebExtensionCommand) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("title"))
	return rv
} /* debug [instance_properties/getter]: title */

// The web extension context associated with the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/Command/webExtensionContext
func (w_ WebExtensionCommand) WebExtensionContext() IWKWebExtensionContext {
	rv := objc.Send[WebExtensionContext](w_.ID, objc.Sel("webExtensionContext"))
	return rv
} /* debug [instance_properties/getter]: webExtensionContext */

// A unique identifier for the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/command/id
func (w_ WebExtensionCommand) Id() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("id"))
	return rv
} /* debug [instance_properties/getter]: id */

// A unique identifier for the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/command/id
func (w_ WebExtensionCommand) SetId(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setId:"), value)
} /* debug [instance_properties/setter]: id */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class WKWebExtensionCommand */
