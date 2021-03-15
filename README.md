# Go_TODO
RESTfull TODO list on Go
Link: https://golang-todo-list.herokuapp.com/
Endpoints:
  1)/auth:
    1)/sign-up POST Body: name, username, password
    2)/sign-in POST Body: username, password Returns JWT token
  2)/api(JWT Token required):
    1)/lists
      1)/ GET Getting all user's lists
      2)/ POST Body: title, description
      3)/:id GET Get list by id
      4)/:id PUT Update list by id. Body: title and/or description
      5)/:id DELETE Delete list by id
      6)/:id/items
        1)/ POST Body: title, description, done(boolean). Create item
        2)/ GET Get all items
    2)/items
      1)/:id GET Get item by id
      2)/:id PUT Body: title and/or description and/or done. Update item
      3)/:id DELETE item by id
