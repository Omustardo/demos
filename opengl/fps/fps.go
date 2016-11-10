// Package fps can be used to keep track of framerate.
// Sample usage:
//
package fps

import (
	"fmt"
	"time"
)

type FPSCounter struct {
	// Smoothing is how much old framerate values affect the current FPS value. Larger smoothing mean that older
	// values count for more. Smoothing must be in the range [0,1]
	smoothing float32
	prevTime  int64
	framerate float32
}

func NewFPSCounter() *FPSCounter {
	return &FPSCounter{
		smoothing: 0.9,
		prevTime:  getTimeMillis(),
		framerate: 1.0,
	}
}

// Update is expected to be called once per frame.
func (f *FPSCounter) Update() {
	currTime := getTimeMillis()
	delta := currTime - f.prevTime
	f.framerate = (float32(delta) * f.smoothing) + (f.framerate * (1.0 - f.smoothing))
	f.prevTime = currTime
}

// GetFPS returns the estimated number of frames shown in the last second.
func (f *FPSCounter) GetFPS() float32 {
	return 1000 / f.framerate
}

// GetFramerate returns an estimate of the average length of each frame.
func (f *FPSCounter) GetFramerate() float32 {
	return f.framerate
}

// GetFPS returns the number of frames shown in the last second.
func (f *FPSCounter) GetFPSString() string {
	return fmt.Sprintf("%.f", f.GetFPS())
}

func getTimeMillis() int64 {
	return time.Now().UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
}
