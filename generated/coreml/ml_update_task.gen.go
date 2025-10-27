// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





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





// An interface definition for the [UpdateTask] class.
type IUpdateTask interface {
	ITask
	

	// properties:


	

	// methods:
	ResumeWithParameters(updateParameters foundation.IDictionary)


}





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






// Creates a task that updates the model at the URL with the training data, and calls the completion handler when the update completes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLUpdateTask/init(forModelAt:trainingData:completionHandler:)
func NewUpdateTaskForModelAtURLTrainingDataCompletionHandlerError(modelURL foundation.foundation.INSURL, trainingData unsafe.Pointer, completionHandler unsafe.Pointer, error_ foundation.foundation.INSError) UpdateTask {
	rv := objc.Send[UpdateTask](objc.ID(getUpdateTaskClass().class), objc.Sel("updateTaskForModelAtURL:trainingData:completionHandler:error:"), modelURL, trainingData, completionHandler, error_)
	return rv
}


// Creates a task that updates the model at the URL with the training data and configuration, and calls the completion handler when the update completes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLUpdateTask/init(forModelAt:trainingData:configuration:completionHandler:)
func NewUpdateTaskForModelAtURLTrainingDataConfigurationCompletionHandlerError(modelURL foundation.foundation.INSURL, trainingData unsafe.Pointer, configuration IMLModelConfiguration, completionHandler unsafe.Pointer, error_ foundation.foundation.INSError) UpdateTask {
	rv := objc.Send[UpdateTask](objc.ID(getUpdateTaskClass().class), objc.Sel("updateTaskForModelAtURL:trainingData:configuration:completionHandler:error:"), modelURL, trainingData, configuration, completionHandler, error_)
	return rv
}


// Creates a task that updates the model at the URL with the training data and configuration, and calls the progress handlers during and after the update.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLUpdateTask/init(forModelAt:trainingData:configuration:progressHandlers:)
func NewUpdateTaskForModelAtURLTrainingDataConfigurationProgressHandlersError(modelURL foundation.foundation.INSURL, trainingData unsafe.Pointer, configuration IMLModelConfiguration, progressHandlers IMLUpdateProgressHandlers, error_ foundation.foundation.INSError) UpdateTask {
	rv := objc.Send[UpdateTask](objc.ID(getUpdateTaskClass().class), objc.Sel("updateTaskForModelAtURL:trainingData:configuration:progressHandlers:error:"), modelURL, trainingData, configuration, progressHandlers, error_)
	return rv
}


// Creates a task that updates the model at the URL with the training data, and calls the progress handlers during and after the update.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLUpdateTask/init(forModelAt:trainingData:progressHandlers:)
func NewUpdateTaskForModelAtURLTrainingDataProgressHandlersError(modelURL foundation.foundation.INSURL, trainingData unsafe.Pointer, progressHandlers IMLUpdateProgressHandlers, error_ foundation.foundation.INSError) UpdateTask {
	rv := objc.Send[UpdateTask](objc.ID(getUpdateTaskClass().class), objc.Sel("updateTaskForModelAtURL:trainingData:progressHandlers:error:"), modelURL, trainingData, progressHandlers, error_)
	return rv
}







// Creates a task that updates the model at the URL with the training data, and calls the completion handler when the update completes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLUpdateTask/init(forModelAt:trainingData:completionHandler:)
func (uc _UpdateTaskClass) UpdateTaskForModelAtURLTrainingDataCompletionHandlerError(modelURL foundation.foundation.INSURL, trainingData unsafe.Pointer, completionHandler unsafe.Pointer, error_ foundation.foundation.INSError) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(uc.class), objc.Sel("updateTaskForModelAtURL:trainingData:completionHandler:error:"), modelURL, trainingData, completionHandler, error_)
	return rv
}


// Creates a task that updates the model at the URL with the training data and configuration, and calls the completion handler when the update completes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLUpdateTask/init(forModelAt:trainingData:configuration:completionHandler:)
func (uc _UpdateTaskClass) UpdateTaskForModelAtURLTrainingDataConfigurationCompletionHandlerError(modelURL foundation.foundation.INSURL, trainingData unsafe.Pointer, configuration IMLModelConfiguration, completionHandler unsafe.Pointer, error_ foundation.foundation.INSError) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(uc.class), objc.Sel("updateTaskForModelAtURL:trainingData:configuration:completionHandler:error:"), modelURL, trainingData, configuration, completionHandler, error_)
	return rv
}


// Creates a task that updates the model at the URL with the training data and configuration, and calls the progress handlers during and after the update.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLUpdateTask/init(forModelAt:trainingData:configuration:progressHandlers:)
func (uc _UpdateTaskClass) UpdateTaskForModelAtURLTrainingDataConfigurationProgressHandlersError(modelURL foundation.foundation.INSURL, trainingData unsafe.Pointer, configuration IMLModelConfiguration, progressHandlers IMLUpdateProgressHandlers, error_ foundation.foundation.INSError) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(uc.class), objc.Sel("updateTaskForModelAtURL:trainingData:configuration:progressHandlers:error:"), modelURL, trainingData, configuration, progressHandlers, error_)
	return rv
}


// Creates a task that updates the model at the URL with the training data, and calls the progress handlers during and after the update.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLUpdateTask/init(forModelAt:trainingData:progressHandlers:)
func (uc _UpdateTaskClass) UpdateTaskForModelAtURLTrainingDataProgressHandlersError(modelURL foundation.foundation.INSURL, trainingData unsafe.Pointer, progressHandlers IMLUpdateProgressHandlers, error_ foundation.foundation.INSError) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(uc.class), objc.Sel("updateTaskForModelAtURL:trainingData:progressHandlers:error:"), modelURL, trainingData, progressHandlers, error_)
	return rv
}












// Resumes a model update with updated parameter values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLUpdateTask/resume(withParameters:)
func (u_ UpdateTask) ResumeWithParameters(updateParameters foundation.IDictionary) {
	objc.Send[objc.ID](u_.ID, objc.Sel("resumeWithParameters:"), updateParameters)
}












