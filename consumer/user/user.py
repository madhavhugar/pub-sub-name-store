import json
import pika
import os

from sqlalchemy import exc

from models.users import UserData
from db.db_conn import db_session


def callback(ch, method, properties, body, db):
    """
    Called whenever message received in queue
    """
    payload = body.decode("utf-8")
    user_data = json.loads(payload)
    email = user_data['email']

    user = UserData(name=user_data['name'],
                    email=user_data['email'])
    try:
        db.add(user)
        db.commit()
        print(f' [x] Added {user}')
    except exc.IntegrityError:
        print(f'{email} already exists')
        db.rollback()


def consume():
    """
    Consume data and run callback whenever necessary
    """
    rabbitmq_host = os.environ.get('RABBITMQ_HOST')
    connection_mq = pika.BlockingConnection(
                 pika.ConnectionParameters(host=rabbitmq_host))
    channel = connection_mq.channel()
    channel.basic_consume(
        queue='users-info-queue',
        on_message_callback=(lambda ch, method, properties, body: callback(
            ch,
            method,
            properties,
            body,
            db_session,
            )),
        auto_ack=True)
    channel.start_consuming()
