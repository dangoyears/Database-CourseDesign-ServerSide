# 系统环境设置与管理员用户登陆

import requests

# 开发环境与生产环境地址
BACKEND = {
    "DEV": "http://localhost:12323",
    "PRO": "https://dbcd.qfstudio.net"
}

# 后端接口地址
URL = BACKEND["DEV"]

# 登陆表单信息
LOGIN_DATA = {
    "user": "dangoyears",
    "pass": "dangoyears",
    "type": "admin"
}

# 尝试登陆
r = requests.post(URL + '/login', json=LOGIN_DATA)
print(r.json())

# 导出用户授权token
POST_DATA = {
    "token": r.json().get("token")
}
