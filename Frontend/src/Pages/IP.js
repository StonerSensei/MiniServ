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
            const response = await axios.get("http://localhost:8080/get_ip_info/");
            setIpInfo(response.data); // Store API repsonse in state
        }
        catch (err){
            setError("Failed to fetch IP Info");
        }
        setLoading(false); // Stop Loading
    };
    // onCLick Button below: First Disabled the button while Loading and if Loading true show Loading.... else show Get IP Info
    // then display error if any
    // If ipInfo is not null display the details
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