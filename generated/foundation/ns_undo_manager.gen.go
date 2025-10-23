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
	UndoManagerClass     _UndoManagerClass
	UndoManagerClassOnce sync.Once
)

func getUndoManagerClass() _UndoManagerClass {
	UndoManagerClassOnce.Do(func() {
		UndoManagerClass = _UndoManagerClass{objc.GetClass("NSUndoManager")}
	})
	return UndoManagerClass
}

type _UndoManagerClass struct {
	class objc.Class
}

// An interface definition for the [UndoManager] class.
type IUndoManager interface {
	objectivec.IObject
	// properties:
	CanRedo() bool /* primitive/slice/pointer. */
	CanUndo() bool /* primitive/slice/pointer. */
	GroupingLevel() int /* primitive/slice/pointer. */
	GroupsByEvent() bool /* primitive/slice/pointer. */
	SetGroupsByEvent(value bool /* primitive/slice/pointer. */)
	Redoing() bool /* primitive/slice/pointer. */
	UndoRegistrationEnabled() bool /* primitive/slice/pointer. */
	Undoing() bool /* primitive/slice/pointer. */
	LevelsOfUndo() uint /* primitive/slice/pointer. */
	SetLevelsOfUndo(value uint /* primitive/slice/pointer. */)
	RedoActionIsDiscardable() bool /* primitive/slice/pointer. */
	RedoActionName() IString
	RedoCount() uint /* primitive/slice/pointer. */
	RedoMenuItemTitle() IString
	RunLoopModes() []string /* primitive/slice/pointer. */
	SetRunLoopModes(value []string /* primitive/slice/pointer. */)
	UndoActionIsDiscardable() bool /* primitive/slice/pointer. */
	UndoActionName() IString
	UndoCount() uint /* primitive/slice/pointer. */
	UndoMenuItemTitle() IString
	NSUndoCloseGroupingRunLoopOrdering() int /* primitive/slice/pointer. */
	NSUndoManagerGroupIsDiscardableKey() IString
	IsRedoing() bool /* primitive/slice/pointer. */
	SetIsRedoing(value bool /* primitive/slice/pointer. */)
	IsUndoRegistrationEnabled() bool /* primitive/slice/pointer. */
	SetIsUndoRegistrationEnabled(value bool /* primitive/slice/pointer. */)
	IsUndoing() bool /* primitive/slice/pointer. */
	SetIsUndoing(value bool /* primitive/slice/pointer. */)
	// methods:
	RegisterUndoWithTargetHandler(target objectivec.IObject, undoHandler func(unsafe.Pointer) /* not a class type */)
	BeginUndoGrouping()
	DisableUndoRegistration()
	EnableUndoRegistration()
	EndUndoGrouping()
	PrepareWithInvocationTarget(target objectivec.IObject) objc.ID
	Redo()
	RedoActionUserInfoValueForKey(key objc.IObject /* cross-framework UndoManagerUserInfoKey */) objc.ID
	RedoMenuTitleForUndoActionName(actionName IString) IString
	RegisterUndoWithTargetSelectorObject(target objectivec.IObject, selector objc.SEL, object objectivec.IObject)
	RemoveAllActions()
	RemoveAllActionsWithTarget(target objectivec.IObject)
	SetActionIsDiscardable(discardable bool /* primitive/slice/pointer. */)
	SetActionName(actionName IString)
	SetActionUserInfoValueForKey(info objectivec.IObject, key objc.IObject /* cross-framework UndoManagerUserInfoKey */)
	Undo()
	UndoActionUserInfoValueForKey(key objc.IObject /* cross-framework UndoManagerUserInfoKey */) objc.ID
	UndoMenuTitleForUndoActionName(actionName IString) IString
	UndoNestedGroup()
}

// A general-purpose recorder of operations that enables undo and redo.
//
// You register an undo operation by calling one of the methods described in Registering undo operations. You specify the name of the object that’s changing (or the owner of that object) and provide a closure, method, or invocation to revert its state. After you register an undo operation, you can call on the undo manager to revert to the state of the last undo operation. When undoing an action, saves the operations you revert to so that you can call automatically. Typically, apps with UI interactions work with . For example, UIKit implements undo and redo in its text view object, making it easy for you to undo and redo actions in objects along the responder chain. also serves as a general-purpose state manager, which you can use to undo and redo many kinds of actions. For example, an interactive command-line utility can use this class to undo the last command run, or a networking library can undo a request by sending another request that invalidates the previous one.


// A general-purpose recorder of operations that enables undo and redo.
//
// [Full Topic]
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



// Records a single undo operation for a given target so that when the manager performs an undo, it executes the specified block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUndoManager/registerUndoWithTarget:handler:
func (u_ UndoManager) RegisterUndoWithTargetHandler(target objectivec.IObject, undoHandler func(unsafe.Pointer) /* not a class type */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("registerUndoWithTarget:handler:"), target, undoHandler)
}


// Marks the beginning of an undo group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UndoManager/beginUndoGrouping()
func (u_ UndoManager) BeginUndoGrouping() {
	objc.Send[objc.ID](u_.ID, objc.Sel("beginUndoGrouping"))
}


// Disables the recording of undo operations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UndoManager/disableUndoRegistration()
func (u_ UndoManager) DisableUndoRegistration() {
	objc.Send[objc.ID](u_.ID, objc.Sel("disableUndoRegistration"))
}


// Enables the recording of undo operations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UndoManager/enableUndoRegistration()
func (u_ UndoManager) EnableUndoRegistration() {
	objc.Send[objc.ID](u_.ID, objc.Sel("enableUndoRegistration"))
}


// Marks the end of an undo group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UndoManager/endUndoGrouping()
func (u_ UndoManager) EndUndoGrouping() {
	objc.Send[objc.ID](u_.ID, objc.Sel("endUndoGrouping"))
}


// Prepares the undo manager for invocation-based undo with the given target as the subject of the next undo operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UndoManager/prepare(withInvocationTarget:)
func (u_ UndoManager) PrepareWithInvocationTarget(target objectivec.IObject) objc.ID {
	rv := objc.Send[objc.ID](u_.ID, objc.Sel("prepareWithInvocationTarget:"), target)
	return rv
}


// Performs the operations in the last group on the redo stack, if there are any, recording them on the undo stack as a single group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UndoManager/redo()
func (u_ UndoManager) Redo() {
	objc.Send[objc.ID](u_.ID, objc.Sel("redo"))
}


// Retrieves the redo action’s user info value for the given key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UndoManager/redoActionUserInfoValue(forKey:)
func (u_ UndoManager) RedoActionUserInfoValueForKey(key objc.IObject /* cross-framework UndoManagerUserInfoKey */) objc.ID {
	rv := objc.Send[objc.ID](u_.ID, objc.Sel("redoActionUserInfoValueForKey:"), key)
	return rv
}


// Returns the localized title of the Redo menu command for the identified action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UndoManager/redoMenuTitle(forUndoActionName:)
func (u_ UndoManager) RedoMenuTitleForUndoActionName(actionName IString) IString {
	rv := objc.Send[String](u_.ID, objc.Sel("redoMenuTitleForUndoActionName:"), actionName)
	return rv
}


// Registers the selector of the specified target to implement a single undo operation that the target receives.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UndoManager/registerUndo(withTarget:selector:object:)
func (u_ UndoManager) RegisterUndoWithTargetSelectorObject(target objectivec.IObject, selector objc.SEL, object objectivec.IObject) {
	objc.Send[objc.ID](u_.ID, objc.Sel("registerUndoWithTarget:selector:object:"), target, selector, object)
}


// Clears the undo and redo stacks and reenables the manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UndoManager/removeAllActions()
func (u_ UndoManager) RemoveAllActions() {
	objc.Send[objc.ID](u_.ID, objc.Sel("removeAllActions"))
}


// Clears the undo and redo stacks of all operations involving the specified target as the recipient of the undo message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UndoManager/removeAllActions(withTarget:)
func (u_ UndoManager) RemoveAllActionsWithTarget(target objectivec.IObject) {
	objc.Send[objc.ID](u_.ID, objc.Sel("removeAllActionsWithTarget:"), target)
}


// Sets whether the next undo or redo action is discardable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UndoManager/setActionIsDiscardable(_:)
func (u_ UndoManager) SetActionIsDiscardable(discardable bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setActionIsDiscardable:"), discardable)
}


// Sets the name of the action associated with the Undo or Redo command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UndoManager/setActionName(_:)-8lzip
func (u_ UndoManager) SetActionName(actionName IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setActionName:"), actionName)
}


// Sets a user info value for an undo or redo action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UndoManager/setActionUserInfoValue(_:forKey:)
func (u_ UndoManager) SetActionUserInfoValueForKey(info objectivec.IObject, key objc.IObject /* cross-framework UndoManagerUserInfoKey */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setActionUserInfoValue:forKey:"), info, key)
}


// Closes the top-level undo group if necessary, and then performs undo operations on the group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UndoManager/undo()
func (u_ UndoManager) Undo() {
	objc.Send[objc.ID](u_.ID, objc.Sel("undo"))
}


// Retrieves the undo action’s user info value for the given key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UndoManager/undoActionUserInfoValue(forKey:)
func (u_ UndoManager) UndoActionUserInfoValueForKey(key objc.IObject /* cross-framework UndoManagerUserInfoKey */) objc.ID {
	rv := objc.Send[objc.ID](u_.ID, objc.Sel("undoActionUserInfoValueForKey:"), key)
	return rv
}


// Returns the localized title of the Undo menu command for the identified action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UndoManager/undoMenuTitle(forUndoActionName:)
func (u_ UndoManager) UndoMenuTitleForUndoActionName(actionName IString) IString {
	rv := objc.Send[String](u_.ID, objc.Sel("undoMenuTitleForUndoActionName:"), actionName)
	return rv
}


// Performs the undo operations in the last undo group (whether top-level or nested), recording the operations on the redo stack as a single group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UndoManager/undoNestedGroup()
func (u_ UndoManager) UndoNestedGroup() {
	objc.Send[objc.ID](u_.ID, objc.Sel("undoNestedGroup"))
}


// A Boolean value that indicates whether the manager has any actions to redo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UndoManager/canRedo
func (u_ UndoManager) CanRedo() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](u_.ID, objc.Sel("canRedo"))
	return rv
}


// A Boolean value that indicates whether the manager has any actions to undo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UndoManager/canUndo
func (u_ UndoManager) CanUndo() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](u_.ID, objc.Sel("canUndo"))
	return rv
}


// The number of nested undo groups (or redo groups, if redo is the most recent operation) in the current event loop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UndoManager/groupingLevel
func (u_ UndoManager) GroupingLevel() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](u_.ID, objc.Sel("groupingLevel"))
	return rv
}


// A Boolean value that indicates whether the manager automatically creates undo groups around each pass of the run loop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UndoManager/groupsByEvent
func (u_ UndoManager) GroupsByEvent() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](u_.ID, objc.Sel("groupsByEvent"))
	return rv
}


// A Boolean value that indicates whether the manager automatically creates undo groups around each pass of the run loop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UndoManager/groupsByEvent
func (u_ UndoManager) SetGroupsByEvent(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setGroupsByEvent:"), value)
}


// Returns a Boolean value that indicates whether the manager is in the process of performing a redo action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UndoManager/isRedoing
func (u_ UndoManager) Redoing() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](u_.ID, objc.Sel("redoing"))
	return rv
}


// A Boolean value that indicates whether the recording of undo operations is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UndoManager/isUndoRegistrationEnabled
func (u_ UndoManager) UndoRegistrationEnabled() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](u_.ID, objc.Sel("undoRegistrationEnabled"))
	return rv
}


// Returns a Boolean value that indicates whether the manager is in the process of performing an undo action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UndoManager/isUndoing
func (u_ UndoManager) Undoing() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](u_.ID, objc.Sel("undoing"))
	return rv
}


// The maximum number of top-level undo groups the undo manager holds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UndoManager/levelsOfUndo
func (u_ UndoManager) LevelsOfUndo() uint /* primitive/slice/pointer. */ {
	rv := objc.Send[uint](u_.ID, objc.Sel("levelsOfUndo"))
	return rv
}


// The maximum number of top-level undo groups the undo manager holds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UndoManager/levelsOfUndo
func (u_ UndoManager) SetLevelsOfUndo(value uint /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setLevelsOfUndo:"), value)
}


// A Boolean value that indicates whether the next redo action is discardable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UndoManager/redoActionIsDiscardable
func (u_ UndoManager) RedoActionIsDiscardable() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](u_.ID, objc.Sel("redoActionIsDiscardable"))
	return rv
}


// The name identifying the redo action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UndoManager/redoActionName
func (u_ UndoManager) RedoActionName() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("redoActionName"))
	return rv
}


// The number of times you can invoke redo before there are no actions left to redo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UndoManager/redoCount
func (u_ UndoManager) RedoCount() uint /* primitive/slice/pointer. */ {
	rv := objc.Send[uint](u_.ID, objc.Sel("redoCount"))
	return rv
}


// The title of the Redo menu command, such as Redo Paste.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UndoManager/redoMenuItemTitle
func (u_ UndoManager) RedoMenuItemTitle() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("redoMenuItemTitle"))
	return rv
}


// The modes governing the types of input to handle during a cycle of the run loop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UndoManager/runLoopModes
func (u_ UndoManager) RunLoopModes() []string /* primitive/slice/pointer. */ {
	rv := objc.Send[[]string](u_.ID, objc.Sel("runLoopModes"))
	return rv
}


// The modes governing the types of input to handle during a cycle of the run loop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UndoManager/runLoopModes
func (u_ UndoManager) SetRunLoopModes(value []string /* primitive/slice/pointer. */) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](u_.ID, objc.Sel("setRunLoopModes:"), nsArray)
}


// A Boolean value that indicates whether the next undo action is discardable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UndoManager/undoActionIsDiscardable
func (u_ UndoManager) UndoActionIsDiscardable() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](u_.ID, objc.Sel("undoActionIsDiscardable"))
	return rv
}


// The name identifying the undo action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UndoManager/undoActionName
func (u_ UndoManager) UndoActionName() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("undoActionName"))
	return rv
}


// The number of times you can invoke undo before there are no actions left to undo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UndoManager/undoCount
func (u_ UndoManager) UndoCount() uint /* primitive/slice/pointer. */ {
	rv := objc.Send[uint](u_.ID, objc.Sel("undoCount"))
	return rv
}


// The title of the Undo menu command, such as Undo Paste.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UndoManager/undoMenuItemTitle
func (u_ UndoManager) UndoMenuItemTitle() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("undoMenuItemTitle"))
	return rv
}


// A priority to use when using a run loop to close an undo group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsundoclosegroupingrunloopordering
func (u_ UndoManager) NSUndoCloseGroupingRunLoopOrdering() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](u_.ID, objc.Sel("NSUndoCloseGroupingRunLoopOrdering"))
	return rv
}


// A key, used in a notification’s user info, that indicates the undo group contains only discardable actions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsundomanagergroupisdiscardablekey
func (u_ UndoManager) NSUndoManagerGroupIsDiscardableKey() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("NSUndoManagerGroupIsDiscardableKey"))
	return rv
}


// Returns a Boolean value that indicates whether the manager is in the process of performing a redo action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/undomanager/isredoing
func (u_ UndoManager) IsRedoing() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](u_.ID, objc.Sel("isRedoing"))
	return rv
}


// Returns a Boolean value that indicates whether the manager is in the process of performing a redo action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/undomanager/isredoing
func (u_ UndoManager) SetIsRedoing(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setIsRedoing:"), value)
}


// A Boolean value that indicates whether the recording of undo operations is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/undomanager/isundoregistrationenabled
func (u_ UndoManager) IsUndoRegistrationEnabled() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](u_.ID, objc.Sel("isUndoRegistrationEnabled"))
	return rv
}


// A Boolean value that indicates whether the recording of undo operations is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/undomanager/isundoregistrationenabled
func (u_ UndoManager) SetIsUndoRegistrationEnabled(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setIsUndoRegistrationEnabled:"), value)
}


// Returns a Boolean value that indicates whether the manager is in the process of performing an undo action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/undomanager/isundoing
func (u_ UndoManager) IsUndoing() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](u_.ID, objc.Sel("isUndoing"))
	return rv
}


// Returns a Boolean value that indicates whether the manager is in the process of performing an undo action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/undomanager/isundoing
func (u_ UndoManager) SetIsUndoing(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setIsUndoing:"), value)
}



