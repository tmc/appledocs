// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [HelpManager] class.
var HelpManagerClass objc.Class

func init() {
	HelpManagerClass = objc.GetClass("NSHelpManager")
}

type HelpManager struct {
	objc.ID
}

func HelpManagerFrom(ptr unsafe.Pointer) HelpManager {
	return HelpManager{
		ID: objc.ID(ptr),
	}
}


// Returns context-sensitive help for an object. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSHelpManager/contextHelp(for:)
func (h_ HelpManager) ContextHelpForObject(object objc.ID) unsafe.Pointer {
	sel := objc.RegisterName("contextHelpForObject:")
	ret := h_.ID.Send(sel, object)
	return unsafe.Pointer(ret)
}
// Performs a search for the specified string in the specified book. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSHelpManager/find(_:inBook:)
func (h_ HelpManager) FindStringInBook(query unsafe.Pointer, book unsafe.Pointer) {
	sel := objc.RegisterName("findString:inBook:")
	h_.ID.Send(sel, query, book)
}
// Finds and displays the text at the given anchor location in the given book. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSHelpManager/openHelpAnchor(_:inBook:)
func (h_ HelpManager) OpenHelpAnchorInBook(anchor unsafe.Pointer, book unsafe.Pointer) {
	sel := objc.RegisterName("openHelpAnchor:inBook:")
	h_.ID.Send(sel, anchor, book)
}
// Registers one or more help books in the given bundle. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSHelpManager/registerBooks(in:)
func (h_ HelpManager) RegisterBooksInBundle(bundle unsafe.Pointer) bool {
	sel := objc.RegisterName("registerBooksInBundle:")
	ret := h_.ID.Send(sel, bundle)
	return ret != 0
}
// Removes the association between an object and its context-sensitive help. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSHelpManager/removeContextHelp(for:)
func (h_ HelpManager) RemoveContextHelpForObject(object objc.ID) {
	sel := objc.RegisterName("removeContextHelpForObject:")
	h_.ID.Send(sel, object)
}
// Associates help content with an object. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSHelpManager/setContextHelp(_:for:)
func (h_ HelpManager) SetContextHelpForObject(attrString unsafe.Pointer, object objc.ID) {
	sel := objc.RegisterName("setContextHelp:forObject:")
	h_.ID.Send(sel, attrString, object)
}
// Displays the context-sensitive help for a given object at or near the point on the screen specified by a given point. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSHelpManager/showContextHelp(for:locationHint:)
func (h_ HelpManager) ShowContextHelpForObjectLocationHint(object objc.ID, pt unsafe.Pointer) bool {
	sel := objc.RegisterName("showContextHelpForObject:locationHint:")
	ret := h_.ID.Send(sel, object, pt)
	return ret != 0
}

