package webroot

func init() {
	webFiles["index.css"] = indexCSS
}

const indexCSS = `
body {
	font-family: "Lato", sans-serif;
	background-color:#1c1c1c;
	color:#eaeaea;
}
ul.tab {
	list-style-type: none;
	margin: 0;
	padding: 0;
	overflow: hidden;
	border: 1px solid #ccc;
	background-color:#1c1c1c;
	color:#eaeaea;
}

/* Float the list items side by side */
ul.tab li {float: left;}

/* Style the links inside the list items */
ul.tab li a {
	display: inline-block;
	text-align: center;
	padding: 14px 16px;
	text-decoration: none;
	transition: 0.3s;
	font-size: 17px;

	color:#eaeaea;
	background-color:#1c1c1c;
}

/* Change background color of links on hover */
ul.tab li a:hover {
	background-color:#282828;
	color:#eaeaea;
}

/* Create an active/current tablink class */
ul.tab li a:focus, .active {
	background-color:#232323;
	color:#eaeaea;
}

/* Style the tab content */
.tabcontent {
	display: none;
	padding: 6px 12px;
	border: 1px solid #ccc;
	border-top: none;
}

textarea {
	width:100%;

	color:#eaeaea;
	background-color:#1c1c1c;
}

/*table {
/*	width:100%;
/*	overflow:scroll;
/*	border:0;
/*}*/

button {
    background-color: #1c1c1c; 
	 border: 2px solid #eaeaea; 
    color: #eaeaea;
    padding: 15px 32px;
    text-align: center;
    text-decoration: none;
    display: inline-block;
}
`
