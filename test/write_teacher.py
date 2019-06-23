import requests
from public import URL, POST_DATA as data

data = {
    **data, **{
        "college": "路由测试学院",
        "name": "南宫问天",
        "jobId": "0000000001",
        "sex": "男",
        "education": "硕士",
        "graduation": "南开大学",
        "birthday": "2000-1-1",
        "age": "34",
        "idCard": "440582199708310612",
        "password": "310612",
        "position": "教务办主任"
    }
}

r = requests.post(URL + '/write/teacher', json=data)
print(r.json())
