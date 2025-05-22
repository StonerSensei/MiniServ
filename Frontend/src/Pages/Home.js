import { useNavigate } from "react-router-dom";
import "./styles/Home.css"; 

function Home() {
  const navigate = useNavigate();

  return (
    <>
      <div className="home-container">
        <h1>Services</h1>
      </div>
      <div className="home-container">
        <button onClick={() => navigate("/qrcode")}>QR Code Generator</button>
      </div>

      <div className="home-container">
        <button onClick={() => navigate("/iplookup")}>IP Lookup</button>
      </div>
      <div className="home-container">
        <button onClick={() => navigate("/dnslookup")}>DNS Lookup</button>
      </div>
      <div className="home-container">
        <button onClick={() => navigate("/fileuploader")}>File Uploader</button>
      </div>
      <div className="home-container">
        <button onClick={() => navigate("/pastebin")}>Paste Bin</button>
      </div>
      <div className="home-container">
        <button onClick={() => navigate("/convert")}>Convert-JSON-YAML-TOML</button>
      </div>
    </>
  );
}

export default Home;
