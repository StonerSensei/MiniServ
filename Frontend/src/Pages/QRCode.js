import { useState } from "react";
import axios from "axios";
import "./styles/QRCode.css"; 

function QRCode() {
  const [content, setContent] = useState("");
  const [fileName, setFileName] = useState("qrcode.png");
  const [qrUrl, setQrUrl] = useState("");

  const generateQR = async () => {
    try {
      const response = await axios.post(
  `${process.env.REACT_APP_API_URL}/generate_qr/`,
  {
    content,
    file_name: fileName,
  }
);

      setQrUrl(response.data.url);
    } catch (error) {
      alert("Error generating QR Code");
    }
  };

  return (
    <div className="container">
      <h2>Generate a QR Code</h2>
      <input
        type="text"
        placeholder="Enter text or URL"
        value={content}
        onChange={(e) => setContent(e.target.value)}
      />
      <button onClick={generateQR}>Generate</button>

      {qrUrl && (
        <div>
          <h3>QR Code:</h3>
          <img src={qrUrl} alt="Generated QR Code" />
          <a href={qrUrl} download>Download</a>
        </div>
      )}
    </div>
  );
}

export default QRCode;
