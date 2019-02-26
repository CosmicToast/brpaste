import brpaste;

import vibe.d;

int main(string[] args)
{
    auto settings = new HTTPServerSettings;
    settings.port = 8080;
    settings.bindAddresses = [];

    readOption("bind|b", &settings.bindAddresses, "Sets the address to bind to");
    readOption("port|p", &settings.port, "Sets the port to listen on");
    if(settings.bindAddresses.length == 0) settings.bindAddresses = [ "127.0.0.1" ];

    auto router = new URLRouter;
    router.registerWebInterface(new BRPaste);
    listenHTTP(settings, router);

    import std.conv;
    logInfo("Please open http://%s:%d/ in your browser.", settings.bindAddresses[0], settings.port);
    return runApplication();
}
