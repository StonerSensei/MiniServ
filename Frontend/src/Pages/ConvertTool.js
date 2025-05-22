import React, { useState } from "react";
import axios from "axios";

function ConvertTool() {
  const [input, setInput] = useState("");
  const [inputFormat, setInputFormat] = useState("json");
  const [outputFormat, setOutputFormat] = useState("yaml");
  const [converted, setConverted] = useState("");
  const [error, setError] = useState("");

  const handleConvert = async () => {
    setError("");
    setConverted("");

    try {
      // Make sure we have valid input
      if (!input.trim()) {
        setError("Please enter some content to convert");
        return;
      }

      // Set the right Content-Type for the input format
      let contentType;
      switch (inputFormat) {
        case "json":
          contentType = "application/json";
          break;
        case "yaml":
          contentType = "application/x-yaml";
          break;
        case "toml":
          contentType = "application/toml";
          break;
        default:
          contentType = "text/plain";
      }

      const res = await axios.post(
        `http://localhost:8080/convert?input=${inputFormat}&output=${outputFormat}`,
        input,
        {
          headers: {
            "Content-Type": contentType,
          },
          withCredentials: true, // Important for CORS with credentials
        }
      );
      
      // Handle the response based on format
      if (typeof res.data === "object") {
        setConverted(JSON.stringify(res.data, null, 2));
      } else {
        setConverted(res.data);
      }
    } catch (err) {
      console.error("Conversion error:", err);
      setError(
        err.response?.data || 
        "Failed to convert. Please check your input format and try again."
      );
    }
  };

  // Example placeholders for each format
  const getPlaceholder = () => {
    switch (inputFormat) {
      case "json":
        return '{\n  "name": "John Doe",\n  "age": 30,\n  "isActive": true\n}';
      case "yaml":
        return 'name: John Doe\nage: 30\nisActive: true';
      case "toml":
        return 'name = "John Doe"\nage = 30\nisActive = true';
      default:
        return "";
    }
  };

  return (
    <div className="container mt-5">
      <h2 className="mb-4">🛠️ Config Converter (JSON ⇄ YAML ⇄ TOML)</h2>

      <div className="mb-3">
        <label>Input Format:</label>
        <select
          className="form-select"
          value={inputFormat}
          onChange={(e) => setInputFormat(e.target.value)}
        >
          <option value="json">JSON</option>
          <option value="yaml">YAML</option>
          <option value="toml">TOML</option>
        </select>
      </div>

      <div className="mb-3">
        <label>Output Format:</label>
        <select
          className="form-select"
          value={outputFormat}
          onChange={(e) => setOutputFormat(e.target.value)}
        >
          <option value="json">JSON</option>
          <option value="yaml">YAML</option>
          <option value="toml">TOML</option>
        </select>
      </div>

      <div className="mb-3">
        <label>Input Content:</label>
        <textarea
          className="form-control"
          rows="6"
          value={input}
          onChange={(e) => setInput(e.target.value)}
          placeholder={getPlaceholder()}
        />
      </div>

      <button className="btn btn-primary" onClick={handleConvert}>
        Convert
      </button>

      {error && (
        <div className="alert alert-danger mt-3">
          <strong>Error:</strong> {error}
        </div>
      )}

      {converted && (
        <div className="mt-4">
          <h5>Converted Output:</h5>
          <pre className="p-3 bg-light border rounded">{converted}</pre>
          <button 
            className="btn btn-sm btn-secondary mt-2"
            onClick={() => {
              navigator.clipboard.writeText(converted);
              alert("Copied to clipboard!");
            }}
          >
            Copy to clipboard
          </button>
        </div>
      )}
    </div>
  );
}

export default ConvertTool;