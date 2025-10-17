// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Scrubber] class.
var ScrubberClass objc.Class

func init() {
	ScrubberClass = objc.GetClass("NSScrubber")
}

type Scrubber struct {
	objc.ID
}

func ScrubberFrom(ptr unsafe.Pointer) Scrubber {
	return Scrubber{
		ID: objc.ID(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (sc Scrubber) Alloc() Scrubber {
	ret := objc.ID(ScrubberClass).Send(objc.RegisterName("alloc"))
	return Scrubber{ret}
}

// Init initializes the instance.
func (s_ Scrubber) Init() Scrubber {
	ret := s_.ID.Send(objc.RegisterName("init"))
	return Scrubber{ret}
}
// Initializes and returns a newly allocated scrubber object from a storyboard or nib file. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubber/init(coder:)
func NewScrubberWithCoder(coder unsafe.Pointer) Scrubber {
	instance := Scrubber{}.Alloc()
	sel := objc.RegisterName("initWithCoder:")
	ret := instance.ID.Send(sel, coder)
	instance = Scrubber{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
// Initializes and returns a newly allocated scrubber object with the specified frame rectangle. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubber/init(frame:)
func NewScrubberWithFrame(frameRect unsafe.Pointer) Scrubber {
	instance := Scrubber{}.Alloc()
	sel := objc.RegisterName("initWithFrame:")
	ret := instance.ID.Send(sel, frameRect)
	instance = Scrubber{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}


// Inserts new items at the specified indexes into the scrubber. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubber/insertItems(at:)
func (s_ Scrubber) InsertItemsAtIndexes(indexes unsafe.Pointer) {
	sel := objc.RegisterName("insertItemsAtIndexes:")
	s_.ID.Send(sel, indexes)
}
// Returns the view for the item at the specified index. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubber/itemViewForItem(at:)
func (s_ Scrubber) ItemViewForItemAtIndex(index int) unsafe.Pointer {
	sel := objc.RegisterName("itemViewForItemAtIndex:")
	ret := s_.ID.Send(sel, index)
	return unsafe.Pointer(ret)
}
// Creates or returns a reusable item object with the specified identifier. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubber/makeItem(withIdentifier:owner:)
func (s_ Scrubber) MakeItemWithIdentifierOwner(itemIdentifier unsafe.Pointer, owner objc.ID) unsafe.Pointer {
	sel := objc.RegisterName("makeItemWithIdentifier:owner:")
	ret := s_.ID.Send(sel, itemIdentifier, owner)
	return unsafe.Pointer(ret)
}
// Moves an item from one index to another in the scrubber. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubber/moveItem(at:to:)
func (s_ Scrubber) MoveItemAtIndexToIndex(oldIndex int, newIndex int) {
	sel := objc.RegisterName("moveItemAtIndex:toIndex:")
	s_.ID.Send(sel, oldIndex, newIndex)
}
// Combines multiple scrubber content updates into a single action. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubber/performSequentialBatchUpdates(_:)
func (s_ Scrubber) PerformSequentialBatchUpdates(updateBlock unsafe.Pointer) {
	sel := objc.RegisterName("performSequentialBatchUpdates:")
	s_.ID.Send(sel, updateBlock)
}
// Registers a class for the scrubber to use when it creates new items. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubber/register(_:forItemIdentifier:)-2rb69
func (s_ Scrubber) RegisterClassForItemIdentifier(itemViewClass objc.Class, itemIdentifier unsafe.Pointer) {
	sel := objc.RegisterName("registerClass:forItemIdentifier:")
	s_.ID.Send(sel, itemViewClass, itemIdentifier)
}
// Registers a nib file for the scrubber to use when it creates new items in the scrubber. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubber/register(_:forItemIdentifier:)-6jye0
func (s_ Scrubber) RegisterNibForItemIdentifier(nib unsafe.Pointer, itemIdentifier unsafe.Pointer) {
	sel := objc.RegisterName("registerNib:forItemIdentifier:")
	s_.ID.Send(sel, nib, itemIdentifier)
}
// Reloads the content of the entire scrubber, and deselects the currently selected item. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubber/reloadData()
func (s_ Scrubber) ReloadData() {
	sel := objc.RegisterName("reloadData")
	s_.ID.Send(sel)
}
// Reloads the items at the specified indexes. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubber/reloadItems(at:)
func (s_ Scrubber) ReloadItemsAtIndexes(indexes unsafe.Pointer) {
	sel := objc.RegisterName("reloadItemsAtIndexes:")
	s_.ID.Send(sel, indexes)
}
// Removes the items at the specified indexes from the scrubber. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubber/removeItems(at:)
func (s_ Scrubber) RemoveItemsAtIndexes(indexes unsafe.Pointer) {
	sel := objc.RegisterName("removeItemsAtIndexes:")
	s_.ID.Send(sel, indexes)
}
// Scrolls an item to a specified alignment within the scrubber. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrubber/scrollItem(at:to:)
func (s_ Scrubber) ScrollItemAtIndexToAlignment(index int, alignment unsafe.Pointer) {
	sel := objc.RegisterName("scrollItemAtIndex:toAlignment:")
	s_.ID.Send(sel, index, alignment)
}

