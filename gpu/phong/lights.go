// Copyright 2024 Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package phong

import (
	"cogentcore.org/core/gpu"
	"cogentcore.org/core/math32"
)

// Number of different lights active.
type NLights struct {
	Ambient int32
	Dir     int32
	Point   int32
	Spot    int32
}

func (nl *NLights) Reset() {
	nl.Ambient = 0
	nl.Dir = 0
	nl.Point = 0
	nl.Spot = 0
}

// AmbientLight provides diffuse uniform lighting. Typically only one of these.
type AmbientLight struct {

	// color of light, which multiplies ambient color of materials.
	Color math32.Vector3
	pad0  float32
}

// DirLight is directional light, which is assumed to project light toward
// the origin based on its position, with no attenuation, like the Sun.
// For rendering, the position is negated and normalized to get the direction
// vector (i.e., absolute distance doesn't matter).
type DirLight struct {

	// color of light at full intensity.
	Color math32.Vector3
	pad0  float32

	// position of light vector. Think of it shining down from
	// this position toward the origin, i.e., the negation of
	// this position is the vector.
	Pos  math32.Vector3
	pad1 float32
}

// PointLight is an omnidirectional light with a position
// and associated decay factors, which divide the light
// intensity as a function of linear and quadratic distance.
// The quadratic factor dominates at longer distances.
type PointLight struct {

	// color of light a full intensity.
	Color math32.Vector3
	pad0  float32

	// position of light in world coordinates.
	Pos  math32.Vector3
	pad1 float32

	// X = Linear distance decay factor (default .1).
	// Y = Quadratic distance quadratic decay factor
	// (default .01, dominates at longer distances)
	Decay math32.Vector3
	pad2  float32
}

// Spotlight is a light with a position and direction and
// associated decay factors and angles,
// which divide the light intensity as a function of
// linear and quadratic distance.
// The quadratic factor dominates at longer distances.
type SpotLight struct {

	// color of light a full intensity.
	Color math32.Vector3
	pad0  float32

	// position of light in world coordinates.
	Pos  math32.Vector3
	pad1 float32

	// direction of light vector.
	Dir  math32.Vector3
	pad2 float32

	// X = Angular decay, in degrees (15 default)
	// Y = CutAngle is cutoff angle in degrees beyond which no light (45 default).
	// Z = LinDecay distance linear decay (.1 default)
	// W = QuadDecay distance Distance quadratic decay factor (0.001 default),
	//     which dominates at longer distances
	Decay math32.Vector4
}

// ResetNLights resets number of lights to 0 -- for reconfig
func (ph *Phong) ResetNLights() {
	ph.NLights.Reset()
}

// AddAmbientLight adds Ambient light at given position
func (ph *Phong) AddAmbientLight(color math32.Vector3) {
	ph.SetAmbientLight(int(ph.NLights.Ambient), color)
	ph.NLights.Ambient++
}

// SetAmbientLight sets Ambient light at index to given position
func (ph *Phong) SetAmbientLight(idx int, color math32.Vector3) {
	ph.Ambient[idx].Color = color
}

// AddDirLight adds directional light
func (ph *Phong) AddDirLight(color, pos math32.Vector3) {
	ph.SetDirLight(int(ph.NLights.Dir), color, pos)
	ph.NLights.Dir++
}

// SetDirLight sets directional light at given index
func (ph *Phong) SetDirLight(idx int, color, pos math32.Vector3) {
	ph.Dir[idx].Color = color
	ph.Dir[idx].Pos = pos
}

// AddPointLight adds point light.
// Defaults: linDecay=.1, quadDecay=.01
func (ph *Phong) AddPointLight(color, pos math32.Vector3, linDecay, quadDecay float32) {
	ph.SetPointLight(int(ph.NLights.Point), color, pos, linDecay, quadDecay)
	ph.NLights.Point++
}

// SetPointLight sets point light at given index.
// Defaults: linDecay=.1, quadDecay=.01
func (ph *Phong) SetPointLight(idx int, color, pos math32.Vector3, linDecay, quadDecay float32) {
	ph.Point[idx].Color = color
	ph.Point[idx].Pos = pos
	ph.Point[idx].Decay = math32.Vector3{X: linDecay, Y: quadDecay}
}

// AddSpotLight adds spot light
// Defaults: angDecay=15, cutAngle=45 (max 90), linDecay=.01, quadDecay=0.001
func (ph *Phong) AddSpotLight(color, pos, dir math32.Vector3, angDecay, cutAngle, linDecay, quadDecay float32) {
	ph.SetSpotLight(int(ph.NLights.Spot), color, pos, dir, angDecay, cutAngle, linDecay, quadDecay)
	ph.NLights.Spot++
}

// SetSpotLight sets spot light at given index
// Defaults: angDecay=15, cutAngle=45 (max 90), linDecay=.01, quadDecay=0.001
func (ph *Phong) SetSpotLight(idx int, color, pos, dir math32.Vector3, angDecay, cutAngle, linDecay, quadDecay float32) {
	ph.Spot[idx].Color = color
	ph.Spot[idx].Pos = pos
	ph.Spot[idx].Dir = dir
	ph.Spot[idx].Decay = math32.Vec4(angDecay, cutAngle, linDecay, quadDecay)
}

// configLights configures the rendering for the lights that have been added.
func (ph *Phong) configLights() {
	sy := ph.Sys
	vs := sy.Vars.Groups[int(LightGroup)]
	gpu.SetValueFrom(vs.ValueByIndex("NLights", 0), []NLights{ph.NLights})
	gpu.SetValueFrom(vs.ValueByIndex("AmbLights", 0), ph.Ambient[:])
	gpu.SetValueFrom(vs.ValueByIndex("DirLights", 0), ph.Dir[:])
	gpu.SetValueFrom(vs.ValueByIndex("PointLights", 0), ph.Point[:])
	gpu.SetValueFrom(vs.ValueByIndex("SpotLights", 0), ph.Spot[:])
}
