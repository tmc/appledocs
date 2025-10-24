// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSTextContentManager */


/* debug [class_header]: Header for NSTextContentManager */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TextContentManager */
// An interface definition for the [TextContentManager] class.
type ITextContentManager interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for TextContentManager */
	// properties:
	AutomaticallySynchronizesTextLayoutManagers() bool
	SetAutomaticallySynchronizesTextLayoutManagers(value bool)
	AutomaticallySynchronizesToBackingStore() bool
	SetAutomaticallySynchronizesToBackingStore(value bool)
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	HasEditingTransaction() bool
	PrimaryTextLayoutManager() ITextLayoutManager
	SetPrimaryTextLayoutManager(value ITextLayoutManager)
	TextLayoutManagers() []TextLayoutManager
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TextContentManager */
	// methods:
	AddTextLayoutManager(textLayoutManager ITextLayoutManager)
	PerformEditingTransactionUsingBlock(transaction unsafe.Pointer)
	RecordEditActionInRangeNewTextRange(originalTextRange ITextRange, newTextRange ITextRange)
	RemoveTextLayoutManager(textLayoutManager ITextLayoutManager)
	SynchronizeTextLayoutManagers(completionHandler unsafe.Pointer)
	TextElementsForRange(range_ ITextRange) []TextElement
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TextContentManager */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TextContentManager */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TextContentManager */

// Creates a new content manager object from data in an unarchiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContentManager/init(coder:)
func NewTextContentManagerWithCoder(coder foundation.Coder) TextContentManager {
	instance := getTextContentManagerClass().Alloc()
	rv := objc.Send[TextContentManager](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewTextContentManagerWithCoder */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TextContentManager */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TextContentManager */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TextContentManager */

// Adds the layout manager you provide to the list of layout managers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContentManager/addTextLayoutManager(_:)
func (t_ TextContentManager) AddTextLayoutManager(textLayoutManager ITextLayoutManager) {
	objc.Send[objc.ID](t_.ID, objc.Sel("addTextLayoutManager:"), textLayoutManager)
}/* debug [instance_methods/method]: AddTextLayoutManager */


// Performs an editing transaction and invokes a block upon completion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContentManager/performEditingTransaction(_:)
func (t_ TextContentManager) PerformEditingTransactionUsingBlock(transaction unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("performEditingTransactionUsingBlock:"), transaction)
}/* debug [instance_methods/method]: PerformEditingTransactionUsingBlock */


// Records information about an edit action to the transaction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContentManager/recordEditAction(in:newTextRange:)
func (t_ TextContentManager) RecordEditActionInRangeNewTextRange(originalTextRange ITextRange, newTextRange ITextRange) {
	objc.Send[objc.ID](t_.ID, objc.Sel("recordEditActionInRange:newTextRange:"), originalTextRange, newTextRange)
}/* debug [instance_methods/method]: RecordEditActionInRangeNewTextRange */


// Removes the layout manager you specifiy from the list of layout managers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContentManager/removeTextLayoutManager(_:)
func (t_ TextContentManager) RemoveTextLayoutManager(textLayoutManager ITextLayoutManager) {
	objc.Send[objc.ID](t_.ID, objc.Sel("removeTextLayoutManager:"), textLayoutManager)
}/* debug [instance_methods/method]: RemoveTextLayoutManager */


// Synchronizes changes to all nonprimary text layout managers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContentManager/synchronizeTextLayoutManagers(_:)
func (t_ TextContentManager) SynchronizeTextLayoutManagers(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("synchronizeTextLayoutManagers:"), completionHandler)
}/* debug [instance_methods/method]: SynchronizeTextLayoutManagers */


// Returns an array of text elements that intersect with the range you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContentManager/textElements(for:)
func (t_ TextContentManager) TextElementsForRange(range_ ITextRange) []TextElement {
	rv := objc.Send[[]TextElement](t_.ID, objc.Sel("textElementsForRange:"), range_)
	return rv
}/* debug [instance_methods/method]: TextElementsForRange */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TextContentManager */

// Determines if the framework should automatically synchronize all text layout managers when exiting an editing transaction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContentManager/automaticallySynchronizesTextLayoutManagers
func (t_ TextContentManager) AutomaticallySynchronizesTextLayoutManagers() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("automaticallySynchronizesTextLayoutManagers"))
	return rv
}/* debug [instance_properties/getter]: automaticallySynchronizesTextLayoutManagers */


// Determines if the framework should automatically synchronize all text layout managers when exiting an editing transaction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContentManager/automaticallySynchronizesTextLayoutManagers
func (t_ TextContentManager) SetAutomaticallySynchronizesTextLayoutManagers(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAutomaticallySynchronizesTextLayoutManagers:"), value)
}/* debug [instance_properties/setter]: automaticallySynchronizesTextLayoutManagers */


// Determines whether to automatically synchronize with the backing store when an editing transaction finishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContentManager/automaticallySynchronizesToBackingStore
func (t_ TextContentManager) AutomaticallySynchronizesToBackingStore() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("automaticallySynchronizesToBackingStore"))
	return rv
}/* debug [instance_properties/getter]: automaticallySynchronizesToBackingStore */


// Determines whether to automatically synchronize with the backing store when an editing transaction finishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContentManager/automaticallySynchronizesToBackingStore
func (t_ TextContentManager) SetAutomaticallySynchronizesToBackingStore(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAutomaticallySynchronizesToBackingStore:"), value)
}/* debug [instance_properties/setter]: automaticallySynchronizesToBackingStore */


// The delegate for the content manager object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContentManager/delegate
func (t_ TextContentManager) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The delegate for the content manager object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContentManager/delegate
func (t_ TextContentManager) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// Indicates there’s an active editing transaction from the primary text layout manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContentManager/hasEditingTransaction
func (t_ TextContentManager) HasEditingTransaction() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("hasEditingTransaction"))
	return rv
}/* debug [instance_properties/getter]: hasEditingTransaction */


// The primary text layout manager for this content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContentManager/primaryTextLayoutManager
func (t_ TextContentManager) PrimaryTextLayoutManager() ITextLayoutManager {
	rv := objc.Send[TextLayoutManager](t_.ID, objc.Sel("primaryTextLayoutManager"))
	return rv
}/* debug [instance_properties/getter]: primaryTextLayoutManager */


// The primary text layout manager for this content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContentManager/primaryTextLayoutManager
func (t_ TextContentManager) SetPrimaryTextLayoutManager(value ITextLayoutManager) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPrimaryTextLayoutManager:"), value)
}/* debug [instance_properties/setter]: primaryTextLayoutManager */


// The array of text layout managers associated with this text content manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContentManager/textLayoutManagers
func (t_ TextContentManager) TextLayoutManagers() []TextLayoutManager {
	rv := objc.Send[[]TextLayoutManager](t_.ID, objc.Sel("textLayoutManagers"))
	return rv
}/* debug [instance_properties/getter]: textLayoutManagers */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSTextContentManager */


