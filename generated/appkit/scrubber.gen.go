
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Scrubber] class.
var ScrubberClass _ScrubberClass

func init() {
	ScrubberClass = _ScrubberClass{objc.GetClass("NSScrubber")}
}

type _ScrubberClass struct {
	objc.Class
}

// An interface definition for the [Scrubber] class.
type IScrubber interface {
	ID() objc.ID
	InsertItemsAtIndexes(indexes unsafe.Pointer)
	ItemViewForItemAtIndex(index int) unsafe.Pointer
	MakeItemWithIdentifierOwner(itemIdentifier unsafe.Pointer, owner objc.ID) unsafe.Pointer
	MoveItemAtIndexToIndex(oldIndex int, newIndex int)
	PerformSequentialBatchUpdates(updateBlock unsafe.Pointer)
	RegisterClassForItemIdentifier(itemViewClass objc.Class, itemIdentifier unsafe.Pointer)
	RegisterNibForItemIdentifier(nib unsafe.Pointer, itemIdentifier unsafe.Pointer)
	ReloadData()
	ReloadItemsAtIndexes(indexes unsafe.Pointer)
	RemoveItemsAtIndexes(indexes unsafe.Pointer)
	ScrollItemAtIndexToAlignment(index int, alignment unsafe.Pointer)
}

type Scrubber struct {
	id objc.ID
}

func ScrubberFrom(ptr unsafe.Pointer) Scrubber {
	return Scrubber{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ Scrubber) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _ScrubberClass) Alloc() Scrubber {
	rv := objc.Send[Scrubber](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _ScrubberClass) New() Scrubber {
	rv := objc.Send[Scrubber](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewScrubber creates and returns a new initialized instance.
func NewScrubber() Scrubber {
	return ScrubberClass.New()
}

// Init initializes the instance.
func (s_ Scrubber) Init() Scrubber {
	rv := objc.Send[Scrubber](s_.ID(), selInit)
	return rv
}
// Inserts new items at the specified indexes into the scrubber. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubber/insertItems(at:)
func (s_ Scrubber) InsertItemsAtIndexes(indexes unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("insertItemsAtIndexes:"), indexes)
}
// Returns the view for the item at the specified index. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubber/itemViewForItem(at:)
func (s_ Scrubber) ItemViewForItemAtIndex(index int) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID(), objc.RegisterName("itemViewForItemAtIndex:"), index)
	return rv
}
// Creates or returns a reusable item object with the specified identifier. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubber/makeItem(withIdentifier:owner:)
func (s_ Scrubber) MakeItemWithIdentifierOwner(itemIdentifier unsafe.Pointer, owner objc.ID) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID(), objc.RegisterName("makeItemWithIdentifier:owner:"), itemIdentifier, owner)
	return rv
}
// Moves an item from one index to another in the scrubber. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubber/moveItem(at:to:)
func (s_ Scrubber) MoveItemAtIndexToIndex(oldIndex int, newIndex int) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("moveItemAtIndex:toIndex:"), oldIndex, newIndex)
}
// Combines multiple scrubber content updates into a single action. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubber/performSequentialBatchUpdates(_:)
func (s_ Scrubber) PerformSequentialBatchUpdates(updateBlock unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("performSequentialBatchUpdates:"), updateBlock)
}
// Registers a class for the scrubber to use when it creates new items. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubber/register(_:forItemIdentifier:)-2rb69
func (s_ Scrubber) RegisterClassForItemIdentifier(itemViewClass objc.Class, itemIdentifier unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("registerClass:forItemIdentifier:"), itemViewClass, itemIdentifier)
}
// Registers a nib file for the scrubber to use when it creates new items in the scrubber. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubber/register(_:forItemIdentifier:)-6jye0
func (s_ Scrubber) RegisterNibForItemIdentifier(nib unsafe.Pointer, itemIdentifier unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("registerNib:forItemIdentifier:"), nib, itemIdentifier)
}
// Reloads the content of the entire scrubber, and deselects the currently selected item. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubber/reloadData()
func (s_ Scrubber) ReloadData() {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("reloadData"))
}
// Reloads the items at the specified indexes. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubber/reloadItems(at:)
func (s_ Scrubber) ReloadItemsAtIndexes(indexes unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("reloadItemsAtIndexes:"), indexes)
}
// Removes the items at the specified indexes from the scrubber. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubber/removeItems(at:)
func (s_ Scrubber) RemoveItemsAtIndexes(indexes unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("removeItemsAtIndexes:"), indexes)
}
// Scrolls an item to a specified alignment within the scrubber. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubber/scrollItem(at:to:)
func (s_ Scrubber) ScrollItemAtIndexToAlignment(index int, alignment unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("scrollItemAtIndex:toAlignment:"), index, alignment)
}
// The color displayed behind the scrubber content. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubber/backgroundColor
func (s_ Scrubber) BackgroundColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID(), objc.RegisterName("backgroundColor"))
	return rv
}
// SetBackgroundColor sets the value of the backgroundColor property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubber/backgroundColor
func (s_ Scrubber) SetBackgroundColor(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setBackgroundColor:"), value)
}
// A view that is displayed behind the scrubber content. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubber/backgroundView
func (s_ Scrubber) BackgroundView() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID(), objc.RegisterName("backgroundView"))
	return rv
}
// SetBackgroundView sets the value of the backgroundView property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubber/backgroundView
func (s_ Scrubber) SetBackgroundView(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setBackgroundView:"), value)
}
// The object that provides the data for the scrubber. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubber/dataSource
func (s_ Scrubber) DataSource() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID(), objc.RegisterName("dataSource"))
	return rv
}
// SetDataSource sets the value of the dataSource property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubber/dataSource
func (s_ Scrubber) SetDataSource(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setDataSource:"), value)
}
// The object that acts as the delegate of the scrubber. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubber/delegate
func (s_ Scrubber) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID(), objc.RegisterName("delegate"))
	return rv
}
// SetDelegate sets the value of the delegate property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubber/delegate
func (s_ Scrubber) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setDelegate:"), value)
}
// A Boolean value that determines the behavior of the item selection decorations as the scrubber’s selection changes. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubber/floatsSelectionViews
func (s_ Scrubber) FloatsSelectionViews() bool {
	rv := objc.Send[bool](s_.ID(), objc.RegisterName("floatsSelectionViews"))
	return rv
}
// SetFloatsSelectionViews sets the value of the floatsSelectionViews property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubber/floatsSelectionViews
func (s_ Scrubber) SetFloatsSelectionViews(value bool) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setFloatsSelectionViews:"), value)
}
// The index of the highlighted item in the scrubber. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubber/highlightedIndex
func (s_ Scrubber) HighlightedIndex() int {
	rv := objc.Send[int](s_.ID(), objc.RegisterName("highlightedIndex"))
	return rv
}
// A Boolean value that, together with the   property, determines scrubber interaction style. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubber/isContinuous
func (s_ Scrubber) Continuous() bool {
	rv := objc.Send[bool](s_.ID(), objc.RegisterName("continuous"))
	return rv
}
// SetContinuous sets the value of the continuous property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubber/isContinuous
func (s_ Scrubber) SetContinuous(value bool) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setContinuous:"), value)
}
// A setting that specifies the snapping behavior of items in the scrubber. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubber/itemAlignment
func (s_ Scrubber) ItemAlignment() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID(), objc.RegisterName("itemAlignment"))
	return rv
}
// SetItemAlignment sets the value of the itemAlignment property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubber/itemAlignment
func (s_ Scrubber) SetItemAlignment(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setItemAlignment:"), value)
}
// A setting that determines whether interaction with the scrubber is fixed or free. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubber/mode-swift.property
func (s_ Scrubber) Mode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID(), objc.RegisterName("mode"))
	return rv
}
// SetMode sets the value of the mode property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubber/mode-swift.property
func (s_ Scrubber) SetMode(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setMode:"), value)
}
// The number of items represented by the scrubber. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubber/numberOfItems
func (s_ Scrubber) NumberOfItems() int {
	rv := objc.Send[int](s_.ID(), objc.RegisterName("numberOfItems"))
	return rv
}
// An object used to describe the layout of items within the scrubber. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubber/scrubberLayout
func (s_ Scrubber) ScrubberLayout() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID(), objc.RegisterName("scrubberLayout"))
	return rv
}
// SetScrubberLayout sets the value of the scrubberLayout property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubber/scrubberLayout
func (s_ Scrubber) SetScrubberLayout(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setScrubberLayout:"), value)
}
// The index of the selected item in the scrubber. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubber/selectedIndex
func (s_ Scrubber) SelectedIndex() int {
	rv := objc.Send[int](s_.ID(), objc.RegisterName("selectedIndex"))
	return rv
}
// SetSelectedIndex sets the value of the selectedIndex property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubber/selectedIndex
func (s_ Scrubber) SetSelectedIndex(value int) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setSelectedIndex:"), value)
}
// The style applied to the background of selected items. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubber/selectionBackgroundStyle
func (s_ Scrubber) SelectionBackgroundStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID(), objc.RegisterName("selectionBackgroundStyle"))
	return rv
}
// SetSelectionBackgroundStyle sets the value of the selectionBackgroundStyle property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubber/selectionBackgroundStyle
func (s_ Scrubber) SetSelectionBackgroundStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setSelectionBackgroundStyle:"), value)
}
// The style overlaid on selected items. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubber/selectionOverlayStyle
func (s_ Scrubber) SelectionOverlayStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID(), objc.RegisterName("selectionOverlayStyle"))
	return rv
}
// SetSelectionOverlayStyle sets the value of the selectionOverlayStyle property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubber/selectionOverlayStyle
func (s_ Scrubber) SetSelectionOverlayStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setSelectionOverlayStyle:"), value)
}
// A Boolean value that specifies whether the scrubber should display the existence of additional items beyond the leading and trailing edges. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubber/showsAdditionalContentIndicators
func (s_ Scrubber) ShowsAdditionalContentIndicators() bool {
	rv := objc.Send[bool](s_.ID(), objc.RegisterName("showsAdditionalContentIndicators"))
	return rv
}
// SetShowsAdditionalContentIndicators sets the value of the showsAdditionalContentIndicators property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubber/showsAdditionalContentIndicators
func (s_ Scrubber) SetShowsAdditionalContentIndicators(value bool) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setShowsAdditionalContentIndicators:"), value)
}
// A Boolean value that specifies whether arrow buttons should be displayed at the leading and trailing edges of the scrubber. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubber/showsArrowButtons
func (s_ Scrubber) ShowsArrowButtons() bool {
	rv := objc.Send[bool](s_.ID(), objc.RegisterName("showsArrowButtons"))
	return rv
}
// SetShowsArrowButtons sets the value of the showsArrowButtons property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubber/showsArrowButtons
func (s_ Scrubber) SetShowsArrowButtons(value bool) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setShowsArrowButtons:"), value)
}
