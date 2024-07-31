// Code generated by "core generate"; DO NOT EDIT.

package shape

import (
	"image/color"

	"cogentcore.org/core/math32"
	"cogentcore.org/core/types"
)

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/gpu/shape.Box", IDName: "box", Doc: "Box is a rectangular-shaped solid (cuboid)", Directives: []types.Directive{{Tool: "types", Directive: "add", Args: []string{"-setters"}}}, Embeds: []types.Field{{Name: "ShapeBase"}}, Fields: []types.Field{{Name: "Size", Doc: "size along each dimension"}, {Name: "Segs", Doc: "number of segments to divide each plane into (enforced to be at least 1) -- may potentially increase rendering quality to have > 1"}}})

// SetSize sets the [Box.Size]:
// size along each dimension
func (t *Box) SetSize(v math32.Vector3) *Box { t.Size = v; return t }

// SetSegs sets the [Box.Segs]:
// number of segments to divide each plane into (enforced to be at least 1) -- may potentially increase rendering quality to have > 1
func (t *Box) SetSegs(v math32.Vector3i) *Box { t.Segs = v; return t }

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/gpu/shape.Capsule", IDName: "capsule", Doc: "Capsule is a generalized capsule shape: a cylinder with hemisphere end caps.\nSupports different radii on each end.\nHeight is along the Y axis -- total height is Height + TopRad + BotRad.", Directives: []types.Directive{{Tool: "types", Directive: "add", Args: []string{"-setters"}}}, Embeds: []types.Field{{Name: "ShapeBase"}}, Fields: []types.Field{{Name: "Height", Doc: "height of the cylinder portion"}, {Name: "TopRad", Doc: "radius of the top -- set to 0 to omit top cap"}, {Name: "BotRad", Doc: "radius of the bottom -- set to 0 to omit bottom cap"}, {Name: "RadialSegs", Doc: "number of radial segments (32 is a reasonable default for full circle)"}, {Name: "HeightSegs", Doc: "number of height segments"}, {Name: "CapSegs", Doc: "number of segments in the hemisphere cap ends (16 is a reasonable default)"}, {Name: "AngStart", Doc: "starting angle in degrees, relative to -1,0,0 left side starting point"}, {Name: "AngLen", Doc: "total angle to generate in degrees (max 360)"}}})

// SetHeight sets the [Capsule.Height]:
// height of the cylinder portion
func (t *Capsule) SetHeight(v float32) *Capsule { t.Height = v; return t }

// SetTopRad sets the [Capsule.TopRad]:
// radius of the top -- set to 0 to omit top cap
func (t *Capsule) SetTopRad(v float32) *Capsule { t.TopRad = v; return t }

// SetBotRad sets the [Capsule.BotRad]:
// radius of the bottom -- set to 0 to omit bottom cap
func (t *Capsule) SetBotRad(v float32) *Capsule { t.BotRad = v; return t }

// SetRadialSegs sets the [Capsule.RadialSegs]:
// number of radial segments (32 is a reasonable default for full circle)
func (t *Capsule) SetRadialSegs(v int) *Capsule { t.RadialSegs = v; return t }

// SetHeightSegs sets the [Capsule.HeightSegs]:
// number of height segments
func (t *Capsule) SetHeightSegs(v int) *Capsule { t.HeightSegs = v; return t }

// SetCapSegs sets the [Capsule.CapSegs]:
// number of segments in the hemisphere cap ends (16 is a reasonable default)
func (t *Capsule) SetCapSegs(v int) *Capsule { t.CapSegs = v; return t }

// SetAngStart sets the [Capsule.AngStart]:
// starting angle in degrees, relative to -1,0,0 left side starting point
func (t *Capsule) SetAngStart(v float32) *Capsule { t.AngStart = v; return t }

// SetAngLen sets the [Capsule.AngLen]:
// total angle to generate in degrees (max 360)
func (t *Capsule) SetAngLen(v float32) *Capsule { t.AngLen = v; return t }

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/gpu/shape.Cylinder", IDName: "cylinder", Doc: "Cylinder is a generalized cylinder shape, including a cone\nor truncated cone by having different size circles at either end.\nHeight is up along the Y axis.", Directives: []types.Directive{{Tool: "types", Directive: "add", Args: []string{"-setters"}}}, Embeds: []types.Field{{Name: "ShapeBase"}}, Fields: []types.Field{{Name: "Height", Doc: "height of the cylinder"}, {Name: "TopRad", Doc: "radius of the top -- set to 0 for a cone"}, {Name: "BotRad", Doc: "radius of the bottom"}, {Name: "RadialSegs", Doc: "number of radial segments (32 is a reasonable default for full circle)"}, {Name: "HeightSegs", Doc: "number of height segments"}, {Name: "Top", Doc: "render the top disc"}, {Name: "Bottom", Doc: "render the bottom disc"}, {Name: "AngStart", Doc: "starting angle in degrees, relative to -1,0,0 left side starting point"}, {Name: "AngLen", Doc: "total angle to generate in degrees (max 360)"}}})

// SetHeight sets the [Cylinder.Height]:
// height of the cylinder
func (t *Cylinder) SetHeight(v float32) *Cylinder { t.Height = v; return t }

// SetTopRad sets the [Cylinder.TopRad]:
// radius of the top -- set to 0 for a cone
func (t *Cylinder) SetTopRad(v float32) *Cylinder { t.TopRad = v; return t }

// SetBotRad sets the [Cylinder.BotRad]:
// radius of the bottom
func (t *Cylinder) SetBotRad(v float32) *Cylinder { t.BotRad = v; return t }

// SetRadialSegs sets the [Cylinder.RadialSegs]:
// number of radial segments (32 is a reasonable default for full circle)
func (t *Cylinder) SetRadialSegs(v int) *Cylinder { t.RadialSegs = v; return t }

// SetHeightSegs sets the [Cylinder.HeightSegs]:
// number of height segments
func (t *Cylinder) SetHeightSegs(v int) *Cylinder { t.HeightSegs = v; return t }

// SetTop sets the [Cylinder.Top]:
// render the top disc
func (t *Cylinder) SetTop(v bool) *Cylinder { t.Top = v; return t }

// SetBottom sets the [Cylinder.Bottom]:
// render the bottom disc
func (t *Cylinder) SetBottom(v bool) *Cylinder { t.Bottom = v; return t }

// SetAngStart sets the [Cylinder.AngStart]:
// starting angle in degrees, relative to -1,0,0 left side starting point
func (t *Cylinder) SetAngStart(v float32) *Cylinder { t.AngStart = v; return t }

// SetAngLen sets the [Cylinder.AngLen]:
// total angle to generate in degrees (max 360)
func (t *Cylinder) SetAngLen(v float32) *Cylinder { t.AngLen = v; return t }

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/gpu/shape.ShapeGroup", IDName: "shape-group", Doc: "ShapeGroup is a group of shapes -- returns summary data for shape elements", Directives: []types.Directive{{Tool: "types", Directive: "add", Args: []string{"-setters"}}}, Embeds: []types.Field{{Name: "ShapeBase"}}, Fields: []types.Field{{Name: "Shapes", Doc: "list of shapes in group"}}})

// SetShapes sets the [ShapeGroup.Shapes]:
// list of shapes in group
func (t *ShapeGroup) SetShapes(v ...Shape) *ShapeGroup { t.Shapes = v; return t }

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/gpu/shape.Lines", IDName: "lines", Doc: "Lines are lines rendered as long thin boxes defined by points\nand width parameters.  The Mesh must be drawn in the XY plane (i.e., use Z = 0\nor a constant unless specifically relevant to have full 3D variation).\nRotate the solid to put into other planes.", Directives: []types.Directive{{Tool: "types", Directive: "add", Args: []string{"-setters"}}}, Embeds: []types.Field{{Name: "ShapeBase"}}, Fields: []types.Field{{Name: "Points", Doc: "line points (must be 2 or more)"}, {Name: "Width", Doc: "line width, Y = height perpendicular to line direction, and X = depth"}, {Name: "Colors", Doc: "optional colors for each point -- actual color interpolates between"}, {Name: "Closed", Doc: "if true, connect the first and last points to form a closed shape"}}})

// SetPoints sets the [Lines.Points]:
// line points (must be 2 or more)
func (t *Lines) SetPoints(v ...math32.Vector3) *Lines { t.Points = v; return t }

// SetWidth sets the [Lines.Width]:
// line width, Y = height perpendicular to line direction, and X = depth
func (t *Lines) SetWidth(v math32.Vector2) *Lines { t.Width = v; return t }

// SetColors sets the [Lines.Colors]:
// optional colors for each point -- actual color interpolates between
func (t *Lines) SetColors(v ...color.Color) *Lines { t.Colors = v; return t }

// SetClosed sets the [Lines.Closed]:
// if true, connect the first and last points to form a closed shape
func (t *Lines) SetClosed(v bool) *Lines { t.Closed = v; return t }

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/gpu/shape.Plane", IDName: "plane", Doc: "Plane is a flat 2D plane, which can be oriented along any\naxis facing either positive or negative", Directives: []types.Directive{{Tool: "types", Directive: "add", Args: []string{"-setters"}}}, Embeds: []types.Field{{Name: "ShapeBase"}}, Fields: []types.Field{{Name: "NormAxis", Doc: "axis along which the normal perpendicular to the plane points.  E.g., if the Y axis is specified, then it is a standard X-Z ground plane -- see also NormNeg for whether it is facing in the positive or negative of the given axis."}, {Name: "NormNeg", Doc: "if false, the plane normal facing in the positive direction along specified NormAxis, otherwise it faces in the negative if true"}, {Name: "Size", Doc: "2D size of plane"}, {Name: "Segs", Doc: "number of segments to divide plane into (enforced to be at least 1) -- may potentially increase rendering quality to have > 1"}, {Name: "Offset", Doc: "offset from origin along direction of normal to the plane"}}})

// SetNormAxis sets the [Plane.NormAxis]:
// axis along which the normal perpendicular to the plane points.  E.g., if the Y axis is specified, then it is a standard X-Z ground plane -- see also NormNeg for whether it is facing in the positive or negative of the given axis.
func (t *Plane) SetNormAxis(v math32.Dims) *Plane { t.NormAxis = v; return t }

// SetNormNeg sets the [Plane.NormNeg]:
// if false, the plane normal facing in the positive direction along specified NormAxis, otherwise it faces in the negative if true
func (t *Plane) SetNormNeg(v bool) *Plane { t.NormNeg = v; return t }

// SetSize sets the [Plane.Size]:
// 2D size of plane
func (t *Plane) SetSize(v math32.Vector2) *Plane { t.Size = v; return t }

// SetSegs sets the [Plane.Segs]:
// number of segments to divide plane into (enforced to be at least 1) -- may potentially increase rendering quality to have > 1
func (t *Plane) SetSegs(v math32.Vector2i) *Plane { t.Segs = v; return t }

// SetOffset sets the [Plane.Offset]:
// offset from origin along direction of normal to the plane
func (t *Plane) SetOffset(v float32) *Plane { t.Offset = v; return t }

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/gpu/shape.ShapeBase", IDName: "shape-base", Doc: "ShapeBase is the base shape element", Directives: []types.Directive{{Tool: "types", Directive: "add", Args: []string{"-setters"}}}, Fields: []types.Field{{Name: "VertexOff", Doc: "vertex offset, in points"}, {Name: "IndexOff", Doc: "index offset, in points"}, {Name: "CBBox", Doc: "cubic bounding box in local coords"}, {Name: "Pos", Doc: "all shapes take a 3D position offset to enable composition"}}})

// SetVertexOff sets the [ShapeBase.VertexOff]:
// vertex offset, in points
func (t *ShapeBase) SetVertexOff(v int) *ShapeBase { t.VertexOff = v; return t }

// SetIndexOff sets the [ShapeBase.IndexOff]:
// index offset, in points
func (t *ShapeBase) SetIndexOff(v int) *ShapeBase { t.IndexOff = v; return t }

// SetCBBox sets the [ShapeBase.CBBox]:
// cubic bounding box in local coords
func (t *ShapeBase) SetCBBox(v math32.Box3) *ShapeBase { t.CBBox = v; return t }

// SetPos sets the [ShapeBase.Pos]:
// all shapes take a 3D position offset to enable composition
func (t *ShapeBase) SetPos(v math32.Vector3) *ShapeBase { t.Pos = v; return t }

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/gpu/shape.Sphere", IDName: "sphere", Doc: "Sphere is a sphere shape (can be a partial sphere too)", Directives: []types.Directive{{Tool: "types", Directive: "add", Args: []string{"-setters"}}}, Embeds: []types.Field{{Name: "ShapeBase"}}, Fields: []types.Field{{Name: "Radius", Doc: "radius of the sphere"}, {Name: "WidthSegs", Doc: "number of segments around the width of the sphere (32 is reasonable default for full circle)"}, {Name: "HeightSegs", Doc: "number of height segments (32 is reasonable default for full height)"}, {Name: "AngStart", Doc: "starting radial angle in degrees, relative to -1,0,0 left side starting point"}, {Name: "AngLen", Doc: "total radial angle to generate in degrees (max = 360)"}, {Name: "ElevStart", Doc: "starting elevation (height) angle in degrees - 0 = top of sphere, and Pi is bottom"}, {Name: "ElevLen", Doc: "total angle to generate in degrees (max = 180)"}}})

// SetRadius sets the [Sphere.Radius]:
// radius of the sphere
func (t *Sphere) SetRadius(v float32) *Sphere { t.Radius = v; return t }

// SetWidthSegs sets the [Sphere.WidthSegs]:
// number of segments around the width of the sphere (32 is reasonable default for full circle)
func (t *Sphere) SetWidthSegs(v int) *Sphere { t.WidthSegs = v; return t }

// SetHeightSegs sets the [Sphere.HeightSegs]:
// number of height segments (32 is reasonable default for full height)
func (t *Sphere) SetHeightSegs(v int) *Sphere { t.HeightSegs = v; return t }

// SetAngStart sets the [Sphere.AngStart]:
// starting radial angle in degrees, relative to -1,0,0 left side starting point
func (t *Sphere) SetAngStart(v float32) *Sphere { t.AngStart = v; return t }

// SetAngLen sets the [Sphere.AngLen]:
// total radial angle to generate in degrees (max = 360)
func (t *Sphere) SetAngLen(v float32) *Sphere { t.AngLen = v; return t }

// SetElevStart sets the [Sphere.ElevStart]:
// starting elevation (height) angle in degrees - 0 = top of sphere, and Pi is bottom
func (t *Sphere) SetElevStart(v float32) *Sphere { t.ElevStart = v; return t }

// SetElevLen sets the [Sphere.ElevLen]:
// total angle to generate in degrees (max = 180)
func (t *Sphere) SetElevLen(v float32) *Sphere { t.ElevLen = v; return t }

var _ = types.AddType(&types.Type{Name: "cogentcore.org/core/gpu/shape.Torus", IDName: "torus", Doc: "Torus is a torus mesh, defined by the radius of the solid tube and the\nlarger radius of the ring.", Directives: []types.Directive{{Tool: "types", Directive: "add", Args: []string{"-setters"}}}, Embeds: []types.Field{{Name: "ShapeBase"}}, Fields: []types.Field{{Name: "Radius", Doc: "larger radius of the torus ring"}, {Name: "TubeRadius", Doc: "radius of the solid tube"}, {Name: "RadialSegs", Doc: "number of segments around the radius of the torus (32 is reasonable default for full circle)"}, {Name: "TubeSegs", Doc: "number of segments for the tube itself (32 is reasonable default for full height)"}, {Name: "AngStart", Doc: "starting radial angle in degrees relative to 1,0,0 starting point"}, {Name: "AngLen", Doc: "total radial angle to generate in degrees (max = 360)"}}})

// SetRadius sets the [Torus.Radius]:
// larger radius of the torus ring
func (t *Torus) SetRadius(v float32) *Torus { t.Radius = v; return t }

// SetTubeRadius sets the [Torus.TubeRadius]:
// radius of the solid tube
func (t *Torus) SetTubeRadius(v float32) *Torus { t.TubeRadius = v; return t }

// SetRadialSegs sets the [Torus.RadialSegs]:
// number of segments around the radius of the torus (32 is reasonable default for full circle)
func (t *Torus) SetRadialSegs(v int) *Torus { t.RadialSegs = v; return t }

// SetTubeSegs sets the [Torus.TubeSegs]:
// number of segments for the tube itself (32 is reasonable default for full height)
func (t *Torus) SetTubeSegs(v int) *Torus { t.TubeSegs = v; return t }

// SetAngStart sets the [Torus.AngStart]:
// starting radial angle in degrees relative to 1,0,0 starting point
func (t *Torus) SetAngStart(v float32) *Torus { t.AngStart = v; return t }

// SetAngLen sets the [Torus.AngLen]:
// total radial angle to generate in degrees (max = 360)
func (t *Torus) SetAngLen(v float32) *Torus { t.AngLen = v; return t }
