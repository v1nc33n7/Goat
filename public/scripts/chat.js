window.addEventListener('load', function(event) {
  const responseDiv = document.getElementById('response');
  const loginPage = document.getElementById('login');
  const chatPage = document.getElementById('chat');

  const response = function(message) {
    const d = document.createElement('div');
    d.textContent = message;
    responseDiv.appendChild(d);
    responseDiv.scroll(0, responseDiv.scrollHeight);
  };

  let ws;
  const loadWs = function(websocket) {
    websocket.onopen = function(event) {
      chatPage.style.display = 'flex';
      chatPage.style.marginTop = 0;
      loginPage.style.display = 'none';

      response('Connected.');
    };

    websocket.onclose = function(event) {
      loginPage.style.display = 'flex';
      chatPage.style.display = 'none';

      ws = null;
    };

    websocket.onmessage = function(event) {
      response(event.data);
    };
  };

  document.getElementById('create').onclick = function(event) {
    if (ws) {
      return false;
    }
    ws = new WebSocket(
        'ws://' +
            window.location.host +
            '/create?username=' +
            document.getElementById('username').value,
    );

    loadWs(ws);

    return false;
  };

  document.getElementById('add').onclick = function(event) {
    if (ws) {
      return false;
    }
    ws = new WebSocket(
        'ws://' +
            window.location.host +
            '/add?username=' +
            document.getElementById('username').value +
            '&roomid=' +
            document.getElementById('roomid').value,
    );

    loadWs(ws);

    return false;
  };

  document
      .getElementById('message')
      .addEventListener('keypress', function(event) {
        if (event.key === 'Enter') {
          document.getElementById('send').click();
        }
      });

  document.getElementById('send').onclick = function(event) {
    if (!ws) {
      return false;
    }
    ws.send(document.getElementById('message').value);

    document.getElementById('message').value = '';

    return false;
  };

  document.getElementById('exit').onclick = function(event) {
    if (!ws) {
      return false;
    }
    ws.close();
    ws = null;

    responseDiv.innerHTML = '';

    return false;
  };
});
