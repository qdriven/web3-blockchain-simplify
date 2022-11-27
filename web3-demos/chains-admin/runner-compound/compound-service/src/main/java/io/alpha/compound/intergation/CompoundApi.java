package io.alpha.compound.intergation;

import io.alpha.compound.intergation.request.AccountRequest;
import io.alpha.compound.intergation.response.AccountServiceResponse;
import io.alpha.compound.intergation.response.CTokenServiceResponse;
import retrofit2.Call;
import retrofit2.http.Body;
import retrofit2.http.GET;
import retrofit2.http.POST;

public interface CompoundApi {
    @POST("/api/v2/account")
    Call<AccountServiceResponse> accountService(@Body AccountRequest accountRequest);

    @GET("/api/v2/ctoken")
    Call<CTokenServiceResponse> ctokenService();
}
