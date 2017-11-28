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

uniform vec4 outlineColor;  // The color of the outline
uniform float outlineWidth; // The width of the outline from the center
uniform float outlineEdge;  // The smoothness of the outline edge
uniform vec2 outlineOffset; // Translates the outline (can be used for shadows)
uniform float width = 0.5;  // The maximum distance from the center
uniform float edge = 0.01;  // The smoothnes of the edge

void main() {
    float distance = 1.0 - texture(tex, fragTexCoord).a;
    float alpha = 1.0 - smoothstep(width,width+ edge,distance);

    if(outlineColor.a > 0 ){ //Render outline
        float outlineDistance = 1.0 - texture(tex, fragTexCoord+outlineOffset).a;
        float outlineAlpha = 1.0 - smoothstep(outlineWidth, outlineWidth+outlineEdge, outlineDistance);
        float combined = alpha +(1-alpha)*outlineAlpha;
        vec4 finColor = mix(outlineColor,v_color,alpha/combined);
        outputColor = vec4(finColor.rgb,finColor.a*combined);
    }else{
        outputColor =  vec4(v_color.rgb,v_color*alpha);
    }
}