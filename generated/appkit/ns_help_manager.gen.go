// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [HelpManager] class.
type IHelpManager interface {
	objectivec.IObject
	// properties:
	// methods:
	ContextHelpForObject(object objectivec.IObject) objc.IObject /* cross-framework: AttributedString */
	FindStringInBook(query objc.IObject /* cross-framework: NSString */, book objc.IObject /* cross-framework: HelpBookName */)
	OpenHelpAnchorInBook(anchor objc.IObject /* cross-framework: HelpAnchorName */, book objc.IObject /* cross-framework: HelpBookName */)
	RegisterBooksInBundle(bundle objc.IObject /* cross-framework: Bundle */) bool
	RemoveContextHelpForObject(object objectivec.IObject)
	SetContextHelpForObject(attrString objc.IObject /* cross-framework: AttributedString */, object objectivec.IObject)
	ShowContextHelpForObjectLocationHint(object objectivec.IObject, pt objc.IObject /* cross-framework: Point */) bool
}

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

// Alloc allocates a new instance without initialization.
func (hc _HelpManagerClass) Alloc() HelpManager {
	rv := objc.Send[HelpManager](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSHelpManager/isContextHelpModeActive
func (hc _HelpManagerClass) ContextHelpModeActive() bool {
	rv := objc.Send[bool](objc.ID(hc.class), objc.Sel("contextHelpModeActive"))
	return rv
}

// Returns the shared instance, creating it if it does not already exist.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSHelpManager/shared
func (hc _HelpManagerClass) SharedHelpManager() HelpManager {
	rv := objc.Send[HelpManager](objc.ID(hc.class), objc.Sel("sharedHelpManager"))
	return rv
}

// Returns context-sensitive help for an object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSHelpManager/contextHelp(for:)
func (h_ HelpManager) ContextHelpForObject(object objectivec.IObject) objc.IObject /* cross-framework: AttributedString */ {
	rv := objc.Send[foundation.AttributedString](h_.ID, objc.Sel("contextHelpForObject:"), object)
	return rv
}


// Performs a search for the specified string in the specified book.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSHelpManager/find(_:inBook:)
func (h_ HelpManager) FindStringInBook(query objc.IObject /* cross-framework: NSString */, book objc.IObject /* cross-framework: HelpBookName */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("findString:inBook:"), query, book)
}


// Finds and displays the text at the given anchor location in the given book.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSHelpManager/openHelpAnchor(_:inBook:)
func (h_ HelpManager) OpenHelpAnchorInBook(anchor objc.IObject /* cross-framework: HelpAnchorName */, book objc.IObject /* cross-framework: HelpBookName */) {
	objc.Send[objc.ID](h_.ID, objc.Sel("openHelpAnchor:inBook:"), anchor, book)
}


// Registers one or more help books in the given bundle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSHelpManager/registerBooks(in:)
func (h_ HelpManager) RegisterBooksInBundle(bundle objc.IObject /* cross-framework: Bundle */) bool {
	rv := objc.Send[bool](h_.ID, objc.Sel("registerBooksInBundle:"), bundle)
	return rv
}


// Removes the association between an object and its context-sensitive help.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSHelpManager/removeContextHelp(for:)
func (h_ HelpManager) RemoveContextHelpForObject(object objectivec.IObject) {
	objc.Send[objc.ID](h_.ID, objc.Sel("removeContextHelpForObject:"), object)
}


// Associates help content with an object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSHelpManager/setContextHelp(_:for:)
func (h_ HelpManager) SetContextHelpForObject(attrString objc.IObject /* cross-framework: AttributedString */, object objectivec.IObject) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setContextHelp:forObject:"), attrString, object)
}


// Displays the context-sensitive help for a given object at or near the point on the screen specified by a given point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSHelpManager/showContextHelp(for:locationHint:)
func (h_ HelpManager) ShowContextHelpForObjectLocationHint(object objectivec.IObject, pt objc.IObject /* cross-framework: Point */) bool {
	rv := objc.Send[bool](h_.ID, objc.Sel("showContextHelpForObject:locationHint:"), object, pt)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSHelpManager/isContextHelpModeActive
func (h_ HelpManager) ContextHelpModeActive() bool {
	rv := objc.Send[bool](h_.ID, objc.Sel("contextHelpModeActive"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSHelpManager/isContextHelpModeActive
func (h_ HelpManager) SetContextHelpModeActive(value bool) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setContextHelpModeActive:"), value)
}


// Returns the shared instance, creating it if it does not already exist.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSHelpManager/shared
func (h_ HelpManager) SharedHelpManager() IHelpManager {
	rv := objc.Send[HelpManager](h_.ID, objc.Sel("sharedHelpManager"))
	return rv
}



