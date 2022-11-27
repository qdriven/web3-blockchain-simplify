#!/usr/bin/env python
import asyncio
from unittest import TestCase

# -*- coding: utf-8 -*-
from raydium.driver import RaydiumFetcher


class TestRaydiumFetcher(TestCase):
    def test_fetch(self):
        loop = asyncio.get_event_loop()
        loop.run_until_complete(RaydiumFetcher().fetch())
        loop.close()
