import redis

class Subscriber:
    def __init__(self, id, channel) -> None:
        self.redisClient = redis.Redis(
            host="127.0.0.1", 
            port=6379,
            password="",
            db=0
        )
        self.id = id
        self.channel = channel

    def listen(self):
        pubsub = self.redisClient.pubsub()
        pubsub.subscribe(self.channel)
        # print(f"[subscriber-{self.id}]: start listening...")
        for message in pubsub.listen():
            if message["type"] == 'message':
                # print(f"[subscriber-{self.id}]: message= {message['data'].decode('utf-8')}")
                pass