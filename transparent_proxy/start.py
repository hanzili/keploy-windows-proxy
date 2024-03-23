from mitmproxy import ctx, http

def request(flow: http.HTTPFlow):
    print(flow.request.port)
    # Example: Intercepting traffic to and from port 3000
    if flow.request.port in [3000]:
        try:
            # Construct data for a request
            data = {
                "type": "request",
                "direction": "outgoing" if flow.request.is_outgoing else "incoming",
                "url": flow.request.url,
                "method": flow.request.method,
                "headers": dict(flow.request.headers),
                "body": flow.request.get_text(strict=False)
            }
            send_to_go_service(data)
        except Exception as e:
            ctx.log.error(str(e))

def response(flow: http.HTTPFlow):
    print(flow.response.port)
    if flow.request.port in [3000]:
        try:
            # Construct data for a response
            data = {
                "type": "response",
                "direction": "incoming" if flow.response.is_outgoing else "outgoing",
                "status_code": flow.response.status_code,
                "headers": dict(flow.response.headers),
                "body": flow.response.get_text(strict=False)
            }
            send_to_go_service(data)
        except Exception as e:
            ctx.log.error(str(e))

def send_to_go_service(data):
    import json
    import requests
    # Convert data to JSON
    json_data = json.dumps(data)
    # Send JSON data to the Go service
    requests.post("http://localhost:9090/", data=json_data, headers={'Content-Type': 'application/json'})
