import { useState } from "react";
import axios from "axios";
import "./styles/Shortner.css";  // ✅ Importing CSS

function Shortener() {
  const [url, setUrl] = useState("");
  const [customId, setCustomId] = useState("");
  const [shortUrl, setShortUrl] = useState("");

  const shortenUrl = async () => {
    try {
      const response = await axios.post("http://localhost:8080/shorturl", {
        url,
        custom_id: customId || undefined,
      });

      setShortUrl(`http://localhost:8080/redirect/${response.data.new_url}`);
    } catch (error) {
      alert("Error shortening URL");
    }
  };

  return (
    <div className="container">
      <h2>Shorten a URL</h2>
      <input
        type="text"
        placeholder="Enter URL"
        value={url}
        onChange={(e) => setUrl(e.target.value)}
      />
      <input
        type="text"
        placeholder="Custom ID (optional)"
        value={customId}
        onChange={(e) => setCustomId(e.target.value)}
      />
      <button onClick={shortenUrl}>Shorten</button>

      {shortUrl && (
        <div>
          <h3>Shortened URL:</h3>
          <a href={shortUrl} target="_blank" rel="noopener noreferrer">
            {shortUrl}
          </a>
        </div>
      )}
    </div>
  );
}

export default Shortener;
