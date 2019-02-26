module brpaste.web;

import brpaste.hash;

import vibe.vibe;

RedisDatabase client;
shared static this() {
    string path = "redis://127.0.0.1";
    readOption("redis|r", &path, "The URL to use to connect to redis");
    URL redis = path;
    client = connectRedisDB(redis);
}

class BRPaste {
    @method(HTTPMethod.REPORT)
    @path("/health")
    void health(HTTPServerResponse res) {
        import std.array;
        res.statusCode = 200;
        auto app = appender!string;

        // is Redis healthy? - FIXME: stack traces over HTTP are fun, I guess
        import std.random;
        long val = uniform!uint;
        auto ech = client.client.echo!(long, long)(val);
        if(val != ech) {
            res.statusCode = 500;
            app.put("Redis: failed.");
        } else app.put("Redis: pass.");

        res.writeBody(app.data);
    }

    void index() {
        render!("index.dt");
    }

    @path("/:id")
    void getId(string _id) {
        if (!client.exists(_id)) throw new HTTPStatusException(404);
        string language = "none";
        // TODO: rewrite the next two lines once #2273 is resolved
        auto req = request;
        if (req.query.length > 0) language = req.query.byKey.front;
        auto data = client.get(_id);
        render!("code.dt", data, language);
    }

    @path("/raw/:id")
    void getRawId(HTTPServerResponse res, string _id) {
        if (!client.exists(_id)) throw new HTTPStatusException(404);
        auto val = client.get(_id);
        res.contentType = "text/plain";
        res.writeBody(val);
    }

    void post(string data) {
        auto hash = data.hash;
        client.set(hash, data);
        status(201);
        response.writeBody(hash);
    }
}
