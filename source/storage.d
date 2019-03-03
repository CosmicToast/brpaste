module brpaste.storage;

import vibe.vibe;

class RedisStorage {
    private RedisDatabase client;

    this(URL url = URL("redis://127.0.0.1")) {
        client = connectRedisDB(url);
    }

    void isDown() {
        enforceHTTP(healthy, HTTPStatus.serviceUnavailable, "Redis is down.");
    }

    bool healthy() {
        try {
            client.client.ping;
        } catch (Exception e) {
            return false;
        }
        return true;
    }

    auto get(in string key) {
        isDown;
        enforceHTTP(client.exists(key), HTTPStatus.notFound, key ~ " not found.");
        return client.get(key);
    }

    void put(in string key, in string data, in bool collision = false) {
        isDown;
        if(collision) enforceHTTP(! client.exists(key), HTTPStatus.unprocessableEntity, key ~ " already exists.");
        client.set(key, data);
    }
}
