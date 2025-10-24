// Code generated from Apple documentation for Automator. DO NOT EDIT.

package automator

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AMWorkspace] class.
var (
	AMWorkspaceClass     _AMWorkspaceClass
	AMWorkspaceClassOnce sync.Once
)

func getAMWorkspaceClass() _AMWorkspaceClass {
	AMWorkspaceClassOnce.Do(func() {
		AMWorkspaceClass = _AMWorkspaceClass{objc.GetClass("AMWorkspace")}
	})
	return AMWorkspaceClass
}

type _AMWorkspaceClass struct {
	class objc.Class
}

// An interface definition for the [AMWorkspace] class.
type IAMWorkspace interface {
	objectivec.IObject
	// properties:
	// methods:
	RunWorkflowAtPathWithInputError(path objc.IObject /* cross-framework: NSString */, input objectivec.IObject, error_ unsafe.Pointer) objc.ID
}

// A workspace for running an Automator workflow.
//
// The class provides access to the shared workspace in the Automator framework, where you can run workflows without a workflow controller. Use to access the shared workspace and to run your workflow in it.


// A workspace for running an Automator workflow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMWorkspace
type AMWorkspace struct {
	objectivec.Object
}

// AMWorkspaceFrom constructs a [AMWorkspace] from an unsafe.Pointer.
//
// A workspace for running an Automator workflow.
func AMWorkspaceFrom(ptr unsafe.Pointer) AMWorkspace {
	return AMWorkspace{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AMWorkspaceClass) Alloc() AMWorkspace {
	rv := objc.Send[AMWorkspace](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AMWorkspaceClass) New() AMWorkspace {
	rv := objc.Send[AMWorkspace](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AMWorkspace) Init() AMWorkspace {
	rv := objc.Send[AMWorkspace](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AMWorkspace) Autorelease() AMWorkspace {
	rv := objc.Send[AMWorkspace](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAMWorkspace creates a new AMWorkspace instance.
func NewAMWorkspace() AMWorkspace {
	return getAMWorkspaceClass().New()
}



// The shared workspace object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMWorkspace/shared
func (ac _AMWorkspaceClass) SharedWorkspace() AMWorkspace {
	rv := objc.Send[AMWorkspace](objc.ID(ac.class), objc.Sel("sharedWorkspace"))
	return rv
}

// Loads and runs the specified workflow file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMWorkspace/runWorkflow(atPath:withInput:)
func (a_ AMWorkspace) RunWorkflowAtPathWithInputError(path objc.IObject /* cross-framework: NSString */, input objectivec.IObject, error_ unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("runWorkflowAtPath:withInput:error:"), path, input, error_)
	return rv
}


// The shared workspace object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Automator/AMWorkspace/shared
func (a_ AMWorkspace) SharedWorkspace() IAMWorkspace {
	rv := objc.Send[AMWorkspace](a_.ID, objc.Sel("sharedWorkspace"))
	return rv
}




