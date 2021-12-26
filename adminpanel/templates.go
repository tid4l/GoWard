package adminpanel

const InternalPage = `
<!DOCTYPE html>
<html lang="en">
<head>
	<style type="text/css">
	@import url(https://fonts.googleapis.com/css?family=Roboto:300);
		*,
	*::before,
	*::after {
	padding: 0;
	margin: 0;
	-webkit-box-sizing: inherit;
	box-sizing: inherit;
	}

	:root {
	--primary-color: #6B0000;
	--secondary-color: #efefef;
	--dark-green-color: #efefef;
	--white-color: #383232;
	--main-content-color: #383232;
	--grey-color: #777;
	}

	::-moz-selection {
	background-color: var(--primary-color);
	color: var(--white-color);
	}

	::selection {
	background-color: var(--primary-color);
	color: var(--white-color);
	}

	html {
	font-size: 62.5%;
	-webkit-box-sizing: border-box;
	box-sizing: border-box;
	background-color: #383232;
	background-image: linear-gradient(90deg,#383232,#8E8E8E);
	}

	body {
	/*line-height: 1.6;*/
	/*font-weight: 300;*/
	font-family: "Roboto", sans-serif;		
	color: var(--grey-color);
	min-height: 100vh;
	}

	.menu__checkbox {
		display: none;
	}

	.side-view {
	height: 100%;
	width: 30rem;
	position: fixed;
	z-index: 999;
	background-image: linear-gradient(
		to right bottom,
		var(--dark-green-color),
		var(--secondary-color)
	);
	color: var(--white-color);
	box-shadow: 0 0 20px 0 rgba(0, 0, 0, 0.2), 0 5px 5px 0 rgba(0, 0, 0, 0.24);
	}


	.admin-view__header {
	display: flex;
	justify-content: center;
	-webkit-box-align: center;
	-ms-flex-align: center;
	align-items: center;
	height: 6rem;
	padding: 0 2.5rem;
	padding-top: 2rem;
	margin-bottom: 4rem;
	}

	.menu__checkbox:checked ~ .side-view {
	width: 9.2rem;
	}

	.menu__checkbox:checked ~ .main-content {
	margin-left: 9.2rem;
	}

	.menu__checkbox:checked ~ .side-view .admin-view__menu .user-profile,
	.menu__checkbox:checked ~ .side-view .admin-view__menu .admin-view__header h3,
	.menu__checkbox:checked ~ .side-view .admin-view__menu .side-nav li a span,
	.menu__checkbox:checked ~ .side-view .footer {
	display: none;
	}

	.admin-view__header h3 {
	font-size: 2.0rem;
	}

	.menu__checkbox:checked ~ .side-view .admin-view__menu .side-nav li a svg {
	position: fixed;
	}

	.menu__checkbox:checked
	~ .side-view
	.admin-view__menu
	.side-nav
	li:not(:last-child) {
	margin-bottom: 3.5rem;
	}

	.user-profile {
	display: -webkit-box;
	display: -ms-flexbox;
	display: flex;
	flex-direction: column;
	-webkit-box-align: center;
	-ms-flex-align: center;
	align-items: center;
	margin-bottom: 5rem;
	}

	.user-profile .admin-name {
	font-size: 1.3rem;
	}

	.user-profile img {
	height: 10rem;
	width: 10.3rem;
	border-radius: 50%;
	margin-bottom: 1.5rem;
	}

	.side-nav {
	list-style: none;
	}

	.side-nav li {
	margin: 1.5rem 0;
	border-left: 0 solid #fff;
	-webkit-transition: all 0.1s;
	transition: all 0.1s;
	}

	.side-nav__active,
	.side-nav li:hover {
	border-left: 4px solid #6B0000 !important;
	}

	.side-nav--active a {
	-webkit-transform: translateX(-3px);
	transform: translateX(-3px);
	}

	.side-nav li a:link,
	.side-nav li a:visited {
	padding: 1rem 3rem;
	display: -webkit-box;
	display: -ms-flexbox;
	display: flex;
	-webkit-box-align: center;
	-ms-flex-align: center;
	align-items: center;
	color: var(--white-color);
	font-size: 1.6rem;
	font-weight: 600;
	text-decoration: none;
	-webkit-transition: all 0.3s;
	transition: all 0.3s;
	}

	.side-nav li a:hover,
	.side-nav li a:active {
	-webkit-transform: translateX(3px);
	transform: translateX(3px);
	}

	.footer {
	width: 100%;
	position: absolute;
	bottom: 1rem;
	text-align: center;
	font-size: 1.1rem;
	font-weight: 400;
	padding: 0 0.5rem;
	}

	.main-content {
	margin-left: 30rem;
	}

	.main-content .header {
	height: 27.55rem;
	padding-top: 10rem;
	padding-left: 9rem;
	}

	.header p {
	font-size: 1.5rem;
	}

	.overview-cards {
	display: grid;
	max-width: 80%;
	grid-template-columns: repeat(3, 1fr);
	grid-column-gap: 12rem;
	padding-left: 15rem;
	transform: translateY(-5rem);
	}

	.overview-cards .card {
	background-image: linear-gradient(
		to right bottom,
		var(--secondary-color),
		var(--dark-green-color)
	);
	color: var(--white-color);
	padding: 2.5rem 3.5rem;

	width: 30rem;
	height: 18rem;
	box-shadow: 0 0 20px 0 rgba(0, 0, 0, 0.2), 0 5px 5px 0 rgba(0, 0, 0, 0.24);
	}

	.overview-cards .card .title {
	font-size: 1.2rem;
	}

	.overview-cards .card .content {
	margin-top: 2.5rem;
	margin-left: 3rem;
	display: -webkit-box;
	display: -ms-flexbox;
	display: flex;
	-webkit-box-align: center;
	-ms-flex-align: center;
	align-items: center;
	}

	.overview-cards .card svg {
	height: 6rem;
	width: 6rem;
	fill: var(--white-color);
	}

	.content svg {
	margin-right: 7rem;
	}

	.content .number {
	font-size: 2.9rem;
	}

	/* Screen resolution and menu post-checked */

	.menu__checkbox:checked ~ .proxy-section {
	margin-left: 9.2rem;
	}
	.menu__checkbox:checked ~ .requests-section {
	margin-left: 9.2rem;
	}

	@media screen and (max-width: 1600px) {
	.side-view {
		width: 12rem;
	}

	.main-content {
		margin-left: 9.2rem;
		width: 142rem;
	}

	.proxy-section {
		margin-left: 9.2rem;
		width: 142rem;
	}

	.requests-section {
		margin-left: 9.2rem;
		width: 142rem;
	}

	.side-view .admin-view__menu .side-nav li a svg {
		position: fixed;
	}
	}

	@media screen and (max-width: 1400px) {
	.main-content {
		margin-left: 3rem;
		width: 151rem;
	}

	.proxy-section {
		margin-left: 3rem;
		width: 151rem;
		height: 100.8rem;
	}
	.requests-section {
		margin-left: 3rem;
		width: 151rem;
		height: 100.8rem;
	}

	.overview-cards {
		display: grid;
		max-width: 80%;
		grid-template-columns: repeat(3, 1fr);
		grid-column-gap: 8rem;
		padding-left: 15rem;
		transform: translateY(-5rem);
	}
	}

	@media screen and (max-width: 1256px) {
	.overview-cards {
		display: grid;
		max-width: 70%;
		grid-template-columns: repeat(3, 1fr);
		grid-column-gap: 8rem;
		padding-left: 15rem;
		transform: translateY(-5rem);
	}

	.overview-cards .card {
		background-image: linear-gradient(
		to right bottom,
		var(--secondary-color),
		var(--dark-green-color)
		);
		color: var(--white-color);
		padding: 2.5rem 3.5rem;
		width: 25rem;
		height: 15rem;
		box-shadow: 0 0 20px 0 rgba(0, 0, 0, 0.2), 0 5px 5px 0 rgba(0, 0, 0, 0.24);
	}
	}

	@media screen and (max-width: 1108px) {
	.overview-cards {
		display: grid;
		max-width: 40%;
		grid-template-columns: repeat(3, 1fr);
		grid-column-gap: 8rem;
		padding-left: 15rem;
		transform: translateY(-5rem);
	}

	.overview-cards .card {
		background-image: linear-gradient(
		to right bottom,
		var(--secondary-color),
		var(--dark-green-color)
		);
		color: var(--white-color);
		padding: 2rem 3rem;
		box-shadow: 0 0 20px 0 rgba(0, 0, 0, 0.2), 0 5px 5px 0 rgba(0, 0, 0, 0.24);
		width: 22rem;
		height: 15rem;
	}
	}

	@media screen and (max-width: 345px) {
	.main-content {
		margin-left: 4rem;
	}

	.overview-cards {
		display: grid;
		max-width: 80%;
		grid-template-columns: repeat(3, 1fr);
		grid-column-gap: 4rem;
		padding-left: 15rem;
		transform: translateY(-5rem);
	}
	}
	.form button {
		font-family: "Roboto", sans-serif;
		text-transform: uppercase;
		outline: 0;
		background-color: #6B0000;
		background-image: linear-gradient(45deg,#6B0000,#A10101);
		width: 80%;
		border: 0;
		padding: 10px;
		color: #FFFFFF;
		font-size: 12px;
		-webkit-transition: all 0.3 ease;
		transition: all 0.3 ease;
		cursor: pointer;
	}
	</style>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
	<!--<meta name="refresh" http-equiv="refresh" content="3" />-->
    <meta name="viewport" content="width=device-width, initial-scale=1,maximum-scale=1">
    <link href="https://fonts.googleapis.com/css?family=Lato:300,400,400i,700" rel="stylesheet">
    <title>Admin Dashboard</title>
</head>
<body>
	<!-- SIDEBAR -->
	<input type="checkbox" class="menu__checkbox" id="sideview-crawl">
	<div class="side-view">
		<nav class="admin-view__menu">
			<div class="admin-view__header">
				<h3 class="site-name">
					<span><i>GoVector</i></span>
				</h3>
				<div class="menu-icon">
					<label for="sideview-crawl" class="menu-bar">
					</label>
				</div>
			</div>
			<ul class="side-nav">
				<li class="side-nav__active">
					<a href="#">
						<span>Home</span>
					</a>
				</li>
				<li class="side-nav__proxy">
					<a href="#">
						<span>Proxy (WIP)</span>
					</a>
				</li>
				<li class="side-nav__requests">
					<a href="#">
						<span>Requests (WIP)</span>
					</a>
				</li>
			</ul>
		</nav>

		<footer class="footer">
			<div class="form">
			<form method="post" action="/logout" class="logout-form">
				<button type="submit">Logout</button>
			</form>
			</div>
		</footer>
	</div>
	<!-- MAIN --->
	<main class="main main-content">
		<div class="header">
			<h1>Welcome back, Admin</h1>
			<p>Last login: {{.LastLogin}}</p>
		</div>
		
		<div class="overview-cards">
			<div class="card status-card">
				<div class="title">
					<h2>Backend C2 Status</h2>
				</div>
				<span class="content status-content">
					<div class="number">
						<h4><span id ="ws-onlineHosts">{{.OnlineHosts}}</span> of <span id="ws-totalHosts">{{.TotalHosts}}</span> online</h4>
					</div>
				</span>
			</div>
			<div class="card connection-card">
				<div class="title">
					<h2>C2 Connections (WIP)</h2>
				</div>
				<span class="content connection-content">
					<div class="number">
						<h4 id="ws-c2">0</h4>
					</div>
				</span>
			</div>
			<div class="card request-card">
				<div class="title">
					<h2>Web Requests</h2>
				</div>
				<span class="content request-content">
					<div class="number">
						<h4 id="ws-requests">{{.RequestCounter}}</h4>
					</div>
				</span>
			</div>
		</div>
	</main>
	<!-- Proxy Section -->
	<section class="proxy-section">
	</section>
	<!-- Requests Section -->
	<section class="requests-section">
	</section>
	<!-- SCRIPTS -->
	<script type="text/javascript">
	(function() {
		var c2Data = document.getElementById("ws-c2");
		var reqData = document.getElementById("ws-requests")
		var onlineHostData = document.getElementById("ws-onlineHosts")
		var totalHostData = document.getElementById("ws-totalHosts")
		var conn = new WebSocket("ws://{{.RequestHost}}:80/ws");
		conn.onclose = function(evt) {
			data.textContent = 'Connection closed';
		}
		conn.onmessage = function(evt) {
			console.log('file updated');
			var messages = evt.data.replace('}','')
			var messages = evt.data.replace('{','')
			messages = messages.split(' ');
			reqData.textContent = messages[0]
			onlineHostData.textContent = messages[1]
			totalHostData.textContent = messages[2]
		}
	})();
	</script> 
	<script type="text/javascript">
	document.querySelector(".side-nav__proxy").addEventListener("click", (e) => {
		document.querySelector(".main").style.display = "none";
		document.querySelector(".requests-section").style.display = "none";
		document.querySelector(".proxy-section").style.display = "block";
	});
	document.querySelector(".side-nav__requests").addEventListener("click", (e) => {
		document.querySelector(".main").style.display = "none";
		document.querySelector(".requests-section").style.display = "block";
		document.querySelector(".proxy-section").style.display = "none";
	});  
	document.querySelector(".side-nav__active").addEventListener("click", (e) => {
		document.querySelector(".main").style.display = "block";
		document.querySelector(".proxy-section").style.display = "none";
		document.querySelector(".requests-section").style.display = "none";
	});
	</script>   
	</body>
</html>
`

const LoginPage = `
<!DOCTYPE html>
<html>
<head>
<style type= text/css>
	@import url(https://fonts.googleapis.com/css?family=Roboto:300);
	header .header{
	background-color: #fff;
	height: 45px;
	}
	header a img{
	width: 134px;
	margin-top: 4px;
	}
	.login-page {
	width: 360px;
	padding: 8% 0 0;
	margin: auto;
	}
	.login-page .form .login{
	margin-top: -31px;
	margin-bottom: 26px;
	}
	.form {
	position: relative;
	z-index: 1;
	background: #FFFFFF;
	max-width: 360px;
	margin: 0 auto 100px;
	padding: 45px;
	text-align: center;
	box-shadow: 0 0 20px 0 rgba(0, 0, 0, 0.2), 0 5px 5px 0 rgba(0, 0, 0, 0.24);
	}
	.form input {
	font-family: "Roboto", sans-serif;
	outline: 0;
	background: #f2f2f2;
	width: 100%;
	border: 0;
	margin: 0 0 15px;
	padding: 15px;
	box-sizing: border-box;
	font-size: 14px;
	}
	.form button {
	font-family: "Roboto", sans-serif;
	text-transform: uppercase;
	outline: 0;
	background-color: #6B0000;
	background-image: linear-gradient(45deg,#6B0000,#A10101);
	width: 100%;
	border: 0;
	padding: 15px;
	color: #FFFFFF;
	font-size: 14px;
	-webkit-transition: all 0.3 ease;
	transition: all 0.3 ease;
	cursor: pointer;
	}
	.form .message {
	margin: 15px 0 0;
	color: #b3b3b3;
	font-size: 12px;
	}
	.form .message a {
	color: #4CAF50;
	text-decoration: none;
	}

	.container {
	position: relative;
	z-index: 1;
	max-width: 300px;
	margin: 0 auto;
	}

	body {
	background-color: #383232;
	background-image: linear-gradient(90deg,#383232,#8E8E8E);
	font-family: "Roboto", sans-serif;
	-webkit-font-smoothing: antialiased;
	-moz-osx-font-smoothing: grayscale;
	}
</style>
<title> Login </title>
</head>
<body>
  <body>
    <div class="login-page">
      <div class="form">
        <div class="login">
          <div class="login-header">
            <h3>LOGIN</h3>
            <p>Please enter your credentials to login.</p>
          </div>
        </div>
        <form method="post" action="/login" class="login-form">
          <input type="text" id="name" name="name" placeholder="username"/>
          <input type="password" id="password" name="password" placeholder="password"/>
          <button type="submit">login</button>
        </form>
      </div>
    </div>
</body>
</body>
</html>`
