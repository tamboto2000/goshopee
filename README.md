# GoShopee

GoShopee is a wraper library for Shopee internal web API. Its still in development, so a lot more features will be added!

# Current Features

  - Lookup item by Shop ID and Item ID
  - Add Item to Cart
  - Checkout all items in cart

### Installation

GoShopee require Go version 1.14 and up

```sh
$ go get gitlab.com/tamboto2000/goshopee
```

### Example

Init Shopee client

```go
import (
    "gitlab.com/tamboto2000/goshopee"
)

func main() {
    sh := goshopee.New()
    // Get your Shopee cookie from browser after login or visiting Shopee web.
    // Make sure to login to Shopee before retrieving the cookie
    sh.SetCookieStr("your_shopee_cookie_string")
    // Do something...
}
```

Item by ID and Shop ID
```go
item, err := sh.ItemByIDAndShopID(6554463078, 155149633)
if err != nil {
	panic(err.Error())
}

// Do something with item...
```

Item by Link
```go
item, err := sh.ItemByLink("https://shopee.co.id/G-W-KULOT-JEANS-WARNA-i.40659222.6060064837")
if err != nil {
	return err
}
```

Add Item to Cart
```go
// info is actually Item, but only descripting the order info
info, err := item.AddToCart(1234, 1) // 1234 is Model ID, 1 is the quantity
if err != nil {
    panic(err.Error())
}

// Do something...
```

### Todos

 - Add payment with Shopeepay
 - Add Travis CI

License
----

MIT
