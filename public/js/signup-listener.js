const form = document.getElementById("form")

form.addEventListener("submit", (event) => {
  event.preventDefault()
})

document.getElementById('image').addEventListener('change', (event) => {
  const file = event.target.files[0];

  if (file) {
    const reader = new FileReader();

    reader.onloadend = function() {
      const imagePreview = document.getElementById('imagePreview');
      const plusSiign = document.getElementById("plus-sign")

      imagePreview.src = reader.result;
      imagePreview.classList.remove("hidden")
      previewContainer.classList.remove('hidden');
      plusSiign.classList.add("hidden");
    };

    reader.readAsDataURL(file);
  }
});

