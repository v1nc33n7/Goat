export function enableBtn(b) {
  b.style.color = 'rgb(255 255 255 / var(--tw-text-opacity))';
  b.style.borderColor = 'rgb(127 29 29 / var(--tw-border-opacity))';
  b.style.backgroundColor = 'rgb(153 27 27 / var(--tw-bg-opacity))';

  b.style.cursor = 'pointer';

  b.addEventListener('mouseover', function() {
    b.style.backgroundColor = 'rgb(220 38 38 / var(--tw-bg-opacity))';
  });

  b.addEventListener('mouseleave', function() {
    b.style.backgroundColor = 'rgb(153 27 27 / var(--tw-bg-opacity))';
  });
}

export function disableBtn(b) {
  b.style.color = 'rgb(82 82 82 / var(--tw-text-opacity))';
  b.style.borderColor = 'rgb(23 23 23 / var(--tw-border-opacity))';
  b.style.backgroundColor = 'rgb(38 38 38 / var(--tw-bg-opacity))';

  b.style.cursor = 'default';

  b.addEventListener('mouseover', function() {
    b.style.backgroundColor = 'rgb(38 38 38 / var(--tw-bg-opacity))';
  });

  b.addEventListener('mouseleave', function() {
    b.style.backgroundColor = 'rgb(38 38 38 / var(--tw-bg-opacity))';
  });
}
