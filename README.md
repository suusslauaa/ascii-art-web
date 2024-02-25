<h1>Text to ASCII Art</h1>
A simple Go web-application to convert your text to ASCII Art.  
This web-application is able to read your text and convert it to each font mentioned below.
<h3>Here is some examples:</h3>
<div>
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
<div class="pre">
    <pre class="pre-style" cols="50" rows="5" id="output">__     __                               _                   _    <br>\ \   / /                              | |                 | |   <br> \ \_/ /    ___    _   _   _ __        | |_    ___  __  __ | |_  <br>  \   /    / _ \  | | | | | '__|       | __|  / _ \ \ \/ / | __| <br>   | |    | (_) | | |_| | | |          \ |_  |  __/  >  <  \ |_  <br>   |_|     \___/   \__,_| |_|           \__|  \___| /_/\_\  \__| <br>                                                                 <br>                                                                 <br></pre>
</div>
<footer class="footer">
    <h3>by suusslauaa</h3>
</footer>
<h2>Usage</h2>
<ol>
    <li>Clone the repository:</li>
    <code>$ git clone https://github.com/suusslauaa/text-to-ascii-art.git</code>
    <li>There is a command to run this web-application:</li>
    <code>$ go run ./web</code>
</ol>
<h2>Author</h2>
<li><a href="https://github.com/suusslauaa" target="_blank">suusslauaa</a></li>