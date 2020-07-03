from sqlalchemy import Column, String
from sqlalchemy.ext.declarative import declarative_base


Base = declarative_base()


class UserData(Base):
    """
    Defines UserData model
    """
    __tablename__ = 'users'

    name = Column(String)
    email = Column(String, primary_key=True)

    def __repr__(self):
        return f'user details: name: {self.name}, email: {self.email}'
