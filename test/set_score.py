import requests
from public import URL, POST_DATA as data

data = {
    **data, **{
        "courseId": "0000000001",
        "students": {
            "1706300004": "23"
        }
    }
}

r = requests.post(URL + '/set/score', json=data)
print(r.json())
