import { useState } from "react";
import axios from "axios"; // For API Calls
import "./styles/IP.css";

function GetIP() {
    const [ipInfo, setIpInfo] = useState(null); // ipInfo stores IP Info from the API
    const [loading, setLoading] = useState(false); // To Show Loading state
    const [error, setError] = useState(null); // 

    const fetchIpInfo = async () => {
        setLoading(true); // Will Start Loading 
        setError(null); // Reseting Perivous Error
        try {
            const response = await axios.get(`${process.env.REACT_APP_API_URL}/get_ip_info/`);
            setIpInfo(response.data); // Store API repsonse in state
        }
        catch (err){
            setError("Failed to fetch IP Info");
        }
        setLoading(false); // Stop Loading
    };
    return (
        <div className="ip-container">
            <p>Find Your Public Ip and GeoLocation details.</p>
            <button onClick={fetchIpInfo} disabled={loading}>{loading ? "Loading...." : "Get IP Info"}</button> 
            {error && <p className="error">{error}</p>}
            
            {ipInfo && (
                <div className="ip-info">
                    <p><strong>IP Address:</strong> {ipInfo.query}</p>
                    <p><strong>Country:</strong> {ipInfo.country}</p>
                    <p><strong>City:</strong> {ipInfo.city}</p>
                    <p><strong>ISP:</strong> {ipInfo.isp}</p>
                    <p><strong>Organization:</strong> {ipInfo.org}</p>
                    <p><strong>Latitude:</strong> {ipInfo.lat}</p>
                    <p><strong>Longitude:</strong> {ipInfo.lon}</p>
                </div>    
            )}
        </div>
    );
}


// How State Updates working
/*
setLoading(true)  start Loading
setIpInfo(resp.data) stores API Response
setLoading(false) stop Loading
*/
export default GetIP;