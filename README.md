# Tori ロボット

[![Go Report Card](https://goreportcard.com/badge/github.com/rfguri/tori)](https://goreportcard.com/report/github.com/rfguri/tori)

`Tori` is an [Airbnb](https://www.airbnb.com/) FAQ scraper used to get data from airbnb faq page to use it as the input data for a AI engine that will use NLP.

## Features

* Fetch URLs in parallel using [`gorutines`](https://golang.org/doc/effective_go.html#goroutines) and [`channels`](https://golang.org/doc/effective_go.html#channels).
* Set a window space time (*`window`*) between *HTTP GET* requests.
* Disable the *SSL verification* to run it against an https-secured site that can’t validate the certificate using [`TLS 1.2`](https://en.wikipedia.org/wiki/Transport_Layer_Security), as specified in *RFC 5246*.

## Tests

`Tori` comes with a fully loaded suite to do some proper test the crawler agains a list of URLs. These are located in [crawler_cases.go](https://github.com/rfguri/tori/blob/master/crawler/crawler_cases.go) and follows a **struct** setup.

##### Example:

```Go
var cases = []struct {
  input  []string
  expect map[string]string
}{
  {
    []string {"https://www.airbnb.com/help/article/221"},
    map[string]string{
      "How do I create an account?": "If you don't have an Airbnb account yet, head to: www.airbnb.com/signup_login Signing up and creating an Airbnb account is totally free, whether you're looking to travel or list your space! After you sign up, please be sure to complete your profile.",
  },
}
```

##### Running tests

```Shell
$ go test github.com/rfguri/tori/crawler -v
```

## Examples

##### Example usage

```Shell
$ tori 0 1 5
```

##### Example Result

	[HTTP Request] https://www.airbnb.com/help/article/5
	[HTTP Request] https://www.airbnb.com/help/article/2
	[HTTP Request] https://www.airbnb.com/help/article/1
	[HTTP Request] https://www.airbnb.com/help/article/4
	[HTTP Request] https://www.airbnb.com/help/article/3

	[HTTP Response] https://www.airbnb.com/help/article/4
	[HTTP Response] https://www.airbnb.com/help/article/5

	[Tori] Found 2 unique questions:

	Q: Does Airbnb screen users?
	A: Airbnb does not routinely perform background checks on its users, although we reserve the right to do so. Instead, we strive to provide our hosts and guests with the right tools to make informed decisions regarding who they interact with on the site and in the real world.Airbnb offers a number of features that help build trust and cultivate a transparent community marketplace, including our secure messaging system, userreviews, theHost Guarantee, andVerified ID.At Airbnb, we want to build the world's mosttrustedcommunity. To help ensure our members' safety, we dedicate a knowledgeable and experienced team to monitor any suspicious activity in our marketplace.Users can also help ensure the safety of the community by flagging activity they feel is suspicious or inappropriate. We investigate each flag on a case-by-case basis and follow up as appropriate. In any user profile, listing, or message thread, you will see an icon with a small flag which allows you to notify our team who can investigate. Here’s an example of what the flag looks like:The user profile flag is located in the lower right-hand corner of any user's description box. The listing flag is located in the upper right-hand corner of any listing's description box. The message thread flag is located in the bottom right-hand corner of any message box.Read more about the tools used to helpverify user information.

	Q: What does the room type of a listing mean?
	A: Hosts on Airbnb offer a wide variety of spaces, ranging from shared rooms to private islands.While you canfilter searchesto look for something more specific (like boats or castles!), all listings are categorized into the following three room types: shared rooms, private rooms, and entire homes/apartments.Shared roomsare for when you’d like a comfy, communal experience, and don't mind sharing a space with others. When you book a shared room, you'll be sharing your bedroom and the entire space with another person or persons. These rooms work best for the flexible traveler looking to make friends, and a budget-friendly stay.Private roomsare great for when you prefer a little privacy, yet still value a local connection. When you book a private room, you'll have a bedroom to yourself, but will share some spaces with others. With a private room, you'll be able to wake up to greet your new friends in the kitchen and have the freedom to bid them adieu at bedtime.Entire homes/apartmentsare for you if you are seeking a home away from home with complete privacy and the freedom to cook breakfast in your pajamas. With an entire home/apartment, you'll have the whole space to yourself. You can be your own host, make your own dinner, and remember to treat your listing with the respect and courtesy you would at your own home.

## Running

To run `tori` just issue:

```Shell
$ go get
$ go build
$ tori <window> <start> <end>
```

## License

`Tori` is Copyright © 2017 Roger Fernandez Guri. It is free software, and may be
redistributed under the terms specified in the [LICENSE](https://github.com/rfguri/tori/blob/master/LICENSE) file.
