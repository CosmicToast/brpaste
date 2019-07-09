module brpaste.web;

import brpaste.hash;
import brpaste.storage;

import vibe.vibe;

import std.functional;
import std.regex;

RedisStorage store;

alias put   = partial!(insert, true);
alias post  = partial!(insert, false);
alias idLng = partial!(id, true);
alias idRaw = partial!(id, false);

void id(bool highlight, HTTPServerRequest req, HTTPServerResponse res) {
    string id = req.params["id"];
    auto data = store.get(id);

    if(!highlight) {
        res.contentType = "text/plain";
        res.writeBody(data);
        return;
    }

    string language = "none";
    // TODO: rewrite the next two lines once #2273 is resolved
    if ("lang" in req.query) language = req.query["lang"];
    else if (req.query.length > 0) language = req.query.byKey.front;

    render!("code.dt", data, language)(res);
}

void insert(bool put, HTTPServerRequest req, HTTPServerResponse res) {
    import std.encoding;

    enforceHTTP("data" in req.form, HTTPStatus.badRequest, "Missing data field.");
    string data = req.form["data"];
    enforceHTTP(data.isValid, HTTPStatus.unsupportedMediaType, "Content contains binary.");
    auto hash = put ? req.params["id"] : data.hash;
    store.put(hash, data, put);

    auto ua = req.headers.get("User-Agent", "");
    if(ua.isBrowser) {
        // TODO: eventually move back to registerWebInterface for redirect()
        res.statusCode = HTTPStatus.seeOther;
        res.headers["Location"] = "/%s".format(hash);
        res.writeBody("");
    } else {
        res.statusCode = HTTPStatus.created;
        res.writeBody(hash);
    }
}

void health(HTTPServerRequest req, HTTPServerResponse res) {
    res.statusCode = HTTPStatus.noContent;
    scope(success) res.writeBody("");

    // Redis
    store.isDown;
}

// tries to match User-Agent against known browsers
static bool isBrowser(string ua) {
    foreach (r; [
            "Firefox/",
            "Chrome/",
            "Safari/",
            "OPR/",
            "Edge/",
            "Trident/"
        ]) {
        if(ua.matchFirst(r)) return true;
    }
    return false;
}

shared static this() {
    // setup redis
    string path;
    readOption("redis|r", &path, "The URL to use to connect to redis");
    store = path.empty ? new RedisStorage : new RedisStorage(URL(path));
}

