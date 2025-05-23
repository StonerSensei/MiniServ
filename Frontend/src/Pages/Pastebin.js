import { useState } from "react";
import axios from "axios";

function Pastebin() {
  const [content, setContent] = useState("");
  const [expiresInMinutes, setExpiresInMinutes] = useState("");
  const [pasteID, setPasteID] = useState(null);
  const [shareableURL, setShareableURL] = useState("");

  const createPaste = async () => {
    try {
      const payload = { content };

      if (expiresInMinutes) {
        payload.expires_in_minutes = parseInt(expiresInMinutes);
      }

      const response = await axios.post(
        `${process.env.REACT_APP_API_URL}/paste`,
        payload
      );

      const id = response.data.id;
      setPasteID(id);
      setShareableURL(`${window.location.origin}/paste/${id}`);
    } catch (error) {
      console.error("Failed to create paste:", error);
    }
  };

  return (
    <div className="container text-center mt-5" style={{ maxWidth: "700px" }}>
      <h2>📜 Pastebin Clone</h2>

      <textarea
        className="form-control mt-3"
        rows="6"
        placeholder="Paste your content here..."
        value={content}
        onChange={(e) => setContent(e.target.value)}
      />

      <input
        type="number"
        className="form-control mt-3"
        placeholder="Expires in (minutes, optional)"
        value={expiresInMinutes}
        onChange={(e) => setExpiresInMinutes(e.target.value)}
      />

      <button className="btn btn-primary mt-3" onClick={createPaste}>
        Create Paste
      </button>

      {pasteID && (
        <div className="mt-4">
          <h5>✅ Paste Created!</h5>
          <a href={shareableURL} target="_blank" rel="noopener noreferrer">
            {shareableURL}
          </a>
        </div>
      )}
    </div>
  );
}

export default Pastebin;
