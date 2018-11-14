#!/usr/bin/env python
# coding=utf-8
import os
import logging
import time

import redis

REDIS_CONFIG = {
    'host': os.environ.get('REDIS_HOST', '127.0.0.1'),
    'port': int(os.environ.get('REDIS_PORT', 6379)),
    'key': b'honorlist'
}

LUA_SCRIPT = """
local key = KEYS[1]
local data = redis.call('HGET', key, 'To')
redis.call('ZINCRBY', 'honorlist', 1, data)
redis.call('DEL', key)
"""


def gen_honor_list():
    rc = redis.StrictRedis(host=REDIS_CONFIG['host'], port=REDIS_CONFIG['port'])
    for k in rc.scan_iter():
        if k != REDIS_CONFIG['key']:
            try:
                rc.eval(LUA_SCRIPT, 1, k)
            except Exception as e:
                logging.error('key: {}, msg: {}'.format(k, e))


if __name__ == '__main__':
    while 1:
        print('update honorlist...')
        gen_honor_list()
        time.sleep(5)
