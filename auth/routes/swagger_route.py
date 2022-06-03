from flask import Blueprint, render_template


bp = Blueprint('swagger_bp', __name__)

@bp.route('/docs', methods=['GET'])
def get_docs():
    print('sending docs')
    return render_template('swaggerui.html')