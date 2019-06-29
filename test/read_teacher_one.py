import requests
import json
from public import URL, POST_DATA as token

data = {
    **token, **{
        "jobId": "0000000002"
    }
}


r = requests.post(URL + '/read/teacher/one', json=data)
print(json.dumps(r.json(), indent=4, ensure_ascii=False))
