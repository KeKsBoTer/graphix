#version 330
uniform mat4 projection;
uniform mat4 camera;
uniform mat4 model;
in vec2 vert;
in vec2 vertTexCoord;
in vec4 a_color;
out vec2 fragTexCoord;
out vec4 v_color;

void main() {
    v_color = a_color;
    fragTexCoord = vertTexCoord;
    gl_Position = projection * camera  * vec4(vert,0, 1);
}