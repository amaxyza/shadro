import { useState } from "react";

interface ToolbarProps {
  code: string;                       // current shader source
  setProgramName: (s: string) => void;
  saveShader: (name: string) => Promise<void>;
}

const downloadShader = (name: string, code: string) => {
  const blob = new Blob([code], { type: "text/plain" });
  const url = URL.createObjectURL(blob);
  const link = document.createElement("a");
  link.href = url;
  link.download = `${name || "shader"}.glsl`;
  document.body.appendChild(link);
  link.click();
  document.body.removeChild(link);
  URL.revokeObjectURL(url);
};

export default function EditorToolbar({
  code,
  setProgramName,
  saveShader,
}: ToolbarProps) {
  const [showModal, setShowModal] = useState(false);
  const [tempName, setTempName] = useState("");

  return (
    <>
      <div className="toolbar">
        <button onClick={() => downloadShader(tempName, code)}>Download</button>
        <button onClick={() => setShowModal(true)}>Save</button>
        <button onClick={() => null}>Run</button>
      </div>

      {showModal && (
        <>
          <div className="backdrop" onClick={() => setShowModal(false)} />
          <div className="modal">
            <h2>Save Shader</h2>
            <input
              type="text"
              placeholder="Enter program name"
              value={tempName}
              onChange={(e) => {
                setTempName(e.target.value);
                setProgramName(e.target.value);
              }}
            />
            <button
              onClick={async () => {
                await saveShader(tempName);
                setShowModal(false);
              }}
            >
              Confirm
            </button>
          </div>
        </>
      )}
    </>
  );
}