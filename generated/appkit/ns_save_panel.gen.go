// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/uniformtypeidentifiers"
)

/* debug [class.gen.go]: Generating class NSSavePanel */


/* debug [class_header]: Header for NSSavePanel */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SavePanel */
// An interface definition for the [SavePanel] class.
type ISavePanel interface {
	IPanel
	
/* debug [class_interface_properties]: Properties for SavePanel */
	// properties:
	AccessoryView() IView
	SetAccessoryView(value IView)
	AllowedContentTypes() []uniformtypeidentifiers.UTType
	SetAllowedContentTypes(value []uniformtypeidentifiers.UTType)
	AllowedFileTypes() []string
	SetAllowedFileTypes(value []string)
	AllowsOtherFileTypes() bool
	SetAllowsOtherFileTypes(value bool)
	CanCreateDirectories() bool
	SetCanCreateDirectories(value bool)
	CanSelectHiddenExtension() bool
	SetCanSelectHiddenExtension(value bool)
	DirectoryURL() objc.IObject /* cross-framework: NSURL */
	SetDirectoryURL(value objc.IObject /* cross-framework: NSURL */)
	Expanded() bool
	ExtensionHidden() bool
	SetExtensionHidden(value bool)
	Message() objc.IObject /* cross-framework: NSString */
	SetMessage(value objc.IObject /* cross-framework: NSString */)
	NameFieldLabel() objc.IObject /* cross-framework: NSString */
	SetNameFieldLabel(value objc.IObject /* cross-framework: NSString */)
	NameFieldStringValue() objc.IObject /* cross-framework: NSString */
	SetNameFieldStringValue(value objc.IObject /* cross-framework: NSString */)
	Prompt() objc.IObject /* cross-framework: NSString */
	SetPrompt(value objc.IObject /* cross-framework: NSString */)
	ShowsHiddenFiles() bool
	SetShowsHiddenFiles(value bool)
	ShowsTagField() bool
	SetShowsTagField(value bool)
	TagNames() []string
	SetTagNames(value []string)
	Title() objc.IObject /* cross-framework: NSString */
	SetTitle(value objc.IObject /* cross-framework: NSString */)
	TreatsFilePackagesAsDirectories() bool
	SetTreatsFilePackagesAsDirectories(value bool)
	URL() objc.IObject /* cross-framework: NSURL */
	CurrentContentType() uniformtypeidentifiers.UTType
	SetCurrentContentType(value uniformtypeidentifiers.UTType)
	Delegate() objc.IObject /* cross-framework: OpenSavePanelDelegate */
	SetDelegate(value objc.IObject /* cross-framework: OpenSavePanelDelegate */)
	Identifier() UserInterfaceItemIdentifier /* typedef */
	SetIdentifier(value UserInterfaceItemIdentifier /* typedef */)
	IsExpanded() bool
	SetIsExpanded(value bool)
	IsExtensionHidden() bool
	SetIsExtensionHidden(value bool)
	ShowsContentTypes() bool
	SetShowsContentTypes(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SavePanel */
	// methods:
	BeginWithCompletionHandler(handler unsafe.Pointer)
	BeginSheetModalForWindowCompletionHandler(window IWindow, handler unsafe.Pointer)
	RunModal() ModalResponse /* typedef */
	ValidateVisibleColumns()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SavePanel */
// Alloc allocates a new instance without initialization.
func (sc _SavePanelClass) Alloc() SavePanel {
	rv := objc.Send[SavePanel](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SavePanel */
// A panel that prompts the user for information about where to save a file.
//
// The Save panel provides an interface for specifying the location to save a file and the name of that file. You present this panel when the user attempts to save a new document, or when the user saves a copy of an existing document to a new location. The panel includes UI for browsing the file system, selecting a directory, and specifying the new name for the file. You can also add custom UI for your app using an accessory view. An object reports user interactions to its associated object, which must adopt the protocol. Use your delegate object to validate the user’s selection and respond to user interactions with the panel. In macOS 10.15, the system always displays the Save dialog in a separate process, regardless of whether the app is sandboxed. When the user saves the document, macOS adds the saved file to the app’s sandbox (if necessary) so that the app can write to the file. Prior to macOS 10.15, the system used a separate process only for sandboxed apps.


// A panel that prompts the user for information about where to save a file.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SavePanel *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SavePanel */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SavePanel */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SavePanel */

// Presents the panel as a modeless window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/begin(completionHandler:)
func (s_ SavePanel) BeginWithCompletionHandler(handler unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("beginWithCompletionHandler:"), handler)
}/* debug [instance_methods/method]: BeginWithCompletionHandler */


// Presents the panel as a sheet modal to the specified window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/beginSheetModal(for:completionHandler:)
func (s_ SavePanel) BeginSheetModalForWindowCompletionHandler(window IWindow, handler unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("beginSheetModalForWindow:completionHandler:"), window, handler)
}/* debug [instance_methods/method]: BeginSheetModalForWindowCompletionHandler */


// Displays the panel and begins its event loop with the current working (or last-selected) directory as the default starting point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/runModal()
func (s_ SavePanel) RunModal() ModalResponse /* typedef */ {
	rv := objc.Send[int](s_.ID, objc.Sel("runModal"))
	return rv
}/* debug [instance_methods/method]: RunModal */


// Validates and reloads the browser columns visible in the panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/validateVisibleColumns()
func (s_ SavePanel) ValidateVisibleColumns() {
	objc.Send[objc.ID](s_.ID, objc.Sel("validateVisibleColumns"))
}/* debug [instance_methods/method]: ValidateVisibleColumns */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SavePanel */

// The custom accessory view for the current app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/accessoryView
func (s_ SavePanel) AccessoryView() IView {
	rv := objc.Send[View](s_.ID, objc.Sel("accessoryView"))
	return rv
}/* debug [instance_properties/getter]: accessoryView */


// The custom accessory view for the current app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/accessoryView
func (s_ SavePanel) SetAccessoryView(value IView) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAccessoryView:"), value)
}/* debug [instance_properties/setter]: accessoryView */


// An array of types that specify the files types to which you can save.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/allowedContentTypes
func (s_ SavePanel) AllowedContentTypes() []uniformtypeidentifiers.UTType {
	rv := objc.Send[[]uniformtypeidentifiers.UTType](s_.ID, objc.Sel("allowedContentTypes"))
	return rv
}/* debug [instance_properties/getter]: allowedContentTypes */


// An array of types that specify the files types to which you can save.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/allowedContentTypes
func (s_ SavePanel) SetAllowedContentTypes(value []uniformtypeidentifiers.UTType) {
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
}/* debug [instance_properties/setter]: allowedContentTypes */


// An array of filename extensions or UTIs that represent the allowed file types for the panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/allowedFileTypes
func (s_ SavePanel) AllowedFileTypes() []string {
	rv := objc.Send[[]string](s_.ID, objc.Sel("allowedFileTypes"))
	return rv
}/* debug [instance_properties/getter]: allowedFileTypes */


// An array of filename extensions or UTIs that represent the allowed file types for the panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/allowedFileTypes
func (s_ SavePanel) SetAllowedFileTypes(value []string) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](s_.ID, objc.Sel("setAllowedFileTypes:"), nsArray)
}/* debug [instance_properties/setter]: allowedFileTypes */


// A Boolean value that indicates whether the panel allows the user to save files with a filename extension that’s not in the list of allowed types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/allowsOtherFileTypes
func (s_ SavePanel) AllowsOtherFileTypes() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("allowsOtherFileTypes"))
	return rv
}/* debug [instance_properties/getter]: allowsOtherFileTypes */


// A Boolean value that indicates whether the panel allows the user to save files with a filename extension that’s not in the list of allowed types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/allowsOtherFileTypes
func (s_ SavePanel) SetAllowsOtherFileTypes(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAllowsOtherFileTypes:"), value)
}/* debug [instance_properties/setter]: allowsOtherFileTypes */


// A Boolean value that indicates whether the panel displays UI for creating directories.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/canCreateDirectories
func (s_ SavePanel) CanCreateDirectories() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("canCreateDirectories"))
	return rv
}/* debug [instance_properties/getter]: canCreateDirectories */


// A Boolean value that indicates whether the panel displays UI for creating directories.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/canCreateDirectories
func (s_ SavePanel) SetCanCreateDirectories(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCanCreateDirectories:"), value)
}/* debug [instance_properties/setter]: canCreateDirectories */


// A Boolean value that indicates whether the panel displays UI for hiding or showing filename extensions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/canSelectHiddenExtension
func (s_ SavePanel) CanSelectHiddenExtension() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("canSelectHiddenExtension"))
	return rv
}/* debug [instance_properties/getter]: canSelectHiddenExtension */


// A Boolean value that indicates whether the panel displays UI for hiding or showing filename extensions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/canSelectHiddenExtension
func (s_ SavePanel) SetCanSelectHiddenExtension(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCanSelectHiddenExtension:"), value)
}/* debug [instance_properties/setter]: canSelectHiddenExtension */


// The current directory shown in the panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/directoryURL
func (s_ SavePanel) DirectoryURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](s_.ID, objc.Sel("directoryURL"))
	return rv
}/* debug [instance_properties/getter]: directoryURL */


// The current directory shown in the panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/directoryURL
func (s_ SavePanel) SetDirectoryURL(value objc.IObject /* cross-framework: NSURL */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDirectoryURL:"), value)
}/* debug [instance_properties/setter]: directoryURL */


// A Boolean value that indicates whether whether the panel is expanded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/isExpanded
func (s_ SavePanel) Expanded() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("expanded"))
	return rv
}/* debug [instance_properties/getter]: expanded */


// A Boolean value that indicates whether to display filename extensions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/isExtensionHidden
func (s_ SavePanel) ExtensionHidden() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("extensionHidden"))
	return rv
}/* debug [instance_properties/getter]: extensionHidden */


// A Boolean value that indicates whether to display filename extensions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/isExtensionHidden
func (s_ SavePanel) SetExtensionHidden(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setExtensionHidden:"), value)
}/* debug [instance_properties/setter]: extensionHidden */


// The message text displayed in the panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/message
func (s_ SavePanel) Message() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("message"))
	return rv
}/* debug [instance_properties/getter]: message */


// The message text displayed in the panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/message
func (s_ SavePanel) SetMessage(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMessage:"), value)
}/* debug [instance_properties/setter]: message */


// The label text displayed in front of the filename text field.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/nameFieldLabel
func (s_ SavePanel) NameFieldLabel() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("nameFieldLabel"))
	return rv
}/* debug [instance_properties/getter]: nameFieldLabel */


// The label text displayed in front of the filename text field.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/nameFieldLabel
func (s_ SavePanel) SetNameFieldLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setNameFieldLabel:"), value)
}/* debug [instance_properties/setter]: nameFieldLabel */


// The user-editable filename currently shown in the name field.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/nameFieldStringValue
func (s_ SavePanel) NameFieldStringValue() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("nameFieldStringValue"))
	return rv
}/* debug [instance_properties/getter]: nameFieldStringValue */


// The user-editable filename currently shown in the name field.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/nameFieldStringValue
func (s_ SavePanel) SetNameFieldStringValue(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setNameFieldStringValue:"), value)
}/* debug [instance_properties/setter]: nameFieldStringValue */


// The text to display in the default button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/prompt
func (s_ SavePanel) Prompt() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("prompt"))
	return rv
}/* debug [instance_properties/getter]: prompt */


// The text to display in the default button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/prompt
func (s_ SavePanel) SetPrompt(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPrompt:"), value)
}/* debug [instance_properties/setter]: prompt */


// A Boolean value that indicates whether the panel displays files that are normally hidden from the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/showsHiddenFiles
func (s_ SavePanel) ShowsHiddenFiles() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("showsHiddenFiles"))
	return rv
}/* debug [instance_properties/getter]: showsHiddenFiles */


// A Boolean value that indicates whether the panel displays files that are normally hidden from the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/showsHiddenFiles
func (s_ SavePanel) SetShowsHiddenFiles(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setShowsHiddenFiles:"), value)
}/* debug [instance_properties/setter]: showsHiddenFiles */


// A Boolean value that indicates whether the panel displays the Tags field.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/showsTagField
func (s_ SavePanel) ShowsTagField() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("showsTagField"))
	return rv
}/* debug [instance_properties/getter]: showsTagField */


// A Boolean value that indicates whether the panel displays the Tags field.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/showsTagField
func (s_ SavePanel) SetShowsTagField(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setShowsTagField:"), value)
}/* debug [instance_properties/setter]: showsTagField */


// The tag names that you want to include on a saved file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/tagNames
func (s_ SavePanel) TagNames() []string {
	rv := objc.Send[[]string](s_.ID, objc.Sel("tagNames"))
	return rv
}/* debug [instance_properties/getter]: tagNames */


// The tag names that you want to include on a saved file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/tagNames
func (s_ SavePanel) SetTagNames(value []string) {
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
}/* debug [instance_properties/setter]: tagNames */


// The title of the panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/title
func (s_ SavePanel) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("title"))
	return rv
}/* debug [instance_properties/getter]: title */


// The title of the panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/title
func (s_ SavePanel) SetTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTitle:"), value)
}/* debug [instance_properties/setter]: title */


// A Boolean value that indicates whether the panel displays file packages as directories.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/treatsFilePackagesAsDirectories
func (s_ SavePanel) TreatsFilePackagesAsDirectories() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("treatsFilePackagesAsDirectories"))
	return rv
}/* debug [instance_properties/getter]: treatsFilePackagesAsDirectories */


// A Boolean value that indicates whether the panel displays file packages as directories.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/treatsFilePackagesAsDirectories
func (s_ SavePanel) SetTreatsFilePackagesAsDirectories(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTreatsFilePackagesAsDirectories:"), value)
}/* debug [instance_properties/setter]: treatsFilePackagesAsDirectories */


// A URL that contains the fully specified location of the targeted file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSavePanel/url
func (s_ SavePanel) URL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](s_.ID, objc.Sel("URL"))
	return rv
}/* debug [instance_properties/getter]: URL */


// :The current type. If set to
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssavepanel/currentcontenttype
func (s_ SavePanel) CurrentContentType() uniformtypeidentifiers.UTType {
	rv := objc.Send[uniformtypeidentifiers.UTType](s_.ID, objc.Sel("currentContentType"))
	return rv
}/* debug [instance_properties/getter]: currentContentType */


// :The current type. If set to
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssavepanel/currentcontenttype
func (s_ SavePanel) SetCurrentContentType(value uniformtypeidentifiers.UTType) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCurrentContentType:"), value)
}/* debug [instance_properties/setter]: currentContentType */


// A custom object you use to manage interactions with an open or save panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssavepanel/delegate
func (s_ SavePanel) Delegate() objc.IObject /* cross-framework: OpenSavePanelDelegate */ {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// A custom object you use to manage interactions with an open or save panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssavepanel/delegate
func (s_ SavePanel) SetDelegate(value objc.IObject /* cross-framework: OpenSavePanelDelegate */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssavepanel/identifier
func (s_ SavePanel) Identifier() UserInterfaceItemIdentifier /* typedef */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("identifier"))
	return rv
}/* debug [instance_properties/getter]: identifier */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssavepanel/identifier
func (s_ SavePanel) SetIdentifier(value UserInterfaceItemIdentifier /* typedef */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIdentifier:"), value)
}/* debug [instance_properties/setter]: identifier */


// A Boolean value that indicates whether whether the panel is expanded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssavepanel/isexpanded
func (s_ SavePanel) IsExpanded() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isExpanded"))
	return rv
}/* debug [instance_properties/getter]: isExpanded */


// A Boolean value that indicates whether whether the panel is expanded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssavepanel/isexpanded
func (s_ SavePanel) SetIsExpanded(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsExpanded:"), value)
}/* debug [instance_properties/setter]: isExpanded */


// A Boolean value that indicates whether to display filename extensions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssavepanel/isextensionhidden
func (s_ SavePanel) IsExtensionHidden() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isExtensionHidden"))
	return rv
}/* debug [instance_properties/getter]: isExtensionHidden */


// A Boolean value that indicates whether to display filename extensions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssavepanel/isextensionhidden
func (s_ SavePanel) SetIsExtensionHidden(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsExtensionHidden:"), value)
}/* debug [instance_properties/setter]: isExtensionHidden */


// : Whether or not to show a control for selecting the type of the saved file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssavepanel/showscontenttypes
func (s_ SavePanel) ShowsContentTypes() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("showsContentTypes"))
	return rv
}/* debug [instance_properties/getter]: showsContentTypes */


// : Whether or not to show a control for selecting the type of the saved file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssavepanel/showscontenttypes
func (s_ SavePanel) SetShowsContentTypes(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setShowsContentTypes:"), value)
}/* debug [instance_properties/setter]: showsContentTypes */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSSavePanel */



