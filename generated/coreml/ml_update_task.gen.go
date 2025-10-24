// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MLUpdateTask */


/* debug [class_header]: Header for MLUpdateTask */
// The class instance for the [UpdateTask] class.
var (
	UpdateTaskClass     _UpdateTaskClass
	UpdateTaskClassOnce sync.Once
)

func getUpdateTaskClass() _UpdateTaskClass {
	UpdateTaskClassOnce.Do(func() {
		UpdateTaskClass = _UpdateTaskClass{objc.GetClass("MLUpdateTask")}
	})
	return UpdateTaskClass
}

type _UpdateTaskClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for UpdateTask */
// An interface definition for the [UpdateTask] class.
type IUpdateTask interface {
	ITask
	
/* debug [class_interface_properties]: Properties for UpdateTask */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for UpdateTask */
	// methods:
	ResumeWithParameters(updateParameters foundation.IDictionary)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for UpdateTask */
// Alloc allocates a new instance without initialization.
func (uc _UpdateTaskClass) Alloc() UpdateTask {
	rv := objc.Send[UpdateTask](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (uc _UpdateTaskClass) New() UpdateTask {
	rv := objc.Send[UpdateTask](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UpdateTask) Init() UpdateTask {
	rv := objc.Send[UpdateTask](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UpdateTask) Autorelease() UpdateTask {
	rv := objc.Send[UpdateTask](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUpdateTask creates a new UpdateTask instance.
func NewUpdateTask() UpdateTask {
	return getUpdateTaskClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for UpdateTask */
// A task that updates a model with additional training data.
//
// Use an to update a machine learning model on a user’s device.


// A task that updates a model with additional training data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLUpdateTask
type UpdateTask struct {
	Task
}

// UpdateTaskFrom constructs a [UpdateTask] from an unsafe.Pointer.
//
// A task that updates a model with additional training data.
func UpdateTaskFrom(ptr unsafe.Pointer) UpdateTask {
	return UpdateTask{
		Task: TaskFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for UpdateTask */

// Creates a task that updates the model at the URL with the training data, and calls the completion handler when the update completes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLUpdateTask/init(forModelAt:trainingData:completionHandler:)
func NewUpdateTaskForModelAtURLTrainingDataCompletionHandlerError(modelURL objc.IObject /* cross-framework: NSURL */, trainingData unsafe.Pointer, completionHandler unsafe.Pointer, error_ objectivec.IObject) UpdateTask {
	rv := objc.Send[UpdateTask](objc.ID(getUpdateTaskClass().class), objc.Sel("updateTaskForModelAtURL:trainingData:completionHandler:error:"), modelURL, trainingData, completionHandler, error_)
	return rv
}/* debug [class_init_methods/constructor]: NewUpdateTaskForModelAtURLTrainingDataCompletionHandlerError */


// Creates a task that updates the model at the URL with the training data and configuration, and calls the completion handler when the update completes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLUpdateTask/init(forModelAt:trainingData:configuration:completionHandler:)
func NewUpdateTaskForModelAtURLTrainingDataConfigurationCompletionHandlerError(modelURL objc.IObject /* cross-framework: NSURL */, trainingData unsafe.Pointer, configuration IMLModelConfiguration, completionHandler unsafe.Pointer, error_ objectivec.IObject) UpdateTask {
	rv := objc.Send[UpdateTask](objc.ID(getUpdateTaskClass().class), objc.Sel("updateTaskForModelAtURL:trainingData:configuration:completionHandler:error:"), modelURL, trainingData, configuration, completionHandler, error_)
	return rv
}/* debug [class_init_methods/constructor]: NewUpdateTaskForModelAtURLTrainingDataConfigurationCompletionHandlerError */


// Creates a task that updates the model at the URL with the training data and configuration, and calls the progress handlers during and after the update.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLUpdateTask/init(forModelAt:trainingData:configuration:progressHandlers:)
func NewUpdateTaskForModelAtURLTrainingDataConfigurationProgressHandlersError(modelURL objc.IObject /* cross-framework: NSURL */, trainingData unsafe.Pointer, configuration IMLModelConfiguration, progressHandlers IMLUpdateProgressHandlers, error_ objectivec.IObject) UpdateTask {
	rv := objc.Send[UpdateTask](objc.ID(getUpdateTaskClass().class), objc.Sel("updateTaskForModelAtURL:trainingData:configuration:progressHandlers:error:"), modelURL, trainingData, configuration, progressHandlers, error_)
	return rv
}/* debug [class_init_methods/constructor]: NewUpdateTaskForModelAtURLTrainingDataConfigurationProgressHandlersError */


// Creates a task that updates the model at the URL with the training data, and calls the progress handlers during and after the update.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLUpdateTask/init(forModelAt:trainingData:progressHandlers:)
func NewUpdateTaskForModelAtURLTrainingDataProgressHandlersError(modelURL objc.IObject /* cross-framework: NSURL */, trainingData unsafe.Pointer, progressHandlers IMLUpdateProgressHandlers, error_ objectivec.IObject) UpdateTask {
	rv := objc.Send[UpdateTask](objc.ID(getUpdateTaskClass().class), objc.Sel("updateTaskForModelAtURL:trainingData:progressHandlers:error:"), modelURL, trainingData, progressHandlers, error_)
	return rv
}/* debug [class_init_methods/constructor]: NewUpdateTaskForModelAtURLTrainingDataProgressHandlersError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for UpdateTask */

// Creates a task that updates the model at the URL with the training data, and calls the completion handler when the update completes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLUpdateTask/init(forModelAt:trainingData:completionHandler:)
func (uc _UpdateTaskClass) UpdateTaskForModelAtURLTrainingDataCompletionHandlerError(modelURL objc.IObject /* cross-framework: NSURL */, trainingData unsafe.Pointer, completionHandler unsafe.Pointer, error_ objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(uc.class), objc.Sel("updateTaskForModelAtURL:trainingData:completionHandler:error:"), modelURL, trainingData, completionHandler, error_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=UpdateTaskForModelAtURLTrainingDataCompletionHandlerError) */


// Creates a task that updates the model at the URL with the training data and configuration, and calls the completion handler when the update completes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLUpdateTask/init(forModelAt:trainingData:configuration:completionHandler:)
func (uc _UpdateTaskClass) UpdateTaskForModelAtURLTrainingDataConfigurationCompletionHandlerError(modelURL objc.IObject /* cross-framework: NSURL */, trainingData unsafe.Pointer, configuration IMLModelConfiguration, completionHandler unsafe.Pointer, error_ objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(uc.class), objc.Sel("updateTaskForModelAtURL:trainingData:configuration:completionHandler:error:"), modelURL, trainingData, configuration, completionHandler, error_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=UpdateTaskForModelAtURLTrainingDataConfigurationCompletionHandlerError) */


// Creates a task that updates the model at the URL with the training data and configuration, and calls the progress handlers during and after the update.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLUpdateTask/init(forModelAt:trainingData:configuration:progressHandlers:)
func (uc _UpdateTaskClass) UpdateTaskForModelAtURLTrainingDataConfigurationProgressHandlersError(modelURL objc.IObject /* cross-framework: NSURL */, trainingData unsafe.Pointer, configuration IMLModelConfiguration, progressHandlers IMLUpdateProgressHandlers, error_ objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(uc.class), objc.Sel("updateTaskForModelAtURL:trainingData:configuration:progressHandlers:error:"), modelURL, trainingData, configuration, progressHandlers, error_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=UpdateTaskForModelAtURLTrainingDataConfigurationProgressHandlersError) */


// Creates a task that updates the model at the URL with the training data, and calls the progress handlers during and after the update.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLUpdateTask/init(forModelAt:trainingData:progressHandlers:)
func (uc _UpdateTaskClass) UpdateTaskForModelAtURLTrainingDataProgressHandlersError(modelURL objc.IObject /* cross-framework: NSURL */, trainingData unsafe.Pointer, progressHandlers IMLUpdateProgressHandlers, error_ objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(uc.class), objc.Sel("updateTaskForModelAtURL:trainingData:progressHandlers:error:"), modelURL, trainingData, progressHandlers, error_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=UpdateTaskForModelAtURLTrainingDataProgressHandlersError) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for UpdateTask */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for UpdateTask */

// Resumes a model update with updated parameter values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLUpdateTask/resume(withParameters:)
func (u_ UpdateTask) ResumeWithParameters(updateParameters foundation.IDictionary) {
	objc.Send[objc.ID](u_.ID, objc.Sel("resumeWithParameters:"), updateParameters)
}/* debug [instance_methods/method]: ResumeWithParameters */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for UpdateTask */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLUpdateTask */


