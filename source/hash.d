module brpaste.hash;

pure string hash(T)(T data) {
    import std.base64;
    import std.digest.murmurhash;
    auto hash = digest!(MurmurHash3!32)(data);
    return Base64URLNoPadding.encode(hash);
}

pure string hash(T : string)(T data) {
    import std.string;
    return hash(data.representation);
}
