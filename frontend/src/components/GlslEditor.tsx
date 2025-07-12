import React, { useState } from 'react'
import './Editor.css'
import CodeMirror from "@uiw/react-codemirror"
import { cpp } from "@codemirror/lang-cpp"

const GlslEditor: React.FC = () => {
  const boilerplate: string = `void mainImage( out vec4 fragColor, in vec2 fragCoord )
{
    // Normalized pixel coordinates (from 0 to 1)
    vec2 uv = fragCoord/iResolution.xy;

    // Time varying pixel color
    vec3 col = 0.5 + 0.5*cos(iTime+uv.xyx+vec3(0,2,4));

    // Output to screen
    fragColor = vec4(col,1.0);
}`

  return (
    <>
      <div className="editor-buttons">
        <button className="editor-button">Download</button>
        <button className="editor-button">Save</button>
        <button className="editor-button run">Run</button>
      </div>
      <div className="editor-layout">
        <CodeMirror
          value={boilerplate} 
          height="75vh"
          width="100vh" 
          theme="dark"  
          extensions={
              [cpp()]
          }
        />
      </div>
    </>
  )
}

export default GlslEditor