window.onload = () => {
  const title = document.getElementById("title");
  let smiling = true;

  setInterval(() => {
    if (smiling) {
      title.textContent = "Hello World";
    } else {
      title.textContent = "Hello World 😊";
    }
    smiling = !smiling;
  }, 1000);
};
