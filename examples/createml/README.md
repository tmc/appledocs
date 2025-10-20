# CreateML Framework Bindings Demo

This example demonstrates that the CreateML framework can be loaded using purego-based bindings without cgo, with basic end-to-end testing.

## Overview

CreateML is Apple's machine learning framework for training custom models on macOS. It provides:
- Easy-to-use APIs for training ML models
- Support for various model types (classifiers, regressors, recommenders)
- Integration with Core ML for deployment
- Playground and Xcode integration

## Building

```bash
go build .
```

## Running

```bash
# Show framework overview
./createml

# Run end-to-end tests
./createml -e2e
```

## Features Demonstrated

1. **Framework Loading** - Demonstrates CreateML framework can be loaded via purego
2. **Basic Testing** - Simple e2e tests to verify framework accessibility
3. **No CGO Required** - Pure Go bindings without C dependencies

## E2E Tests

The `-e2e` flag runs basic tests that verify:

1. **Framework loads successfully** - CreateML framework can be loaded
2. **Package is accessible** - Generated bindings are importable

These tests ensure that the Go bindings are correctly generated and the framework is available on the system.

## CreateML Framework

CreateML enables training machine learning models without requiring deep ML expertise:

### Supported Model Types

#### Image Classification
- Train models to classify images into categories
- Transfer learning from pre-trained models
- Data augmentation support
- Validation metrics and confusion matrix

#### Text Classification
- Sentiment analysis
- Topic categorization
- Custom text classifiers
- Natural language processing

#### Tabular Data
- Regression models (predict continuous values)
- Classification models (predict categories)
- Feature engineering
- Missing data handling

#### Sound Classification
- Audio event detection
- Environmental sound classification
- Custom sound categories

#### Activity Classification
- Motion and sensor data classification
- Exercise/activity recognition
- Gesture recognition

#### Object Detection
- Detect and locate objects in images
- Bounding box predictions
- Multiple object detection

#### Style Transfer
- Artistic style application
- Custom style training

#### Recommenders
- User-item recommendations
- Content-based filtering
- Collaborative filtering

### Key Features

**Ease of Use**
- Swift-based API (high-level)
- Playground support for experimentation
- Xcode integration
- Automatic validation and metrics

**Performance**
- Hardware acceleration (GPU, Neural Engine)
- Optimized for Apple Silicon
- Efficient training algorithms
- Transfer learning support

**Integration**
- Exports to Core ML format
- iOS/macOS deployment
- Real-time inference
- On-device processing

## Machine Learning Workflow

```
┌─────────────────────────────────────┐
│      1. Prepare Training Data       │
│   (Images, Text, Tabular, Audio)    │
└────────────────┬────────────────────┘
                 │
┌────────────────┴────────────────────┐
│    2. Create ML Training Task       │
│   (MLImageClassifier, MLTextClass)  │
└────────────────┬────────────────────┘
                 │
┌────────────────┴────────────────────┐
│       3. Train Model                │
│  (CreateML handles complexity)      │
└────────────────┬────────────────────┘
                 │
┌────────────────┴────────────────────┐
│    4. Validate & Test Model         │
│   (Accuracy, Precision, Recall)     │
└────────────────┬────────────────────┘
                 │
┌────────────────┴────────────────────┐
│    5. Export to Core ML             │
│      (.mlmodel file)                │
└────────────────┬────────────────────┘
                 │
┌────────────────┴────────────────────┐
│    6. Deploy in Application         │
│  (iOS, macOS, watchOS, tvOS)        │
└─────────────────────────────────────┘
```

## Use Cases

### Image Classification
- Photo organization (cats vs dogs)
- Quality control (defect detection)
- Medical imaging assistance
- Plant/animal species identification

### Text Analysis
- Customer sentiment analysis
- Spam detection
- Document categorization
- Intent recognition for chatbots

### Recommendation Systems
- Product recommendations
- Content suggestions
- Personalized experiences
- Music/movie recommendations

### Time Series & Tabular Data
- Sales forecasting
- Stock price prediction
- Customer churn prediction
- Demand forecasting

### Sound Classification
- Environmental sound detection
- Voice command recognition
- Music genre classification
- Acoustic event detection

## Requirements

- macOS 10.14+ (Mojave or later)
- Go 1.24.1 or later
- CreateML framework (included in macOS 10.14+)
- Xcode 10+ (for full CreateML features)

## Limitations

This example demonstrates framework loading only. It does not:
- Train actual machine learning models
- Process training data
- Export Core ML models
- Perform model evaluation

### Training Real Models

To create functional ML models with CreateML, you would typically:

1. **In Swift/Xcode**
   - Use MLImageClassifier.train()
   - Use MLTextClassifier.train()
   - Use MLRegressor.train()
   - Configure training parameters

2. **Data Preparation**
   - Organize training data properly
   - Create training/validation splits
   - Format data according to model type
   - Handle data augmentation

3. **Model Training**
   - Configure hyperparameters
   - Monitor training progress
   - Evaluate validation metrics
   - Prevent overfitting

4. **Model Export**
   - Export to Core ML format (.mlmodel)
   - Add metadata and descriptions
   - Test model inference
   - Optimize for deployment

## Swift-Only Limitations

CreateML is primarily a Swift framework with limited Objective-C bridging:
- Most APIs are Swift-only (structs, protocols, generics)
- Training APIs use Swift-specific features
- Data handling uses Swift types
- Callback/completion handlers use Swift closures

For Go-based ML workflows, consider:
- Training models in Python (TensorFlow, PyTorch)
- Converting to Core ML format (coremltools)
- Using Core ML from Go for inference only
- Or using Go ML libraries (Gorgonia, GoLearn)

## CreateML vs Core ML

| Feature | CreateML | Core ML |
|---------|----------|---------|
| Purpose | Training models | Running models |
| Platform | macOS only | All Apple platforms |
| Language | Swift (mostly) | Objective-C + Swift |
| Use Case | Development | Production deployment |
| Complexity | High-level, easy | Lower-level APIs |

## Alternative ML Training

For training models accessible from Go:

### Python ML Stack
```python
# Train in Python
from sklearn import tree
model = tree.DecisionTreeClassifier()
model.fit(X_train, y_train)

# Convert to Core ML
import coremltools
coreml_model = coremltools.converters.sklearn.convert(model)
coreml_model.save('Model.mlmodel')
```

### TensorFlow/PyTorch
- Train complex neural networks
- Convert to Core ML format
- Deploy via Core ML on Apple platforms
- Inference via Core ML from Go (if needed)

## References

- [CreateML Framework](https://developer.apple.com/documentation/createml)
- [Training ML Models](https://developer.apple.com/documentation/createml/creating_an_image_classifier_model)
- [Core ML](https://developer.apple.com/documentation/coreml)
- [Core ML Tools](https://coremltools.readme.io/)
- [Create ML App](https://developer.apple.com/machine-learning/create-ml/)

## See Also

- `../../generated/createml/` - Generated CreateML bindings
- Core ML examples for model inference
