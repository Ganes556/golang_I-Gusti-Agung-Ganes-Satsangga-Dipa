# **Rangkuman Introduction Echo Golang**

## **Framework Echo**

- performa tinggi, ekstensibilitas, dan minimalist
- keuntungan:
  - optimized router
  - middleware
  - data rendering
  - scalable
  - data binding
- struktur pada framework echo bebas yang bisa diintegrasikan dengan third party lib lainnya

## **Basic Echo**

- `echo.New()` untuk membuat instance echo
- routing adalah url yang akan diakses oleh user, dapat berupa GET, POST, PUT, DELETE, dan lainnya

  - Contoh pada Echo:
    - `e.GET("/hello", hello)`
    - `e.POST("/hello", hello)`
    - `e.PUT("/hello", hello)`
    - `e.DELETE("/hello", hello)`
    - etc

- controller pada echo menggunakan `echo.Context` untuk mempermudah routing
  - Contoh:
    ```go
    func hello(c echo.Context) error {
    	return c.String(http.StatusOK, "Hello, World!")
    }
    ```
- mengambil param dari url
  - Contoh:
    ```go
    func hello(c echo.Context) error {
    	name := c.Param("name")
    	return c.String(http.StatusOK, "Hello, "+name)
    }
    ```
- mengambil query dari url
  - Contoh:
    ```go
    func hello(c echo.Context) error {
    	name := c.QueryParam("name")
    	return c.String(http.StatusOK, "Hello, "+name)
    }
    ```
- mengambil data dari form
  - Contoh:
    ```go
    func hello(c echo.Context) error {
    	name := c.FormValue("name")
    	return c.String(http.StatusOK, "Hello, "+name)
    }
    ```
- binding data json ke struct

  - Contoh:

    ```go
    type User struct {
    	Name  string `json:"name"`
    	Email string `json:"email"`
    }

    func hello(c echo.Context) error {
    	u := new(User)
    	if err := c.Bind(u); err != nil {
    		return err
    	}
    	return c.JSON(http.StatusOK, u)
    }
    ```
