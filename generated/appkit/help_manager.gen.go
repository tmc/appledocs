
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [HelpManager] class.
var HelpManagerClass _HelpManagerClass

func init() {
	HelpManagerClass = _HelpManagerClass{objc.GetClass("NSHelpManager")}
}

type _HelpManagerClass struct {
	objc.Class
}

// An interface definition for the [HelpManager] class.
type IHelpManager interface {
	ID() objc.ID
	ContextHelpForObject(object objc.ID) unsafe.Pointer
	FindStringInBook(query unsafe.Pointer, book unsafe.Pointer)
	OpenHelpAnchorInBook(anchor unsafe.Pointer, book unsafe.Pointer)
	RegisterBooksInBundle(bundle unsafe.Pointer) bool
	RemoveContextHelpForObject(object objc.ID)
	SetContextHelpForObject(attrString unsafe.Pointer, object objc.ID)
	ShowContextHelpForObjectLocationHint(object objc.ID, pt unsafe.Pointer) bool
}

type HelpManager struct {
	id objc.ID
}

func HelpManagerFrom(ptr unsafe.Pointer) HelpManager {
	return HelpManager{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (h_ HelpManager) ID() objc.ID {
	return h_.id
}

// Alloc allocates a new instance without initialization.
func (hc _HelpManagerClass) Alloc() HelpManager {
	rv := objc.Send[HelpManager](objc.ID(hc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (hc _HelpManagerClass) New() HelpManager {
	rv := objc.Send[HelpManager](objc.ID(hc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewHelpManager creates and returns a new initialized instance.
func NewHelpManager() HelpManager {
	return HelpManagerClass.New()
}

// Init initializes the instance.
func (h_ HelpManager) Init() HelpManager {
	rv := objc.Send[HelpManager](h_.ID(), selInit)
	return rv
}
// Returns context-sensitive help for an object. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSHelpManager/contextHelp(for:)
func (h_ HelpManager) ContextHelpForObject(object objc.ID) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID(), objc.RegisterName("contextHelpForObject:"), object)
	return rv
}
// Performs a search for the specified string in the specified book. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSHelpManager/find(_:inBook:)
func (h_ HelpManager) FindStringInBook(query unsafe.Pointer, book unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID(), objc.RegisterName("findString:inBook:"), query, book)
}
// Finds and displays the text at the given anchor location in the given book. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSHelpManager/openHelpAnchor(_:inBook:)
func (h_ HelpManager) OpenHelpAnchorInBook(anchor unsafe.Pointer, book unsafe.Pointer) {
	objc.Send[objc.ID](h_.ID(), objc.RegisterName("openHelpAnchor:inBook:"), anchor, book)
}
// Registers one or more help books in the given bundle. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSHelpManager/registerBooks(in:)
func (h_ HelpManager) RegisterBooksInBundle(bundle unsafe.Pointer) bool {
	rv := objc.Send[bool](h_.ID(), objc.RegisterName("registerBooksInBundle:"), bundle)
	return rv
}
// Removes the association between an object and its context-sensitive help. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSHelpManager/removeContextHelp(for:)
func (h_ HelpManager) RemoveContextHelpForObject(object objc.ID) {
	objc.Send[objc.ID](h_.ID(), objc.RegisterName("removeContextHelpForObject:"), object)
}
// Associates help content with an object. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSHelpManager/setContextHelp(_:for:)
func (h_ HelpManager) SetContextHelpForObject(attrString unsafe.Pointer, object objc.ID) {
	objc.Send[objc.ID](h_.ID(), objc.RegisterName("setContextHelp:forObject:"), attrString, object)
}
// Displays the context-sensitive help for a given object at or near the point on the screen specified by a given point. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSHelpManager/showContextHelp(for:locationHint:)
func (h_ HelpManager) ShowContextHelpForObjectLocationHint(object objc.ID, pt unsafe.Pointer) bool {
	rv := objc.Send[bool](h_.ID(), objc.RegisterName("showContextHelpForObject:locationHint:"), object, pt)
	return rv
}
