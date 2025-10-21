// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	ResumeWithParameters(updateParameters unsafe.Pointer)
}

// A task that updates a model with additional training data.
//
// Use an to update a machine learning model on a user’s device.
//
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

// Alloc allocates a new instance without initialization.
func (uc _UpdateTaskClass) Alloc() UpdateTask {
	rv := objc.Send[UpdateTask](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Creates a task that updates the model at the URL with the training data, and calls the completion handler when the update completes.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLUpdateTask/init(forModelAt:trainingData:completionHandler:)
func NewUpdateTaskForModelAtURLTrainingDataCompletionHandlerError(modelURL unsafe.Pointer, trainingData objc.ID, completionHandler unsafe.Pointer, error_ unsafe.Pointer) UpdateTask {
	rv := objc.Send[UpdateTask](objc.ID(getUpdateTaskClass().class), objc.Sel("updateTaskForModelAtURL:trainingData:completionHandler:error:"), modelURL, trainingData, completionHandler, error_)
	return rv
}



// Creates a task that updates the model at the URL with the training data and configuration, and calls the completion handler when the update completes.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLUpdateTask/init(forModelAt:trainingData:configuration:completionHandler:)
func NewUpdateTaskForModelAtURLTrainingDataConfigurationCompletionHandlerError(modelURL unsafe.Pointer, trainingData objc.ID, configuration unsafe.Pointer, completionHandler unsafe.Pointer, error_ unsafe.Pointer) UpdateTask {
	rv := objc.Send[UpdateTask](objc.ID(getUpdateTaskClass().class), objc.Sel("updateTaskForModelAtURL:trainingData:configuration:completionHandler:error:"), modelURL, trainingData, configuration, completionHandler, error_)
	return rv
}



// Creates a task that updates the model at the URL with the training data and configuration, and calls the progress handlers during and after the update.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLUpdateTask/init(forModelAt:trainingData:configuration:progressHandlers:)
func NewUpdateTaskForModelAtURLTrainingDataConfigurationProgressHandlersError(modelURL unsafe.Pointer, trainingData objc.ID, configuration unsafe.Pointer, progressHandlers unsafe.Pointer, error_ unsafe.Pointer) UpdateTask {
	rv := objc.Send[UpdateTask](objc.ID(getUpdateTaskClass().class), objc.Sel("updateTaskForModelAtURL:trainingData:configuration:progressHandlers:error:"), modelURL, trainingData, configuration, progressHandlers, error_)
	return rv
}



// Creates a task that updates the model at the URL with the training data, and calls the progress handlers during and after the update.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLUpdateTask/init(forModelAt:trainingData:progressHandlers:)
func NewUpdateTaskForModelAtURLTrainingDataProgressHandlersError(modelURL unsafe.Pointer, trainingData objc.ID, progressHandlers unsafe.Pointer, error_ unsafe.Pointer) UpdateTask {
	rv := objc.Send[UpdateTask](objc.ID(getUpdateTaskClass().class), objc.Sel("updateTaskForModelAtURL:trainingData:progressHandlers:error:"), modelURL, trainingData, progressHandlers, error_)
	return rv
}


// Creates a task that updates the model at the URL with the training data, and calls the completion handler when the update completes.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLUpdateTask/init(forModelAt:trainingData:completionHandler:)
func (uc _UpdateTaskClass) UpdateTaskForModelAtURLTrainingDataCompletionHandlerError(modelURL unsafe.Pointer, trainingData objc.ID, completionHandler unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("updateTaskForModelAtURL:trainingData:completionHandler:error:"), modelURL, trainingData, completionHandler, error_)
	return rv
}

// Creates a task that updates the model at the URL with the training data and configuration, and calls the completion handler when the update completes.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLUpdateTask/init(forModelAt:trainingData:configuration:completionHandler:)
func (uc _UpdateTaskClass) UpdateTaskForModelAtURLTrainingDataConfigurationCompletionHandlerError(modelURL unsafe.Pointer, trainingData objc.ID, configuration unsafe.Pointer, completionHandler unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("updateTaskForModelAtURL:trainingData:configuration:completionHandler:error:"), modelURL, trainingData, configuration, completionHandler, error_)
	return rv
}

// Creates a task that updates the model at the URL with the training data and configuration, and calls the progress handlers during and after the update.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLUpdateTask/init(forModelAt:trainingData:configuration:progressHandlers:)
func (uc _UpdateTaskClass) UpdateTaskForModelAtURLTrainingDataConfigurationProgressHandlersError(modelURL unsafe.Pointer, trainingData objc.ID, configuration unsafe.Pointer, progressHandlers unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("updateTaskForModelAtURL:trainingData:configuration:progressHandlers:error:"), modelURL, trainingData, configuration, progressHandlers, error_)
	return rv
}

// Creates a task that updates the model at the URL with the training data, and calls the progress handlers during and after the update.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLUpdateTask/init(forModelAt:trainingData:progressHandlers:)
func (uc _UpdateTaskClass) UpdateTaskForModelAtURLTrainingDataProgressHandlersError(modelURL unsafe.Pointer, trainingData objc.ID, progressHandlers unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(uc.class), objc.Sel("updateTaskForModelAtURL:trainingData:progressHandlers:error:"), modelURL, trainingData, progressHandlers, error_)
	return rv
}

// Resumes a model update with updated parameter values.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLUpdateTask/resume(withParameters:)
func (u_ UpdateTask) ResumeWithParameters(updateParameters unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("resumeWithParameters:"), updateParameters)
}


