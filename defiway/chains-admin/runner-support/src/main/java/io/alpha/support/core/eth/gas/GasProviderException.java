package io.alpha.support.core.eth.gas;
/*
 *  Copyright 2020-2021 QMETA - this repo owner
 *
 *  Licensed under unkown License
 */

public class GasProviderException extends RuntimeException{
    public GasProviderException() {
    }

    public GasProviderException(String message) {
        super(message);
    }

    public GasProviderException(String message, Throwable cause) {
        super(message, cause);
    }

    public GasProviderException(Throwable cause) {
        super(cause);
    }

    public GasProviderException(String message, Throwable cause, boolean enableSuppression, boolean writableStackTrace) {
        super(message, cause, enableSuppression, writableStackTrace);
    }
}
