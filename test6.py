import requests
import getpass
import json
import urllib3

urllib3.disable_warnings(urllib3.exceptions.InsecureRequestWarning)
username = input("Username: ")
password = getpass.getpass()

print("-----------------")

token_data=json.dumps({"username": username, "password": password})

token = requests.post("https://securitycenter/rest/token", verify='\\pit-darizola\CA\Certificates-Current User\Personal\Certificates')
#print(token.status_code)
#print(token.json())
#print(token.cookies)
#print(token_num)
token_num = token.json()['response']['token']

headers={"X-SecurityCenter": str(token_num)}
assets = requests.get("https://securitycenter/rest/asset?id=209&fields=id,name,viewableIPs", headers=headers,
                      verify=False, cookies=token.cookies)

print(assets.json())
