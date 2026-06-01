// Состояние toggle-кнопок
let isSeries = null;       // true / false / null
let isHorizontal = null;   // true / false / null

function toggle(val) {
  isSeries = val === 'yes';
  document.getElementById('yes-btn').className = 'toggle-btn' + (isSeries ? ' active-yes' : '');
  document.getElementById('no-btn').className = 'toggle-btn' + (!isSeries ? ' active-no' : '');
}

function toggleH(val) {
  isHorizontal = val === 'yes';
  document.getElementById('h-yes-btn').className = 'toggle-btn' + (isHorizontal ? ' active-yes' : '');
  document.getElementById('h-no-btn').className = 'toggle-btn' + (!isHorizontal ? ' active-no' : '');
}

// Подсветка ошибки на поле
function setError(el, message) {
  el.style.borderColor = '#E24B4A';
  let hint = el.parentElement.querySelector('.error-hint');
  if (!hint) {
    hint = document.createElement('span');
    hint.className = 'error-hint';
    hint.style.cssText = 'display:block;font-size:12px;color:#f08080;margin-top:4px';
    el.parentElement.appendChild(hint);
  }
  hint.textContent = message;
}

// Сброс ошибки
function clearError(el) {
  el.style.borderColor = '';
  const hint = el.parentElement.querySelector('.error-hint');
  if (hint) hint.remove();
}

// Валидация всех полей
function validate() {
  let valid = true;

  const title = document.getElementById('title');
  const description = document.getElementById('description');
  const trailer = document.getElementById('trailer');
  const poster = document.getElementById('poster');
  const logo = document.getElementById('logo');

  // Сбрасываем старые ошибки
  [title, description, trailer, poster, logo].forEach(clearError);

  if (!title.value.trim()) {
    setError(title, 'Введите название проекта');
    valid = false;
  }

  if (!description.value.trim()) {
    setError(description, 'Введите описание');
    valid = false;
  }

  if (!trailer.value.trim()) {
    setError(trailer, 'Укажите путь к трейлеру');
    valid = false;
  }

  if (!poster.value.trim()) {
    setError(poster, 'Укажите путь к постеру');
    valid = false;
  }

  if (!logo.value.trim()) {
    setError(logo, 'Укажите путь к логотипу');
    valid = false;
  }

  if (isSeries === null) {
    alert('Укажите: это сериал или нет');
    valid = false;
  }

  if (isHorizontal === null) {
    alert('Укажите: постер горизонтальный или нет');
    valid = false;
  }

  return valid;
}

// Сбор данных
function getData() {
    t={
        path: document.getElementById('trailer').value.trim(),
    }
    c={
        path: document.getElementById('poster').value.trim(),
        is_horizontal: isHorizontal,
    }
    l={
        path: document.getElementById('logo').value.trim(),
    }
  return {
    title: document.getElementById('title').value.trim(),
    description: document.getElementById('description').value.trim(),
    is_serial: isSeries,
    trailer: t,
    card:c,
    logo:l,
  };
}

// Отправка
async function send() {
  if (!validate()) return;

  const data = getData();
  console.log('Данные для отправки:', data);

  try {
    const res = await fetch('/api/add', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(data),
    });

    if (!res.ok) throw new Error(`Ошибка сервера: ${res.status}`);

    const json = await res.json();
    console.log('Ответ сервера:', json);
    showToast('Данные отправлены');
    window.location.href="admin"
  } catch (err) {
    console.error(err);
    alert('Не удалось отправить данные: ' + err.message);
  }
}

// Toast
function showToast(msg) {
  const t = document.getElementById('toast');
  t.textContent = msg;
  t.className = 'toast show';
  setTimeout(() => t.className = 'toast', 2500);
}

function exit() {
  if (confirm('Выйти без сохранения?'))window.location.href="admin";
}