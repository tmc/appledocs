// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	AlwaysUsesMultipleValuesMarker() bool /* primitive/slice/pointer. */
	SetAlwaysUsesMultipleValuesMarker(value bool /* primitive/slice/pointer. */)
	ArrangedObjects() objc.IObject /* cross-framework: TreeNode */
	SetArrangedObjects(value objc.IObject /* cross-framework: TreeNode */)
	AvoidsEmptySelection() bool /* primitive/slice/pointer. */
	SetAvoidsEmptySelection(value bool /* primitive/slice/pointer. */)
	CanAddChild() bool /* primitive/slice/pointer. */
	SetCanAddChild(value bool /* primitive/slice/pointer. */)
	CanInsert() bool /* primitive/slice/pointer. */
	SetCanInsert(value bool /* primitive/slice/pointer. */)
	CanInsertChild() bool /* primitive/slice/pointer. */
	SetCanInsertChild(value bool /* primitive/slice/pointer. */)
	ChildrenKeyPath() objc.IObject /* cross-framework: NSString */
	SetChildrenKeyPath(value objc.IObject /* cross-framework: NSString */)
	Content() unsafe.Pointer
	SetContent(value unsafe.Pointer)
	CountKeyPath() objc.IObject /* cross-framework: NSString */
	SetCountKeyPath(value objc.IObject /* cross-framework: NSString */)
	LeafKeyPath() objc.IObject /* cross-framework: NSString */
	SetLeafKeyPath(value objc.IObject /* cross-framework: NSString */)
	PreservesSelection() bool /* primitive/slice/pointer. */
	SetPreservesSelection(value bool /* primitive/slice/pointer. */)
	SelectedNodes() objc.IObject /* cross-framework: TreeNode */
	SetSelectedNodes(value objc.IObject /* cross-framework: TreeNode */)
	SelectedObjects() unsafe.Pointer
	SetSelectedObjects(value unsafe.Pointer)
	SelectionIndexPath() objc.IObject /* cross-framework: IndexPath */
	SetSelectionIndexPath(value objc.IObject /* cross-framework: IndexPath */)
	SelectionIndexPaths() objc.IObject /* cross-framework: IndexPath */
	SetSelectionIndexPaths(value objc.IObject /* cross-framework: IndexPath */)
	SelectsInsertedObjects() bool /* primitive/slice/pointer. */
	SetSelectsInsertedObjects(value bool /* primitive/slice/pointer. */)
	SortDescriptors() objc.IObject /* cross-framework: SortDescriptor */
	SetSortDescriptors(value objc.IObject /* cross-framework: SortDescriptor */)
	// methods:
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

// Alloc allocates a new instance without initialization.
func (tc _TreeControllerClass) Alloc() TreeController {
	rv := objc.Send[TreeController](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// A Boolean value that indicates whether the tree controller always returns the multiple values marker when multiple objects are selected, even if the selected items have the same value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/alwaysusesmultiplevaluesmarker
func (t_ TreeController) AlwaysUsesMultipleValuesMarker() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](t_.ID, objc.Sel("alwaysUsesMultipleValuesMarker"))
	return rv
}


// A Boolean value that indicates whether the tree controller always returns the multiple values marker when multiple objects are selected, even if the selected items have the same value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/alwaysusesmultiplevaluesmarker
func (t_ TreeController) SetAlwaysUsesMultipleValuesMarker(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAlwaysUsesMultipleValuesMarker:"), value)
}


// The tree controller’s sorted content objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/arrangedobjects
func (t_ TreeController) ArrangedObjects() objc.IObject /* cross-framework: TreeNode */ {
	rv := objc.Send[TreeNode](t_.ID, objc.Sel("arrangedObjects"))
	return rv
}


// The tree controller’s sorted content objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/arrangedobjects
func (t_ TreeController) SetArrangedObjects(value objc.IObject /* cross-framework: TreeNode */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setArrangedObjects:"), value)
}


// A Boolean value that indicates whether the tree controller requires the content array to attempt to maintain a selection at all times, avoiding an empty selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/avoidsemptyselection
func (t_ TreeController) AvoidsEmptySelection() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](t_.ID, objc.Sel("avoidsEmptySelection"))
	return rv
}


// A Boolean value that indicates whether the tree controller requires the content array to attempt to maintain a selection at all times, avoiding an empty selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/avoidsemptyselection
func (t_ TreeController) SetAvoidsEmptySelection(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAvoidsEmptySelection:"), value)
}


// A Boolean value that indicates if a child object can be added to the tree controller’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/canaddchild
func (t_ TreeController) CanAddChild() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](t_.ID, objc.Sel("canAddChild"))
	return rv
}


// A Boolean value that indicates if a child object can be added to the tree controller’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/canaddchild
func (t_ TreeController) SetCanAddChild(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setCanAddChild:"), value)
}


// A Boolean value that indicates if an object can be inserted into the tree controller’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/caninsert
func (t_ TreeController) CanInsert() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](t_.ID, objc.Sel("canInsert"))
	return rv
}


// A Boolean value that indicates if an object can be inserted into the tree controller’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/caninsert
func (t_ TreeController) SetCanInsert(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setCanInsert:"), value)
}


// A Boolean value that indicates if a child object can be inserted into the tree controller’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/caninsertchild
func (t_ TreeController) CanInsertChild() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](t_.ID, objc.Sel("canInsertChild"))
	return rv
}


// A Boolean value that indicates if a child object can be inserted into the tree controller’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/caninsertchild
func (t_ TreeController) SetCanInsertChild(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setCanInsertChild:"), value)
}


// The key path used to find the children in the tree controller’s objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/childrenkeypath
func (t_ TreeController) ChildrenKeyPath() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("childrenKeyPath"))
	return rv
}


// The key path used to find the children in the tree controller’s objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/childrenkeypath
func (t_ TreeController) SetChildrenKeyPath(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setChildrenKeyPath:"), value)
}


// The tree controller’s content object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/content
func (t_ TreeController) Content() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("content"))
	return rv
}


// The tree controller’s content object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/content
func (t_ TreeController) SetContent(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setContent:"), value)
}


// The key path used to find the number of children for a node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/countkeypath
func (t_ TreeController) CountKeyPath() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("countKeyPath"))
	return rv
}


// The key path used to find the number of children for a node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/countkeypath
func (t_ TreeController) SetCountKeyPath(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setCountKeyPath:"), value)
}


// The key path used by the tree controller to determine if a node is a leaf key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/leafkeypath
func (t_ TreeController) LeafKeyPath() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("leafKeyPath"))
	return rv
}


// The key path used by the tree controller to determine if a node is a leaf key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/leafkeypath
func (t_ TreeController) SetLeafKeyPath(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLeafKeyPath:"), value)
}


// A Boolean value that indicates whether the tree controller will attempt to preserve the current selection when the content changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/preservesselection
func (t_ TreeController) PreservesSelection() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](t_.ID, objc.Sel("preservesSelection"))
	return rv
}


// A Boolean value that indicates whether the tree controller will attempt to preserve the current selection when the content changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/preservesselection
func (t_ TreeController) SetPreservesSelection(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPreservesSelection:"), value)
}


// An array containing the tree controller’s selected tree nodes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/selectednodes
func (t_ TreeController) SelectedNodes() objc.IObject /* cross-framework: TreeNode */ {
	rv := objc.Send[TreeNode](t_.ID, objc.Sel("selectedNodes"))
	return rv
}


// An array containing the tree controller’s selected tree nodes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/selectednodes
func (t_ TreeController) SetSelectedNodes(value objc.IObject /* cross-framework: TreeNode */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSelectedNodes:"), value)
}


// An array containing the currently selected objects in the tree controller’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/selectedobjects
func (t_ TreeController) SelectedObjects() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("selectedObjects"))
	return rv
}


// An array containing the currently selected objects in the tree controller’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/selectedobjects
func (t_ TreeController) SetSelectedObjects(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSelectedObjects:"), value)
}


// The index path of the first selected object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/selectionindexpath
func (t_ TreeController) SelectionIndexPath() objc.IObject /* cross-framework: IndexPath */ {
	rv := objc.Send[foundation.IndexPath](t_.ID, objc.Sel("selectionIndexPath"))
	return rv
}


// The index path of the first selected object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/selectionindexpath
func (t_ TreeController) SetSelectionIndexPath(value objc.IObject /* cross-framework: IndexPath */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSelectionIndexPath:"), value)
}


// An array containing the index paths of the currently selected objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/selectionindexpaths
func (t_ TreeController) SelectionIndexPaths() objc.IObject /* cross-framework: IndexPath */ {
	rv := objc.Send[foundation.IndexPath](t_.ID, objc.Sel("selectionIndexPaths"))
	return rv
}


// An array containing the index paths of the currently selected objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/selectionindexpaths
func (t_ TreeController) SetSelectionIndexPaths(value objc.IObject /* cross-framework: IndexPath */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSelectionIndexPaths:"), value)
}


// A Boolean value that indicates whether the tree controller automatically selects objects as they are inserted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/selectsinsertedobjects
func (t_ TreeController) SelectsInsertedObjects() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](t_.ID, objc.Sel("selectsInsertedObjects"))
	return rv
}


// A Boolean value that indicates whether the tree controller automatically selects objects as they are inserted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/selectsinsertedobjects
func (t_ TreeController) SetSelectsInsertedObjects(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSelectsInsertedObjects:"), value)
}


// An array containing the sort descriptors used to arrange the tree controller’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/sortdescriptors
func (t_ TreeController) SortDescriptors() objc.IObject /* cross-framework: SortDescriptor */ {
	rv := objc.Send[SortDescriptor](t_.ID, objc.Sel("sortDescriptors"))
	return rv
}


// An array containing the sort descriptors used to arrange the tree controller’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/sortdescriptors
func (t_ TreeController) SetSortDescriptors(value objc.IObject /* cross-framework: SortDescriptor */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSortDescriptors:"), value)
}



