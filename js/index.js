'use strict';

const dgram = require('dgram');

class UDPLogger {
  constructor(addr, port, max) {
    this.addr = addr;
    this.port = port;
    this.conn = dgram.createSocket('udp4');
    this.max = (max ? max : 1472) - 36;
  }

  send(msg, cb) {
    if (!msg) return;
    if (!msg.check) msg.check = 0;
    if (!msg.group) msg.group = 0;
    if (!msg.size) msg.size = 0;
    if (!msg.identity) msg.identity = 0;
    if (!msg.duration) msg.duration = 0;
    if (!msg.timestamp) msg.timestamp = 0;
    if (!msg.remaining) msg.remaining = '';

    const buf1 = Buffer.alloc(36);
    buf1.writeUInt16BE(msg.check, 0);
    buf1.writeUInt16BE(msg.group, 2);
    buf1.writeUInt32BE(msg.code, 4);
    buf1.writeUInt32BE(msg.size, 8);
    buf1.writeBigUInt64BE(BigInt(msg.identity), 12);
    buf1.writeBigUInt64BE(BigInt(msg.duration), 20);
    buf1.writeBigUInt64BE(BigInt(msg.timestamp), 28);

    const buf2 = msg.remaining.length <= this.max ?
      Buffer.from(msg.remaining, 'utf8') :
      Buffer.from(msg.remaining.substring(0, this.max), 'utf8');

    this.conn.send([buf1, buf2], this.port, this.addr, cb);
  }
}

exports.UDPLogger = UDPLogger;