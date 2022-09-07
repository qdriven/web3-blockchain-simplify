#!/usr/bin/env python
# -*- coding: utf-8 -*-
import asyncio
import os
from pprint import pprint

from apscheduler.schedulers.asyncio import AsyncIOScheduler

from raydium.driver import RaydiumFetcher


def run_raydium_job():
    scheduler = AsyncIOScheduler()

    scheduler.add_job(RaydiumFetcher().fetch, 'cron', minute="*")
    scheduler.start()
    print('Press Ctrl+{0} to exit'.format('Break' if os.name == 'nt' else 'C'))
    # Execution will block here until Ctrl+C (Ctrl+Break on Windows) is pressed.
    try:
        asyncio.get_event_loop().run_forever()
    except (KeyboardInterrupt, SystemExit) as e:
        pprint(e)


if __name__ == '__main__':
    run_raydium_job()
