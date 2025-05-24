import { useState } from "react";
import axios from "axios"; 
import "./styles/DNS.css";

function GetDNS(){
    const [dnsInfo, setDnsInfo] = useState(null);
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState(null);
    const [domain, setDomain] = useState("");  

    const fetchDnsInfo = async () => {
        if (!domain) {
            setError("Please enter a domain.");
            return;
        }

        setLoading(true);
        setError(null);

        try {
  const response = await axios.get(
    `${process.env.REACT_APP_API_URL}/get_dns/?domain=${domain}`
  );
  setDnsInfo(response.data);  
} catch (err) {
  setError("Failed To Fetch DNS Info");
}

        
        setLoading(false);
    };

    return (
        <div className="dns-container">
            <h2>DNS Lookup</h2>
            <input
                type="text"
                placeholder="Enter Domain (example.com)"
                value={domain}
                onChange={(e) => setDomain(e.target.value)}
            />
            <button onClick={fetchDnsInfo} disabled={loading}>
                {loading ? "Loading..." : "Get DNS Info"}
            </button>
            
            {error && <p className="error">{error}</p>}

            {dnsInfo && (
                <div className="dns-info">
                    <p><strong>Domain:</strong> {dnsInfo.domain}</p>

                    <p><strong>IP Addresses:</strong></p>
                    {dnsInfo.ip_addresses.length > 0 ? (
                        <ul>
                            {dnsInfo.ip_addresses.map((ip, index) => (
                                <li key={index}>{ip}</li>
                            ))}
                        </ul>
                    ) : <p>N/A</p>}

                    <p><strong>Name Servers:</strong></p>
                    {dnsInfo.name_servers.length > 0 ? (
                        <ul>
                            {dnsInfo.name_servers.map((server, index) => (
                                <li key={index}>{server}</li>
                            ))}
                        </ul>
                    ) : <p>N/A</p>}

                    <p><strong>Mail Servers (MX):</strong></p>
                    {dnsInfo.mail_servers.length > 0 ? (
                        <ul>
                            {dnsInfo.mail_servers.map((mx, index) => (
                                <li key={index}>{mx}</li>
                            ))}
                        </ul>
                    ) : <p>N/A</p>}

                    <p><strong>CNAME:</strong> {dnsInfo.cname || "N/A"}</p>

                    <p><strong>TXT Records:</strong></p>
                    {dnsInfo.txt_records.length > 0 ? (
                        <ul>
                            {dnsInfo.txt_records.map((txt, index) => (
                                <li key={index}>{txt}</li>
                            ))}
                        </ul>
                    ) : <p>N/A</p>}
                </div>
            )}
        </div>
    );
}

export default GetDNS;
 