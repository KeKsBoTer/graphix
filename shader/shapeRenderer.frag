#version 330
#ifdef GL_ES
    #define LOWP lowp
    precision mediump float;
#else
    #define LOWP
#endif
uniform sampler2D tex;
in vec4 LOWP v_color;
out vec4 outputColor;
void main() {
    outputColor = v_color;
}