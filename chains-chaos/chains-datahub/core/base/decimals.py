#!/usr/bin/env python
# -*- coding: utf-8 -*-

__all__ = [
    'TRUNCATE',
    'ROUND',
    'ROUND_UP',
    'ROUND_DOWN',
    'DECIMAL_PLACES',
    'SIGNIFICANT_DIGITS',
    'TICK_SIZE',
    'NO_PADDING',
    'PAD_WITH_ZERO',
    'decimals',
]

# rounding mode
TRUNCATE = 0
ROUND = 1
ROUND_UP = 2
ROUND_DOWN = 3

# digits counting mode
DECIMAL_PLACES = 2
SIGNIFICANT_DIGITS = 3
TICK_SIZE = 4

# padding mode
NO_PADDING = 5
PAD_WITH_ZERO = 6


def decimal_to_precision(n, rounding_mode=ROUND, precision=None,
                         counting_mode=DECIMAL_PLACES, padding_mode=NO_PADDING):
    assert precision is not None
    if counting_mode == TICK_SIZE:
        pass
