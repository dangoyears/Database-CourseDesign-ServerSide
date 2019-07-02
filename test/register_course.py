import requests
from public import URL, POST_DATA as data

data = {
    **data, **{
        "studentId": "1706300004",
        "courseId": "0000000003"
    }
}

r = requests.post(URL + '/register/course', json=data)
print(r.json())
