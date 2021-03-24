package main;

import java.io.*;
import java.net.*;
import java.nio.charset.StandardCharsets;

import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class HttpRequest {
    private String url;
    private HttpRequestMethod httpRequestMethod;
    private HashMap<String, String> headers;
    private HashMap<String, String> params;
    private HashMap<String, String> body;


    public HttpRequest(String url) {
        this(url, HttpRequestMethod.GET, new HashMap<>(), new HashMap<>(), new HashMap<>());
    }

    public HttpRequest(String url, HttpRequestMethod httpRequestMethod) {
        this(url, httpRequestMethod, new HashMap<>(), new HashMap<>(), new HashMap<>());
    }

    public HttpRequest(String url, HttpRequestMethod httpRequestMethod, HashMap<String, 
            String> headers, HashMap<String, String> params, HashMap<String, String> body) {
        this.url = url;
        this.httpRequestMethod = httpRequestMethod;
        this.headers = headers;
        this.params = params;
        this.body = body;
    }

    public String getUrl() {
        return this.url;
    }

    public void setUrl(String url) {
        this.url = url;
    }

    public HttpRequestMethod getHttpRequestMethod() {
        return this.httpRequestMethod;
    }

    public void setHttpRequestMethod(HttpRequestMethod httpRequestMethod) {
        this.httpRequestMethod = httpRequestMethod;
    }

    public HashMap<String, String> getHeaders() {
        return this.headers;
    }

    public void setHeaders(HashMap<String, String> headers) {
        this.headers = headers;
    }

    public HashMap<String, String> getParams() {
        return this.params;
    }

    public void setParams(HashMap<String, String> params) {
        this.params = params;
    }

    public HashMap<String, String> getBody() {
        return this.body;
    }

    public void setBody(HashMap<String, String> body) {
        this.body = body;
    }

    public String apply_parameters_to_url(HashMap<String, String> params, String base_url) {

        // A function to apply parameters to the base url
        String url = base_url;
        if (!params.isEmpty()){
            url += "?";
            for (Map.Entry<String, String> entry : params.entrySet()){
                url += entry.getKey() + "=" + entry.getValue() + "&";
            }
            url = url.substring(0, url.length() - 1);
        }

        return url;
    }

    public HttpResponse request() throws HttpException, IOException {

        //  Apply the parameters to base url
        String url_in_string = apply_parameters_to_url(params, this.url);
        // Declare our connection
        HttpURLConnection connection = null;
        // Get an instance of URL
        URL url = new URL(url_in_string);
        // Open the connection
        connection = (HttpURLConnection) url.openConnection();
        // Set the request method
        connection.setRequestMethod(httpRequestMethod.toString());
        // Set a limit for the connection time out
        connection.setConnectTimeout(6000);
        // Apply user headers
        for (Map.Entry<String, String> entry : headers.entrySet()) {
            connection.setRequestProperty(entry.getKey(), entry.getValue());
        }

        if (httpRequestMethod == HttpRequestMethod.GET
                                        || httpRequestMethod == HttpRequestMethod.HEAD
                                        || httpRequestMethod == HttpRequestMethod.OPTIONS) {
            connection.setDoOutput(false);
            connection.setDoInput(true);
        } else { // POST or DELETE or PUT
            // Set the "content-type" request header to "application/json" to send request content in JSON form.
            connection.setRequestProperty("Content-Type", "application/json; utf-8");
            // Set the "Accept" request header to "application/json" to read the response in the desired format.
            connection.setRequestProperty("Accept", "application/json");
            // Ensure the connection will be used to send content
            connection.setDoOutput(true);
            // Create the request body
            try (OutputStream os = connection.getOutputStream()) {
                byte[] input = body.toString().getBytes("utf-8");
                os.write(input, 0, input.length);
            }
        }

        // Get response code
        int status = connection.getResponseCode();
        // Read response
        BufferedReader reader;
        String line = "";
        StringBuffer responseContent = new StringBuffer();
        String body = "";
        Map<String, List<String>> headers;
        HttpResponse response;

        if (status == 200 || status == 201 || status == 204) { // Successful responses
            reader = new BufferedReader(new InputStreamReader(connection.getInputStream()));
            while ((line = reader.readLine()) != null) {
                responseContent.append(line);
            }
            reader.close();
            // Set the body, headers & finally response
            body = responseContent.toString();
            headers = connection.getHeaderFields();
            response = new HttpResponse(status, headers, body);

        } else { // Fail responses
            reader = new BufferedReader(new InputStreamReader(connection.getErrorStream()));
            while ((line = reader.readLine()) != null) {
                responseContent.append(line);
            }
            reader.close();
            // Set the body, headers & finally response
            body = responseContent.toString();
            headers = connection.getHeaderFields();
            response = new HttpResponse(status, headers, body);

            connection.disconnect();
            throw new HttpException(status + ": " + connection.getResponseMessage(), response);
        }
        connection.disconnect();
        return response;
    }
}
