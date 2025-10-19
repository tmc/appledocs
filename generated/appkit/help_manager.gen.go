// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [HelpManager] class.
var helpManagerClass = _HelpManagerClass{objc.GetClass("NSHelpManager")}

type _HelpManagerClass struct {
	class objc.Class
}

// An interface definition for the [HelpManager] class.
type IHelpManager interface {
	objectivec.IObject
	ContextHelpForObject(object objc.ID) unsafe.Pointer
	FindStringInBook(query string, book unsafe.Pointer)
	OpenHelpAnchorInBook(anchor unsafe.Pointer, book unsafe.Pointer)
	RegisterBooksInBundle(bundle unsafe.Pointer) bool
	RemoveContextHelpForObject(object objc.ID)
	SetContextHelpForObject(attrString unsafe.Pointer, object objc.ID)
	ShowContextHelpForObjectLocationHint(object objc.ID, pt unsafe.Pointer) bool
}

// An object for displaying online help for an app. [Full Topic]
//
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

// New creates and returns a new instance with a +1 retain count.
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
	return helpManagerClass.New()
}


// Returns context-sensitive help for an object. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSHelpManager/contextHelp(for:)
func (h_ HelpManager) ContextHelpForObject(object objc.ID) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("contextHelpForObject:"), object)
	return rv
}
// Performs a search for the specified string in the specified book. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSHelpManager/find(_:inBook:)
func (h_ HelpManager) FindStringInBook(query string, book unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("findString:inBook:"), objc.String(query), book)
}
// Finds and displays the text at the given anchor location in the given book. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSHelpManager/openHelpAnchor(_:inBook:)
func (h_ HelpManager) OpenHelpAnchorInBook(anchor unsafe.Pointer, book unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID, objc.Sel("openHelpAnchor:inBook:"), anchor, book)
}
// Registers one or more help books in the given bundle. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSHelpManager/registerBooks(in:)
func (h_ HelpManager) RegisterBooksInBundle(bundle unsafe.Pointer) bool {
	rv := objc.Send[bool](h_.ID, objc.Sel("registerBooksInBundle:"), bundle)
	return rv
}
// Removes the association between an object and its context-sensitive help. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSHelpManager/removeContextHelp(for:)
func (h_ HelpManager) RemoveContextHelpForObject(object objc.ID) {
	objc.Send[objc.ID](h_.ID, objc.Sel("removeContextHelpForObject:"), object)
}
// Associates help content with an object. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSHelpManager/setContextHelp(_:for:)
func (h_ HelpManager) SetContextHelpForObject(attrString unsafe.Pointer, object objc.ID) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setContextHelp:forObject:"), attrString, object)
}
// Displays the context-sensitive help for a given object at or near the point on the screen specified by a given point. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSHelpManager/showContextHelp(for:locationHint:)
func (h_ HelpManager) ShowContextHelpForObjectLocationHint(object objc.ID, pt unsafe.Pointer) bool {
	rv := objc.Send[bool](h_.ID, objc.Sel("showContextHelpForObject:locationHint:"), object, pt)
	return rv
}


