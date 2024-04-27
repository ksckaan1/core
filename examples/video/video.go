package main

import (
	"cogentcore.org/core/core"
	"cogentcore.org/core/gox/errors"
	"cogentcore.org/core/styles"
	"cogentcore.org/core/video"
	_ "cogentcore.org/core/views"
)

func main() {
	b := core.NewBody("Basic Video Example")
	bx := core.NewLayout(b).Style(func(s *styles.Style) {
		s.Grow.Set(1, 1)
	})
	core.NewText(bx).SetText("video:").Style(func(s *styles.Style) {
		s.SetTextWrap(false)
	})
	v := video.NewVideo(bx)
	v.Style(func(s *styles.Style) {
		s.Min.X.Px(200)
		s.Grow.Set(1, 1)
	})
	core.NewText(bx).SetText("filler:").Style(func(s *styles.Style) {
		s.SetTextWrap(false)
	})
	core.NewText(b).SetText("footer:")
	// errors.Log(v.Open("deer.mp4"))
	// errors.Log(v.Open("countdown.mp4"))
	errors.Log(v.Open("randy_first_360.mov")) // note: not uploaded -- good test case tho
	v.Rotation = -90
	w := b.RunWindow()
	v.Play(0, 0)
	w.Wait()
}
