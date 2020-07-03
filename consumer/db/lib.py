import os

import sqlalchemy as db
from sqlalchemy.orm import sessionmaker


def get_db_connection():
    """
    Returns db engine and connection string
    """
    user = os.environ.get('POSTGRES_USER')
    password = os.environ.get('POSTGRES_PASSWORD')
    host = os.environ.get('POSTGRES_HOST')
    postgres_db = os.environ.get('POSTGRES_DB')
    db_connection_string = f'postgres+psycopg2://{user}:{password}@{host}/{postgres_db}'
    db_engine = db.create_engine(db_connection_string)
    return db_engine, db_connection_string


def get_db_session(db_engine):
    """
    Returns db session
    """
    Session = sessionmaker(bind=db_engine)
    return Session()
