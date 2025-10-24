// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSTreeNode */


/* debug [class_header]: Header for NSTreeNode */
// The class instance for the [TreeNode] class.
var (
	TreeNodeClass     _TreeNodeClass
	TreeNodeClassOnce sync.Once
)

func getTreeNodeClass() _TreeNodeClass {
	TreeNodeClassOnce.Do(func() {
		TreeNodeClass = _TreeNodeClass{objc.GetClass("NSTreeNode")}
	})
	return TreeNodeClass
}

type _TreeNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TreeNode */
// An interface definition for the [TreeNode] class.
type ITreeNode interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for TreeNode */
	// properties:
	ChildNodes() []TreeNode
	IndexPath() foundation.IndexPath
	Leaf() bool
	MutableChildNodes() unsafe.Pointer
	ParentNode() ITreeNode
	RepresentedObject() objc.ID
	Children() ITreeNode
	SetChildren(value ITreeNode)
	IsLeaf() bool
	SetIsLeaf(value bool)
	MutableChildren() foundation.MutableArray
	SetMutableChildren(value foundation.MutableArray)
	Parent() ITreeNode
	SetParent(value ITreeNode)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TreeNode */
	// methods:
	DescendantNodeAtIndexPath(indexPath foundation.IndexPath) ITreeNode
	SortWithSortDescriptorsRecursively(sortDescriptors []objc.IObject, recursively bool)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TreeNode */
// Alloc allocates a new instance without initialization.
func (tc _TreeNodeClass) Alloc() TreeNode {
	rv := objc.Send[TreeNode](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TreeNodeClass) New() TreeNode {
	rv := objc.Send[TreeNode](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TreeNode) Init() TreeNode {
	rv := objc.Send[TreeNode](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TreeNode) Autorelease() TreeNode {
	rv := objc.Send[TreeNode](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTreeNode creates a new TreeNode instance.
func NewTreeNode() TreeNode {
	return getTreeNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TreeNode */
// A node in a tree of nodes.
//
// simplifies the creation and management of trees of objects. Each tree node represents a model object. A tree node with as its parent node is considered the root of the tree.


// A node in a tree of nodes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeNode
type TreeNode struct {
	objectivec.Object
}

// TreeNodeFrom constructs a [TreeNode] from an unsafe.Pointer.
//
// A node in a tree of nodes.
func TreeNodeFrom(ptr unsafe.Pointer) TreeNode {
	return TreeNode{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TreeNode */

// Initializes a newly allocated tree node that represents the specified object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeNode/init(representedObject:)
func NewTreeNodeWithRepresentedObject(modelObject objc.IObject) TreeNode {
	instance := getTreeNodeClass().Alloc()
	rv := objc.Send[TreeNode](instance.ID, objc.Sel("initWithRepresentedObject:"), modelObject)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewTreeNodeWithRepresentedObject */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TreeNode */

// Creates and returns a tree node that represents the specified object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeNode/treeNodeWithRepresentedObject:
func (tc _TreeNodeClass) TreeNodeWithRepresentedObject(modelObject objc.IObject) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(tc.class), objc.Sel("treeNodeWithRepresentedObject:"), modelObject)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=TreeNodeWithRepresentedObject) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TreeNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TreeNode */

// Returns the receiver’s descendant at the specified index path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeNode/descendant(at:)
func (t_ TreeNode) DescendantNodeAtIndexPath(indexPath foundation.IndexPath) ITreeNode {
	rv := objc.Send[TreeNode](t_.ID, objc.Sel("descendantNodeAtIndexPath:"), indexPath)
	return rv
}/* debug [instance_methods/method]: DescendantNodeAtIndexPath */


// Sorts the receiver’s subtree using the values of the represented objects with the specified sort descriptors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeNode/sort(with:recursively:)
func (t_ TreeNode) SortWithSortDescriptorsRecursively(sortDescriptors []objc.IObject, recursively bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("sortWithSortDescriptors:recursively:"), sortDescriptors, recursively)
}/* debug [instance_methods/method]: SortWithSortDescriptorsRecursively */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TreeNode */

// An array containing receiver’s child nodes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeNode/children
func (t_ TreeNode) ChildNodes() []TreeNode {
	rv := objc.Send[[]TreeNode](t_.ID, objc.Sel("childNodes"))
	return rv
}/* debug [instance_properties/getter]: childNodes */


// The position of the receiver relative to its root parent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeNode/indexPath
func (t_ TreeNode) IndexPath() foundation.IndexPath {
	rv := objc.Send[foundation.IndexPath](t_.ID, objc.Sel("indexPath"))
	return rv
}/* debug [instance_properties/getter]: indexPath */


// A Boolean that indicates whether the receiver is a leaf node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeNode/isLeaf
func (t_ TreeNode) Leaf() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("leaf"))
	return rv
}/* debug [instance_properties/getter]: leaf */


// A mutable array that provides read-write access to the receiver’s child nodes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeNode/mutableChildren
func (t_ TreeNode) MutableChildNodes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("mutableChildNodes"))
	return rv
}/* debug [instance_properties/getter]: mutableChildNodes */


// The receiver’s parent node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeNode/parent
func (t_ TreeNode) ParentNode() ITreeNode {
	rv := objc.Send[TreeNode](t_.ID, objc.Sel("parentNode"))
	return rv
}/* debug [instance_properties/getter]: parentNode */


// The object the tree node represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTreeNode/representedObject
func (t_ TreeNode) RepresentedObject() objc.ID {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("representedObject"))
	return rv
}/* debug [instance_properties/getter]: representedObject */


// An array containing receiver’s child nodes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreenode/children
func (t_ TreeNode) Children() ITreeNode {
	rv := objc.Send[TreeNode](t_.ID, objc.Sel("children"))
	return rv
}/* debug [instance_properties/getter]: children */


// An array containing receiver’s child nodes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreenode/children
func (t_ TreeNode) SetChildren(value ITreeNode) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setChildren:"), value)
}/* debug [instance_properties/setter]: children */


// A Boolean that indicates whether the receiver is a leaf node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreenode/isleaf
func (t_ TreeNode) IsLeaf() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isLeaf"))
	return rv
}/* debug [instance_properties/getter]: isLeaf */


// A Boolean that indicates whether the receiver is a leaf node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreenode/isleaf
func (t_ TreeNode) SetIsLeaf(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsLeaf:"), value)
}/* debug [instance_properties/setter]: isLeaf */


// A mutable array that provides read-write access to the receiver’s child nodes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreenode/mutablechildren
func (t_ TreeNode) MutableChildren() foundation.MutableArray {
	rv := objc.Send[foundation.MutableArray](t_.ID, objc.Sel("mutableChildren"))
	return rv
}/* debug [instance_properties/getter]: mutableChildren */


// A mutable array that provides read-write access to the receiver’s child nodes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreenode/mutablechildren
func (t_ TreeNode) SetMutableChildren(value foundation.MutableArray) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setMutableChildren:"), value)
}/* debug [instance_properties/setter]: mutableChildren */


// The receiver’s parent node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreenode/parent
func (t_ TreeNode) Parent() ITreeNode {
	rv := objc.Send[TreeNode](t_.ID, objc.Sel("parent"))
	return rv
}/* debug [instance_properties/getter]: parent */


// The receiver’s parent node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstreenode/parent
func (t_ TreeNode) SetParent(value ITreeNode) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setParent:"), value)
}/* debug [instance_properties/setter]: parent */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSTreeNode */


