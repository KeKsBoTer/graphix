package graphics

import (
	"github.com/KeKsBoTer/graphix/utils"
	"math/rand"
)

type PlayMode int

const (
	Normal       PlayMode = iota
	Reversed     PlayMode = iota
	Loop         PlayMode = iota
	LoopReversed PlayMode = iota
	LoopPingPong PlayMode = iota
	Random       PlayMode = iota
)

type Animation struct {
	regions            []TextureRegion
	duration, perFrame float64
	mode               PlayMode
	lastFrame          int
	lastStateTime      float64
}

func NewAnimation(regions []TextureRegion, mode PlayMode, duration float64) *Animation {
	a := Animation{
		regions:   regions,
		perFrame:  duration / float64(len(regions)),
		duration:  duration,
		mode:      mode,
		lastFrame: 0,
	}
	return &a
}

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
