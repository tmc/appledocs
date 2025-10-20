// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [UndoManager] class.
var (
	undoManagerClass     _UndoManagerClass
	undoManagerClassOnce sync.Once
)

func getUndoManagerClass() _UndoManagerClass {
	undoManagerClassOnce.Do(func() {
		undoManagerClass = _UndoManagerClass{objc.GetClass("NSUndoManager")}
	})
	return undoManagerClass
}

type _UndoManagerClass struct {
	class objc.Class
}

// An interface definition for the [UndoManager] class.
type IUndoManager interface {
	objectivec.IObject
	EndUndoGrouping()
	Redo()
	RegisterUndoWithTargetSelectorObject(target objc.ID, selector objc.SEL, object objc.ID)
	RemoveAllActions()
	SetActionName(actionName string)
	SetActionUserInfoValueForKey(info objc.ID, key unsafe.Pointer)
	Undo()
	UndoActionUserInfoValueForKey(key unsafe.Pointer) objc.ID
	UndoMenuTitleForUndoActionName(actionName string) unsafe.Pointer
	UndoNestedGroup()
}

// A general-purpose recorder of operations that enables undo and redo.
//
// You register an undo operation by calling one of the methods described in Registering undo operations. You specify the name of the object that’s changing (or the owner of that object) and provide a closure, method, or invocation to revert its state. After you register an undo operation, you can call on the undo manager to revert to the state of the last undo operation. When undoing an action, saves the operations you revert to so that you can call automatically. Typically, apps with UI interactions work with . For example, UIKit implements undo and redo in its text view object, making it easy for you to undo and redo actions in objects along the responder chain. also serves as a general-purpose state manager, which you can use to undo and redo many kinds of actions. For example, an interactive command-line utility can use this class to undo the last command run, or a networking library can undo a request by sending another request that invalidates the previous one.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UndoManager
type UndoManager struct {
	objectivec.Object
}

// UndoManagerFrom constructs a [UndoManager] from an unsafe.Pointer.
//
// A general-purpose recorder of operations that enables undo and redo.
func UndoManagerFrom(ptr unsafe.Pointer) UndoManager {
	return UndoManager{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (uc _UndoManagerClass) Alloc() UndoManager {
	rv := objc.Send[UndoManager](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _UndoManagerClass) New() UndoManager {
	rv := objc.Send[UndoManager](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UndoManager) Init() UndoManager {
	rv := objc.Send[UndoManager](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UndoManager) Autorelease() UndoManager {
	rv := objc.Send[UndoManager](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUndoManager creates a new UndoManager instance.
func NewUndoManager() UndoManager {
	return getUndoManagerClass().New()
}


// Marks the end of an undo group.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UndoManager/endUndoGrouping()
func (u_ UndoManager) EndUndoGrouping() {
	objc.Send[objc.ID](u_.ID, objc.Sel("endUndoGrouping"))
}

// Performs the operations in the last group on the redo stack, if there are any, recording them on the undo stack as a single group.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UndoManager/redo()
func (u_ UndoManager) Redo() {
	objc.Send[objc.ID](u_.ID, objc.Sel("redo"))
}

// Registers the selector of the specified target to implement a single undo operation that the target receives.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UndoManager/registerUndo(withTarget:selector:object:)
func (u_ UndoManager) RegisterUndoWithTargetSelectorObject(target objc.ID, selector objc.SEL, object objc.ID) {
	objc.Send[objc.ID](u_.ID, objc.Sel("registerUndoWithTarget:selector:object:"), target, selector, object)
}

// Clears the undo and redo stacks and reenables the manager.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UndoManager/removeAllActions()
func (u_ UndoManager) RemoveAllActions() {
	objc.Send[objc.ID](u_.ID, objc.Sel("removeAllActions"))
}

// Sets the name of the action associated with the Undo or Redo command.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UndoManager/setActionName(_:)-8lzip
func (u_ UndoManager) SetActionName(actionName string) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setActionName:"), objc.String(actionName))
}

// Sets a user info value for an undo or redo action.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UndoManager/setActionUserInfoValue(_:forKey:)
func (u_ UndoManager) SetActionUserInfoValueForKey(info objc.ID, key unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setActionUserInfoValue:forKey:"), info, key)
}

// Closes the top-level undo group if necessary, and then performs undo operations on the group.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UndoManager/undo()
func (u_ UndoManager) Undo() {
	objc.Send[objc.ID](u_.ID, objc.Sel("undo"))
}

// Retrieves the undo action’s user info value for the given key.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UndoManager/undoActionUserInfoValue(forKey:)
func (u_ UndoManager) UndoActionUserInfoValueForKey(key unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](u_.ID, objc.Sel("undoActionUserInfoValueForKey:"), key)
	return rv
}

// Returns the localized title of the Undo menu command for the identified action.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UndoManager/undoMenuTitle(forUndoActionName:)
func (u_ UndoManager) UndoMenuTitleForUndoActionName(actionName string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("undoMenuTitleForUndoActionName:"), objc.String(actionName))
	return rv
}

// Performs the undo operations in the last undo group (whether top-level or nested), recording the operations on the redo stack as a single group.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UndoManager/undoNestedGroup()
func (u_ UndoManager) UndoNestedGroup() {
	objc.Send[objc.ID](u_.ID, objc.Sel("undoNestedGroup"))
}

// The number of nested undo groups (or redo groups, if redo is the most recent operation) in the current event loop.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UndoManager/groupingLevel
func (u_ UndoManager) GroupingLevel() int {
	rv := objc.Send[int](u_.ID, objc.Sel("groupingLevel"))
	return rv
}


// The name identifying the undo action.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UndoManager/undoActionName
func (u_ UndoManager) UndoActionName() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("undoActionName"))
	return rv
}




