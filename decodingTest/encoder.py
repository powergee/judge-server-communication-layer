import struct
import zlib
import json

to_send = {
    "a" : "b",
    "c" : "d"
}

size_pack = struct.Struct("!I")
raw = zlib.compress(json.dumps(to_send).encode('utf-8'))
writer = open("encoded", "wb")
print(len(raw))
writer.write(size_pack.pack(len(raw)))
writer.write(raw)
writer.close()