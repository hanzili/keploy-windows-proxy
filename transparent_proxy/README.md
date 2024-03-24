### enable transparent proxying on windows
https://docs.mitmproxy.org/stable/howto-transparent/#windows


### start the mitmproxy in transparent mode
run the command at admin previliege.
mitmproxy --mode transparent --showhost -s start.py

### send a request through proxy in regular mode
curl --proxy http://localhost:8080 -k https://google.com