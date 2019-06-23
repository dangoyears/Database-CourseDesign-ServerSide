import requests
from public import URL, POST_DATA as data

data = {
    **data, **{
        "college": "路由测试学院",
        "specialty": "路由测试专业",
        "grade": "17",
        "class": "1",
        "name": "测试生",
        "studentId": "1706300000",
        "status": "在读本科生",
        "sex": "女",
        "birthday": "1970-3-23",
        "age": "21",
        "idCard": "440582199708310612",
        "password": "whoami",
        "yearSystem": "4"
    }
}

r = requests.post(URL + '/write/student', json=data)
print(r.json())
