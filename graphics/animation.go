package graphics

import (
	"github.com/KeKsBoTer/graphix/utils"
	"math/rand"
)

// The PlayMode defines in which order the frames of an Animation are displayed
// Possible options are:
//  - Normal: 		The Animation is played once in the order of given list of images
//  - Reversed:		The Animation is played once in the reversed order of the list of images
//  - Loop:			Like Normal, but the Animation is played infinite times
//  - LoopReversed: Like Reversed, but the Animation is played infinite times
//  - LoopPingPong: The Animation is played alternately Normal and Reversed
//  - LoopRandom:	The images are shown in a random order
type PlayMode int

const (
	Normal       PlayMode = iota
	Reversed
	Loop
	LoopReversed
	LoopPingPong
	Random
)

// An Animation holds a list off TextureRegions and settings, which can be used to render an Animation elegant.
// Components:
//  - regions:  	 An array of TextureRegions. Every TextureRegion represents a frame in the animation.
//  - duration: 	 The total amount of time the Animation takes to play all frames (order depends on PlayMode)
//  - perFrame: 	 The amount of time, that on frame is visible
//  - mode:			 The PlayMode defines in which order the frames are displayed
//  - lastFrame:	 The last frame that was played in the Animation (only used for the random PlayMode)
//  - lastStateTime: The last time, the Animation was updated
//
// How to use:
//  - Create a new animation with the function NewAnimation(...)
//  - To render the animation call .GetRegion(...) every frame
type Animation struct {
	regions            []TextureRegion
	duration, perFrame float64
	mode               PlayMode
	lastFrame          int
	lastStateTime      float64
}

// The constructor for a new animation.
// The function creates a new animation with the given parameters and returns a pointer the created animation.
// Arguments:
//  - regions: An array of TextureRegions, where every item represents on frame in the animation
//  - mode: The PlayMode, which is used to play the animation
//  - duration: This value is used to calculate the amount of time every frame is visible (duration divided by number of frames)
func NewAnimation(regions []TextureRegion, mode PlayMode, duration float64) *Animation {
	a := Animation{
		regions:       regions,
		perFrame:      duration / float64(len(regions)),
		duration:      duration,
		mode:          mode,
		lastFrame:     0,
		lastStateTime: 0.0,
	}
	return &a
}

// The function returns the right frame for the given time
// stateTime: The current time (can be relative). The value should only increase for every call to ensure a smooth animation.
func (a *Animation) GetRegion(stateTime float64) TextureRegion {
	index := int(stateTime / a.perFrame)
	frameCount := len(a.regions)
	switch a.mode {
	case Normal:
		index = utils.Min(frameCount-1, index)
		break
	case Loop:
		index %= frameCount
		break
	case Reversed:
		index = utils.Max(frameCount-1-index, 0)
		break
	case LoopReversed:
		index = frameCount - 1 - index%frameCount
		break
	case LoopPingPong:
		index = index % ((frameCount * 2) - 2)
		if index >= frameCount {
			index = frameCount - 2 - (index - frameCount)
		}
		break
	case Random:
		lastFrameNumber := int((a.lastStateTime) / a.perFrame)
		if lastFrameNumber != index {
			index = rand.Intn(frameCount - 1)
		} else {
			index = a.lastFrame
		}
		break
	default:
		index = 0
	}

	a.lastFrame = index
	a.lastStateTime = stateTime
	return a.regions[index]
}
