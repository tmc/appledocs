// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UndoManager] class.
var UndoManagerClass = _UndoManagerClass{objc.GetClass("NSUndoManager")}

type _UndoManagerClass struct {
	class objc.Class
}

type UndoManager struct {
	objc.ID
}

func UndoManagerFrom(ptr unsafe.Pointer) UndoManager {
	return UndoManager{
		ID: objc.ID(ptr),
	}
}


// Marks the end of an undo group. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UndoManager/endUndoGrouping()
func (u_ UndoManager) EndUndoGrouping() {
	objc.Send[objc.ID](u_.ID, objc.Sel("endUndoGrouping"))
}
// Performs the operations in the last group on the redo stack, if there are any, recording them on the undo stack as a single group. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UndoManager/redo()
func (u_ UndoManager) Redo() {
	objc.Send[objc.ID](u_.ID, objc.Sel("redo"))
}
// Registers the selector of the specified target to implement a single undo operation that the target receives. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UndoManager/registerUndo(withTarget:selector:object:)
func (u_ UndoManager) RegisterUndoWithTargetSelectorObject(target objc.ID, selector objc.SEL, object objc.ID) {
	objc.Send[objc.ID](u_.ID, objc.Sel("registerUndoWithTarget:selector:object:"), target, selector, object)
}
// Clears the undo and redo stacks and reenables the manager. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UndoManager/removeAllActions()
func (u_ UndoManager) RemoveAllActions() {
	objc.Send[objc.ID](u_.ID, objc.Sel("removeAllActions"))
}
// Sets the name of the action associated with the Undo or Redo command. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UndoManager/setActionName(_:)-8lzip
func (u_ UndoManager) SetActionName(actionName string) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setActionName:"), actionName)
}
// Sets a user info value for an undo or redo action. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UndoManager/setActionUserInfoValue(_:forKey:)
func (u_ UndoManager) SetActionUserInfoValueForKey(info objc.ID, key unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setActionUserInfoValue:forKey:"), info, key)
}
// Closes the top-level undo group if necessary, and then performs undo operations on the group. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UndoManager/undo()
func (u_ UndoManager) Undo() {
	objc.Send[objc.ID](u_.ID, objc.Sel("undo"))
}
// Retrieves the undo action’s user info value for the given key. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UndoManager/undoActionUserInfoValue(forKey:)
func (u_ UndoManager) UndoActionUserInfoValueForKey(key unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](u_.ID, objc.Sel("undoActionUserInfoValueForKey:"), key)
	return rv
}
// Returns the localized title of the Undo menu command for the identified action. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UndoManager/undoMenuTitle(forUndoActionName:)
func (u_ UndoManager) UndoMenuTitleForUndoActionName(actionName string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("undoMenuTitleForUndoActionName:"), actionName)
	return rv
}
// Performs the undo operations in the last undo group (whether top-level or nested), recording the operations on the redo stack as a single group. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UndoManager/undoNestedGroup()
func (u_ UndoManager) UndoNestedGroup() {
	objc.Send[objc.ID](u_.ID, objc.Sel("undoNestedGroup"))
}


