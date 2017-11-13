#version 330
uniform mat4 projection;
uniform mat4 camera;
uniform mat4 model;
in vec3 vert;
in vec2 vertTexCoord;
out vec2 fragTexCoord;
void main() {
    fragTexCoord = vertTexCoord;
    gl_Position = projection * camera  * vec4(vert, 1);
}
/*
attribute vec4 a_position;
attribute vec4  a_color;
attribute vec2  a_texCoord0;
uniform mat4 u_projTrans;
varying vec4 v_color;
varying vec2 v_texCoords;

void main(){
    v_color =  a_color;
    v_color.a = v_color.a * (255.0/254.0);
    v_texCoords =  a_texCoord0;
    gl_Position =  u_projTrans *  a_position;
}*/