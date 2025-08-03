import { useEffect, useState } from 'react'
import './Editor.css'
import CodeMirror from "@uiw/react-codemirror"
import { cpp } from "@codemirror/lang-cpp"

interface GlslEditorProps {
  program_id?: number;
  code?: string;
  onChange: (value: string) => void;
  programName: string;
}

const GlslEditor = ( {program_id = -1, code = `void mainImage(out vec4 fragColor, in vec2 fragCoord) {
    vec2 uv = fragCoord / iResolution.xy;
    fragColor = vec4(uv, 0.5 + 0.5 * sin(iTime), 1.0);
}`, onChange, programName}: GlslEditorProps ) => {
  const [shaderCode, setShaderCode] = useState(code);

  useEffect( () => {
    const replace_program = async () => {
      if (program_id !== -1) {
        const res = await fetch('/api/programs/' + program_id)
        if (res.ok) {
          const { name, glsl } = await res.json()
          setShaderCode(glsl)
          programName = name
        }
      } else {
        console.error('unable to retrive program, reverting to default display')
      }
    }
    replace_program();
  }, [])

  return (
    <>
      <div className="editor-layout" style={{ marginTop: "0%" }}>
        <div className="program-name">
          <h2>{programName}</h2>
        </div>
        <CodeMirror
          value={shaderCode} 
          height="75vh"
          width="100vh" 
          theme="dark"  
          extensions={
              [cpp()]
          }
          onChange={(value: string) => { onChange(value) }}
        />
      </div>
    </>
  )
}

export default GlslEditor