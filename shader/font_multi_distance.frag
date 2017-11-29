#version 330
#ifdef GL_ES
    #define LOWP lowp
    precision mediump float;
#else
    #define LOWP
#endif
uniform sampler2D tex;
in vec2 fragTexCoord;
in vec4 LOWP v_color;
out vec4 outputColor;

float median(float r, float g, float b) {
    return max(min(r, g), min(max(r, g), b));
}

void main() {
    vec3 sampleColor = texture(tex, fragTexCoord).rgb;
    float sigDist = median(sampleColor.r, sampleColor.g, sampleColor.b) - 0.5;
    float opacity = clamp(sigDist/fwidth(sigDist) + 0.5, 0.0, 1.0);

    outputColor = vec4(v_color.rgb,(1- opacity)*v_color.a);
}