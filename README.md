<!DOCTYPE html>
<font color="black">
    <h1>Text to ASCII Art</h1>
    A simple Go web-application to convert your text to ASCII Art.  
    This web-application is able to read your text and convert it to each font mentioned below.
    <h3>Here is some examples:</h3>
</font>
<head>
    <link rel="stylesheet" href="ui/static/css/home.css">
    <title>Text to ASCII Art</title>
</head>
<body>
    <header class="header">
        <div class="header_inner">
            <a href="/">
                <h1 class="header_title">Text to ASCII Art</h1>
            </a>
        </div>  
    </header>
    <main class="main">
        <div class="content">
            <form id="text" action="/text-to-ascii-art" method="POST">
                <div>
                    <!-- mb later I'll add a placeholder="" maxlength="" autofocus name="" -->
                    <textarea class="input" autofocus name="input" cols="50" rows="5" id="input">Your text</textarea>
                </div>
                <div class="settings">
                    <select class="fonts" id="font" name="font">
                        <option value="program/banners/standard.txt">Standard</option>
                        <option value="program/banners/shadow.txt">Shadow</option>
                        <option value="program/banners/thinkertoy.txt">Thinkertoy</option>
                    </select>
                </div>
                <div>
                    <button class="button">Convert</button>
                </div>
            </form>
            <div class="pre">
                <pre class="pre-style" cols="50" rows="5" id="output">__     __                               _                   _    <br>\ \   / /                              | |                 | |   <br> \ \_/ /    ___    _   _   _ __        | |_    ___  __  __ | |_  <br>  \   /    / _ \  | | | | | '__|       | __|  / _ \ \ \/ / | __| <br>   | |    | (_) | | |_| | | |          \ |_  |  __/  >  <  \ |_  <br>   |_|     \___/   \__,_| |_|           \__|  \___| /_/\_\  \__| <br>                                                                 <br>                                                                 <br></pre>
            </div>
        </div>
    </main>
    <footer class="footer">
        <h3>by suusslauaa</h3>
    </footer>
    <!-- <script src="../static/js/main.js" type="text/javascript"></script> -->
</body>
<font color="black">
    <h2>Usage</h2>
    <ol>
        <li>Clone the repository:</li>
        <code><p style="background-color:black;color:white;">$ git clone https://github.com/suusslauaa/text-to-ascii-art.git</p></code>
        <li>There is a command to run this web-application:</li>
        <code><p style="background-color:black;color:white;">$ go run ./web</p></code>
    </ol>
    <h2>Author</h2>
    <li><a href="https://github.com/suusslauaa" target="_blank">suusslauaa</a></li>
</font>