from db import connection
from dto.user_dto import User


def insert_user(user: User):

    query = """
    INSERT INTO user (phone, name, role, password) VALUES (?,?,?,?)
    """

    print(user.name)
    print(user.password)
    conn = connection.get_db_connection()
    conn.execute(query, (user.phone, user.name, user.role, user.password))
    conn.commit()
    conn.close()
