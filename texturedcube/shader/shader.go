package shader

import (
	"github.com/goxjs/gl"
)

// Online live shader editor: http://shdr.bkcore.com/
// gman's explanation is great: http://stackoverflow.com/questions/30364213/shaders-in-webgl-vs-opengl
// GLSL (GL Shading Language) Reference: http://www.shaderific.com/glsl/   Particularly the qualifiers section.

var (
	Texture *texture
)

func Initialize() error {
	errs := make(chan error, 10)
	errs <- setupTextureShader()
	close(errs)
	for err := range errs {
		if err != nil {
			return err
		}
	}
	Texture.SetDefaults()
	gl.UseProgram(Texture.Program)
	return nil
}
