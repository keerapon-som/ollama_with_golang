import requests
import json

url = "http://localhost:8080/api/generateCompletionStream"

payload = json.dumps({
    "model": "llama3.2:1b",
    "prompt": "this is eiei right ?",
    "Stream": True
})

headers = {
    'Content-Type': 'application/json'
}

s = requests.Session()
r = s.post(url, data=payload, headers=headers, stream=True)
print(r.status_code)
for line in r.iter_content(chunk_size=10):
    if line:
        print(line)