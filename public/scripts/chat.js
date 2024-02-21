window.addEventListener('load', function(event) {
  const responseDiv = document.getElementById('response');
  const loginPage = document.getElementById('login');
  const chatPage = document.getElementById('chat');

  const msgBox = document.getElementById('message');
  const sendButton = document.getElementById('send');

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

  const disableBtn = function(b) {
    b.style.color = 'rgb(82 82 82 / var(--tw-text-opacity))';
    b.style.borderColor = 'rgb(23 23 23 / var(--tw-border-opacity))';
    b.style.backgroundColor = 'rgb(38 38 38 / var(--tw-bg-opacity))';

    b.style.cursor = 'default';
  };

  const enableBtn = function(b) {
    b.style.color = 'rgb(255 255 255 / var(--tw-text-opacity))';
    b.style.borderColor = 'rgb(127 29 29 / var(--tw-border-opacity))';
    b.style.backgroundColor = 'rgb(153 27 27 / var(--tw-bg-opacity))';

    b.style.cursor = 'pointer';
  };

  const hoverBtn = function(b) {
    b.style.backgroundColor = 'rgb(220 38 38 / var(--tw-bg-opacity))';
  };

  const unhoverBtn = function(b) {
    b.style.backgroundColor = 'rgb(153 27 27 / var(--tw-bg-opacity))';
  };

  msgBox
      .addEventListener('keypress', function(event) {
        if (event.key === 'Enter') {
          if (msgBox.value == '') {
            return false;
          }

          sendButton.click();
        }
      });

  msgBox
      .addEventListener('input', function(event) {
        if (msgBox.value != '') {
          enableBtn(sendButton);
        } else {
          disableBtn(sendButton);
        }
      });

  sendButton
      .addEventListener('mouseover', function(event) {
        if (msgBox.value != '') {
          hoverBtn(sendButton);
        }
      });

  sendButton
      .addEventListener('mouseleave', function(event) {
        if (msgBox.value != '') {
          unhoverBtn(sendButton);
        }
      });

  sendButton.onclick = function(event) {
    if (!ws) {
      return false;
    }
    ws.send(msgBox.value);

    msgBox.value = '';
    disableBtn(sendButton);

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
