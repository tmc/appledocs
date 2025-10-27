// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [TreeController] class.
var (
	TreeControllerClass     _TreeControllerClass
	TreeControllerClassOnce sync.Once
)

func getTreeControllerClass() _TreeControllerClass {
	TreeControllerClassOnce.Do(func() {
		TreeControllerClass = _TreeControllerClass{objc.GetClass("NSTreeController")}
	})
	return TreeControllerClass
}

type _TreeControllerClass struct {
	class objc.Class
}





// An interface definition for the [TreeController] class.
type ITreeController interface {
	IObjectController
	

	// properties:
	AlwaysUsesMultipleValuesMarker() bool
	SetAlwaysUsesMultipleValuesMarker(value bool)
	ArrangedObjects() ITreeNode
	AvoidsEmptySelection() bool
	SetAvoidsEmptySelection(value bool)
	CanAddChild() bool
	CanInsert() bool
	CanInsertChild() bool
	ChildrenKeyPath() foundation.foundation.INSString
	SetChildrenKeyPath(value foundation.foundation.INSString)
	Content() objc.ID
	SetContent(value objc.ID)
	CountKeyPath() foundation.foundation.INSString
	SetCountKeyPath(value foundation.foundation.INSString)
	LeafKeyPath() foundation.foundation.INSString
	SetLeafKeyPath(value foundation.foundation.INSString)
	PreservesSelection() bool
	SetPreservesSelection(value bool)
	SelectedNodes() []TreeNode
	SelectedObjects() foundation.foundation.INSArray
	SelectionIndexPath() foundation.foundation.INSIndexPath
	SelectionIndexPaths() []foundation.IndexPath
	SelectsInsertedObjects() bool
	SetSelectsInsertedObjects(value bool)
	SortDescriptors() []foundation.SortDescriptor
	SetSortDescriptors(value []foundation.SortDescriptor)


	

	// methods:
	Add(sender objectivec.IObject)
	AddChild(sender objectivec.IObject)
	AddSelectionIndexPaths(indexPaths []foundation.IndexPath) bool
	ChildrenKeyPathForNode(node ITreeNode) foundation.String
	CountKeyPathForNode(node ITreeNode) foundation.String
	Insert(sender objectivec.IObject)
	InsertObjectAtArrangedObjectIndexPath(object objectivec.IObject, indexPath foundation.foundation.INSIndexPath)
	InsertObjectsAtArrangedObjectIndexPaths(objects foundation.foundation.INSArray, indexPaths []foundation.IndexPath)
	InsertChild(sender objectivec.IObject)
	LeafKeyPathForNode(node ITreeNode) foundation.String
	MoveNodesToIndexPath(nodes []TreeNode, startingIndexPath foundation.foundation.INSIndexPath)
	MoveNodeToIndexPath(node ITreeNode, indexPath foundation.foundation.INSIndexPath)
	RearrangeObjects()
	Remove(sender objectivec.IObject)
	RemoveObjectAtArrangedObjectIndexPath(indexPath foundation.foundation.INSIndexPath)
	RemoveObjectsAtArrangedObjectIndexPaths(indexPaths []foundation.IndexPath)
	RemoveSelectionIndexPaths(indexPaths []foundation.IndexPath) bool


}





// Alloc allocates a new instance without initialization.
func (tc _TreeControllerClass) Alloc() TreeController {
	rv := objc.Send[TreeController](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TreeControllerClass) New() TreeController {
	rv := objc.Send[TreeController](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TreeController) Init() TreeController {
	rv := objc.Send[TreeController](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TreeController) Autorelease() TreeController {
	rv := objc.Send[TreeController](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTreeController creates a new TreeController instance.
func NewTreeController() TreeController {
	return getTreeControllerClass().New()
}





// A bindings-compatible controller that manages a tree of objects.
//
// The class provides selection and sort management. Its primary purpose is to act as the controller when binding and instances to a hierarchical collection of objects. The root content object of the tree can be a single object, or an array of objects. An object requires that you describe how the tree of objects is traversed by specifying the key-path for child objects specified by . All child objects for the tree must be key-value coding compliant for the same child key path. If necessary, you should add properties to your model classes that map the child key name to the appropriate class-specific property name. Child objects can implement a count method (specified to the tree controller using ) that, if provided, returns the number of child objects available. Your model objects are expected to update the value of the count key path in a key-value observing compliant method. Optionally, you can also provide a leaf key path using that specifies a key in your model object that returns if the object is a leaf node, and if it is not. Changes to the leaf node value of the child object should be made in a key-value observing compliant manner. Providing the leaf node key path can improve performance, because it prevents the from having to examine the child object to determine if it is a leaf node. For more information about using NSTreeController in your app, see .


// A bindings-compatible controller that manages a tree of objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController
type TreeController struct {
	ObjectController
}

// TreeControllerFrom constructs a [TreeController] from an unsafe.Pointer.
//
// A bindings-compatible controller that manages a tree of objects.
func TreeControllerFrom(ptr unsafe.Pointer) TreeController {
	return TreeController{
		ObjectController: ObjectControllerFrom(ptr),
	}
}




















// Adds an object to the tree controller’s content after the current selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/add(_:)
func (t_ TreeController) Add(sender objectivec.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("add:"), sender)
}


// Adds a child object to the currently selected item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/addChild(_:)
func (t_ TreeController) AddChild(sender objectivec.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("addChild:"), sender)
}


// Adds the objects at the specified in the tree controller’s content to the current selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/addSelectionIndexPaths(_:)
func (t_ TreeController) AddSelectionIndexPaths(indexPaths []foundation.IndexPath) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("addSelectionIndexPaths:"), indexPaths)
	return rv
}


// Returns the key path used to find the children in the specified tree node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/childrenKeyPath(for:)
func (t_ TreeController) ChildrenKeyPathForNode(node ITreeNode) foundation.String {
	rv := objc.Send[foundation.String](t_.ID, objc.Sel("childrenKeyPathForNode:"), node)
	return rv
}


// Returns the key path that provides the number of children for a specified node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/countKeyPath(for:)
func (t_ TreeController) CountKeyPathForNode(node ITreeNode) foundation.String {
	rv := objc.Send[foundation.String](t_.ID, objc.Sel("countKeyPathForNode:"), node)
	return rv
}


// Creates a new object of the class specified by and inserts it into the tree controller’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/insert(_:)
func (t_ TreeController) Insert(sender objectivec.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("insert:"), sender)
}


// Inserts into the tree controller’s arranged objects array at the location specified by , and adds it to the tree controller’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/insert(_:atArrangedObjectIndexPath:)
func (t_ TreeController) InsertObjectAtArrangedObjectIndexPath(object objectivec.IObject, indexPath foundation.foundation.INSIndexPath) {
	objc.Send[objc.ID](t_.ID, objc.Sel("insertObject:atArrangedObjectIndexPath:"), object, indexPath)
}


// Inserts into the tree controller’s arranged objects array at the locations specified in , and adds them to the tree controller’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/insert(_:atArrangedObjectIndexPaths:)
func (t_ TreeController) InsertObjectsAtArrangedObjectIndexPaths(objects foundation.foundation.INSArray, indexPaths []foundation.IndexPath) {
	objc.Send[objc.ID](t_.ID, objc.Sel("insertObjects:atArrangedObjectIndexPaths:"), objects, indexPaths)
}


// Creates a new object of the class specified by and inserts it into the tree controller’s content as a child of the current selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/insertChild(_:)
func (t_ TreeController) InsertChild(sender objectivec.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("insertChild:"), sender)
}


// Returns the key path that specifies whether the node is a leaf node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/leafKeyPath(for:)
func (t_ TreeController) LeafKeyPathForNode(node ITreeNode) foundation.String {
	rv := objc.Send[foundation.String](t_.ID, objc.Sel("leafKeyPathForNode:"), node)
	return rv
}


// Moves the specified tree nodes to the new index path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/move(_:to:)-moi9
func (t_ TreeController) MoveNodesToIndexPath(nodes []TreeNode, startingIndexPath foundation.foundation.INSIndexPath) {
	objc.Send[objc.ID](t_.ID, objc.Sel("moveNodes:toIndexPath:"), nodes, startingIndexPath)
}


// Moves the specified tree node to the new index path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/move(_:to:)-s5xp
func (t_ TreeController) MoveNodeToIndexPath(node ITreeNode, indexPath foundation.foundation.INSIndexPath) {
	objc.Send[objc.ID](t_.ID, objc.Sel("moveNode:toIndexPath:"), node, indexPath)
}


// Use this method to trigger reordering of the tree controller’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/rearrangeObjects()
func (t_ TreeController) RearrangeObjects() {
	objc.Send[objc.ID](t_.ID, objc.Sel("rearrangeObjects"))
}


// Removes the tree controller’s selected objects from the content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/remove(_:)
func (t_ TreeController) Remove(sender objectivec.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("remove:"), sender)
}


// Removes the object at the specified in the tree controller’s arranged objects from the tree controller’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/removeObject(atArrangedObjectIndexPath:)
func (t_ TreeController) RemoveObjectAtArrangedObjectIndexPath(indexPath foundation.foundation.INSIndexPath) {
	objc.Send[objc.ID](t_.ID, objc.Sel("removeObjectAtArrangedObjectIndexPath:"), indexPath)
}


// Removes the objects at the specified in the tree controller’s arranged objects from the tree controller’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/removeObjects(atArrangedObjectIndexPaths:)
func (t_ TreeController) RemoveObjectsAtArrangedObjectIndexPaths(indexPaths []foundation.IndexPath) {
	objc.Send[objc.ID](t_.ID, objc.Sel("removeObjectsAtArrangedObjectIndexPaths:"), indexPaths)
}


// Removes the objects at the specified index paths from the tree controller’s current selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/removeSelectionIndexPaths(_:)
func (t_ TreeController) RemoveSelectionIndexPaths(indexPaths []foundation.IndexPath) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("removeSelectionIndexPaths:"), indexPaths)
	return rv
}







// A Boolean value that indicates whether the tree controller always returns the multiple values marker when multiple objects are selected, even if the selected items have the same value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/alwaysUsesMultipleValuesMarker
func (t_ TreeController) AlwaysUsesMultipleValuesMarker() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("alwaysUsesMultipleValuesMarker"))
	return rv
}


// A Boolean value that indicates whether the tree controller always returns the multiple values marker when multiple objects are selected, even if the selected items have the same value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/alwaysUsesMultipleValuesMarker
func (t_ TreeController) SetAlwaysUsesMultipleValuesMarker(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAlwaysUsesMultipleValuesMarker:"), value)
}


// The tree controller’s sorted content objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/arrangedObjects
func (t_ TreeController) ArrangedObjects() ITreeNode {
	rv := objc.Send[TreeNode](t_.ID, objc.Sel("arrangedObjects"))
	return rv
}


// A Boolean value that indicates whether the tree controller requires the content array to attempt to maintain a selection at all times, avoiding an empty selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/avoidsEmptySelection
func (t_ TreeController) AvoidsEmptySelection() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("avoidsEmptySelection"))
	return rv
}


// A Boolean value that indicates whether the tree controller requires the content array to attempt to maintain a selection at all times, avoiding an empty selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/avoidsEmptySelection
func (t_ TreeController) SetAvoidsEmptySelection(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAvoidsEmptySelection:"), value)
}


// A Boolean value that indicates if a child object can be added to the tree controller’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/canAddChild
func (t_ TreeController) CanAddChild() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("canAddChild"))
	return rv
}


// A Boolean value that indicates if an object can be inserted into the tree controller’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/canInsert
func (t_ TreeController) CanInsert() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("canInsert"))
	return rv
}


// A Boolean value that indicates if a child object can be inserted into the tree controller’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/canInsertChild
func (t_ TreeController) CanInsertChild() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("canInsertChild"))
	return rv
}


// The key path used to find the children in the tree controller’s objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/childrenKeyPath
func (t_ TreeController) ChildrenKeyPath() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("childrenKeyPath"))
	return rv
}


// The key path used to find the children in the tree controller’s objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/childrenKeyPath
func (t_ TreeController) SetChildrenKeyPath(value foundation.foundation.INSString) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setChildrenKeyPath:"), value)
}


// The tree controller’s content object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/content
func (t_ TreeController) Content() objc.ID {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("content"))
	return rv
}


// The tree controller’s content object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/content
func (t_ TreeController) SetContent(value objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setContent:"), value)
}


// The key path used to find the number of children for a node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/countKeyPath
func (t_ TreeController) CountKeyPath() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("countKeyPath"))
	return rv
}


// The key path used to find the number of children for a node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/countKeyPath
func (t_ TreeController) SetCountKeyPath(value foundation.foundation.INSString) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setCountKeyPath:"), value)
}


// The key path used by the tree controller to determine if a node is a leaf key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/leafKeyPath
func (t_ TreeController) LeafKeyPath() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("leafKeyPath"))
	return rv
}


// The key path used by the tree controller to determine if a node is a leaf key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/leafKeyPath
func (t_ TreeController) SetLeafKeyPath(value foundation.foundation.INSString) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLeafKeyPath:"), value)
}


// A Boolean value that indicates whether the tree controller will attempt to preserve the current selection when the content changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/preservesSelection
func (t_ TreeController) PreservesSelection() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("preservesSelection"))
	return rv
}


// A Boolean value that indicates whether the tree controller will attempt to preserve the current selection when the content changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/preservesSelection
func (t_ TreeController) SetPreservesSelection(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPreservesSelection:"), value)
}


// An array containing the tree controller’s selected tree nodes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/selectedNodes
func (t_ TreeController) SelectedNodes() []TreeNode {
	rv := objc.Send[[]TreeNode](t_.ID, objc.Sel("selectedNodes"))
	return rv
}


// An array containing the currently selected objects in the tree controller’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/selectedObjects
func (t_ TreeController) SelectedObjects() foundation.foundation.INSArray {
	rv := objc.Send[foundation.NSArray](t_.ID, objc.Sel("selectedObjects"))
	return rv
}


// The index path of the first selected object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/selectionIndexPath
func (t_ TreeController) SelectionIndexPath() foundation.foundation.INSIndexPath {
	rv := objc.Send[foundation.NSIndexPath](t_.ID, objc.Sel("selectionIndexPath"))
	return rv
}


// An array containing the index paths of the currently selected objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/selectionIndexPaths
func (t_ TreeController) SelectionIndexPaths() []foundation.IndexPath {
	rv := objc.Send[[]foundation.IndexPath](t_.ID, objc.Sel("selectionIndexPaths"))
	return rv
}


// A Boolean value that indicates whether the tree controller automatically selects objects as they are inserted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/selectsInsertedObjects
func (t_ TreeController) SelectsInsertedObjects() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("selectsInsertedObjects"))
	return rv
}


// A Boolean value that indicates whether the tree controller automatically selects objects as they are inserted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/selectsInsertedObjects
func (t_ TreeController) SetSelectsInsertedObjects(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSelectsInsertedObjects:"), value)
}


// An array containing the sort descriptors used to arrange the tree controller’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/sortDescriptors
func (t_ TreeController) SortDescriptors() []foundation.SortDescriptor {
	rv := objc.Send[[]foundation.SortDescriptor](t_.ID, objc.Sel("sortDescriptors"))
	return rv
}


// An array containing the sort descriptors used to arrange the tree controller’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/sortDescriptors
func (t_ TreeController) SetSortDescriptors(value []foundation.SortDescriptor) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](t_.ID, objc.Sel("setSortDescriptors:"), nsArray)
}








