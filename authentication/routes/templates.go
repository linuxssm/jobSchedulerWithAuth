package routes

const BaseTmplStr = `
{{ define "base" }}
<html>
<head>
	<title>Node Authentication</title>
	<link rel="stylesheet" href="//netdna.bootstrapcdn.com/bootstrap/3.0.2/css/bootstrap.min.css"> <!-- load bootstrap css -->
	<link rel="stylesheet" href="//netdna.bootstrapcdn.com/font-awesome/4.0.3/css/font-awesome.min.css"> <!-- load fontawesome -->
	<style>
		body 		{ padding-top:80px; }
	</style>
</head>
<body>
    <div class="container">
        <div class="jumbotron text-center">
            <h1><span class="fa fa-lock"></span> Node Authentication</h1>

            <p>Login or Register with:</p>

            <a href="/login" class="btn btn-default"><span class="fa fa-user"></span> Local Login</a>
            <a href="/signup" class="btn btn-default"><span class="fa fa-user"></span> Local Signup</a>
        </div>
    </div>
</body>
</html>

{{ end }}
`

const LoginTmplStr = `
{{ define "base" }}
<html>
<head>
	<title>Node Authentication</title>
	<link rel="stylesheet" href="//netdna.bootstrapcdn.com/bootstrap/3.0.2/css/bootstrap.min.css"> <!-- load bootstrap css -->
	<link rel="stylesheet" href="//netdna.bootstrapcdn.com/font-awesome/4.0.3/css/font-awesome.min.css"> <!-- load fontawesome -->
	<style>
		body 		{ padding-top:80px; }
	</style>
</head>
<body>
    <div class="container">
        <div class="col-sm-6 col-sm-offset-3">
            <h1><span class="fa fa-sign-in"></span> Login</h1>

            <!-- show any messages that come back with authentication
            //<% if (message.length > 0) { %>
            //    <div class="alert alert-danger"><%= message %></div>
            //<% } %>
-->
            <!-- LOGIN FORM -->
            <form action="/login" method="post">
                <div class="form-group">
                    <label>Username</label>
                    <input type="text" class="form-control" name="username">
                </div>
                <div class="form-group">
                    <label>Password</label>
                    <input type="password" class="form-control" name="password">
                </div>
                <div class="form-group">
                    <label>Remember Me</label>
                    <input type="checkbox" class="form-control" name="remember" value="yes">
                </div>

                <button type="submit" class="btn btn-warning btn-lg">Login</button>
            </form>

            <hr>

            <p>Need an account? <a href="/signup">Signup</a></p>
            <p>Or go <a href="/">home</a>.</p>
        </div>
    </div>
</body>
</html>
{{ end }}
`

const profileTmplStr = `
{{ define "base" }}
<html>
<head>
	<title>Node Authentication</title>
	<link rel="stylesheet" href="//netdna.bootstrapcdn.com/bootstrap/3.0.2/css/bootstrap.min.css">
	<link rel="stylesheet" href="//netdna.bootstrapcdn.com/font-awesome/4.0.3/css/font-awesome.min.css">
	<style>
		body 		{ padding-top:80px; word-wrap:break-word; }
	</style>
</head>
<body>
<div class="container">

	<div class="page-header text-center">
		<h1><span class="fa fa-anchor"></span> Profile Page</h1>
		<a href="/logout" class="btn btn-default btn-sm">Logout</a>
	</div>

	<div class="row">

		<!-- LOCAL INFORMATION -->
		<div class="col-sm-6">
			<div class="well">
				<h3><span class="fa fa-user"></span> Local</h3>

					<p>
						<strong>id</strong>: <%= user.id %><br>
						<strong>username</strong>: <%= user.username %><br>
						<strong>password</strong>: <%= user.password %>
					</p>

			</div>
		</div>

	</div>

</div>
</body>
</html>
{{ end }}
`

const signupTmplStr = `
{{ define "base" }}
<html>
<head>
	<title>Node Authentication</title>
	<link rel="stylesheet" href="//netdna.bootstrapcdn.com/bootstrap/3.0.2/css/bootstrap.min.css"> <!-- load bootstrap css -->
	<link rel="stylesheet" href="//netdna.bootstrapcdn.com/font-awesome/4.0.3/css/font-awesome.min.css"> <!-- load fontawesome -->
	<style>
		body 		{ padding-top:80px; }
	</style>
</head>
<body>
<div class="container">

<div class="col-sm-6 col-sm-offset-3">

	<h1><span class="fa fa-sign-in"></span> Signup</h1>

	<!-- show any messages that come back with authentication -->
	<% if (message.length > 0) { %>
		<div class="alert alert-danger"><%= message %></div>
	<% } %>

	<!-- LOGIN FORM -->
	<form action="/signup" method="post">
		<div class="form-group">
			<label>Username</label>
			<input type="text" class="form-control" name="username">
		</div>
		<div class="form-group">
			<label>Password</label>
			<input type="password" class="form-control" name="password">
		</div>

		<button type="submit" class="btn btn-warning btn-lg">Signup</button>
	</form>

	<hr>

	<p>Already have an account? <a href="/login">Login</a></p>
	<p>Or go <a href="/">home</a>.</p>

</div>

</div>
</body>
</html>

{{ end }}
`
