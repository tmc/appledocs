// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [Scrubber] class.
var (
	scrubberClass     _ScrubberClass
	scrubberClassOnce sync.Once
)

func getScrubberClass() _ScrubberClass {
	scrubberClassOnce.Do(func() {
		scrubberClass = _ScrubberClass{objc.GetClass("NSScrubber")}
	})
	return scrubberClass
}

type _ScrubberClass struct {
	class objc.Class
}

// An interface definition for the [Scrubber] class.
type IScrubber interface {
	IView
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

// A customizable item picker control for the Touch Bar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubber

type Scrubber struct {
	View
}

// ScrubberFrom constructs a [Scrubber] from an unsafe.Pointer.
//
// A customizable item picker control for the Touch Bar.
func ScrubberFrom(ptr unsafe.Pointer) Scrubber {
	return Scrubber{
		View: ViewFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (sc _ScrubberClass) Alloc() Scrubber {
	rv := objc.Send[Scrubber](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _ScrubberClass) New() Scrubber {
	rv := objc.Send[Scrubber](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ Scrubber) Init() Scrubber {
	rv := objc.Send[Scrubber](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ Scrubber) Autorelease() Scrubber {
	rv := objc.Send[Scrubber](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewScrubber creates a new Scrubber instance.
func NewScrubber() Scrubber {
	return getScrubberClass().New()
}


// Initializes and returns a newly allocated scrubber object from a storyboard or nib file. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubber/init(coder:)
func NewScrubberWithCoder(coder unsafe.Pointer) Scrubber {
	// Instance methods (init*) require Autorelease() to balance the +1 from alloc
	instance := getScrubberClass().Alloc()
	rv := objc.Send[Scrubber](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}
// Initializes and returns a newly allocated scrubber object with the specified frame rectangle. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubber/init(frame:)
func NewScrubberWithFrame(frameRect unsafe.Pointer) Scrubber {
	// Instance methods (init*) require Autorelease() to balance the +1 from alloc
	instance := getScrubberClass().Alloc()
	rv := objc.Send[Scrubber](instance.ID, objc.Sel("initWithFrame:"), frameRect)
	rv.Autorelease()
	return rv
}


// Inserts new items at the specified indexes into the scrubber. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubber/insertItems(at:)
func (s_ Scrubber) InsertItemsAtIndexes(indexes unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("insertItemsAtIndexes:"), indexes)
}
// Returns the view for the item at the specified index. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubber/itemViewForItem(at:)
func (s_ Scrubber) ItemViewForItemAtIndex(index int) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("itemViewForItemAtIndex:"), index)
	return rv
}
// Creates or returns a reusable item object with the specified identifier. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubber/makeItem(withIdentifier:owner:)
func (s_ Scrubber) MakeItemWithIdentifierOwner(itemIdentifier unsafe.Pointer, owner objc.ID) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("makeItemWithIdentifier:owner:"), itemIdentifier, owner)
	return rv
}
// Moves an item from one index to another in the scrubber. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubber/moveItem(at:to:)
func (s_ Scrubber) MoveItemAtIndexToIndex(oldIndex int, newIndex int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("moveItemAtIndex:toIndex:"), oldIndex, newIndex)
}
// Combines multiple scrubber content updates into a single action. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubber/performSequentialBatchUpdates(_:)
func (s_ Scrubber) PerformSequentialBatchUpdates(updateBlock unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("performSequentialBatchUpdates:"), updateBlock)
}
// Registers a class for the scrubber to use when it creates new items. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubber/register(_:forItemIdentifier:)-2rb69
func (s_ Scrubber) RegisterClassForItemIdentifier(itemViewClass objc.Class, itemIdentifier unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("registerClass:forItemIdentifier:"), itemViewClass, itemIdentifier)
}
// Registers a nib file for the scrubber to use when it creates new items in the scrubber. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubber/register(_:forItemIdentifier:)-6jye0
func (s_ Scrubber) RegisterNibForItemIdentifier(nib unsafe.Pointer, itemIdentifier unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("registerNib:forItemIdentifier:"), nib, itemIdentifier)
}
// Reloads the content of the entire scrubber, and deselects the currently selected item. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubber/reloadData()
func (s_ Scrubber) ReloadData() {
	objc.Send[objc.ID](s_.ID, objc.Sel("reloadData"))
}
// Reloads the items at the specified indexes. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubber/reloadItems(at:)
func (s_ Scrubber) ReloadItemsAtIndexes(indexes unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("reloadItemsAtIndexes:"), indexes)
}
// Removes the items at the specified indexes from the scrubber. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubber/removeItems(at:)
func (s_ Scrubber) RemoveItemsAtIndexes(indexes unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("removeItemsAtIndexes:"), indexes)
}
// Scrolls an item to a specified alignment within the scrubber. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubber/scrollItem(at:to:)
func (s_ Scrubber) ScrollItemAtIndexToAlignment(index int, alignment unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("scrollItemAtIndex:toAlignment:"), index, alignment)
}

