package main

import (
	"fmt"
	"image/color"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Rectangle struct {
	Points [4]Point
}

type spiner struct {
	angle float32
}

type player struct {
	opponents        []Point
	x, y             int32
	targetX, targetY int32
	targetLocked     bool
}

func (p *player) Update() {
	if rl.IsKeyDown(rl.KeyW) {
		p.y -= 2
	}
	if rl.IsKeyDown(rl.KeyA) {
		p.x -= 2
	}
	if rl.IsKeyDown(rl.KeyD) {
		p.x += 2
	}
	if rl.IsKeyDown(rl.KeyS) {
		p.y += 2
	}

	lockedFound := false
	for i := range p.opponents {
		if distance(Point{float64(p.x), float64(p.y)}, Point{p.opponents[i].X, p.opponents[i].Y}) < 40 {
			p.targetLocked = true
			lockedFound = true
			p.targetX = int32(p.opponents[i].X)
			p.targetY = int32(p.opponents[i].Y)
			break
		}
	}

	if !lockedFound {
		p.targetLocked = false
	}
}

func (p *player) Draw() {

	if p.targetLocked {
		angle := calculateAngle(Point{float64(p.x), float64(p.y)}, Point{float64(p.targetX), float64(p.targetY)})

		rotation := calculateRotation(angle)

		if p.x >= p.targetX {
			rotation += 180
		}

		if p.x == p.targetX && p.y > p.targetY {

			rotation += 180
		}

		rl.DrawRectanglePro(rl.NewRectangle(float32(p.x), float32(p.y), float32(60), 80), rl.Vector2{
			X: 0,
			Y: 40,
		}, -float32(rotation), rl.Red)

		DrawPoints(getRectangleCorners(Point{float64(p.x), float64(p.y)}, Point{float64(p.targetX), float64(p.targetY)}, 80, 60))
		pnts := CreateRectangle(getRectangleCorners(Point{float64(p.x), float64(p.y)}, Point{float64(p.targetX), float64(p.targetY)}, 80, 60))
		tRectangle := rl.Rectangle{330, 340, 30, 30}
		other := CreateRectangle(Point{330, 340}, Point{360, 340}, Point{330, 370}, Point{360, 370})

		collisionColor := rl.Gold
		if rectCollision(pnts, other) {
			collisionColor = rl.Red
		}

		rl.DrawRectanglePro(tRectangle, rl.Vector2{}, 0, collisionColor)

	}

	newp := Point{200, 20}
	rotation := Point{200, 50}
	DrawPoint(newp, rl.Pink)
	DrawPoint(rotation, rl.Green)
	rotated := ownRotate(newp, rotation, 90)
	DrawPoint(rotated, rl.Yellow)

	// new stuff
	newA := Point{400, 400}
	newB := Point{432, 472}
	width := 90
	length := 220
	iThinkTopRight := Point{newA.X + float64(width/2), newA.Y}
	iThinkBotRight := Point{newA.X + float64(width/2), newA.Y + float64(length)}
	newAngle := calculateAngle(newA, newB)
	newRotation := calculateRotation(newAngle)

	plusAngle := math.Tan((float64(width/2) / float64(length))) * (180 / math.Pi)

	rpa := ownRotate(iThinkTopRight, newA, -float64(newRotation-270))
	otherRpa := ownRotate(iThinkTopRight, newA, -float64(newRotation-270+180))
	bottomRight := ownRotate(iThinkBotRight, newA, -float64(newRotation-270))
	bottomLeft := ownRotate(iThinkBotRight, newA, -float64(newRotation-270-plusAngle-plusAngle))

	rl.DrawRectanglePro(rl.NewRectangle(float32(newA.X), float32(newA.Y), float32(width/2), float32(length)), rl.Vector2{}, -float32(newRotation-270), rl.Pink)
	DrawPoint(newA, rl.Green)
	DrawPoint(newB, rl.Blue)
	DrawPoint(iThinkTopRight, rl.DarkGreen)
	DrawPoint(rpa, rl.Yellow)
	DrawPoint(otherRpa, rl.Green)
	DrawPoint(bottomRight, rl.Gold)
	DrawPoint(bottomLeft, rl.Red)

	DrawPoint(Point{float64(p.x), float64(p.y)}, rl.Green)
	for i := range p.opponents {
		DrawPoint(p.opponents[i], rl.Yellow)
	}
}

func (s *spiner) Draw() {
	s.angle = 110
	rl.DrawRectanglePro(rl.NewRectangle(40, 40, 80, 80), rl.Vector2{10, 10}, s.angle, rl.Lime)
}

func main() {

	rl.InitWindow(1400, 1400, "rectangle test")

	defer rl.CloseWindow()

	rl.SetTargetFPS(60)
	rl.SetExitKey(13123213)

	s := spiner{}

	p := player{}
	p.opponents = append(p.opponents, Point{300, 300})

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		a := Point{200, 300}
		b := Point{90, 40}
		c := Point{
			X: 90 - 36,
			Y: 40 + 15,
		}

		rl.DrawRectanglePro(rl.NewRectangle(float32(b.X), float32(b.Y), float32(distance(a, b)), 80), rl.Vector2{
			X: 0,
			Y: 40,
		}, float32(calculateAngle(b, a)), rl.Yellow)

		DrawPoint(a, rl.Green)
		DrawPoint(b, rl.Blue)
		DrawPoint(c, rl.Pink)

		// hypotenuse := 40.0
		// angleAlpha := 180.0 - 67.0 // in degrees
		//
		// // Convert the angle to radians as the math package functions use radians
		// angleAlphaRad := angleAlpha * (math.Pi / 180.0)
		//
		// // Calculate the adjacent side (b) using cosine
		// adjacent := hypotenuse * math.Cos(angleAlphaRad)
		//
		// // Calculate the opposite side (a) using sine
		// opposite := hypotenuse * math.Sin(angleAlphaRad)
		//
		// // Calculate the angles back from the sides
		// angleBetaRad := math.Atan(opposite / adjacent)
		// angleBeta := angleBetaRad * (180.0 / math.Pi) // convert back to degrees

		// Angle gamma is 90 degrees in a right-angled triangle
		// angleGamma := 90.0

		// Angle alpha can also be confirmed by the sides using atan (should be same as input)
		// angleAlphaConfirmedRad := math.Atan(adjacent / opposite)
		// ang eAlphaConfirmed := angleAlphaConfirmedRad * (180.0 / math.Pi) // convert back to degrees

		// Print the results
		// fmt.Printf("Hypotenuse (C): %.2f\n", hypotenuse)
		// fmt.Printf("Given Angle Alpha: %.2f degrees\n", angleAlpha)
		// fmt.Printf("Calculated Adjacent side (b): %.2f\n", adjacent)
		// fmt.Printf("Calculated Opposite side (a): %.2f\n", opposite)
		//fmt.Printf("Calculated Angle Beta: %.2f degrees\n", angleBeta)
		// fmt.Printf("Calculated Angle Gamma: %.2f degrees\n", angleGamma)
		// fmt.Printf("Confirmed Angle Alpha: %.2f degrees\n", angleAlphaConfirmed)

		s.Draw()
		p.Update()
		p.Draw()
		rl.ClearBackground(rl.Black)

		rl.EndDrawing()
	}
}

type Point struct {
	X, Y float64
}

func calculateAngle(a, b Point) float64 {
	// Calculate the differences in x and y coordinates
	dx := b.X - a.X
	dy := b.Y - a.Y

	// Calculate the angle using atan2
	angle := math.Atan2(dy, dx)

	// Convert the angle from radians to degrees
	angleDegrees := angle * (180 / math.Pi)

	return angleDegrees
}

func DrawPoints(a, b, c, d Point) {
	rl.DrawCircle(int32(a.X), int32(a.Y), 5, rl.Yellow)
	rl.DrawCircle(int32(b.X), int32(b.Y), 5, rl.Orange)
	rl.DrawCircle(int32(c.X), int32(c.Y), 5, rl.Pink)
	rl.DrawCircle(int32(d.X), int32(d.Y), 5, rl.Pink)
}

func DrawPoint(a Point, color color.RGBA) {
	rl.DrawCircle(int32(a.X), int32(a.Y), 4, color)
}

// perpendicular calculates a perpendicular vector.
func perpendicular(a Point) Point {
	return Point{-a.Y, a.X}
}

// calculateRectangle calculates the coordinates of the rectangle.
func calculateRectangle(a, b Point, w, l float64) (Point, Point, Point, Point) {
	// Calculate the direction vector from A to B
	dir := Point{b.X - a.X, b.Y - a.Y}

	// Calculate the length of the direction vector
	dirLength := math.Sqrt(dir.X*dir.X + dir.Y*dir.Y)

	// Calculate the unit vector in the direction of A to B
	unitDir := Point{dir.X / dirLength, dir.Y / dirLength}

	// Calculate the perpendicular vector to the direction vector
	perpDir := perpendicular(unitDir)

	// Scale the perpendicular vector by half the length of the short side
	halfW := w / 2
	perpDir.X *= halfW
	perpDir.Y *= halfW

	// Scale the direction vector by half the length of the long side
	halfL := l / 2
	unitDir.X *= halfL
	unitDir.Y *= halfL

	// Calculate the four corners of the rectangle
	p1 := Point{a.X + perpDir.X, a.Y + perpDir.Y}
	p2 := Point{a.X - perpDir.X, a.Y - perpDir.Y}
	p3 := Point{b.X + perpDir.X, b.Y + perpDir.Y}
	p4 := Point{b.X - perpDir.X, b.Y - perpDir.Y}

	return p1, p2, p3, p4
}

func distance(a, b Point) float64 {
	return math.Sqrt((b.X-a.X)*(b.X-a.X) + (b.Y-a.Y)*(b.Y-a.Y))
}

// normalize normalizes a vector.
func normalize(a Point) Point {
	mag := math.Sqrt(a.X*a.X + a.Y*a.Y)
	return Point{a.X / mag, a.Y / mag}
}

// calculateC calculates the coordinates of point C.
func calculateC(a, b Point, dist float64) Point {
	// Calculate the direction vector from A to B
	dir := Point{b.X - a.X, b.Y - a.Y}

	// Normalize the direction vector
	normDir := normalize(dir)

	// Scale the normalized direction vector by the distance X
	scaledDir := Point{normDir.X * dist, normDir.Y * dist}

	// Calculate point C
	c := Point{a.X + scaledDir.X, a.Y + scaledDir.Y}
	return c
}

func calculateRotation(angle float64) float64 {
	//half width of short side
	hypotenuse := 40.0

	angleAlpha := 180.0 - angle // in degrees

	// Convert the angle to radians as the math package functions use radians
	angleAlphaRad := angleAlpha * (math.Pi / 180.0)

	// Calculate the adjacent side (b) using cosine
	adjacent := hypotenuse * math.Cos(angleAlphaRad)

	// Calculate the opposite side (a) using sine
	opposite := hypotenuse * math.Sin(angleAlphaRad)

	// Calculate the angles back from the sides
	angleBetaRad := math.Atan(opposite / adjacent)
	angleBeta := angleBetaRad * (180.0 / math.Pi) // convert back to degrees

	if angleBeta < 0 && angleBeta < -90 {
		angleBeta += 180
	}

	return angleBeta

}

func getRectangleCorners(a, b Point, width, height float64) (Point, Point, Point, Point) {
	iThinkTopRight := Point{a.X + float64(width/2), a.Y}
	iThinkBotRight := Point{a.X + float64(width/2), a.Y + height}
	newAngle := calculateAngle(a, b)
	newRotation := calculateRotation(newAngle)

	plusAngle := math.Tan((float64(width/2) / height)) * (180 / math.Pi)

	rpa := ownRotate(iThinkTopRight, a, -float64(newRotation-270))
	otherRpa := ownRotate(iThinkTopRight, a, -float64(newRotation-270+180))

	if a.X > b.X || (a.X == b.X && a.Y < b.Y) {
		newRotation += 180
	}
	bottomRight := ownRotate(iThinkBotRight, a, -float64(newRotation-270))
	bottomLeft := ownRotate(iThinkBotRight, a, -float64(newRotation-270-plusAngle-plusAngle/2))

	fmt.Println(newRotation)

	fmt.Println(bottomLeft, bottomRight)

	return rpa, otherRpa, bottomRight, bottomLeft
}

func rotatePoint(p, center Point, angle float64) Point {
	s := math.Sin(angle)
	c := math.Cos(angle)

	// Translate point to origin
	p.X -= center.X
	p.Y -= center.Y

	// Rotate point
	xnew := p.X*c - p.Y*s
	ynew := p.X*s + p.Y*c

	// Translate point back
	p.X = xnew + center.X
	p.Y = ynew + center.Y
	return p
}

func ownRotate(p, center Point, angle float64) Point {
	// where c, s are the cosine and sine of the angle.
	//
	// A rotation around an arbitrary point (u, v) is
	//
	// X = c (x - u) - s (y - v) + u
	// Y = s (x - u) + c (y - v) + v

	c := math.Cos(angle * math.Pi / 180)
	s := math.Sin(angle * math.Pi / 180)

	x := c*(p.X-center.X) - s*(p.Y-center.Y) + center.X
	y := s*(p.X-center.X) + c*(p.Y-center.Y) + center.Y

	return Point{x, y}

}

func dotProduct(v1, v2 Point) float64 {
	return v1.X*v2.X + v1.Y*v2.Y
}

// subtractPoints subtracts one point from another to get a vector
func subtractPoints(p1, p2 Point) Point {
	return Point{p1.X - p2.X, p1.Y - p2.Y}
}

// perpendicularVector returns a vector that is perpendicular to the given vector
func perpendicularVector(v Point) Point {
	return Point{-v.Y, v.X}
}

// projectPoints projects the points of a rectangle onto an axis and returns the minimum and maximum values
func projectPoints(rect Rectangle, axis Point) (float64, float64) {
	min := dotProduct(subtractPoints(rect.Points[0], Point{0, 0}), axis)
	max := min
	for _, point := range rect.Points[1:] {
		projection := dotProduct(subtractPoints(point, Point{0, 0}), axis)
		if projection < min {
			min = projection
		}
		if projection > max {
			max = projection
		}
	}
	return min, max
}

// overlap checks if two projection ranges overlap
func overlap(minA, maxA, minB, maxB float64) bool {
	return !(minA > maxB || minB > maxA)
}

// rectCollision checks if two rectangles collide using the Separating Axis Theorem
func rectCollision(rect1, rect2 Rectangle) bool {
	axes := []Point{
		subtractPoints(rect1.Points[1], rect1.Points[0]),
		subtractPoints(rect1.Points[3], rect1.Points[0]),
		subtractPoints(rect2.Points[1], rect2.Points[0]),
		subtractPoints(rect2.Points[3], rect2.Points[0]),
	}

	for _, axis := range axes {
		perpAxis := perpendicularVector(axis)
		minA, maxA := projectPoints(rect1, perpAxis)
		minB, maxB := projectPoints(rect2, perpAxis)
		if !overlap(minA, maxA, minB, maxB) {
			return false
		}
	}
	return true
}

func CreateRectangle(a, b, c, d Point) Rectangle {
	return Rectangle{Points: [4]Point{a, b, c, d}}
}
