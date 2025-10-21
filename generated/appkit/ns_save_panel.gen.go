// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/uniformtypeidentifiers"
)

// The class instance for the [SavePanel] class.
var (
	SavePanelClass     _SavePanelClass
	SavePanelClassOnce sync.Once
)

func getSavePanelClass() _SavePanelClass {
	SavePanelClassOnce.Do(func() {
		SavePanelClass = _SavePanelClass{objc.GetClass("NSSavePanel")}
	})
	return SavePanelClass
}

type _SavePanelClass struct {
	class objc.Class
}

// An interface definition for the [SavePanel] class.
type ISavePanel interface {
	IPanel
	BeginWithCompletionHandler(handler unsafe.Pointer)
	BeginSheetModalForWindowCompletionHandler(window unsafe.Pointer, handler unsafe.Pointer)
	Cancel(sender objc.ID)
	Ok(sender objc.ID)
	RunModal() unsafe.Pointer
	ValidateVisibleColumns()
}

// A panel that prompts the user for information about where to save a file.
//
// The Save panel provides an interface for specifying the location to save a file and the name of that file. You present this panel when the user attempts to save a new document, or when the user saves a copy of an existing document to a new location. The panel includes UI for browsing the file system, selecting a directory, and specifying the new name for the file. You can also add custom UI for your app using an accessory view. An object reports user interactions to its associated object, which must adopt the protocol. Use your delegate object to validate the user’s selection and respond to user interactions with the panel. In macOS 10.15, the system always displays the Save dialog in a separate process, regardless of whether the app is sandboxed. When the user saves the document, macOS adds the saved file to the app’s sandbox (if necessary) so that the app can write to the file. Prior to macOS 10.15, the system used a separate process only for sandboxed apps.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel
type SavePanel struct {
	Panel
}

// SavePanelFrom constructs a [SavePanel] from an unsafe.Pointer.
//
// A panel that prompts the user for information about where to save a file.
func SavePanelFrom(ptr unsafe.Pointer) SavePanel {
	return SavePanel{
		Panel: PanelFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (sc _SavePanelClass) Alloc() SavePanel {
	rv := objc.Send[SavePanel](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SavePanelClass) New() SavePanel {
	rv := objc.Send[SavePanel](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SavePanel) Init() SavePanel {
	rv := objc.Send[SavePanel](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SavePanel) Autorelease() SavePanel {
	rv := objc.Send[SavePanel](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSavePanel creates a new SavePanel instance.
func NewSavePanel() SavePanel {
	return getSavePanelClass().New()
}

// Creates a new Save panel and initializes it with default information.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/savePanel
func (sc _SavePanelClass) SavePanel() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("savePanel"))
	return rv
}

// Presents the panel as a modeless window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/begin(completionHandler:)
func (s_ SavePanel) BeginWithCompletionHandler(handler unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("beginWithCompletionHandler:"), handler)
}

// Presents the panel as a sheet modal to the specified window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/beginSheetModal(for:completionHandler:)
func (s_ SavePanel) BeginSheetModalForWindowCompletionHandler(window unsafe.Pointer, handler unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("beginSheetModalForWindow:completionHandler:"), window, handler)
}

// The action method that the panel calls when the user clicks the Cancel button.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/cancel(_:)
func (s_ SavePanel) Cancel(sender objc.ID) {
	objc.Send[objc.ID](s_.ID, objc.Sel("cancel:"), sender)
}

// The action method that the panel calls when the user clicks the OK button.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/ok(_:)
func (s_ SavePanel) Ok(sender objc.ID) {
	objc.Send[objc.ID](s_.ID, objc.Sel("ok:"), sender)
}

// Displays the panel and begins its event loop with the current working (or last-selected) directory as the default starting point.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/runModal()
func (s_ SavePanel) RunModal() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("runModal"))
	return rv
}

// Validates and reloads the browser columns visible in the panel.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/validateVisibleColumns()
func (s_ SavePanel) ValidateVisibleColumns() {
	objc.Send[objc.ID](s_.ID, objc.Sel("validateVisibleColumns"))
}

// The custom accessory view for the current app.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/accessoryView
func (s_ SavePanel) AccessoryView() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("accessoryView"))
	return rv
}

// SetAccessoryView sets the value of the accessoryView property.
// The custom accessory view for the current app.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/accessoryView
func (s_ SavePanel) SetAccessoryView(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAccessoryView:"), value)
}

// An array of types that specify the files types to which you can save.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/allowedContentTypes
func (s_ SavePanel) AllowedContentTypes() []uniformtypeidentifiers.UTType {
	rv := objc.Send[[]uniformtypeidentifiers.UTType](s_.ID, objc.Sel("allowedContentTypes"))
	return rv
}

// SetAllowedContentTypes sets the value of the allowedContentTypes property.
// An array of types that specify the files types to which you can save.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/allowedContentTypes
func (s_ SavePanel) SetAllowedContentTypes(value []uniformtypeidentifiers.UTType) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](s_.ID, objc.Sel("setAllowedContentTypes:"), nsArray)
}

// A Boolean value that indicates whether the panel allows the user to save files with a filename extension that’s not in the list of allowed types.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/allowsOtherFileTypes
func (s_ SavePanel) AllowsOtherFileTypes() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("allowsOtherFileTypes"))
	return rv
}

// SetAllowsOtherFileTypes sets the value of the allowsOtherFileTypes property.
// A Boolean value that indicates whether the panel allows the user to save files with a filename extension that’s not in the list of allowed types.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/allowsOtherFileTypes
func (s_ SavePanel) SetAllowsOtherFileTypes(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAllowsOtherFileTypes:"), value)
}

// A Boolean value that indicates whether the panel displays UI for creating directories.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/canCreateDirectories
func (s_ SavePanel) CanCreateDirectories() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("canCreateDirectories"))
	return rv
}

// SetCanCreateDirectories sets the value of the canCreateDirectories property.
// A Boolean value that indicates whether the panel displays UI for creating directories.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/canCreateDirectories
func (s_ SavePanel) SetCanCreateDirectories(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCanCreateDirectories:"), value)
}

// A Boolean value that indicates whether the panel displays UI for hiding or showing filename extensions.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/canSelectHiddenExtension
func (s_ SavePanel) CanSelectHiddenExtension() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("canSelectHiddenExtension"))
	return rv
}

// SetCanSelectHiddenExtension sets the value of the canSelectHiddenExtension property.
// A Boolean value that indicates whether the panel displays UI for hiding or showing filename extensions.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/canSelectHiddenExtension
func (s_ SavePanel) SetCanSelectHiddenExtension(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCanSelectHiddenExtension:"), value)
}

// :The current type. If set to , resets to the first allowed content type. Returns if is empty. : Not used.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/currentContentType
func (s_ SavePanel) CurrentContentType() uniformtypeidentifiers.UTType {
	rv := objc.Send[uniformtypeidentifiers.UTType](s_.ID, objc.Sel("currentContentType"))
	return rv
}

// SetCurrentContentType sets the value of the currentContentType property.
// :The current type. If set to , resets to the first allowed content type. Returns if is empty. : Not used.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/currentContentType
func (s_ SavePanel) SetCurrentContentType(value uniformtypeidentifiers.UTType) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCurrentContentType:"), value)
}

// A custom object you use to manage interactions with an open or save panel.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/delegate
func (s_ SavePanel) Delegate() objc.ID {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("delegate"))
	return rv
}

// SetDelegate sets the value of the delegate property.
// A custom object you use to manage interactions with an open or save panel.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/delegate
func (s_ SavePanel) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDelegate:"), value)
}

// The current directory shown in the panel.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/directoryURL
func (s_ SavePanel) DirectoryURL() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("directoryURL"))
	return rv
}

// SetDirectoryURL sets the value of the directoryURL property.
// The current directory shown in the panel.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/directoryURL
func (s_ SavePanel) SetDirectoryURL(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDirectoryURL:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/identifier
func (s_ SavePanel) Identifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("identifier"))
	return rv
}

// SetIdentifier sets the value of the identifier property.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/identifier
func (s_ SavePanel) SetIdentifier(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIdentifier:"), value)
}

// A Boolean value that indicates whether whether the panel is expanded.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/isExpanded
func (s_ SavePanel) Expanded() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("expanded"))
	return rv
}

// A Boolean value that indicates whether to display filename extensions.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/isExtensionHidden
func (s_ SavePanel) ExtensionHidden() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("extensionHidden"))
	return rv
}

// SetExtensionHidden sets the value of the extensionHidden property.
// A Boolean value that indicates whether to display filename extensions.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/isExtensionHidden
func (s_ SavePanel) SetExtensionHidden(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setExtensionHidden:"), value)
}

// The message text displayed in the panel.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/message
func (s_ SavePanel) Message() string {
	rv := objc.Send[string](s_.ID, objc.Sel("message"))
	return rv
}

// SetMessage sets the value of the message property.
// The message text displayed in the panel.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/message
func (s_ SavePanel) SetMessage(value string) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMessage:"), objc.String(value))
}

// The label text displayed in front of the filename text field.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/nameFieldLabel
func (s_ SavePanel) NameFieldLabel() string {
	rv := objc.Send[string](s_.ID, objc.Sel("nameFieldLabel"))
	return rv
}

// SetNameFieldLabel sets the value of the nameFieldLabel property.
// The label text displayed in front of the filename text field.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/nameFieldLabel
func (s_ SavePanel) SetNameFieldLabel(value string) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setNameFieldLabel:"), objc.String(value))
}

// The user-editable filename currently shown in the name field.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/nameFieldStringValue
func (s_ SavePanel) NameFieldStringValue() string {
	rv := objc.Send[string](s_.ID, objc.Sel("nameFieldStringValue"))
	return rv
}

// SetNameFieldStringValue sets the value of the nameFieldStringValue property.
// The user-editable filename currently shown in the name field.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/nameFieldStringValue
func (s_ SavePanel) SetNameFieldStringValue(value string) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setNameFieldStringValue:"), objc.String(value))
}

// The text to display in the default button.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/prompt
func (s_ SavePanel) Prompt() string {
	rv := objc.Send[string](s_.ID, objc.Sel("prompt"))
	return rv
}

// SetPrompt sets the value of the prompt property.
// The text to display in the default button.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/prompt
func (s_ SavePanel) SetPrompt(value string) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPrompt:"), objc.String(value))
}

// : Whether or not to show a control for selecting the type of the saved file. The control shows the types in . Default is . : Not used.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/showsContentTypes
func (s_ SavePanel) ShowsContentTypes() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("showsContentTypes"))
	return rv
}

// SetShowsContentTypes sets the value of the showsContentTypes property.
// : Whether or not to show a control for selecting the type of the saved file. The control shows the types in . Default is . : Not used.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/showsContentTypes
func (s_ SavePanel) SetShowsContentTypes(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setShowsContentTypes:"), value)
}

// A Boolean value that indicates whether the panel displays files that are normally hidden from the user.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/showsHiddenFiles
func (s_ SavePanel) ShowsHiddenFiles() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("showsHiddenFiles"))
	return rv
}

// SetShowsHiddenFiles sets the value of the showsHiddenFiles property.
// A Boolean value that indicates whether the panel displays files that are normally hidden from the user.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/showsHiddenFiles
func (s_ SavePanel) SetShowsHiddenFiles(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setShowsHiddenFiles:"), value)
}

// A Boolean value that indicates whether the panel displays the Tags field.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/showsTagField
func (s_ SavePanel) ShowsTagField() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("showsTagField"))
	return rv
}

// SetShowsTagField sets the value of the showsTagField property.
// A Boolean value that indicates whether the panel displays the Tags field.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/showsTagField
func (s_ SavePanel) SetShowsTagField(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setShowsTagField:"), value)
}

// The tag names that you want to include on a saved file.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/tagNames
func (s_ SavePanel) TagNames() []string {
	rv := objc.Send[[]string](s_.ID, objc.Sel("tagNames"))
	return rv
}

// SetTagNames sets the value of the tagNames property.
// The tag names that you want to include on a saved file.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/tagNames
func (s_ SavePanel) SetTagNames(value []string) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](s_.ID, objc.Sel("setTagNames:"), nsArray)
}

// The title of the panel.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/title
func (s_ SavePanel) Title() string {
	rv := objc.Send[string](s_.ID, objc.Sel("title"))
	return rv
}

// SetTitle sets the value of the title property.
// The title of the panel.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/title
func (s_ SavePanel) SetTitle(value string) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTitle:"), objc.String(value))
}

// A Boolean value that indicates whether the panel displays file packages as directories.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/treatsFilePackagesAsDirectories
func (s_ SavePanel) TreatsFilePackagesAsDirectories() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("treatsFilePackagesAsDirectories"))
	return rv
}

// SetTreatsFilePackagesAsDirectories sets the value of the treatsFilePackagesAsDirectories property.
// A Boolean value that indicates whether the panel displays file packages as directories.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/treatsFilePackagesAsDirectories
func (s_ SavePanel) SetTreatsFilePackagesAsDirectories(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTreatsFilePackagesAsDirectories:"), value)
}

// A URL that contains the fully specified location of the targeted file.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/url
func (s_ SavePanel) URL() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("URL"))
	return rv
}
