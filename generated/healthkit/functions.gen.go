// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// HealthKit Functions (11 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_HKAppleSleepingBreathingDisturbancesMinimumQuantityForClassification func(unsafe.Pointer) unsafe.Pointer
	_HKAppleSleepingBreathingDisturbancesClassificationForQuantity func(unsafe.Pointer) unsafe.Pointer
	_HKAppleWalkingSteadinessClassificationForQuantity func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_HKAppleWalkingSteadinessMaximumQuantityForClassification func(unsafe.Pointer) unsafe.Pointer
	_HKAppleWalkingSteadinessMinimumQuantityForClassification func(unsafe.Pointer) unsafe.Pointer
	_HKCategoryValueSleepAnalysisAsleepValues func() unsafe.Pointer
	_HKMaximumScoreForGAD7AssessmentRisk func(unsafe.Pointer) unsafe.Pointer
	_HKMaximumScoreForPHQ9AssessmentRisk func(unsafe.Pointer) unsafe.Pointer
	_HKMinimumScoreForGAD7AssessmentRisk func(unsafe.Pointer) unsafe.Pointer
	_HKMinimumScoreForPHQ9AssessmentRisk func(unsafe.Pointer) unsafe.Pointer
	_HKStateOfMindValenceClassificationForValence func(unsafe.Pointer) unsafe.Pointer
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_HKAppleSleepingBreathingDisturbancesMinimumQuantityForClassification, lib, "HKAppleSleepingBreathingDisturbancesMinimumQuantityForClassification")
	tryRegister(&_HKAppleSleepingBreathingDisturbancesClassificationForQuantity, lib, "HKAppleSleepingBreathingDisturbancesClassificationForQuantity")
	tryRegister(&_HKAppleWalkingSteadinessClassificationForQuantity, lib, "HKAppleWalkingSteadinessClassificationForQuantity")
	tryRegister(&_HKAppleWalkingSteadinessMaximumQuantityForClassification, lib, "HKAppleWalkingSteadinessMaximumQuantityForClassification")
	tryRegister(&_HKAppleWalkingSteadinessMinimumQuantityForClassification, lib, "HKAppleWalkingSteadinessMinimumQuantityForClassification")
	tryRegister(&_HKCategoryValueSleepAnalysisAsleepValues, lib, "HKCategoryValueSleepAnalysisAsleepValues")
	tryRegister(&_HKMaximumScoreForGAD7AssessmentRisk, lib, "HKMaximumScoreForGAD7AssessmentRisk")
	tryRegister(&_HKMaximumScoreForPHQ9AssessmentRisk, lib, "HKMaximumScoreForPHQ9AssessmentRisk")
	tryRegister(&_HKMinimumScoreForGAD7AssessmentRisk, lib, "HKMinimumScoreForGAD7AssessmentRisk")
	tryRegister(&_HKMinimumScoreForPHQ9AssessmentRisk, lib, "HKMinimumScoreForPHQ9AssessmentRisk")
	tryRegister(&_HKStateOfMindValenceClassificationForValence, lib, "HKStateOfMindValenceClassificationForValence")
}

// tryRegister attempts to register a function, silently ignoring failures.
// This allows the library to load even if some symbols are missing.
func tryRegister(fn interface{}, lib uintptr, name string) {
	defer func() {
		if r := recover(); r != nil {
			// Symbol not found - function will remain nil and panic when called
			// This is expected for inline functions, macros, or version-specific APIs
		}
	}()
	purego.RegisterLibFunc(fn, lib, name)
}



// HKAppleSleepingBreathingDisturbancesMinimumQuantityForClassification is a HealthKit function. [Full Topic]
//
// Added in macOS 15.0.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAppleSleepingBreathingDisturbancesClassification/minimum
func HKAppleSleepingBreathingDisturbancesMinimumQuantityForClassification(classification unsafe.Pointer) unsafe.Pointer {
	return _HKAppleSleepingBreathingDisturbancesMinimumQuantityForClassification(classification)
	}


// HKAppleSleepingBreathingDisturbancesClassificationForQuantity is a HealthKit function. [Full Topic]
//
// Added in macOS 15.0.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAppleSleepingBreathingDisturbancesClassificationForQuantity
func HKAppleSleepingBreathingDisturbancesClassificationForQuantity(value unsafe.Pointer) unsafe.Pointer {
	return _HKAppleSleepingBreathingDisturbancesClassificationForQuantity(value)
	}


// Provides a classification for a score that measures the steadiness of the user’s gait. [Full Topic]
//
// Added in macOS 13.0.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAppleWalkingSteadinessClassificationForQuantity
func HKAppleWalkingSteadinessClassificationForQuantity(value unsafe.Pointer, classificationOut unsafe.Pointer, errorOut unsafe.Pointer) unsafe.Pointer {
	return _HKAppleWalkingSteadinessClassificationForQuantity(value, classificationOut, errorOut)
	}


// Returns the maximum score for the steadiness of the user’s gait based on the provided classification. [Full Topic]
//
// Added in macOS 13.0.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAppleWalkingSteadinessMaximumQuantityForClassification
func HKAppleWalkingSteadinessMaximumQuantityForClassification(classification unsafe.Pointer) unsafe.Pointer {
	return _HKAppleWalkingSteadinessMaximumQuantityForClassification(classification)
	}


// Returns the minimum score for the steadiness of the user’s gait based on the provided classification. [Full Topic]
//
// Added in macOS 13.0.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKAppleWalkingSteadinessMinimumQuantityForClassification
func HKAppleWalkingSteadinessMinimumQuantityForClassification(classification unsafe.Pointer) unsafe.Pointer {
	return _HKAppleWalkingSteadinessMinimumQuantityForClassification(classification)
	}


// HKCategoryValueSleepAnalysisAsleepValues is a HealthKit function. [Full Topic]
//
// Added in macOS 13.0.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKCategoryValueSleepAnalysisAsleepValues
func HKCategoryValueSleepAnalysisAsleepValues() unsafe.Pointer {
	return _HKCategoryValueSleepAnalysisAsleepValues()
	}


// Returns the upper bound of the score range for the given GAD-7 risk classification. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKMaximumScoreForGAD7AssessmentRisk
func HKMaximumScoreForGAD7AssessmentRisk(risk unsafe.Pointer) unsafe.Pointer {
	return _HKMaximumScoreForGAD7AssessmentRisk(risk)
	}


// Returns the upper bound of the score range for the given PHQ-9 risk classification. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKMaximumScoreForPHQ9AssessmentRisk
func HKMaximumScoreForPHQ9AssessmentRisk(risk unsafe.Pointer) unsafe.Pointer {
	return _HKMaximumScoreForPHQ9AssessmentRisk(risk)
	}


// Returns the lower bound of the score range for the given GAD-7 risk classification. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKMinimumScoreForGAD7AssessmentRisk
func HKMinimumScoreForGAD7AssessmentRisk(risk unsafe.Pointer) unsafe.Pointer {
	return _HKMinimumScoreForGAD7AssessmentRisk(risk)
	}


// Returns the lower bound of the score range for the given PHQ-9 risk classification. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKMinimumScoreForPHQ9AssessmentRisk
func HKMinimumScoreForPHQ9AssessmentRisk(risk unsafe.Pointer) unsafe.Pointer {
	return _HKMinimumScoreForPHQ9AssessmentRisk(risk)
	}


// HKStateOfMindValenceClassificationForValence is a HealthKit function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKStateOfMindValenceClassificationForValence
func HKStateOfMindValenceClassificationForValence(valence unsafe.Pointer) unsafe.Pointer {
	return _HKStateOfMindValenceClassificationForValence(valence)
	}




