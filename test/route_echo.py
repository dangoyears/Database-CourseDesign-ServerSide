import requests
from public import URL, POST_DATA as data

data = {
    **data, **{
        "anyint": 123,
        "anystr": "string"
    }
}

r = requests.post(URL + '/echo?token=TO_BE_REPLACED_BY_TOKEN_IN_POST_DATA', json=data)
print(r.json())
