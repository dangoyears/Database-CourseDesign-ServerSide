import requests
from public import URL, POST_DATA as data

data = {
    **data, **{
        "anything": "something"
    }
}

r = requests.get(URL + '/admin', params=data)
print(r.json())

r = requests.post(URL + '/logout', json=data)
print(r.json())

r = requests.get(URL + '/admin', params=data)
print(r.json())
