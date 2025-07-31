import React, { useEffect, useState } from 'react'
import './Editor.css'
import CodeMirror, { keymap, ViewUpdate } from "@uiw/react-codemirror"
import { cpp } from "@codemirror/lang-cpp"

interface User {
  id: number;
  username: string;
}

/*
type glslProgram struct {
	ID       int       `json:"id"`
	Owner_id int       `json:"owner_id"`
	Name     string    `json:"name"`
	Glsl     string    `json:"glsl"`
	Created  time.Time `json:"created"`
	Updated  time.Time `json:"updated"`
}
*/

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

interface GlslEditorProps {
  program_id?: number;
  code?: string;
  onChange: (value: string) => void;
}

const GlslEditor = ( {program_id = -1, code = `void mainImage(out vec4 fragColor, in vec2 fragCoord) {
    vec2 uv = fragCoord / iResolution.xy;
    fragColor = vec4(uv, 0.5 + 0.5 * sin(iTime), 1.0);
}`, onChange}: GlslEditorProps ) => {
  const [shaderCode, setShaderCode] = useState(code);

  const [showModal, setShowModal] = useState(false);
  const [programName, setProgramName] = useState('');
  const [finalName, setFinalName] = useState('Untitled Shader');
  const [user, setUser] = useState<User | null>();

  useEffect( () => {
    const replace_program = async () => {
      if (program_id !== -1) {
        const res = await fetch('/api/programs/' + program_id)
        if (res.ok) {
          const { name, glsl } = await res.json()
          setShaderCode(glsl)
          setProgramName(name)
          setFinalName(name)
        }
      }
      else {
        console.error('unable to retrive program, reverting to default display')
      }
    }

    replace_program();
  }, [])

  return (
    <>
      
      <div className="editor-layout" style={{ marginTop: "0%" }}>
        <div className="program-name">
          <h2>{finalName || "Untitled Shader"}</h2>
        </div>
        <CodeMirror
          value={shaderCode} 
          height="75vh"
          width="100vh" 
          theme="dark"  
          extensions={
              [cpp()]
          }
          onChange={(value: string, viewUpdate: ViewUpdate) => { onChange(value) }}
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
            <button onClick={async () => {
              // api call to /me to verify user login
              const res = await fetch('/api/me', {credentials: 'include'})
              if (res.ok) {
                const data: User = await res.json();
                setUser(data);
              }
              else {
                console.log("failed to get login")
                setUser(null)
              }
              setFinalName(programName);
              if (user) {
                // api call to POST program with user ID, user name, and shader source code as inputs
                const program_input: {owner_id: Number, program_name: string, source: string} = {
                  owner_id: user.id,
                  program_name: finalName,
                  source: shaderCode
                }

                const res = await fetch('/api/programs', {
                  method: 'POST',
                  headers: { 'Content-Type': 'application/json'},
                  body: JSON.stringify(program_input)
                })

                if (res.ok) {
                  console.log('Saved as:', programName, "for user", user.username);
                }
                else {
                  console.error('error with post request to backend server')
                }
                
              }
              setShowModal(false);
            }}>Confirm</button>
          </div>
        </>
      )}
    </>
  )
}

export default GlslEditor