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
	MData           unsafe.Pointer // A pointer to a buffer of audio data.
	MDataByteSize   unsafe.Pointer // The number of bytes in the buffer.
	MNumberChannels unsafe.Pointer // The number of interleaved channels in the buffer.
} /* debug [types.gen.go/struct]: AudioBuffer */

// AudioBufferList - A structure that stores a variable-length array of audio buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioBufferList
type AudioBufferList struct {
	MBuffers       AudioBuffer    // A variable-length array of audio buffers.
	MNumberBuffers unsafe.Pointer // The number of audio buffers in the list.
} /* debug [types.gen.go/struct]: AudioBufferList */

// AudioChannelDescription - A structure that describes a channel of audio data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioChannelDescription
type AudioChannelDescription struct {
	MChannelFlags AudioChannelFlags // The audio channel flags that indicate how to interpret the channel coordinates.
	MChannelLabel AudioChannelLabel // A label that describes the audio channel.
	MCoordinates  unsafe.Pointer    // The coordinates that specify a precise speaker location.
} /* debug [types.gen.go/struct]: AudioChannelDescription */

// AudioChannelLayout - A structure that specifies a channel layout in a file or in hardware.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioChannelLayout
type AudioChannelLayout struct {
	MChannelBitmap             AudioChannelBitmap      // If   is set to  , this field is the channel-use bitmap.
	MChannelDescriptions       AudioChannelDescription // A variable length array of   elements that describes a layout. If the   field is set to  , use this field to describe the layout.
	MChannelLayoutTag          AudioChannelLayoutTag   // The   value that indicates the layout. See   for possible values.
	MNumberChannelDescriptions unsafe.Pointer          // The number of items in the   array.
} /* debug [types.gen.go/struct]: AudioChannelLayout */

// AudioClassDescription - A structure that describes an audio codec.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioClassDescription
type AudioClassDescription struct {
	MManufacturer unsafe.Pointer // A four character code that identifies a codec manufacturer.
	MSubType      unsafe.Pointer // A four character code that a manufacturer defines for a codec subtype.
	MType         unsafe.Pointer // A four character code that a manufacturer defines for a codec type.
} /* debug [types.gen.go/struct]: AudioClassDescription */

// AudioFormatListItem
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioFormatListItem
type AudioFormatListItem struct {
	MASBD             AudioStreamBasicDescription
	MChannelLayoutTag AudioChannelLayoutTag
} /* debug [types.gen.go/struct]: AudioFormatListItem */

// AudioStreamBasicDescription - A format specification for an audio stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioStreamBasicDescription
type AudioStreamBasicDescription struct {
	MBitsPerChannel   unsafe.Pointer   // The number of bits for one audio sample.
	MBytesPerFrame    unsafe.Pointer   // The number of bytes from the start of one frame to the start of the next frame in an audio buffer.
	MBytesPerPacket   unsafe.Pointer   // The number of bytes in a packet of audio data.
	MChannelsPerFrame unsafe.Pointer   // The number of channels in each frame of audio data.
	MFormatFlags      AudioFormatFlags // Format-specific flags to specify details of the format.
	MFormatID         AudioFormatID    // An identifier specifying the general audio data format in the stream.
	MFramesPerPacket  unsafe.Pointer   // The number of frames in a packet of audio data.
	MReserved         unsafe.Pointer   // The amount to pad the structure to force an even 8-byte alignment.
	MSampleRate       unsafe.Pointer   // The number of frames per second of the data in the stream, when playing the stream at normal speed.
} /* debug [types.gen.go/struct]: AudioStreamBasicDescription */

// AudioStreamPacketDependencyDescription
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioStreamPacketDependencyDescription
type AudioStreamPacketDependencyDescription struct {
	MFlags                    unsafe.Pointer
	MIsIndependentlyDecodable unsafe.Pointer
	MPreRollCount             unsafe.Pointer
	MReserved                 unsafe.Pointer
} /* debug [types.gen.go/struct]: AudioStreamPacketDependencyDescription */

// AudioStreamPacketDescription - A value that describes a packet in a buffer of audio data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioStreamPacketDescription
type AudioStreamPacketDescription struct {
	MDataByteSize           unsafe.Pointer // The number of bytes in the packet.
	MStartOffset            unsafe.Pointer // The number of bytes from the start of the buffer to the beginning of the packet.
	MVariableFramesInPacket unsafe.Pointer // The number of sample frames of data in the packet.
} /* debug [types.gen.go/struct]: AudioStreamPacketDescription */

// AudioTimeStamp - A structure that represents a timestamp value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioTimeStamp
type AudioTimeStamp struct {
	MFlags         AudioTimeStampFlags // A set of flags indicating which representations of the time are valid; see   and  .
	MHostTime      unsafe.Pointer      // The host machine’s time base (see  ).
	MRateScalar    unsafe.Pointer      // The ratio of actual host ticks per sample frame to the nominal host ticks per sample frame.
	MReserved      unsafe.Pointer      // Pads the structure out to force an even 8-byte alignment.
	MSampleTime    unsafe.Pointer      // The absolute sample frame time.
	MSMPTETime     PTETime             // The SMPTE time (see  ).
	MWordClockTime unsafe.Pointer      // The word clock time.
} /* debug [types.gen.go/struct]: AudioTimeStamp */

// AudioValueRange - A structure that represents a continuous range of values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioValueRange
type AudioValueRange struct {
	MMaximum unsafe.Pointer // The maximum value.
	MMinimum unsafe.Pointer // The minimum value.
} /* debug [types.gen.go/struct]: AudioValueRange */

// AudioValueTranslation - A structure that stores buffers to use in translation operations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/AudioValueTranslation
type AudioValueTranslation struct {
	MInputData      unsafe.Pointer // The buffer containing the data to be translated.
	MInputDataSize  unsafe.Pointer // The number of bytes in the buffer pointed at by  .
	MOutputData     unsafe.Pointer // The buffer to hold the result of the translation.
	MOutputDataSize unsafe.Pointer // The number of bytes in the buffer pointed at by  .
} /* debug [types.gen.go/struct]: AudioValueTranslation */

// SMPTETime - A structure that defines an SMPTE time value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudioTypes/SMPTETime
type SMPTETime struct {
	MCounter         unsafe.Pointer // The total number of messages received. It takes 8 messages to carry a full SMPTE time code.
	MFlags           PTETimeFlags   // A set of flags that indicate the SMPTE state (see  ).
	MFrames          unsafe.Pointer // The value of the frames portion of the SMPTE time.
	MHours           unsafe.Pointer // The value of the hours portion of the SMPTE time.
	MMinutes         unsafe.Pointer // The value of the minutes portion of the SMPTE time.
	MSeconds         unsafe.Pointer // The value of the seconds portion of the SMPTE time.
	MSubframeDivisor unsafe.Pointer // The number of subframes per video frame (typically 80).
	MSubframes       unsafe.Pointer // A subframe offset to the HH:MM:SS:FF time. You can use this field to position a time marker somewhere within the time span represented by a video frame, if necessary.
	MType            PTETimeType    // A SMPTE time type constant indicating the kind of SMPTE time used (see  ).
} /* debug [types.gen.go/struct]: SMPTETime */
