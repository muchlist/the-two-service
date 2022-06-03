from dto.user_dto import User
from flask import Blueprint, jsonify, request
from flask_jwt_extended import get_jwt, jwt_required
from marshmallow import ValidationError
from service import user_service
from validate.user_schema import UserLoginSchema, UserRegisterSchema

bp = Blueprint('user_bp', __name__)


"""
------------------------------------------------------------------------------
login
------------------------------------------------------------------------------
"""


@bp.route('/login', methods=['POST'])
def login_handler():
    schema = UserLoginSchema()
    try:
        data = schema.load(request.get_json())
    except ValidationError as err:
        return jsonify(data=None, error=str(err.messages)), 400

    result = user_service.login(data['phone'], data['password'])
    return jsonify(data=result.data, error=result.error), result.code


"""
------------------------------------------------------------------------------
register
------------------------------------------------------------------------------
"""


@bp.route('/register', methods=['POST'])
def register_user():

    schema = UserRegisterSchema()
    try:
        data = schema.load(request.get_json())
    except ValidationError as err:
        return jsonify(data=None, error=str(err.messages)), 400

    result = user_service.register(
        User(phone=data['phone'],
             name=data['name'],
             role=data['role'],
             password='',))

    return jsonify(data=result.data, error=result.error), result.code


"""
------------------------------------------------------------------------------
profil
------------------------------------------------------------------------------
"""


@bp.route('/profil', methods=['GET'])
@jwt_required()
def profil():
    return jsonify(data=get_jwt(), error=None), 200


"""
------------------------------------------------------------------------------
list
------------------------------------------------------------------------------
"""


@bp.route('/users', methods=['GET'])
@jwt_required()
def find_user():
    claims = get_jwt()
    if claims['role'].lower() != 'admin':
        return jsonify(data=None, error='need user with role admin to access'), 400

    result = user_service.get_all()
    return jsonify(data=result.data, error=result.error), result.code
