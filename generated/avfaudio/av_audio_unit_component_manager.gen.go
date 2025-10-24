// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/audiotoolbox"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVAudioUnitComponentManager */


/* debug [class_header]: Header for AVAudioUnitComponentManager */
// The class instance for the [AudioUnitComponentManager] class.
var (
	AudioUnitComponentManagerClass     _AudioUnitComponentManagerClass
	AudioUnitComponentManagerClassOnce sync.Once
)

func getAudioUnitComponentManagerClass() _AudioUnitComponentManagerClass {
	AudioUnitComponentManagerClassOnce.Do(func() {
		AudioUnitComponentManagerClass = _AudioUnitComponentManagerClass{objc.GetClass("AVAudioUnitComponentManager")}
	})
	return AudioUnitComponentManagerClass
}

type _AudioUnitComponentManagerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AudioUnitComponentManager */
// An interface definition for the [AudioUnitComponentManager] class.
type IAudioUnitComponentManager interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AudioUnitComponentManager */
	// properties:
	StandardLocalizedTagNames() []string
	TagNames() []string
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AudioUnitComponentManager */
	// methods:
	ComponentsMatchingPredicate(predicate foundation.Predicate) []AudioUnitComponent
	ComponentsMatchingDescription(desc audiotoolbox.AudioComponentDescription) []AudioUnitComponent
	ComponentsPassingTest(testHandler unsafe.Pointer) []AudioUnitComponent
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AudioUnitComponentManager */
// Alloc allocates a new instance without initialization.
func (ac _AudioUnitComponentManagerClass) Alloc() AudioUnitComponentManager {
	rv := objc.Send[AudioUnitComponentManager](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AudioUnitComponentManagerClass) New() AudioUnitComponentManager {
	rv := objc.Send[AudioUnitComponentManager](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioUnitComponentManager) Init() AudioUnitComponentManager {
	rv := objc.Send[AudioUnitComponentManager](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioUnitComponentManager) Autorelease() AudioUnitComponentManager {
	rv := objc.Send[AudioUnitComponentManager](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioUnitComponentManager creates a new AudioUnitComponentManager instance.
func NewAudioUnitComponentManager() AudioUnitComponentManager {
	return getAudioUnitComponentManagerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AudioUnitComponentManager */
// An object that provides a way to search and query audio components that the system registers.
//
// The component manager has methods to find various information about the audio components without opening them. Currently, you can only search audio components that are audio units. The class supports system tags and arbitrary user tags. You can tag each audio unit as part of its definition. Audio unit hosts, such as Logic or GarageBand, can present groupings of audio units according to the tags. You can search for audio units in the following ways: Using a instance that contains search strings for tags or descriptions Using a block to match on a custom criteria Using an


// An object that provides a way to search and query audio components that the system registers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitComponentManager
type AudioUnitComponentManager struct {
	objectivec.Object
}

// AudioUnitComponentManagerFrom constructs a [AudioUnitComponentManager] from an unsafe.Pointer.
//
// An object that provides a way to search and query audio components that the system registers.
func AudioUnitComponentManagerFrom(ptr unsafe.Pointer) AudioUnitComponentManager {
	return AudioUnitComponentManager{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AudioUnitComponentManager *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AudioUnitComponentManager */

// Gets the shared component manager instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitComponentManager/shared()
func (ac _AudioUnitComponentManagerClass) SharedAudioUnitComponentManager() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ac.class), objc.Sel("sharedAudioUnitComponentManager"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SharedAudioUnitComponentManager) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AudioUnitComponentManager */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AudioUnitComponentManager */

// Gets an array of audio component objects that match the search predicate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitComponentManager/components(matching:)-96l2c
func (a_ AudioUnitComponentManager) ComponentsMatchingPredicate(predicate foundation.Predicate) []AudioUnitComponent {
	rv := objc.Send[[]AudioUnitComponent](a_.ID, objc.Sel("componentsMatchingPredicate:"), predicate)
	return rv
}/* debug [instance_methods/method]: ComponentsMatchingPredicate */


// Gets an array of audio component objects that match the description.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitComponentManager/components(matching:)-9qt94
func (a_ AudioUnitComponentManager) ComponentsMatchingDescription(desc audiotoolbox.AudioComponentDescription) []AudioUnitComponent {
	rv := objc.Send[[]AudioUnitComponent](a_.ID, objc.Sel("componentsMatchingDescription:"), desc)
	return rv
}/* debug [instance_methods/method]: ComponentsMatchingDescription */


// Gets an array of audio components that pass the block method.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitComponentManager/components(passingTest:)
func (a_ AudioUnitComponentManager) ComponentsPassingTest(testHandler unsafe.Pointer) []AudioUnitComponent {
	rv := objc.Send[[]AudioUnitComponent](a_.ID, objc.Sel("componentsPassingTest:"), testHandler)
	return rv
}/* debug [instance_methods/method]: ComponentsPassingTest */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AudioUnitComponentManager */

// An array of the localized standard system tags the audio units define.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitComponentManager/standardLocalizedTagNames
func (a_ AudioUnitComponentManager) StandardLocalizedTagNames() []string {
	rv := objc.Send[[]string](a_.ID, objc.Sel("standardLocalizedTagNames"))
	return rv
}/* debug [instance_properties/getter]: standardLocalizedTagNames */


// An array of all tags the audio unit associates with the current user, and the system tags the audio units define.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitComponentManager/tagNames
func (a_ AudioUnitComponentManager) TagNames() []string {
	rv := objc.Send[[]string](a_.ID, objc.Sel("tagNames"))
	return rv
}/* debug [instance_properties/getter]: tagNames */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVAudioUnitComponentManager */



