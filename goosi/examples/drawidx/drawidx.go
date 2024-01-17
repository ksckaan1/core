// Copyright 2023 Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"embed"
	"fmt"
	"image"
	"time"
	"unsafe"

	vk "github.com/goki/vulkan"

	"cogentcore.org/core/events"
	"cogentcore.org/core/goosi"
	_ "cogentcore.org/core/goosi/driver"
	"cogentcore.org/core/mat32"
	"cogentcore.org/core/vgpu"
)

//go:embed *.spv
var content embed.FS

type CamView struct {
	Model mat32.Mat4
	View  mat32.Mat4
	Prjn  mat32.Mat4
}

func main() {
	opts := &goosi.NewWindowOptions{
		Size:      image.Pt(1024, 768),
		StdPixels: true,
		Title:     "Goosi Test Window",
	}
	w, err := goosi.TheApp.NewWindow(opts)
	if err != nil {
		panic(err)
	}
	// w.SetFPS(20) // 60 default

	var sf *vgpu.Surface
	var sy *vgpu.System
	var pl *vgpu.Pipeline
	var cam *vgpu.Val
	var camo CamView

	make := func() {

		// note: drawer is always created and ready to go
		// we are creating an additional rendering system here.
		sf = w.Drawer().Surface().(*vgpu.Surface)
		sy = sf.GPU.NewGraphicsSystem("drawidx", &sf.Device)

		destroy := func() {
			sy.Destroy()
		}
		w.SetDestroyGPUResourcesFunc(destroy)

		pl = sy.NewPipeline("drawidx")
		// sf.Format.SetMultisample(1)
		sy.ConfigRender(&sf.Format, vgpu.Depth32)
		sf.SetRender(&sy.Render)
		sy.SetClearColor(0.2, 0.2, 0.2, 1)
		sy.SetRasterization(vk.PolygonModeFill, vk.CullModeNone, vk.FrontFaceCounterClockwise, 1.0)

		pl.AddShaderEmbed("indexed", vgpu.VertexShader, content, "indexed.spv")
		pl.AddShaderEmbed("vtxcolor", vgpu.FragmentShader, content, "vtxcolor.spv")

		vars := sy.Vars()
		vset := vars.AddVertexSet()
		set := vars.AddSet()

		nPts := 3

		posv := vset.Add("Pos", vgpu.Float32Vec3, nPts, vgpu.Vertex, vgpu.VertexShader)
		clrv := vset.Add("Color", vgpu.Float32Vec3, nPts, vgpu.Vertex, vgpu.VertexShader)
		// note: always put indexes last so there isn't a gap in the location indexes!
		// just the fact of adding one (and only one) Index type triggers indexed render
		idxv := vset.Add("Index", vgpu.Uint16, nPts, vgpu.Index, vgpu.VertexShader)

		camv := set.Add("Camera", vgpu.Struct, 1, vgpu.Uniform, vgpu.VertexShader)
		camv.SizeOf = vgpu.Float32Mat4.Bytes() * 3 // no padding for these

		vset.ConfigVals(1) // one val per var
		set.ConfigVals(1)  // one val per var
		sy.Config()

		triPos, _ := posv.Vals.ValByIdxTry(0)
		triPosA := triPos.Floats32()
		triPosA.Set(0,
			-0.5, 0.5, 0.0,
			0.5, 0.5, 0.0,
			0.0, -0.5, 0.0) // negative point is UP in native Vulkan
		triPos.SetMod()

		triClr, _ := clrv.Vals.ValByIdxTry(0)
		triClrA := triClr.Floats32()
		triClrA.Set(0,
			1.0, 0.0, 0.0,
			0.0, 1.0, 0.0,
			0.0, 0.0, 1.0)
		triClr.SetMod()

		triIdx, _ := idxv.Vals.ValByIdxTry(0)
		idxs := []uint16{0, 1, 2}
		triIdx.CopyFromBytes(unsafe.Pointer(&idxs[0]))

		// This is the standard camera view projection computation
		cam, _ = camv.Vals.ValByIdxTry(0)
		campos := mat32.V3(0, 0, 2)
		target := mat32.V3(0, 0, 0)
		var lookq mat32.Quat
		lookq.SetFromRotationMatrix(mat32.NewLookAt(campos, target, mat32.V3(0, 1, 0)))
		scale := mat32.V3(1, 1, 1)
		var cview mat32.Mat4
		cview.SetTransform(campos, lookq, scale)
		view, _ := cview.Inverse()

		camo.Model.SetIdentity()
		camo.View.CopyFrom(view)
		aspect := float32(sf.Format.Size.X) / float32(sf.Format.Size.Y)
		fmt.Printf("aspect: %g\n", aspect)
		// VkPerspective version automatically flips Y axis and shifts depth
		// into a 0..1 range instead of -1..1, so original GL based geometry
		// will render identically here.
		camo.Prjn.SetVkPerspective(45, aspect, 0.01, 100)

		cam.CopyFromBytes(unsafe.Pointer(&camo)) // sets mod

		sy.Mem.SyncToGPU()

		vars.BindDynVal(0, camv, cam)
	}

	frameCount := 0
	stTime := time.Now()

	renderFrame := func() {
		if sf == nil {
			make()
		}
		// fmt.Printf("frame: %d\n", frameCount)
		// rt := time.Now()
		camo.Model.SetRotationY(.1 * float32(frameCount))
		cam.CopyFromBytes(unsafe.Pointer(&camo)) // sets mod
		sy.Mem.SyncToGPU()

		idx := sf.AcquireNextImage()
		// fmt.Printf("\nacq: %v\n", time.Now().Sub(rt))
		descIdx := 0 // if running multiple frames in parallel, need diff sets
		cmd := sy.CmdPool.Buff
		sy.ResetBeginRenderPass(cmd, sf.Frames[idx], descIdx)
		pl.BindDrawVertex(cmd, descIdx)
		sy.EndRenderPass(cmd)
		// fmt.Printf("cmd %v\n", time.Now().Sub(rt))
		sf.SubmitRender(cmd) // this is where it waits for the 16 msec
		// fmt.Printf("submit %v\n", time.Now().Sub(rt))
		sf.PresentImage(idx)
		// fmt.Printf("present %v\n\n", time.Now().Sub(rt))
		frameCount++
		eTime := time.Now()
		dur := float64(eTime.Sub(stTime)) / float64(time.Second)
		if dur > 10 {
			fps := float64(frameCount) / dur
			fmt.Printf("fps: %.0f\n", fps)
			frameCount = 0
			stTime = eTime
		}
	}

	haveKeyboard := false
	go func() {
		for {
			evi := w.EventMgr().Deque.NextEvent()
			et := evi.Type()
			if et != events.WindowPaint && et != events.MouseMove {
				fmt.Println("got event", evi)
			}
			switch et {
			case events.Window:
				ev := evi.(*events.WindowEvent)
				fmt.Println("got window event", ev)
				switch ev.Action {
				case events.WinShow:
					make()
				case events.WinClose:
					fmt.Println("got events.Close; quitting")
					goosi.TheApp.Quit()
				}
			case events.WindowPaint:
				if w.IsVisible() {
					renderFrame()
				} else {
					fmt.Println("skipping paint event")
				}
			case events.MouseDown:
				if haveKeyboard {
					goosi.TheApp.HideVirtualKeyboard()
				} else {
					goosi.TheApp.ShowVirtualKeyboard(goosi.DefaultKeyboard)
				}
				haveKeyboard = !haveKeyboard
			}
		}
	}()
	goosi.TheApp.MainLoop()
}
