from user.user import callback
from models.users import UserData


def test_consumer_data_to_db(postgres):
    url, engine, db_session = postgres
    body = '{"name": "session", "email": "session@test.com"}'.encode('utf-8')
    callback(None, None, None, body, db_session)
    user_data = db_session.query(UserData).first()
    assert user_data.name == "session"
    assert user_data.email == "session@test.com"
    db_session.query(UserData).delete()


def test_duplicate_consumer_data_to_db(postgres):
    url, engine, db_session = postgres
    body = '{"name": "session", "email": "session@test.com"}'.encode('utf-8')
    callback(None, None, None, body, db_session)
    body = '{"name": "different", "email": "session@test.com"}'.encode('utf-8')
    callback(None, None, None, body, db_session)
    user_data = db_session.query(UserData)\
        .filter(UserData.email == 'session@test.com')\
        .all()
    assert len(user_data) == 1
    assert user_data[0].name == 'session'
    assert user_data[0].email == 'session@test.com'
    db_session.query(UserData).delete()
