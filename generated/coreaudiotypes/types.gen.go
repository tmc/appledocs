// Code generated from Apple documentation for CoreAudioTypes. DO NOT EDIT.

package coreaudiotypes
import (
	"unsafe"
)


// C struct types
// AudioBuffer - A structure that holds a buffer of audio data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioBuffer
type AudioBuffer struct {
	MData unsafe.Pointer // A pointer to a buffer of audio data.
	MNumberChannels unsafe.Pointer // The number of interleaved channels in the buffer.
}// AudioBufferList - A structure that stores a variable-length array of audio buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioBufferList
type AudioBufferList struct {
	MBuffers unsafe.Pointer // A variable-length array of audio buffers.
}// AudioChannelDescription - A structure that describes a channel of audio data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioChannelDescription
type AudioChannelDescription struct {
}// AudioChannelLayout - A structure that specifies a channel layout in a file or in hardware.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioChannelLayout
type AudioChannelLayout struct {
}// AudioClassDescription - A structure that describes an audio codec.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioClassDescription
type AudioClassDescription struct {
}// AudioFormatListItem
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioFormatListItem
type AudioFormatListItem struct {
	MASBD unsafe.Pointer
	MChannelLayoutTag AudioChannelLayoutTag
}// AudioStreamBasicDescription - A format specification for an audio stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioStreamBasicDescription
type AudioStreamBasicDescription struct {
	MBitsPerChannel unsafe.Pointer // The number of bits for one audio sample.
	MBytesPerFrame unsafe.Pointer // The number of bytes from the start of one frame to the start of the next frame in an audio buffer.
	MBytesPerPacket unsafe.Pointer // The number of bytes in a packet of audio data.
	MChannelsPerFrame unsafe.Pointer // The number of channels in each frame of audio data.
	MFormatFlags AudioFormatFlags // Format-specific flags to specify details of the format.
	MFormatID AudioFormatID // An identifier specifying the general audio data format in the stream.
	MFramesPerPacket unsafe.Pointer // The number of frames in a packet of audio data.
	MReserved unsafe.Pointer // The amount to pad the structure to force an even 8-byte alignment.
	MSampleRate unsafe.Pointer // The number of frames per second of the data in the stream, when playing the stream at normal speed.
}// AudioStreamPacketDependencyDescription
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioStreamPacketDependencyDescription
type AudioStreamPacketDependencyDescription struct {
	MPreRollCount unsafe.Pointer
}// AudioStreamPacketDescription - A value that describes a packet in a buffer of audio data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioStreamPacketDescription
type AudioStreamPacketDescription struct {
	MDataByteSize unsafe.Pointer // The number of bytes in the packet.
	MStartOffset unsafe.Pointer // The number of bytes from the start of the buffer to the beginning of the packet.
	MVariableFramesInPacket unsafe.Pointer // The number of sample frames of data in the packet.
}// AudioTimeStamp - A structure that represents a timestamp value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioTimeStamp
type AudioTimeStamp struct {
	MFlags unsafe.Pointer // A set of flags indicating which representations of the time are valid; see   and  .
	MHostTime unsafe.Pointer // The host machine’s time base (see  ).
	MRateScalar unsafe.Pointer // The ratio of actual host ticks per sample frame to the nominal host ticks per sample frame.
	MReserved unsafe.Pointer // Pads the structure out to force an even 8-byte alignment.
	MSMPTETime unsafe.Pointer // The SMPTE time (see  ).
	MSampleTime unsafe.Pointer // The absolute sample frame time.
	MWordClockTime unsafe.Pointer // The word clock time.
}// AudioValueRange - A structure that represents a continuous range of values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioValueRange
type AudioValueRange struct {
	MMinimum unsafe.Pointer // The minimum value.
}// AudioValueTranslation - A structure that stores buffers to use in translation operations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioValueTranslation
type AudioValueTranslation struct {
	MOutputDataSize unsafe.Pointer // The number of bytes in the buffer pointed at by  .
}// SMPTETime - A structure that defines an SMPTE time value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/SMPTETime
type SMPTETime struct {
	MFrames unsafe.Pointer // The value of the frames portion of the SMPTE time.
	MSubframeDivisor unsafe.Pointer // The number of subframes per video frame (typically 80).
	MSubframes unsafe.Pointer // A subframe offset to the HH:MM:SS:FF time. You can use this field to position a time marker somewhere within the time span represented by a video frame, if necessary.
	MType unsafe.Pointer // A SMPTE time type constant indicating the kind of SMPTE time used (see  ).
}



