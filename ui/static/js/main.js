function text (input, output) {
  var txtArea = document.getElementById(input);

  txtArea.addEventListener('input', function () {
    var div = document.getElementById(output);
    div.innerHTML = txtArea.value;
  })
}

text('input','output');
