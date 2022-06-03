from db import connection


def get_user(phone: str):

    query = """
    SELECT * FROM user WHERE phone = ?
    """

    conn = connection.get_db_connection()
    user = conn.execute(query, (phone,)).fetchone()
    conn.close()
    return user
