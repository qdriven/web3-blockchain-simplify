package io.alpha.compound.intergation.response;


import lombok.Data;

import java.util.List;

@Data
public class CTokenServiceResponse {
    List<CTokenResponse> cToken;
    private Object error;
    private Object request;
}
