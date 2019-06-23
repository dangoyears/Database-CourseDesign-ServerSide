import requests
from public import URL, POST_DATA as data

data = {
    **data, **{
        "college": "路由测试学院",
        "specialty": "路由测试专业",
        "grade": "17",
        "class": "1",
    }
}

r = requests.post(URL + '/delete/class', json=data)
print(r.json())
