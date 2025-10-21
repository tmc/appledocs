// Code generated from Apple documentation for FinderSync. DO NOT EDIT.

package findersync

// Enum types and constants
// FIMenuKind - The different kinds of custom menus that the Finder Sync extension can
//
// [Full Topic]: https://developer.apple.com/documentation/FinderSync/FIMenuKind
type FIMenuKind uint

const (
	// FIMenuKindContextualMenuForContainer - A shortcut menu created when the user control-clicks on the Finder   window’s background.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FinderSync/FIMenuKind/contextualMenuForContainer
	FIMenuKindContextualMenuForContainer FIMenuKind = 0
	// FIMenuKindContextualMenuForItems - A shortcut menu created when the user control-clicks on an item or a   group of selected items inside the Finder window.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FinderSync/FIMenuKind/contextualMenuForItems
	FIMenuKindContextualMenuForItems FIMenuKind = 0
	// FIMenuKindContextualMenuForSidebar - A shortcut menu created when the user control-clicks on an item in the   sidebar.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FinderSync/FIMenuKind/contextualMenuForSidebar
	FIMenuKindContextualMenuForSidebar FIMenuKind = 0
	// FIMenuKindToolbarItemMenu - A menu created when the user clicks on the extension’s toolbar button.
	//
	// [Full Topic]: https://developer.apple.com/documentation/FinderSync/FIMenuKind/toolbarItemMenu
	FIMenuKindToolbarItemMenu FIMenuKind = 0
)


