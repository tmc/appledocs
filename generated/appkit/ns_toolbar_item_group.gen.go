// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSToolbarItemGroup */


/* debug [class_header]: Header for NSToolbarItemGroup */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ToolbarItemGroup */
// An interface definition for the [ToolbarItemGroup] class.
type IToolbarItemGroup interface {
	IToolbarItem
	
/* debug [class_interface_properties]: Properties for ToolbarItemGroup */
	// properties:
	ControlRepresentation() ToolbarItemGroupControlRepresentation
	SetControlRepresentation(value ToolbarItemGroupControlRepresentation)
	SelectedIndex() int
	SetSelectedIndex(value int)
	SelectionMode() ToolbarItemGroupSelectionMode
	SetSelectionMode(value ToolbarItemGroupSelectionMode)
	Subitems() []ToolbarItem
	SetSubitems(value []ToolbarItem)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ToolbarItemGroup */
	// methods:
	IsSelectedAtIndex(index int) bool
	SetSelectedAtIndex(selected bool, index int)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ToolbarItemGroup */
// Alloc allocates a new instance without initialization.
func (tc _ToolbarItemGroupClass) Alloc() ToolbarItemGroup {
	rv := objc.Send[ToolbarItemGroup](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ToolbarItemGroup */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ToolbarItemGroup */

// Creates a grouped toolbar item with images.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItemGroup/init(itemIdentifier:images:selectionMode:labels:target:action:)
func NewToolbarItemGroupWithItemIdentifierImagesSelectionModeLabelsTargetAction(itemIdentifier ToolbarItemIdentifier /* typedef */, images []Image, selectionMode ToolbarItemGroupSelectionMode, labels []string, target objc.IObject, action objc.SEL) ToolbarItemGroup {
	rv := objc.Send[ToolbarItemGroup](objc.ID(getToolbarItemGroupClass().class), objc.Sel("groupWithItemIdentifier:images:selectionMode:labels:target:action:"), itemIdentifier, images, selectionMode, labels, target, action)
	return rv
}/* debug [class_init_methods/constructor]: NewToolbarItemGroupWithItemIdentifierImagesSelectionModeLabelsTargetAction */


// Creates a grouped toolbar item with labels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItemGroup/init(itemIdentifier:titles:selectionMode:labels:target:action:)
func NewToolbarItemGroupWithItemIdentifierTitlesSelectionModeLabelsTargetAction(itemIdentifier ToolbarItemIdentifier /* typedef */, titles []string, selectionMode ToolbarItemGroupSelectionMode, labels []string, target objc.IObject, action objc.SEL) ToolbarItemGroup {
	rv := objc.Send[ToolbarItemGroup](objc.ID(getToolbarItemGroupClass().class), objc.Sel("groupWithItemIdentifier:titles:selectionMode:labels:target:action:"), itemIdentifier, titles, selectionMode, labels, target, action)
	return rv
}/* debug [class_init_methods/constructor]: NewToolbarItemGroupWithItemIdentifierTitlesSelectionModeLabelsTargetAction */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ToolbarItemGroup */

// Creates a grouped toolbar item with images.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItemGroup/init(itemIdentifier:images:selectionMode:labels:target:action:)
func (tc _ToolbarItemGroupClass) GroupWithItemIdentifierImagesSelectionModeLabelsTargetAction(itemIdentifier ToolbarItemIdentifier /* typedef */, images []Image, selectionMode ToolbarItemGroupSelectionMode, labels []string, target objc.IObject, action objc.SEL) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(tc.class), objc.Sel("groupWithItemIdentifier:images:selectionMode:labels:target:action:"), itemIdentifier, images, selectionMode, labels, target, action)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=GroupWithItemIdentifierImagesSelectionModeLabelsTargetAction) */


// Creates a grouped toolbar item with labels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItemGroup/init(itemIdentifier:titles:selectionMode:labels:target:action:)
func (tc _ToolbarItemGroupClass) GroupWithItemIdentifierTitlesSelectionModeLabelsTargetAction(itemIdentifier ToolbarItemIdentifier /* typedef */, titles []string, selectionMode ToolbarItemGroupSelectionMode, labels []string, target objc.IObject, action objc.SEL) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(tc.class), objc.Sel("groupWithItemIdentifier:titles:selectionMode:labels:target:action:"), itemIdentifier, titles, selectionMode, labels, target, action)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=GroupWithItemIdentifierTitlesSelectionModeLabelsTargetAction) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ToolbarItemGroup */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ToolbarItemGroup */

// Indicates whether a specified index is currently selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItemGroup/isSelected(at:)
func (t_ ToolbarItemGroup) IsSelectedAtIndex(index int) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isSelectedAtIndex:"), index)
	return rv
}/* debug [instance_methods/method]: IsSelectedAtIndex */


// Sets the selected state of a subitem in a grouped toolbar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItemGroup/setSelected(_:at:)
func (t_ ToolbarItemGroup) SetSelectedAtIndex(selected bool, index int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSelected:atIndex:"), selected, index)
}/* debug [instance_methods/method]: SetSelectedAtIndex */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ToolbarItemGroup */

// A value that represents how a toolbar displays a grouped toolbar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItemGroup/controlRepresentation-swift.property
func (t_ ToolbarItemGroup) ControlRepresentation() ToolbarItemGroupControlRepresentation {
	rv := objc.Send[ToolbarItemGroupControlRepresentation](t_.ID, objc.Sel("controlRepresentation"))
	return rv
}/* debug [instance_properties/getter]: controlRepresentation */


// A value that represents how a toolbar displays a grouped toolbar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItemGroup/controlRepresentation-swift.property
func (t_ ToolbarItemGroup) SetControlRepresentation(value ToolbarItemGroupControlRepresentation) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setControlRepresentation:"), value)
}/* debug [instance_properties/setter]: controlRepresentation */


// The index value for the most recently selected subitem of a grouped toolbar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItemGroup/selectedIndex
func (t_ ToolbarItemGroup) SelectedIndex() int {
	rv := objc.Send[int](t_.ID, objc.Sel("selectedIndex"))
	return rv
}/* debug [instance_properties/getter]: selectedIndex */


// The index value for the most recently selected subitem of a grouped toolbar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItemGroup/selectedIndex
func (t_ ToolbarItemGroup) SetSelectedIndex(value int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSelectedIndex:"), value)
}/* debug [instance_properties/setter]: selectedIndex */


// The selection mode of the grouped toolbar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItemGroup/selectionMode-swift.property
func (t_ ToolbarItemGroup) SelectionMode() ToolbarItemGroupSelectionMode {
	rv := objc.Send[ToolbarItemGroupSelectionMode](t_.ID, objc.Sel("selectionMode"))
	return rv
}/* debug [instance_properties/getter]: selectionMode */


// The selection mode of the grouped toolbar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItemGroup/selectionMode-swift.property
func (t_ ToolbarItemGroup) SetSelectionMode(value ToolbarItemGroupSelectionMode) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSelectionMode:"), value)
}/* debug [instance_properties/setter]: selectionMode */


// The subitems of the grouped toolbar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItemGroup/subitems
func (t_ ToolbarItemGroup) Subitems() []ToolbarItem {
	rv := objc.Send[[]ToolbarItem](t_.ID, objc.Sel("subitems"))
	return rv
}/* debug [instance_properties/getter]: subitems */


// The subitems of the grouped toolbar item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItemGroup/subitems
func (t_ ToolbarItemGroup) SetSubitems(value []ToolbarItem) {
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
}/* debug [instance_properties/setter]: subitems */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSToolbarItemGroup */


