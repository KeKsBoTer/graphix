#version 330
uniform mat4 projection;
uniform mat4 camera;
uniform mat4 model;
in vec2 vert;
in vec4 a_color;
out vec4 v_color;

void main() {
    v_color = a_color;
    v_color.a = v_color.a * (255.0/254.0);
    gl_Position = projection * camera  * vec4(vert,0, 1);
}