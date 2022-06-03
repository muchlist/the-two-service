from db import connection


def get_user(phone: str):

    query = """
    SELECT * FROM user WHERE phone = ?
    """

    conn = connection.get_db_connection()
    user = conn.execute(query, (phone,)).fetchone()
    conn.close()
    return user


def get_users():

    query = """
    SELECT * FROM user ORDER BY id DESC
    """

    conn = connection.get_db_connection()
    cursor = conn.execute(query).fetchall()
    users = [dict(ix) for ix in cursor]
    conn.commit()
    conn.close()

    return users
