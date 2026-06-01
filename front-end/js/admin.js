async function loadData() {
  try {
    const res = await fetch('/api/films');
    if (!res.ok) throw new Error(res.status);
    const films = await res.json();
    renderTable(films);
  } catch (err) {
    console.error('Ошибка загрузки:', err);
  }
}

function renderTable(films) {
  const tbody = document.getElementById('tbody');

  if (!films.length) {
    tbody.innerHTML = `<tr><td colspan="9" class="empty">Нет данных</td></tr>`;
    return;
  }

  console.log(films)

  tbody.innerHTML = films.map(f => `
    <tr>
      <td class="id">#${f.id}</td>
      <td style="font-weight:500">${f.title}</td>
      <td><span class="badge ${f.is_serial? 'badge-yes' : 'badge-no'}">${f.is_serial ? 'Да' : 'Нет'}</span></td>
      <td class="desc-cell">${f.description}</td>
      <td class="path" title="${f.trailer ? f.trailer.path: "—"}">${f.trailer ? f.trailer.path: "—"}</td>
      <td class="path" title="${f.card ? f.card.path: "—"}">${f.card ? f.card.path: "—"}</td>
      <td><span class="badge ${f.card && f.card.is_horizontal ? 'badge-yes' : 'badge-no'}">${f.card && f.card.is_horizontal ? 'Да' : 'Нет'}</span></td>
      <td class="path" title="${f.logo ? f.logo.path: "—"}">${f.logo ? f.logo.path: "—"}</td>
      <td><button class="del-btn" onclick="deleteFilm(${f.id})">Удалить</button></td>
    </tr>
  `).join('');
}

// async function deleteFilm(id) {
//   if (!confirm(`Удалить проект #${id}?`)) return;
//   try {
//     const res = await fetch(`/api/films/${id}`, { method: 'DELETE' });
//     if (!res.ok) throw new Error(res.status);
//     loadData(); // перезагружаем таблицу
//   } catch (err) {
//     alert('Ошибка удаления: ' + err.message);
//   }
// }

loadData();