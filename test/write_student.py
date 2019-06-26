import requests
from public import URL, POST_DATA as token

data = {
    **token, **{
        "college": "计算机科学与网络学院",
        "specialty": "软件工程",
        "grade": "17",
        "class": "1",
        "name": "学生1",
        "studentId": "1706300000",
        "status": "在读本科生",
        "sex": "女",
        "birthday": "1970-3-23",
        "age": "21",
        "idCard": "440582199708310610",
        "password": "whoami",
        "yearSystem": "4"
    }
}

r = requests.post(URL + '/write/student', json=data)
print(r.json())

data = {
    **token, **{
        "college": "计算机科学与网络学院",
        "specialty": "软件工程",
        "grade": "17",
        "class": "1",
        "name": "学生2",
        "studentId": "1706300004",
        "status": "在读本科生",
        "sex": "女",
        "birthday": "1970-3-23",
        "age": "21",
        "idCard": "440582199708310688",
        "password": "whoami",
        "yearSystem": "4"
    }
}

r = requests.post(URL + '/write/student', json=data)
print(r.json())
