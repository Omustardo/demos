Implementation of tutorials at https://open.gl/introduction

Note that I wrote them using the goxjs/gl and goxjs/glfw libraries rather than the standard go-gl/gl and go-gl/glfw.

goxjs allows for context sensitive compilation, so you can compile Go code to Javascript and have it run in a browser 
using WebGL. Unfortunately WebGL doesn't support the modern GLSL syntax used in these tutorials, so you can't actually 
take advantage of that in this case. It still runs on desktop though.
 
Since my other projects are using goxjs and I mostly worked through these for practice, I decided not to switch to the
standard gl and glfw libraries.
