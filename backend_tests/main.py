import requests


def health_check(s: requests.Session):
    printline()
    r = s.get("http://localhost:2140/health")
    if r.status_code != 200:
        fail("Health Check Failed")
    else:
        success("Health Check Passed")


def login(s: requests.Session):
    printline()
    r = s.post(
        "http://localhost:2140/api/login",
        json={"email": "ethansteere1@gmail.com", "pass": "password"},
    )
    if r.status_code == 200:
        success("Login Successful")
    else:
        fail(f"Login failed with code {r.status_code}")


def auth_check_after_login(s: requests.Session):
    printline()
    r = s.get("http://localhost:2140/api/auth")
    if r.status_code == 200:
        success("Auth Check Success")
    elif r.status_code == 401:
        fail("Auth is invalid after login")
    else:
        fail(f"Unkown error checking auth: {r.status_code}")


def logout(s: requests.Session):
    printline()
    r = s.post("http://localhost:2140/api/logout")
    if r.status_code == 200:
        success("Logout Successful")
    else:
        fail(f"Unkown error logging out. Recieved Code {r.status_code}")


def auth_check_after_logout(s: requests.Session):
    printline()
    r = s.get("http://localhost:2140/api/auth")
    if r.status_code == 401:
        success("Auth Check Failed as expected after Logout")
    elif r.status_code == 200:
        fail("Auth is still valid after logout")
    else:
        fail(f"Unkown error checking auth: {r.status_code}")


def get_user(s: requests.Session):
    printline()
    r = s.get("http://localhost:2140/api/user")
    if r.status_code == 200:
        success("Fetching User details successful")
    elif r.status_code == 401:
        fail("Unauthorized to retrieve user details")
    else:
        fail(f"Unkown error occured fetching user details. Code: {r.status_code}")
    b = r.json()
    if b["FirstName"] != "Ethan" or b["LastName"] != "Steere":
        fail(f"User body malformed: First: '{b.first_name}' and Last '{b.last_name}'")
    success("Successfully retrieved user body")


def list_chats(s: requests.Session):
    printline()
    r = s.get("http://localhost:2140/api/chat")
    if r.status_code == 200:
        success("Fetching Chats Success")
    else:
        fail(f"Fetching Chats returned status {r.status_code}")
    if r.json() == []:
        success("No Chats as expected")
    else:
        fail("Unexpected chats were found in this request")


def create_chat(s: requests.Session) -> int:
    printline()
    r = s.post(
        "http://localhost:2140/api/chat", json={"Title": "My Fantastic New Chat"}
    )
    if r.status_code == 200:
        success(f"Successfully created new chat with ID: {r.json()['ID']}")
    else:
        fail(f"Error occured creating new chat! Status: {r.status_code}")
    return r.json()["ID"]


def list_chats_after_create(s: requests.Session):
    printline()
    r = s.get("http://localhost:2140/api/chat")
    if r.status_code == 200:
        success("Successfully listed chats")
    else:
        fail(f"Error occured creating new chat! Status: {r.status_code}")

    chats = r.json()
    if chats[0]["Title"] == "My Fantastic New Chat":
        success("Successfully retrieved my new chat")
    else:
        fail("Previously inserted chat was not found!")


def update_chat(s: requests.Session, chatid: int):
    printline()
    r = s.patch(
        "http://localhost:2140/api/chat",
        json={"new_title": "My Great New Chat"},
        params={"chatid": chatid},
    )
    if r.status_code == 200:
        success("Successfully updated chat")
    else:
        fail(f"Error ocurred creating new chat! Status: {r.status_code}")


def list_chats_after_update(s: requests.Session):
    printline()
    r = s.get("http://localhost:2140/api/chat")
    if r.status_code == 200:
        success("Successfully listed chats")
    else:
        fail(f"Error ocurred listing chats! Status: {r.status_code}")
    chats = r.json()
    if chats[0]["Title"] == "My Great New Chat":
        success("Chat title shows updated value")
    else:
        fail("Chat title does not show updated value")


def delete_chat(s: requests.Session, chatid: int):
    printline()
    r = s.delete("http://localhost:2140/api/chat", params={"chatid": chatid})
    if r.status_code == 200:
        success("Succesfully deleted chat")
    else:
        fail(f"Error deleting chat: Status: {r.status_code}")


def list_messages(s: requests.Session, chatid: int):
    printline()
    r = s.get(f"http://localhost:2140/api/chat/{chatid}")
    if r.status_code == 200:
        success("Successfully retrieved messages")
    else:
        fail(f"error occurred listing messages. Status: {r.status_code}")
    if r.json() == []:
        success("Got 0 messages as expected")
    else:
        fail("Unexpected messages were found in this chat!")


def test():
    with requests.Session() as s:
        health_check(s)
        login(s)
        auth_check_after_login(s)
        logout(s)
        auth_check_after_logout(s)
        login(s)
        get_user(s)
        list_chats(s)
        created_chat_id = create_chat(s)
        list_chats_after_create(s)
        update_chat(s, created_chat_id)
        list_chats_after_update(s)
        delete_chat(s, created_chat_id)
        list_chats(s)
        created_chat_id = create_chat(s)
        list_messages(s, created_chat_id)
        delete_chat(s, created_chat_id)

        printline()
        print()


def success(s: str):
    print(f"\033[92m{s}\033[0m")


def fail(s: str):
    print(f"\033[91m{s}\033[0m")
    exit(1)


def printline():
    print("\u2500" * 40)


if __name__ == "__main__":
    test()
