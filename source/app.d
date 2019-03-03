import brpaste.web;

import vibe.d;

shared static this() {
    // HTTP settings
    auto settings = new HTTPServerSettings;
    settings.port = 8080;
    settings.bindAddresses = [];

    readOption("bind|b", &settings.bindAddresses, "Sets the addresses to bind to [127.0.0.1 ::1]");
    readOption("port|p", &settings.port, "Sets the port to listen on [8080]");
    if(settings.bindAddresses.empty) settings.bindAddresses = [ "127.0.0.1", "::1" ];

    // setup router
    auto router = new URLRouter;

    router.match(HTTPMethod.REPORT, "/health", &health);
    router.get("/health", &health);

    router.get("/", staticTemplate!"index.dt");

    router.post("/", &post);
    router.put("/:id", &put);

    router.get("/:id", &id);
    router.get("/raw/:id", &rawId);

    listenHTTP(settings, router);
}
