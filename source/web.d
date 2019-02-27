module brpaste.web;

import brpaste.hash;

import vibe.vibe;

RedisDatabase client;

void id(HTTPServerRequest req, HTTPServerResponse res) {
    string id = req.params["id"];
    string language = "none";
    // TODO: rewrite the next two lines once #2273 is resolved
    if (req.query.length > 0) language = req.query.byKey.front;
    enforceHTTP(client.exists(id), HTTPStatus.notFound, "No paste under " ~ id ~ ".");

    auto data = client.get(id);
    render!("code.dt", data, language)(res);
}

void rawId(HTTPServerRequest req, HTTPServerResponse res) {
    string id = req.params["id"];
    enforceHTTP(client.exists(id), HTTPStatus.notFound, "No paste under " ~ id  ~ ".");

    auto data = client.get(id);
    res.contentType = "text/plain";
    res.writeBody(data);
}

void post(HTTPServerRequest req, HTTPServerResponse res) {
    enforceHTTP("data" in req.form, HTTPStatus.badRequest, "Missing data field.");
    auto data = req.form["data"];

    auto hash = data.hash;
    client.set(hash, data);
    res.statusCode = HTTPStatus.created;
    res.writeBody(hash);
}

void health(HTTPServerRequest req, HTTPServerResponse res) {
    res.statusCode = HTTPStatus.noContent;
    scope(exit) res.writeBody("");

    // Redis
    try {
        client.client.ping;
    } catch (Exception e) {
        logCritical("Redis is down!");
        res.statusCode = HTTPStatus.serviceUnavailable;
        res.statusPhrase = "Backend Storage Unavailable";
        res.headers["Retry-After"] = "60";
    }
}

shared static this() {
    // setup redis
    string path = "redis://127.0.0.1";
    readOption("redis|r", &path, "The URL to use to connect to redis");
    URL redis = path;
    client = connectRedisDB(redis);
}

