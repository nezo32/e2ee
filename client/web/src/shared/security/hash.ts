export function FNV1a(data: Uint8Array<ArrayBuffer>) {
    const prime = 0x01000193;
    let hash = 0x811c9dc5;

    for (let i = 0; i < data.length; i++) {
        hash ^= data[i];
        hash = (hash * prime) >>> 0;
    }

    return hash >>> 0;
}