// Code generated from Apple documentation for PencilKit. DO NOT EDIT.

package pencilkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [LassoTool] class.
var (
	LassoToolClass     _LassoToolClass
	LassoToolClassOnce sync.Once
)

func getLassoToolClass() _LassoToolClass {
	LassoToolClassOnce.Do(func() {
		LassoToolClass = _LassoToolClass{objc.GetClass("PKLassoTool")}
	})
	return LassoToolClass
}

type _LassoToolClass struct {
	class objc.Class
}

// An interface definition for the [LassoTool] class.
type ILassoTool interface {
	ITool
}

// A tool for selecting stroked lines and shapes in a canvas view.
//
// A object supports the selection of content on a . When active, the canvas uses incoming touch events to determine what content to add to the selection. Create a lasso tool programmatically or display a object from which the user selects the tool. Assign the resulting object to the property of your object. The canvas uses any subsequent touch sequences to select content on the canvas.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKLassoToolReference
type LassoTool struct {
	Tool
}

// LassoToolFrom constructs a [LassoTool] from an unsafe.Pointer.
//
// A tool for selecting stroked lines and shapes in a canvas view.
func LassoToolFrom(ptr unsafe.Pointer) LassoTool {
	return LassoTool{
		Tool: ToolFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (lc _LassoToolClass) Alloc() LassoTool {
	rv := objc.Send[LassoTool](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (lc _LassoToolClass) New() LassoTool {
	rv := objc.Send[LassoTool](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ LassoTool) Init() LassoTool {
	rv := objc.Send[LassoTool](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ LassoTool) Autorelease() LassoTool {
	rv := objc.Send[LassoTool](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewLassoTool creates a new LassoTool instance.
func NewLassoTool() LassoTool {
	return getLassoToolClass().New()
}




