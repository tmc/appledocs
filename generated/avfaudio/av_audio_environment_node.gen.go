// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVAudioEnvironmentNode */


/* debug [class_header]: Header for AVAudioEnvironmentNode */
// The class instance for the [AudioEnvironmentNode] class.
var (
	AudioEnvironmentNodeClass     _AudioEnvironmentNodeClass
	AudioEnvironmentNodeClassOnce sync.Once
)

func getAudioEnvironmentNodeClass() _AudioEnvironmentNodeClass {
	AudioEnvironmentNodeClassOnce.Do(func() {
		AudioEnvironmentNodeClass = _AudioEnvironmentNodeClass{objc.GetClass("AVAudioEnvironmentNode")}
	})
	return AudioEnvironmentNodeClass
}

type _AudioEnvironmentNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AudioEnvironmentNode */
// An interface definition for the [AudioEnvironmentNode] class.
type IAudioEnvironmentNode interface {
	IAudioNode
	
/* debug [class_interface_properties]: Properties for AudioEnvironmentNode */
	// properties:
	ApplicableRenderingAlgorithms() []foundation.Number
	DistanceAttenuationParameters() IAVAudioEnvironmentDistanceAttenuationParameters
	ListenerHeadTrackingEnabled() bool
	SetListenerHeadTrackingEnabled(value bool)
	ListenerAngularOrientation() objc.IObject /* cross-framework: AVAudio3DAngularOrientation */
	SetListenerAngularOrientation(value objc.IObject /* cross-framework: AVAudio3DAngularOrientation */)
	ListenerPosition() objc.IObject /* cross-framework: AVAudio3DPoint */
	SetListenerPosition(value objc.IObject /* cross-framework: AVAudio3DPoint */)
	ListenerVectorOrientation() objc.IObject /* cross-framework: AVAudio3DVectorOrientation */
	SetListenerVectorOrientation(value objc.IObject /* cross-framework: AVAudio3DVectorOrientation */)
	NextAvailableInputBus() AudioNodeBus /* typedef */
	OutputType() AudioEnvironmentOutputType
	SetOutputType(value AudioEnvironmentOutputType)
	OutputVolume() float32
	SetOutputVolume(value float32)
	ReverbParameters() IAVAudioEnvironmentReverbParameters
	IsListenerHeadTrackingEnabled() bool
	SetIsListenerHeadTrackingEnabled(value bool)
	KAudioChannelLayoutTag_AudioUnit_4() objectivec.IObject
	SetKAudioChannelLayoutTag_AudioUnit_4(value objectivec.IObject)
	KAudioChannelLayoutTag_AudioUnit_5_0() objectivec.IObject
	SetKAudioChannelLayoutTag_AudioUnit_5_0(value objectivec.IObject)
	KAudioChannelLayoutTag_AudioUnit_6_0() objectivec.IObject
	SetKAudioChannelLayoutTag_AudioUnit_6_0(value objectivec.IObject)
	KAudioChannelLayoutTag_AudioUnit_7_0() objectivec.IObject
	SetKAudioChannelLayoutTag_AudioUnit_7_0(value objectivec.IObject)
	KAudioChannelLayoutTag_AudioUnit_7_0_Front() objectivec.IObject
	SetKAudioChannelLayoutTag_AudioUnit_7_0_Front(value objectivec.IObject)
	KAudioChannelLayoutTag_AudioUnit_8() objectivec.IObject
	SetKAudioChannelLayoutTag_AudioUnit_8(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AudioEnvironmentNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AudioEnvironmentNode */
// Alloc allocates a new instance without initialization.
func (ac _AudioEnvironmentNodeClass) Alloc() AudioEnvironmentNode {
	rv := objc.Send[AudioEnvironmentNode](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AudioEnvironmentNodeClass) New() AudioEnvironmentNode {
	rv := objc.Send[AudioEnvironmentNode](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioEnvironmentNode) Init() AudioEnvironmentNode {
	rv := objc.Send[AudioEnvironmentNode](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioEnvironmentNode) Autorelease() AudioEnvironmentNode {
	rv := objc.Send[AudioEnvironmentNode](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioEnvironmentNode creates a new AudioEnvironmentNode instance.
func NewAudioEnvironmentNode() AudioEnvironmentNode {
	return getAudioEnvironmentNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AudioEnvironmentNode */
// An object that simulates a 3D audio environment.
//
// The class is a mixer node that simulates a 3D audio environment. Any node that conforms to can act as a source node, such as . The environment node has an implicit listener. You set the listener’s position and orientation, and the system then controls the way the user experiences the virtual world. To help characterize the environment, this class defines properties for distance attenuation and reverberation. affects how inputs with different channel configurations render. Spatialization applies only to inputs with a mono channel connection format. This class doesn’t spatialize stereo inputs or support inputs with connection formats of more than two channels. To set the node’s output to a multichannel format, use an that has one of the following :


// An object that simulates a 3D audio environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentNode
type AudioEnvironmentNode struct {
	AudioNode
}

// AudioEnvironmentNodeFrom constructs a [AudioEnvironmentNode] from an unsafe.Pointer.
//
// An object that simulates a 3D audio environment.
func AudioEnvironmentNodeFrom(ptr unsafe.Pointer) AudioEnvironmentNode {
	return AudioEnvironmentNode{
		AudioNode: AudioNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AudioEnvironmentNode */
/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AudioEnvironmentNode */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AudioEnvironmentNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AudioEnvironmentNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AudioEnvironmentNode */

// An array of rendering algorithms applicable to the environment node.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentNode/applicableRenderingAlgorithms
func (a_ AudioEnvironmentNode) ApplicableRenderingAlgorithms() []foundation.Number {
	rv := objc.Send[[]foundation.Number](a_.ID, objc.Sel("applicableRenderingAlgorithms"))
	return rv
}/* debug [instance_properties/getter]: applicableRenderingAlgorithms */


// The distance attenuation parameters for the environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentNode/distanceAttenuationParameters
func (a_ AudioEnvironmentNode) DistanceAttenuationParameters() IAVAudioEnvironmentDistanceAttenuationParameters {
	rv := objc.Send[AudioEnvironmentDistanceAttenuationParameters](a_.ID, objc.Sel("distanceAttenuationParameters"))
	return rv
}/* debug [instance_properties/getter]: distanceAttenuationParameters */


// A Boolean value that indicates whether the listener orientation is automatically rotated based on head orientation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentNode/isListenerHeadTrackingEnabled
func (a_ AudioEnvironmentNode) ListenerHeadTrackingEnabled() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("listenerHeadTrackingEnabled"))
	return rv
}/* debug [instance_properties/getter]: listenerHeadTrackingEnabled */


// A Boolean value that indicates whether the listener orientation is automatically rotated based on head orientation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentNode/isListenerHeadTrackingEnabled
func (a_ AudioEnvironmentNode) SetListenerHeadTrackingEnabled(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setListenerHeadTrackingEnabled:"), value)
}/* debug [instance_properties/setter]: listenerHeadTrackingEnabled */


// The listener’s angular orientation in the environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentNode/listenerAngularOrientation
func (a_ AudioEnvironmentNode) ListenerAngularOrientation() objc.IObject /* cross-framework: AVAudio3DAngularOrientation */ {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("listenerAngularOrientation"))
	return rv
}/* debug [instance_properties/getter]: listenerAngularOrientation */


// The listener’s angular orientation in the environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentNode/listenerAngularOrientation
func (a_ AudioEnvironmentNode) SetListenerAngularOrientation(value objc.IObject /* cross-framework: AVAudio3DAngularOrientation */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setListenerAngularOrientation:"), value)
}/* debug [instance_properties/setter]: listenerAngularOrientation */


// The listener’s position in the 3D environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentNode/listenerPosition
func (a_ AudioEnvironmentNode) ListenerPosition() objc.IObject /* cross-framework: AVAudio3DPoint */ {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("listenerPosition"))
	return rv
}/* debug [instance_properties/getter]: listenerPosition */


// The listener’s position in the 3D environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentNode/listenerPosition
func (a_ AudioEnvironmentNode) SetListenerPosition(value objc.IObject /* cross-framework: AVAudio3DPoint */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setListenerPosition:"), value)
}/* debug [instance_properties/setter]: listenerPosition */


// The listener’s vector orientation in the environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentNode/listenerVectorOrientation
func (a_ AudioEnvironmentNode) ListenerVectorOrientation() objc.IObject /* cross-framework: AVAudio3DVectorOrientation */ {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("listenerVectorOrientation"))
	return rv
}/* debug [instance_properties/getter]: listenerVectorOrientation */


// The listener’s vector orientation in the environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentNode/listenerVectorOrientation
func (a_ AudioEnvironmentNode) SetListenerVectorOrientation(value objc.IObject /* cross-framework: AVAudio3DVectorOrientation */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setListenerVectorOrientation:"), value)
}/* debug [instance_properties/setter]: listenerVectorOrientation */


// An unused input bus.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentNode/nextAvailableInputBus
func (a_ AudioEnvironmentNode) NextAvailableInputBus() AudioNodeBus /* typedef */ {
	rv := objc.Send[uint](a_.ID, objc.Sel("nextAvailableInputBus"))
	return rv
}/* debug [instance_properties/getter]: nextAvailableInputBus */


// The type of output hardware.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentNode/outputType
func (a_ AudioEnvironmentNode) OutputType() AudioEnvironmentOutputType {
	rv := objc.Send[AudioEnvironmentOutputType](a_.ID, objc.Sel("outputType"))
	return rv
}/* debug [instance_properties/getter]: outputType */


// The type of output hardware.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentNode/outputType
func (a_ AudioEnvironmentNode) SetOutputType(value AudioEnvironmentOutputType) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setOutputType:"), value)
}/* debug [instance_properties/setter]: outputType */


// The mixer’s output volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentNode/outputVolume
func (a_ AudioEnvironmentNode) OutputVolume() float32 {
	rv := objc.Send[float32](a_.ID, objc.Sel("outputVolume"))
	return rv
}/* debug [instance_properties/getter]: outputVolume */


// The mixer’s output volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentNode/outputVolume
func (a_ AudioEnvironmentNode) SetOutputVolume(value float32) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setOutputVolume:"), value)
}/* debug [instance_properties/setter]: outputVolume */


// The reverb parameters for the environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentNode/reverbParameters
func (a_ AudioEnvironmentNode) ReverbParameters() IAVAudioEnvironmentReverbParameters {
	rv := objc.Send[AudioEnvironmentReverbParameters](a_.ID, objc.Sel("reverbParameters"))
	return rv
}/* debug [instance_properties/getter]: reverbParameters */


// A Boolean value that indicates whether the listener orientation is automatically rotated based on head orientation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioenvironmentnode/islistenerheadtrackingenabled
func (a_ AudioEnvironmentNode) IsListenerHeadTrackingEnabled() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isListenerHeadTrackingEnabled"))
	return rv
}/* debug [instance_properties/getter]: isListenerHeadTrackingEnabled */


// A Boolean value that indicates whether the listener orientation is automatically rotated based on head orientation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfaudio/avaudioenvironmentnode/islistenerheadtrackingenabled
func (a_ AudioEnvironmentNode) SetIsListenerHeadTrackingEnabled(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsListenerHeadTrackingEnabled:"), value)
}/* debug [instance_properties/setter]: isListenerHeadTrackingEnabled */


// A quadraphonic symmetrical layout, recommended for use by audio units.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/kAudioChannelLayoutTag_AudioUnit_4
func (a_ AudioEnvironmentNode) KAudioChannelLayoutTag_AudioUnit_4() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](a_.ID, objc.Sel("kAudioChannelLayoutTag_AudioUnit_4"))
	return rv
}/* debug [instance_properties/getter]: kAudioChannelLayoutTag_AudioUnit_4 */


// A quadraphonic symmetrical layout, recommended for use by audio units.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/kAudioChannelLayoutTag_AudioUnit_4
func (a_ AudioEnvironmentNode) SetKAudioChannelLayoutTag_AudioUnit_4(value objectivec.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setKAudioChannelLayoutTag_AudioUnit_4:"), value)
}/* debug [instance_properties/setter]: kAudioChannelLayoutTag_AudioUnit_4 */


// A 5-channel surround-based layout, recommended for use by audio units.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/kAudioChannelLayoutTag_AudioUnit_5_0
func (a_ AudioEnvironmentNode) KAudioChannelLayoutTag_AudioUnit_5_0() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](a_.ID, objc.Sel("kAudioChannelLayoutTag_AudioUnit_5_0"))
	return rv
}/* debug [instance_properties/getter]: kAudioChannelLayoutTag_AudioUnit_5_0 */


// A 5-channel surround-based layout, recommended for use by audio units.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/kAudioChannelLayoutTag_AudioUnit_5_0
func (a_ AudioEnvironmentNode) SetKAudioChannelLayoutTag_AudioUnit_5_0(value objectivec.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setKAudioChannelLayoutTag_AudioUnit_5_0:"), value)
}/* debug [instance_properties/setter]: kAudioChannelLayoutTag_AudioUnit_5_0 */


// A 6-channel surround-based layout, recommended for use by audio units.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/kAudioChannelLayoutTag_AudioUnit_6_0
func (a_ AudioEnvironmentNode) KAudioChannelLayoutTag_AudioUnit_6_0() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](a_.ID, objc.Sel("kAudioChannelLayoutTag_AudioUnit_6_0"))
	return rv
}/* debug [instance_properties/getter]: kAudioChannelLayoutTag_AudioUnit_6_0 */


// A 6-channel surround-based layout, recommended for use by audio units.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/kAudioChannelLayoutTag_AudioUnit_6_0
func (a_ AudioEnvironmentNode) SetKAudioChannelLayoutTag_AudioUnit_6_0(value objectivec.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setKAudioChannelLayoutTag_AudioUnit_6_0:"), value)
}/* debug [instance_properties/setter]: kAudioChannelLayoutTag_AudioUnit_6_0 */


// A 7-channel surround-based layout, recommended for use by audio units.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/kAudioChannelLayoutTag_AudioUnit_7_0
func (a_ AudioEnvironmentNode) KAudioChannelLayoutTag_AudioUnit_7_0() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](a_.ID, objc.Sel("kAudioChannelLayoutTag_AudioUnit_7_0"))
	return rv
}/* debug [instance_properties/getter]: kAudioChannelLayoutTag_AudioUnit_7_0 */


// A 7-channel surround-based layout, recommended for use by audio units.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/kAudioChannelLayoutTag_AudioUnit_7_0
func (a_ AudioEnvironmentNode) SetKAudioChannelLayoutTag_AudioUnit_7_0(value objectivec.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setKAudioChannelLayoutTag_AudioUnit_7_0:"), value)
}/* debug [instance_properties/setter]: kAudioChannelLayoutTag_AudioUnit_7_0 */


// An alternate 7-channel surround-based layout, for use by audio units.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/kAudioChannelLayoutTag_AudioUnit_7_0_Front
func (a_ AudioEnvironmentNode) KAudioChannelLayoutTag_AudioUnit_7_0_Front() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](a_.ID, objc.Sel("kAudioChannelLayoutTag_AudioUnit_7_0_Front"))
	return rv
}/* debug [instance_properties/getter]: kAudioChannelLayoutTag_AudioUnit_7_0_Front */


// An alternate 7-channel surround-based layout, for use by audio units.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/kAudioChannelLayoutTag_AudioUnit_7_0_Front
func (a_ AudioEnvironmentNode) SetKAudioChannelLayoutTag_AudioUnit_7_0_Front(value objectivec.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setKAudioChannelLayoutTag_AudioUnit_7_0_Front:"), value)
}/* debug [instance_properties/setter]: kAudioChannelLayoutTag_AudioUnit_7_0_Front */


// An octagonal symmetrical layout, recommended for use by audio units.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/kAudioChannelLayoutTag_AudioUnit_8
func (a_ AudioEnvironmentNode) KAudioChannelLayoutTag_AudioUnit_8() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](a_.ID, objc.Sel("kAudioChannelLayoutTag_AudioUnit_8"))
	return rv
}/* debug [instance_properties/getter]: kAudioChannelLayoutTag_AudioUnit_8 */


// An octagonal symmetrical layout, recommended for use by audio units.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/kAudioChannelLayoutTag_AudioUnit_8
func (a_ AudioEnvironmentNode) SetKAudioChannelLayoutTag_AudioUnit_8(value objectivec.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setKAudioChannelLayoutTag_AudioUnit_8:"), value)
}/* debug [instance_properties/setter]: kAudioChannelLayoutTag_AudioUnit_8 */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVAudioEnvironmentNode */


