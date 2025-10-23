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
	AlwaysUsesMultipleValuesMarker() bool
	SetAlwaysUsesMultipleValuesMarker(value bool)
	ArrangedObjects() TreeNode
	SetArrangedObjects(value TreeNode)
	AvoidsEmptySelection() bool
	SetAvoidsEmptySelection(value bool)
	CanAddChild() bool
	SetCanAddChild(value bool)
	CanInsert() bool
	SetCanInsert(value bool)
	CanInsertChild() bool
	SetCanInsertChild(value bool)
	ChildrenKeyPath() string
	SetChildrenKeyPath(value string)
	Content() unsafe.Pointer
	SetContent(value unsafe.Pointer)
	CountKeyPath() string
	SetCountKeyPath(value string)
	LeafKeyPath() string
	SetLeafKeyPath(value string)
	PreservesSelection() bool
	SetPreservesSelection(value bool)
	SelectedNodes() TreeNode
	SetSelectedNodes(value TreeNode)
	SelectedObjects() unsafe.Pointer
	SetSelectedObjects(value unsafe.Pointer)
	SelectionIndexPath() foundation.IndexPath
	SetSelectionIndexPath(value foundation.IndexPath)
	SelectionIndexPaths() foundation.IndexPath
	SetSelectionIndexPaths(value foundation.IndexPath)
	SelectsInsertedObjects() bool
	SetSelectsInsertedObjects(value bool)
	SortDescriptors() foundation.SortDescriptor
	SetSortDescriptors(value foundation.SortDescriptor)
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
func (t_ TreeController) AlwaysUsesMultipleValuesMarker() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("alwaysUsesMultipleValuesMarker"))
	return rv
}


// A Boolean value that indicates whether the tree controller always returns the multiple values marker when multiple objects are selected, even if the selected items have the same value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/alwaysusesmultiplevaluesmarker
func (t_ TreeController) SetAlwaysUsesMultipleValuesMarker(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAlwaysUsesMultipleValuesMarker:"), value)
}


// The tree controller’s sorted content objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/arrangedobjects
func (t_ TreeController) ArrangedObjects() TreeNode {
	rv := objc.Send[TreeNode](t_.ID, objc.Sel("arrangedObjects"))
	return rv
}


// The tree controller’s sorted content objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/arrangedobjects
func (t_ TreeController) SetArrangedObjects(value TreeNode) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setArrangedObjects:"), value)
}


// A Boolean value that indicates whether the tree controller requires the content array to attempt to maintain a selection at all times, avoiding an empty selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/avoidsemptyselection
func (t_ TreeController) AvoidsEmptySelection() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("avoidsEmptySelection"))
	return rv
}


// A Boolean value that indicates whether the tree controller requires the content array to attempt to maintain a selection at all times, avoiding an empty selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/avoidsemptyselection
func (t_ TreeController) SetAvoidsEmptySelection(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAvoidsEmptySelection:"), value)
}


// A Boolean value that indicates if a child object can be added to the tree controller’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/canaddchild
func (t_ TreeController) CanAddChild() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("canAddChild"))
	return rv
}


// A Boolean value that indicates if a child object can be added to the tree controller’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/canaddchild
func (t_ TreeController) SetCanAddChild(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setCanAddChild:"), value)
}


// A Boolean value that indicates if an object can be inserted into the tree controller’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/caninsert
func (t_ TreeController) CanInsert() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("canInsert"))
	return rv
}


// A Boolean value that indicates if an object can be inserted into the tree controller’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/caninsert
func (t_ TreeController) SetCanInsert(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setCanInsert:"), value)
}


// A Boolean value that indicates if a child object can be inserted into the tree controller’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/caninsertchild
func (t_ TreeController) CanInsertChild() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("canInsertChild"))
	return rv
}


// A Boolean value that indicates if a child object can be inserted into the tree controller’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/caninsertchild
func (t_ TreeController) SetCanInsertChild(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setCanInsertChild:"), value)
}


// The key path used to find the children in the tree controller’s objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/childrenkeypath
func (t_ TreeController) ChildrenKeyPath() string {
	rv := objc.Send[string](t_.ID, objc.Sel("childrenKeyPath"))
	return rv
}


// The key path used to find the children in the tree controller’s objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/childrenkeypath
func (t_ TreeController) SetChildrenKeyPath(value string) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setChildrenKeyPath:"), objc.String(value))
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
func (t_ TreeController) CountKeyPath() string {
	rv := objc.Send[string](t_.ID, objc.Sel("countKeyPath"))
	return rv
}


// The key path used to find the number of children for a node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/countkeypath
func (t_ TreeController) SetCountKeyPath(value string) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setCountKeyPath:"), objc.String(value))
}


// The key path used by the tree controller to determine if a node is a leaf key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/leafkeypath
func (t_ TreeController) LeafKeyPath() string {
	rv := objc.Send[string](t_.ID, objc.Sel("leafKeyPath"))
	return rv
}


// The key path used by the tree controller to determine if a node is a leaf key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/leafkeypath
func (t_ TreeController) SetLeafKeyPath(value string) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLeafKeyPath:"), objc.String(value))
}


// A Boolean value that indicates whether the tree controller will attempt to preserve the current selection when the content changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/preservesselection
func (t_ TreeController) PreservesSelection() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("preservesSelection"))
	return rv
}


// A Boolean value that indicates whether the tree controller will attempt to preserve the current selection when the content changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/preservesselection
func (t_ TreeController) SetPreservesSelection(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPreservesSelection:"), value)
}


// An array containing the tree controller’s selected tree nodes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/selectednodes
func (t_ TreeController) SelectedNodes() TreeNode {
	rv := objc.Send[TreeNode](t_.ID, objc.Sel("selectedNodes"))
	return rv
}


// An array containing the tree controller’s selected tree nodes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/selectednodes
func (t_ TreeController) SetSelectedNodes(value TreeNode) {
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
func (t_ TreeController) SelectionIndexPath() foundation.IndexPath {
	rv := objc.Send[foundation.IndexPath](t_.ID, objc.Sel("selectionIndexPath"))
	return rv
}


// The index path of the first selected object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/selectionindexpath
func (t_ TreeController) SetSelectionIndexPath(value foundation.IndexPath) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSelectionIndexPath:"), value)
}


// An array containing the index paths of the currently selected objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/selectionindexpaths
func (t_ TreeController) SelectionIndexPaths() foundation.IndexPath {
	rv := objc.Send[foundation.IndexPath](t_.ID, objc.Sel("selectionIndexPaths"))
	return rv
}


// An array containing the index paths of the currently selected objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/selectionindexpaths
func (t_ TreeController) SetSelectionIndexPaths(value foundation.IndexPath) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSelectionIndexPaths:"), value)
}


// A Boolean value that indicates whether the tree controller automatically selects objects as they are inserted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/selectsinsertedobjects
func (t_ TreeController) SelectsInsertedObjects() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("selectsInsertedObjects"))
	return rv
}


// A Boolean value that indicates whether the tree controller automatically selects objects as they are inserted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/selectsinsertedobjects
func (t_ TreeController) SetSelectsInsertedObjects(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSelectsInsertedObjects:"), value)
}


// An array containing the sort descriptors used to arrange the tree controller’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/sortdescriptors
func (t_ TreeController) SortDescriptors() foundation.SortDescriptor {
	rv := objc.Send[foundation.SortDescriptor](t_.ID, objc.Sel("sortDescriptors"))
	return rv
}


// An array containing the sort descriptors used to arrange the tree controller’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreecontroller/sortdescriptors
func (t_ TreeController) SetSortDescriptors(value foundation.SortDescriptor) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSortDescriptors:"), value)
}



