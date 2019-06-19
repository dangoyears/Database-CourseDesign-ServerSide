import requests

BACKEND = {
    "DEV": "http://localhost:12323",
    "PRO": "https://dbcd.qfstudio.net"
}

URL = BACKEND["DEV"]

LOGIN_DATA = {
    "user": "dangoyears",
    "pass": "dangoyears",
    "type": "admin"
}

r = requests.post(URL + '/login', json=LOGIN_DATA)
print(r.json())
