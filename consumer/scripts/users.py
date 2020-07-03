from sqlalchemy_utils import create_database, database_exists

from models.users import UserData


def create_db(conn_string):
    """
    Creates db if it doesn't exist
    """
    if not database_exists(conn_string):
        create_database(conn_string)


def create_user_table(engine):
    """
    Creates user table
    """
    user_table = UserData.__tablename__
    if not engine.dialect.has_table(engine, user_table):
        UserData.__table__.create(bind=engine, checkfirst=True)
