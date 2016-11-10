package parallax
import "github.com/go-gl/mathgl/mgl32"
// Parallax interface allows for easy implementation of a Parallax effect.
// Essentially, objects far away appear to move more slowly relative to the camera, so
// to simulate this, an object should have a movement ratio, such as 0.5 - in which case you simply take
// the camera's X,Y coordinate and move the parallax background object to 50% of that. In most cases this will
// make the object not appear on screen, so if you have an object that can be tiled, you should mod the camera's
// position by the object's width or height (for x,y coords respectively).
type Parallax interface {
	GetParallaxPosition() mgl32.Vec2
}
