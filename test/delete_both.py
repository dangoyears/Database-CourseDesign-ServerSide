import requests
from public import URL, POST_DATA as token

r = requests.post(URL + '/delete/both', json={
    **token, **{
        "type": "student",
        "id": "1706300000"
    }
})
print(r.json())


r = requests.post(URL + '/delete/both', json={
    **token, **{
        "type": "teacher",
        "id": "0000000001"
    }
})
print(r.json())

r = requests.post(URL + '/delete/both', json={
    **token, **{
        "type": "unkown",
        "id": "123456"
    }
})
print(r.json())
