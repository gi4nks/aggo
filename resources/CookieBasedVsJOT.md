#Cookies vs Tokens. Getting auth right with Angular.JS
* * * 

##Introduction

There are basically two different ways of implementing server side authentication for apps with a frontend and an API:

* The most adopted one, is *Cookie-Based Authentication* (you can find an example here) that uses server side cookies to authenticate the user on every request.

* A newer approach, *Token-Based Authentication*, relies on a signed token that is sent to the server on each request.

![Image](./Cookie-based Auth vs Token-based Auth.png)


##Token based vs. Cookie based

What are the benefits of using a token-based approach?

* Cross-domain / CORS: cookies + CORS don't play well across different domains. A token-based approach allows you to make AJAX calls to any server, on any domain because you use an HTTP header to transmit the user information.
Stateless (a.k.a. Server side scalability): there is no need to keep a session store, the token is a self-contanined entity that conveys all the user information. The rest of the state lives in cookies or local storage on the client side.
* CDN: you can serve all the assets of your app from a CDN (e.g. javascript, HTML, images, etc.), and your server side is just the API.
* Decoupling: you are not tied to a particular authentication scheme. The token might be generated anywhere, hence your API can be called from anywhere with a single way of authenticating those calls.
Mobile ready: when you start working on a native platform (iOS, Android, Windows 8, etc.) cookies are not ideal when consuming a secure API (you have to deal with cookie containers). Adopting a token-based approach simplifies this a lot.
* CSRF: since you are not relying on cookies, you don't need to protect against cross site requests (e.g. it would not be possible to `<iframe>` your site, generate a POST request and re-use the existing authentication cookie because there will be none).
Performance: we are not presenting any hard perf benchmarks here, but a network roundtrip (e.g. finding a session on database) is likely to take more time than calculating an HMACSHA256 to validate a token and parsing its contents.
* Login page is not an special case: If you are using Protractor to write your functional tests, you don't need to handle any special case for login.
* Standard-based: your API could accepts a standard *JSON Web Token* (JWT). This is a standard and there are multiple backend libraries (.NET, Ruby, Java, Python, PHP) and companies backing their infrastructure (e.g. Firebase, Google, Microsoft). As an example, Firebase allows their customers to use any authentication mechanism, as long as you generate a JWT with certain pre-defined properties, and signed with the shared secret to call their API.

What's JSON Web Token? JSON Web Token (JWT, pronounced _jot_) is a relatively new token format used in space-constrained environments such as HTTP Authorization headers. JWT is architected as a method for transferring security claims based between parties.

##Implementation

Asuming you have a node.js app, below you can find the components of this architecture.

###Server Side

Let's start by installing express-jwt:

	$ npm install express-jwt

Configure the express middleware to protect every call to /api.

	var expressJwt = require('express-jwt');
	
	// We are going to protect /api routes with JWT
	app.use('/api', expressJwt({secret: secret}));
	
	app.use(express.json());
	app.use(express.urlencoded());
	The angular app will perform a POST through AJAX with the user's credentials:
	
	app.post('/authenticate', function (req, res) {
	  //TODO validate req.body.username and req.body.password
	  //if is invalid, return 401
	  if (!(req.body.username === 'john.doe' && req.body.password === 'foobar')) {
	    res.send(401, 'Wrong user or password');
	    return;
	  }
	
	  var profile = {
	    first_name: 'John',
	    last_name: 'Doe',
	    email: 'john@doe.com',
	    id: 123
	  };
	
	  // We are sending the profile inside the token
	  var token = jwt.sign(profile, secret, { expiresInMinutes: 60*5 });
	
	  res.json({ token: token });
	});

GET'ing a resource named /api/restricted is straight forward. Notice that the credentials check is performed by the expressJwt middleware.

	app.get('/api/restricted', function (req, res) {
	  console.log('user ' + req.user.email + ' is calling /api/restricted');
	  res.json({
	    name: 'foo'
	  });
	});

##Angular Side

The first step on the client side using AngularJS is to retrieve the JWT Token. In order to do that we will need user credentials. We will start by creating a view with a form where the user can input its username and password.
	
	<div ng-controller="UserCtrl">
	  <span></span>
	  <form ng-submit="submit()">
	    <input ng-model="user.username" type="text" name="user" placeholder="Username" />
	    <input ng-model="user.password" type="password" name="pass" placeholder="Password" />
	    <input type="submit" value="Login" />
	  </form>
	</div>

And a controller where to handle the submit action:

	myApp.controller('UserCtrl', function ($scope, $http, $window) {
	  $scope.user = {username: 'john.doe', password: 'foobar'};
	  $scope.message = '';
	  $scope.submit = function () {
	    $http
	      .post('/authenticate', $scope.user)
	      .success(function (data, status, headers, config) {
	        $window.sessionStorage.token = data.token;
	        $scope.message = 'Welcome';
	      })
	      .error(function (data, status, headers, config) {
	        // Erase the token if the user fails to log in
	        delete $window.sessionStorage.token;
	
	        // Handle login errors here
	        $scope.message = 'Error: Invalid user or password';
	      });
	  };
	});

Now we have the JWT saved on sessionStorage. If the token is set, we are going to set the Authorization header for every outgoing request done using $http. As value part of that header we are going to use Bearer <token>.

*sessionStorage*: Although is not supported in all browsers (you can use a polyfill) is a good idea to use it instead of cookies ($cookies, $cookieStore) and localStorage: The data persisted there lives until the browser tab is closed.

	myApp.factory('authInterceptor', function ($rootScope, $q, $window) {
	  return {
	    request: function (config) {
	      config.headers = config.headers || {};
	      if ($window.sessionStorage.token) {
	        config.headers.Authorization = 'Bearer ' + $window.sessionStorage.token;
	      }
	      return config;
	    },
	    response: function (response) {
	      if (response.status === 401) {
	        // handle the case where the user is not authenticated
	      }
	      return response || $q.when(response);
	    }
	  };
	});
	
	myApp.config(function ($httpProvider) {
	  $httpProvider.interceptors.push('authInterceptor');
	});
	After that, we can send a request to a restricted resource:
	
	$http({url: '/api/restricted', method: 'GET'})
	.success(function (data, status, headers, config) {
	  console.log(data.name); // Should log 'foo'
	});

The server logged to the console:

	user foo@bar.com is calling /api/restricted

The source code is here together with an AngularJS seed app.