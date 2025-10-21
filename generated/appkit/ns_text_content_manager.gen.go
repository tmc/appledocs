// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
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
	PerformEditingTransactionUsingBlock(transaction unsafe.Pointer)
	RecordEditActionInRangeNewTextRange(originalTextRange unsafe.Pointer, newTextRange unsafe.Pointer)
}

// An abstract class that defines the interface and a default implementation for managing the text document contents.
//
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

// Alloc allocates a new instance without initialization.
func (tc _TextContentManagerClass) Alloc() TextContentManager {
	rv := objc.Send[TextContentManager](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Performs an editing transaction and invokes a block upon completion.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContentManager/performEditingTransaction(_:)
func (t_ TextContentManager) PerformEditingTransactionUsingBlock(transaction unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("performEditingTransactionUsingBlock:"), transaction)
}

// Records information about an edit action to the transaction.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContentManager/recordEditAction(in:newTextRange:)
func (t_ TextContentManager) RecordEditActionInRangeNewTextRange(originalTextRange unsafe.Pointer, newTextRange unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("recordEditActionInRange:newTextRange:"), originalTextRange, newTextRange)
}

// Determines if the framework should automatically synchronize all text layout managers when exiting an editing transaction.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContentManager/automaticallySynchronizesTextLayoutManagers
func (t_ TextContentManager) AutomaticallySynchronizesTextLayoutManagers() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("automaticallySynchronizesTextLayoutManagers"))
	return rv
}


// SetAutomaticallySynchronizesTextLayoutManagers sets the value of the automaticallySynchronizesTextLayoutManagers property.
// Determines if the framework should automatically synchronize all text layout managers when exiting an editing transaction.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContentManager/automaticallySynchronizesTextLayoutManagers
func (t_ TextContentManager) SetAutomaticallySynchronizesTextLayoutManagers(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAutomaticallySynchronizesTextLayoutManagers:"), value)
}

// Indicates there’s an active editing transaction from the primary text layout manager.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContentManager/hasEditingTransaction
func (t_ TextContentManager) HasEditingTransaction() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("hasEditingTransaction"))
	return rv
}

// The array of text layout managers associated with this text content manager.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContentManager/textLayoutManagers
func (t_ TextContentManager) TextLayoutManagers() []TextLayoutManager {
	rv := objc.Send[[]TextLayoutManager](t_.ID, objc.Sel("textLayoutManagers"))
	return rv
}



