import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Navbar from "./Components/Navbar";
import Home from "./Pages/Home";
import QRCode from "./Pages/QRCode";
import Shortener from "./Pages/Shortener";
import IP from "./Pages/IP";
import DNS from "./Pages/DNS";
import "./styles/global.css";  // ✅ Importing global CSS
import FileUploader from "./Pages/FileUploader";
import Pastebin from "./Pages/Pastebin";
import ViewPaste from "./Pages/ViewPaste";


function App() {
  return (
    <Router>
      <Navbar />
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/shortener" element={<Shortener />} />
        <Route path="/qrcode" element={<QRCode />} />
        <Route path="/iplookup" element={<IP/>}/>
        <Route path="/dnslookup" element={<DNS/>}/>
        <Route path="/fileuploader" element={<FileUploader/>}/>
        <Route path="/pastebin" element={<Pastebin/>}/>
        <Route path="/paste/:pasteID" element={<ViewPaste />} />
      </Routes>
    </Router>
  );
}
export default App;
