import re

from marshmallow import Schema, ValidationError, fields


def validate_phone(n):
    msg = u"invalid phone number."

    try:
        int(n)
    except ValueError as err:
        raise ValidationError(msg)

    rule = re.compile(
        r'\d{3}[-\.\s]??\d{3}[-\.\s]??\d{4}|\(\d{3}\)\s*\d{3}[-\.\s]??\d{4}|\d{3}[-\.\s]??\d{4}')
    if not rule.search(n):
        raise ValidationError(msg)


class UserLoginSchema(Schema):
    phone = fields.Str(required=True)
    password = fields.Str(required=True)


class UserRegisterSchema(Schema):
    phone = fields.Str(required=True, validate=validate_phone)
    name = fields.Str(required=True)
    role = fields.Str(required=True)
