import pytest
from sqlalchemy import create_engine
import testing.postgresql

from db.lib import get_db_session
from scripts.users import create_db, create_user_table
from sqlalchemy_utils import database_exists, drop_database


def remove_db(engine):
    if database_exists(engine.url):
        drop_database(engine.url)

@pytest.fixture(scope="session")
def postgres():
    """
    Postgres fixture - starts a postgres instance inside a temp directory
    and closes it after tests are done
    """
    with testing.postgresql.Postgresql() as postgresql:
        conn_str = postgresql.url()
        engine = create_engine(conn_str)
        db_session = get_db_session(engine)
        create_db(conn_str)
        create_user_table(engine)
        yield conn_str, engine, db_session
        db_session.rollback()
        remove_db(engine)
        engine.dispose()
