package io.alpha.support.core.http;

import cn.hutool.core.bean.BeanUtil;
import cn.hutool.json.JSONUtil;
import lombok.extern.slf4j.Slf4j;
import okhttp3.MediaType;
import okhttp3.OkHttpClient;
import okhttp3.RequestBody;
import okhttp3.Response;

import java.io.IOException;
import java.net.InetSocketAddress;
import java.net.Proxy;
import java.util.Map;

@Slf4j
public final class OkHttpClientWrapper {
    private static OkHttpClient client = createClient();
    private okhttp3.Request.Builder builder = new okhttp3.Request.Builder();
    private String requestUrl;


    private static void resetClient() {
        client.connectionPool().evictAll();
        client = createClient();
    }

    private static OkHttpClient createClient() {
        //防止SSLPeerUnverifiedException问题，加verifier
        return new OkHttpClient.Builder()
                .hostnameVerifier((hostname, session) -> true).build();
    }

    public OkHttpClientWrapper addQueryParameters(Map<String, Object> queryParameters) {
        this.requestUrl = this.requestUrl + "?";
        for (Map.Entry<String, Object> entry : queryParameters.entrySet()) {
            this.requestUrl = this.requestUrl + entry.getKey() + "=" + entry.getValue().toString() + "&";
        }
        builder.url(this.requestUrl);
        return this;
    }

    public OkHttpClientWrapper get(Object queryBean) {
        Map<String, Object> queryParameters = BeanUtil.beanToMap(queryBean);
        return addQueryParameters(queryParameters);
    }

    public OkHttpClientWrapper url(String url) {
        this.requestUrl = url;
        builder.url(this.requestUrl);
        return this;
    }

    public OkHttpClientWrapper addHeader(String name, String value) {
        builder.addHeader(name, value);
        return this;
    }

    public OkHttpClientWrapper addHeaders(Map<String, Object> headers) {
        for (Map.Entry<String, Object> entry : headers.entrySet()) {
            builder.addHeader(entry.getKey(), entry.getValue().toString());
        }
        return this;
    }

    public OkHttpClientWrapper addHeader(String name, boolean value) {
        builder.addHeader(name, Boolean.toString(value));
        return this;
    }

    public OkHttpClientWrapper get() {
        builder.get();
        return this;
    }

    public OkHttpClientWrapper post(RequestBody body) {
        builder.post(body);
        return this;
    }

    public <T> OkHttpClientWrapper post(T object) {
        builder.post(RequestBody.create(JSONUtil.toJsonStr(object),MediaType.parse("application/json; charset=utf-8")));
        return this;
    }

    public OkHttpClientWrapper put(RequestBody body) {
        builder.put(body);
        return this;
    }

    public OkHttpClientWrapper delete() {
        builder.delete();
        return this;
    }

    public OkHttpClientWrapper delete(RequestBody body) {
        builder.delete(body);
        return this;
    }

    public static void setProxy(String hostname, int port) {
        InetSocketAddress address = new InetSocketAddress(hostname, port);
        Proxy proxy = new Proxy(Proxy.Type.HTTP, address);
        client.newBuilder().proxy(proxy);
    }

    public static void cleanProxy() {
        client.newBuilder().proxy(Proxy.NO_PROXY);
    }


    public Response execute() {
        okhttp3.Request request = builder.build();
        try {
            return client.newCall(request).execute();
        } catch (IOException e) {
            OkHttpClientWrapper.resetClient();
            try {
                return client.newCall(request).execute();
            } catch (IOException ioException) {
                log.error("http request failed,", e);
            }
        }
        throw new RuntimeException("Http Call failed in " + request.toString());
    }
}
