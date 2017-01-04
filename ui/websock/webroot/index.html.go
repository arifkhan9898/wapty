package webroot

func init() {
	webFiles[""] = indexHtml
}

const indexHtml = `
<HTML>
	<HEAD>
		<link rel="stylesheet" type="text/css" href="/index.css">
		<link rel="stylesheet" type="text/css" href="/colresizable.css">
		<!--TODO embed minified version of these too, using it live just for tests-->
		<script src="https://ajax.googleapis.com/ajax/libs/jquery/3.1.0/jquery.js"></script>
		<script src="/tablesorter.js"></script>
		<script src="/colResizable.js"></script>
	</HEAD>
	<BODY>
		<ul class="tab">
			<li><a href="javascript:void(0)" class="tablinks" onclick="openTab(event, 'proxyTab')" id="defaultOpen">Proxy</a></li>
			<li><a href="javascript:void(0)" class="tablinks" onclick="openTab(event, 'historyTab')">History</a></li>
		</ul>
		<div id="proxyTab" class="tabcontent">
			<button id="forwardOriginal" type="button" onclick="clickhandler();">Forward Original</button>
			<button id="forwardModified" type="button" onclick="clickhandler();">Forward Modified</button>
			<button id="drop" type="button" onclick="clickhandler();">Drop</button>
			<button id="provideResponse" type="button" onclick="clickhandler();">Provide Response</button>
			<label style="border: 2px solid #eaeaea; padding: 15px 32px;"><input type="checkbox" id="interceptToggle" value="Intercept" onclick="toggler();">Intercept</label>
			<button id="menu" type="button">Action</button>
			<br>
			<div id="endpointIndicator"></div>
			<br>

			<textarea id="proxybuffer" name="proxybuffer" rows="40"></textarea>
		</div>

		<div id="historyTab" class="tabcontent">
			<table id="historyTable" border="0" cellpadding="0" cellspacing="0" class="JPadding JColResizer overflowContainer">
				<tr id="historyHeader">
				</tr>
			</table>
		</div>

		<SCRIPT src="/index.js"></SCRIPT>
		<script>
function openTab(evt, tabName) {
	var i, tabcontent, tablinks;
	tabcontent = document.getElementsByClassName("tabcontent");
	for (i = 0; i < tabcontent.length; i++) {
		tabcontent[i].style.display = "none";
	}
	tablinks = document.getElementsByClassName("tablinks");
	for (i = 0; i < tablinks.length; i++) {
		tablinks[i].className = tablinks[i].className.replace(" active", "");
	}
	document.getElementById(tabName).style.display = "block";
	evt.currentTarget.className += " active";
}

// Get the element with id="defaultOpen" and click on it
document.getElementById("defaultOpen").click();

		</script>

	</BODY>
</HTML>
`
