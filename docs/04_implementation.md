# Implementation

## Challenges

- **N+1 query problem:**  
    During implementation of the front-end page, I realized that the existing design `/user/media` returning only relationships and `/media` returning titles would require many API calls to fetch all media items.  
    **Solution:** Updated `/user/media` to join with `media_items` so the response now includes the full media items, reducing the number of queries.
    
- **Missing media image field:**  
    Noticed that a core part of the media the image was missing from the database.  
    **Solution:** Added an `image_url` field to the `media_items` table to store media images.