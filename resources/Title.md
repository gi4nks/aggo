<html xmlns="http://www.w3.org/1999/xhtml" xml:lang="en" lang="en">
<head>
<script type="text/javascript" src="http://ajax.aspnetcdn.com/ajax/jQuery/jquery-1.6.2.min.js"></script>
<script type="text/javascript" src="waypoints.min.js"></script>
<script type="text/javascript">
$(function(){
	var link = document.getElementById('thisone');
	link.href = link.href.replace("email.me", "gmail.com")
		.replace("enable_js_to", "glurgle");
});
</script>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
<meta name="generator" content="Docutils 0.8: http://docutils.sourceforge.net/" />
<title>A short guide to Connect Middleware</title>
<link rel="stylesheet" href="index.css" type="text/css" />
</head>
<body>
<div class="document" id="a-short-guide-to-connect-middleware">
<h1 class="title">A short guide to Connect Middleware</h1>

<p>This guide will introduce you to <a class="reference external" href="http://senchalabs.github.com/connect">connect</a> and the concept of
middleware. If you are familiar with middleware from other systems such as WSGI
or Rack, you can probably skip the middle section and go straight to the
<a class="reference internal" href="#examples">Examples</a>.</p>
<div class="section" id="what-is-connect">
<h1>What is connect?</h1>
<p>From the <a class="reference external" href="http://github.com/senchalabs/connect">README</a>:</p>
<blockquote>
Connect is an extensible HTTP server framework for node,
providing high performance &quot;plugins&quot; known as middleware.</blockquote>
<p>More specifically, connect wraps the <a class="reference external" href="http://nodejs.org/docs/v0.4.12/api/http.html#http.Server">Server</a>, <a class="reference external" href="http://nodejs.org/docs/v0.4.12/api/http.html#http.ServerRequest">ServerRequest</a>, and
<a class="reference external" href="http://nodejs.org/docs/v0.4.12/api/http.html#http.ServerResponse">ServerResponse</a> objects of node.js' standard <tt class="docutils literal">http</tt> module, giving them a few
nice extra features, one of which is allowing the <tt class="docutils literal">Server</tt> object to <tt class="docutils literal">use</tt>
a stack of <cite>middleware</cite>.</p>
</div>
<div class="section" id="what-is-middleware">
<h1>What is middleware?</h1>
<p>Simply put, middleware are functions that handle requests. A server created
by <tt class="docutils literal">connect.createServer</tt> can have a stack of middleware associated with it.
When a request comes in, it is passed off to the first middleware function,
along with a wrapped <a class="reference external" href="http://nodejs.org/docs/v0.4.12/api/http.html#http.ServerResponse">ServerResponse</a> object and a <tt class="docutils literal">next</tt> callback. Each
middleware can decide to respond by calling methods on the response object, and/or
pass the request off to the next layer in the stack by calling <tt class="docutils literal">next()</tt>. A simple no-op middleware would look like this:</p>
<div class="highlight" style="background: #f0f0f0"><pre style="line-height: 125%"><span style="color: #007020; font-weight: bold">function</span> uselessMiddleware(req, res, next) { next() }
</pre></div>
<p>A middleware can also signal an error by passing it as the first argument to
<tt class="docutils literal">next</tt>:</p>
<div class="highlight" style="background: #f0f0f0"><pre style="line-height: 125%"><span style="color: #60a0b0; font-style: italic">// A middleware that simply interrupts every request</span>
<span style="color: #007020; font-weight: bold">function</span> worseThanUselessMiddleware(req, res, next) {
  next(<span style="color: #4070a0">&quot;Hey are you busy?&quot;</span>)
}
</pre></div>
<p>When a middleware returns an error like this, all subsequent middleware will be
skipped until connect can find an error handler. (See <a class="reference internal" href="#error-handling">Error Handling</a> for an
example).</p>
<p>To add a middleware to the stack of middleware for a server, we <tt class="docutils literal">use</tt> it like
so:</p>
<div class="highlight" style="background: #f0f0f0"><pre style="line-height: 125%">connect <span style="color: #666666">=</span> require(<span style="color: #4070a0">&#39;connect&#39;</span>)
stephen <span style="color: #666666">=</span> connect.createServer()
stephen.use(worseThanUselessMiddleware)
</pre></div>
<p>Finally, you can also specify a path prefix when adding a middleware, and the
middleware will only be asked to handle requests that match the prefix:</p>
<div class="highlight" style="background: #f0f0f0"><pre style="line-height: 125%">connect <span style="color: #666666">=</span> require(<span style="color: #4070a0">&#39;connect&#39;</span>)
bob <span style="color: #666666">=</span> connect.createServer()
bob.use(<span style="color: #4070a0">&#39;/attention&#39;</span>, worseThanUselessMiddleware)
</pre></div>
</div>
<div class="section" id="what-can-i-use-it-for">
<h1>What can I use it for?</h1>
<p>Plenty of things! Common examples are logging, serving static files, and error
handling. Note that all three of the above (and more) are standard middleware
included with connect itself, so you probably won't need to implement them
yourself. Another common middleware is routing requests to controllers or
callback methods (for this and <em>way</em> more, check out TJ Holowaychuk's <a class="reference external" href="http://expressjs.com">express</a>).</p>
<p>Really, you can use middleware anywhere you might want to have some sort of
generic request handling logic applied to all requests. For example, in my
<a class="reference external" href="http://github.com/BetSmartMedia/Lazorse">Lazorse</a> project, request routing and response rendering are separate
middleware, because it's designed around building API's that may want to use
different rendering backends depending on the clients <tt class="docutils literal">Accept</tt> header.</p>
</div>
<div class="section" id="examples">
<h1>Examples</h1>
<div class="section" id="url-based-authentication-policy">
<h2>URL based authentication policy</h2>
<p>Authentication policy is often app-specific, so this
middleware wraps the <tt class="docutils literal">connect.basicAuth</tt> middleware with a list of URL
patterns that should be authenticated.</p>
<div class="highlight" style="background: #f0f0f0"><pre style="line-height: 125%"><span style="color: #007020; font-weight: bold">function</span> authenticateUrls(urls <span style="color: #60a0b0; font-style: italic">/* basicAuth args*/</span>) {
  basicAuthArgs <span style="color: #666666">=</span> <span style="color: #007020">Array</span>.slice.call(arguments, <span style="color: #40a070">1</span>)
  basicAuth <span style="color: #666666">=</span> require(<span style="color: #4070a0">&#39;connect&#39;</span>).basicAuth.apply(basicAuthArgs)
  <span style="color: #007020; font-weight: bold">function</span> authenticate(req, res, next) {
    <span style="color: #60a0b0; font-style: italic">// Warning! assumes that urls is amenable to iteration</span>
    <span style="color: #007020; font-weight: bold">for</span> (<span style="color: #007020; font-weight: bold">var</span> pattern <span style="color: #007020; font-weight: bold">in</span> urls) {
      <span style="color: #007020; font-weight: bold">if</span> (req.path.match(pattern)) {
        basicAuth(req, res, next)
        <span style="color: #007020; font-weight: bold">return</span>
      }
    }
    next()
  }
  <span style="color: #007020; font-weight: bold">return</span> authenticate
}
</pre></div>
<p>See the <a class="reference external" href="http://senchalabs.github.com/connect/middleware-basicAuth.html">basicAuth</a> docs for the various options it takes.</p>
</div>
<div class="section" id="role-base-authorization">
<h2>Role base authorization</h2>
<div class="highlight" style="background: #f0f0f0"><pre style="line-height: 125%"><span style="color: #60a0b0; font-style: italic">// @urls - an object mapping patterns to lists of roles.</span>
<span style="color: #60a0b0; font-style: italic">// @roles - an object mapping usernames to lists of roles</span>
<span style="color: #007020; font-weight: bold">function</span> authorizeUrls(urls, roles) {
  <span style="color: #007020; font-weight: bold">function</span> authorize(req, res, next) {
    <span style="color: #007020; font-weight: bold">for</span> (<span style="color: #007020; font-weight: bold">var</span> pattern <span style="color: #007020; font-weight: bold">in</span> urls) {
      <span style="color: #007020; font-weight: bold">if</span> (req.path.match(pattern)) {
        <span style="color: #007020; font-weight: bold">for</span> (<span style="color: #007020; font-weight: bold">var</span> role <span style="color: #007020; font-weight: bold">in</span> urls[pattern]) {
          <span style="color: #007020; font-weight: bold">if</span> (users[req.remoteUser].indexOf(role) <span style="color: #666666">&lt;</span> <span style="color: #40a070">0</span>) {
            next(<span style="color: #007020; font-weight: bold">new</span> <span style="color: #007020">Error</span>(<span style="color: #4070a0">&quot;unauthorized&quot;</span>))
            <span style="color: #007020; font-weight: bold">return</span>
          }
        }
      }
    }
    next()
  }
  <span style="color: #007020; font-weight: bold">return</span> authorize
}
</pre></div>
<p>These examples demonstrate how middleware can help isolate cross-cutting request
logic into modules that can be swapped out. If we later decide to replace
basicAuth with OAuth, the only dependency our authorization module has is on the
<tt class="docutils literal">.remoteUser</tt> property.</p>
</div>
<div class="section" id="error-handling">
<h2>Error handling</h2>
<p>Another common cross-cutting concern is error handling. Again, connect ships
with an handler that will serve up nice error responses to clients, but finding
out about errors in production via your customers is usually not a good
business practice ;). To help with that, let's implement a simple error
notification middleware using <a class="reference external" href="https://github.com/marak/node_mailer">node_mailer</a>:</p>
<div class="highlight" style="background: #f0f0f0"><pre style="line-height: 125%">email <span style="color: #666666">=</span> require(<span style="color: #4070a0">&#39;node_mailer&#39;</span>)

<span style="color: #007020; font-weight: bold">function</span> emailErrorNotifier(generic_opts, escalate) {
  <span style="color: #007020; font-weight: bold">function</span> notifyViaEmail(err, req, res, next) {
    <span style="color: #007020; font-weight: bold">if</span> (err) {
      <span style="color: #007020; font-weight: bold">var</span> opts <span style="color: #666666">=</span> {
        subject<span style="color: #666666">:</span> <span style="color: #4070a0">&quot;ERROR: &quot;</span> <span style="color: #666666">+</span> err.constructor.name,
        body<span style="color: #666666">:</span> err.stack,
      }
      opts.__proto__ <span style="color: #666666">=</span> generic_opts
      email.send(opts, escalate)
    }
    next()
  }
}
</pre></div>
<p>Connect checks the arity of middleware functions and considers the function to
be an error handler if it is 4, meaning errors returned by earlier middleware
will be passed as the first argument for inspection.</p>
</div>
<div class="section" id="putting-it-all-together">
<h2>Putting it all together</h2>
<p>Here's a simple app that combines all of the above middleware:</p>
<div class="highlight" style="background: #f0f0f0"><pre style="line-height: 125%">private_urls <span style="color: #666666">=</span> {
  <span style="color: #4070a0">&#39;^/attention&#39;</span><span style="color: #666666">:</span> [<span style="color: #4070a0">&#39;coworker&#39;</span>, <span style="color: #4070a0">&#39;girlfriend&#39;</span>],
  <span style="color: #4070a0">&#39;^/bank_balance&#39;</span><span style="color: #666666">:</span> [<span style="color: #4070a0">&#39;me&#39;</span>],
}

roles <span style="color: #666666">=</span> {
  stephen<span style="color: #666666">:</span> [<span style="color: #4070a0">&#39;me&#39;</span>],
  erin<span style="color: #666666">:</span>    [<span style="color: #4070a0">&#39;girlfriend&#39;</span>],
  judd<span style="color: #666666">:</span>    [<span style="color: #4070a0">&#39;coworker&#39;</span>],
  bob<span style="color: #666666">:</span>     [<span style="color: #4070a0">&#39;coworker&#39;</span>],
}

passwords <span style="color: #666666">=</span> {
  me<span style="color: #666666">:</span>   <span style="color: #4070a0">&#39;doofus&#39;</span>,
  erin<span style="color: #666666">:</span> <span style="color: #4070a0">&#39;greatest&#39;</span>,
  judd<span style="color: #666666">:</span> <span style="color: #4070a0">&#39;daboss&#39;</span>,
  bob<span style="color: #666666">:</span>  <span style="color: #4070a0">&#39;anachronistic discomBOBulation&#39;</span>,
}

<span style="color: #007020; font-weight: bold">function</span> authCallback(name, password) { <span style="color: #007020; font-weight: bold">return</span> passwords[name] <span style="color: #666666">===</span> password }

stephen <span style="color: #666666">=</span> require(<span style="color: #4070a0">&#39;connect&#39;</span>).createServer()
stephen.use(authenticateUrls(private_urls, authCallback))
stephen.use(authorizeUrls(private_urls, roles))
stephen.use(<span style="color: #4070a0">&#39;/attention&#39;</span>, worseThanUselessMiddleware)
stephen.use(emailErrorNotifier({to<span style="color: #666666">:</span> <span style="color: #4070a0">&#39;stephen@betsmartmedia.com&#39;</span>}))
stephen.use(<span style="color: #4070a0">&#39;/bank_balance&#39;</span>, <span style="color: #007020; font-weight: bold">function</span> (req, res, next) {
  res.end(<span style="color: #4070a0">&quot;Don&#39;t be Seb-ish&quot;</span>)
})
stephen.use(<span style="color: #4070a0">&#39;/&#39;</span>, <span style="color: #007020; font-weight: bold">function</span> (req, res, next) {
  res.end(<span style="color: #4070a0">&quot;I&#39;m out of coffee&quot;</span>)
})
</pre></div>
</div>
</div>
<div class="section" id="but-i-want-to-use-objects">
<h1>But I want to use objects!</h1>
<p>Perfectly reasonable, pass an object with a <tt class="docutils literal">handle</tt> method to <tt class="docutils literal">use</tt> and
connect will call that method in the exact same manner.</p>
</div>
<div class="section" id="contact">
<h1>Contact</h1>
<p>Please submit all feedback, flames and compliments to
  <a id="thisone" href="mailto:enable_js_to@email.me">Stephen</a>.
</p></div>
</div>
</body>
</html>
