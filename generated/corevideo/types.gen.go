// Code generated from Apple documentation for CoreVideo. DO NOT EDIT.

package corevideo
import (
	"unsafe"
)


// C struct types
// CVFillExtendedPixelsCallBackData - A structure for holding information that describes a custom extended pixel fill algorithm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVFillExtendedPixelsCallBackData
type CVFillExtendedPixelsCallBackData struct {
	FillCallBack FillExtendedPixelsCallBack
	RefCon unsafe.Pointer // A pointer to application-defined data that is passed to your custom pixel fill function.
	Version Index // The version of this fill algorithm.
}/* debug [types.gen.go/struct]: CVFillExtendedPixelsCallBackData */

// CVPlanarComponentInfo - A structure for describing planar components.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPlanarComponentInfo
type CVPlanarComponentInfo struct {
	Offset int32 // The offset from the main base address to the base address of this plane. (big-endian)
	RowBytes uint32 // The number of bytes per row of this plane. (big-endian)
}/* debug [types.gen.go/struct]: CVPlanarComponentInfo */

// CVPlanarPixelBufferInfo - A structure for describing planar buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPlanarPixelBufferInfo
type CVPlanarPixelBufferInfo struct {
	ComponentInfo PlanarComponentInfo // An array containing a   structure for each plane of the buffer.
}/* debug [types.gen.go/struct]: CVPlanarPixelBufferInfo */

// CVPlanarPixelBufferInfo_YCbCrBiPlanar - A structure for describing YCbCr biplanar buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPlanarPixelBufferInfo_YCbCrBiPlanar
type CVPlanarPixelBufferInfo_YCbCrBiPlanar struct {
	ComponentInfoCbCr PlanarComponentInfo // A   structure containing information on the Cb/Cr component of the buffer.
	ComponentInfoY PlanarComponentInfo // A   structure containing information on the Y component of the buffer.
}/* debug [types.gen.go/struct]: CVPlanarPixelBufferInfo_YCbCrBiPlanar */

// CVPlanarPixelBufferInfo_YCbCrPlanar - A structure for describing YCbCr planar buffers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVPlanarPixelBufferInfo_YCbCrPlanar
type CVPlanarPixelBufferInfo_YCbCrPlanar struct {
	ComponentInfoCb PlanarComponentInfo // A   structure containing information on the Cb component of the buffer.
	ComponentInfoCr PlanarComponentInfo // A   structure containing information on the Cr component of the buffer.
	ComponentInfoY PlanarComponentInfo // A   structure containing information on the Y component of the buffer.
}/* debug [types.gen.go/struct]: CVPlanarPixelBufferInfo_YCbCrPlanar */

// CVSMPTETime - A structure for holding an SMPTE time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVSMPTETime
type CVSMPTETime struct {
	Counter unsafe.Pointer // The total number of messages received.
	Flags unsafe.Pointer // A set of flags that indicate the SMPTE state.
	Frames unsafe.Pointer // The number of frames in the full message.
	Hours unsafe.Pointer // The number of hours in the full message.
	Minutes unsafe.Pointer // The number of minutes in the full message.
	Seconds unsafe.Pointer // The number of seconds in the full message.
	SubframeDivisor unsafe.Pointer // The number of subframes per frame (typically, 80).
	Subframes unsafe.Pointer // The number of subframes in the full message.
	Type unsafe.Pointer // The kind of SMPTE time type.
}/* debug [types.gen.go/struct]: CVSMPTETime */

// CVTime - A structure for reporting Core Video time values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVTime
type CVTime struct {
	Flags int32 // The flags associated with the   value. See   for possible values. If   is set, you should not use any of the other fields in this structure.
	TimeScale int32 // The time scale for this value.
	TimeValue int64 // The time value.
}/* debug [types.gen.go/struct]: CVTime */

// CVTimeStamp - A structure for defining a display timestamp.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/CVTimeStamp
type CVTimeStamp struct {
	Flags uint64 // A bit field containing additional information about the timestamp.
	HostTime uint64 // The system time measured by the timestamp.
	RateScalar float64 // The current rate of the device as measured by the timestamps, divided by the nominal rate.
	Reserved uint64 // Reserved. Do not use.
	SmpteTime SMPTETime // The SMPTE time representation of the timestamp.
	Version uint32 // The current   structure is version 0. Some functions require you to specify a version when passing in a timestamp structure to be filled.
	VideoRefreshPeriod int64
	VideoTime int64 // The start of a frame (or field for interlaced video).
	VideoTimeScale int32 // The scale (in units per second) of the   and   fields.
}/* debug [types.gen.go/struct]: CVTimeStamp */





