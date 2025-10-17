// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [UndoManager] class.
var UndoManagerClass objc.Class

func init() {
	UndoManagerClass = objc.GetClass("NSUndoManager")
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
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/UndoManager/endUndoGrouping()
func (u_ UndoManager) EndUndoGrouping() {
	sel := objc.RegisterName("endUndoGrouping")
	u_.ID.Send(sel)
}
// Performs the operations in the last group on the redo stack, if there are any, recording them on the undo stack as a single group. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/UndoManager/redo()
func (u_ UndoManager) Redo() {
	sel := objc.RegisterName("redo")
	u_.ID.Send(sel)
}
// Registers the selector of the specified target to implement a single undo operation that the target receives. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/UndoManager/registerUndo(withTarget:selector:object:)
func (u_ UndoManager) RegisterUndoWithTargetSelectorObject(target objc.ID, selector objc.SEL, object objc.ID) {
	sel := objc.RegisterName("registerUndoWithTarget:selector:object:")
	u_.ID.Send(sel, target, selector, object)
}
// Clears the undo and redo stacks and reenables the manager. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/UndoManager/removeAllActions()
func (u_ UndoManager) RemoveAllActions() {
	sel := objc.RegisterName("removeAllActions")
	u_.ID.Send(sel)
}
// Sets the name of the action associated with the Undo or Redo command. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/UndoManager/setActionName(_:)-8lzip
func (u_ UndoManager) SetActionName(actionName string) {
	sel := objc.RegisterName("setActionName:")
	u_.ID.Send(sel, actionName)
}
// Sets a user info value for an undo or redo action. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/UndoManager/setActionUserInfoValue(_:forKey:)
func (u_ UndoManager) SetActionUserInfoValueForKey(info objc.ID, key unsafe.Pointer) {
	sel := objc.RegisterName("setActionUserInfoValue:forKey:")
	u_.ID.Send(sel, info, key)
}
// Closes the top-level undo group if necessary, and then performs undo operations on the group. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/UndoManager/undo()
func (u_ UndoManager) Undo() {
	sel := objc.RegisterName("undo")
	u_.ID.Send(sel)
}
// Retrieves the undo action’s user info value for the given key. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/UndoManager/undoActionUserInfoValue(forKey:)
func (u_ UndoManager) UndoActionUserInfoValueForKey(key unsafe.Pointer) objc.ID {
	sel := objc.RegisterName("undoActionUserInfoValueForKey:")
	ret := u_.ID.Send(sel, key)
	return ret
}
// Returns the localized title of the Undo menu command for the identified action. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/UndoManager/undoMenuTitle(forUndoActionName:)
func (u_ UndoManager) UndoMenuTitleForUndoActionName(actionName string) unsafe.Pointer {
	sel := objc.RegisterName("undoMenuTitleForUndoActionName:")
	ret := u_.ID.Send(sel, actionName)
	return unsafe.Pointer(ret)
}
// Performs the undo operations in the last undo group (whether top-level or nested), recording the operations on the redo stack as a single group. [Full Topic]

//
// [Full Topic]: doc://com.apple.foundation/documentation/Foundation/UndoManager/undoNestedGroup()
func (u_ UndoManager) UndoNestedGroup() {
	sel := objc.RegisterName("undoNestedGroup")
	u_.ID.Send(sel)
}


