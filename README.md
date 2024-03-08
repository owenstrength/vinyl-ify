# Vinylify

Vinylify is a web application that allows you to search your Spotify top artists and find where to buy vinyl records for your favorite artists.

## Project Status
The project is currently being developed
✅ means the feature is implemented
❌ means the feature is planned but not implemented yet

## Features

- Connects to your Spotify account to retrieve your top artists. ✅
- Provides a search functionality to find vinyl records for your favorite artists. ❌
- Displays information about the vinyl records, including price, availability, and seller information. ❌
- Allows you to add vinyl records to your wishlist for future reference.❌
- Provides links to purchase the vinyl records directly from the sellers' websites.❌


## Usage

1. Log in to your Spotify account using the provided login button. ✅
2. Once logged in, your favorite artists appear. ✅
3. Additonally you can enter the name of your favorite artist in the search bar and click "Search".❌
4. Browse through the search results and click on a vinyl record to view more details.❌
5. To add a vinyl record to your wishlist, click on the "Add to Wishlist" button.❌
6. To purchase a vinyl record, click on the "Buy Now" button to be redirected to the seller's website.❌

## How it works

Using the Spotify API, the application retrieves the user's top artists. The user's top artists are displayed on the home page. The user can also search for vinyl records for their favorite artists. While the Discogs API seems like a good option, I want this project to benefit the artists, the record labels, and local record stores. Therefore, I will be using the Spotify API to retrieve the artist's information and then use the artist's name to search for vinyl records on the artist's official website or the record label's website or a local record store's website.

Each search will first look to see if it is in stock at a local record store
If it is not in stock at a local record store, it will look to see if it is in stock at the record label's website
If it is not in stock at the record label's website, it will look to see if it is in stock at the artist's official website

in order to reduce the number of API calls, the search results will be cached for a certain amount of time. This is helpful for popular artists who have a lot of searches. But it is also helpful for less popular artists who have fewer searches. 

in order to increase speed, the backend is written in go and will utilize goroutines to make multiple API calls at once.

## Contributing

Contributions are welcome! If you have any ideas, suggestions, or bug reports, please open an issue or submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE).