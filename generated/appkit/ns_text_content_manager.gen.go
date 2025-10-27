// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [TextContentManager] class.
var (
	TextContentManagerClass     _TextContentManagerClass
	TextContentManagerClassOnce sync.Once
)

func getTextContentManagerClass() _TextContentManagerClass {
	TextContentManagerClassOnce.Do(func() {
		TextContentManagerClass = _TextContentManagerClass{objc.GetClass("NSTextContentManager")}
	})
	return TextContentManagerClass
}

type _TextContentManagerClass struct {
	class objc.Class
}





// An interface definition for the [TextContentManager] class.
type ITextContentManager interface {
	objectivec.IObject
	

	// properties:
	AutomaticallySynchronizesTextLayoutManagers() bool
	SetAutomaticallySynchronizesTextLayoutManagers(value bool)
	AutomaticallySynchronizesToBackingStore() bool
	SetAutomaticallySynchronizesToBackingStore(value bool)
	HasEditingTransaction() bool
	PrimaryTextLayoutManager() ITextLayoutManager
	SetPrimaryTextLayoutManager(value ITextLayoutManager)
	TextLayoutManagers() []TextLayoutManager


	

	// methods:
	AddTextLayoutManager(textLayoutManager ITextLayoutManager)
	PerformEditingTransactionUsingBlock(transaction unsafe.Pointer)
	RecordEditActionInRangeNewTextRange(originalTextRange ITextRange, newTextRange ITextRange)
	RemoveTextLayoutManager(textLayoutManager ITextLayoutManager)
	SynchronizeTextLayoutManagers(completionHandler unsafe.Pointer)
	TextElementsForRange(range_ ITextRange) []TextElement


}





// Alloc allocates a new instance without initialization.
func (tc _TextContentManagerClass) Alloc() TextContentManager {
	rv := objc.Send[TextContentManager](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TextContentManagerClass) New() TextContentManager {
	rv := objc.Send[TextContentManager](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TextContentManager) Init() TextContentManager {
	rv := objc.Send[TextContentManager](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TextContentManager) Autorelease() TextContentManager {
	rv := objc.Send[TextContentManager](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextContentManager creates a new TextContentManager instance.
func NewTextContentManager() TextContentManager {
	return getTextContentManagerClass().New()
}





// An abstract class that defines the interface and a default implementation for managing the text document contents.


// An abstract class that defines the interface and a default implementation for managing the text document contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContentManager
type TextContentManager struct {
	objectivec.Object
}

// TextContentManagerFrom constructs a [TextContentManager] from an unsafe.Pointer.
//
// An abstract class that defines the interface and a default implementation for managing the text document contents.
func TextContentManagerFrom(ptr unsafe.Pointer) TextContentManager {
	return TextContentManager{objectivec.Object{objc.ID(ptr)}}
}






// Creates a new content manager object from data in an unarchiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContentManager/init(coder:)
func NewTextContentManagerWithCoder(coder foundation.foundation.INSCoder) TextContentManager {
	instance := getTextContentManagerClass().Alloc()
	rv := objc.Send[TextContentManager](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}

















// Adds the layout manager you provide to the list of layout managers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContentManager/addTextLayoutManager(_:)
func (t_ TextContentManager) AddTextLayoutManager(textLayoutManager ITextLayoutManager) {
	objc.Send[objc.ID](t_.ID, objc.Sel("addTextLayoutManager:"), textLayoutManager)
}


// Performs an editing transaction and invokes a block upon completion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContentManager/performEditingTransaction(_:)
func (t_ TextContentManager) PerformEditingTransactionUsingBlock(transaction unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("performEditingTransactionUsingBlock:"), transaction)
}


// Records information about an edit action to the transaction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContentManager/recordEditAction(in:newTextRange:)
func (t_ TextContentManager) RecordEditActionInRangeNewTextRange(originalTextRange ITextRange, newTextRange ITextRange) {
	objc.Send[objc.ID](t_.ID, objc.Sel("recordEditActionInRange:newTextRange:"), originalTextRange, newTextRange)
}


// Removes the layout manager you specifiy from the list of layout managers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContentManager/removeTextLayoutManager(_:)
func (t_ TextContentManager) RemoveTextLayoutManager(textLayoutManager ITextLayoutManager) {
	objc.Send[objc.ID](t_.ID, objc.Sel("removeTextLayoutManager:"), textLayoutManager)
}


// Synchronizes changes to all nonprimary text layout managers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContentManager/synchronizeTextLayoutManagers(_:)
func (t_ TextContentManager) SynchronizeTextLayoutManagers(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("synchronizeTextLayoutManagers:"), completionHandler)
}


// Returns an array of text elements that intersect with the range you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContentManager/textElements(for:)
func (t_ TextContentManager) TextElementsForRange(range_ ITextRange) []TextElement {
	rv := objc.Send[[]TextElement](t_.ID, objc.Sel("textElementsForRange:"), range_)
	return rv
}







// Determines if the framework should automatically synchronize all text layout managers when exiting an editing transaction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContentManager/automaticallySynchronizesTextLayoutManagers
func (t_ TextContentManager) AutomaticallySynchronizesTextLayoutManagers() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("automaticallySynchronizesTextLayoutManagers"))
	return rv
}


// Determines if the framework should automatically synchronize all text layout managers when exiting an editing transaction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContentManager/automaticallySynchronizesTextLayoutManagers
func (t_ TextContentManager) SetAutomaticallySynchronizesTextLayoutManagers(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAutomaticallySynchronizesTextLayoutManagers:"), value)
}


// Determines whether to automatically synchronize with the backing store when an editing transaction finishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContentManager/automaticallySynchronizesToBackingStore
func (t_ TextContentManager) AutomaticallySynchronizesToBackingStore() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("automaticallySynchronizesToBackingStore"))
	return rv
}


// Determines whether to automatically synchronize with the backing store when an editing transaction finishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContentManager/automaticallySynchronizesToBackingStore
func (t_ TextContentManager) SetAutomaticallySynchronizesToBackingStore(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAutomaticallySynchronizesToBackingStore:"), value)
}


// Indicates there’s an active editing transaction from the primary text layout manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContentManager/hasEditingTransaction
func (t_ TextContentManager) HasEditingTransaction() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("hasEditingTransaction"))
	return rv
}


// The primary text layout manager for this content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContentManager/primaryTextLayoutManager
func (t_ TextContentManager) PrimaryTextLayoutManager() ITextLayoutManager {
	rv := objc.Send[TextLayoutManager](t_.ID, objc.Sel("primaryTextLayoutManager"))
	return rv
}


// The primary text layout manager for this content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContentManager/primaryTextLayoutManager
func (t_ TextContentManager) SetPrimaryTextLayoutManager(value ITextLayoutManager) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPrimaryTextLayoutManager:"), value)
}


// The array of text layout managers associated with this text content manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContentManager/textLayoutManagers
func (t_ TextContentManager) TextLayoutManagers() []TextLayoutManager {
	rv := objc.Send[[]TextLayoutManager](t_.ID, objc.Sel("textLayoutManagers"))
	return rv
}







