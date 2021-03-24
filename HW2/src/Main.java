package main;
import java.io.IOException;
import java.net.HttpURLConnection;
import java.util.HashMap;

public class Main {

    private static HttpURLConnection connection;

    public static void main(String[] args) throws IOException, HttpException {
        // Set the main url
        String url = "https://www.tensorflow.org";
        // Set the query params
        HashMap<String, String> params = new HashMap<String, String>();
        params.put("name", "mohammad");
        params.put("family", "hashemi");
        // Set the header properties
        HashMap<String, String> headers = new HashMap<String, String>();
        headers.put("User-Agent", "Java client");
        headers.put("Content-Type", "application/x-www-form-urlencoded");
        // Set the body content
        HashMap<String, String> body = new HashMap<String, String>();
        body.put("job", "data-scientist");
        body.put("hometown", "Tehran, Iran");

        HttpRequest req = new HttpRequest(url, HttpRequestMethod.GET, headers, params, body);

        HttpResponse res = req.request();
        System.out.println("Status: " + res.getStatus());
        System.out.println("Header: " + res.getHeaders());
        System.out.println("Body: " + res.getBody());

    }
}
