import React, { useState } from "react"
import { useParams } from "react-router-dom"
import Header from "../components/Header"
import GlslEditor from "../components/GlslEditor"
import ShaderPreview from "../components/ShaderPreview"
import EditorToolbar from "../components/EditorToolbar"

const Create: React.FC = () => { 
    const { id } = useParams()
    const [code, setCode] = useState(`void mainImage(out vec4 fragColor, in vec2 fragCoord) {
    vec2 uv = fragCoord / iResolution.xy;
    fragColor = vec4(uv, 0.5 + 0.5 * sin(iTime), 1.0);
  }`);
    const [programName, setProgramName] = useState("Untitled Shader");

    const saveShader = async (name: string) => {
      const meRes = await fetch("/api/me", { credentials: "include" });
      if (!meRes.ok) return console.error("login check failed");
      const user = await meRes.json();

      const body = {
        owner_id: user.id,
        program_name: name,
        source: code,
    };

    const postRes = await fetch("/api/programs", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(body),
    });

    if (!postRes.ok) return console.error("POST /api/programs failed");
    console.log("Saved as", name, "for user", user.username);
  };

    return (
        <>
            
            <Header />
            <EditorToolbar code={code} setProgramName={setProgramName} saveShader={saveShader} />
            <div style={{ display: "flex", height: "100vh" }}>
                <GlslEditor program_id={id ? parseInt(id, 10) : undefined} code={code} onChange={setCode} programName={programName} />
                <ShaderPreview source={code} />
            </div>
        </>
    )
}

export default Create