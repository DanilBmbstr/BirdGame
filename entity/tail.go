package entity

/*
⢠⣶⠶⢶⡆⠀⠀⠐⢶⣆⢠⣶⠆⢠⣶⠿⢷⣦⡀⢰⡆⠀⢰⣶⠀⢶⣆⢀⣶⠆
⠘⢿⣦⣼⡇⠀⠀⠀⠈⣿⣿⡃⠀⣿⡇⠀⠀⣿⡇⢸⣷⣤⣼⣿⠀⠈⣿⣾⡟⠀
⢠⣿⠃⢸⡇⠀⠀⢀⣼⡟⠹⣷⡀⠻⣷⣤⣴⡿⠃⠀⠉⠉⢹⣿⠀⣤⣼⡟⠀⠀
⢀⣴⡶⠶⢶⠀⢰⡶⠶⣶⣄⠀⠀⢠⣶⣶⡀⠀⠰⠶⢶⣶⠶⠦⢰⣶⠀⠀⠀⠀
⣾⡏⠀⠀⠀⠀⢸⡇⠀⣸⣿⠀⠀⣾⡇⢻⣧⠀⠀⠀⢸⣿⠀⠀⢸⣿⣤⣤⣄⡀
⢿⣇⠀⠀⠀⠀⢸⡟⠛⠛⠁⠀⢰⣿⣤⣼⣿⡆⠀⠀⢸⣿⠀⠀⢸⣿⠁⠈⢹⣷
⠈⠻⠷⠶⠾⠀⠸⠇⠀⠀⠀⠠⠿⠃⠀⠀⠸⠿⠀⠀⠸⠿⠀⠀⠸⠿⠶⠶⠿
*/
import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/quasilyte/gmath"
)

type tailSegment struct {
	start        gmath.Vec
	direction    gmath.Vec
	segmentWidht float32
}

type Tail struct {
	segments []tailSegment

	tailWidth     float32
	widthNarrow   float32
	SegmentLenght float64
	RotationSpeed float64
	MovingSpeed   float64
}

func (ts *tailSegment) getSegmentFlanks() (gmath.Vec, gmath.Vec) {

	left := ts.start.Add(ts.direction.Rotated(gmath.DegToRad(-90)).Mulf(float64(ts.segmentWidht)))
	right := ts.start.Add(ts.direction.Rotated(gmath.DegToRad(90)).Mulf(float64(ts.segmentWidht)))

	return left, right
}

func (t *Tail) getSegmentEnd(segmentStart gmath.Vec) gmath.Vec {

	return segmentStart.Add(gmath.Vec{X: -t.SegmentLenght, Y: 0})
}

func (ts *tailSegment) getSegmentEnd(segmentLenght float64) gmath.Vec {

	return ts.start.Add(gmath.Vec{X: -segmentLenght, Y: 0}.Mul(ts.direction))
}
func createSegment(start gmath.Vec, width float32) tailSegment {
	return tailSegment{start, gmath.Vec{X: -1, Y: 0}, width}
}

func (t *Tail) CreateTail(point gmath.Vec, max int) {

	t.segments = append(t.segments, createSegment(point, t.tailWidth))
	t.tailWidth *= t.widthNarrow
	if max > 0 {
		t.CreateTail(t.getSegmentEnd(point), max-1)

	}
}

func (t *Tail) DrawTail(screen *ebiten.Image, xOffset, yOffset int) {
	path := &vector.Path{}
	for i := 0; i < len(t.segments)-1; i++ {
		left, right := t.segments[i].getSegmentFlanks()
		left2, right2 := t.segments[i+1].getSegmentFlanks()

		path.MoveTo(left.AsVec32().X-float32(xOffset), left.AsVec32().Y-float32(yOffset))

		path.LineTo(right.AsVec32().X-float32(xOffset), right.AsVec32().Y-float32(yOffset))
		path.LineTo(right2.AsVec32().X-float32(xOffset), right2.AsVec32().Y-float32(yOffset))
		path.LineTo(left2.AsVec32().X-float32(xOffset), left2.AsVec32().Y-float32(yOffset))
		path.Close()

	}

	cs := ebiten.ColorScale{}
	cs.Scale(0.5, 0.5, 0.5, 0.5)
	vector.FillPath(screen, path, &vector.FillOptions{FillRule: vector.FillRuleEvenOdd}, &vector.DrawPathOptions{AntiAlias: true,
		ColorScale: cs})
}

func (ts *tailSegment) moveSegment(speed float64) {
	ts.start = ts.start.Add(ts.direction.Normalized().Mulf(speed))
}

func (ts *tailSegment) rotateSegment(point gmath.Vec, speed float64) {
	ts.direction = ts.direction.LinearInterpolate(ts.start.DirectionTo(point), speed)

}

func (t *Tail) GetTailDistanceTo(to gmath.Vec) float64 {
	return t.segments[0].start.DistanceTo(to)
}

func (t *Tail) MoveTail(point gmath.Vec) {

	t.segments[0].rotateSegment(t.getSegmentEnd(point), t.RotationSpeed)
	t.segments[0].moveSegment(t.MovingSpeed)
	for i := 1; i < len(t.segments); i++ {
		t.segments[i].rotateSegment(t.segments[i-1].getSegmentEnd(t.SegmentLenght), t.RotationSpeed)
		t.segments[i].moveSegment(t.MovingSpeed)
	}
}
