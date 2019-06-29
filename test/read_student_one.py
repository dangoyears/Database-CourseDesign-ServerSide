import requests
import json
from public import URL, POST_DATA as token

data = {
    **token, **{
        "studentId": "1706300004c"
    }
}


r = requests.post(URL + '/read/student/one', json=data)
print(json.dumps(r.json(), indent=4, ensure_ascii=False))
