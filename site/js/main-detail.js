// main-detail.js
$(document).ready(function() {
    // Load data from JSON file
    $.getJSON('../data/data.json', function(data) {
      let bookDetails = $('#bookDetails');
  
      // Get the book index from the URL query parameter
      const urlParams = new URLSearchParams(window.location.search);
      const bookIndex = urlParams.get('id');
  
      if (bookIndex !== null && bookIndex >= 0 && bookIndex < data.books.length) {
        // Display details for the selected book
        let book = data.books[bookIndex];
        bookDetails.append('<h2>' + book.Title + '</h2>');
        bookDetails.append('<p><strong>Description:</strong> ' + book.Description + '</p>');
        bookDetails.append('<p><strong>Author:</strong> ' + book.Author + '</p>');
        bookDetails.append('<p><strong>Rating:</strong> ' + book.Rating + '</p>');
        bookDetails.append('<p><strong>Images:</strong></p>');
  
        // Display all images for the book
        let imagesContainer = $('<div class="row">');
        $.each(book.Images, function(i, imageUrl) {
          let image = $('<div class="col-md-3">');
          image.append('<img src="' + imageUrl + '" alt="' + book.Title + '" class="img-fluid">');
          imagesContainer.append(image);
        });
        bookDetails.append(imagesContainer);
      } else {
        // Handle invalid book index
        bookDetails.append('<p>Invalid book index</p>');
      }
    });
  });
  