import time

from scripts.users import create_db, create_user_table
from db.db_conn import db_engine, db_connection_string
from user.user import consume


if __name__ == '__main__':
    print('Consuming messages')
    time.sleep(15)
    create_db(conn_string=db_connection_string)
    create_user_table(db_engine)
    consume()
