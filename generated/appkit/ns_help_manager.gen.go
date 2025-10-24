// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/vision"
)

/* debug [class.gen.go]: Generating class NSHelpManager */


/* debug [class_header]: Header for NSHelpManager */
// The class instance for the [HelpManager] class.
var (
	HelpManagerClass     _HelpManagerClass
	HelpManagerClassOnce sync.Once
)

func getHelpManagerClass() _HelpManagerClass {
	HelpManagerClassOnce.Do(func() {
		HelpManagerClass = _HelpManagerClass{objc.GetClass("NSHelpManager")}
	})
	return HelpManagerClass
}

type _HelpManagerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HelpManager */
// An interface definition for the [HelpManager] class.
type IHelpManager interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for HelpManager */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HelpManager */
	// methods:
	ContextHelpForObject(object objc.IObject) foundation.AttributedString
	FindStringInBook(query objc.IObject /* cross-framework: NSString */, book HelpBookName /* typedef */)
	OpenHelpAnchorInBook(anchor HelpAnchorName /* typedef */, book HelpBookName /* typedef */)
	RegisterBooksInBundle(bundle foundation.Bundle) bool
	RemoveContextHelpForObject(object objc.IObject)
	SetContextHelpForObject(attrString foundation.AttributedString, object objc.IObject)
	ShowContextHelpForObjectLocationHint(object objc.IObject, pt vision.Point) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HelpManager */
// Alloc allocates a new instance without initialization.
func (hc _HelpManagerClass) Alloc() HelpManager {
	rv := objc.Send[HelpManager](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _HelpManagerClass) New() HelpManager {
	rv := objc.Send[HelpManager](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HelpManager) Init() HelpManager {
	rv := objc.Send[HelpManager](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HelpManager) Autorelease() HelpManager {
	rv := objc.Send[HelpManager](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHelpManager creates a new HelpManager instance.
func NewHelpManager() HelpManager {
	return getHelpManagerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HelpManager */
// An object for displaying online help for an app.
//
// The class provides an approach to displaying online help. An app contains one object.


// An object for displaying online help for an app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSHelpManager
type HelpManager struct {
	objectivec.Object
}

// HelpManagerFrom constructs a [HelpManager] from an unsafe.Pointer.
//
// An object for displaying online help for an app.
func HelpManagerFrom(ptr unsafe.Pointer) HelpManager {
	return HelpManager{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HelpManager *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HelpManager */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HelpManager */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSHelpManager/isContextHelpModeActive
func (hc _HelpManagerClass) ContextHelpModeActive() bool {
	rv := objc.Send[bool](objc.ID(hc.class), objc.Sel("contextHelpModeActive"))
	return rv
}/* debug [class_properties_class/property]: contextHelpModeActive */

// Returns the shared instance, creating it if it does not already exist.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSHelpManager/shared
func (hc _HelpManagerClass) SharedHelpManager() HelpManager {
	rv := objc.Send[HelpManager](objc.ID(hc.class), objc.Sel("sharedHelpManager"))
	return rv
}/* debug [class_properties_class/property]: sharedHelpManager */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HelpManager */

// Returns context-sensitive help for an object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSHelpManager/contextHelp(for:)
func (h_ HelpManager) ContextHelpForObject(object objc.IObject) foundation.AttributedString {
	rv := objc.Send[foundation.AttributedString](h_.ID, objc.Sel("contextHelpForObject:"), object)
	return rv
}/* debug [instance_methods/method]: ContextHelpForObject */


// Performs a search for the specified string in the specified book.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSHelpManager/find(_:inBook:)
func (h_ HelpManager) FindStringInBook(query objc.IObject /* cross-framework: NSString */, book HelpBookName /* typedef */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("findString:inBook:"), query, book)
}/* debug [instance_methods/method]: FindStringInBook */


// Finds and displays the text at the given anchor location in the given book.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSHelpManager/openHelpAnchor(_:inBook:)
func (h_ HelpManager) OpenHelpAnchorInBook(anchor HelpAnchorName /* typedef */, book HelpBookName /* typedef */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("openHelpAnchor:inBook:"), anchor, book)
}/* debug [instance_methods/method]: OpenHelpAnchorInBook */


// Registers one or more help books in the given bundle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSHelpManager/registerBooks(in:)
func (h_ HelpManager) RegisterBooksInBundle(bundle foundation.Bundle) bool {
	rv := objc.Send[bool](h_.ID, objc.Sel("registerBooksInBundle:"), bundle)
	return rv
}/* debug [instance_methods/method]: RegisterBooksInBundle */


// Removes the association between an object and its context-sensitive help.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSHelpManager/removeContextHelp(for:)
func (h_ HelpManager) RemoveContextHelpForObject(object objc.IObject) {
	objc.Send[objc.ID](h_.ID, objc.Sel("removeContextHelpForObject:"), object)
}/* debug [instance_methods/method]: RemoveContextHelpForObject */


// Associates help content with an object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSHelpManager/setContextHelp(_:for:)
func (h_ HelpManager) SetContextHelpForObject(attrString foundation.AttributedString, object objc.IObject) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setContextHelp:forObject:"), attrString, object)
}/* debug [instance_methods/method]: SetContextHelpForObject */


// Displays the context-sensitive help for a given object at or near the point on the screen specified by a given point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSHelpManager/showContextHelp(for:locationHint:)
func (h_ HelpManager) ShowContextHelpForObjectLocationHint(object objc.IObject, pt vision.Point) bool {
	rv := objc.Send[bool](h_.ID, objc.Sel("showContextHelpForObject:locationHint:"), object, pt)
	return rv
}/* debug [instance_methods/method]: ShowContextHelpForObjectLocationHint */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HelpManager */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSHelpManager/isContextHelpModeActive
func (h_ HelpManager) ContextHelpModeActive() bool {
	rv := objc.Send[bool](h_.ID, objc.Sel("contextHelpModeActive"))
	return rv
}/* debug [instance_properties/getter]: contextHelpModeActive */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSHelpManager/isContextHelpModeActive
func (h_ HelpManager) SetContextHelpModeActive(value bool) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setContextHelpModeActive:"), value)
}/* debug [instance_properties/setter]: contextHelpModeActive */


// Returns the shared instance, creating it if it does not already exist.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSHelpManager/shared
func (h_ HelpManager) SharedHelpManager() IHelpManager {
	rv := objc.Send[HelpManager](h_.ID, objc.Sel("sharedHelpManager"))
	return rv
}/* debug [instance_properties/getter]: sharedHelpManager */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSHelpManager */



