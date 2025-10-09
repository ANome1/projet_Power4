(function () {
  const overlay = document.getElementById('diffPop');
  function close() {
    if (overlay) overlay.style.display = 'none';
  }

  if (overlay) {
    overlay.addEventListener('click', (e) => {
      if (e.target === overlay) close();
    });
  }

  document.addEventListener('keydown', (e) => {
    if (e.key === 'Escape') close();
  });

  const first = document.querySelector('.difficulty-button');
  if (first) first.focus();
})();