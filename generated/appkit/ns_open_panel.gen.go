// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSOpenPanel */


/* debug [class_header]: Header for NSOpenPanel */
// The class instance for the [OpenPanel] class.
var (
	OpenPanelClass     _OpenPanelClass
	OpenPanelClassOnce sync.Once
)

func getOpenPanelClass() _OpenPanelClass {
	OpenPanelClassOnce.Do(func() {
		OpenPanelClass = _OpenPanelClass{objc.GetClass("NSOpenPanel")}
	})
	return OpenPanelClass
}

type _OpenPanelClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for OpenPanel */
// An interface definition for the [OpenPanel] class.
type IOpenPanel interface {
	ISavePanel
	
/* debug [class_interface_properties]: Properties for OpenPanel */
	// properties:
	AllowsMultipleSelection() bool
	SetAllowsMultipleSelection(value bool)
	CanChooseDirectories() bool
	SetCanChooseDirectories(value bool)
	CanChooseFiles() bool
	SetCanChooseFiles(value bool)
	AccessoryViewDisclosed() bool
	SetAccessoryViewDisclosed(value bool)
	ResolvesAliases() bool
	SetResolvesAliases(value bool)
	URLs() []foundation.URL
	CanDownloadUbiquitousContents() bool
	SetCanDownloadUbiquitousContents(value bool)
	CanResolveUbiquitousConflicts() bool
	SetCanResolveUbiquitousConflicts(value bool)
	IsAccessoryViewDisclosed() bool
	SetIsAccessoryViewDisclosed(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for OpenPanel */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for OpenPanel */
// Alloc allocates a new instance without initialization.
func (oc _OpenPanelClass) Alloc() OpenPanel {
	rv := objc.Send[OpenPanel](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (oc _OpenPanelClass) New() OpenPanel {
	rv := objc.Send[OpenPanel](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ OpenPanel) Init() OpenPanel {
	rv := objc.Send[OpenPanel](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ OpenPanel) Autorelease() OpenPanel {
	rv := objc.Send[OpenPanel](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewOpenPanel creates a new OpenPanel instance.
func NewOpenPanel() OpenPanel {
	return getOpenPanelClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for OpenPanel */
// A panel that prompts the user to select a file to open.
//
// Apps use the Open panel as a convenient way to query the user for the name of a file to open. In macOS 10.15 and later, the system always draws Open panels in a separate process, regardless of whether the app is sandboxed. When the user chooses a file to open, macOS adds that file to the app’s sandbox. Prior to macOS 10.15, the system drew the panels in a separate process only for sandboxed apps.


// A panel that prompts the user to select a file to open.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenPanel
type OpenPanel struct {
	SavePanel
}

// OpenPanelFrom constructs a [OpenPanel] from an unsafe.Pointer.
//
// A panel that prompts the user to select a file to open.
func OpenPanelFrom(ptr unsafe.Pointer) OpenPanel {
	return OpenPanel{
		SavePanel: SavePanelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for OpenPanel *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for OpenPanel */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for OpenPanel */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for OpenPanel */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for OpenPanel */

// A Boolean that indicates whether the user may select multiple files and directories.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenPanel/allowsMultipleSelection
func (o_ OpenPanel) AllowsMultipleSelection() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("allowsMultipleSelection"))
	return rv
}/* debug [instance_properties/getter]: allowsMultipleSelection */


// A Boolean that indicates whether the user may select multiple files and directories.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenPanel/allowsMultipleSelection
func (o_ OpenPanel) SetAllowsMultipleSelection(value bool) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAllowsMultipleSelection:"), value)
}/* debug [instance_properties/setter]: allowsMultipleSelection */


// A Boolean that indicates whether the user can choose directories in the panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenPanel/canChooseDirectories
func (o_ OpenPanel) CanChooseDirectories() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("canChooseDirectories"))
	return rv
}/* debug [instance_properties/getter]: canChooseDirectories */


// A Boolean that indicates whether the user can choose directories in the panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenPanel/canChooseDirectories
func (o_ OpenPanel) SetCanChooseDirectories(value bool) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setCanChooseDirectories:"), value)
}/* debug [instance_properties/setter]: canChooseDirectories */


// A Boolean that indicates whether the user can choose files in the panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenPanel/canChooseFiles
func (o_ OpenPanel) CanChooseFiles() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("canChooseFiles"))
	return rv
}/* debug [instance_properties/getter]: canChooseFiles */


// A Boolean that indicates whether the user can choose files in the panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenPanel/canChooseFiles
func (o_ OpenPanel) SetCanChooseFiles(value bool) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setCanChooseFiles:"), value)
}/* debug [instance_properties/setter]: canChooseFiles */


// A Boolean value that indicates whether the panel’s accessory view is visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenPanel/isAccessoryViewDisclosed
func (o_ OpenPanel) AccessoryViewDisclosed() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("accessoryViewDisclosed"))
	return rv
}/* debug [instance_properties/getter]: accessoryViewDisclosed */


// A Boolean value that indicates whether the panel’s accessory view is visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenPanel/isAccessoryViewDisclosed
func (o_ OpenPanel) SetAccessoryViewDisclosed(value bool) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAccessoryViewDisclosed:"), value)
}/* debug [instance_properties/setter]: accessoryViewDisclosed */


// A Boolean that indicates whether the panel resolves aliases.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenPanel/resolvesAliases
func (o_ OpenPanel) ResolvesAliases() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("resolvesAliases"))
	return rv
}/* debug [instance_properties/getter]: resolvesAliases */


// A Boolean that indicates whether the panel resolves aliases.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenPanel/resolvesAliases
func (o_ OpenPanel) SetResolvesAliases(value bool) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setResolvesAliases:"), value)
}/* debug [instance_properties/setter]: resolvesAliases */


// An array of URLs, each of which contains the fully specified location of a selected file or directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenPanel/urls
func (o_ OpenPanel) URLs() []foundation.URL {
	rv := objc.Send[[]foundation.URL](o_.ID, objc.Sel("URLs"))
	return rv
}/* debug [instance_properties/getter]: URLs */


// A Boolean value that indicates how the panel responds to iCloud documents that aren’t fully downloaded locally.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsopenpanel/candownloadubiquitouscontents
func (o_ OpenPanel) CanDownloadUbiquitousContents() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("canDownloadUbiquitousContents"))
	return rv
}/* debug [instance_properties/getter]: canDownloadUbiquitousContents */


// A Boolean value that indicates how the panel responds to iCloud documents that aren’t fully downloaded locally.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsopenpanel/candownloadubiquitouscontents
func (o_ OpenPanel) SetCanDownloadUbiquitousContents(value bool) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setCanDownloadUbiquitousContents:"), value)
}/* debug [instance_properties/setter]: canDownloadUbiquitousContents */


// A Boolean value that indicates how the panel responds to iCloud documents that have conflicting versions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsopenpanel/canresolveubiquitousconflicts
func (o_ OpenPanel) CanResolveUbiquitousConflicts() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("canResolveUbiquitousConflicts"))
	return rv
}/* debug [instance_properties/getter]: canResolveUbiquitousConflicts */


// A Boolean value that indicates how the panel responds to iCloud documents that have conflicting versions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsopenpanel/canresolveubiquitousconflicts
func (o_ OpenPanel) SetCanResolveUbiquitousConflicts(value bool) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setCanResolveUbiquitousConflicts:"), value)
}/* debug [instance_properties/setter]: canResolveUbiquitousConflicts */


// A Boolean value that indicates whether the panel’s accessory view is visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsopenpanel/isaccessoryviewdisclosed
func (o_ OpenPanel) IsAccessoryViewDisclosed() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("isAccessoryViewDisclosed"))
	return rv
}/* debug [instance_properties/getter]: isAccessoryViewDisclosed */


// A Boolean value that indicates whether the panel’s accessory view is visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsopenpanel/isaccessoryviewdisclosed
func (o_ OpenPanel) SetIsAccessoryViewDisclosed(value bool) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setIsAccessoryViewDisclosed:"), value)
}/* debug [instance_properties/setter]: isAccessoryViewDisclosed */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSOpenPanel */



