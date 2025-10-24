// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ConstraintLayoutManager] class.
var (
	ConstraintLayoutManagerClass     _ConstraintLayoutManagerClass
	ConstraintLayoutManagerClassOnce sync.Once
)

func getConstraintLayoutManagerClass() _ConstraintLayoutManagerClass {
	ConstraintLayoutManagerClassOnce.Do(func() {
		ConstraintLayoutManagerClass = _ConstraintLayoutManagerClass{objc.GetClass("CAConstraintLayoutManager")}
	})
	return ConstraintLayoutManagerClass
}

type _ConstraintLayoutManagerClass struct {
	class objc.Class
}

// An interface definition for the [ConstraintLayoutManager] class.
type IConstraintLayoutManager interface {
	objectivec.IObject
	// properties:
	LayoutManager() objectivec.IObject
	SetLayoutManager(value objectivec.IObject)
	Name() objc.IObject /* cross-framework: NSString */
	SetName(value objc.IObject /* cross-framework: NSString */)
	// methods:
}

// An object that provides a constraint-based layout manager.
//
// You use the shared instance of this object by assigning it to the property of any layer objects to which you have added constraints. During a layout update, Core Animation uses the layout manager to update the size and position of the sublayers based on the registered set of constraints. Constraints let you define a set of geometric relationships between a layer and its sibling layers or between a layer and its superlayer. These relationships are expressed using constraint objects, which are instances of the class. When creating constraints, you can reference a layer by name using that object’s property. You can also use the special name to refer to the layer’s superlayer. The following example shows how you can use to create a layer containing two constrained sublayers: and . A series of objects are created so that the sublayers match their superlayer’s height and are half of its width. matches the attribute and matches the attribute. The end result is that the two sublayers are always laid out so that fills the left half of and fills the right half of layer. This class is not meant to be subclassed.


// An object that provides a constraint-based layout manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAConstraintLayoutManager
type ConstraintLayoutManager struct {
	objectivec.Object
}

// ConstraintLayoutManagerFrom constructs a [ConstraintLayoutManager] from an unsafe.Pointer.
//
// An object that provides a constraint-based layout manager.
func ConstraintLayoutManagerFrom(ptr unsafe.Pointer) ConstraintLayoutManager {
	return ConstraintLayoutManager{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _ConstraintLayoutManagerClass) Alloc() ConstraintLayoutManager {
	rv := objc.Send[ConstraintLayoutManager](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _ConstraintLayoutManagerClass) New() ConstraintLayoutManager {
	rv := objc.Send[ConstraintLayoutManager](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ConstraintLayoutManager) Init() ConstraintLayoutManager {
	rv := objc.Send[ConstraintLayoutManager](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ConstraintLayoutManager) Autorelease() ConstraintLayoutManager {
	rv := objc.Send[ConstraintLayoutManager](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewConstraintLayoutManager creates a new ConstraintLayoutManager instance.
func NewConstraintLayoutManager() ConstraintLayoutManager {
	return getConstraintLayoutManagerClass().New()
}



// Returns the shared layout manager object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAConstraintLayoutManager/layoutManager
func (cc _ConstraintLayoutManagerClass) LayoutManager() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("layoutManager"))
	return rv
}


// The object responsible for laying out the layer’s sublayers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/layoutmanager
func (c_ ConstraintLayoutManager) LayoutManager() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("layoutManager"))
	return rv
}


// The object responsible for laying out the layer’s sublayers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/layoutmanager
func (c_ ConstraintLayoutManager) SetLayoutManager(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLayoutManager:"), value)
}


// The name of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/name
func (c_ ConstraintLayoutManager) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("name"))
	return rv
}


// The name of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/name
func (c_ ConstraintLayoutManager) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setName:"), value)
}



