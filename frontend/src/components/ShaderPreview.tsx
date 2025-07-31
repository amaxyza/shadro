import { useEffect, useRef } from "react";

/**
 * Live WebGL preview for Shadro fragment shaders.
 *
 * EXPECTED USAGE:
 *
 *  <ShaderPreview source={code} />
 *
 * Where `code` is the raw GLSL fragment shader body supplied by the user.
 * The component injects standard Shadertoy‑style uniforms (iTime, iResolution)
 * and recompiles whenever the source changes.
 */
interface ShaderPreviewProps {
  /** Raw GLSL fragment shader source typed by the user */
  source: string;
}

// Simple full‑screen triangle vertex shader
const VS_SOURCE = `#version 300 es
precision highp float;
layout(location = 0) in vec2 aPosition;
void main() {
    gl_Position = vec4(aPosition, 0.0, 1.0);
}`;

/** Wrap user code in a Shadertoy‑compatible program so we can call mainImage(). */
function buildFragmentSource(userSource: string): string {
  return `#version 300 es
precision highp float;

out vec4 fragColor;
uniform vec3 iResolution; // (width, height, pixelRatio)
uniform float iTime;      // seconds since start

// >>> User code >>>
${userSource.trim()}
// <<< User code <<<

void main() {
    vec2 fragCoord = gl_FragCoord.xy;
    mainImage(fragColor, fragCoord);
}`;
}

const ShaderPreview: React.FC<ShaderPreviewProps> = ({ source }) => {
  const canvasRef = useRef<HTMLCanvasElement>(null);
  const glRef = useRef<WebGL2RenderingContext | null>(null);
  const programRef = useRef<WebGLProgram | null>(null);
  const startTime = useRef<number>(performance.now());

  /** Compile a shader of given type, returning null on error */
  const compileShader = (
    gl: WebGL2RenderingContext,
    type: number,
    src: string
  ): WebGLShader | null => {
    const shader = gl.createShader(type);
    if (!shader) return null;
    gl.shaderSource(shader, src);
    gl.compileShader(shader);
    if (!gl.getShaderParameter(shader, gl.COMPILE_STATUS)) {
      console.warn(gl.getShaderInfoLog(shader));
      return null;
    }
    return shader;
  };

  /** Link shaders into a program */
  const linkProgram = (
    gl: WebGL2RenderingContext,
    vs: WebGLShader,
    fs: WebGLShader
  ): WebGLProgram | null => {
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
  };

  /** Build quad geometry just once */
  const initGeometry = (gl: WebGL2RenderingContext) => {
    const vao = gl.createVertexArray();
    gl.bindVertexArray(vao);
    const vbo = gl.createBuffer();
    gl.bindBuffer(gl.ARRAY_BUFFER, vbo);
    // Two‑triangle full‑screen quad
    gl.bufferData(
      gl.ARRAY_BUFFER,
      new Float32Array([
        -1, -1, 1, -1, -1, 1, -1, 1, 1, -1, 1, 1,
      ]),
      gl.STATIC_DRAW
    );
    gl.enableVertexAttribArray(0);
    gl.vertexAttribPointer(0, 2, gl.FLOAT, false, 0, 0);
  };

  /** Recompile program whenever the user source changes */
  useEffect(() => {
    const canvas = canvasRef.current;
    if (!canvas) return;
    const gl = (glRef.current ||= canvas.getContext("webgl2"));
    if (!gl) return;

    // Initialize VBO/VAO on first run
    if (!programRef.current) {
      initGeometry(gl);
    }

    const vs = compileShader(gl, gl.VERTEX_SHADER, VS_SOURCE);
    const fs = compileShader(gl, gl.FRAGMENT_SHADER, buildFragmentSource(source));
    if (!vs || !fs) return;

    const prog = linkProgram(gl, vs, fs);
    programRef.current = prog;
  }, [source]);

  /** Main render loop */
  useEffect(() => {
    const canvas = canvasRef.current;
    if (!canvas) return;
    const gl = (glRef.current ||= canvas.getContext("webgl2"));
    if (!gl) return;

    let running = true;
    const render = () => {
      if (!running) return;
      // Resize
      const dpr = window.devicePixelRatio || 1;
      const displayWidth = Math.floor(canvas.clientWidth * dpr);
      const displayHeight = Math.floor(canvas.clientHeight * dpr);
      if (canvas.width !== displayWidth || canvas.height !== displayHeight) {
        canvas.width = displayWidth;
        canvas.height = displayHeight;
      }
      gl.viewport(0, 0, canvas.width, canvas.height);

      const prog = programRef.current;
      if (prog) {
        gl.useProgram(prog);
        const resLoc = gl.getUniformLocation(prog, "iResolution");
        const timeLoc = gl.getUniformLocation(prog, "iTime");
        if (resLoc) gl.uniform3f(resLoc, canvas.width, canvas.height, dpr);
        if (timeLoc) gl.uniform1f(timeLoc, (performance.now() - startTime.current) / 1000);
        gl.drawArrays(gl.TRIANGLES, 0, 6);
      }
      requestAnimationFrame(render);
    };
    requestAnimationFrame(render);
    return () => {
      running = false;
    };
  }, []);

  return (
    <canvas
      ref={canvasRef}
      style={{ width: "100%", height: "100%", display: "block", background: "#000" }}
    />
  );
};

export default ShaderPreview;
