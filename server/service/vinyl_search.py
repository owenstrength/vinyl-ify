import sys
import webbrowser
import requests
from bs4 import BeautifulSoup

def search_vinyl(artist):
    query = artist + ' vinyl'
    url = 'https://www.google.com/search?q=' + query

    # Send a GET request to Google search
    response = requests.get(url)

    # Parse the HTML content using BeautifulSoup
    soup = BeautifulSoup(response.text, 'html.parser')

    first_link = None
    res = []
    search_results = soup.find_all('a')
    for link in search_results:
        href = link.get('href')
        if href.startswith('/url?q='):
            res.append(href[7:].split('&')[0])
        if len(res) == 3:
            break
            
    return res[2]

if __name__ == '__main__':
    # Get the artist's title from command-line arguments
    artist_title = ' '.join(sys.argv[1:])

    # Call the search_vinyl function with the artist's title
    print(search_vinyl(artist_title))