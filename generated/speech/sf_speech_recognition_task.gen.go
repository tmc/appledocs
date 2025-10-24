// Code generated from Apple documentation for Speech. DO NOT EDIT.

package speech

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coretelephony"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [SFSpeechRecognitionTask] class.
type ISFSpeechRecognitionTask interface {
	objectivec.IObject
	// properties:
	Error() objc.IObject /* cross-framework: Error */
	SetError(value objc.IObject /* cross-framework: Error */)
	IsCancelled() bool
	SetIsCancelled(value bool)
	IsFinishing() bool
	SetIsFinishing(value bool)
	State() SFSpeechRecognitionTaskState
	SetState(value SFSpeechRecognitionTaskState)
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (sc _SFSpeechRecognitionTaskClass) Alloc() SFSpeechRecognitionTask {
	rv := objc.Send[SFSpeechRecognitionTask](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// An error object that specifies the error that occurred during a speech recognition task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/speech/sfspeechrecognitiontask/error
func (s_ SFSpeechRecognitionTask) Error() objc.IObject /* cross-framework: Error */ {
	rv := objc.Send[coretelephony.Error](s_.ID, objc.Sel("error"))
	return rv
}


// An error object that specifies the error that occurred during a speech recognition task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/speech/sfspeechrecognitiontask/error
func (s_ SFSpeechRecognitionTask) SetError(value objc.IObject /* cross-framework: Error */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setError:"), value)
}


// A Boolean value that indicates whether the speech recognition task was canceled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/speech/sfspeechrecognitiontask/iscancelled
func (s_ SFSpeechRecognitionTask) IsCancelled() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isCancelled"))
	return rv
}


// A Boolean value that indicates whether the speech recognition task was canceled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/speech/sfspeechrecognitiontask/iscancelled
func (s_ SFSpeechRecognitionTask) SetIsCancelled(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsCancelled:"), value)
}


// A Boolean value that indicates whether audio input has stopped.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/speech/sfspeechrecognitiontask/isfinishing
func (s_ SFSpeechRecognitionTask) IsFinishing() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isFinishing"))
	return rv
}


// A Boolean value that indicates whether audio input has stopped.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/speech/sfspeechrecognitiontask/isfinishing
func (s_ SFSpeechRecognitionTask) SetIsFinishing(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsFinishing:"), value)
}


// The current state of the speech recognition task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/speech/sfspeechrecognitiontask/state
func (s_ SFSpeechRecognitionTask) State() SFSpeechRecognitionTaskState {
	rv := objc.Send[SFSpeechRecognitionTaskState](s_.ID, objc.Sel("state"))
	return rv
}


// The current state of the speech recognition task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/speech/sfspeechrecognitiontask/state
func (s_ SFSpeechRecognitionTask) SetState(value SFSpeechRecognitionTaskState) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setState:"), value)
}



