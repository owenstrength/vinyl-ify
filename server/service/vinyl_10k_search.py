# search 10000hz records for the artists and return all of the vinyls for that artist
# search the artists on the website, the website will return everything that is related to that artists,
# included different artists with the same producer or record label

# so each vinyl will be listed as artist name - album name
# so check all of them and return the ones that are the same as the artist name
  
# the website is https://www.10000hzrecords.com/search?q=artist_name
# replace the spaces with + and the website will return the search results

# the website also shows if the vinyl is in stock or not
# if the out of stock message is not found, then the vinyl is in stock
# return the vinyls that are in stock

import json
import sys
import webbrowser
import requests
from bs4 import BeautifulSoup

def search_vinyl(artist):
    query = artist.replace(' ', '+')
    url = 'https://10000hzrecords.com/search?q=' + query

    # Send a GET request to Google search
    response = requests.get(url)

    # Parse the HTML content using BeautifulSoup
    soup = BeautifulSoup(response.text, 'html.parser')

    first_link = None
    search_results = soup.find_all('a')
    res_list = [str(result) for result in search_results if artist.lower() in result.text.lower()]
    album_dict = {}

    for link in res_list:
        album_name_start = link.lower().find('">' + artist[0].lower()) + 2
        album_name_end = link.find('</span>')
        album_name = link[album_name_start:album_name_end]
        
        href_start = link.find('href="') + 6
        href_end = link.split('>')[0]
        href = link[href_start:len(href_end)- 1]
        
        album_dict[album_name.split(" - ")[1]] = "https://10000hzrecords.com" + href

    album_json = json.dumps(album_dict, indent=4)
    return album_json

if __name__ == '__main__':
    # Get the artist's title from command-line arguments
    artist_title = ' '.join(sys.argv[1:])

    # Call the search_vinyl function with the artist's title
    print(search_vinyl(artist_title))