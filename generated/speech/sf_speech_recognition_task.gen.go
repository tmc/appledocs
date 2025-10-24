// Code generated from Apple documentation for Speech. DO NOT EDIT.

package speech

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/coretelephony"
	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class SFSpeechRecognitionTask */

/* debug [class_header]: Header for SFSpeechRecognitionTask */
// The class instance for the [SFSpeechRecognitionTask] class.
var (
	SFSpeechRecognitionTaskClass     _SFSpeechRecognitionTaskClass
	SFSpeechRecognitionTaskClassOnce sync.Once
)

func getSFSpeechRecognitionTaskClass() _SFSpeechRecognitionTaskClass {
	SFSpeechRecognitionTaskClassOnce.Do(func() {
		SFSpeechRecognitionTaskClass = _SFSpeechRecognitionTaskClass{objc.GetClass("SFSpeechRecognitionTask")}
	})
	return SFSpeechRecognitionTaskClass
}

type _SFSpeechRecognitionTaskClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for SFSpeechRecognitionTask */
// An interface definition for the [SFSpeechRecognitionTask] class.
type ISFSpeechRecognitionTask interface {
	objectivec.IObject

	/* debug [class_interface_properties]: Properties for SFSpeechRecognitionTask */
	// properties:
	Error() objc.IObject /* cross-framework: Error */
	Cancelled() bool
	Finishing() bool
	State() SFSpeechRecognitionTaskState
	IsCancelled() bool
	SetIsCancelled(value bool)
	IsFinishing() bool
	SetIsFinishing(value bool)
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for SFSpeechRecognitionTask */
	// methods:
	Cancel()
	Finish()
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for SFSpeechRecognitionTask */
// Alloc allocates a new instance without initialization.
func (sc _SFSpeechRecognitionTaskClass) Alloc() SFSpeechRecognitionTask {
	rv := objc.Send[SFSpeechRecognitionTask](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SFSpeechRecognitionTaskClass) New() SFSpeechRecognitionTask {
	rv := objc.Send[SFSpeechRecognitionTask](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SFSpeechRecognitionTask) Init() SFSpeechRecognitionTask {
	rv := objc.Send[SFSpeechRecognitionTask](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SFSpeechRecognitionTask) Autorelease() SFSpeechRecognitionTask {
	rv := objc.Send[SFSpeechRecognitionTask](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSFSpeechRecognitionTask creates a new SFSpeechRecognitionTask instance.
func NewSFSpeechRecognitionTask() SFSpeechRecognitionTask {
	return getSFSpeechRecognitionTaskClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for SFSpeechRecognitionTask */
// A task object for monitoring the speech recognition progress.
//
// Use an object to determine the state of a speech recognition task, to cancel an ongoing task, or to signal the end of the task. You don’t create speech recognition task objects directly. Instead, you receive one of these objects after calling or on your object.

// A task object for monitoring the speech recognition progress.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionTask
type SFSpeechRecognitionTask struct {
	objectivec.Object
}

// SFSpeechRecognitionTaskFrom constructs a [SFSpeechRecognitionTask] from an unsafe.Pointer.
//
// A task object for monitoring the speech recognition progress.
func SFSpeechRecognitionTaskFrom(ptr unsafe.Pointer) SFSpeechRecognitionTask {
	return SFSpeechRecognitionTask{objectivec.Object{objc.ID(ptr)}}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for SFSpeechRecognitionTask */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for SFSpeechRecognitionTask */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for SFSpeechRecognitionTask */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for SFSpeechRecognitionTask */

// Cancels the current speech recognition task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionTask/cancel()
func (s_ SFSpeechRecognitionTask) Cancel() {
	objc.Send[objc.ID](s_.ID, objc.Sel("cancel"))
} /* debug [instance_methods/method]: Cancel */

// Stops accepting new audio and finishes processing on the audio input that has already been accepted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionTask/finish()
func (s_ SFSpeechRecognitionTask) Finish() {
	objc.Send[objc.ID](s_.ID, objc.Sel("finish"))
} /* debug [instance_methods/method]: Finish */

/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for SFSpeechRecognitionTask */

// An error object that specifies the error that occurred during a speech recognition task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionTask/error
func (s_ SFSpeechRecognitionTask) Error() objc.IObject /* cross-framework: Error */ {
	rv := objc.Send[coretelephony.Error](s_.ID, objc.Sel("error"))
	return rv
} /* debug [instance_properties/getter]: error */

// A Boolean value that indicates whether the speech recognition task was canceled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionTask/isCancelled
func (s_ SFSpeechRecognitionTask) Cancelled() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("cancelled"))
	return rv
} /* debug [instance_properties/getter]: cancelled */

// A Boolean value that indicates whether audio input has stopped.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionTask/isFinishing
func (s_ SFSpeechRecognitionTask) Finishing() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("finishing"))
	return rv
} /* debug [instance_properties/getter]: finishing */

// The current state of the speech recognition task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Speech/SFSpeechRecognitionTask/state
func (s_ SFSpeechRecognitionTask) State() SFSpeechRecognitionTaskState {
	rv := objc.Send[SFSpeechRecognitionTaskState](s_.ID, objc.Sel("state"))
	return rv
} /* debug [instance_properties/getter]: state */

// A Boolean value that indicates whether the speech recognition task was canceled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/speech/sfspeechrecognitiontask/iscancelled
func (s_ SFSpeechRecognitionTask) IsCancelled() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isCancelled"))
	return rv
} /* debug [instance_properties/getter]: isCancelled */

// A Boolean value that indicates whether the speech recognition task was canceled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/speech/sfspeechrecognitiontask/iscancelled
func (s_ SFSpeechRecognitionTask) SetIsCancelled(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsCancelled:"), value)
} /* debug [instance_properties/setter]: isCancelled */

// A Boolean value that indicates whether audio input has stopped.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/speech/sfspeechrecognitiontask/isfinishing
func (s_ SFSpeechRecognitionTask) IsFinishing() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isFinishing"))
	return rv
} /* debug [instance_properties/getter]: isFinishing */

// A Boolean value that indicates whether audio input has stopped.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/speech/sfspeechrecognitiontask/isfinishing
func (s_ SFSpeechRecognitionTask) SetIsFinishing(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsFinishing:"), value)
} /* debug [instance_properties/setter]: isFinishing */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class SFSpeechRecognitionTask */
