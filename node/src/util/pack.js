/**
 * @author  zhaoliang.liang
 * @date  2024/12/30 12:28
 */

const { Buffer } = require('buffer');

// Utility functions to pack data into the byte stream

function packUint16(w, n) {
    const buf = Buffer.alloc(2);
    buf.writeUInt16LE(n, 0);
    w.push(buf);
}

function packUint32(w, n) {
    const buf = Buffer.alloc(4);
    buf.writeUInt32LE(n, 0);
    w.push(buf);
}

function packUint64(w, n) {
    const buf = Buffer.alloc(8);
    buf.writeBigUInt64LE(BigInt(n), 0);
    w.push(buf);
}

function packString(w, s) {
    const len = Buffer.alloc(2);
    len.writeUInt16LE(s.length, 0);
    w.push(len);
    w.push(Buffer.from(s, 'utf-8'));
}

function packHexString(w, s) {
    const buf = Buffer.from(s, 'hex');
    return packString(w, buf.toString('utf8'));
}

function packExtra(w, extra) {
    const keys = Object.keys(extra).map(k => parseInt(k));
    keys.sort((a, b) => a - b); // sort keys in ascending order

    packUint16(w, keys.length);

    for (const k of keys) {
        packUint16(w, k);
        packString(w, extra[k]);
    }
}

function packMapUint32(w, extra) {
    const keys = Object.keys(extra).map(k => parseInt(k));
    keys.sort((a, b) => a - b); // sort keys in ascending order

    packUint16(w, keys.length);

    for (const k of keys) {
        packUint16(w, k);
        packUint32(w, extra[k]);
    }
}

class Packer {
    static packUint16(w,n) {
        packUint16(w, n);
        return Buffer.concat(w);
    }
    static packUint32(w,n) {
        packUint32(w, n);
        return Buffer.concat(w);
    }
    static packUint64(w,n) {
        packUint64(w, n);
        return Buffer.concat(w);
    }
    static packString(w,s) {
        packString(w, s);
        return Buffer.concat(w);
    }
    static packHexString(w,s) {
        packHexString(w, s);
        return Buffer.concat(w);
    }
    static packExtra(w,extra) {
        packExtra(w, extra);
        return Buffer.concat(w);
    }
    static packMapUint32(w,extra) {
        packMapUint32(w, extra);
        return Buffer.concat(w);
    }

}
module.exports = Packer;
