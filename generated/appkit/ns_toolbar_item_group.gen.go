// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ToolbarItemGroup] class.
var (
	ToolbarItemGroupClass     _ToolbarItemGroupClass
	ToolbarItemGroupClassOnce sync.Once
)

func getToolbarItemGroupClass() _ToolbarItemGroupClass {
	ToolbarItemGroupClassOnce.Do(func() {
		ToolbarItemGroupClass = _ToolbarItemGroupClass{objc.GetClass("NSToolbarItemGroup")}
	})
	return ToolbarItemGroupClass
}

type _ToolbarItemGroupClass struct {
	class objc.Class
}

// An interface definition for the [ToolbarItemGroup] class.
type IToolbarItemGroup interface {
	IToolbarItem
	// properties:
	ControlRepresentation() ToolbarItemGroupControlRepresentation
	SetControlRepresentation(value ToolbarItemGroupControlRepresentation)
	SelectedIndex() int
	SetSelectedIndex(value int)
	SelectionMode() ToolbarItemGroupSelectionMode
	SetSelectionMode(value ToolbarItemGroupSelectionMode)
	Subitems() []ToolbarItem
	SetSubitems(value []ToolbarItem)
	// methods:
	IsSelectedAtIndex(index int) bool
	SetSelectedAtIndex(selected bool, index int)
}

// A group of subitems in a toolbar item.
//
// An represents a collection set of subitems in a toolbar that the system displays based on available space and settings that you specify. The system uses the views and labels of the subitems, but the parent’s attributes take precedence. This differs from other objects because they’re attached — the user drags them together as a single item rather than separately. If a subitem of the group has an action set on it, the group uses that action instead of its own when the user clicks or taps on that item. The system prefers the subitem’s action if it exists, otherwise it uses the group’s action. To configure an instance of , you first create the individual toolbar subitems: Then, you put them in a grouped item: In this configuration, you get two grouped items, and two labels. If you set a label on the parent item, you get two grouped items with one shared label: If instead you set a view on the parent item, you get two labels with one shared view:


// A group of subitems in a toolbar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItemGroup
type ToolbarItemGroup struct {
	ToolbarItem
}

// ToolbarItemGroupFrom constructs a [ToolbarItemGroup] from an unsafe.Pointer.
//
// A group of subitems in a toolbar item.
func ToolbarItemGroupFrom(ptr unsafe.Pointer) ToolbarItemGroup {
	return ToolbarItemGroup{
		ToolbarItem: ToolbarItemFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (tc _ToolbarItemGroupClass) Alloc() ToolbarItemGroup {
	rv := objc.Send[ToolbarItemGroup](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _ToolbarItemGroupClass) New() ToolbarItemGroup {
	rv := objc.Send[ToolbarItemGroup](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ ToolbarItemGroup) Init() ToolbarItemGroup {
	rv := objc.Send[ToolbarItemGroup](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ ToolbarItemGroup) Autorelease() ToolbarItemGroup {
	rv := objc.Send[ToolbarItemGroup](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewToolbarItemGroup creates a new ToolbarItemGroup instance.
func NewToolbarItemGroup() ToolbarItemGroup {
	return getToolbarItemGroupClass().New()
}



// Creates a grouped toolbar item with images.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItemGroup/init(itemIdentifier:images:selectionMode:labels:target:action:)
func NewToolbarItemGroupWithItemIdentifierImagesSelectionModeLabelsTargetAction(itemIdentifier objc.IObject /* cross-framework: ToolbarItemIdentifier */, images []Image, selectionMode ToolbarItemGroupSelectionMode, labels []string, target objc.IObject, action objc.SEL) ToolbarItemGroup {
	rv := objc.Send[ToolbarItemGroup](objc.ID(getToolbarItemGroupClass().class), objc.Sel("groupWithItemIdentifier:images:selectionMode:labels:target:action:"), itemIdentifier, images, selectionMode, labels, target, action)
	return rv
}


// Creates a grouped toolbar item with labels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItemGroup/init(itemIdentifier:titles:selectionMode:labels:target:action:)
func NewToolbarItemGroupWithItemIdentifierTitlesSelectionModeLabelsTargetAction(itemIdentifier objc.IObject /* cross-framework: ToolbarItemIdentifier */, titles []string, selectionMode ToolbarItemGroupSelectionMode, labels []string, target objc.IObject, action objc.SEL) ToolbarItemGroup {
	rv := objc.Send[ToolbarItemGroup](objc.ID(getToolbarItemGroupClass().class), objc.Sel("groupWithItemIdentifier:titles:selectionMode:labels:target:action:"), itemIdentifier, titles, selectionMode, labels, target, action)
	return rv
}



// Creates a grouped toolbar item with images.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItemGroup/init(itemIdentifier:images:selectionMode:labels:target:action:)
func (tc _ToolbarItemGroupClass) GroupWithItemIdentifierImagesSelectionModeLabelsTargetAction(itemIdentifier objc.IObject /* cross-framework: ToolbarItemIdentifier */, images []Image, selectionMode ToolbarItemGroupSelectionMode, labels []string, target objc.IObject, action objc.SEL) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(tc.class), objc.Sel("groupWithItemIdentifier:images:selectionMode:labels:target:action:"), itemIdentifier, images, selectionMode, labels, target, action)
	return rv
}


// Creates a grouped toolbar item with labels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItemGroup/init(itemIdentifier:titles:selectionMode:labels:target:action:)
func (tc _ToolbarItemGroupClass) GroupWithItemIdentifierTitlesSelectionModeLabelsTargetAction(itemIdentifier objc.IObject /* cross-framework: ToolbarItemIdentifier */, titles []string, selectionMode ToolbarItemGroupSelectionMode, labels []string, target objc.IObject, action objc.SEL) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(tc.class), objc.Sel("groupWithItemIdentifier:titles:selectionMode:labels:target:action:"), itemIdentifier, titles, selectionMode, labels, target, action)
	return rv
}


// Indicates whether a specified index is currently selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItemGroup/isSelected(at:)
func (t_ ToolbarItemGroup) IsSelectedAtIndex(index int) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isSelectedAtIndex:"), index)
	return rv
}


// Sets the selected state of a subitem in a grouped toolbar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItemGroup/setSelected(_:at:)
func (t_ ToolbarItemGroup) SetSelectedAtIndex(selected bool, index int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSelected:atIndex:"), selected, index)
}


// A value that represents how a toolbar displays a grouped toolbar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItemGroup/controlRepresentation-swift.property
func (t_ ToolbarItemGroup) ControlRepresentation() ToolbarItemGroupControlRepresentation {
	rv := objc.Send[ToolbarItemGroupControlRepresentation](t_.ID, objc.Sel("controlRepresentation"))
	return rv
}


// A value that represents how a toolbar displays a grouped toolbar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItemGroup/controlRepresentation-swift.property
func (t_ ToolbarItemGroup) SetControlRepresentation(value ToolbarItemGroupControlRepresentation) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setControlRepresentation:"), value)
}


// The index value for the most recently selected subitem of a grouped toolbar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItemGroup/selectedIndex
func (t_ ToolbarItemGroup) SelectedIndex() int {
	rv := objc.Send[int](t_.ID, objc.Sel("selectedIndex"))
	return rv
}


// The index value for the most recently selected subitem of a grouped toolbar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItemGroup/selectedIndex
func (t_ ToolbarItemGroup) SetSelectedIndex(value int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSelectedIndex:"), value)
}


// The selection mode of the grouped toolbar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItemGroup/selectionMode-swift.property
func (t_ ToolbarItemGroup) SelectionMode() ToolbarItemGroupSelectionMode {
	rv := objc.Send[ToolbarItemGroupSelectionMode](t_.ID, objc.Sel("selectionMode"))
	return rv
}


// The selection mode of the grouped toolbar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItemGroup/selectionMode-swift.property
func (t_ ToolbarItemGroup) SetSelectionMode(value ToolbarItemGroupSelectionMode) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSelectionMode:"), value)
}


// The subitems of the grouped toolbar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItemGroup/subitems
func (t_ ToolbarItemGroup) Subitems() []ToolbarItem {
	rv := objc.Send[[]ToolbarItem](t_.ID, objc.Sel("subitems"))
	return rv
}


// The subitems of the grouped toolbar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItemGroup/subitems
func (t_ ToolbarItemGroup) SetSubitems(value []ToolbarItem) {
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
	objc.Send[objc.ID](t_.ID, objc.Sel("setSubitems:"), nsArray)
}


