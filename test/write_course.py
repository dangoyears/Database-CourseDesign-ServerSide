import requests
from public import URL, POST_DATA as data

data = {
    **data, **{
        "name": "数据结构",
        "id": "0000000001",
        "credit": "2",
        "nature": "专业必修课",
        "accommodate": "50",
        "selectedSum": "50",
        "time": "第7-14周,第4-6节",
        "teachers": "['xxx', 'yyy']",
        "courseLeader": "yyy",
        "address": "理科南教学楼710",
        "class": "['计算机科学与网络学院-软件工程-171']"
    }
}

r = requests.post(URL + '/write/course', json=data)
print(r.json())
