import { useEffect, useRef } from "react";

// Vertex shader: full‑screen triangle
const VS_SOURCE = `#version 300 es
precision highp float;
layout(location = 0) in vec2 aPos;
void main() {
  gl_Position = vec4(aPos, 0.0, 1.0);
}`;

/* ────────────────────────────────────────────────
 * Three subtle fragment shaders.
 * Each keeps luminance low so foreground text stays readable.
 * One is picked at component‑mount time.
 *
 * 0 – Horizontal flowing waves (was the original)
 * 1 – Radial pulse rings
 * 2 – Gentle diagonal drift
 * ──────────────────────────────────────────────── */
const FRAG_SHADERS = [
  // 0 ▸ horizontal sine bands
  `#version 300 es
precision highp float;
out vec4 fragColor;
uniform vec3 iResolution; uniform float iTime;
void mainImage(out vec4 c, in vec2 fc) {
  vec2 uv = fc / iResolution.xy;
  float wave = sin(uv.x * 10.0 - iTime * 0.8);
  float band = smoothstep(0.48, 0.5, abs(uv.y + wave * 0.02 - 0.5));
  vec3 base = vec3(0.07, 0.09, 0.11);
  vec3 color = mix(base, base + 0.05, band);
  c = vec4(color, 1.0);
}
void main() { mainImage(fragColor, gl_FragCoord.xy); }`,

  // 1 ▸ concentric dim rings that fade in/out
  `#version 300 es
precision highp float;
out vec4 fragColor;
uniform vec3 iResolution; uniform float iTime;
void main() {
  vec2 uv = (gl_FragCoord.xy - 0.5 * iResolution.xy) / iResolution.y;
  float d = length(uv);
  float rings = 0.5 + 0.5 * sin(10.0 * d - iTime * 0.6);
  vec3 col = vec3(0.06,0.08,0.1) + rings * 0.04;
  fragColor = vec4(col, 1.0);
}`,

  // 2 ▸ diagonal moving stripes
  `#version 300 es
precision highp float;
out vec4 fragColor;
uniform vec3 iResolution; uniform float iTime;
void main() {
  vec2 uv = gl_FragCoord.xy / iResolution.xy;
  float stripe = 0.5 + 0.5 * sin((uv.x + uv.y) * 12.0 - iTime * 0.7);
  vec3 col = vec3(0.08,0.1,0.12) + stripe * 0.03;
  fragColor = vec4(col, 1.0);
}`
];

function compileShader(gl: WebGL2RenderingContext, type: number, src: string) {
  const sh = gl.createShader(type);
  if (!sh) return null;
  gl.shaderSource(sh, src);
  gl.compileShader(sh);
  if (!gl.getShaderParameter(sh, gl.COMPILE_STATUS)) {
    console.warn(gl.getShaderInfoLog(sh));
    return null;
  }
  return sh;
}

function linkProgram(gl: WebGL2RenderingContext, vs: WebGLShader, fs: WebGLShader) {
  const prog = gl.createProgram();
  if (!prog) return null;
  gl.attachShader(prog, vs);
  gl.attachShader(prog, fs);
  gl.linkProgram(prog);
  if (!gl.getProgramParameter(prog, gl.LINK_STATUS)) {
    console.warn(gl.getProgramInfoLog(prog));
    return null;
  }
  return prog;
}

function initGeometry(gl: WebGL2RenderingContext) {
  const vao = gl.createVertexArray();
  gl.bindVertexArray(vao);
  const vbo = gl.createBuffer();
  gl.bindBuffer(gl.ARRAY_BUFFER, vbo);
  gl.bufferData(gl.ARRAY_BUFFER, new Float32Array([
    -1, -1, 1, -1, -1, 1,
    -1, 1, 1, -1, 1, 1,
  ]), gl.STATIC_DRAW);
  gl.enableVertexAttribArray(0);
  gl.vertexAttribPointer(0, 2, gl.FLOAT, false, 0, 0);
}

const BackgroundShader: React.FC = () => {
  const canvasRef = useRef<HTMLCanvasElement>(null);
  const glRef = useRef<WebGL2RenderingContext | null>(null);
  const progRef = useRef<WebGLProgram | null>(null);
  const start = useRef<number>(performance.now());

  // Pick a shader once per mount
  const chosenFS = useRef<string>(
    FRAG_SHADERS[Math.floor(Math.random() * FRAG_SHADERS.length)],
  );

  useEffect(() => {
    const canvas = canvasRef.current;
    if (!canvas) return;
    const gl = (glRef.current ||= canvas.getContext("webgl2")!);
    if (!gl) return;

    initGeometry(gl);
    const vs = compileShader(gl, gl.VERTEX_SHADER, VS_SOURCE);
    const fs = compileShader(gl, gl.FRAGMENT_SHADER, chosenFS.current);
    if (!vs || !fs) return;
    progRef.current = linkProgram(gl, vs, fs);
  }, []);

  useEffect(() => {
    const canvas = canvasRef.current;
    const gl = glRef.current;
    const prog = progRef.current;
    if (!canvas || !gl || !prog) return;

    const render = () => {
      const dpr = window.devicePixelRatio || 1;
      const w = canvas.clientWidth * dpr;
      const h = canvas.clientHeight * dpr;
      if (canvas.width !== w || canvas.height !== h) {
        canvas.width = w;
        canvas.height = h;
      }
      gl.viewport(0, 0, canvas.width, canvas.height);
      gl.useProgram(prog);
      gl.uniform3f(
        gl.getUniformLocation(prog, "iResolution"),
        canvas.width,
        canvas.height,
        dpr,
      );
      gl.uniform1f(
        gl.getUniformLocation(prog, "iTime"),
        (performance.now() - start.current) * 0.001,
      );
      gl.drawArrays(gl.TRIANGLES, 0, 6);
      requestAnimationFrame(render);
    };
    requestAnimationFrame(render);
  }, []);

  return (
    <canvas
      ref={canvasRef}
      style={{
        position: "fixed",
        inset: 0,
        zIndex: -1,
        width: "100%",
        height: "100%",
        display: "block",
        background: "#0d0d0d",
      }}
    />
  );
};

export default BackgroundShader;
