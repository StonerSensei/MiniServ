import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import axios from "axios";

function ViewPaste() {
  const { pasteID } = useParams(); // ✅ Extract pasteID from URL
  const [content, setContent] = useState("");
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchPaste = async () => {
      try {
        const response = await axios.get(`http://localhost:8080/paste/${pasteID}`);
        setContent(response.data); // ✅ Load paste content
      } catch (err) {
        setError("Paste not found");
      }
    };

    fetchPaste();
  }, [pasteID]);

  return (
    <div className="container text-center mt-5">
      <h2>📄 Viewing Paste: {pasteID}</h2>
      {error ? (
        <p className="text-danger">{error}</p>
      ) : (
        <pre className="border p-3 text-start">{content}</pre> // ✅ Show content in a nice box
      )}
    </div>
  );
}

export default ViewPaste;
