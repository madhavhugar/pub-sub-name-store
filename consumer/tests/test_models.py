from models.users import UserData


def test_users_data(postgres):
    url, engine, db_session = postgres
    user = UserData(name='kukido',
                    email='kukido@ki.com')
    db_session.add(user)
    db_session.commit()
    user_data = db_session.query(UserData).first()
    assert user_data.name == 'kukido'
    assert user_data.email == 'kukido@ki.com'
    db_session.query(UserData).delete()
