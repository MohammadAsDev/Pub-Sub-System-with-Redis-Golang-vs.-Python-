import redis

class Publisher:
    def __init__(self, id, channel) -> None:
        self.id = id
        self.channel = channel
        self.redisClient = redis.Redis(
            host="127.0.0.1", 
            port=6379,
            password="",
            db=0
        )

    def publish(self, message):
        # print(f"[publisher-{self.id}]: publishing {message}")
        self.redisClient.publish(self.channel, message)