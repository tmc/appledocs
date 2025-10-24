// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [AudioCompressedBuffer] class.
var (
	AudioCompressedBufferClass     _AudioCompressedBufferClass
	AudioCompressedBufferClassOnce sync.Once
)

func getAudioCompressedBufferClass() _AudioCompressedBufferClass {
	AudioCompressedBufferClassOnce.Do(func() {
		AudioCompressedBufferClass = _AudioCompressedBufferClass{objc.GetClass("AVAudioCompressedBuffer")}
	})
	return AudioCompressedBufferClass
}

type _AudioCompressedBufferClass struct {
	class objc.Class
}





// An interface definition for the [AudioCompressedBuffer] class.
type IAudioCompressedBuffer interface {
	IAudioBuffer
	

	// properties:
	ByteCapacity() uint32 /* not a class type */
	ByteLength() uint32 /* not a class type */
	SetByteLength(value uint32 /* not a class type */)
	Data() objectivec.IObject
	MaximumPacketSize() int
	PacketCapacity() AudioPacketCount /* typedef */
	PacketCount() AudioPacketCount /* typedef */
	SetPacketCount(value AudioPacketCount /* typedef */)
	PacketDependencies() objc.IObject
	PacketDescriptions() objc.IObject


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (ac _AudioCompressedBufferClass) Alloc() AudioCompressedBuffer {
	rv := objc.Send[AudioCompressedBuffer](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AudioCompressedBufferClass) New() AudioCompressedBuffer {
	rv := objc.Send[AudioCompressedBuffer](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AudioCompressedBuffer) Init() AudioCompressedBuffer {
	rv := objc.Send[AudioCompressedBuffer](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AudioCompressedBuffer) Autorelease() AudioCompressedBuffer {
	rv := objc.Send[AudioCompressedBuffer](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAudioCompressedBuffer creates a new AudioCompressedBuffer instance.
func NewAudioCompressedBuffer() AudioCompressedBuffer {
	return getAudioCompressedBufferClass().New()
}





// An object that represents an audio buffer that you use for compressed audio formats.


// An object that represents an audio buffer that you use for compressed audio formats.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioCompressedBuffer
type AudioCompressedBuffer struct {
	AudioBuffer
}

// AudioCompressedBufferFrom constructs a [AudioCompressedBuffer] from an unsafe.Pointer.
//
// An object that represents an audio buffer that you use for compressed audio formats.
func AudioCompressedBufferFrom(ptr unsafe.Pointer) AudioCompressedBuffer {
	return AudioCompressedBuffer{
		AudioBuffer: AudioBufferFrom(ptr),
	}
}






// Creates a buffer that contains constant bytes per packet of audio data in a compressed state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioCompressedBuffer/init(format:packetCapacity:)
func NewAudioCompressedBufferWithFormatPacketCapacity(format IAVAudioFormat, packetCapacity AudioPacketCount /* typedef */) AudioCompressedBuffer {
	instance := getAudioCompressedBufferClass().Alloc()
	rv := objc.Send[AudioCompressedBuffer](instance.ID, objc.Sel("initWithFormat:packetCapacity:"), format, packetCapacity)
	rv.Autorelease()
	return rv
}


// Creates a buffer that contains audio data in a compressed state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioCompressedBuffer/init(format:packetCapacity:maximumPacketSize:)
func NewAudioCompressedBufferWithFormatPacketCapacityMaximumPacketSize(format IAVAudioFormat, packetCapacity AudioPacketCount /* typedef */, maximumPacketSize int) AudioCompressedBuffer {
	instance := getAudioCompressedBufferClass().Alloc()
	rv := objc.Send[AudioCompressedBuffer](instance.ID, objc.Sel("initWithFormat:packetCapacity:maximumPacketSize:"), format, packetCapacity, maximumPacketSize)
	rv.Autorelease()
	return rv
}






















// The number of packets the buffer contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioCompressedBuffer/byteCapacity
func (a_ AudioCompressedBuffer) ByteCapacity() uint32 /* not a class type */ {
	rv := objc.Send[uint32](a_.ID, objc.Sel("byteCapacity"))
	return rv
}


// The number of valid bytes in the buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioCompressedBuffer/byteLength
func (a_ AudioCompressedBuffer) ByteLength() uint32 /* not a class type */ {
	rv := objc.Send[uint32](a_.ID, objc.Sel("byteLength"))
	return rv
}


// The number of valid bytes in the buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioCompressedBuffer/byteLength
func (a_ AudioCompressedBuffer) SetByteLength(value uint32 /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setByteLength:"), value)
}


// The audio buffer’s data bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioCompressedBuffer/data
func (a_ AudioCompressedBuffer) Data() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](a_.ID, objc.Sel("data"))
	return rv
}


// The maximum size of a packet, in bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioCompressedBuffer/maximumPacketSize
func (a_ AudioCompressedBuffer) MaximumPacketSize() int {
	rv := objc.Send[int](a_.ID, objc.Sel("maximumPacketSize"))
	return rv
}


// The total number of packets that the buffer can contain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioCompressedBuffer/packetCapacity
func (a_ AudioCompressedBuffer) PacketCapacity() AudioPacketCount /* typedef */ {
	rv := objc.Send[uint32](a_.ID, objc.Sel("packetCapacity"))
	return rv
}


// The number of packets currently in the buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioCompressedBuffer/packetCount
func (a_ AudioCompressedBuffer) PacketCount() AudioPacketCount /* typedef */ {
	rv := objc.Send[uint32](a_.ID, objc.Sel("packetCount"))
	return rv
}


// The number of packets currently in the buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioCompressedBuffer/packetCount
func (a_ AudioCompressedBuffer) SetPacketCount(value AudioPacketCount /* typedef */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPacketCount:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioCompressedBuffer/packetDependencies-5oae6
func (a_ AudioCompressedBuffer) PacketDependencies() objc.IObject {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("packetDependencies"))
	return rv
}


// The buffer’s array of packet descriptions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioCompressedBuffer/packetDescriptions
func (a_ AudioCompressedBuffer) PacketDescriptions() objc.IObject {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("packetDescriptions"))
	return rv
}







