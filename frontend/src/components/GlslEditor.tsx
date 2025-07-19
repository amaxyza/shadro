import React, { useState } from 'react'
import './Editor.css'
import CodeMirror, { keymap, ViewUpdate } from "@uiw/react-codemirror"
import { cpp } from "@codemirror/lang-cpp"

const downloadShader = (name: string, code: string) => {
  const blob = new Blob([code], {type: "text/plain"});
  const url = URL.createObjectURL(blob);

  const link = document.createElement('a');
  link.href = url;
  link.download = name + '.glsl';

  document.body.appendChild(link);
  link.click();
  document.body.removeChild(link);

  URL.revokeObjectURL(url);
}

const GlslEditor: React.FC = () => {
  const [shaderCode, setShaderCode] = useState(`void mainImage( out vec4 fragColor, in vec2 fragCoord )
{
    // Normalized pixel coordinates (from 0 to 1)
    vec2 uv = fragCoord/iResolution.xy;

    // Time varying pixel color
    vec3 col = 0.5 + 0.5*cos(iTime+uv.xyx+vec3(0,2,4));

    // Output to screen
    fragColor = vec4(col,1.0);
}`);

  const [showModal, setShowModal] = useState(false)
  const [programName, setProgramName] = useState('')
  const [finalName, setFinalName] = useState('Untitled Shader')
  return (
    <>
      <div className="editor-buttons">
        <p className="editor-text">{finalName}</p>
        <button className="editor-button" onClick= {() => {downloadShader(finalName, shaderCode)}}>Download</button>
        <button className="editor-button" onClick= {() => { setShowModal(true)}}>Save</button>
        <button className="editor-button run">Run</button>
      </div>
      <div className="editor-layout">
        <CodeMirror
          value={shaderCode} 
          height="75vh"
          width="100vh" 
          theme="dark"  
          extensions={
              [cpp()]
          }
          onChange={(value: string, viewUpdate: ViewUpdate) => { setShaderCode(value)}}
        />
      </div>

      {showModal && (
        <>
          <div className="backdrop" onClick={() => setShowModal(false)} />
          <div className="modal">
            <h2>Save Shader</h2>
            <input
              type="text"
              placeholder="Enter program name"
              value={programName}
              onChange={(e) => setProgramName(e.target.value)}
            />
            <button onClick={() => {
              // api call to /me to verify user login
              // api call to POST program with user ID and shader source code as inputs
              console.log('Saved as:', programName);
              
              setFinalName(programName);
              setShowModal(false);
            }}>Confirm</button>
          </div>
        </>
      )}
    </>
  )
}

export default GlslEditor