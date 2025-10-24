// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSTreeController */


/* debug [class_header]: Header for NSTreeController */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TreeController */
// An interface definition for the [TreeController] class.
type ITreeController interface {
	IObjectController
	
/* debug [class_interface_properties]: Properties for TreeController */
	// properties:
	AlwaysUsesMultipleValuesMarker() bool
	SetAlwaysUsesMultipleValuesMarker(value bool)
	ArrangedObjects() ITreeNode
	AvoidsEmptySelection() bool
	SetAvoidsEmptySelection(value bool)
	CanAddChild() bool
	CanInsert() bool
	CanInsertChild() bool
	ChildrenKeyPath() objc.IObject /* cross-framework: NSString */
	SetChildrenKeyPath(value objc.IObject /* cross-framework: NSString */)
	Content() objc.ID
	SetContent(value objc.ID)
	CountKeyPath() objc.IObject /* cross-framework: NSString */
	SetCountKeyPath(value objc.IObject /* cross-framework: NSString */)
	LeafKeyPath() objc.IObject /* cross-framework: NSString */
	SetLeafKeyPath(value objc.IObject /* cross-framework: NSString */)
	PreservesSelection() bool
	SetPreservesSelection(value bool)
	SelectedNodes() []TreeNode
	SelectedObjects() objc.IObject /* cross-framework: NSArray */
	SelectionIndexPath() foundation.IndexPath
	SelectionIndexPaths() []foundation.IndexPath
	SelectsInsertedObjects() bool
	SetSelectsInsertedObjects(value bool)
	SortDescriptors() []objc.IObject
	SetSortDescriptors(value []objc.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TreeController */
	// methods:
	Add(sender objc.IObject)
	AddChild(sender objc.IObject)
	AddSelectionIndexPaths(indexPaths []foundation.IndexPath) bool
	ChildrenKeyPathForNode(node ITreeNode) foundation.String
	CountKeyPathForNode(node ITreeNode) foundation.String
	Insert(sender objc.IObject)
	InsertObjectAtArrangedObjectIndexPath(object objc.IObject, indexPath foundation.IndexPath)
	InsertObjectsAtArrangedObjectIndexPaths(objects objc.IObject /* cross-framework: NSArray */, indexPaths []foundation.IndexPath)
	InsertChild(sender objc.IObject)
	LeafKeyPathForNode(node ITreeNode) foundation.String
	MoveNodesToIndexPath(nodes []TreeNode, startingIndexPath foundation.IndexPath)
	MoveNodeToIndexPath(node ITreeNode, indexPath foundation.IndexPath)
	RearrangeObjects()
	Remove(sender objc.IObject)
	RemoveObjectAtArrangedObjectIndexPath(indexPath foundation.IndexPath)
	RemoveObjectsAtArrangedObjectIndexPaths(indexPaths []foundation.IndexPath)
	RemoveSelectionIndexPaths(indexPaths []foundation.IndexPath) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TreeController */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TreeController */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TreeController *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TreeController */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TreeController */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TreeController */

// Adds an object to the tree controller’s content after the current selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/add(_:)
func (t_ TreeController) Add(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("add:"), sender)
}/* debug [instance_methods/method]: Add */


// Adds a child object to the currently selected item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/addChild(_:)
func (t_ TreeController) AddChild(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("addChild:"), sender)
}/* debug [instance_methods/method]: AddChild */


// Adds the objects at the specified in the tree controller’s content to the current selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/addSelectionIndexPaths(_:)
func (t_ TreeController) AddSelectionIndexPaths(indexPaths []foundation.IndexPath) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("addSelectionIndexPaths:"), indexPaths)
	return rv
}/* debug [instance_methods/method]: AddSelectionIndexPaths */


// Returns the key path used to find the children in the specified tree node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/childrenKeyPath(for:)
func (t_ TreeController) ChildrenKeyPathForNode(node ITreeNode) foundation.String {
	rv := objc.Send[foundation.String](t_.ID, objc.Sel("childrenKeyPathForNode:"), node)
	return rv
}/* debug [instance_methods/method]: ChildrenKeyPathForNode */


// Returns the key path that provides the number of children for a specified node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/countKeyPath(for:)
func (t_ TreeController) CountKeyPathForNode(node ITreeNode) foundation.String {
	rv := objc.Send[foundation.String](t_.ID, objc.Sel("countKeyPathForNode:"), node)
	return rv
}/* debug [instance_methods/method]: CountKeyPathForNode */


// Creates a new object of the class specified by and inserts it into the tree controller’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/insert(_:)
func (t_ TreeController) Insert(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("insert:"), sender)
}/* debug [instance_methods/method]: Insert */


// Inserts into the tree controller’s arranged objects array at the location specified by , and adds it to the tree controller’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/insert(_:atArrangedObjectIndexPath:)
func (t_ TreeController) InsertObjectAtArrangedObjectIndexPath(object objc.IObject, indexPath foundation.IndexPath) {
	objc.Send[objc.ID](t_.ID, objc.Sel("insertObject:atArrangedObjectIndexPath:"), object, indexPath)
}/* debug [instance_methods/method]: InsertObjectAtArrangedObjectIndexPath */


// Inserts into the tree controller’s arranged objects array at the locations specified in , and adds them to the tree controller’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/insert(_:atArrangedObjectIndexPaths:)
func (t_ TreeController) InsertObjectsAtArrangedObjectIndexPaths(objects objc.IObject /* cross-framework: NSArray */, indexPaths []foundation.IndexPath) {
	objc.Send[objc.ID](t_.ID, objc.Sel("insertObjects:atArrangedObjectIndexPaths:"), objects, indexPaths)
}/* debug [instance_methods/method]: InsertObjectsAtArrangedObjectIndexPaths */


// Creates a new object of the class specified by and inserts it into the tree controller’s content as a child of the current selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/insertChild(_:)
func (t_ TreeController) InsertChild(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("insertChild:"), sender)
}/* debug [instance_methods/method]: InsertChild */


// Returns the key path that specifies whether the node is a leaf node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/leafKeyPath(for:)
func (t_ TreeController) LeafKeyPathForNode(node ITreeNode) foundation.String {
	rv := objc.Send[foundation.String](t_.ID, objc.Sel("leafKeyPathForNode:"), node)
	return rv
}/* debug [instance_methods/method]: LeafKeyPathForNode */


// Moves the specified tree nodes to the new index path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/move(_:to:)-moi9
func (t_ TreeController) MoveNodesToIndexPath(nodes []TreeNode, startingIndexPath foundation.IndexPath) {
	objc.Send[objc.ID](t_.ID, objc.Sel("moveNodes:toIndexPath:"), nodes, startingIndexPath)
}/* debug [instance_methods/method]: MoveNodesToIndexPath */


// Moves the specified tree node to the new index path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/move(_:to:)-s5xp
func (t_ TreeController) MoveNodeToIndexPath(node ITreeNode, indexPath foundation.IndexPath) {
	objc.Send[objc.ID](t_.ID, objc.Sel("moveNode:toIndexPath:"), node, indexPath)
}/* debug [instance_methods/method]: MoveNodeToIndexPath */


// Use this method to trigger reordering of the tree controller’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/rearrangeObjects()
func (t_ TreeController) RearrangeObjects() {
	objc.Send[objc.ID](t_.ID, objc.Sel("rearrangeObjects"))
}/* debug [instance_methods/method]: RearrangeObjects */


// Removes the tree controller’s selected objects from the content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/remove(_:)
func (t_ TreeController) Remove(sender objc.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("remove:"), sender)
}/* debug [instance_methods/method]: Remove */


// Removes the object at the specified in the tree controller’s arranged objects from the tree controller’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/removeObject(atArrangedObjectIndexPath:)
func (t_ TreeController) RemoveObjectAtArrangedObjectIndexPath(indexPath foundation.IndexPath) {
	objc.Send[objc.ID](t_.ID, objc.Sel("removeObjectAtArrangedObjectIndexPath:"), indexPath)
}/* debug [instance_methods/method]: RemoveObjectAtArrangedObjectIndexPath */


// Removes the objects at the specified in the tree controller’s arranged objects from the tree controller’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/removeObjects(atArrangedObjectIndexPaths:)
func (t_ TreeController) RemoveObjectsAtArrangedObjectIndexPaths(indexPaths []foundation.IndexPath) {
	objc.Send[objc.ID](t_.ID, objc.Sel("removeObjectsAtArrangedObjectIndexPaths:"), indexPaths)
}/* debug [instance_methods/method]: RemoveObjectsAtArrangedObjectIndexPaths */


// Removes the objects at the specified index paths from the tree controller’s current selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/removeSelectionIndexPaths(_:)
func (t_ TreeController) RemoveSelectionIndexPaths(indexPaths []foundation.IndexPath) bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("removeSelectionIndexPaths:"), indexPaths)
	return rv
}/* debug [instance_methods/method]: RemoveSelectionIndexPaths */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TreeController */

// A Boolean value that indicates whether the tree controller always returns the multiple values marker when multiple objects are selected, even if the selected items have the same value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/alwaysUsesMultipleValuesMarker
func (t_ TreeController) AlwaysUsesMultipleValuesMarker() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("alwaysUsesMultipleValuesMarker"))
	return rv
}/* debug [instance_properties/getter]: alwaysUsesMultipleValuesMarker */


// A Boolean value that indicates whether the tree controller always returns the multiple values marker when multiple objects are selected, even if the selected items have the same value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/alwaysUsesMultipleValuesMarker
func (t_ TreeController) SetAlwaysUsesMultipleValuesMarker(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAlwaysUsesMultipleValuesMarker:"), value)
}/* debug [instance_properties/setter]: alwaysUsesMultipleValuesMarker */


// The tree controller’s sorted content objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/arrangedObjects
func (t_ TreeController) ArrangedObjects() ITreeNode {
	rv := objc.Send[TreeNode](t_.ID, objc.Sel("arrangedObjects"))
	return rv
}/* debug [instance_properties/getter]: arrangedObjects */


// A Boolean value that indicates whether the tree controller requires the content array to attempt to maintain a selection at all times, avoiding an empty selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/avoidsEmptySelection
func (t_ TreeController) AvoidsEmptySelection() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("avoidsEmptySelection"))
	return rv
}/* debug [instance_properties/getter]: avoidsEmptySelection */


// A Boolean value that indicates whether the tree controller requires the content array to attempt to maintain a selection at all times, avoiding an empty selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/avoidsEmptySelection
func (t_ TreeController) SetAvoidsEmptySelection(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAvoidsEmptySelection:"), value)
}/* debug [instance_properties/setter]: avoidsEmptySelection */


// A Boolean value that indicates if a child object can be added to the tree controller’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/canAddChild
func (t_ TreeController) CanAddChild() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("canAddChild"))
	return rv
}/* debug [instance_properties/getter]: canAddChild */


// A Boolean value that indicates if an object can be inserted into the tree controller’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/canInsert
func (t_ TreeController) CanInsert() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("canInsert"))
	return rv
}/* debug [instance_properties/getter]: canInsert */


// A Boolean value that indicates if a child object can be inserted into the tree controller’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/canInsertChild
func (t_ TreeController) CanInsertChild() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("canInsertChild"))
	return rv
}/* debug [instance_properties/getter]: canInsertChild */


// The key path used to find the children in the tree controller’s objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/childrenKeyPath
func (t_ TreeController) ChildrenKeyPath() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("childrenKeyPath"))
	return rv
}/* debug [instance_properties/getter]: childrenKeyPath */


// The key path used to find the children in the tree controller’s objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/childrenKeyPath
func (t_ TreeController) SetChildrenKeyPath(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setChildrenKeyPath:"), value)
}/* debug [instance_properties/setter]: childrenKeyPath */


// The tree controller’s content object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/content
func (t_ TreeController) Content() objc.ID {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("content"))
	return rv
}/* debug [instance_properties/getter]: content */


// The tree controller’s content object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/content
func (t_ TreeController) SetContent(value objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setContent:"), value)
}/* debug [instance_properties/setter]: content */


// The key path used to find the number of children for a node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/countKeyPath
func (t_ TreeController) CountKeyPath() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("countKeyPath"))
	return rv
}/* debug [instance_properties/getter]: countKeyPath */


// The key path used to find the number of children for a node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/countKeyPath
func (t_ TreeController) SetCountKeyPath(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setCountKeyPath:"), value)
}/* debug [instance_properties/setter]: countKeyPath */


// The key path used by the tree controller to determine if a node is a leaf key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/leafKeyPath
func (t_ TreeController) LeafKeyPath() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("leafKeyPath"))
	return rv
}/* debug [instance_properties/getter]: leafKeyPath */


// The key path used by the tree controller to determine if a node is a leaf key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/leafKeyPath
func (t_ TreeController) SetLeafKeyPath(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLeafKeyPath:"), value)
}/* debug [instance_properties/setter]: leafKeyPath */


// A Boolean value that indicates whether the tree controller will attempt to preserve the current selection when the content changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/preservesSelection
func (t_ TreeController) PreservesSelection() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("preservesSelection"))
	return rv
}/* debug [instance_properties/getter]: preservesSelection */


// A Boolean value that indicates whether the tree controller will attempt to preserve the current selection when the content changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/preservesSelection
func (t_ TreeController) SetPreservesSelection(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPreservesSelection:"), value)
}/* debug [instance_properties/setter]: preservesSelection */


// An array containing the tree controller’s selected tree nodes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/selectedNodes
func (t_ TreeController) SelectedNodes() []TreeNode {
	rv := objc.Send[[]TreeNode](t_.ID, objc.Sel("selectedNodes"))
	return rv
}/* debug [instance_properties/getter]: selectedNodes */


// An array containing the currently selected objects in the tree controller’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/selectedObjects
func (t_ TreeController) SelectedObjects() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](t_.ID, objc.Sel("selectedObjects"))
	return rv
}/* debug [instance_properties/getter]: selectedObjects */


// The index path of the first selected object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/selectionIndexPath
func (t_ TreeController) SelectionIndexPath() foundation.IndexPath {
	rv := objc.Send[foundation.IndexPath](t_.ID, objc.Sel("selectionIndexPath"))
	return rv
}/* debug [instance_properties/getter]: selectionIndexPath */


// An array containing the index paths of the currently selected objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/selectionIndexPaths
func (t_ TreeController) SelectionIndexPaths() []foundation.IndexPath {
	rv := objc.Send[[]foundation.IndexPath](t_.ID, objc.Sel("selectionIndexPaths"))
	return rv
}/* debug [instance_properties/getter]: selectionIndexPaths */


// A Boolean value that indicates whether the tree controller automatically selects objects as they are inserted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/selectsInsertedObjects
func (t_ TreeController) SelectsInsertedObjects() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("selectsInsertedObjects"))
	return rv
}/* debug [instance_properties/getter]: selectsInsertedObjects */


// A Boolean value that indicates whether the tree controller automatically selects objects as they are inserted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/selectsInsertedObjects
func (t_ TreeController) SetSelectsInsertedObjects(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSelectsInsertedObjects:"), value)
}/* debug [instance_properties/setter]: selectsInsertedObjects */


// An array containing the sort descriptors used to arrange the tree controller’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/sortDescriptors
func (t_ TreeController) SortDescriptors() []objc.IObject {
	rv := objc.Send[[]objc.ID](t_.ID, objc.Sel("sortDescriptors"))
	return rv
}/* debug [instance_properties/getter]: sortDescriptors */


// An array containing the sort descriptors used to arrange the tree controller’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeController/sortDescriptors
func (t_ TreeController) SetSortDescriptors(value []objc.IObject) {
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
}/* debug [instance_properties/setter]: sortDescriptors */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSTreeController */



