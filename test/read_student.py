import requests
import json
from public import URL, POST_DATA as token


r = requests.post(URL + '/read/student', json=token)
print(json.dumps(r.json(), indent=4, ensure_ascii=False))
