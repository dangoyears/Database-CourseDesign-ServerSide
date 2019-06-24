import requests
from public import URL, POST_DATA as token

r = requests.post(URL + '/read/college', json=token)
print(r.json())
