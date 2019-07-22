package model

import (
	"math"

	"github.com/go-gl/mathgl/mgl32"
	opensimplex "github.com/ojrac/opensimplex-go"
)

var (
	sim = opensimplex.NewWithSeed(0)
)

// Abs returns the absolute value
func Abs(x float32) float32 {
	return float32(math.Abs(float64(x)))
}

func Round(x float32) float32 {
	return float32(math.Round(float64(x)))
}

func Sin(x float32) float32 {
	return float32(math.Sin(float64(x)))
}

func Cos(x float32) float32 {
	return float32(math.Cos(float64(x)))
}

func Radian(angle float32) float32 {
	return mgl32.DegToRad(angle)
}

func Max(a, b float32) float32 {
	if a > b {
		return a
	}
	return b
}

func Min(a, b float32) float32 {
	if a < b {
		return a
	}
	return b
}

func Mix(a, b, factor float32) float32 {
	return a*(1-factor) + factor*b
}

func Noise2(x, y float32, octaves int, persistence, lacunarity float32) float32 {
	var (
		freq  float32 = 1
		amp   float32 = 1
		max   float32 = 1
		total         = sim.Eval2(float64(x), float64(y))
	)
	for i := 0; i < octaves; i++ {
		freq *= lacunarity
		amp *= persistence
		max += amp
		total += sim.Eval2(float64(x*freq), float64(y*freq)) * float64(amp)
	}
	return (1 + float32(total)/max) / 2
}

func Noise3(x, y, z float32, octaves int, persistence, lacunarity float32) float32 {
	var (
		freq  float32 = 1
		amp   float32 = 1
		max   float32 = 1
		total         = sim.Eval3(float64(x), float64(y), float64(z))
	)
	for i := 0; i < octaves; i++ {
		freq *= lacunarity
		amp *= persistence
		max += amp
		total += sim.Eval3(float64(x*freq), float64(y*freq), float64(z*freq)) * float64(amp)
	}
	return (1 + float32(total)/max) / 2
}
