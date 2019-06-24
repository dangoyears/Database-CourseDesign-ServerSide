import requests

# 开发环境与生产环境地址
BACKEND = {
    "DEV": "http://localhost:12323",
    "PRO": "https://dbcd.qfstudio.net"
}

# 后端接口地址
URL = BACKEND["DEV"]

r = requests.post(URL + '/login', json={
    "user": "dangoyears",
    "pass": "dangoyears",
    "type": "admin"
})
print(r.json())

r = requests.post(URL + '/login', json={
    "user": "1706300000",
    "pass": "whoami",
    "type": "student"
})
print(r.json())

r = requests.post(URL + '/login', json={
    "user": "0000000001",
    "pass": "310612",
    "type": "teacher"
})
print(r.json())
