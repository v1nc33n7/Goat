import {disableBtn, enableBtn} from './buttons.js';
import {joinRoom, createRoom, appendMessage} from './chat.js';

window.addEventListener('load', function() {
  const username = document.getElementById('username');
  const roomid = document.getElementById('roomid');
  const joinBtn = document.getElementById('join');
  const createBtn = document.getElementById('create');

  const chatPage = document.getElementById('chat');
  const loginPage = document.getElementById('login');

  const chatBox = document.getElementById('response');
  const sendInput = document.getElementById('message');
  const sendBtn = document.getElementById('send');
  const exitBtn = document.getElementById('exit');

  const error = document.getElementById('error');

  let ws;

  username.addEventListener('input', function() {
    if (username.value != '') {
      enableBtn(createBtn);
    } else {
      disableBtn(createBtn);
    }

    if (username.value != '' && roomid.value != '') {
      enableBtn(joinBtn);
    } else {
      disableBtn(joinBtn);
    }
  });

  roomid.addEventListener('input', function() {
    if (username.value != '' && roomid.value != '') {
      enableBtn(joinBtn);
    } else {
      disableBtn(joinBtn);
    }
  });

  sendInput.addEventListener('input', function() {
    if (sendInput.value != '') {
      enableBtn(sendBtn);
    } else {
      disableBtn(sendBtn);
    }
  });

  document.addEventListener('keypress', function(event) {
    if (event.key == 'Enter') {
      if (ws) {
        sendBtn.click();
      } else {
        if (username.value != '' && roomid.value == '') {
          createBtn.click();
        }
        if (username.value != '' && roomid.value != '') {
          joinBtn.click();
        }
      }
    }
  });

  const wsListeners = function() {
    ws.onopen = function() {
      error.innerHTML = '';

      chatPage.style.display = 'flex';
      loginPage.style.display = 'none';
      chatPage.style.marginTop = 0;

      appendMessage('Connected.', chatBox);

      sendInput.select();
    };

    ws.onclose = function() {
      loginPage.style.display = 'flex';
      chatPage.style.display = 'none';

      ws = null;
      chatBox.innerHTML = '';
    };

    ws.onmessage = function(msg) {
      if (msg.data == 'Couldn\'t find room with this id.') {
        error.innerHTML = 'Couldn\'t find room with this id.';

        return false;
      }
      appendMessage(msg.data, chatBox);
    };
  };

  joinBtn.onclick = function() {
    if (ws) {
      return false;
    }
    if (username.value == '' || roomid.value == '') {
      error.innerHTML = 'Please provide username and Room ID.';
      return false;
    }

    ws = joinRoom(ws);
    wsListeners();

    return false;
  };

  createBtn.onclick = function() {
    if (ws) {
      return false;
    }
    if (username.value == '') {
      error.innerHTML = 'Please provide username.';
      return false;
    }

    ws = createRoom(ws);
    wsListeners();

    return false;
  };

  sendBtn.onclick = function() {
    if (!ws) {
      return false;
    }
    ws.send(sendInput.value);

    sendInput.value = '';
    disableBtn(sendBtn);

    return false;
  };

  exitBtn.onclick = function() {
    if (!ws) {
      return false;
    }
    ws.close();
    error.innerHTML = 'Disconnected.';

    return false;
  };
});
