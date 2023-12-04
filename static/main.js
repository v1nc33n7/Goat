window.addEventListener("load", function (event) {
  var blocker = document.getElementById("blocker");
  var login = document.getElementById("login");

  var output = document.getElementById("output");
  var input = document.getElementById("input");
  var ws;

  var response = function (message) {
    var d = document.createElement("div");
    d.textContent = message;
    output.appendChild(d);
    output.scroll(0, output.scrollHeight);
  };

  document.getElementById("create").onclick = function (event) {
    if (ws) {
      return false;
    }

    ws = new WebSocket(
      "wss://" + window.location.host + "/create?username=" + username.value,
    );
    ws.onopen = function (event) {
      login.style.visibility = "hidden";
      blocker.style.visibility = "hidden";
      response("Connecting...");
    };
    ws.onclose = function (event) {
      response("Connection closed.");
      login.style.visibility = "";
      blocker.style.visibility = "";
      ws = null;
    };
    ws.onmessage = function (event) {
      response(event.data);
    };
    return false;
  };

  document.getElementById("open").onclick = function (event) {
    if (ws) {
      return false;
    }

    ws = new WebSocket(
      "wss://" +
        window.location.host +
        "/add?roomid=" +
        roomid.value +
        "&username=" +
        username.value,
    );
    ws.onopen = function (event) {
      login.style.visibility = "hidden";
      blocker.style.visibility = "hidden";
      response("Connecting...");
    };
    ws.onclose = function (event) {
      response("Connection closed.");
      login.style.visibility = "";
      blocker.style.visibility = "";
      ws = null;
    };
    ws.onmessage = function (event) {
      response(event.data);
    };
    return false;
  };

  document.getElementById("send").onclick = function (event) {
    if (!ws) {
      return false;
    }
    ws.send(input.value);
    return false;
  };
});
