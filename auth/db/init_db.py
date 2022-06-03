import os
import sqlite3


def init_database():
    connection = sqlite3.connect(os.path.join('db_data', 'database.db'))
    with open(os.path.join('db', 'schema.sql')) as f:
        connection.executescript(f.read())
    connection.commit()
    connection.close()


if __name__ == '__main__':
    init_database()
