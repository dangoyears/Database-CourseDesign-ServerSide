import requests
from public import URL, POST_DATA as token

data = {
    **token, **{
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

data = {
    **token, **{
        "college": "路由测试学院",
        "name": "xxx",
        "jobId": "0000000002",
        "sex": "男",
        "education": "硕士",
        "graduation": "南开大学",
        "birthday": "2000-1-1",
        "age": "34",
        "idCard": "440582199708310608",
        "password": "310612",
        "position": "教务办主任"
    }
}

r = requests.post(URL + '/write/teacher', json=data)
print(r.json())

data = {
    **token, **{
        "college": "路由测试学院",
        "name": "yyy",
        "jobId": "0000000003",
        "sex": "男",
        "education": "硕士",
        "graduation": "南开大学",
        "birthday": "2018-3-14",
        "age": "34",
        "idCard": "440582199708310609",
        "password": "310612",
        "position": "教务办主任"
    }
}

r = requests.post(URL + '/write/teacher', json=data)
print(r.json())
