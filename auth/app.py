import os

from dotenv import load_dotenv
from flask import Flask, jsonify
from flask_cors import CORS
from flask_jwt_extended import JWTManager

from routes.swagger_route import bp as swagger_bp
from routes.user_route import bp as user_bp
from utils.my_bcrypt import bcrypt

load_dotenv()

# init flask app
app = Flask(__name__)
# cors allow all for easy demo
cors = CORS(app, resources={r"/*": {"origins": "*"}})
app.config['JWT_SECRET_KEY'] = os.environ.get("JWT_SECRET_KEY")
app.url_map.strict_slashes = False

# init jwt
jwt = JWTManager(app)
# init bcrypt
bcrypt.init_app(app)

# overide flask error


@app.errorhandler(404)
def page_not_found(e):
    return jsonify(data=None, error=str(e)), 404


@app.errorhandler(500)
def internal_error(e):
    return jsonify(data=None, error=str(e)), 500

# overide jwt error


@jwt.expired_token_loader
def my_expired_token_callback(jwt_header, jwt_payload):
    return jsonify(data=None, error="token has expired"), 401


@jwt.invalid_token_loader
def my_invalid_token_callback(x):
    return jsonify(data=None, error="invalid token"), 422


# register router
app.register_blueprint(user_bp)
app.register_blueprint(swagger_bp)

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=8080, debug=True)
