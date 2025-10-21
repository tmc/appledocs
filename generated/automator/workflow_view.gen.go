// Code generated from Apple documentation for Automator. DO NOT EDIT.

package automator

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [workflowView] class.
var (
	WorkflowViewClass     _workflowViewClass
	WorkflowViewClassOnce sync.Once
)

func getworkflowViewClass() _workflowViewClass {
	WorkflowViewClassOnce.Do(func() {
		WorkflowViewClass = _workflowViewClass{objc.GetClass("workflowView")}
	})
	return WorkflowViewClass
}

type _workflowViewClass struct {
	class objc.Class
}

// An interface definition for the [workflowView] class.
type IworkflowView interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMWorkflowController/workflowView-c.ivar
type workflowView struct {
	objectivec.Object
}

// workflowViewFrom constructs a [workflowView] from an unsafe.Pointer.
func workflowViewFrom(ptr unsafe.Pointer) workflowView {
	return workflowView{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (wc _workflowViewClass) Alloc() workflowView {
	rv := objc.Send[workflowView](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (wc _workflowViewClass) New() workflowView {
	rv := objc.Send[workflowView](objc.ID(wc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (w_ workflowView) Init() workflowView {
	rv := objc.Send[workflowView](w_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w_ workflowView) Autorelease() workflowView {
	rv := objc.Send[workflowView](w_.ID, objc.Sel("autorelease"))
	return rv
}

// NewworkflowView creates a new workflowView instance.
func NewworkflowView() workflowView {
	return getworkflowViewClass().New()
}




