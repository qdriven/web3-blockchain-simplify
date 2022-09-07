package io.alpha.compound.intergation.request;
/*
 *  Copyright 2020-2021 QMETA - this repo owner
 *
 *  Licensed under unkown License
 */

import lombok.Data;

import java.util.List;

@Data
public class CTokenRequest {
    private List<String> addresses;
    private long block_timestamp;
}
