
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [SavePanel] class.
var SavePanelClass _SavePanelClass

func init() {
	SavePanelClass = _SavePanelClass{objc.GetClass("NSSavePanel")}
}

type _SavePanelClass struct {
	objc.Class
}

// An interface definition for the [SavePanel] class.
type ISavePanel interface {
	ID() objc.ID
	BeginSheetModalForWindowCompletionHandler(window unsafe.Pointer, handler unsafe.Pointer)
	BeginWithCompletionHandler(handler unsafe.Pointer)
	Cancel(sender objc.ID)
	Ok(sender objc.ID)
	RunModal() unsafe.Pointer
	ValidateVisibleColumns()
}

type SavePanel struct {
	id objc.ID
}

func SavePanelFrom(ptr unsafe.Pointer) SavePanel {
	return SavePanel{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ SavePanel) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _SavePanelClass) Alloc() SavePanel {
	rv := objc.Send[SavePanel](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _SavePanelClass) New() SavePanel {
	rv := objc.Send[SavePanel](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewSavePanel creates and returns a new initialized instance.
func NewSavePanel() SavePanel {
	return SavePanelClass.New()
}

// Init initializes the instance.
func (s_ SavePanel) Init() SavePanel {
	rv := objc.Send[SavePanel](s_.ID(), selInit)
	return rv
}
// Creates a new Save panel and initializes it with default information. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSavePanel/savePanel
func (sc _SavePanelClass) SavePanel() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.Class), objc.RegisterName("savePanel"))
	return rv
}

// SavePanel_SavePanel creates a new instance via class method. [Full Topic]
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSavePanel/savePanel
func SavePanel_SavePanel() unsafe.Pointer {
	return SavePanelClass.SavePanel()
}
// Presents the panel as a modeless window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSavePanel/begin(completionHandler:)
func (s_ SavePanel) BeginWithCompletionHandler(handler unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("beginWithCompletionHandler:"), handler)
}
// Presents the panel as a sheet modal to the specified window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSavePanel/beginSheetModal(for:completionHandler:)
func (s_ SavePanel) BeginSheetModalForWindowCompletionHandler(window unsafe.Pointer, handler unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("beginSheetModalForWindow:completionHandler:"), window, handler)
}
// The action method that the panel calls when the user clicks the Cancel button. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSavePanel/cancel(_:)
func (s_ SavePanel) Cancel(sender objc.ID) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("cancel:"), sender)
}
// The action method that the panel calls when the user clicks the OK button. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSavePanel/ok(_:)
func (s_ SavePanel) Ok(sender objc.ID) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("ok:"), sender)
}
// Displays the panel and begins its event loop with the current working (or last-selected) directory as the default starting point. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSavePanel/runModal()
func (s_ SavePanel) RunModal() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID(), objc.RegisterName("runModal"))
	return rv
}
// Validates and reloads the browser columns visible in the panel. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSavePanel/validateVisibleColumns()
func (s_ SavePanel) ValidateVisibleColumns() {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("validateVisibleColumns"))
}
// The custom accessory view for the current app. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSavePanel/accessoryView
func (s_ SavePanel) AccessoryView() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID(), objc.RegisterName("accessoryView"))
	return rv
}
// SetAccessoryView sets the value of the accessoryView property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSavePanel/accessoryView
func (s_ SavePanel) SetAccessoryView(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setAccessoryView:"), value)
}
// An array of types that specify the files types to which you can save. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSavePanel/allowedContentTypes
func (s_ SavePanel) AllowedContentTypes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID(), objc.RegisterName("allowedContentTypes"))
	return rv
}
// SetAllowedContentTypes sets the value of the allowedContentTypes property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSavePanel/allowedContentTypes
func (s_ SavePanel) SetAllowedContentTypes(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setAllowedContentTypes:"), value)
}
// A Boolean value that indicates whether the panel allows the user to save files with a filename extension that’s not in the list of allowed types. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSavePanel/allowsOtherFileTypes
func (s_ SavePanel) AllowsOtherFileTypes() bool {
	rv := objc.Send[bool](s_.ID(), objc.RegisterName("allowsOtherFileTypes"))
	return rv
}
// SetAllowsOtherFileTypes sets the value of the allowsOtherFileTypes property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSavePanel/allowsOtherFileTypes
func (s_ SavePanel) SetAllowsOtherFileTypes(value bool) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setAllowsOtherFileTypes:"), value)
}
// A Boolean value that indicates whether the panel displays UI for creating directories. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSavePanel/canCreateDirectories
func (s_ SavePanel) CanCreateDirectories() bool {
	rv := objc.Send[bool](s_.ID(), objc.RegisterName("canCreateDirectories"))
	return rv
}
// SetCanCreateDirectories sets the value of the canCreateDirectories property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSavePanel/canCreateDirectories
func (s_ SavePanel) SetCanCreateDirectories(value bool) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setCanCreateDirectories:"), value)
}
// A Boolean value that indicates whether the panel displays UI for hiding or showing filename extensions. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSavePanel/canSelectHiddenExtension
func (s_ SavePanel) CanSelectHiddenExtension() bool {
	rv := objc.Send[bool](s_.ID(), objc.RegisterName("canSelectHiddenExtension"))
	return rv
}
// SetCanSelectHiddenExtension sets the value of the canSelectHiddenExtension property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSavePanel/canSelectHiddenExtension
func (s_ SavePanel) SetCanSelectHiddenExtension(value bool) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setCanSelectHiddenExtension:"), value)
}
// :The current type. If set to  , resets to the first allowed content type. Returns   if   is empty.   : Not used. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSavePanel/currentContentType
func (s_ SavePanel) CurrentContentType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID(), objc.RegisterName("currentContentType"))
	return rv
}
// SetCurrentContentType sets the value of the currentContentType property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSavePanel/currentContentType
func (s_ SavePanel) SetCurrentContentType(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setCurrentContentType:"), value)
}
// A custom object you use to manage interactions with an open or save panel. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSavePanel/delegate
func (s_ SavePanel) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID(), objc.RegisterName("delegate"))
	return rv
}
// SetDelegate sets the value of the delegate property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSavePanel/delegate
func (s_ SavePanel) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setDelegate:"), value)
}
// The current directory shown in the panel. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSavePanel/directoryURL
func (s_ SavePanel) DirectoryURL() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID(), objc.RegisterName("directoryURL"))
	return rv
}
// SetDirectoryURL sets the value of the directoryURL property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSavePanel/directoryURL
func (s_ SavePanel) SetDirectoryURL(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setDirectoryURL:"), value)
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSavePanel/identifier
func (s_ SavePanel) Identifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID(), objc.RegisterName("identifier"))
	return rv
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSavePanel/identifier
func (s_ SavePanel) SetIdentifier(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setIdentifier:"), value)
}
// A Boolean value that indicates whether whether the panel is expanded. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSavePanel/isExpanded
func (s_ SavePanel) Expanded() bool {
	rv := objc.Send[bool](s_.ID(), objc.RegisterName("expanded"))
	return rv
}
// A Boolean value that indicates whether to display filename extensions. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSavePanel/isExtensionHidden
func (s_ SavePanel) ExtensionHidden() bool {
	rv := objc.Send[bool](s_.ID(), objc.RegisterName("extensionHidden"))
	return rv
}
// SetExtensionHidden sets the value of the extensionHidden property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSavePanel/isExtensionHidden
func (s_ SavePanel) SetExtensionHidden(value bool) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setExtensionHidden:"), value)
}
// The message text displayed in the panel. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSavePanel/message
func (s_ SavePanel) Message() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID(), objc.RegisterName("message"))
	return rv
}
// SetMessage sets the value of the message property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSavePanel/message
func (s_ SavePanel) SetMessage(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setMessage:"), value)
}
// The label text displayed in front of the filename text field. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSavePanel/nameFieldLabel
func (s_ SavePanel) NameFieldLabel() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID(), objc.RegisterName("nameFieldLabel"))
	return rv
}
// SetNameFieldLabel sets the value of the nameFieldLabel property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSavePanel/nameFieldLabel
func (s_ SavePanel) SetNameFieldLabel(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setNameFieldLabel:"), value)
}
// The user-editable filename currently shown in the name field. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSavePanel/nameFieldStringValue
func (s_ SavePanel) NameFieldStringValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID(), objc.RegisterName("nameFieldStringValue"))
	return rv
}
// SetNameFieldStringValue sets the value of the nameFieldStringValue property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSavePanel/nameFieldStringValue
func (s_ SavePanel) SetNameFieldStringValue(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setNameFieldStringValue:"), value)
}
// The text to display in the default button. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSavePanel/prompt
func (s_ SavePanel) Prompt() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID(), objc.RegisterName("prompt"))
	return rv
}
// SetPrompt sets the value of the prompt property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSavePanel/prompt
func (s_ SavePanel) SetPrompt(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setPrompt:"), value)
}
// : Whether or not to show a control for selecting the type of the saved file.   The control shows the types in  . Default is  .   : Not used. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSavePanel/showsContentTypes
func (s_ SavePanel) ShowsContentTypes() bool {
	rv := objc.Send[bool](s_.ID(), objc.RegisterName("showsContentTypes"))
	return rv
}
// SetShowsContentTypes sets the value of the showsContentTypes property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSavePanel/showsContentTypes
func (s_ SavePanel) SetShowsContentTypes(value bool) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setShowsContentTypes:"), value)
}
// A Boolean value that indicates whether the panel displays files that are normally hidden from the user. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSavePanel/showsHiddenFiles
func (s_ SavePanel) ShowsHiddenFiles() bool {
	rv := objc.Send[bool](s_.ID(), objc.RegisterName("showsHiddenFiles"))
	return rv
}
// SetShowsHiddenFiles sets the value of the showsHiddenFiles property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSavePanel/showsHiddenFiles
func (s_ SavePanel) SetShowsHiddenFiles(value bool) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setShowsHiddenFiles:"), value)
}
// A Boolean value that indicates whether the panel displays the Tags field. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSavePanel/showsTagField
func (s_ SavePanel) ShowsTagField() bool {
	rv := objc.Send[bool](s_.ID(), objc.RegisterName("showsTagField"))
	return rv
}
// SetShowsTagField sets the value of the showsTagField property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSavePanel/showsTagField
func (s_ SavePanel) SetShowsTagField(value bool) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setShowsTagField:"), value)
}
// The tag names that you want to include on a saved file. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSavePanel/tagNames
func (s_ SavePanel) TagNames() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID(), objc.RegisterName("tagNames"))
	return rv
}
// SetTagNames sets the value of the tagNames property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSavePanel/tagNames
func (s_ SavePanel) SetTagNames(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setTagNames:"), value)
}
// The title of the panel. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSavePanel/title
func (s_ SavePanel) Title() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID(), objc.RegisterName("title"))
	return rv
}
// SetTitle sets the value of the title property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSavePanel/title
func (s_ SavePanel) SetTitle(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setTitle:"), value)
}
// A Boolean value that indicates whether the panel displays file packages as directories. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSavePanel/treatsFilePackagesAsDirectories
func (s_ SavePanel) TreatsFilePackagesAsDirectories() bool {
	rv := objc.Send[bool](s_.ID(), objc.RegisterName("treatsFilePackagesAsDirectories"))
	return rv
}
// SetTreatsFilePackagesAsDirectories sets the value of the treatsFilePackagesAsDirectories property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSavePanel/treatsFilePackagesAsDirectories
func (s_ SavePanel) SetTreatsFilePackagesAsDirectories(value bool) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setTreatsFilePackagesAsDirectories:"), value)
}
// A URL that contains the fully specified location of the targeted file. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSSavePanel/url
func (s_ SavePanel) URL() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID(), objc.RegisterName("URL"))
	return rv
}
