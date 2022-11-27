#!/usr/bin/env python
# -*- coding: utf-8 -*-
import asyncio
from pprint import pprint

import aiohttp
from aiolimiter import AsyncLimiter

from core.fetcher import ChainDataFetcher
from raydium.lp_config import RaydiumPoolInfo


class RaydiumFetcher(ChainDataFetcher):
    def __init__(self):
        pass

    async def fetch(self):
        limiter = AsyncLimiter(15, 1)
        task = []
        async with aiohttp.ClientSession() as session:
            async with limiter:
                LP = RaydiumPoolInfo(session)
                fee_apy_task = asyncio.create_task(LP.RAYDIUM.get_pair())
                fee_apy = await fee_apy_task
                for farm in LP.farms_info:
                    task.append(LP.get_APR(farm))
                result = await asyncio.gather(*task)

                for farm in result:
                    farm_name = list(farm.keys())[0]
                    farm[farm_name].update({"Fee_APR": fee_apy[farm_name]})
                    pprint(farm)
