module github.com/tmc/appledocs/examples/screencapturekit-darwinkit

go 1.24.1

require github.com/progrium/darwinkit v0.5.0

require github.com/ebitengine/purego v0.9.0 // indirect

replace (
	github.com/progrium/darwinkit => /Volumes/tmc/go/src/github.com/progrium/darwinkit
	github.com/tmc/appledocs => ../..
)
