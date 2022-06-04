from datetime import timedelta

from dao import user_query, user_update
from dto.result_dto import Result
from dto.user_dto import User
from flask_jwt_extended import create_access_token
from utils import password_gen
from utils.my_bcrypt import bcrypt

EXPIRED_TOKEN = 15  # in minute


def user_exist(phone) -> bool:
    if user_query.get_user(phone):
        return True
    return False


def login(phone: str, password: str) -> Result:
    # get user data from repo
    try:
        user = user_query.get_user(phone)
    except Exception as err:
        return Result(data=None, error=str(err), code=500)

    # check if password is same (we not using hashing method in this simple apps)
    if not (user and bcrypt.check_password_hash(user["password"], password)):
        return Result(data=None, error="user or password not valid", code=400)

    # create access token
    access_token = create_access_token(
        identity=user["phone"],
        expires_delta=timedelta(minutes=EXPIRED_TOKEN),
        fresh=True,
        additional_claims={
            'name': user['name'],
            'phone': user['phone'],
            'role': user['role'],
            'timestamp': user['timestamp']
        }
    )

    response = Result(data={
        'access_token': access_token,
        'name': user['name'],
        'phone': user['phone'],
        'role': user['role'],
        'timestamp': user['timestamp'],
    }, error=None, code=200)

    return response


def register(data: User) -> Result:
    # generate password
    pw_generated = password_gen.generate_password()

    # hash password
    pw_hash = bcrypt.generate_password_hash(pw_generated).decode("utf-8")
    data.password = pw_hash

    # check is user exist
    if user_exist(data.phone):
        return Result(data=None, error="phone not available", code=400)

    try:
        user_update.insert_user(data)
    except Exception as err:
        return Result(data=None, error="failed store data to database " + str(err), code=500)

    return Result(
        data={"message": "success register phone number",
              "phone": data.phone,
              "password": pw_generated
              },
        error=None,
        code=201
    )


def get_all() -> Result:
    try:
        users = user_query.get_users()
    except Exception as err:
        return Result(data=None, error="failed get data from database " + str(err), code=500)

    return Result(
        data=users,
        error=None,
        code=201
    )
