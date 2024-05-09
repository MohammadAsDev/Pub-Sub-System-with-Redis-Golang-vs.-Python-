import subscriber
import publisher
from multiprocessing.pool import ThreadPool
import time
import cProfile
import yaml

TESTING_CHANNEL = "testing_python3_channel"
PUB_MAX_WORKERS = 5
SUB_MAX_WORKERS = 5

def generate_publishers(n_publishers : int) -> list:
    publishers = []
    for i in range(n_publishers):
        publishers.append(publisher.Publisher(i , TESTING_CHANNEL))
    return publishers

def generate_subscribers(n_subscribers : int) -> list:
    subscribers = []
    for i in range(n_subscribers):
        subscribers.append(subscriber.Subscriber(i , TESTING_CHANNEL))
    return subscribers

def load_config(config_path : str) -> dict:
    with open(config_path, 'r') as file:
        data = yaml.safe_load(file)

    return data

def main():

    config = load_config("../config.yaml")

    n_pubs  = config["n_pubs"]
    n_subs = config["n_subs"]
    n_messages = config["n_messages"]
    pubs = generate_publishers(n_pubs)
    subs = generate_subscribers(n_subs)

    pub_pool = ThreadPool(processes=n_pubs)
    sub_pool = ThreadPool(processes=n_subs)

    for pub in pubs:
        for i in range(n_messages):
            pub_pool.apply_async(
                pub.publish , 
                kwds={"message" : f"[publisher-{pub.id}]: this is a message-{i + 1}"}
            )
            
    
    for sub in subs:
        sub_pool.apply_async(sub.listen)


    time.sleep(60)  # running the script for a random amount of time

if __name__ == "__main__":
    """
        Note: many print oeprations have been ignored
        if you want to know how this system work, 
        just uncomment print operations
        in 'publisher.py' and 'subscriber.py'
    """
    # cProfile.run("main()")    # using cProfile
    main()