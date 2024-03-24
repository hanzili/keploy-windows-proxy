from mitmproxy import http
import requests
import logging

# Set up basic logging
logging.basicConfig(level=logging.INFO, format='%(asctime)s - %(message)s')
logger = logging.getLogger()

class SimpleForwardRequest:
    def __init__(self):
        self.target_url = "http://localhost:9090"
        
    def response(self, flow: http.HTTPFlow) -> None:
        # Simplified forwarding logic
        try:
            forwarded_response = {
                "method": flow.request.method,
                "url": flow.request.url,
                "headers": dict(flow.request.headers.items()),
                "data": flow.request.content,
                "params": flow.request.query
            }
            
            requests.post("http://localhost:9090", json=forwarded_response)

        except Exception as e:
            logger.error(f"Error forwarding request: {e}")

    def request(self, flow: http.HTTPFlow) -> None:
        # Simplified forwarding logic
        try:
            forwarded_response = {
                "method": flow.request.method,
                "url": flow.request.url,
                "headers": dict(flow.request.headers.items()),
                "data": flow.request.content,
                "params": flow.request.query
            }
            requests.post("http://localhost:9090", json=forwarded_response)

        except Exception as e:
            logger.error(f"Error forwarding request: {e}")

addons = [
    SimpleForwardRequest()
]
