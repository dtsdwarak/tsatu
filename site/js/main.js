// main.js
$(document).ready(function() {
    // Load data from JSON file
    $.getJSON('../data/data.json', function(data) {
      let bookList = $('#bookList');
  
      // Display thumbnails on the main page
      $.each(data.books, function(index, book) {
        let thumbnail = $('<div class="col-md-3 thumbnail">');
        thumbnail.append('<img src="' + book.Images[0] + '" alt="' + book.Title + '" class="img-fluid">');
        thumbnail.append('<p>' + book.Title + '</p>');
        thumbnail.click(function() {
          // Redirect to detail page with book index
          window.location.href = '../detail.html?id=' + index;
        });
        bookList.append(thumbnail);
      });
    });
  });
  