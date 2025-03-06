const signUp = () => {
  const name = document.getElementById("name").value
  const email = document.getElementById("email").value
  const password = document.getElementById("password").value
  const passwordCheck = document.getElementById("password-confirmation").value
  const image = document.getElementById("image").files[0]

  if (password !== passwordCheck) {
    return
  }

  const json = {
    name: name,
    email: email,
    password: password
  }

  const form = new FormData();

  form.set("json", JSON.stringify(json))
  form.set("profile-image", image)

  fetch("/users/sign-up", {
    cors: "no-cors",
    method: "POST",
    body: form
  }).then((resp) => {
    if (resp.ok) {
      window.location = "/home"
    } else {
      return resp.json()
    }
  }).then((json) => {
    console.log(json)
  }).catch((error) => {
    console.log(error)
  })
}
