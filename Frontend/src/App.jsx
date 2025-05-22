import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Navbar from "./Components/Navbar";
import Home from "./Pages/Home";
import QRCode from "./Pages/QRCode";
import IP from "./Pages/IP";
import DNS from "./Pages/DNS";
import "./styles/global.css";
import FileUploader from "./Pages/FileUploader";
import Pastebin from "./Pages/Pastebin";
import ViewPaste from "./Pages/ViewPaste";
import ConvertTool from "./Pages/ConvertTool";


function App() {
  return (
    <Router>
      <Navbar />
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/qrcode" element={<QRCode />} />
        <Route path="/iplookup" element={<IP/>}/>
        <Route path="/dnslookup" element={<DNS/>}/>
        <Route path="/fileuploader" element={<FileUploader/>}/>
        <Route path="/pastebin" element={<Pastebin/>}/>
        <Route path="/paste/:pasteID" element={<ViewPaste />} />
        <Route path="/convert" element={<ConvertTool/>}/>
      </Routes>
    </Router>
  );
}
export default App;
