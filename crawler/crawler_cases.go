package crawler

var cases = []struct {
	input  []string
	expect map[string]string
}{
	{
		[]string{"https://www.airbnb.com/help/article/5"},
		map[string]string{
			"What does the room type of a listing mean?": "Hosts on Airbnb offer a wide variety of spaces, ranging from shared rooms to private islands. While you can filter searches to look for something more specific (like boats or castles!), all listings are categorized into the following three room types: shared rooms, private rooms, and entire homes/apartments. Shared rooms are for when you’d like a comfy, communal experience, and don't mind sharing a space with others. When you book a shared room, you'll be sharing your bedroom and the entire space with another person or persons. These rooms work best for the flexible traveler looking to make friends, and a budget-friendly stay. Private rooms are great for when you prefer a little privacy, yet still value a local connection. When you book a private room, you'll have a bedroom to yourself, but will share some spaces with others. With a private room, you'll be able to wake up to greet your new friends in the kitchen and have the freedom to bid them adieu at bedtime. Entire homes/apartments are for you if you are seeking a home away from home with complete privacy and the freedom to cook breakfast in your pajamas. With an entire home/apartment, you'll have the whole space to yourself. You can be your own host, make your own dinner, and remember to treat your listing with the respect and courtesy you would at your own home.",
		},
	},
	{
		[]string{"https://www.airbnb.com/help/article/13"},
		map[string]string{
			"How do reviews work?": "All the reviews on Airbnb are written by hosts and travelers from our community, so any review you see is based on a stay that a guest had in a host's listing. Writing a review To leave a review for a recent trip, click on the reminder in the Alerts section of your dashboard, or go to Edit Profile > Reviews > Reviews By You . They are limited to 500 words and must follow Airbnb's review guidelines . It's important that you don't leave personally identifiable information in a review, like a person's last name or address. Editing a review After writing a review, if your host or guest didn't complete their review yet, you can edit your review for up to 48 hours, or until your host or guest completes their review. You can edit a review at Edit Profile > Reviews > Reviews by You > Edit . Previous reviews To see your review history, go to Edit Profile > Reviews . You'll see past Reviews About You and Reviews By You . Among the Reviews About You , you'll also see any Private Feedback that people have left you. This is also where you can also leave a public response to a review you've received within the past 2 weeks. Star ratings In addition to written reviews, the guests who stay in a listing can submit a primary star rating and a set of category star ratings. The number of stars displayed at the top of a listing page is an aggregate of the primary scores guests have given for that listing. At the bottom of a listing page there's an aggregate for each category rating. A host needs to receive star ratings from at least 3 guests before their aggregate score appears. Our community relies on honest, transparent reviews. We will remove or alter a review if we find that it violates our review guidelines .",
		},
	},
	{
		[]string{"https://www.airbnb.com/help/article/4"},
		map[string]string{
			"Does Airbnb screen users?": "Airbnb does not routinely perform background checks on its users, although we reserve the right to do so. Instead, we strive to provide our hosts and guests with the right tools to make informed decisions regarding who they interact with on the site and in the real world. Airbnb offers a number of features that help build trust and cultivate a transparent community marketplace, including our secure messaging system, user reviews , the Host Guarantee , and Verified ID . At Airbnb, we want to build the world's most trusted community. To help ensure our members' safety, we dedicate a knowledgeable and experienced team to monitor any suspicious activity in our marketplace. Users can also help ensure the safety of the community by flagging activity they feel is suspicious or inappropriate. We investigate each flag on a case-by-case basis and follow up as appropriate. In any user profile, listing, or message thread, you will see an icon with a small flag which allows you to notify our team who can investigate. Here’s an example of what the flag looks like: The user profile flag is located in the lower right-hand corner of any user's description box. The listing flag is located in the upper right-hand corner of any listing's description box. The message thread flag is located in the bottom right-hand corner of any message box. Read more about the tools used to help verify user information .",
		},
	},
	{
		[]string{"https://www.airbnb.com/help/article/10"},
		map[string]string{
			"I'm a member of the media, who should I correspond with?": "For information on recent press and media coverage, please see our Press page . If you have a specific media question, please contact: press@airbnb.com . If you're interested in applying for a job at Airbnb, please visit https://www.airbnb.com/jobs .",
		},
	},
	{
		[]string{"https://www.airbnb.com/help/article/18"},
		map[string]string{
			"Who can host on Airbnb?": "Almost anyone can be a host! It's free to sign up and list your space. The listings available on the site are as diverse as the hosts who list them, so you can post airbeds in apartments, entire houses, rooms in bed-and-breakfasts, tree houses in the woods, boats on the water, or enchanted castles. Find out more about room types to see how to describe your space. You can list your space in almost any location worldwide. While we would like to keep the Airbnb marketplace open to the entire world, we are required to comply with US federal regulations that restrict the use of our site by residents of certain countries.",
		},
	},
	{
		[]string{"https://www.airbnb.com/help/article/14"},
		map[string]string{
			"How do text message and push notifications work?": "Text message (SMS) notifications We can let you know when you have new Airbnb messages, reservation requests, and more via text message (SMS) alerts. These notifications may be sent at any time, based on your notification settings. To turn text message notifications on or off and manage the types of mobile notifications you receive, visit the notifications menu in your account settings : Click your name in the top-right corner of airbnb.com Select Account Click Notifications in the menu on the left side of your screen Look for the Mobile Notifications / Text Messages pane. You can manage your settings here. When you opt in to text message (SMS) notifications, we'll send you a text to confirm your signup. We only support sending SMS on carriers in North America and Europe. Standard message and data rates may apply for any messages sent or received, so it's best to contact your wireless provider to confirm. Participating carriers: AT&T, CBW, nTelos, Sprint, Nextel, Boost, Virgin Mobile, Verizon Wireless, T-Mobile, and Alltel. Push notifications If you use Airbnb's mobile app , we can also send push notifications straight to your smart phone. You can enable them from the app's Settings menu.",
		},
	},
	{
		[]string{"https://www.airbnb.com/help/article/9"},
		map[string]string{
			"How do I contact Airbnb?": "To talk with someone from Airbnb, visit our contact page and select the issue you need help with. We're available 24 hours a day, 7 days a week. Most hosts and guests are able to work out reservation-related issues on their own. If you have a reservation-related issue, try sending a message to your host or guest first. This is often the quickest way to resolve problems that come up before, during, or after your stay. We do our best to keep our phone lines open for people with urgent situations. If you need to call us, use the local phone number included in your original reservation confirmation email, or click Urgent after you select your issue on our contact page.",
		},
	},
}
