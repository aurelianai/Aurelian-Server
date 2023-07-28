import requests

def health_check(s: requests.Session):
   r = s.get("http://localhost:2140/health")
   if r.status_code != 200:
      print("Health Check Failed")
      exit(1) 
   else:
      print("Health Check Passed")

def login(s: requests.Session):
   r = s.post(
      'http://localhost:2140/api/login', 
      json={"email":"ethansteere1@gmail.com", "pass": "password"}
   )
   if r.status_code == 200:
      print("Login Successful")
   else:
      print(f"Login failed with code {r.status_code}")

def auth_check_after_login(s: requests.Session):
   r = s.get(
      'http://localhost:2140/api/auth',
   )
   if r.status_code == 200:
      print("Auth Check After Login Success")
   elif r.status_code == 401:
      print("Auth is invalid after login")
      exit(1)
   elif r.status_code == 500:
      print("Error ocurred after auth")
      exit(1)
   else:
      print(f"Unkown error checking auth: {r.status_code}")
      exit(1)

def logout(s: requests.Session):
   r = s.post(
      'http://localhost:2140/api/logout'
   )
   if r.status_code == 200:
      print("Logout Successful")
   else:
      print(f'Unkown error logging out. Revieved Code {r.status_code}')

def auth_check_after_logout(s: requests.Session):
   r = s.get(
      'http://localhost:2140/api/auth',
   )
   if r.status_code == 401:
      print("Auth Check Failed as expected after Logout")
   elif r.status_code == 200:
      print("Auth is still valid after logout")
      exit(1)
   elif r.status_code == 500:
      print("Error ocurred checking auth after logout")
      exit(1)
   else:
      print(f"Unkown error checking auth: {r.status_code}")
      exit(1)

def get_user(s: requests.Session):
   r = s.get(
      'http://localhost:2140/api/user'
   )
   r_body = r.text
   print(r_body)
   return
   if r_body.first_name != "Ethan" or r_body.last_name != "Steere":
      print(f"User body malformed: First: '{r_body.first_name}' and Last '{r_body.last_name}'")
      exit(1)
   if r.status_code == 200:
      print("Fetching User details successful")
   elif r.status_code == 401:
      print("Auth error getting user details")
      exit(1)
   else:
      print(f"Unkown error occured fetching user details. Code: {r.status_code}")
      exit(1)


def test():
   s = requests.Session()

   # Passing
   health_check(s)
   login(s)
   '''
   auth_check_after_login(s)
   logout(s)
   auth_check_after_logout(s)
   login(s)
   get_user(s)
   '''
if __name__ == "__main__":
   test()