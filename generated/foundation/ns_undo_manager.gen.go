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
	NSUndoCloseGroupingRunLoopOrdering() int
	NSUndoManagerGroupIsDiscardableKey() IString
	CanRedo() bool
	SetCanRedo(value bool)
	CanUndo() bool
	SetCanUndo(value bool)
	GroupingLevel() int
	SetGroupingLevel(value int)
	GroupsByEvent() bool
	SetGroupsByEvent(value bool)
	IsRedoing() bool
	SetIsRedoing(value bool)
	IsUndoRegistrationEnabled() bool
	SetIsUndoRegistrationEnabled(value bool)
	IsUndoing() bool
	SetIsUndoing(value bool)
	LevelsOfUndo() int
	SetLevelsOfUndo(value int)
	RedoActionIsDiscardable() bool
	SetRedoActionIsDiscardable(value bool)
	RedoActionName() IString
	SetRedoActionName(value IString)
	RedoCount() int
	SetRedoCount(value int)
	RedoMenuItemTitle() IString
	SetRedoMenuItemTitle(value IString)
	RunLoopModes() unsafe.Pointer
	SetRunLoopModes(value unsafe.Pointer)
	UndoActionIsDiscardable() bool
	SetUndoActionIsDiscardable(value bool)
	UndoActionName() IString
	SetUndoActionName(value IString)
	UndoCount() int
	SetUndoCount(value int)
	UndoMenuItemTitle() IString
	SetUndoMenuItemTitle(value IString)
	// methods:
	Undo()
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



// Closes the top-level undo group if necessary, and then performs undo operations on the group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UndoManager/undo()
func (u_ UndoManager) Undo() {
	objc.Send[objc.ID](u_.ID, objc.Sel("undo"))
}


// A priority to use when using a run loop to close an undo group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsundoclosegroupingrunloopordering
func (u_ UndoManager) NSUndoCloseGroupingRunLoopOrdering() int {
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


// A Boolean value that indicates whether the manager has any actions to redo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/undomanager/canredo
func (u_ UndoManager) CanRedo() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("canRedo"))
	return rv
}


// A Boolean value that indicates whether the manager has any actions to redo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/undomanager/canredo
func (u_ UndoManager) SetCanRedo(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setCanRedo:"), value)
}


// A Boolean value that indicates whether the manager has any actions to undo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/undomanager/canundo
func (u_ UndoManager) CanUndo() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("canUndo"))
	return rv
}


// A Boolean value that indicates whether the manager has any actions to undo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/undomanager/canundo
func (u_ UndoManager) SetCanUndo(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setCanUndo:"), value)
}


// The number of nested undo groups (or redo groups, if redo is the most recent operation) in the current event loop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/undomanager/groupinglevel
func (u_ UndoManager) GroupingLevel() int {
	rv := objc.Send[int](u_.ID, objc.Sel("groupingLevel"))
	return rv
}


// The number of nested undo groups (or redo groups, if redo is the most recent operation) in the current event loop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/undomanager/groupinglevel
func (u_ UndoManager) SetGroupingLevel(value int) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setGroupingLevel:"), value)
}


// A Boolean value that indicates whether the manager automatically creates undo groups around each pass of the run loop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/undomanager/groupsbyevent
func (u_ UndoManager) GroupsByEvent() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("groupsByEvent"))
	return rv
}


// A Boolean value that indicates whether the manager automatically creates undo groups around each pass of the run loop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/undomanager/groupsbyevent
func (u_ UndoManager) SetGroupsByEvent(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setGroupsByEvent:"), value)
}


// Returns a Boolean value that indicates whether the manager is in the process of performing a redo action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/undomanager/isredoing
func (u_ UndoManager) IsRedoing() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("isRedoing"))
	return rv
}


// Returns a Boolean value that indicates whether the manager is in the process of performing a redo action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/undomanager/isredoing
func (u_ UndoManager) SetIsRedoing(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setIsRedoing:"), value)
}


// A Boolean value that indicates whether the recording of undo operations is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/undomanager/isundoregistrationenabled
func (u_ UndoManager) IsUndoRegistrationEnabled() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("isUndoRegistrationEnabled"))
	return rv
}


// A Boolean value that indicates whether the recording of undo operations is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/undomanager/isundoregistrationenabled
func (u_ UndoManager) SetIsUndoRegistrationEnabled(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setIsUndoRegistrationEnabled:"), value)
}


// Returns a Boolean value that indicates whether the manager is in the process of performing an undo action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/undomanager/isundoing
func (u_ UndoManager) IsUndoing() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("isUndoing"))
	return rv
}


// Returns a Boolean value that indicates whether the manager is in the process of performing an undo action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/undomanager/isundoing
func (u_ UndoManager) SetIsUndoing(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setIsUndoing:"), value)
}


// The maximum number of top-level undo groups the undo manager holds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/undomanager/levelsofundo
func (u_ UndoManager) LevelsOfUndo() int {
	rv := objc.Send[int](u_.ID, objc.Sel("levelsOfUndo"))
	return rv
}


// The maximum number of top-level undo groups the undo manager holds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/undomanager/levelsofundo
func (u_ UndoManager) SetLevelsOfUndo(value int) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setLevelsOfUndo:"), value)
}


// A Boolean value that indicates whether the next redo action is discardable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/undomanager/redoactionisdiscardable
func (u_ UndoManager) RedoActionIsDiscardable() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("redoActionIsDiscardable"))
	return rv
}


// A Boolean value that indicates whether the next redo action is discardable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/undomanager/redoactionisdiscardable
func (u_ UndoManager) SetRedoActionIsDiscardable(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setRedoActionIsDiscardable:"), value)
}


// The name identifying the redo action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/undomanager/redoactionname
func (u_ UndoManager) RedoActionName() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("redoActionName"))
	return rv
}


// The name identifying the redo action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/undomanager/redoactionname
func (u_ UndoManager) SetRedoActionName(value IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setRedoActionName:"), value)
}


// The number of times you can invoke redo before there are no actions left to redo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/undomanager/redocount
func (u_ UndoManager) RedoCount() int {
	rv := objc.Send[int](u_.ID, objc.Sel("redoCount"))
	return rv
}


// The number of times you can invoke redo before there are no actions left to redo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/undomanager/redocount
func (u_ UndoManager) SetRedoCount(value int) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setRedoCount:"), value)
}


// The title of the Redo menu command, such as Redo Paste.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/undomanager/redomenuitemtitle
func (u_ UndoManager) RedoMenuItemTitle() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("redoMenuItemTitle"))
	return rv
}


// The title of the Redo menu command, such as Redo Paste.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/undomanager/redomenuitemtitle
func (u_ UndoManager) SetRedoMenuItemTitle(value IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setRedoMenuItemTitle:"), value)
}


// The modes governing the types of input to handle during a cycle of the run loop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/undomanager/runloopmodes
func (u_ UndoManager) RunLoopModes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("runLoopModes"))
	return rv
}


// The modes governing the types of input to handle during a cycle of the run loop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/undomanager/runloopmodes
func (u_ UndoManager) SetRunLoopModes(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setRunLoopModes:"), value)
}


// A Boolean value that indicates whether the next undo action is discardable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/undomanager/undoactionisdiscardable
func (u_ UndoManager) UndoActionIsDiscardable() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("undoActionIsDiscardable"))
	return rv
}


// A Boolean value that indicates whether the next undo action is discardable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/undomanager/undoactionisdiscardable
func (u_ UndoManager) SetUndoActionIsDiscardable(value bool) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setUndoActionIsDiscardable:"), value)
}


// The name identifying the undo action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/undomanager/undoactionname
func (u_ UndoManager) UndoActionName() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("undoActionName"))
	return rv
}


// The name identifying the undo action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/undomanager/undoactionname
func (u_ UndoManager) SetUndoActionName(value IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setUndoActionName:"), value)
}


// The number of times you can invoke undo before there are no actions left to undo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/undomanager/undocount
func (u_ UndoManager) UndoCount() int {
	rv := objc.Send[int](u_.ID, objc.Sel("undoCount"))
	return rv
}


// The number of times you can invoke undo before there are no actions left to undo.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/undomanager/undocount
func (u_ UndoManager) SetUndoCount(value int) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setUndoCount:"), value)
}


// The title of the Undo menu command, such as Undo Paste.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/undomanager/undomenuitemtitle
func (u_ UndoManager) UndoMenuItemTitle() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("undoMenuItemTitle"))
	return rv
}


// The title of the Undo menu command, such as Undo Paste.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/undomanager/undomenuitemtitle
func (u_ UndoManager) SetUndoMenuItemTitle(value IString) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setUndoMenuItemTitle:"), value)
}



