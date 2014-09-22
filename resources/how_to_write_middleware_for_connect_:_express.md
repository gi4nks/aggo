# How to Write Middleware for Connect / Express.js

## Writing Middleware for Connect / Express.js

If have reached to the point of wanting to write your own Express.js middleware, probably you know that Express.js is actually Connect (with additional features). If you didn't know that already, probably you should take it easy and learn more about [Connect][1] first.

If you know what Connect is and want to write a middleware for it, you have come to the right place. Also I will show you how to use your brand new middleware with Express.js.

Before we can write a middleware, we need to know what is one.

Normally, a request to a Connect server is passed though many functions before the HTTP response is generated. These functions are conventionally called "middleware" in Connect terms. That's it, they are that simple.

Working with Expess.js or Connect, you might have come across the `req`, `res`, and `next` objects, the middleware functions do their job and pass on the `req`, `res`, and `next` to the next function in line, if required. That's how middleware work.

Middleware functions are plugged in to the request flow, using `connect.use()` in Connect, and `app.use()` in Express. Take a look at this Connect server example.


    var http = require('http');
    var connect = require('connect');

    var app = connect();
    app.use(function(req, res) {
    res.end('Hello!');
    });

    http.createServer(app).listen(3000);


There is only one middleware, and it just prints 'Hello!' for every request to the server.

Though middlewares are functions that can be plugged in using the `connect.use()` method, it is best to develop them as Node.js modules instead of defining the function within `connect.use()`, to make for a more re-usable and cleaner code.

Let's create a middleware module to ban IP addresses, we will call it `ipban`. Create a file named `ipban.js` with the following content:


    // list of banned IPs
    var banned = [
    '127.0.0.1',
    '192.168.2.12'
    ];

    // the middleware function
    module.exports = function() {
    
    return function(req, res, next) {
    if (banned.indexOf(req.connection.remoteAddress) &gt; -1) {
    res.end('Banned');
    }
    else { next(); }
    }
    
    };


Now, modify the Connect server code to use our middleware:


    var ipban = require('./ipban.js');
    var app = connect();
    app.use(ipban());
    app.use(function(req, res) {
    res.end('Hello!');
    });


Run the app and try loading it on your browser. BAM! Banned!

Following me? A custom middleware module SHOULD return a function, which ideally accepts three parameters - `req`, `res`, and `next`.

You might have noticed that most middlewares accept options when initializing it. It is only natural that we would `ipban` to be able to do so too. By passing 'on' or 'off', the `ipban` will be enabled or disabled.

Modify the `ipban.js` file:


    // list of banned IPs
    var banned = [
    '127.0.0.1',
    '192.168.2.12'
    ];

    // middleware enabled or not
    var enabled = true;

    // the middleware function
    module.exports = function(onoff) {
    
    enabled = (onoff == 'on') ? true : false;
    
    return function(req, res, next) {
    if (enabled &amp;&amp; banned.indexOf(req.connection.remoteAddress) &gt; -1) {
    res.end('Banned');
    }
    else { next(); }
    }
    
    };


Consequently, the usage is also changed:


    app.use(ipban('off'));


Restart the app, and load it on your browser. Banned no more!

If your middleware is not supposed to be the end point of the HTTP request, make sure you call `next()`, else you will end up with a hung application.

To use the middleware with Express.js, we would do this:


    var ipfilter = require('./ipfilter');
    app.configure(function(){
    app.use(ipfilter('on'));
    ...


Note that the order of middleware inclusion in very important. It can make or break your app.

I hope this example helped you understand what middlewares are and how to author them. Could you not follow something in this post? Need to know more about middlewares? Ping me in the comments.

   [1]: https://github.com/senchalabs/connect
  