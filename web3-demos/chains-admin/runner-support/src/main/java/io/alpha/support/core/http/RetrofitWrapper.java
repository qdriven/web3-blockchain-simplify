package io.alpha.support.core.http;

import io.alpha.support.core.http.interceptors.LoggingInterceptor;
import io.alpha.support.core.http.interceptors.RetryInterceptor;
import lombok.extern.slf4j.Slf4j;
import okhttp3.Interceptor;
import okhttp3.OkHttpClient;
import retrofit2.Call;
import retrofit2.Response;
import retrofit2.Retrofit;
import retrofit2.converter.gson.GsonConverterFactory;
import retrofit2.converter.scalars.ScalarsConverterFactory;

import java.io.IOException;
import java.time.Duration;
import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;

@Slf4j
public class RetrofitWrapper {

    private RetrofitWrapper() {
    }

    private static Interceptor loggingInterceptor =
            new LoggingInterceptor();
    private static OkHttpClient defaultHttpClient = defaultHttpClient();
    private static Map<String, Retrofit> retrofitMap = new ConcurrentHashMap<>();

    public static Retrofit getRetrofit(String baseUrl) {

        retrofitMap.putIfAbsent(baseUrl, new Retrofit.Builder()
                .addConverterFactory(ScalarsConverterFactory.create())
                .addConverterFactory(GsonConverterFactory.create())
                .baseUrl(baseUrl)
                .client(defaultHttpClient)
                .build());
        return retrofitMap.get(baseUrl);
    }

    public static OkHttpClient defaultHttpClient() {
        return new OkHttpClient.Builder()
                .addInterceptor(loggingInterceptor)
                .addInterceptor(new RetryInterceptor())
                .retryOnConnectionFailure(false)
                .connectTimeout(Duration.ofSeconds(15))
                .build();
    }

    public static <T> Response<T> send(Call<T> request) {
        try {
            return request.execute();
        } catch (IOException e) {
            throw new HttpException(e);
        }
    }

    public static <T> T sendAndReturnBody(Call<T> request) {
        try {
            return request.execute().body();
        } catch (IOException e) {
            throw new HttpException(e);
        }
    }

}
