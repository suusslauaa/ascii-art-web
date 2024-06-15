function text (input, output) {
  var textArea = document.getElementById(input);

  textArea.addEventListener('input', function () {
    var div = document.getElementById(output);
    div.innerHTML = textArea.value;
  })
}

text('input','output');
