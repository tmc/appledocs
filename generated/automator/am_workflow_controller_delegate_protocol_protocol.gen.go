// Code generated from Apple documentation for Automator. DO NOT EDIT.

package automator

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/coretelephony"
)

// PAMWorkflowControllerDelegate is the AMWorkflowControllerDelegate protocol interface.
//
// A set of optional methods that a delegate of a workflow controller implements.
//
// Availability:
//   - Mac Catalyst 14.0+
//   - macOS 10.4+
//
// See: doc://com.apple.automator/documentation/Automator/AMWorkflowControllerDelegate
type PAMWorkflowControllerDelegate interface {
	// Optional methods
	WorkflowControllerDidError(controller IAMWorkflowController, error_ objc.IObject /* cross-framework: Error */)
	HasWorkflowControllerDidError() bool
	WorkflowControllerDidRunAction(controller IAMWorkflowController, action IAMAction)
	HasWorkflowControllerDidRunAction() bool
	WorkflowControllerWillRunAction(controller IAMWorkflowController, action IAMAction)
	HasWorkflowControllerWillRunAction() bool
	WorkflowControllerDidRun(controller IAMWorkflowController)
	HasWorkflowControllerDidRun() bool
	WorkflowControllerDidStop(controller IAMWorkflowController)
	HasWorkflowControllerDidStop() bool
	WorkflowControllerWillRun(controller IAMWorkflowController)
	HasWorkflowControllerWillRun() bool
	WorkflowControllerWillStop(controller IAMWorkflowController)
	HasWorkflowControllerWillStop() bool
}

// AMWorkflowControllerDelegate is a delegate implementation builder for the PAMWorkflowControllerDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type AMWorkflowControllerDelegate struct {
	_WorkflowControllerDidError func(controller IAMWorkflowController, error_ objc.IObject /* cross-framework: Error */)
	_WorkflowControllerDidRunAction func(controller IAMWorkflowController, action IAMAction)
	_WorkflowControllerWillRunAction func(controller IAMWorkflowController, action IAMAction)
	_WorkflowControllerDidRun func(controller IAMWorkflowController)
	_WorkflowControllerDidStop func(controller IAMWorkflowController)
	_WorkflowControllerWillRun func(controller IAMWorkflowController)
	_WorkflowControllerWillStop func(controller IAMWorkflowController)
}

// SetWorkflowControllerDidError sets the handler for the WorkflowControllerDidError delegate method.
//
// Notifies the delegate when the workflow encounters an error.
func (d *AMWorkflowControllerDelegate) SetWorkflowControllerDidError(f func(controller IAMWorkflowController, error_ objc.IObject /* cross-framework: Error */)) {
	d._WorkflowControllerDidError = f
}

// SetWorkflowControllerDidRunAction sets the handler for the WorkflowControllerDidRunAction delegate method.
//
// Notifies the delegate when the specified action finishes running.
func (d *AMWorkflowControllerDelegate) SetWorkflowControllerDidRunAction(f func(controller IAMWorkflowController, action IAMAction)) {
	d._WorkflowControllerDidRunAction = f
}

// SetWorkflowControllerWillRunAction sets the handler for the WorkflowControllerWillRunAction delegate method.
//
// Notifies the delegate when the specified action is about to run.
func (d *AMWorkflowControllerDelegate) SetWorkflowControllerWillRunAction(f func(controller IAMWorkflowController, action IAMAction)) {
	d._WorkflowControllerWillRunAction = f
}

// SetWorkflowControllerDidRun sets the handler for the WorkflowControllerDidRun delegate method.
//
// Notifies the delegate when the workflow controller object finishes running.
func (d *AMWorkflowControllerDelegate) SetWorkflowControllerDidRun(f func(controller IAMWorkflowController)) {
	d._WorkflowControllerDidRun = f
}

// SetWorkflowControllerDidStop sets the handler for the WorkflowControllerDidStop delegate method.
//
// Tells the delegate that the workflow controller object has stopped.
func (d *AMWorkflowControllerDelegate) SetWorkflowControllerDidStop(f func(controller IAMWorkflowController)) {
	d._WorkflowControllerDidStop = f
}

// SetWorkflowControllerWillRun sets the handler for the WorkflowControllerWillRun delegate method.
//
// Notifies the delegate when the workflow controller object is about to run.
func (d *AMWorkflowControllerDelegate) SetWorkflowControllerWillRun(f func(controller IAMWorkflowController)) {
	d._WorkflowControllerWillRun = f
}

// SetWorkflowControllerWillStop sets the handler for the WorkflowControllerWillStop delegate method.
//
// Tells the delegate that the workflow controller object is about to stop.
func (d *AMWorkflowControllerDelegate) SetWorkflowControllerWillStop(f func(controller IAMWorkflowController)) {
	d._WorkflowControllerWillStop = f
}

// WorkflowControllerDidError implements the PAMWorkflowControllerDelegate interface.
func (d *AMWorkflowControllerDelegate) WorkflowControllerDidError(controller IAMWorkflowController, error_ objc.IObject /* cross-framework: Error */) {
	if d._WorkflowControllerDidError != nil {
		d._WorkflowControllerDidError(controller, error_)
	}
}

// HasWorkflowControllerDidError returns true if a handler for WorkflowControllerDidError has been set.
func (d *AMWorkflowControllerDelegate) HasWorkflowControllerDidError() bool {
	return d._WorkflowControllerDidError != nil
}

// WorkflowControllerDidRunAction implements the PAMWorkflowControllerDelegate interface.
func (d *AMWorkflowControllerDelegate) WorkflowControllerDidRunAction(controller IAMWorkflowController, action IAMAction) {
	if d._WorkflowControllerDidRunAction != nil {
		d._WorkflowControllerDidRunAction(controller, action)
	}
}

// HasWorkflowControllerDidRunAction returns true if a handler for WorkflowControllerDidRunAction has been set.
func (d *AMWorkflowControllerDelegate) HasWorkflowControllerDidRunAction() bool {
	return d._WorkflowControllerDidRunAction != nil
}

// WorkflowControllerWillRunAction implements the PAMWorkflowControllerDelegate interface.
func (d *AMWorkflowControllerDelegate) WorkflowControllerWillRunAction(controller IAMWorkflowController, action IAMAction) {
	if d._WorkflowControllerWillRunAction != nil {
		d._WorkflowControllerWillRunAction(controller, action)
	}
}

// HasWorkflowControllerWillRunAction returns true if a handler for WorkflowControllerWillRunAction has been set.
func (d *AMWorkflowControllerDelegate) HasWorkflowControllerWillRunAction() bool {
	return d._WorkflowControllerWillRunAction != nil
}

// WorkflowControllerDidRun implements the PAMWorkflowControllerDelegate interface.
func (d *AMWorkflowControllerDelegate) WorkflowControllerDidRun(controller IAMWorkflowController) {
	if d._WorkflowControllerDidRun != nil {
		d._WorkflowControllerDidRun(controller)
	}
}

// HasWorkflowControllerDidRun returns true if a handler for WorkflowControllerDidRun has been set.
func (d *AMWorkflowControllerDelegate) HasWorkflowControllerDidRun() bool {
	return d._WorkflowControllerDidRun != nil
}

// WorkflowControllerDidStop implements the PAMWorkflowControllerDelegate interface.
func (d *AMWorkflowControllerDelegate) WorkflowControllerDidStop(controller IAMWorkflowController) {
	if d._WorkflowControllerDidStop != nil {
		d._WorkflowControllerDidStop(controller)
	}
}

// HasWorkflowControllerDidStop returns true if a handler for WorkflowControllerDidStop has been set.
func (d *AMWorkflowControllerDelegate) HasWorkflowControllerDidStop() bool {
	return d._WorkflowControllerDidStop != nil
}

// WorkflowControllerWillRun implements the PAMWorkflowControllerDelegate interface.
func (d *AMWorkflowControllerDelegate) WorkflowControllerWillRun(controller IAMWorkflowController) {
	if d._WorkflowControllerWillRun != nil {
		d._WorkflowControllerWillRun(controller)
	}
}

// HasWorkflowControllerWillRun returns true if a handler for WorkflowControllerWillRun has been set.
func (d *AMWorkflowControllerDelegate) HasWorkflowControllerWillRun() bool {
	return d._WorkflowControllerWillRun != nil
}

// WorkflowControllerWillStop implements the PAMWorkflowControllerDelegate interface.
func (d *AMWorkflowControllerDelegate) WorkflowControllerWillStop(controller IAMWorkflowController) {
	if d._WorkflowControllerWillStop != nil {
		d._WorkflowControllerWillStop(controller)
	}
}

// HasWorkflowControllerWillStop returns true if a handler for WorkflowControllerWillStop has been set.
func (d *AMWorkflowControllerDelegate) HasWorkflowControllerWillStop() bool {
	return d._WorkflowControllerWillStop != nil
}
