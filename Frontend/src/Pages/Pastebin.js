import { useState } from "react";
import axios from "axios";

function Pastebin() {
  const [content, setContent] = useState("");
  const [pasteID, setPasteID] = useState(null);
  const [retrievedPaste, setRetrievedPaste] = useState("");

  // Create paste
  const createPaste = async () => {
    try {
      const response = await axios.post("http://localhost:8080/paste", { content });
      setPasteID(response.data.id);
    } catch (error) {
      console.error("Failed to create paste:", error);
    }
  };

  // Get paste
  const getPaste = async (id) => {
    try {
      const response = await axios.get(`http://localhost:8080/paste/${id}`);
      setRetrievedPaste(response.data.content);
    } catch (error) {
      console.error("Paste not found:", error);
    }
  };

  return (
    <div className="container text-center mt-5">
      <h2>📜 Pastebin Clone</h2>
      <textarea className="form-control mt-3" rows="4" onChange={(e) => setContent(e.target.value)} />
      <button className="btn btn-primary mt-3" onClick={createPaste}>Create Paste</button>

      {pasteID && (
        <div className="mt-4">
          <h5>✅ Paste Created! Share this ID: {pasteID}</h5>
          <button className="btn btn-success mt-2" onClick={() => getPaste(pasteID)}>View Paste</button>
        </div>
      )}

      {retrievedPaste && (
        <div className="mt-4">
          <h5>📄 Retrieved Paste:</h5>
          <p className="border p-3">{retrievedPaste}</p>
        </div>
      )}
    </div>
  );
}

export default Pastebin;
