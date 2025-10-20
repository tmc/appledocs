// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
}

// An object that represents an audio buffer that you use for compressed audio formats.
//
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

// Alloc allocates a new instance without initialization.
func (ac _AudioCompressedBufferClass) Alloc() AudioCompressedBuffer {
	rv := objc.Send[AudioCompressedBuffer](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Creates a buffer that contains constant bytes per packet of audio data in a compressed state.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioCompressedBuffer/init(format:packetCapacity:)
func NewAudioCompressedBufferWithFormatPacketCapacity(format unsafe.Pointer, packetCapacity unsafe.Pointer) AudioCompressedBuffer {
	instance := getAudioCompressedBufferClass().Alloc()
	rv := objc.Send[AudioCompressedBuffer](instance.ID, objc.Sel("initWithFormat:packetCapacity:"), format, packetCapacity)
	rv.Autorelease()
	return rv
}

// Creates a buffer that contains audio data in a compressed state.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioCompressedBuffer/init(format:packetCapacity:maximumPacketSize:)
func NewAudioCompressedBufferWithFormatPacketCapacityMaximumPacketSize(format unsafe.Pointer, packetCapacity unsafe.Pointer, maximumPacketSize int) AudioCompressedBuffer {
	instance := getAudioCompressedBufferClass().Alloc()
	rv := objc.Send[AudioCompressedBuffer](instance.ID, objc.Sel("initWithFormat:packetCapacity:maximumPacketSize:"), format, packetCapacity, maximumPacketSize)
	rv.Autorelease()
	return rv
}


// The number of packets the buffer contains.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioCompressedBuffer/byteCapacity
func (a_ AudioCompressedBuffer) ByteCapacity() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("byteCapacity"))
	return rv
}

// The number of valid bytes in the buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioCompressedBuffer/byteLength
func (a_ AudioCompressedBuffer) ByteLength() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("byteLength"))
	return rv
}


// SetByteLength sets the value of the byteLength property.
// The number of valid bytes in the buffer.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioCompressedBuffer/byteLength
func (a_ AudioCompressedBuffer) SetByteLength(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setByteLength:"), value)
}
// The maximum size of a packet, in bytes.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioCompressedBuffer/maximumPacketSize
func (a_ AudioCompressedBuffer) MaximumPacketSize() int {
	rv := objc.Send[int](a_.ID, objc.Sel("maximumPacketSize"))
	return rv
}

// The total number of packets that the buffer can contain.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioCompressedBuffer/packetCapacity
func (a_ AudioCompressedBuffer) PacketCapacity() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("packetCapacity"))
	return rv
}

// The number of packets currently in the buffer.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioCompressedBuffer/packetCount
func (a_ AudioCompressedBuffer) PacketCount() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("packetCount"))
	return rv
}


// SetPacketCount sets the value of the packetCount property.
// The number of packets currently in the buffer.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioCompressedBuffer/packetCount
func (a_ AudioCompressedBuffer) SetPacketCount(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPacketCount:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioCompressedBuffer/packetDependencies-5oae6
func (a_ AudioCompressedBuffer) PacketDependencies() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("packetDependencies"))
	return rv
}

// The buffer’s array of packet descriptions.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudioCompressedBuffer/packetDescriptions
func (a_ AudioCompressedBuffer) PacketDescriptions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("packetDescriptions"))
	return rv
}


