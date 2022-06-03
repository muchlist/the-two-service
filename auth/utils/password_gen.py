import random
import string

# characters to generate password from
characters = list(string.ascii_letters + string.digits + "!@#$%^&*()")
length = 4


def generate_password() -> str:
    # shuffling the characters
    random.shuffle(characters)

    # picking random characters from the list
    password = []
    for i in range(length):
        password.append(random.choice(characters))

    # shuffling the resultant password
    random.shuffle(password)

    # converting the list to string
    # printing the list
    return "".join(password)
