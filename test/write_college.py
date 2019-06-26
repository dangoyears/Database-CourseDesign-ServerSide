import requests
from public import URL, POST_DATA as data

data = {
    **data, **{
        "college": "计算机科学与网络学院",
        "specialty": "软件工程",
        "grade": "17",
        "class": "1",
    }
}

r = requests.post(URL + '/write/college', json=data)
print(r.json())
