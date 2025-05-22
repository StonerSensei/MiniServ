import { useState } from "react";
import axios from "axios";
import "bootstrap/dist/css/bootstrap.min.css";

function FileUploader() {
  const [file, setFile] = useState(null);
  const [uploadURL, setUploadURL] = useState("");
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);

  // Handle file selection
  const handleFileChange = (event) => {
    setFile(event.target.files[0]);
  };

  // Upload file to the backend
  const handleUpload = async () => {
    if (!file) {
      alert("Please select a file to upload!");
      return;
    }
  
    setLoading(true);
    setError(null);
    const formData = new FormData();
    formData.append("file", file);
  
    try {
      const response = await axios.post("http://localhost:8080/upload", formData, {
        headers: {
          "Content-Type": "multipart/form-data",
        }
      });
  
      setUploadURL(response.data.url);
    } catch (err) {
      setError("Error uploading file. Please try again.");
      console.error("Upload error:", err.response || err.message);
    }
  
    setLoading(false);
  };
  
  

  return (
    <div className="container mt-5 text-center">
      <h2>📂 Temporary File Uploader</h2>
      <p>Upload a file and get a temporary download link.</p>

      <input type="file" className="form-control mt-3" onChange={handleFileChange} />
      <button className="btn btn-primary mt-3" onClick={handleUpload} disabled={loading}>
        {loading ? "Uploading..." : "Upload File"}
      </button>

      {error && <p className="text-danger mt-3">{error}</p>}

      {uploadURL && (
        <div className="mt-4">
          <h5>✅ File Uploaded Successfully!</h5>
          <p><strong>Download Link (Expires in 5 minutes):</strong></p>
          <a href={uploadURL} target="_blank" rel="noopener noreferrer">
            {uploadURL}
          </a>
        </div>
      )}
    </div>
  );
}

export default FileUploader;
