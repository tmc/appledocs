// Code generated from Apple documentation for CoreAudio. DO NOT EDIT.

package coreaudio


// Enum types and constants

// AudioDeviceClockAlgorithmSelector enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioDeviceClockAlgorithmSelector
type AudioDeviceClockAlgorithmSelector uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioDeviceClockAlgorithmSelector/kAudioDeviceClockAlgorithm12PtMovingWindowAverage
	kAudioDeviceClockAlgorithm12PtMovingWindowAverage AudioDeviceClockAlgorithmSelector = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioDeviceClockAlgorithmSelector/kAudioDeviceClockAlgorithmRaw
	kAudioDeviceClockAlgorithmRaw AudioDeviceClockAlgorithmSelector = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioDeviceClockAlgorithmSelector/kAudioDeviceClockAlgorithmSimpleIIR
	kAudioDeviceClockAlgorithmSimpleIIR AudioDeviceClockAlgorithmSelector = 0
)


// AudioHardwarePowerHint enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioHardwarePowerHint
type AudioHardwarePowerHint uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioHardwarePowerHint/favorSavingPower
	kAudioHardwarePowerHintFavorSavingPower AudioHardwarePowerHint = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioHardwarePowerHint/none
	kAudioHardwarePowerHintNone AudioHardwarePowerHint = 0
)


// AudioLevelControlTransferFunction enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioLevelControlTransferFunction
type AudioLevelControlTransferFunction uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioLevelControlTransferFunction/tranferFunction10Over1
	kAudioLevelControlTranferFunction10Over1 AudioLevelControlTransferFunction = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioLevelControlTransferFunction/tranferFunction11Over1
	kAudioLevelControlTranferFunction11Over1 AudioLevelControlTransferFunction = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioLevelControlTransferFunction/tranferFunction12Over1
	kAudioLevelControlTranferFunction12Over1 AudioLevelControlTransferFunction = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioLevelControlTransferFunction/tranferFunction1Over2
	kAudioLevelControlTranferFunction1Over2 AudioLevelControlTransferFunction = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioLevelControlTransferFunction/tranferFunction1Over3
	kAudioLevelControlTranferFunction1Over3 AudioLevelControlTransferFunction = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioLevelControlTransferFunction/tranferFunction2Over1
	kAudioLevelControlTranferFunction2Over1 AudioLevelControlTransferFunction = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioLevelControlTransferFunction/tranferFunction3Over1
	kAudioLevelControlTranferFunction3Over1 AudioLevelControlTransferFunction = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioLevelControlTransferFunction/tranferFunction3Over2
	kAudioLevelControlTranferFunction3Over2 AudioLevelControlTransferFunction = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioLevelControlTransferFunction/tranferFunction3Over4
	kAudioLevelControlTranferFunction3Over4 AudioLevelControlTransferFunction = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioLevelControlTransferFunction/tranferFunction4Over1
	kAudioLevelControlTranferFunction4Over1 AudioLevelControlTransferFunction = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioLevelControlTransferFunction/tranferFunction5Over1
	kAudioLevelControlTranferFunction5Over1 AudioLevelControlTransferFunction = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioLevelControlTransferFunction/tranferFunction6Over1
	kAudioLevelControlTranferFunction6Over1 AudioLevelControlTransferFunction = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioLevelControlTransferFunction/tranferFunction7Over1
	kAudioLevelControlTranferFunction7Over1 AudioLevelControlTransferFunction = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioLevelControlTransferFunction/tranferFunction8Over1
	kAudioLevelControlTranferFunction8Over1 AudioLevelControlTransferFunction = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioLevelControlTransferFunction/tranferFunction9Over1
	kAudioLevelControlTranferFunction9Over1 AudioLevelControlTransferFunction = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioLevelControlTransferFunction/tranferFunctionLinear
	kAudioLevelControlTranferFunctionLinear AudioLevelControlTransferFunction = 0
)


// AudioServerPlugInIOOperation enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioServerPlugInIOOperation
type AudioServerPlugInIOOperation uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioServerPlugInIOOperation/kAudioServerPlugInIOOperationConvertInput
	kAudioServerPlugInIOOperationConvertInput AudioServerPlugInIOOperation = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioServerPlugInIOOperation/kAudioServerPlugInIOOperationConvertMix
	kAudioServerPlugInIOOperationConvertMix AudioServerPlugInIOOperation = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioServerPlugInIOOperation/kAudioServerPlugInIOOperationCycle
	kAudioServerPlugInIOOperationCycle AudioServerPlugInIOOperation = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioServerPlugInIOOperation/kAudioServerPlugInIOOperationMixOutput
	kAudioServerPlugInIOOperationMixOutput AudioServerPlugInIOOperation = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioServerPlugInIOOperation/kAudioServerPlugInIOOperationProcessInput
	kAudioServerPlugInIOOperationProcessInput AudioServerPlugInIOOperation = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioServerPlugInIOOperation/kAudioServerPlugInIOOperationProcessMix
	kAudioServerPlugInIOOperationProcessMix AudioServerPlugInIOOperation = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioServerPlugInIOOperation/kAudioServerPlugInIOOperationProcessOutput
	kAudioServerPlugInIOOperationProcessOutput AudioServerPlugInIOOperation = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioServerPlugInIOOperation/kAudioServerPlugInIOOperationReadInput
	kAudioServerPlugInIOOperationReadInput AudioServerPlugInIOOperation = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioServerPlugInIOOperation/kAudioServerPlugInIOOperationThread
	kAudioServerPlugInIOOperationThread AudioServerPlugInIOOperation = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioServerPlugInIOOperation/kAudioServerPlugInIOOperationWriteMix
	kAudioServerPlugInIOOperationWriteMix AudioServerPlugInIOOperation = 0
)


// TapMuteBehavior enum type
//
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/CATapMuteBehavior
type TapMuteBehavior uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/CATapMuteBehavior/muted
	TapMuted TapMuteBehavior = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/CATapMuteBehavior/mutedWhenTapped
	TapMutedWhenTapped TapMuteBehavior = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/CATapMuteBehavior/unmuted
	TapUnmuted TapMuteBehavior = 0
)


