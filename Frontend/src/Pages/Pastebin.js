import { useState } from "react";
import axios from "axios";

function Pastebin() {
  const [content, setContent] = useState("");
  const [pasteID, setPasteID] = useState(null);
  const [shareableURL, setShareableURL] = useState("");

  // Create a new paste
  const createPaste = async () => {
    try {
      const response = await axios.post("http://localhost:8080/paste", { content });
      setPasteID(response.data.id);
      setShareableURL(`http://localhost:3000/paste/${response.data.id}`); // ✅ React URL
    } catch (error) {
      console.error("Failed to create paste:", error);
    }
  };

  return (
    <div className="container text-center mt-5">
      <h2>📜 Pastebin Clone</h2>
      <textarea className="form-control mt-3" rows="4" onChange={(e) => setContent(e.target.value)} />
      <button className="btn btn-primary mt-3" onClick={createPaste}>Create Paste</button>

      {pasteID && (
        <div className="mt-4">
          <h5>✅ Paste Created!</h5>
          <a href={shareableURL} target="_blank" rel="noopener noreferrer">{shareableURL}</a>
        </div>
      )}
    </div>
  );
}

export default Pastebin;
