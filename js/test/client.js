'use strict';

const { UDPLogger } = require('../index');

const logger = new UDPLogger('127.0.0.1', 7777);

const send = () => logger.send({
  check: 1,
  group: 1,
  size: 10,
  code: 200,
  identity: 666666,
  duration: 1000,
  timestamp:1611855906,
  remaining: '~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~end'
}, send)

send()