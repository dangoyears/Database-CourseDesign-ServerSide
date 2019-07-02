import requests
from public import URL, POST_DATA as data

data = {
    **data, **{
        "id": "0000000001"
    }
}

r = requests.post(URL + '/delete/course', json=data)
print(r.json())
