module brpaste.web;

import brpaste.hash;
import brpaste.storage;

import vibe.vibe;

import std.functional;

RedisStorage store;

alias put  = partial!(insert, true);
alias post = partial!(insert, false);

string idCommon(in HTTPServerRequest req) {
    string id = req.params["id"];
    return store.get(id);
}

void id(HTTPServerRequest req, HTTPServerResponse res) {
    string language = "none";
    // TODO: rewrite the next two lines once #2273 is resolved
    if ("lang" in req.query) language = req.query["lang"];
    else if (req.query.length > 0) language = req.query.byKey.front;

    auto data = idCommon(req);
    render!("code.dt", data, language)(res);
}

void rawId(HTTPServerRequest req, HTTPServerResponse res) {
    res.contentType = "text/plain";

    auto data = idCommon(req);
    res.writeBody(data);
}

void insert(bool put, HTTPServerRequest req, HTTPServerResponse res) {
    enforceHTTP("data" in req.form, HTTPStatus.badRequest, "Missing data field.");
    auto data = req.form["data"];

    auto hash = put ? req.params["id"] : data.hash;
    store.put(hash, data, put);
    res.statusCode = HTTPStatus.created;
    res.writeBody(hash);
}

void health(HTTPServerRequest req, HTTPServerResponse res) {
    res.statusCode = HTTPStatus.noContent;
    scope(success) res.writeBody("");

    // Redis
    store.isDown;
}

shared static this() {
    // setup redis
    string path;
    readOption("redis|r", &path, "The URL to use to connect to redis");
    store = path.empty ? new RedisStorage : new RedisStorage(URL(path));
}

