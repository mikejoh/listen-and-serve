package sparkexample;

import static spark.Spark.get;

public class Headers {
    public static void main(String[] args) {
        get("/headers", (request, response) -> {
            return request.headers();
        });
    }
}
