import { Link } from "react-router-dom";
import "./styles/Navbar.css"; 

function Navbar() {
  return (
    <nav className="navbar">
      <h2 className="navbarText">Services</h2>
      <div>
        <Link to="/">Home</Link>
        <Link to="/qrcode">QR Code</Link>
        <Link to="/iplookup">IP Lookup</Link> 
        <Link to="/dnslookup">DNS Lookup</Link>
        <Link to="/fileuploader">File Uploader</Link>
        <Link to="/pastebin">Paste Bin</Link>
        <Link to="/convert">Convert-JSON-YAML-TOML</Link>
      </div>
    </nav>
  );
}

export default Navbar;
