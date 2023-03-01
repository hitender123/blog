package controller
import (
    "net/http"
    "net/http/httptest"
    "strconv"
    "testing"

    "github.com/gin-gonic/gin"
    "blog/model"
)

func TestGetArticleById(t *testing.T) {
    // Create a new Gin router
    router := gin.New()

    // Create a mock model
    mockModel := &model.MockModel{}

    // Add the GetArticleById route to the router
    router.GET("/articles/:id", func(c *gin.Context) {
        GetArticleById(c, mockModel)
    })

    // Test cases
    testCases := []struct {
        name string
        id   int
        code int
    }{
        {"Valid ID", 1, http.StatusOK},
        {"Missing ID", 0, http.StatusOK},
        {"Negative ID", -1, http.StatusOK},
    }

    // Run the test cases
    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            // Create a new HTTP request for the test case
            req, err := http.NewRequest("GET", "/articles/"+strconv.Itoa(tc.id), nil)
            if err != nil {
                t.Fatalf("Could not create request: %v", err)
            }

            // Create a new HTTP response recorder
            resp := httptest.NewRecorder()

            // Perform the request using the router
            router.ServeHTTP(resp, req)

            // Check the response code
            if resp.Code != tc.code {
                t.Errorf("Expected status code %d but got %d", tc.code, resp.Code)
            }
        })
    }
}

func TestGetArticles(t *testing.T) {
    // Create a new Gin router
    router := gin.New()

    // Create a mock model
    mockModel := &model.MockModel{}

    // Add the GetArticles route to the router
    router.GET("/articles/:l/:o", func(c *gin.Context) {
        GetArticles(c, mockModel)
    })

    // Test cases
    testCases := []struct {
        name         string
        limit        string
        offset       string
        expectedCode int
    }{
        {"Valid limit and offset", "10", "0", http.StatusOK},
        {"Invalid limit", "20", "0", http.StatusOK},
        {"Missing limit", "", "0", http.StatusOK},
        {"Negative limit", "-1", "0", http.StatusOK},
    }

    // Run the test cases
    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            // Create a new HTTP request for the test case
            req, err := http.NewRequest("GET", "/articles/"+tc.limit+"/"+tc.offset, nil)
            if err != nil {
                t.Fatalf("Could not create request: %v", err)
            }

            // Create a new HTTP response recorder
            resp := httptest.NewRecorder()

            // Perform the request using the router
            router.ServeHTTP(resp, req)

            // Check the response code
            if resp.Code != tc.expectedCode {
                t.Errorf("Expected status code %d but got %d", tc.expectedCode, resp.Code)
            }
        })
    }
}
