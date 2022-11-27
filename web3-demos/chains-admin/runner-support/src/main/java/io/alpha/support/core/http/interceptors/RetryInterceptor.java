package io.alpha.support.core.http.interceptors;

import lombok.AllArgsConstructor;
import okhttp3.Interceptor;
import okhttp3.Request;
import okhttp3.Response;
import org.slf4j.LoggerFactory;

import java.io.IOException;
import java.lang.invoke.MethodHandles;

@AllArgsConstructor
public class RetryInterceptor implements Interceptor {
    private static final org.slf4j.Logger log =
            LoggerFactory.getLogger(MethodHandles.lookup().lookupClass().getSimpleName());
    private final int retryTimes = 3;

    @Override
    public Response intercept(Chain chain) throws IOException {

        Request request = chain.request();
        Response response = doRequest(chain, request);
        int tryCount = 0;
        while (response == null && tryCount < retryTimes) {
            Request newRequest = request.newBuilder().build();
            response = doRequest(chain, newRequest);
            tryCount++;
            log.warn("Request failed, retry to acquire a new connection, {} in {} times", tryCount, retryTimes);
        }
        if (response == null) { // Important ,should throw an exception here
            throw new IOException();
        }
        return response;
    }

    /**
     * do retry
     * @param chain
     * @param request
     * @return
     */
    private Response doRequest(Chain chain, Request request) {
        Response response = null;
        try {
            response = chain.proceed(request);
        } catch (Exception e) {
            log.error("", e);
        }
        return response;
    }
}
