import { useState } from "react";
import axios from "axios";

function Pastebin() {
  const [content, setContent] = useState("");
  const [expiresInMinutes, setExpiresInMinutes] = useState("");
  const [pasteID, setPasteID] = useState(null);
  const [shareableURL, setShareableURL] = useState("");

  // Create a new paste
  const createPaste = async () => {
    try {
      const payload = {
        content,
      };

      // If the user provided expiration, include it
      if (expiresInMinutes) {
        payload.expires_in_minutes = parseInt(expiresInMinutes);
      }

      const response = await axios.post("http://localhost:8080/paste", payload);
      setPasteID(response.data.id);
      setShareableURL(`http://localhost:3000/paste/${response.data.id}`); // React side link
    } catch (error) {
      console.error("Failed to create paste:", error);
    }
  };

  return (
    <div className="container text-center mt-5" style={{ maxWidth: "700px" }}>
      <h2>📜 Pastebin Clone</h2>

      {/* Paste Content */}
      <textarea
        className="form-control mt-3"
        rows="6"
        placeholder="Paste your content here..."
        value={content}
        onChange={(e) => setContent(e.target.value)}
      />

      {/* Expiration Input */}
      <input
        type="number"
        className="form-control mt-3"
        placeholder="Expires in (minutes, optional)"
        value={expiresInMinutes}
        onChange={(e) => setExpiresInMinutes(e.target.value)}
      />

      {/* Create Paste Button */}
      <button className="btn btn-primary mt-3" onClick={createPaste}>
        Create Paste
      </button>

      {/* Success Display */}
      {pasteID && (
        <div className="mt-4">
          <h5> Paste Created!</h5>
          <a href={shareableURL} target="_blank" rel="noopener noreferrer">
            {shareableURL}
          </a>
        </div>
      )}
    </div>
  );
}

export default Pastebin;
