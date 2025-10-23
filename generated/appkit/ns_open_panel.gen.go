// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [OpenPanel] class.
type IOpenPanel interface {
	ISavePanel
	// properties:
	CanChooseFiles() bool /* primitive/slice/pointer. */
	SetCanChooseFiles(value bool /* primitive/slice/pointer. */)
	AllowsMultipleSelection() bool /* primitive/slice/pointer. */
	SetAllowsMultipleSelection(value bool /* primitive/slice/pointer. */)
	CanChooseDirectories() bool /* primitive/slice/pointer. */
	SetCanChooseDirectories(value bool /* primitive/slice/pointer. */)
	CanDownloadUbiquitousContents() bool /* primitive/slice/pointer. */
	SetCanDownloadUbiquitousContents(value bool /* primitive/slice/pointer. */)
	CanResolveUbiquitousConflicts() bool /* primitive/slice/pointer. */
	SetCanResolveUbiquitousConflicts(value bool /* primitive/slice/pointer. */)
	IsAccessoryViewDisclosed() bool /* primitive/slice/pointer. */
	SetIsAccessoryViewDisclosed(value bool /* primitive/slice/pointer. */)
	ResolvesAliases() bool /* primitive/slice/pointer. */
	SetResolvesAliases(value bool /* primitive/slice/pointer. */)
	Urls() foundation.objc.IObject /* cross-framework: URL */
	SetUrls(value foundation.objc.IObject /* cross-framework: URL */)
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (oc _OpenPanelClass) Alloc() OpenPanel {
	rv := objc.Send[OpenPanel](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// A Boolean that indicates whether the user can choose files in the panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenPanel/canChooseFiles
func (o_ OpenPanel) CanChooseFiles() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](o_.ID, objc.Sel("canChooseFiles"))
	return rv
}


// A Boolean that indicates whether the user can choose files in the panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenPanel/canChooseFiles
func (o_ OpenPanel) SetCanChooseFiles(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setCanChooseFiles:"), value)
}


// A Boolean that indicates whether the user may select multiple files and directories.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsopenpanel/allowsmultipleselection
func (o_ OpenPanel) AllowsMultipleSelection() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](o_.ID, objc.Sel("allowsMultipleSelection"))
	return rv
}


// A Boolean that indicates whether the user may select multiple files and directories.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsopenpanel/allowsmultipleselection
func (o_ OpenPanel) SetAllowsMultipleSelection(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAllowsMultipleSelection:"), value)
}


// A Boolean that indicates whether the user can choose directories in the panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsopenpanel/canchoosedirectories
func (o_ OpenPanel) CanChooseDirectories() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](o_.ID, objc.Sel("canChooseDirectories"))
	return rv
}


// A Boolean that indicates whether the user can choose directories in the panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsopenpanel/canchoosedirectories
func (o_ OpenPanel) SetCanChooseDirectories(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setCanChooseDirectories:"), value)
}


// A Boolean value that indicates how the panel responds to iCloud documents that aren’t fully downloaded locally.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsopenpanel/candownloadubiquitouscontents
func (o_ OpenPanel) CanDownloadUbiquitousContents() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](o_.ID, objc.Sel("canDownloadUbiquitousContents"))
	return rv
}


// A Boolean value that indicates how the panel responds to iCloud documents that aren’t fully downloaded locally.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsopenpanel/candownloadubiquitouscontents
func (o_ OpenPanel) SetCanDownloadUbiquitousContents(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setCanDownloadUbiquitousContents:"), value)
}


// A Boolean value that indicates how the panel responds to iCloud documents that have conflicting versions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsopenpanel/canresolveubiquitousconflicts
func (o_ OpenPanel) CanResolveUbiquitousConflicts() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](o_.ID, objc.Sel("canResolveUbiquitousConflicts"))
	return rv
}


// A Boolean value that indicates how the panel responds to iCloud documents that have conflicting versions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsopenpanel/canresolveubiquitousconflicts
func (o_ OpenPanel) SetCanResolveUbiquitousConflicts(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setCanResolveUbiquitousConflicts:"), value)
}


// A Boolean value that indicates whether the panel’s accessory view is visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsopenpanel/isaccessoryviewdisclosed
func (o_ OpenPanel) IsAccessoryViewDisclosed() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](o_.ID, objc.Sel("isAccessoryViewDisclosed"))
	return rv
}


// A Boolean value that indicates whether the panel’s accessory view is visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsopenpanel/isaccessoryviewdisclosed
func (o_ OpenPanel) SetIsAccessoryViewDisclosed(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setIsAccessoryViewDisclosed:"), value)
}


// A Boolean that indicates whether the panel resolves aliases.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsopenpanel/resolvesaliases
func (o_ OpenPanel) ResolvesAliases() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](o_.ID, objc.Sel("resolvesAliases"))
	return rv
}


// A Boolean that indicates whether the panel resolves aliases.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsopenpanel/resolvesaliases
func (o_ OpenPanel) SetResolvesAliases(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setResolvesAliases:"), value)
}


// An array of URLs, each of which contains the fully specified location of a selected file or directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsopenpanel/urls
func (o_ OpenPanel) Urls() foundation.objc.IObject /* cross-framework: URL */ {
	rv := objc.Send[foundation.URL](o_.ID, objc.Sel("urls"))
	return rv
}


// An array of URLs, each of which contains the fully specified location of a selected file or directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsopenpanel/urls
func (o_ OpenPanel) SetUrls(value foundation.objc.IObject /* cross-framework: URL */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setUrls:"), value)
}



