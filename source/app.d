import brpaste;

import vibe.d;

shared static this() {
    auto settings = new HTTPServerSettings;
    settings.port = 8080;
    settings.bindAddresses = [];

    readOption("bind|b", &settings.bindAddresses, "Sets the addresses to bind to [127.0.0.1 ::1]");
    readOption("port|p", &settings.port, "Sets the port to listen on [8080]");
    if(settings.bindAddresses.empty) settings.bindAddresses = [ "127.0.0.1", "::1" ];

    auto router = new URLRouter;
    router.registerWebInterface(new BRPaste);
    listenHTTP(settings, router);
}
